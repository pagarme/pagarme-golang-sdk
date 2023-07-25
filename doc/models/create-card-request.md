
# Create Card Request

Card data

## Structure

`CreateCardRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `number` | `*string` | Optional | Credit card number |
| `holderName` | `*string` | Optional | Holder name, as written on the card |
| `expMonth` | `*int` | Optional | The expiration month |
| `expYear` | `*int` | Optional | The expiration year, that can be informed with 2 or 4 digits |
| `cvv` | `*string` | Optional | The card's security code |
| `billingAddress` | [`*models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Optional | Card's billing address |
| `brand` | `*string` | Optional | Card brand |
| `billingAddressId` | `*string` | Optional | The address id for the billing address |
| `metadata` | `map[string]*string` | Optional | Metadata |
| `mType` | `*string` | Optional | Card type<br>**Default**: `"credit"` |
| `options` | [`*models.CreateCardOptionsRequest`](../../doc/models/create-card-options-request.md) | Optional | Options for creating the card |
| `holderDocument` | `*string` | Optional | Document number for the card's holder |
| `privateLabel` | `*bool` | Optional | Indicates whether it is a private label card |
| `label` | `*string` | Optional | - |
| `id` | `*string` | Optional | Identifier |
| `token` | `*string` | Optional | token identifier |

## Example (as JSON)

```json
{
  "type": "credit",
  "number": "number2",
  "holder_name": "holder_name4",
  "exp_month": 42,
  "exp_year": 254,
  "cvv": "cvv2"
}
```

