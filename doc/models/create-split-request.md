
# Create Split Request

Split

## Structure

`CreateSplitRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mType` | `string` | Required | Split type |
| `amount` | `int` | Required | Amount |
| `recipientId` | `string` | Required | Recipient id |
| `options` | [`*models.CreateSplitOptionsRequest`](../../doc/models/create-split-options-request.md) | Optional | The split options request |
| `splitRuleId` | `*string` | Optional | Rule code used in cancellation. |

## Example (as JSON)

```json
{
  "type": "type0",
  "amount": 46,
  "recipient_id": "recipient_id0",
  "options": {
    "liable": false,
    "charge_processing_fee": false,
    "charge_remainder_fee": false
  },
  "split_rule_id": "split_rule_id2"
}
```

