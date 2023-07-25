
# Get Plan Response

Response object for getting a plan

## Structure

`GetPlanResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `name` | `types.Optional[string]` | Optional | - |
| `description` | `types.Optional[string]` | Optional | - |
| `url` | `types.Optional[string]` | Optional | - |
| `statementDescriptor` | `types.Optional[string]` | Optional | - |
| `interval` | `types.Optional[string]` | Optional | - |
| `intervalCount` | `types.Optional[int]` | Optional | - |
| `billingType` | `types.Optional[string]` | Optional | - |
| `paymentMethods` | `types.Optional[[]string]` | Optional | - |
| `installments` | `types.Optional[[]int]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `currency` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `items` | [`types.Optional[[]models.GetPlanItemResponse]`](../../doc/models/get-plan-item-response.md) | Optional | - |
| `billingDays` | `types.Optional[[]int]` | Optional | - |
| `shippable` | `types.Optional[bool]` | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `trialPeriodDays` | `types.Optional[int]` | Optional | - |
| `minimumPrice` | `types.Optional[int]` | Optional | - |
| `deletedAt` | `types.Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "name": "name0",
  "description": "description0",
  "url": "url4",
  "statement_descriptor": "statement_descriptor0"
}
```

