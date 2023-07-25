
# Get Transaction Response

Generic response object for getting a transaction.

## Structure

`GetTransactionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `gatewayId` | `types.Optional[string]` | Optional | Gateway transaction id |
| `amount` | `types.Optional[int]` | Optional | Amount in cents |
| `status` | `types.Optional[string]` | Optional | Transaction status |
| `success` | `types.Optional[bool]` | Optional | Indicates if the transaction ocurred successfuly |
| `createdAt` | `types.Optional[time.Time]` | Optional | Creation date |
| `updatedAt` | `types.Optional[time.Time]` | Optional | Last update date |
| `attemptCount` | `types.Optional[int]` | Optional | Number of attempts tried |
| `maxAttempts` | `types.Optional[int]` | Optional | Max attempts |
| `splits` | [`types.Optional[[]models.GetSplitResponse]`](../../doc/models/get-split-response.md) | Optional | Splits |
| `nextAttempt` | `types.Optional[time.Time]` | Optional | Date and time of the next attempt |
| `transactionType` | `*string` | Optional | - |
| `id` | `types.Optional[string]` | Optional | Código da transação |
| `gatewayResponse` | [`types.Optional[models.GetGatewayResponseResponse]`](../../doc/models/get-gateway-response-response.md) | Optional | The Gateway Response |
| `antifraudResponse` | [`types.Optional[models.GetAntifraudResponse]`](../../doc/models/get-antifraud-response.md) | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `split` | [`types.Optional[[]models.GetSplitResponse]`](../../doc/models/get-split-response.md) | Optional | - |
| `interest` | [`types.Optional[models.GetInterestResponse]`](../../doc/models/get-interest-response.md) | Optional | - |
| `fine` | [`types.Optional[models.GetFineResponse]`](../../doc/models/get-fine-response.md) | Optional | - |
| `maxDaysToPayPastDue` | `types.Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id0",
  "amount": 46,
  "status": "status8",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "transaction_type": "transaction"
}
```

