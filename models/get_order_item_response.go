package models

import (
    "encoding/json"
)

// GetOrderItemResponse represents a GetOrderItemResponse struct.
// Response object for getting an order item
type GetOrderItemResponse struct {
    // Id
    Id          Optional[string] `json:"id"`
    Amount      Optional[int]    `json:"amount"`
    Description Optional[string] `json:"description"`
    Quantity    Optional[int]    `json:"quantity"`
    // Category
    Category    Optional[string] `json:"category"`
    // Code
    Code        Optional[string] `json:"code"`
}

// MarshalJSON implements the json.Marshaler interface for GetOrderItemResponse.
// It customizes the JSON marshaling process for GetOrderItemResponse objects.
func (g *GetOrderItemResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetOrderItemResponse object to a map representation for JSON marshaling.
func (g *GetOrderItemResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if g.Id.IsValueSet() {
        structMap["id"] = g.Id.Value()
    }
    if g.Amount.IsValueSet() {
        structMap["amount"] = g.Amount.Value()
    }
    if g.Description.IsValueSet() {
        structMap["description"] = g.Description.Value()
    }
    if g.Quantity.IsValueSet() {
        structMap["quantity"] = g.Quantity.Value()
    }
    if g.Category.IsValueSet() {
        structMap["category"] = g.Category.Value()
    }
    if g.Code.IsValueSet() {
        structMap["code"] = g.Code.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOrderItemResponse.
// It customizes the JSON unmarshaling process for GetOrderItemResponse objects.
func (g *GetOrderItemResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id          Optional[string] `json:"id"`
        Amount      Optional[int]    `json:"amount"`
        Description Optional[string] `json:"description"`
        Quantity    Optional[int]    `json:"quantity"`
        Category    Optional[string] `json:"category"`
        Code        Optional[string] `json:"code"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Amount = temp.Amount
    g.Description = temp.Description
    g.Quantity = temp.Quantity
    g.Category = temp.Category
    g.Code = temp.Code
    return nil
}
