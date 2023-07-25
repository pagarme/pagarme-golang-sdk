
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `serviceRefererName` | `string` |  |
| `httpConfiguration` | `https.HttpConfiguration` | Configurable http client options. |
| `basicAuthUserName` | `string` | The username to use with basic authentication |
| `basicAuthPassword` | `string` | The password to use with basic authentication |

The API client can be initialized as follows:

```go
config := pagarmeapisdk.CreateConfiguration(
    pagarmeapisdk.WithServiceRefererName("ServiceRefererName"),
    pagarmeapisdk.WithBasicAuthUserName("BasicAuthUserName"),
    pagarmeapisdk.WithBasicAuthPassword("BasicAuthPassword"),
)
client := pagarmeapisdk.NewClient(config)
```

## PagarmeApiSDK Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| plans | Gets PlansController |
| subscriptions | Gets SubscriptionsController |
| invoices | Gets InvoicesController |
| orders | Gets OrdersController |
| customers | Gets CustomersController |
| recipients | Gets RecipientsController |
| charges | Gets ChargesController |
| transfers | Gets TransfersController |
| tokens | Gets TokensController |
| transactions | Gets TransactionsController |
| payables | Gets PayablesController |
| balanceOperations | Gets BalanceOperationsController |

