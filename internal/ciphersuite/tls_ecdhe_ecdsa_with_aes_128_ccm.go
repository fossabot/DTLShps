package ciphersuite

import (
	"github.com/zjw1111/DTLShps/pkg/crypto/ciphersuite"
	"github.com/zjw1111/DTLShps/pkg/crypto/clientcertificate"
)

// NewTLSEcdheEcdsaWithAes128Ccm constructs a TLS_ECDHE_ECDSA_WITH_AES_128_CCM Cipher
func NewTLSEcdheEcdsaWithAes128Ccm() *Aes128Ccm {
	return newAes128Ccm(clientcertificate.ECDSASign, TLS_ECDHE_ECDSA_WITH_AES_128_CCM, false, ciphersuite.CCMTagLength)
}
