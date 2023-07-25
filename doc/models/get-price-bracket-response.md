
# Get Price Bracket Response

Response object for getting a price bracket

## Structure

`GetPriceBracketResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `startQuantity` | `types.Optional[int]` | Optional | - |
| `price` | `types.Optional[int]` | Optional | - |
| `endQuantity` | `types.Optional[int]` | Optional | - |
| `overagePrice` | `types.Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "start_quantity": 46,
  "price": 16,
  "end_quantity": 54,
  "overage_price": 68
}
```

