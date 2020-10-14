package handler

// ICMP防护策略
type ICMP struct {
	Filter uint32 `json:"filter"`
	Thresh uint32 `json:"thresh"`
}
