
# Create Boleto Payment Request

Contains the settings for creating a boleto payment

## Structure

`CreateBoletoPaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `retries` | `int` | Required | Number of retries |
| `bank` | `types.Optional[string]` | Optional | The bank code, containing three characters. The available codes are on the API specification |
| `instructions` | `string` | Required | The instructions field that will be printed on the boleto. |
| `dueAt` | `types.Optional[time.Time]` | Optional | Boleto due date |
| `billingAddress` | [`models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Required | Card's billing address |
| `billingAddressId` | `types.Optional[string]` | Optional | The address id for the billing address |
| `nossoNumero` | `types.Optional[string]` | Optional | Customer identification number with the bank |
| `documentNumber` | `string` | Required | Boleto identification |
| `statementDescriptor` | `string` | Required | Soft Descriptor |
| `interest` | [`types.Optional[models.CreateInterestRequest]`](../../doc/models/create-interest-request.md) | Optional | - |
| `fine` | [`types.Optional[models.CreateFineRequest]`](../../doc/models/create-fine-request.md) | Optional | - |
| `maxDaysToPayPastDue` | `types.Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "retries": 230,
  "bank": "bank8",
  "instructions": "instructions2",
  "due_at": "2016-03-13T12:52:32.123Z",
  "billing_address": {
    "street": "street8",
    "number": "number4",
    "zip_code": "zip_code2",
    "neighborhood": "neighborhood4",
    "city": "city2",
    "state": "state6",
    "country": "country2",
    "complement": "complement6",
    "metadata": {
      "key0": "metadata5",
      "key1": "metadata6"
    },
    "line_1": "line_18",
    "line_2": "line_26"
  },
  "billing_address_id": "billing_address_id6",
  "nosso_numero": "nosso_numero0",
  "document_number": "document_number6",
  "statement_descriptor": "statement_descriptor0",
  "interest": {
    "days": 156,
    "type": "type0",
    "amount": 230
  }
}
```

