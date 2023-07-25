
# Get Invoice Response

Response object for getting an invoice

## Structure

`GetInvoiceResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `code` | `types.Optional[string]` | Optional | - |
| `url` | `types.Optional[string]` | Optional | - |
| `amount` | `types.Optional[int]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `paymentMethod` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `items` | [`types.Optional[[]models.GetInvoiceItemResponse]`](../../doc/models/get-invoice-item-response.md) | Optional | - |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `charge` | [`types.Optional[models.GetChargeResponse]`](../../doc/models/get-charge-response.md) | Optional | - |
| `installments` | `types.Optional[int]` | Optional | - |
| `billingAddress` | [`types.Optional[models.GetBillingAddressResponse]`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `subscription` | [`types.Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `cycle` | [`types.Optional[models.GetPeriodResponse]`](../../doc/models/get-period-response.md) | Optional | - |
| `shipping` | [`types.Optional[models.GetShippingResponse]`](../../doc/models/get-shipping-response.md) | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `dueAt` | `types.Optional[time.Time]` | Optional | - |
| `canceledAt` | `types.Optional[time.Time]` | Optional | - |
| `billingAt` | `types.Optional[time.Time]` | Optional | - |
| `seenAt` | `types.Optional[time.Time]` | Optional | - |
| `totalDiscount` | `types.Optional[int]` | Optional | Total discounted value |
| `totalIncrement` | `types.Optional[int]` | Optional | Total discounted value |
| `subscriptionId` | `types.Optional[string]` | Optional | Subscription Id |

## Example (as JSON)

```json
{
  "id": "id0",
  "code": "code8",
  "url": "url4",
  "amount": 46,
  "status": "status8"
}
```

