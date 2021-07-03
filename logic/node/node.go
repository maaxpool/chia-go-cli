package node

import (
	"crypto/tls"
	"fmt"
	"os"
)

type Node struct {
	m *Manage

	nodeData *nodeData
}

func newNode(manage *Manage, name string) (*Node, error) {
	node := &Node{
		m: manage,
	}

	err := node.open(name)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (n *Node) open(nodeName string) error {
	data := &nodeData{
		NodeName: nodeName,
	}

	err := data.load(n.m)
	if err != nil {
		return err
	}

	n.nodeData = data
	return nil
}

func (n *Node) Name() string {
	return n.nodeData.NodeName
}

func (n *Node) NodeUrl() string {
	return n.nodeData.NodeUrl
}

func (n *Node) StoredCerts() []string {
	retList := make([]string, 0, len(n.nodeData.HasCerts))
	for certName, has := range n.nodeData.HasCerts {
		if has {
			retList = append(retList, certName)
		}
	}
	return retList
}

func (n *Node) AddCert(typ CertType, certFile string, keyFile string) error {
	certPEMBlock, err := os.ReadFile(certFile)
	if err != nil {
		return err
	}

	keyPEMBlock, err := os.ReadFile(keyFile)
	if err != nil {
		return err
	}

	_, err = tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return err
	}

	filePrex := fmt.Sprintf("Node-%s-Cert-%s", n.Name(), typ)

	err = n.m.writeEncryptFile(fmt.Sprintf("%s-Cert", filePrex), certPEMBlock)
	if err != nil {
		return err
	}

	err = n.m.writeEncryptFile(fmt.Sprintf("%s-Key", filePrex), keyPEMBlock)
	if err != nil {
		return err
	}

	n.nodeData.HasCerts[string(typ)] = true
	return n.nodeData.save(n.m)
}

func (n *Node) GetCert(typ CertType) (*tls.Certificate, error) {
	if !n.nodeData.HasCerts[string(typ)] {
		return nil, fmt.Errorf("current not has %s cert, please add it first", typ)
	}

	filePrex := fmt.Sprintf("Node-%s-Cert-%s", n.Name(), typ)
	certPEMBlock, err := n.m.readEncryptFile(fmt.Sprintf("%s-Cert", filePrex))
	if err != nil {
		return nil, err
	}

	keyPEMBlock, err := n.m.readEncryptFile(fmt.Sprintf("%s-Key", filePrex))
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, err
	}

	return &cert, nil
}

func (n *Node) RemoveCert(typ CertType) error {
	filePrex := fmt.Sprintf("Node-%s-Cert-%s", n.Name(), typ)
	err := n.m.removeFile(fmt.Sprintf("%s-Cert", filePrex))
	if err != nil {
		return err
	}

	return n.m.removeFile(fmt.Sprintf("%s-Key", filePrex))
}

func (n *Node) remove() error {
	for _, certName := range n.StoredCerts() {
		_ = n.RemoveCert(CertType(certName))
	}

	return n.m.removeFile(fmt.Sprintf("Node-%s-MetaData", n.nodeData.NodeName))
}
