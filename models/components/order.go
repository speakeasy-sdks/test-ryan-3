// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Status - The status of the order.
type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusComplete   Status = "complete"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "processing":
		fallthrough
	case "complete":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// Order - An order for a drink or ingredient.
type Order struct {
	// The product code of the drink or ingredient.
	ProductCode string `json:"productCode"`
	// The number of units of the drink or ingredient to order.
	Quantity int64 `json:"quantity"`
	// The status of the order.
	Status Status `json:"status"`
	// The type of order.
	Type OrderType `json:"type"`
}

func (o *Order) GetProductCode() string {
	if o == nil {
		return ""
	}
	return o.ProductCode
}

func (o *Order) GetQuantity() int64 {
	if o == nil {
		return 0
	}
	return o.Quantity
}

func (o *Order) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}

func (o *Order) GetType() OrderType {
	if o == nil {
		return OrderType("")
	}
	return o.Type
}