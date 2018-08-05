// Code generated by "kmipenums "; DO NOT EDIT.

package kmip

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

// CredentialType
// 9.1.3.2.1 Credential Type Enumeration
type CredentialType uint32

const (
	CredentialTypeAttestation         CredentialType = 0x00000003
	CredentialTypeDevice              CredentialType = 0x00000002
	CredentialTypeUsernameAndPassword CredentialType = 0x00000001
)

var _CredentialTypeNameToValueMap = map[string]CredentialType{
	"Attestation":         CredentialTypeAttestation,
	"Device":              CredentialTypeDevice,
	"UsernameAndPassword": CredentialTypeUsernameAndPassword,
}

var _CredentialTypeValueToNameMap = map[CredentialType]string{
	CredentialTypeAttestation:         "Attestation",
	CredentialTypeDevice:              "Device",
	CredentialTypeUsernameAndPassword: "UsernameAndPassword",
}

func (c CredentialType) String() string {
	if s, ok := _CredentialTypeValueToNameMap[c]; ok {
		return s
	}
	return fmt.Sprintf("%#08x", c)
}

func ParseCredentialType(s string) (CredentialType, error) {
	if strings.HasPrefix(s, "0x") && len(s) == 10 {
		b, err := hex.DecodeString(s[2:])
		if err != nil {
			return 0, err
		}
		return CredentialType(binary.BigEndian.Uint32(b)), nil
	}
	if v, ok := _CredentialTypeNameToValueMap[s]; ok {
		return v, nil
	} else {
		var v CredentialType
		return v, fmt.Errorf("%s is not a valid CredentialType", s)
	}
}

func (c CredentialType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CredentialType) UnmarshalText(text []byte) (err error) {
	*c, err = ParseCredentialType(string(text))
	return
}

func (c CredentialType) EnumValue() uint32 {
	return uint32(c)
}

func (c CredentialType) MarshalTTLVEnum() uint32 {
	return uint32(c)
}
