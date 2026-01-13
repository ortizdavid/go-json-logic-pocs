+--------------------+
|   Frontend React   |
|  (PDV do cliente)  |
+--------------------+
          |
          | 1. Envia payload do pedido
          |    Exemplo:
          |    {
          |      "order": {
          |        "customerType": "PF",
          |        "discountPercent": 12,
          |        "grossTotal": 1200,
          |        "deliveryType": "ENTREGA_DOMICILIO",
          |        "deliveryFee": 10,
          |        "managerApproved": false
          |      }
          |    }
          v
+--------------------+
|     Backend Go     |
|  (API multi-tenant)|
+--------------------+
          |
          | 2. Seleciona regras JsonLogic
          |    Baseado no tenant (cliente)
          |    - discounts.json
          |    - restrictions.json
          v
+--------------------+
|    JsonLogic       |
|   Engine Go        |
+--------------------+
          |
          | 3. Avalia regras sobre o payload
          |    - Calcula desconto final
          |    - Checa frete permitido
          |    - Determina se precisa aprovação
          v
+--------------------+
| Resultado Backend  |
|  (JSON pronto)     |
+--------------------+
          |
          | 4. Retorna para o front
          |    Exemplo:
          |    {
          |      "discountPercent": 10,
          |      "allowed": true,
          |      "needsApproval": false
          |    }
          v
+--------------------+
|   Frontend React   |
| Atualiza tela/UX   |
|  - Mostra valores  |
|  - Bloqueia checkout se necessário |
+--------------------+
