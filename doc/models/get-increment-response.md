
# Get Increment Response

Response object for getting a increment

## Structure

`GetIncrementResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `value` | `types.Optional[float64]` | Optional | - |
| `incrementType` | `types.Optional[string]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `cycles` | `types.Optional[int]` | Optional | - |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |
| `description` | `types.Optional[string]` | Optional | - |
| `subscription` | [`types.Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `subscriptionItem` | [`types.Optional[models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | The Subscription Item |

## Example (as JSON)

```json
{
  "id": "id0",
  "value": 251.52,
  "increment_type": "increment_type8",
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

