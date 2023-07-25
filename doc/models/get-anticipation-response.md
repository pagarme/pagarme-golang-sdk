
# Get Anticipation Response

Anticipation

## Structure

`GetAnticipationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | Id |
| `requestedAmount` | `types.Optional[int]` | Optional | Requested amount |
| `approvedAmount` | `types.Optional[int]` | Optional | Approved amount |
| `recipient` | [`types.Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `pgid` | `types.Optional[string]` | Optional | Anticipation id on the gateway |
| `createdAt` | `types.Optional[time.Time]` | Optional | Creation date |
| `updatedAt` | `types.Optional[time.Time]` | Optional | Last update date |
| `paymentDate` | `types.Optional[time.Time]` | Optional | Payment date |
| `status` | `types.Optional[string]` | Optional | Status |
| `timeframe` | `types.Optional[string]` | Optional | Timeframe |

## Example (as JSON)

```json
{
  "id": "id0",
  "requested_amount": 246,
  "approved_amount": 212,
  "recipient": {
    "id": "id8",
    "name": "name8",
    "email": "email8",
    "document": "document8",
    "description": "description2"
  },
  "pgid": "pgid4"
}
```

