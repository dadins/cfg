package handler

// UDP防护策略
type UDP struct {
	Filter uint32 `json:"filter"`
	Thresh uint32 `json:"thresh"`
}
