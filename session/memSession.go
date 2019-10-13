package session

import "fmt"

/**
sesson memory实现
*/

//Get ...
func (d *Data) Get(key string) (data interface{}, err error) {
	defer d.rwLock.RUnlock()
	d.rwLock.RLock()
	var ok bool
	data, ok = d.Data[key]
	if !ok {
		err = fmt.Errorf("key is not exist")
		return
	}
	return
}

//Set ...
func (d *Data) Set(key string, value interface{}) {
	defer d.rwLock.Unlock()
	d.rwLock.Lock()
	d.Data[key] = value
}

//Del ...
func (d *Data) Del(key string) {
	defer d.rwLock.Unlock()
	d.rwLock.Lock()
	delete(d.Data, key)
}
