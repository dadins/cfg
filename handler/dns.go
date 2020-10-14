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

// DNS防护策略
type DNS struct {
	Filter     uint16 `json:"filter"`
	Enable     uint8  `json:"enable"`
	LimitLevel uint8  `json:"limit_level"`
	Algo       uint32 `json:"algo"`
	TotalLimit uint32 `json:"total_limit"`
}

var DefaultDNS = new(DNS)

func GetDefaultDNS() *DNS {
	return DefaultDNS
}

func (d *DNS) Handler(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body)
	nd, err := d.Decode(string(s))
	if err != nil {
		return
	}

	nd.Valid()
	nd.SerializeAndSend()
	w.Write(s)
}

func (d *DNS) Decode(buf string) (*DNS, error) {
	nd := new(DNS)
	err := json.Unmarshal([]byte(buf), nd)
	if err != nil {
		glog.Error("json unmarshal error:", err)
	}
	return nd, err
}

func (d *DNS) Valid() {
	return
}

func (d *DNS) SerializeAndSend() {
	var hdr = Header{
		Type:   1,
		Len:    uint32(unsafe.Sizeof(*d)),
		Cmd:    1,
		Status: 1,
		Id:     1,
	}

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, hdr)
	err = binary.Write(&buf, binary.LittleEndian, *d)
	glog.Info(buf, err)
	// TODO: send policy to the engine
	return
}
