package node

import "errors"

type CertType string

const (
	CertTypePrivateFullNode = "private_full_node"
	CertTypePrivateWallet   = "private_wallet"
)

func GetCertTypeFromString(name string) (CertType, error) {
	switch name {
	case CertTypePrivateFullNode:
		return CertTypePrivateFullNode, nil
	case CertTypePrivateWallet:
		return CertTypePrivateWallet, nil
	default:
		return "", errors.New("unknown cert type")
	}
}
