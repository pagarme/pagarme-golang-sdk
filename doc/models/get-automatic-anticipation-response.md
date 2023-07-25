
# Get Automatic Anticipation Response

## Structure

`GetAutomaticAnticipationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `enabled` | `types.Optional[bool]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | - |
| `volumePercentage` | `types.Optional[int]` | Optional | - |
| `delay` | `types.Optional[int]` | Optional | - |
| `days` | `types.Optional[[]int]` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "type0",
  "volume_percentage": 62,
  "delay": 228,
  "days": [
    188,
    189,
    190
  ]
}
```

