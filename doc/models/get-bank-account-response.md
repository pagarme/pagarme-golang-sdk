
# Get Bank Account Response

## Structure

`GetBankAccountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | Id |
| `holderName` | `types.Optional[string]` | Optional | Holder name |
| `holderType` | `types.Optional[string]` | Optional | Holder type |
| `bank` | `types.Optional[string]` | Optional | Bank |
| `branchNumber` | `types.Optional[string]` | Optional | Branch number |
| `branchCheckDigit` | `types.Optional[string]` | Optional | Branch check digit |
| `accountNumber` | `types.Optional[string]` | Optional | Account number |
| `accountCheckDigit` | `types.Optional[string]` | Optional | Account check digit |
| `mType` | `types.Optional[string]` | Optional | Bank account type |
| `status` | `types.Optional[string]` | Optional | Bank account status |
| `createdAt` | `types.Optional[time.Time]` | Optional | Creation date |
| `updatedAt` | `types.Optional[time.Time]` | Optional | Last update date |
| `deletedAt` | `types.Optional[time.Time]` | Optional | Deletion date |
| `recipient` | [`types.Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `metadata` | `types.Optional[map[string]string]` | Optional | Metadata |
| `pixKey` | `types.Optional[string]` | Optional | Pix Key |

## Example (as JSON)

```json
{
  "id": "id0",
  "holder_name": "holder_name4",
  "holder_type": "holder_type2",
  "bank": "bank8",
  "branch_number": "branch_number6"
}
```

