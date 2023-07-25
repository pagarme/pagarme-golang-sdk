
# Get Balance Operation Response

Generic response object for getting a BalanceOperation.

## Structure

`GetBalanceOperationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `balanceAmount` | `types.Optional[string]` | Optional | - |
| `balanceOldAmount` | `types.Optional[string]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | - |
| `amount` | `types.Optional[string]` | Optional | - |
| `fee` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[string]` | Optional | - |
| `movementObject` | [`*models.GetMovementObjectBaseResponseInterface`](../../doc/models/get-movement-object-base-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "status": "status8",
  "balance_amount": "balance_amount0",
  "balance_old_amount": "balance_old_amount2",
  "type": "type0"
}
```

