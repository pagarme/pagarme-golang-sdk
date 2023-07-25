
# Create Credit Card Payment Request

The settings for creating a credit card payment

## Structure

`CreateCreditCardPaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `installments` | `*int` | Optional | Number of installments<br>**Default**: `1` |
| `statementDescriptor` | `*string` | Optional | The text that will be shown on the credit card's statement |
| `card` | [`*models.CreateCardRequest`](../../doc/models/create-card-request.md) | Optional | Credit card data |
| `cardId` | `*string` | Optional | The credit card id |
| `cardToken` | `*string` | Optional | - |
| `recurrence` | `*bool` | Optional | Indicates a recurrence |
| `capture` | `*bool` | Optional | Indicates if the operation should be only authorization or auth and capture.<br>**Default**: `true` |
| `extendedLimitEnabled` | `*bool` | Optional | Indicates whether the extended label (private label) is enabled |
| `extendedLimitCode` | `*string` | Optional | Extended Limit Code |
| `merchantCategoryCode` | `*int64` | Optional | Customer business segment code |
| `authentication` | [`*models.CreatePaymentAuthenticationRequest`](../../doc/models/create-payment-authentication-request.md) | Optional | The payment authentication request |
| `contactless` | [`*models.CreateCardPaymentContactlessRequest`](../../doc/models/create-card-payment-contactless-request.md) | Optional | The Credit card payment contactless request |
| `autoRecovery` | `*bool` | Optional | Indicates whether a particular payment will enter the offline retry flow |
| `operationType` | `*string` | Optional | AuthOnly, AuthAndCapture, PreAuth |
| `recurrencyCycle` | `*string` | Optional | Defines whether the card has been used one or more times. |

## Example (as JSON)

```json
{
  "installments": 1,
  "capture": true,
  "recurrency_cycle": "\"first\" or \"subsequent\"",
  "statement_descriptor": "statement_descriptor0",
  "card": {
    "number": "number6",
    "holder_name": "holder_name2",
    "exp_month": 228,
    "exp_year": 68,
    "cvv": "cvv4"
  },
  "card_id": "card_id4",
  "card_token": "card_token0"
}
```

