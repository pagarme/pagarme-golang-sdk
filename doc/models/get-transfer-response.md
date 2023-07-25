
# Get Transfer Response

Transfer response

## Structure

`GetTransferResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | Id |
| `amount` | `types.Optional[int]` | Optional | Transfer amount |
| `status` | `types.Optional[string]` | Optional | Transfer status |
| `createdAt` | `types.Optional[time.Time]` | Optional | Transfer creation date |
| `updatedAt` | `types.Optional[time.Time]` | Optional | Transfer last update date |
| `bankAccount` | [`types.Optional[models.GetBankAccountResponse]`](../../doc/models/get-bank-account-response.md) | Optional | Bank account |
| `metadata` | `types.Optional[map[string]string]` | Optional | Metadata |

## Example (as JSON)

```json
{
  "id": "id0",
  "amount": 46,
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z"
}
```

