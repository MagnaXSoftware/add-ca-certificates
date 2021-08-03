package management

import (
	"bytes"
	"encoding/pem"
	"testing"
)

func TestCertificateSet_AddCertificate(t *testing.T) {
	set := NewCertificateSet()
	if l := set.Len(); l != 0 {
		t.Errorf("Expected new certificateSet to be empty, got %d", l)
	}

	var block certificate = &pem.Block{
		Type:    "CERTIFICATE",
		Headers: make(map[string]string),
		Bytes:   []byte{0, 1, 2, 3, 4, 5, 6, 7},
	}

	set.AddCertificate(block)
	if l := set.Len(); l != 1 {
		t.Errorf("Expected certificateSet to have one entry, got %d", l)
	}

	set.AddCertificate(block)
	if l := set.Len(); l != 1 {
		t.Errorf("Expected certificateSet to have one entry, got %d", l)
	}
}

func TestCertificateSet_WriteTo(t *testing.T) {
	set := NewCertificateSet()

	var block certificate = &pem.Block{
		Type:    "CERTIFICATE",
		Headers: make(map[string]string),
		Bytes:   []byte{0, 1, 2, 3, 4, 5, 6, 7},
	}
	set.AddCertificate(block)

	buf := &bytes.Buffer{}

	n, err := set.WriteTo(buf)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("Expected 1 certificate to be printed, got %d", n)
	}

	const expected = "-----BEGIN CERTIFICATE-----\nAAECAwQFBgc=\n-----END CERTIFICATE-----\n\n"
	if actual := string(buf.Bytes()); expected != actual {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
