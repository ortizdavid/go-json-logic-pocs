package models


// Struct do pedido
type Order struct {
	CustomerType    string  `json:"customerType"`
	DiscountPercent float64 `json:"discountPercent"`
	GrossTotal      float64 `json:"grossTotal"`
	DeliveryType    string  `json:"deliveryType"`
	DeliveryFee     float64 `json:"deliveryFee"`
	ManagerApproved bool    `json:"managerApproved"`
}

// Payload esperado: { "order": {...} }
type Payload struct {
	Order Order `json:"order"`
}

