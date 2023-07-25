
# Get Period Response

Response object for getting a period

## Structure

`GetPeriodResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `startAt` | `types.Optional[time.Time]` | Optional | - |
| `endAt` | `types.Optional[time.Time]` | Optional | - |
| `id` | `types.Optional[string]` | Optional | - |
| `billingAt` | `types.Optional[time.Time]` | Optional | - |
| `subscription` | [`types.Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `duration` | `types.Optional[int]` | Optional | - |
| `createdAt` | `types.Optional[string]` | Optional | - |
| `updatedAt` | `types.Optional[string]` | Optional | - |
| `cycle` | `types.Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "start_at": "2016-03-13T12:52:32.123Z",
  "end_at": "2016-03-13T12:52:32.123Z",
  "id": "id0",
  "billing_at": "2016-03-13T12:52:32.123Z",
  "subscription": {
    "id": "id4",
    "code": "code2",
    "start_at": "2016-03-13T12:52:32.123Z",
    "interval": "interval2",
    "interval_count": 234
  }
}
```

