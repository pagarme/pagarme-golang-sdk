
# Get Card Token Response

Card token data

## Structure

`GetCardTokenResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `lastFourDigits` | `types.Optional[string]` | Optional | - |
| `holderName` | `types.Optional[string]` | Optional | - |
| `holderDocument` | `types.Optional[string]` | Optional | - |
| `expMonth` | `types.Optional[int]` | Optional | - |
| `expYear` | `types.Optional[int]` | Optional | - |
| `brand` | `types.Optional[string]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | - |
| `label` | `types.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "last_four_digits": "last_four_digits6",
  "holder_name": "holder_name4",
  "holder_document": "holder_document6",
  "exp_month": 42,
  "exp_year": 254
}
```

