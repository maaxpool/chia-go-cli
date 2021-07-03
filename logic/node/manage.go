package node

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofrs/flock"
	"io"
	"os"
	"path/filepath"
)

var defaultTestData = []byte("github.com/LuttyYang/chia-go-cli")

type Manage struct {
	lock     *flock.Flock
	dir      string
	password string
}

func NewManage(dir string, password string) (*Manage, error) {
	manage := &Manage{dir: dir, password: password}

	return manage.init()
}

func (m *Manage) init() (*Manage, error) {
	if _, err := os.Stat(m.dir); os.IsNotExist(err) {
		err := os.MkdirAll(m.dir, 0700)
		if err != nil {
			return nil, err
		}

		err = m.writeEncryptFile(m.getRealFilePath("init"), defaultTestData)
		if err != nil {
			return nil, err
		}
	}

	testContent, err := m.readEncryptFile(m.getRealFilePath("init"))
	if err != nil {
		return nil, err
	}

	if bytes.Compare(testContent, defaultTestData) != 0 {
		return nil, errors.New("data is corrupted, maybe your password is not valid")
	}

	lock := flock.New(m.getRealFilePath("LOCK"))
	isLock, err := lock.TryLock()
	if err != nil {
		return nil, err
	}

	if !isLock {
		return nil, errors.New("lock fail, maybe has another process")
	}

	m.lock = lock

	return m, nil
}

func (m *Manage) getRealFilePath(fileName string) string {
	hash := md5.Sum([]byte(fileName))

	return filepath.Join(m.dir, hex.EncodeToString(hash[:]))
}

func (m *Manage) getAesCipher() (cipher.AEAD, error) {
	hash := md5.Sum([]byte(m.password))

	c, err := aes.NewCipher(hash[:])
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(c)
}

func (m *Manage) hasFile(fileName string) bool {
	if _, err := os.Stat(m.getRealFilePath(fileName)); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (m *Manage) removeFile(fileName string) error {
	return os.Remove(m.getRealFilePath(fileName))
}

func (m *Manage) readEncryptFile(fileName string) ([]byte, error) {
	encryptData, err := os.ReadFile(m.getRealFilePath(fileName))
	if err != nil {
		return nil, err
	}

	gcm, err := m.getAesCipher()
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptData) < gcm.NonceSize() {
		return nil, errors.New("bad src data, maybe data is corrupted")
	}

	nonce, ciphertext := encryptData[:nonceSize], encryptData[nonceSize:]
	oriByte, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return oriByte, nil
}

func (m *Manage) writeEncryptFile(fileName string, data []byte) error {
	gcm, err := m.getAesCipher()
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	encryptData := gcm.Seal(nonce, nonce, data, nil)
	return os.WriteFile(m.getRealFilePath(fileName), encryptData, 0600)
}

func (m *Manage) Close() error {
	if m.lock != nil {
		err := m.lock.Unlock()
		if err != nil {
			return err
		}

		m.lock = nil
	}

	return nil
}

func (m *Manage) getNodeList() (list []string, err error) {
	if !m.hasFile("NodeList") {
		return []string{}, nil
	}

	nodeList := make([]string, 0)
	listJsonByte, err := m.readEncryptFile("NodeList")
	if err != nil {

		return nil, err
	}

	err = json.Unmarshal(listJsonByte, &nodeList)
	if err != nil {
		return nil, err
	}

	return nodeList, err
}

func (m *Manage) addNameToNodeList(name string) (err error) {
	nodeList, err := m.getNodeList()
	if err != nil {
		return err
	}

	for _, s := range nodeList {
		if s == name {
			return nil
		}
	}

	nodeList = append(nodeList, name)
	return m.saveNodeList(nodeList)
}

func (m *Manage) removeNameFromNodeList(name string) (err error) {
	nodeList, err := m.getNodeList()
	if err != nil {
		return err
	}

	for i, s := range nodeList {
		if s == name {
			return m.saveNodeList(append(nodeList[:i], nodeList[i+1:]...))
		}
	}

	return nil
}

func (m *Manage) saveNodeList(list []string) error {
	listJsonByte, err := json.Marshal(list)
	if err != nil {
		return err
	}

	return m.writeEncryptFile("NodeList", listJsonByte)
}

func (m *Manage) GetNodeList() (list []*Node, err error) {
	nodeList, err := m.getNodeList()
	if err != nil {
		return nil, err
	}

	list = make([]*Node, 0, len(nodeList))
	for _, data := range nodeList {
		node, err := newNode(m, data)
		if err != nil {
			return nil, err
		}

		list = append(list, node)
	}

	return list, nil
}

func (m *Manage) GetNode(nodeName string) (node *Node, err error) {
	return newNode(m, nodeName)
}

func (m *Manage) HasNode(nodeName string) bool {
	return m.hasFile(fmt.Sprintf("Node-%s-MetaData", nodeName))
}

func (m *Manage) CreateNode(nodeName string, nodeUrl string) (node *Node, err error) {
	if m.HasNode(nodeName) {
		return nil, errors.New("node is already exists")
	}

	fileName := fmt.Sprintf("Node-%s-MetaData", nodeName)
	data := &nodeData{
		NodeName: nodeName,
		NodeUrl:  nodeUrl,
		HasCerts: make(map[string]bool, 0),
	}

	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = m.writeEncryptFile(fileName, dataByte)
	if err != nil {
		return nil, err
	}

	err = m.addNameToNodeList(nodeName)
	if err != nil {
		return nil, err
	}

	return m.GetNode(nodeName)
}

func (m *Manage) GetOrCreateNode(nodeName string, nodeUrl string) (node *Node, err error) {
	if m.HasNode(nodeName) {
		return m.GetNode(nodeName)
	}

	return m.CreateNode(nodeName, nodeUrl)
}

func (m *Manage) RemoveNode(nodeName string) error {
	node, err := m.GetNode(nodeName)
	if err != nil {
		return err
	}

	err = m.removeNameFromNodeList(nodeName)
	if err != nil {
		return err
	}

	return node.remove()
}
