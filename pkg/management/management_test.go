package management

import "testing"

const mixedFile = `-----BEGIN CERTIFICATE-----
AAECAwQFBgc=
-----END CERTIFICATE-----
-----BEGIN PRIVATE KEY-----
AAECAwQFBgc=
-----END PRIVATE KEY-----
-----BEGIN PUBLIC KEY-----
AAECAwQFBgc=
-----END PUBLIC KEY-----
-----BEGIN CERTIFICATE REQUEST-----
AAECAwQFBgc=
-----END CERTIFICATE REQUEST-----
-----BEGIN CERTIFICATE-----
VGhpcyBpcyB2YWxpZCBkYXRh
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
AAECAwQFBgc=
-----END CERTIFICATE-----
`

func TestManager_ExtractCertificates(t *testing.T) {
	m := NewManager("", []string{""})
	data := []byte("")

	m.extractCertificates(data)
	if l := m.certs.Len(); l != 0 {
		t.Errorf("Expected no certificates, got %d", l)
	}

	m.extractCertificates([]byte(mixedFile))
	if l := m.certs.Len(); l != 2 {
		t.Errorf("Expected 2 certificates, got %d", l)
	}
}
