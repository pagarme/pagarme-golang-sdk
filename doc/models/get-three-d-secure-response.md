
# Get Three D Secure Response

3D-S payment authentication response

## Structure

`GetThreeDSecureResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mpi` | `types.Optional[string]` | Optional | MPI Vendor |
| `eci` | `types.Optional[string]` | Optional | Electronic Commerce Indicator (ECI) (Opcional) |
| `cavv` | `types.Optional[string]` | Optional | Online payment cryptogram, definido pelo 3-D Secure. |
| `transactionId` | `types.Optional[string]` | Optional | Identificador da transação (XID) |
| `successUrl` | `types.Optional[string]` | Optional | Url de redirecionamento de sucessso |

## Example (as JSON)

```json
{
  "mpi": "mpi2",
  "eci": "eci0",
  "cavv": "cavv4",
  "transaction_Id": "transaction_Id4",
  "success_url": "success_url2"
}
```

