package xml

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"
)

func TestDigitalSignature(t *testing.T) {
	keyName := "magic"
	commonName := "Johnny"
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()

	certTmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"Acme Co"},
		},
		NotBefore:    now.Add(-5 * time.Minute).UTC(),
		NotAfter:     now.AddDate(1, 0, 0).UTC(),
		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &certTmpl, &certTmpl,
		&privKey.PublicKey, privKey)
	if err != nil {
		t.Fatal(err)
	}
	certPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certBytes,
		},
	)
	privPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privKey),
		},
	)
	doc := NewDoc("1.0")
	defer doc.Free()
	node := NewNode(nil, "blackbox")
	node.SetContent("magic")
	doc.AddChild(node)
	if err := DigitallySign(doc, node, keyName, privPem, certPem); err != nil {
		t.Fatal(err)
	}
	if signed, err := VerifySignature(node, keyName, certPem); err != nil {
		t.Fatal(err)
	} else if !signed {
		t.Errorf("expected verify to be true")
	}
}
