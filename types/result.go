package types

type Result struct {
	Payload *Payload `json:"payload"`
	Result  bool     `json:"result"`
}
