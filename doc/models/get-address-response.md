
# Get Address Response

Response object for getting an Address

## Structure

`GetAddressResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `street` | `types.Optional[string]` | Optional | - |
| `number` | `types.Optional[string]` | Optional | - |
| `complement` | `types.Optional[string]` | Optional | - |
| `zipCode` | `types.Optional[string]` | Optional | - |
| `neighborhood` | `types.Optional[string]` | Optional | - |
| `city` | `types.Optional[string]` | Optional | - |
| `state` | `types.Optional[string]` | Optional | - |
| `country` | `types.Optional[string]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `line1` | `types.Optional[string]` | Optional | Line 1 for address |
| `line2` | `types.Optional[string]` | Optional | Line 2 for address |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "street": "street0",
  "number": "number2",
  "complement": "complement4",
  "zip_code": "zip_code4"
}
```

