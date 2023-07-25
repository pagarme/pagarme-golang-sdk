
# Get Voucher Transaction Response

Response for voucher transactions

## Structure

`GetVoucherTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `types.Optional[string]` | Optional | Text that will appear on the voucher's statement |
| `acquirerName` | `types.Optional[string]` | Optional | Acquirer name |
| `acquirerAffiliationCode` | `types.Optional[string]` | Optional | Acquirer affiliation code |
| `acquirerTid` | `types.Optional[string]` | Optional | Acquirer TID |
| `acquirerNsu` | `types.Optional[string]` | Optional | Acquirer NSU |
| `acquirerAuthCode` | `types.Optional[string]` | Optional | Acquirer authorization code |
| `acquirerMessage` | `types.Optional[string]` | Optional | acquirer_message |
| `acquirerReturnCode` | `types.Optional[string]` | Optional | Acquirer return code |
| `operationType` | `types.Optional[string]` | Optional | Operation type |
| `card` | [`types.Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | Card data |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id4",
  "amount": 24,
  "status": "status2",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "transaction_type": "voucher",
  "statement_descriptor": "statement_descriptor0",
  "acquirer_name": "acquirer_name4",
  "acquirer_affiliation_code": "acquirer_affiliation_code8",
  "acquirer_tid": "acquirer_tid0",
  "acquirer_nsu": "acquirer_nsu0"
}
```

