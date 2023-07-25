
# Get Usage Response

Response object for getting a usage

## Structure

`GetUsageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | Id |
| `quantity` | `types.Optional[int]` | Optional | Quantity |
| `description` | `types.Optional[string]` | Optional | Description |
| `usedAt` | `types.Optional[time.Time]` | Optional | Used at |
| `createdAt` | `types.Optional[time.Time]` | Optional | Creation date |
| `status` | `types.Optional[string]` | Optional | Status |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |
| `subscriptionItem` | [`types.Optional[models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | Subscription item |
| `code` | `types.Optional[string]` | Optional | Identification code in the client system |
| `group` | `types.Optional[string]` | Optional | Identification group in the client system |
| `amount` | `types.Optional[int]` | Optional | Field used in item scheme type 'Percent' |

## Example (as JSON)

```json
{
  "id": "id0",
  "quantity": 68,
  "description": "description0",
  "used_at": "2016-03-13T12:52:32.123Z",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

