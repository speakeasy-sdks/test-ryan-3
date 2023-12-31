// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package webhooks

import (
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
	"net/http"
)

type StockUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// An unknown error occurred interacting with the API.
	Error *components.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *StockUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StockUpdateResponse) GetError() *components.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *StockUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StockUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

type StockUpdateRequestBody struct {
	Drink      *components.DrinkInput      `json:"drink,omitempty"`
	Ingredient *components.IngredientInput `json:"ingredient,omitempty"`
}

func (o *StockUpdateRequestBody) GetDrink() *components.DrinkInput {
	if o == nil {
		return nil
	}
	return o.Drink
}

func (o *StockUpdateRequestBody) GetIngredient() *components.IngredientInput {
	if o == nil {
		return nil
	}
	return o.Ingredient
}
