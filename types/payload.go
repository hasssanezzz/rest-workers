package types

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Payload struct {
	Number *big.Int `json:"number"`
}

// Implementing json.Unmarshaler interface for Payload
// implement your own when customizing the payload object
func (p *Payload) UnmarshalJSON(data []byte) error {
	var jsonData struct {
		Number string `json:"number"`
	}

	if err := json.Unmarshal(data, &jsonData); err != nil {
		println(string(data))
		return err
	}

	p.Number = new(big.Int)
	_, ok := p.Number.SetString(jsonData.Number, 10)
	if !ok {
		return fmt.Errorf("invalid number format")
	}

	return nil
}
