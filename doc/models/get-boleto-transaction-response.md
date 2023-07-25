
# Get Boleto Transaction Response

Response object for getting a boleto transaction

## Structure

`GetBoletoTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `url` | `types.Optional[string]` | Optional | - |
| `barcode` | `types.Optional[string]` | Optional | - |
| `nossoNumero` | `types.Optional[string]` | Optional | - |
| `bank` | `types.Optional[string]` | Optional | - |
| `documentNumber` | `types.Optional[string]` | Optional | - |
| `instructions` | `types.Optional[string]` | Optional | - |
| `billingAddress` | [`types.Optional[models.GetBillingAddressResponse]`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `dueAt` | `types.Optional[time.Time]` | Optional | - |
| `qrCode` | `types.Optional[string]` | Optional | - |
| `line` | `types.Optional[string]` | Optional | - |
| `pdfPassword` | `types.Optional[string]` | Optional | - |
| `pdf` | `types.Optional[string]` | Optional | - |
| `paidAt` | `types.Optional[time.Time]` | Optional | - |
| `paidAmount` | `types.Optional[string]` | Optional | - |
| `mType` | `types.Optional[string]` | Optional | - |
| `creditAt` | `types.Optional[time.Time]` | Optional | - |
| `statementDescriptor` | `types.Optional[string]` | Optional | Soft Descriptor |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id4",
  "amount": 250,
  "status": "status2",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "transaction_type": "boleto",
  "url": "url4",
  "barcode": "barcode0",
  "nosso_numero": "nosso_numero0",
  "bank": "bank8",
  "document_number": "document_number6"
}
```

