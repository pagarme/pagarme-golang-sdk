
# Get Movement Object Transfer Response

## Structure

`GetMovementObjectTransferResponse`

## Inherits From

[`GetMovementObjectBaseResponse`](../../doc/models/get-movement-object-base-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `sourceType` | `types.Optional[string]` | Optional | - |
| `sourceId` | `types.Optional[string]` | Optional | - |
| `targetType` | `types.Optional[string]` | Optional | - |
| `targetId` | `types.Optional[string]` | Optional | - |
| `fee` | `types.Optional[string]` | Optional | - |
| `fundingDate` | `types.Optional[string]` | Optional | - |
| `fundingEstimatedDate` | `types.Optional[string]` | Optional | - |
| `bankAccount` | `types.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "object": "transfer",
  "id": "id6",
  "status": "status2",
  "amount": "amount8",
  "created_at": "created_at4",
  "source_type": "source_type0",
  "source_id": "source_id6",
  "target_type": "target_type2",
  "target_id": "target_id0",
  "fee": "fee2"
}
```

