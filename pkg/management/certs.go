package management

import (
	"bytes"
	"encoding/pem"
	"io"
	"sync"
)

const (
	mapSetSize = 256
)

// hashCert creates a hash of the DER bytes for a given pem Block.
// Inspired by alpine's update-ca-certificate hash_string.
func hashCert(block certificate) uint8 {
	var hash uint64 = 5381
	for _, b := range block.Bytes {
		hash = (hash << 5) + hash + uint64(b)
	}
	return uint8(hash % mapSetSize)
}

type certificateSet struct {
	certMap   [mapSetSize]certificateList
	certOrder []certIdx

	mutex sync.Mutex
}

type certificateList []certificate
type certificate *pem.Block
type certIdx struct {
	hashCode uint8
	idx      int
}

func NewCertificateSet() *certificateSet {
	return &certificateSet{}
}

// Len calculates and returns the size of the certificate set.
func (s *certificateSet) Len() (n int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	n = len(s.certOrder)
	return
}

// AddCertificate stores the given certificate block in the set.
//
// Because of set semantics, only one copy of a given certificate is stored.
func (s *certificateSet) AddCertificate(block certificate) {
	hash := hashCert(block)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.certMap[hash] != nil {
		for _, p := range s.certMap[hash] {
			if len(p.Bytes) == len(block.Bytes) && bytes.Equal(p.Bytes, block.Bytes) {
				return
			}
		}
	}
	s.certMap[hash] = append(s.certMap[hash], block)
	s.certOrder = append(s.certOrder, certIdx{hash, len(s.certMap[hash]) - 1})
}

func (s *certificateSet) WriteTo(w io.Writer) (n int64, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, idx := range s.certOrder {
		err = pem.Encode(w, s.certMap[idx.hashCode][idx.idx])
		if err != nil {
			return
		}
		n++
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return
		}
	}

	return
}

var _ io.WriterTo = &certificateSet{}
