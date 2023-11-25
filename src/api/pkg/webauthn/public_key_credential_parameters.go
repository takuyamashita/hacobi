package webauthn

const (
	PublicKeyCredentialTypePublicKey PublicKeyCredentialType = "public-key"

	// AlgES256 ECDSA with SHA-256.
	AlgES256 COSEAlgorithmIdentifier = -7

	// AlgES384 ECDSA with SHA-384.
	AlgES384 COSEAlgorithmIdentifier = -35

	// AlgES512 ECDSA with SHA-512.
	AlgES512 COSEAlgorithmIdentifier = -36

	// AlgRS1 RSASSA-PKCS1-v1_5 with SHA-1.
	AlgRS1 COSEAlgorithmIdentifier = -65535

	// AlgRS256 RSASSA-PKCS1-v1_5 with SHA-256.
	AlgRS256 COSEAlgorithmIdentifier = -257

	// AlgRS384 RSASSA-PKCS1-v1_5 with SHA-384.
	AlgRS384 COSEAlgorithmIdentifier = -258

	// AlgRS512 RSASSA-PKCS1-v1_5 with SHA-512.
	AlgRS512 COSEAlgorithmIdentifier = -259

	// AlgPS256 RSASSA-PSS with SHA-256.
	AlgPS256 COSEAlgorithmIdentifier = -37

	// AlgPS384 RSASSA-PSS with SHA-384.
	AlgPS384 COSEAlgorithmIdentifier = -38

	// AlgPS512 RSASSA-PSS with SHA-512.
	AlgPS512 COSEAlgorithmIdentifier = -39

	// AlgEdDSA EdDSA.
	AlgEdDSA COSEAlgorithmIdentifier = -8

	// AlgES256K is ECDSA using secp256k1 curve and SHA-256.
	AlgES256K COSEAlgorithmIdentifier = -47
)

type PublicKeyCredentialType string

// https://www.iana.org/assignments/cose/cose.xhtml#algorithms
type COSEAlgorithmIdentifier int

type PublicKeyCredentialParameters struct {
	Type PublicKeyCredentialType `json:"type"`
	Alg  COSEAlgorithmIdentifier `json:"alg"`
}
