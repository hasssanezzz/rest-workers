package types

import "math/big"

type Payload struct {
	Number *big.Int `json:"number"`
}
