
# Get Subscription Item Response

## Structure

`GetSubscriptionItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `description` | `types.Optional[string]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `pricingScheme` | [`types.Optional[models.GetPricingSchemeResponse]`](../../doc/models/get-pricing-scheme-response.md) | Optional | - |
| `discounts` | [`types.Optional[[]models.GetDiscountResponse]`](../../doc/models/get-discount-response.md) | Optional | - |
| `increments` | [`types.Optional[[]models.GetIncrementResponse]`](../../doc/models/get-increment-response.md) | Optional | - |
| `subscription` | [`types.Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `name` | `types.Optional[string]` | Optional | Item name |
| `quantity` | `types.Optional[int]` | Optional | - |
| `cycles` | `types.Optional[int]` | Optional | - |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "description": "description0",
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z"
}
```

