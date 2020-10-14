package handler

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"unsafe"

	"github.com/golang/glog"
)

// TCP防护策略
type TCP struct {
	Filter  uint32 `json:"filter"`
	Algo    uint32 `json:"algorithm"`
	Thresh1 uint32 `json:"thresh1"`
	Thresh2 uint32 `json:"thresh2"`
}

var DefaultTCP = new(TCP)

func GetDefaultTCP() *TCP {
	return DefaultTCP
}

func (t *TCP) Handler(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body)
	nt, err := t.Decode(string(s))
	if err != nil {
		return
	}

	nt.Valid()
	nt.SerializeAndSend()
	w.Write(s)
}

func (t *TCP) Decode(buf string) (*TCP, error) {
	nt := new(TCP)
	err := json.Unmarshal([]byte(buf), nt)
	if err != nil {
		glog.Error("json unmarshal error:", err)
	}
	return nt, err
}

func (t *TCP) Valid() {
	return
}

func (t *TCP) SerializeAndSend() {
	var hdr = Header{
		Type:   1,
		Len:    uint32(unsafe.Sizeof(*t)),
		Cmd:    1,
		Status: 1,
		Id:     1,
	}

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, hdr)
	err = binary.Write(&buf, binary.LittleEndian, *t)
	glog.Info(buf, err)
	// TODO: send policy to the engine
	return
}
