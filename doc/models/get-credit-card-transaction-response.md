
# Get Credit Card Transaction Response

Response object for getting a credit card transaction

## Structure

`GetCreditCardTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `types.Optional[string]` | Optional | Text that will appear on the credit card's statement |
| `acquirerName` | `*string` | Optional | Acquirer name |
| `acquirerAffiliationCode` | `types.Optional[string]` | Optional | Aquirer affiliation code |
| `acquirerTid` | `*string` | Optional | Acquirer TID |
| `acquirerNsu` | `*string` | Optional | Acquirer NSU |
| `acquirerAuthCode` | `types.Optional[string]` | Optional | Acquirer authorization code |
| `operationType` | `types.Optional[string]` | Optional | Operation type |
| `card` | [`types.Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | Card data |
| `acquirerMessage` | `types.Optional[string]` | Optional | Acquirer message |
| `acquirerReturnCode` | `types.Optional[string]` | Optional | Acquirer Return Code |
| `installments` | `types.Optional[int]` | Optional | Number of installments |
| `threedAuthenticationUrl` | `types.Optional[string]` | Optional | 3D-S authentication Url |
| `fundingSource` | `types.Optional[string]` | Optional | Identify when a card is prepaid, credit or debit. |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id8",
  "amount": 0,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "transaction_type": "credit_card",
  "statement_descriptor": "statement_descriptor0",
  "acquirer_name": "acquirer_name4",
  "acquirer_affiliation_code": "acquirer_affiliation_code8",
  "acquirer_tid": "acquirer_tid0",
  "acquirer_nsu": "acquirer_nsu0"
}
```

