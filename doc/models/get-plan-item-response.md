
# Get Plan Item Response

Response object for getting a plan item

## Structure

`GetPlanItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `name` | `types.Optional[string]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `pricingScheme` | [`types.Optional[models.GetPricingSchemeResponse]`](../../doc/models/get-pricing-scheme-response.md) | Optional | - |
| `description` | `types.Optional[string]` | Optional | - |
| `plan` | [`types.Optional[models.GetPlanResponse]`](../../doc/models/get-plan-response.md) | Optional | - |
| `quantity` | `types.Optional[int]` | Optional | - |
| `cycles` | `types.Optional[int]` | Optional | - |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "name": "name0",
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z"
}
```

