
# Get Order Response

Response object for getting an Order

## Structure

`GetOrderResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `code` | `types.Optional[string]` | Optional | - |
| `currency` | `types.Optional[string]` | Optional | - |
| `items` | [`types.Optional[[]models.GetOrderItemResponse]`](../../doc/models/get-order-item-response.md) | Optional | - |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `charges` | [`types.Optional[[]models.GetChargeResponse]`](../../doc/models/get-charge-response.md) | Optional | - |
| `invoiceUrl` | `types.Optional[string]` | Optional | - |
| `shipping` | [`types.Optional[models.GetShippingResponse]`](../../doc/models/get-shipping-response.md) | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `checkouts` | [`types.Optional[[]models.GetCheckoutPaymentResponse]`](../../doc/models/get-checkout-payment-response.md) | Optional | Checkout Payment Settings Response |
| `ip` | `types.Optional[string]` | Optional | Ip address |
| `sessionId` | `types.Optional[string]` | Optional | Session id |
| `location` | [`types.Optional[models.GetLocationResponse]`](../../doc/models/get-location-response.md) | Optional | Location |
| `device` | [`types.Optional[models.GetDeviceResponse]`](../../doc/models/get-device-response.md) | Optional | Device's informations |
| `closed` | `types.Optional[bool]` | Optional | Indicates whether the order is closed |

## Example (as JSON)

```json
{
  "id": "id0",
  "code": "code8",
  "currency": "currency0",
  "items": [
    {
      "id": "id7",
      "amount": 13,
      "description": "description7",
      "quantity": 127,
      "category": "category5"
    },
    {
      "id": "id8",
      "amount": 14,
      "description": "description8",
      "quantity": 128,
      "category": "category6"
    }
  ],
  "customer": {
    "id": "id0",
    "name": "name0",
    "email": "email6",
    "delinquent": false,
    "created_at": "2016-03-13T12:52:32.123Z"
  }
}
```

