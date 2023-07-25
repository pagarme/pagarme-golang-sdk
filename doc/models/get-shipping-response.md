
# Get Shipping Response

Response object for getting the shipping data

## Structure

`GetShippingResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `amount` | `types.Optional[int]` | Optional | - |
| `description` | `types.Optional[string]` | Optional | - |
| `recipientName` | `types.Optional[string]` | Optional | - |
| `recipientPhone` | `types.Optional[string]` | Optional | - |
| `address` | [`types.Optional[models.GetAddressResponse]`](../../doc/models/get-address-response.md) | Optional | - |
| `maxDeliveryDate` | `types.Optional[time.Time]` | Optional | Data máxima de entrega |
| `estimatedDeliveryDate` | `types.Optional[time.Time]` | Optional | Prazo estimado de entrega |
| `mType` | `types.Optional[string]` | Optional | Shipping Type |

## Example (as JSON)

```json
{
  "amount": 46,
  "description": "description0",
  "recipient_name": "recipient_name8",
  "recipient_phone": "recipient_phone2",
  "address": {
    "id": "id6",
    "street": "street6",
    "number": "number4",
    "complement": "complement2",
    "zip_code": "zip_code0"
  }
}
```

