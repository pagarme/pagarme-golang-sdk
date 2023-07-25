
# Get Billing Address Response

Response object for getting a billing address

## Structure

`GetBillingAddressResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `street` | `types.Optional[string]` | Optional | - |
| `number` | `types.Optional[string]` | Optional | - |
| `zipCode` | `types.Optional[string]` | Optional | - |
| `neighborhood` | `types.Optional[string]` | Optional | - |
| `city` | `types.Optional[string]` | Optional | - |
| `state` | `types.Optional[string]` | Optional | - |
| `country` | `types.Optional[string]` | Optional | - |
| `complement` | `types.Optional[string]` | Optional | - |
| `line1` | `types.Optional[string]` | Optional | Line 1 for address |
| `line2` | `types.Optional[string]` | Optional | Line 2 for address |

## Example (as JSON)

```json
{
  "street": "street0",
  "number": "number2",
  "zip_code": "zip_code4",
  "neighborhood": "neighborhood6",
  "city": "city0"
}
```

