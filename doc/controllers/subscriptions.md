# Subscriptions

```go
subscriptionsController := client.SubscriptionsController
```

## Class Name

`SubscriptionsController`

## Methods

* [Renew Subscription](../../doc/controllers/subscriptions.md#renew-subscription)
* [Update Subscription Card](../../doc/controllers/subscriptions.md#update-subscription-card)
* [Delete Usage](../../doc/controllers/subscriptions.md#delete-usage)
* [Create Discount](../../doc/controllers/subscriptions.md#create-discount)
* [Create an Usage](../../doc/controllers/subscriptions.md#create-an-usage)
* [Update Current Cycle Status](../../doc/controllers/subscriptions.md#update-current-cycle-status)
* [Delete Discount](../../doc/controllers/subscriptions.md#delete-discount)
* [Get Subscription Items](../../doc/controllers/subscriptions.md#get-subscription-items)
* [Update Subscription Payment Method](../../doc/controllers/subscriptions.md#update-subscription-payment-method)
* [Get Subscription Item](../../doc/controllers/subscriptions.md#get-subscription-item)
* [Get Subscriptions](../../doc/controllers/subscriptions.md#get-subscriptions)
* [Cancel Subscription](../../doc/controllers/subscriptions.md#cancel-subscription)
* [Create Increment](../../doc/controllers/subscriptions.md#create-increment)
* [Create Usage](../../doc/controllers/subscriptions.md#create-usage)
* [Get Discount by Id](../../doc/controllers/subscriptions.md#get-discount-by-id)
* [Create Subscription](../../doc/controllers/subscriptions.md#create-subscription)
* [Get Increment by Id](../../doc/controllers/subscriptions.md#get-increment-by-id)
* [Update Subscription Affiliation Id](../../doc/controllers/subscriptions.md#update-subscription-affiliation-id)
* [Update Subscription Metadata](../../doc/controllers/subscriptions.md#update-subscription-metadata)
* [Delete Increment](../../doc/controllers/subscriptions.md#delete-increment)
* [Get Subscription Cycles](../../doc/controllers/subscriptions.md#get-subscription-cycles)
* [Get Discounts](../../doc/controllers/subscriptions.md#get-discounts)
* [Update Subscription Billing Date](../../doc/controllers/subscriptions.md#update-subscription-billing-date)
* [Delete Subscription Item](../../doc/controllers/subscriptions.md#delete-subscription-item)
* [Get Increments](../../doc/controllers/subscriptions.md#get-increments)
* [Update Subscription Due Days](../../doc/controllers/subscriptions.md#update-subscription-due-days)
* [Update Subscription Start At](../../doc/controllers/subscriptions.md#update-subscription-start-at)
* [Update Subscription Item](../../doc/controllers/subscriptions.md#update-subscription-item)
* [Create Subscription Item](../../doc/controllers/subscriptions.md#create-subscription-item)
* [Get Subscription](../../doc/controllers/subscriptions.md#get-subscription)
* [Get Usages](../../doc/controllers/subscriptions.md#get-usages)
* [Update Latest Period End At](../../doc/controllers/subscriptions.md#update-latest-period-end-at)
* [Update Subscription Minium Price](../../doc/controllers/subscriptions.md#update-subscription-minium-price)
* [Get Subscription Cycle by Id](../../doc/controllers/subscriptions.md#get-subscription-cycle-by-id)
* [Get Usage Report](../../doc/controllers/subscriptions.md#get-usage-report)
* [Update Split Subscription](../../doc/controllers/subscriptions.md#update-split-subscription)


# Renew Subscription

```go
RenewSubscription(
    subscriptionId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetPeriodResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | - |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPeriodResponse`](../../doc/models/get-period-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.RenewSubscription(subscriptionId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Card

Updates the credit card from a subscription

```go
UpdateSubscriptionCard(
    subscriptionId string,
    request models.UpdateSubscriptionCardRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.UpdateSubscriptionCardRequest`](../../doc/models/update-subscription-card-request.md) | Body, Required | Request for updating a card |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

requestCard := models.CreateCardRequest{}
requestCardType := "credit"
requestCard.Type = &requestCardType

request := models.UpdateSubscriptionCardRequest{
    CardId: "card_id2",
    Card:   requestCard,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionCard(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Usage

Deletes a usage

```go
DeleteUsage(
    subscriptionId string,
    itemId string,
    usageId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetUsageResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `itemId` | `string` | Template, Required | The subscription item id |
| `usageId` | `string` | Template, Required | The usage id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetUsageResponse`](../../doc/models/get-usage-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
itemId := "item_id0"
usageId := "usage_id0"

apiResponse, err := subscriptionsController.DeleteUsage(subscriptionId, itemId, usageId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Discount

Creates a discount

```go
CreateDiscount(
    subscriptionId string,
    request models.CreateDiscountRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetDiscountResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.CreateDiscountRequest`](../../doc/models/create-discount-request.md) | Body, Required | Request for creating a discount |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetDiscountResponse`](../../doc/models/get-discount-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.CreateDiscountRequest{
    Value:        185.28,
    DiscountType: "discount_type4",
    ItemId:       "item_id6",
}

apiResponse, err := subscriptionsController.CreateDiscount(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create an Usage

Create Usage

```go
CreateAnUsage(
    subscriptionId string,
    itemId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetUsageResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `itemId` | `string` | Template, Required | Item id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetUsageResponse`](../../doc/models/get-usage-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
itemId := "item_id0"

apiResponse, err := subscriptionsController.CreateAnUsage(subscriptionId, itemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Current Cycle Status

```go
UpdateCurrentCycleStatus(
    subscriptionId string,
    request models.UpdateCurrentCycleStatusRequest,
    idempotencyKey *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `request` | [`models.UpdateCurrentCycleStatusRequest`](../../doc/models/update-current-cycle-status-request.md) | Body, Required | Request for updating the end date of the subscription current status |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

``

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.UpdateCurrentCycleStatusRequest{
    Status: "status8",
}

resp, err := subscriptionsController.UpdateCurrentCycleStatus(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Delete Discount

Deletes a discount

```go
DeleteDiscount(
    subscriptionId string,
    discountId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetDiscountResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `discountId` | `string` | Template, Required | Discount Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetDiscountResponse`](../../doc/models/get-discount-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
discountId := "discount_id8"

apiResponse, err := subscriptionsController.DeleteDiscount(subscriptionId, discountId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Items

Get Subscription Items

```go
GetSubscriptionItems(
    subscriptionId string,
    page *int,
    size *int,
    name *string,
    code *string,
    status *string,
    description *string,
    createdSince *string,
    createdUntil *string) (
    https.ApiResponse[models.ListSubscriptionItemsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `name` | `*string` | Query, Optional | The item name |
| `code` | `*string` | Query, Optional | Identification code in the client system |
| `status` | `*string` | Query, Optional | The item statis |
| `description` | `*string` | Query, Optional | The item description |
| `createdSince` | `*string` | Query, Optional | Filter for item's creation date start range |
| `createdUntil` | `*string` | Query, Optional | Filter for item's creation date end range |

## Response Type

[`models.ListSubscriptionItemsResponse`](../../doc/models/list-subscription-items-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.GetSubscriptionItems(subscriptionId, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Payment Method

Updates the payment method from a subscription

```go
UpdateSubscriptionPaymentMethod(
    subscriptionId string,
    request models.UpdateSubscriptionPaymentMethodRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.UpdateSubscriptionPaymentMethodRequest`](../../doc/models/update-subscription-payment-method-request.md) | Body, Required | Request for updating the paymentmethod from a subscription |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

requestCard := models.CreateCardRequest{}
requestCardType := "credit"
requestCard.Type = &requestCardType

request := models.UpdateSubscriptionPaymentMethodRequest{
    PaymentMethod: "payment_method4",
    CardId:        "card_id2",
    Card:          requestCard,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionPaymentMethod(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Item

Get Subscription Item

```go
GetSubscriptionItem(
    subscriptionId string,
    itemId string) (
    https.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `itemId` | `string` | Template, Required | Item id |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
itemId := "item_id0"

apiResponse, err := subscriptionsController.GetSubscriptionItem(subscriptionId, itemId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscriptions

Gets all subscriptions

```go
GetSubscriptions(
    page *int,
    size *int,
    code *string,
    billingType *string,
    customerId *string,
    planId *string,
    cardId *string,
    status *string,
    nextBillingSince *time.Time,
    nextBillingUntil *time.Time,
    createdSince *time.Time,
    createdUntil *time.Time) (
    https.ApiResponse[models.ListSubscriptionsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Filter for subscription's code |
| `billingType` | `*string` | Query, Optional | Filter for subscription's billing type |
| `customerId` | `*string` | Query, Optional | Filter for subscription's customer id |
| `planId` | `*string` | Query, Optional | Filter for subscription's plan id |
| `cardId` | `*string` | Query, Optional | Filter for subscription's card id |
| `status` | `*string` | Query, Optional | Filter for subscription's status |
| `nextBillingSince` | `*time.Time` | Query, Optional | Filter for subscription's next billing date start range |
| `nextBillingUntil` | `*time.Time` | Query, Optional | Filter for subscription's next billing date end range |
| `createdSince` | `*time.Time` | Query, Optional | Filter for subscription's creation date start range |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for subscriptions creation date end range |

## Response Type

[`models.ListSubscriptionsResponse`](../../doc/models/list-subscriptions-response.md)

## Example Usage

```go
apiResponse, err := subscriptionsController.GetSubscriptions(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Cancel Subscription

Cancels a subscription

```go
CancelSubscription(
    subscriptionId string,
    request *models.CreateCancelSubscriptionRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`*models.CreateCancelSubscriptionRequest`](../../doc/models/create-cancel-subscription-request.md) | Body, Optional | Request for cancelling a subscription |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.CreateCancelSubscriptionRequest{
    CancelPendingInvoices: true,
}

apiResponse, err := subscriptionsController.CancelSubscription(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Increment

Creates a increment

```go
CreateIncrement(
    subscriptionId string,
    request models.CreateIncrementRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetIncrementResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.CreateIncrementRequest`](../../doc/models/create-increment-request.md) | Body, Required | Request for creating a increment |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetIncrementResponse`](../../doc/models/get-increment-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.CreateIncrementRequest{
    Value:         185.28,
    IncrementType: "increment_type8",
    ItemId:        "item_id6",
}

apiResponse, err := subscriptionsController.CreateIncrement(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Usage

Creates a usage

```go
CreateUsage(
    subscriptionId string,
    itemId string,
    body models.CreateUsageRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetUsageResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `itemId` | `string` | Template, Required | Item id |
| `body` | [`models.CreateUsageRequest`](../../doc/models/create-usage-request.md) | Body, Required | Request for creating a usage |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetUsageResponse`](../../doc/models/get-usage-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
itemId := "item_id0"

bodyUsedAt, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
body := models.CreateUsageRequest{
    Quantity:    156,
    Description: "description4",
    UsedAt:      bodyUsedAt,
}

apiResponse, err := subscriptionsController.CreateUsage(subscriptionId, itemId, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Discount by Id

```go
GetDiscountById(
    subscriptionId string,
    discountId string) (
    https.ApiResponse[models.GetDiscountResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `discountId` | `string` | Template, Required | - |

## Response Type

[`models.GetDiscountResponse`](../../doc/models/get-discount-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
discountId := "discountId0"

apiResponse, err := subscriptionsController.GetDiscountById(subscriptionId, discountId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Subscription

Creates a new subscription

```go
CreateSubscription(
    body models.CreateSubscriptionRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`models.CreateSubscriptionRequest`](../../doc/models/create-subscription-request.md) | Body, Required | Request for creating a subscription |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
bodyCustomerAddress := models.CreateAddressRequest{
    Street:       "street0",
    Number:       "number8",
    ZipCode:      "zip_code4",
    Neighborhood: "neighborhood6",
    City:         "city0",
    State:        "state6",
    Country:      "country4",
    Complement:   "complement6",
    Line1:        "line_16",
    Line2:        "line_28",
}

bodyCustomerPhones := models.CreatePhonesRequest{}

bodyCustomer := models.CreateCustomerRequest{
    Name:     "{\n    \"name\": \"Tony Stark\"\n}",
    Email:    "email2",
    Document: "document2",
    Type:     "type6",
    Metadata: map[string]*string{
"key0" : "metadata9",
"key1" : "metadata0",
},
    Code:     "code2",
    Address:  bodyCustomerAddress,
    Phones:   bodyCustomerPhones,
}

bodyCard := models.CreateCardRequest{}
bodyCardType := "credit"
bodyCard.Type = &bodyCardType

bodyPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType: "scheme_type2",
}

bodyItems0PricingScheme := models.CreatePricingSchemeRequest{
    SchemeType: "scheme_type5",
}


bodyItems0Discounts0 := models.CreateDiscountRequest{
    Value:        65.46,
    DiscountType: "discount_type2",
    ItemId:       "item_id4",
}
bodyItems0Discounts := []models.CreateDiscountRequest{bodyItems0Discounts0}
bodyItems0 := models.CreateSubscriptionItemRequest{
    Description:   "description3",
    Id:            "id3",
    PlanItemId:    "plan_item_id3",
    Name:          "name3",
    PricingScheme: bodyItems0PricingScheme,
    Discounts:     bodyItems0Discounts,
}

bodyItems := []models.CreateSubscriptionItemRequest{bodyItems0}
bodyShippingAddress := models.CreateAddressRequest{
    Street:       "street6",
    Number:       "number4",
    ZipCode:      "zip_code0",
    Neighborhood: "neighborhood2",
    City:         "city6",
    State:        "state2",
    Country:      "country0",
    Complement:   "complement2",
    Line1:        "line_10",
    Line2:        "line_24",
}

bodyShipping := models.CreateShippingRequest{
    Amount:         140,
    Description:    "description0",
    RecipientName:  "recipient_name8",
    RecipientPhone: "recipient_phone2",
    AddressId:      "address_id0",
    Type:           "type0",
    Address:        bodyShippingAddress,
}


bodyDiscounts0 := models.CreateDiscountRequest{
    Value:        95.59,
    DiscountType: "discount_type5",
    ItemId:       "item_id7",
}
bodyDiscounts := []models.CreateDiscountRequest{bodyDiscounts0}

bodyIncrements0 := models.CreateIncrementRequest{
    Value:         38.83,
    IncrementType: "increment_type3",
    ItemId:        "item_id9",
}
bodyIncrements := []models.CreateIncrementRequest{bodyIncrements0}
body := models.CreateSubscriptionRequest{
    Code:                "code4",
    PaymentMethod:       "payment_method4",
    BillingType:         "billing_type0",
    StatementDescriptor: "statement_descriptor6",
    Description:         "description4",
    Currency:            "currency6",
    Interval:            "interval6",
    IntervalCount:       170,
    Metadata:            map[string]*string{
"key0" : "metadata7",
"key1" : "metadata8",
},
    Customer:            bodyCustomer,
    Card:                bodyCard,
    PricingScheme:       bodyPricingScheme,
    Items:               bodyItems,
    Shipping:            bodyShipping,
    Discounts:           bodyDiscounts,
    Increments:          bodyIncrements,
}

apiResponse, err := subscriptionsController.CreateSubscription(&body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Increment by Id

```go
GetIncrementById(
    subscriptionId string,
    incrementId string) (
    https.ApiResponse[models.GetIncrementResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription Id |
| `incrementId` | `string` | Template, Required | The increment Id |

## Response Type

[`models.GetIncrementResponse`](../../doc/models/get-increment-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
incrementId := "increment_id8"

apiResponse, err := subscriptionsController.GetIncrementById(subscriptionId, incrementId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Affiliation Id

```go
UpdateSubscriptionAffiliationId(
    subscriptionId string,
    request models.UpdateSubscriptionAffiliationIdRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | - |
| `request` | [`models.UpdateSubscriptionAffiliationIdRequest`](../../doc/models/update-subscription-affiliation-id-request.md) | Body, Required | Request for updating a subscription affiliation id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.UpdateSubscriptionAffiliationIdRequest{
    GatewayAffiliationId: "gateway_affiliation_id2",
}

apiResponse, err := subscriptionsController.UpdateSubscriptionAffiliationId(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Metadata

Updates the metadata from a subscription

```go
UpdateSubscriptionMetadata(
    subscriptionId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the subscrption metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.UpdateMetadataRequest{
    Metadata: map[string]*string{
"key0" : "metadata3",
},
}

apiResponse, err := subscriptionsController.UpdateSubscriptionMetadata(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Increment

Deletes a increment

```go
DeleteIncrement(
    subscriptionId string,
    incrementId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetIncrementResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `incrementId` | `string` | Template, Required | Increment id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetIncrementResponse`](../../doc/models/get-increment-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
incrementId := "increment_id8"

apiResponse, err := subscriptionsController.DeleteIncrement(subscriptionId, incrementId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Cycles

```go
GetSubscriptionCycles(
    subscriptionId string,
    page string,
    size string) (
    https.ApiResponse[models.ListCyclesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `page` | `string` | Query, Required | Page number |
| `size` | `string` | Query, Required | Page size |

## Response Type

[`models.ListCyclesResponse`](../../doc/models/list-cycles-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
page := "page8"
size := "size0"

apiResponse, err := subscriptionsController.GetSubscriptionCycles(subscriptionId, page, size)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Discounts

```go
GetDiscounts(
    subscriptionId string,
    page int,
    size int) (
    https.ApiResponse[models.ListDiscountsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `page` | `int` | Query, Required | Page number |
| `size` | `int` | Query, Required | Page size |

## Response Type

[`models.ListDiscountsResponse`](../../doc/models/list-discounts-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
page := 30
size := 18

apiResponse, err := subscriptionsController.GetDiscounts(subscriptionId, page, size)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Billing Date

Updates the billing date from a subscription

```go
UpdateSubscriptionBillingDate(
    subscriptionId string,
    request models.UpdateSubscriptionBillingDateRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `request` | [`models.UpdateSubscriptionBillingDateRequest`](../../doc/models/update-subscription-billing-date-request.md) | Body, Required | Request for updating the subscription billing date |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

requestNextBillingAt, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
request := models.UpdateSubscriptionBillingDateRequest{
    NextBillingAt: requestNextBillingAt,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionBillingDate(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Subscription Item

Deletes a subscription item

```go
DeleteSubscriptionItem(
    subscriptionId string,
    subscriptionItemId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `subscriptionItemId` | `string` | Template, Required | Subscription item id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
subscriptionItemId := "subscription_item_id4"

apiResponse, err := subscriptionsController.DeleteSubscriptionItem(subscriptionId, subscriptionItemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Increments

```go
GetIncrements(
    subscriptionId string,
    page *int,
    size *int) (
    https.ApiResponse[models.ListIncrementsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |

## Response Type

[`models.ListIncrementsResponse`](../../doc/models/list-increments-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.GetIncrements(subscriptionId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Due Days

Updates the boleto due days from a subscription

```go
UpdateSubscriptionDueDays(
    subscriptionId string,
    request models.UpdateSubscriptionDueDaysRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `request` | [`models.UpdateSubscriptionDueDaysRequest`](../../doc/models/update-subscription-due-days-request.md) | Body, Required | - |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.UpdateSubscriptionDueDaysRequest{
    BoletoDueDays: 226,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionDueDays(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Start At

Updates the start at date from a subscription

```go
UpdateSubscriptionStartAt(
    subscriptionId string,
    request models.UpdateSubscriptionStartAtRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `request` | [`models.UpdateSubscriptionStartAtRequest`](../../doc/models/update-subscription-start-at-request.md) | Body, Required | Request for updating the subscription start date |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

requestStartAt, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
request := models.UpdateSubscriptionStartAtRequest{
    StartAt: requestStartAt,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionStartAt(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Item

Updates a subscription item

```go
UpdateSubscriptionItem(
    subscriptionId string,
    itemId string,
    body models.UpdateSubscriptionItemRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `itemId` | `string` | Template, Required | Item id |
| `body` | [`models.UpdateSubscriptionItemRequest`](../../doc/models/update-subscription-item-request.md) | Body, Required | Request for updating a subscription item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
itemId := "item_id0"

bodyPricingSchemePriceBrackets0 := models.UpdatePriceBracketRequest{
    StartQuantity: 31,
    Price:         225,
}

bodyPricingSchemePriceBrackets := []models.UpdatePriceBracketRequest{bodyPricingSchemePriceBrackets0}
bodyPricingScheme := models.UpdatePricingSchemeRequest{
    SchemeType:    "scheme_type2",
    PriceBrackets: bodyPricingSchemePriceBrackets,
}

body := models.UpdateSubscriptionItemRequest{
    Description:   "description4",
    Status:        "status2",
    Name:          "name6",
    PricingScheme: bodyPricingScheme,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionItem(subscriptionId, itemId, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Subscription Item

Creates a new Subscription item

```go
CreateSubscriptionItem(
    subscriptionId string,
    request models.CreateSubscriptionItemRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.CreateSubscriptionItemRequest`](../../doc/models/create-subscription-item-request.md) | Body, Required | Request for creating a subscription item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
requestPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType: "scheme_type2",
}


requestDiscounts0 := models.CreateDiscountRequest{
    Value:        199.99,
    DiscountType: "discount_type5",
    ItemId:       "item_id7",
}
requestDiscounts := []models.CreateDiscountRequest{requestDiscounts0}
request := models.CreateSubscriptionItemRequest{
    Description:   "description6",
    Id:            "id6",
    PlanItemId:    "plan_item_id6",
    Name:          "name6",
    PricingScheme: requestPricingScheme,
    Discounts:     requestDiscounts,
}


apiResponse, err := subscriptionsController.CreateSubscriptionItem(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription

Gets a subscription

```go
GetSubscription(
    subscriptionId string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.GetSubscription(subscriptionId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Usages

Lists all usages from a subscription item

```go
GetUsages(
    subscriptionId string,
    itemId string,
    page *int,
    size *int,
    code *string,
    group *string,
    usedSince *time.Time,
    usedUntil *time.Time) (
    https.ApiResponse[models.ListUsagesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `itemId` | `string` | Template, Required | The subscription item id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Identification code in the client system |
| `group` | `*string` | Query, Optional | Identification group in the client system |
| `usedSince` | `*time.Time` | Query, Optional | - |
| `usedUntil` | `*time.Time` | Query, Optional | - |

## Response Type

[`models.ListUsagesResponse`](../../doc/models/list-usages-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
itemId := "item_id0"

apiResponse, err := subscriptionsController.GetUsages(subscriptionId, itemId, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Latest Period End At

```go
UpdateLatestPeriodEndAt(
    subscriptionId string,
    request models.UpdateCurrentCycleEndDateRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | - |
| `request` | [`models.UpdateCurrentCycleEndDateRequest`](../../doc/models/update-current-cycle-end-date-request.md) | Body, Required | Request for updating the end date of the current signature cycle |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.UpdateCurrentCycleEndDateRequest{}

apiResponse, err := subscriptionsController.UpdateLatestPeriodEndAt(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Minium Price

Atualização do valor mínimo da assinatura

```go
UpdateSubscriptionMiniumPrice(
    subscriptionId string,
    request models.UpdateSubscriptionMinimumPriceRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `request` | [`models.UpdateSubscriptionMinimumPriceRequest`](../../doc/models/update-subscription-minimum-price-request.md) | Body, Required | Request da requisição com o valor mínimo que será configurado |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"

request := models.UpdateSubscriptionMinimumPriceRequest{}

apiResponse, err := subscriptionsController.UpdateSubscriptionMiniumPrice(subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Cycle by Id

```go
GetSubscriptionCycleById(
    subscriptionId string,
    cycleId string) (
    https.ApiResponse[models.GetPeriodResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `cycleId` | `string` | Template, Required | - |

## Response Type

[`models.GetPeriodResponse`](../../doc/models/get-period-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
cycleId := "cycleId0"

apiResponse, err := subscriptionsController.GetSubscriptionCycleById(subscriptionId, cycleId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Usage Report

```go
GetUsageReport(
    subscriptionId string,
    periodId string) (
    https.ApiResponse[models.GetUsageReportResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription Id |
| `periodId` | `string` | Template, Required | The period Id |

## Response Type

[`models.GetUsageReportResponse`](../../doc/models/get-usage-report-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
periodId := "period_id0"

apiResponse, err := subscriptionsController.GetUsageReport(subscriptionId, periodId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Split Subscription

```go
UpdateSplitSubscription(
    id string,
    request models.UpdateSubscriptionSplitRequest) (
    https.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `string` | Template, Required | Subscription's id |
| `request` | [`models.UpdateSubscriptionSplitRequest`](../../doc/models/update-subscription-split-request.md) | Body, Required | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
id := "id0"

requestRules0 := models.CreateSplitRequest{
    Type:        "type6",
    Amount:      222,
    RecipientId: "recipient_id6",
}

requestRules := []models.CreateSplitRequest{requestRules0}
request := models.UpdateSubscriptionSplitRequest{
    Enabled: false,
    Rules:   requestRules,
}

apiResponse, err := subscriptionsController.UpdateSplitSubscription(id, &request)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

