
# Get Balance Response

Balance

## Structure

`GetBalanceResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `currency` | `types.Optional[string]` | Optional | Currency |
| `availableAmount` | `types.Optional[int64]` | Optional | Amount available for transferring |
| `recipient` | [`types.Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `transferredAmount` | `types.Optional[int64]` | Optional | - |
| `waitingFundsAmount` | `types.Optional[int64]` | Optional | - |

## Example (as JSON)

```json
{
  "currency": "currency0",
  "available_amount": 182,
  "recipient": {
    "id": "id8",
    "name": "name8",
    "email": "email8",
    "document": "document8",
    "description": "description2"
  },
  "transferred_amount": 228,
  "waiting_funds_amount": 252
}
```

