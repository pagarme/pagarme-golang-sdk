
# Get Split Response

Split response

## Structure

`GetSplitResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mType` | `types.Optional[string]` | Optional | Type |
| `amount` | `types.Optional[int]` | Optional | Amount |
| `recipient` | [`types.Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `gatewayId` | `types.Optional[string]` | Optional | The split rule gateway id |
| `options` | [`types.Optional[models.GetSplitOptionsResponse]`](../../doc/models/get-split-options-response.md) | Optional | - |
| `id` | `types.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "type": "type0",
  "amount": 46,
  "recipient": {
    "id": "id8",
    "name": "name8",
    "email": "email8",
    "document": "document8",
    "description": "description2"
  },
  "gateway_id": "gateway_id0",
  "options": {
    "liable": false,
    "charge_processing_fee": false,
    "charge_remainder_fee": "charge_remainder_fee0"
  }
}
```

