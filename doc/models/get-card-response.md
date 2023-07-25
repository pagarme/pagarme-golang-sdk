
# Get Card Response

Response object for getting a credit card

## Structure

`GetCardResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `lastFourDigits` | `types.Optional[string]` | Optional | - |
| `brand` | `types.Optional[string]` | Optional | - |
| `holderName` | `types.Optional[string]` | Optional | - |
| `expMonth` | `types.Optional[int]` | Optional | - |
| `expYear` | `types.Optional[int]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `billingAddress` | [`types.Optional[models.GetBillingAddressResponse]`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | Card type |
| `holderDocument` | `types.Optional[string]` | Optional | Document number for the card's holder |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |
| `firstSixDigits` | `types.Optional[string]` | Optional | First six digits |
| `label` | `types.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "last_four_digits": "last_four_digits6",
  "brand": "brand4",
  "holder_name": "holder_name4",
  "exp_month": 42
}
```

