package session

import (
	"errors"
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

var manager *Manager

func init() {
	fmt.Println("init manager ")
	manager = &Manager{
		Session: make(map[string]*Data, 1024),
	}
}

//Data ...
type Data struct {
	ID     string
	Data   map[string]interface{}
	rwLock sync.RWMutex
}

//Manager ...
type Manager struct {
	Session map[string]*Data
	rwLock  sync.RWMutex
}

//NewData ...
func NewData(key string) *Data {
	return &Data{
		ID:   key,
		Data: make(map[string]interface{}),
	}
}

//GetSession ...
func (m *Manager) GetSession(key string) (s *Data, err error) {
	defer m.rwLock.RUnlock()
	m.rwLock.RLock()
	var ok bool
	s, ok = m.Session[key]
	if !ok {
		err = errors.New("SessionID is not exist")
		return
	}
	return
}

//Create ...
func (m *Manager) Create() *Data {
	//1.create session id
	uuid := uuid.NewV4()
	//2.new Data
	data := NewData(uuid.String())
	return data
}
