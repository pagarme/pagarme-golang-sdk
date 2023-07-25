# Charges

```go
chargesController := client.ChargesController
```

## Class Name

`ChargesController`

## Methods

* [Update Charge Metadata](../../doc/controllers/charges.md#update-charge-metadata)
* [Update Charge Payment Method](../../doc/controllers/charges.md#update-charge-payment-method)
* [Get Charge Transactions](../../doc/controllers/charges.md#get-charge-transactions)
* [Update Charge Due Date](../../doc/controllers/charges.md#update-charge-due-date)
* [Get Charges](../../doc/controllers/charges.md#get-charges)
* [Capture Charge](../../doc/controllers/charges.md#capture-charge)
* [Update Charge Card](../../doc/controllers/charges.md#update-charge-card)
* [Get Charge](../../doc/controllers/charges.md#get-charge)
* [Get Charges Summary](../../doc/controllers/charges.md#get-charges-summary)
* [Retry Charge](../../doc/controllers/charges.md#retry-charge)
* [Cancel Charge](../../doc/controllers/charges.md#cancel-charge)
* [Create Charge](../../doc/controllers/charges.md#create-charge)
* [Confirm Payment](../../doc/controllers/charges.md#confirm-payment)


# Update Charge Metadata

Updates the metadata from a charge

```go
UpdateChargeMetadata(
    chargeId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | The charge id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the charge metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

request := models.UpdateMetadataRequest{
    Metadata: map[string]*string{
"key0" : "metadata3",
},
}

apiResponse, err := chargesController.UpdateChargeMetadata(chargeId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Charge Payment Method

Updates a charge's payment method

```go
UpdateChargePaymentMethod(
    chargeId string,
    request models.UpdateChargePaymentMethodRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`models.UpdateChargePaymentMethodRequest`](../../doc/models/update-charge-payment-method-request.md) | Body, Required | Request for updating the payment method from a charge |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

requestCreditCard := models.CreateCreditCardPaymentRequest{}
requestCreditCardInstallments := 1
requestCreditCard.Installments = &requestCreditCardInstallments
requestCreditCardCapture := true
requestCreditCard.Capture = &requestCreditCardCapture
requestCreditCardRecurrencyCycle := "\"first\" or \"subsequent\""
requestCreditCard.RecurrencyCycle = &requestCreditCardRecurrencyCycle

requestDebitCard := models.CreateDebitCardPaymentRequest{}

requestBoletoBillingAddress := models.CreateAddressRequest{
    Street:       "street8",
    Number:       "number4",
    ZipCode:      "zip_code2",
    Neighborhood: "neighborhood4",
    City:         "city2",
    State:        "state6",
    Country:      "country2",
    Complement:   "complement6",
    Line1:        "line_18",
    Line2:        "line_26",
}

requestBoleto := models.CreateBoletoPaymentRequest{
    Retries:             10,
    Instructions:        "instructions4",
    DocumentNumber:      "document_number0",
    StatementDescriptor: "statement_descriptor6",
    BillingAddress:      requestBoletoBillingAddress,
}

requestVoucher := models.CreateVoucherPaymentRequest{}
requestVoucherRecurrencyCycle := "\"first\" or \"subsequent\""
requestVoucher.RecurrencyCycle = &requestVoucherRecurrencyCycle

requestCash := models.CreateCashPaymentRequest{
    Description: "description6",
    Confirm:     false,
}

requestBankTransfer := models.CreateBankTransferPaymentRequest{
    Bank:    "bank4",
    Retries: 204,
}

requestPrivateLabel := models.CreatePrivateLabelPaymentRequest{}
requestPrivateLabelInstallments := 1
requestPrivateLabel.Installments = &requestPrivateLabelInstallments
requestPrivateLabelCapture := true
requestPrivateLabel.Capture = &requestPrivateLabelCapture
requestPrivateLabelRecurrencyCycle := "\"first\" or \"subsequent\""
requestPrivateLabel.RecurrencyCycle = &requestPrivateLabelRecurrencyCycle

request := models.UpdateChargePaymentMethodRequest{
    UpdateSubscription: false,
    PaymentMethod:      "payment_method4",
    CreditCard:         requestCreditCard,
    DebitCard:          requestDebitCard,
    Boleto:             requestBoleto,
    Voucher:            requestVoucher,
    Cash:               requestCash,
    BankTransfer:       requestBankTransfer,
    PrivateLabel:       requestPrivateLabel,
}

apiResponse, err := chargesController.UpdateChargePaymentMethod(chargeId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charge Transactions

```go
GetChargeTransactions(
    chargeId string,
    page *int,
    size *int) (
    https.ApiResponse[models.ListChargeTransactionsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge Id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |

## Response Type

[`models.ListChargeTransactionsResponse`](../../doc/models/list-charge-transactions-response.md)

## Example Usage

```go
chargeId := "charge_id8"

apiResponse, err := chargesController.GetChargeTransactions(chargeId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Charge Due Date

Updates the due date from a charge

```go
UpdateChargeDueDate(
    chargeId string,
    request models.UpdateChargeDueDateRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge Id |
| `request` | [`models.UpdateChargeDueDateRequest`](../../doc/models/update-charge-due-date-request.md) | Body, Required | Request for updating the due date |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

request := models.UpdateChargeDueDateRequest{}

apiResponse, err := chargesController.UpdateChargeDueDate(chargeId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charges

Lists all charges

```go
GetCharges(
    page *int,
    size *int,
    code *string,
    status *string,
    paymentMethod *string,
    customerId *string,
    orderId *string,
    createdSince *time.Time,
    createdUntil *time.Time) (
    https.ApiResponse[models.ListChargesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Filter for charge's code |
| `status` | `*string` | Query, Optional | Filter for charge's status |
| `paymentMethod` | `*string` | Query, Optional | Filter for charge's payment method |
| `customerId` | `*string` | Query, Optional | Filter for charge's customer id |
| `orderId` | `*string` | Query, Optional | Filter for charge's order id |
| `createdSince` | `*time.Time` | Query, Optional | Filter for the beginning of the range for charge's creation |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for the end of the range for charge's creation |

## Response Type

[`models.ListChargesResponse`](../../doc/models/list-charges-response.md)

## Example Usage

```go
apiResponse, err := chargesController.GetCharges(nil, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Capture Charge

Captures a charge

```go
CaptureCharge(
    chargeId string,
    request *models.CreateCaptureChargeRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`*models.CreateCaptureChargeRequest`](../../doc/models/create-capture-charge-request.md) | Body, Optional | Request for capturing a charge |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

apiResponse, err := chargesController.CaptureCharge(chargeId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Charge Card

Updates the card from a charge

```go
UpdateChargeCard(
    chargeId string,
    request models.UpdateChargeCardRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`models.UpdateChargeCardRequest`](../../doc/models/update-charge-card-request.md) | Body, Required | Request for updating a charge's card |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

requestCard := models.CreateCardRequest{}
requestCardType := "credit"
requestCard.Type = &requestCardType

request := models.UpdateChargeCardRequest{
    UpdateSubscription: false,
    CardId:             "card_id2",
    Recurrence:         false,
    Card:               requestCard,
}

apiResponse, err := chargesController.UpdateChargeCard(chargeId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charge

Get a charge from its id

```go
GetCharge(
    chargeId string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

apiResponse, err := chargesController.GetCharge(chargeId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Charges Summary

```go
GetChargesSummary(
    status string,
    createdSince *time.Time,
    createdUntil *time.Time) (
    https.ApiResponse[models.GetChargesSummaryResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `status` | `string` | Query, Required | - |
| `createdSince` | `*time.Time` | Query, Optional | - |
| `createdUntil` | `*time.Time` | Query, Optional | - |

## Response Type

[`models.GetChargesSummaryResponse`](../../doc/models/get-charges-summary-response.md)

## Example Usage

```go
status := "status8"

apiResponse, err := chargesController.GetChargesSummary(status, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Retry Charge

Retries a charge

```go
RetryCharge(
    chargeId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

apiResponse, err := chargesController.RetryCharge(chargeId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Cancel Charge

Cancel a charge

```go
CancelCharge(
    chargeId string,
    request *models.CreateCancelChargeRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | Charge id |
| `request` | [`*models.CreateCancelChargeRequest`](../../doc/models/create-cancel-charge-request.md) | Body, Optional | Request for cancelling a charge |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

apiResponse, err := chargesController.CancelCharge(chargeId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Charge

Creates a new charge

```go
CreateCharge(
    request models.CreateChargeRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `request` | [`models.CreateChargeRequest`](../../doc/models/create-charge-request.md) | Body, Required | Request for creating a charge |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
requestPayment := models.CreatePaymentRequest{
    PaymentMethod: "payment_method2",
}

request := models.CreateChargeRequest{
    Amount:  242,
    OrderId: "order_id0",
    Payment: requestPayment,
}

apiResponse, err := chargesController.CreateCharge(&request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Confirm Payment

```go
ConfirmPayment(
    chargeId string,
    request *models.CreateConfirmPaymentRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetChargeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargeId` | `string` | Template, Required | - |
| `request` | [`*models.CreateConfirmPaymentRequest`](../../doc/models/create-confirm-payment-request.md) | Body, Optional | Request for confirm payment |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetChargeResponse`](../../doc/models/get-charge-response.md)

## Example Usage

```go
chargeId := "charge_id8"

apiResponse, err := chargesController.ConfirmPayment(chargeId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

