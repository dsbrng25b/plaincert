// Code generated by go generate; DO NOT EDIT.
package pcert

import (
	"crypto/x509"
)

var KeyUsage = map[string]x509.KeyUsage{
	"CRLSign":           x509.KeyUsageCRLSign,
	"CertSign":          x509.KeyUsageCertSign,
	"ContentCommitment": x509.KeyUsageContentCommitment,
	"DataEncipherment":  x509.KeyUsageDataEncipherment,
	"DecipherOnly":      x509.KeyUsageDecipherOnly,
	"DigitalSignature":  x509.KeyUsageDigitalSignature,
	"EncipherOnly":      x509.KeyUsageEncipherOnly,
	"KeyAgreement":      x509.KeyUsageKeyAgreement,
	"KeyEncipherment":   x509.KeyUsageKeyEncipherment,
}

var ExtKeyUsage = map[string]x509.ExtKeyUsage{
	"Any":                            x509.ExtKeyUsageAny,
	"ClientAuth":                     x509.ExtKeyUsageClientAuth,
	"CodeSigning":                    x509.ExtKeyUsageCodeSigning,
	"EmailProtection":                x509.ExtKeyUsageEmailProtection,
	"IPSECEndSystem":                 x509.ExtKeyUsageIPSECEndSystem,
	"IPSECTunnel":                    x509.ExtKeyUsageIPSECTunnel,
	"IPSECUser":                      x509.ExtKeyUsageIPSECUser,
	"MicrosoftCommercialCodeSigning": x509.ExtKeyUsageMicrosoftCommercialCodeSigning,
	"MicrosoftKernelCodeSigning":     x509.ExtKeyUsageMicrosoftKernelCodeSigning,
	"MicrosoftServerGatedCrypto":     x509.ExtKeyUsageMicrosoftServerGatedCrypto,
	"NetscapeServerGatedCrypto":      x509.ExtKeyUsageNetscapeServerGatedCrypto,
	"OCSPSigning":                    x509.ExtKeyUsageOCSPSigning,
	"ServerAuth":                     x509.ExtKeyUsageServerAuth,
	"TimeStamping":                   x509.ExtKeyUsageTimeStamping,
}