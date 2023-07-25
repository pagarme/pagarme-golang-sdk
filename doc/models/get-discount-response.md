
# Get Discount Response

Response object for getting a discount

## Structure

`GetDiscountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `value` | `types.Optional[float64]` | Optional | - |
| `discountType` | `types.Optional[string]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `cycles` | `types.Optional[int]` | Optional | - |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |
| `description` | `types.Optional[string]` | Optional | - |
| `subscription` | [`types.Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `subscriptionItem` | [`types.Optional[models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | The subscription item |

## Example (as JSON)

```json
{
  "id": "id0",
  "value": 251.52,
  "discount_type": "discount_type8",
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

