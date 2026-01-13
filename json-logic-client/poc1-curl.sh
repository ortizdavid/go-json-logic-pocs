curl -X POST http://localhost:8080/apply-restrictions \
  -H "Content-Type: application/json" \
  -d @poc1-payload.json


curl -X POST http://localhost:8080/apply-discounts \
  -H "Content-Type: application/json" \
  -d @poc1-payload.json

