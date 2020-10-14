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

// 群组策略
type Group struct {
	Tcp  TCP  `json:"tcp"`
	Udp  UDP  `json:"udp"`
	Icmp ICMP `json:"icmp"`
	Dns  DNS  `json:"dns"`
}

var DefaultGroup = new(Group)

func GetDefaultGroup() *Group {
	return DefaultGroup
}

func (g *Group) Handler(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body)
	ng, err := g.Decode(string(s))
	if err != nil {
		return
	}

	glog.Info(ng)
	ng.Valid()
	ng.SerializeAndSend()
	w.Write(s)
}

func (g *Group) Decode(buf string) (*Group, error) {
	ng := new(Group)
	err := json.Unmarshal([]byte(buf), ng)
	if err != nil {
		glog.Error("json unmarshal error:", err)
	}
	return ng, err
}

func (g *Group) Valid() {
	return
}

func (g *Group) SerializeAndSend() {
	var hdr = Header{
		Type:   1,
		Len:    uint32(unsafe.Sizeof(*g)),
		Cmd:    1,
		Status: 1,
		Id:     1,
	}
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, hdr)
	err = binary.Write(&buf, binary.LittleEndian, *g)
	// TODO: send policy to the engine
	glog.Info(buf, err)
}
