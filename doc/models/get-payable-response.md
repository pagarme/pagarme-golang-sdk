
# Get Payable Response

Response object for getting an payable

## Structure

`GetPayableResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[int64]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `amount` | `types.Optional[int]` | Optional | - |
| `fee` | `types.Optional[int]` | Optional | - |
| `anticipationFee` | `types.Optional[int]` | Optional | - |
| `fraudCoverageFee` | `types.Optional[int]` | Optional | - |
| `installment` | `types.Optional[int]` | Optional | - |
| `gatewayId` | `types.Optional[int]` | Optional | - |
| `chargeId` | `types.Optional[string]` | Optional | - |
| `splitId` | `types.Optional[string]` | Optional | - |
| `bulkAnticipationId` | `types.Optional[string]` | Optional | - |
| `anticipationId` | `types.Optional[string]` | Optional | - |
| `recipientId` | `types.Optional[string]` | Optional | - |
| `originatorModel` | `types.Optional[string]` | Optional | - |
| `originatorModelId` | `types.Optional[string]` | Optional | - |
| `paymentDate` | `types.Optional[time.Time]` | Optional | - |
| `originalPaymentDate` | `types.Optional[time.Time]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | - |
| `paymentMethod` | `types.Optional[string]` | Optional | - |
| `accrualAt` | `types.Optional[time.Time]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `liquidationArrangementId` | `types.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": 112,
  "status": "status8",
  "amount": 46,
  "fee": 168,
  "anticipation_fee": 140
}
```

