
# Get Private Label Transaction Response

Response object for getting a private label transaction

## Structure

`GetPrivateLabelTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `types.Optional[string]` | Optional | Text that will appear on the credit card's statement |
| `acquirerName` | `types.Optional[string]` | Optional | Acquirer name |
| `acquirerAffiliationCode` | `types.Optional[string]` | Optional | Aquirer affiliation code |
| `acquirerTid` | `types.Optional[string]` | Optional | Acquirer TID |
| `acquirerNsu` | `types.Optional[string]` | Optional | Acquirer NSU |
| `acquirerAuthCode` | `types.Optional[string]` | Optional | Acquirer authorization code |
| `operationType` | `types.Optional[string]` | Optional | Operation type |
| `card` | [`types.Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | Card data |
| `acquirerMessage` | `types.Optional[string]` | Optional | Acquirer message |
| `acquirerReturnCode` | `types.Optional[string]` | Optional | Acquirer Return Code |
| `installments` | `types.Optional[int]` | Optional | Number of installments |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id6",
  "amount": 20,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "transaction_type": "private_label",
  "statement_descriptor": "statement_descriptor0",
  "acquirer_name": "acquirer_name4",
  "acquirer_affiliation_code": "acquirer_affiliation_code8",
  "acquirer_tid": "acquirer_tid0",
  "acquirer_nsu": "acquirer_nsu0"
}
```

