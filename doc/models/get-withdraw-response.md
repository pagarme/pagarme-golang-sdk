
# Get Withdraw Response

## Structure

`GetWithdrawResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `gatewayId` | `types.Optional[string]` | Optional | - |
| `amount` | `types.Optional[int]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `metadata` | `types.Optional[[]string]` | Optional | - |
| `fee` | `types.Optional[int]` | Optional | - |
| `fundingDate` | `types.Optional[time.Time]` | Optional | - |
| `fundingEstimatedDate` | `types.Optional[time.Time]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | - |
| `source` | [`types.Optional[models.GetWithdrawSourceResponse]`](../../doc/models/get-withdraw-source-response.md) | Optional | - |
| `target` | [`types.Optional[models.GetWithdrawTargetResponse]`](../../doc/models/get-withdraw-target-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "gateway_id": "gateway_id0",
  "amount": 46,
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

