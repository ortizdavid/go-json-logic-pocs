curl -X POST http://localhost:8080/apply-discounts \
-H "Content-Type: application/json" \
-d '{
  "order": {
    "customerType": "PF",
    "discountPercent": 12,
    "grossTotal": 1200,
    "deliveryType": "ENTREGA_DOMICILIO",
    "deliveryFee": 10,
    "managerApproved": false,
    "items": [
      { "product": "Produto A", "quantity": 2, "unitPrice": 500, "subtotal": 1000 },
      { "product": "Produto B", "quantity": 1, "unitPrice": 200, "subtotal": 200 }
    ]
  }
}'
