package kmip

import (
	"github.com/gemalto/kmip-go/ttlv"
	"math/big"
)

// 2.2

// 2.2.1

type Certificate struct {
	CertificateType  ttlv.CertificateType
	CertificateValue []byte
}

// 2.2.2

type SymmetricKey struct {
	KeyBlock KeyBlock
}

// 2.2.3

type PublicKey struct {
	KeyBlock KeyBlock
}

// 2.2.4

type PrivateKey struct {
	KeyBlock KeyBlock
}

// 2.2.5

type SplitKey struct {
	SplitKeyParts     int
	KeyPartIdentifier int
	SplitKeyThreshold int
	SplitKeyMethod    ttlv.SplitKeyMethod
	PrimeFieldSize    *big.Int `ttlv:",omitempty"`
	KeyBlock          KeyBlock
}

// 2.2.6

type Template struct {
	Attribute []Attribute
}

// 2.2.7

type SecretData struct {
	SecretDataType ttlv.SecretDataType
	KeyBlock       KeyBlock
}

// 2.2.8

type OpaqueObject struct {
	OpaqueDataType  ttlv.OpaqueDataType
	OpaqueDataValue []byte
}

// 2.2.9

type PGPKey struct {
	PGPKeyVersion int
	KeyBlock      KeyBlock
}
