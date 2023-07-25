
# Get Charge Response

Response object for getting a charge

## Structure

`GetChargeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `code` | `types.Optional[string]` | Optional | - |
| `gatewayId` | `types.Optional[string]` | Optional | - |
| `amount` | `types.Optional[int]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `currency` | `types.Optional[string]` | Optional | - |
| `paymentMethod` | `types.Optional[string]` | Optional | - |
| `dueAt` | `types.Optional[time.Time]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `lastTransaction` | [`types.Optional[models.GetTransactionResponseInterface]`](../../doc/models/get-transaction-response.md) | Optional | - |
| `invoice` | [`types.Optional[models.GetInvoiceResponse]`](../../doc/models/get-invoice-response.md) | Optional | - |
| `order` | [`types.Optional[models.GetOrderResponse]`](../../doc/models/get-order-response.md) | Optional | - |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `paidAt` | `types.Optional[time.Time]` | Optional | - |
| `canceledAt` | `types.Optional[time.Time]` | Optional | - |
| `canceledAmount` | `types.Optional[int]` | Optional | Canceled Amount |
| `paidAmount` | `types.Optional[int]` | Optional | Paid amount |
| `interestAndFinePaid` | `types.Optional[int]` | Optional | interest and fine paid |
| `recurrencyCycle` | `types.Optional[string]` | Optional | Defines whether the card has been used one or more times. |

## Example (as JSON)

```json
{
  "recurrency_cycle": "\"first\" or \"subsequent\"",
  "id": "id0",
  "code": "code8",
  "gateway_id": "gateway_id0",
  "amount": 46,
  "status": "status8"
}
```

