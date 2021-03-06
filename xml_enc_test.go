package xml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const TEST_SAML_FILE = "test/upitt_saml.xml"

const DECRYPTION_KEY string = `-----BEGIN PRIVATE KEY-----
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
yRVhYDIfRH/lgvQwjj6E3ADiihT9wA==
-----END PRIVATE KEY-----`

func TestDecrypt(t *testing.T) {
	doc := ParseEntity(TEST_SAML_FILE)
	node := doc.Root().Children()
	for node != nil {
		if node.Name() == "EncryptedAssertion" {
			err := Decrypt(node, []byte(DECRYPTION_KEY))
			assert.Nil(t, err)
			return
		}
		node = node.NextSibling()
	}
	t.Fatal("EncryptedAssertion node wasn't found")
}

func TestDecryptWithEmptyKey(t *testing.T) {
	doc := ParseEntity(TEST_SAML_FILE)
	node := doc.Root().Children()
	for node != nil {
		if node.Name() == "EncryptedAssertion" {
			err := Decrypt(node, []byte(""))
			assert.Equal(t, err.Error(), "cannot decrypt XML due to empty decryption key")
			return
		}
		node = node.NextSibling()
	}
	t.Fatal("EncryptedAssertion node wasn't found")
}
