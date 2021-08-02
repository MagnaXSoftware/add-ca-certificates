package management

import (
	"encoding/pem"
	"io/fs"
	"os"
	"strings"
	"sync"
	"syscall"
)

type Manager struct {
	CertBundlePath string
	LocalPath      string

	certs *certificateSet
}

func NewManager(certBundle, localPath string) *Manager {
	manager := &Manager{
		CertBundlePath: certBundle,
		LocalPath:      localPath,

		certs: NewCertificateSet(),
	}

	return manager
}

func (m *Manager) BuildBundle() error {
	sys := &fileSystem{}
	err := m.parseCertBundle(sys)
	if err != nil {
		return err
	}

	// We use the specific Sub implementation because the fs.Sub one rejects rooted paths, which we support.
	sub, err := sys.Sub(m.LocalPath)
	if err != nil {
		return err
	}
	return m.parseLocalPath(sub)
}

func (m *Manager) WriteBundle() (int64, error) {
	outFile, err := os.OpenFile(m.CertBundlePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	return m.certs.WriteTo(outFile)
}

func (m *Manager) extractCertificates(data []byte) {
	var wg sync.WaitGroup
	for {
		block, newData := pem.Decode(data)
		if block == nil {
			break
		}

		if block.Type != "CERTIFICATE" {
			continue
		}

		wg.Add(1)
		go func(p *pem.Block) {
			defer wg.Done()
			m.certs.AddCertificate(p)
		}(block)

		data = newData
	}

	wg.Wait()
}

func (m *Manager) parseCertBundle(sys fs.FS) error {
	data, err := fs.ReadFile(sys, m.CertBundlePath)
	if err != nil {
		if ferr, ok := err.(*fs.PathError); ok && ferr.Err == syscall.Errno(2) {
			return nil
		} else if err == fs.ErrNotExist {
			return nil
		}
		return err
	}

	m.extractCertificates(data)

	return nil
}

func (m *Manager) parseLocalPath(sys fs.FS) error {
	dirEntries, err := fs.ReadDir(sys, ".")
	if err != nil {
		if ferr, ok := err.(*fs.PathError); ok && ferr.Err == syscall.Errno(2) {
			return nil
		} else if err == fs.ErrNotExist {
			return nil
		}
		return err
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			sub, err := fs.Sub(sys, dirEntry.Name())
			if err != nil {
				return err
			}
			err = m.parseLocalPath(sub)
			if err != nil {
				return err
			}
			continue
		}

		// It's a file!
		if !strings.HasSuffix(dirEntry.Name(), ".crt") {
			continue
		}

		data, err := fs.ReadFile(sys, dirEntry.Name())
		if err != nil {
			return err
		}
		m.extractCertificates(data)
	}

	return nil
}
