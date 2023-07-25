
# Get Customer Response

Response object for getting a customer

## Structure

`GetCustomerResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `name` | `types.Optional[string]` | Optional | - |
| `email` | `types.Optional[string]` | Optional | - |
| `delinquent` | `types.Optional[bool]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `document` | `types.Optional[string]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | - |
| `fbAccessToken` | `types.Optional[string]` | Optional | - |
| `address` | [`types.Optional[models.GetAddressResponse]`](../../doc/models/get-address-response.md) | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `phones` | [`types.Optional[models.GetPhonesResponse]`](../../doc/models/get-phones-response.md) | Optional | - |
| `fbId` | `types.Optional[int64]` | Optional | - |
| `code` | `types.Optional[string]` | Optional | Código de referência do cliente no sistema da loja. Max: 52 caracteres |
| `documentType` | `types.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "name": "name0",
  "email": "email6",
  "delinquent": false,
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

