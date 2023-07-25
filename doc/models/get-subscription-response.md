
# Get Subscription Response

## Structure

`GetSubscriptionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `code` | `types.Optional[string]` | Optional | - |
| `startAt` | `types.Optional[time.Time]` | Optional | - |
| `interval` | `types.Optional[string]` | Optional | - |
| `intervalCount` | `types.Optional[int]` | Optional | - |
| `billingType` | `types.Optional[string]` | Optional | - |
| `currentCycle` | [`types.Optional[models.GetPeriodResponse]`](../../doc/models/get-period-response.md) | Optional | - |
| `paymentMethod` | `types.Optional[string]` | Optional | - |
| `currency` | `types.Optional[string]` | Optional | - |
| `installments` | `types.Optional[int]` | Optional | - |
| `status` | `types.Optional[string]` | Optional | - |
| `createdAt` | `types.Optional[time.Time]` | Optional | - |
| `updatedAt` | `types.Optional[time.Time]` | Optional | - |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `card` | [`types.Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | - |
| `items` | [`types.Optional[[]models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | - |
| `statementDescriptor` | `types.Optional[string]` | Optional | - |
| `metadata` | `types.Optional[map[string]string]` | Optional | - |
| `setup` | [`types.Optional[models.GetSetupResponse]`](../../doc/models/get-setup-response.md) | Optional | - |
| `gatewayAffiliationId` | `types.Optional[string]` | Optional | Affiliation Code |
| `nextBillingAt` | `types.Optional[time.Time]` | Optional | - |
| `billingDay` | `types.Optional[int]` | Optional | - |
| `minimumPrice` | `types.Optional[int]` | Optional | - |
| `canceledAt` | `types.Optional[time.Time]` | Optional | - |
| `discounts` | [`types.Optional[[]models.GetDiscountResponse]`](../../doc/models/get-discount-response.md) | Optional | Subscription discounts |
| `increments` | [`types.Optional[[]models.GetIncrementResponse]`](../../doc/models/get-increment-response.md) | Optional | Subscription increments |
| `boletoDueDays` | `types.Optional[int]` | Optional | Days until boleto expires |
| `split` | [`types.Optional[models.GetSubscriptionSplitResponse]`](../../doc/models/get-subscription-split-response.md) | Optional | Subscription's split response |
| `boleto` | [`types.Optional[models.GetSubscriptionBoletoResponse]`](../../doc/models/get-subscription-boleto-response.md) | Optional | - |
| `manualBilling` | `types.Optional[bool]` | Optional | - |

## Example (as JSON)

```json
{
  "boleto": {
    "interest": {
      "days": 2,
      "type": "percentage",
      "amount": 20
    },
    "fine": {
      "days": 2,
      "type": "flat",
      "amount": 10
    },
    "max_days_to_pay_past_due": 2
  },
  "id": "id0",
  "code": "code8",
  "start_at": "2016-03-13T12:52:32.123Z",
  "interval": "interval2",
  "interval_count": 82
}
```

