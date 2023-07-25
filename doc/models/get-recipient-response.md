
# Get Recipient Response

Recipient response

## Structure

`GetRecipientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | Id |
| `name` | `types.Optional[string]` | Optional | Name |
| `email` | `types.Optional[string]` | Optional | Email |
| `document` | `types.Optional[string]` | Optional | Document |
| `description` | `types.Optional[string]` | Optional | Description |
| `mType` | `types.Optional[string]` | Optional | Type |
| `status` | `types.Optional[string]` | Optional | Status |
| `createdAt` | `types.Optional[time.Time]` | Optional | Creation date |
| `updatedAt` | `types.Optional[time.Time]` | Optional | Last update date |
| `deletedAt` | `types.Optional[time.Time]` | Optional | Deletion date |
| `defaultBankAccount` | [`types.Optional[models.GetBankAccountResponse]`](../../doc/models/get-bank-account-response.md) | Optional | Default bank account |
| `gatewayRecipients` | [`types.Optional[[]models.GetGatewayRecipientResponse]`](../../doc/models/get-gateway-recipient-response.md) | Optional | Info about the recipient on the gateway |
| `metadata` | `types.Optional[map[string]string]` | Optional | Metadata |
| `automaticAnticipationSettings` | [`types.Optional[models.GetAutomaticAnticipationResponse]`](../../doc/models/get-automatic-anticipation-response.md) | Optional | - |
| `transferSettings` | [`types.Optional[models.GetTransferSettingsResponse]`](../../doc/models/get-transfer-settings-response.md) | Optional | - |
| `code` | `types.Optional[string]` | Optional | Recipient code |
| `paymentMode` | `types.Optional[string]` | Optional | Payment mode<br>**Default**: `"bank_transfer"` |

## Example (as JSON)

```json
{
  "payment_mode": "bank_transfer",
  "id": "id0",
  "name": "name0",
  "email": "email6",
  "document": "document6",
  "description": "description0"
}
```

