
# Get Safety Pay Transaction Response

Response object for getting a safety pay transaction

## Structure

`GetSafetyPayTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `url` | `types.Optional[string]` | Optional | Payment url |
| `bankTid` | `types.Optional[string]` | Optional | Transaction identifier on bank |
| `paidAt` | `types.Optional[time.Time]` | Optional | Payment date |
| `paidAmount` | `types.Optional[int]` | Optional | Paid amount |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id6",
  "amount": 62,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "transaction_type": "safetypay",
  "url": "url4",
  "bank_tid": "bank_tid4",
  "paid_at": "2016-03-13T12:52:32.123Z",
  "paid_amount": 210
}
```
