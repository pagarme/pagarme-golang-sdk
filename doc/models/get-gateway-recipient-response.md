
# Get Gateway Recipient Response

Information about the recipient on the gateway

## Structure

`GetGatewayRecipientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `gateway` | `types.Optional[string]` | Optional | Gateway name |
| `status` | `types.Optional[string]` | Optional | Status of the recipient on the gateway |
| `pgid` | `types.Optional[string]` | Optional | Recipient id on the gateway |
| `createdAt` | `types.Optional[string]` | Optional | Creation date |
| `updatedAt` | `types.Optional[string]` | Optional | Last update date |

## Example (as JSON)

```json
{
  "gateway": "gateway0",
  "status": "status8",
  "pgid": "pgid4",
  "created_at": "created_at2",
  "updated_at": "updated_at4"
}
```

