package handler

// module_com header
type Header struct {
	Type   uint32 `json:"type"`
	Len    uint32 `json:"len"`
	Cmd    uint32 `json:"cmd"`
	Status int32  `json:"status"`
	Id     uint32 `json:"id"`
	Unused uint32 `json:"unused"`
}
