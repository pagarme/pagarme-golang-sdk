# Orders

```go
ordersController := client.OrdersController
```

## Class Name

`OrdersController`

## Methods

* [Get Orders](../../doc/controllers/orders.md#get-orders)
* [Update Order Item](../../doc/controllers/orders.md#update-order-item)
* [Delete All Order Items](../../doc/controllers/orders.md#delete-all-order-items)
* [Delete Order Item](../../doc/controllers/orders.md#delete-order-item)
* [Close Order](../../doc/controllers/orders.md#close-order)
* [Create Order](../../doc/controllers/orders.md#create-order)
* [Create Order Item](../../doc/controllers/orders.md#create-order-item)
* [Get Order Item](../../doc/controllers/orders.md#get-order-item)
* [Update Order Metadata](../../doc/controllers/orders.md#update-order-metadata)
* [Get Order](../../doc/controllers/orders.md#get-order)


# Get Orders

Gets all orders

```go
GetOrders(
    page *int,
    size *int,
    code *string,
    status *string,
    createdSince *time.Time,
    createdUntil *time.Time,
    customerId *string) (
    https.ApiResponse[models.ListOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Filter for order's code |
| `status` | `*string` | Query, Optional | Filter for order's status |
| `createdSince` | `*time.Time` | Query, Optional | Filter for order's creation date start range |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for order's creation date end range |
| `customerId` | `*string` | Query, Optional | Filter for order's customer id |

## Response Type

[`models.ListOrderResponse`](../../doc/models/list-order-response.md)

## Example Usage

```go
apiResponse, err := ordersController.GetOrders(nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Order Item

```go
UpdateOrderItem(
    orderId string,
    itemId string,
    request models.UpdateOrderItemRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |
| `request` | [`models.UpdateOrderItemRequest`](../../doc/models/update-order-item-request.md) | Body, Required | Item Model |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
itemId := "itemId8"

request := models.UpdateOrderItemRequest{
    Amount:      242,
    Description: "description6",
    Quantity:    100,
    Category:    "category4",
}

apiResponse, err := ordersController.UpdateOrderItem(orderId, itemId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete All Order Items

```go
DeleteAllOrderItems(
    orderId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
orderId := "orderId2"

apiResponse, err := ordersController.DeleteAllOrderItems(orderId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Order Item

```go
DeleteOrderItem(
    orderId string,
    itemId string,
    idempotencyKey *string) (
    https.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
itemId := "itemId8"

apiResponse, err := ordersController.DeleteOrderItem(orderId, itemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Close Order

```go
CloseOrder(
    id string,
    request models.UpdateOrderStatusRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `string` | Template, Required | Order Id |
| `request` | [`models.UpdateOrderStatusRequest`](../../doc/models/update-order-status-request.md) | Body, Required | Update Order Model |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
id := "id0"

request := models.UpdateOrderStatusRequest{
    Status: "status8",
}

apiResponse, err := ordersController.CloseOrder(id, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Order

Creates a new Order

```go
CreateOrder(
    body models.CreateOrderRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`models.CreateOrderRequest`](../../doc/models/create-order-request.md) | Body, Required | Request for creating an order |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
bodyItems0 := models.CreateOrderItemRequest{
    Amount:      101,
    Description: "description3",
    Quantity:    215,
    Category:    "category1",
}

bodyItems := []models.CreateOrderItemRequest{bodyItems0}
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

bodyPayments0 := models.CreatePaymentRequest{
    PaymentMethod: "payment_method0",
}

bodyPayments := []models.CreatePaymentRequest{bodyPayments0}
body := models.CreateOrderRequest{
    Code:     "code4",
    Closed:   true,
    Items:    bodyItems,
    Customer: bodyCustomer,
    Payments: bodyPayments,
}

apiResponse, err := ordersController.CreateOrder(&body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Order Item

```go
CreateOrderItem(
    orderId string,
    request models.CreateOrderItemRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `request` | [`models.CreateOrderItemRequest`](../../doc/models/create-order-item-request.md) | Body, Required | Order Item Model |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
request := models.CreateOrderItemRequest{
    Amount:      242,
    Description: "description6",
    Quantity:    100,
    Category:    "category4",
}


apiResponse, err := ordersController.CreateOrderItem(orderId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Order Item

```go
GetOrderItem(
    orderId string,
    itemId string) (
    https.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
orderId := "orderId2"
itemId := "itemId8"

apiResponse, err := ordersController.GetOrderItem(orderId, itemId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Order Metadata

Updates the metadata from an order

```go
UpdateOrderMetadata(
    orderId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    https.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | The order id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the order metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
orderId := "order_id6"

request := models.UpdateMetadataRequest{
    Metadata: map[string]*string{
"key0" : "metadata3",
},
}

apiResponse, err := ordersController.UpdateOrderMetadata(orderId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Order

Gets an order

```go
GetOrder(
    orderId string) (
    https.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order id |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
orderId := "order_id6"

apiResponse, err := ordersController.GetOrder(orderId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

