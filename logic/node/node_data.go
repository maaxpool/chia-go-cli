package node

import (
	"encoding/json"
	"errors"
	"fmt"
)

type nodeData struct {
	NodeName string          `json:"node_name"`
	NodeUrl  string          `json:"node_url"`
	HasCerts map[string]bool `json:"has_certs"`
}

func (d *nodeData) save(m *Manage) error {
	nodeDataByte, err := json.Marshal(d)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("Node-%s-MetaData", d.NodeName)
	return m.writeEncryptFile(fileName, nodeDataByte)
}

func (d *nodeData) load(m *Manage) error {
	fileName := fmt.Sprintf("Node-%s-MetaData", d.NodeName)
	if !m.hasFile(fileName) {
		return errors.New("node is not exists")
	}

	nodeDataByte, err := m.readEncryptFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(nodeDataByte, d)
}
