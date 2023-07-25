# Balance Operations

```go
balanceOperationsController := client.BalanceOperationsController
```

## Class Name

`BalanceOperationsController`

## Methods

* [Get Balance Operations](../../doc/controllers/balance-operations.md#get-balance-operations)
* [Get Balance Operation by Id](../../doc/controllers/balance-operations.md#get-balance-operation-by-id)


# Get Balance Operations

```go
GetBalanceOperations(
    status *string,
    createdSince *time.Time,
    createdUntil *time.Time) (
    https.ApiResponse[models.ListBalanceOperationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `status` | `*string` | Query, Optional | - |
| `createdSince` | `*time.Time` | Query, Optional | - |
| `createdUntil` | `*time.Time` | Query, Optional | - |

## Response Type

[`models.ListBalanceOperationResponse`](../../doc/models/list-balance-operation-response.md)

## Example Usage

```go
apiResponse, err := balanceOperationsController.GetBalanceOperations(nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Balance Operation by Id

```go
GetBalanceOperationById(
    id int64) (
    https.ApiResponse[models.GetBalanceOperationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int64` | Template, Required | - |

## Response Type

[`models.GetBalanceOperationResponse`](../../doc/models/get-balance-operation-response.md)

## Example Usage

```go
id := int64(112)

apiResponse, err := balanceOperationsController.GetBalanceOperationById(id)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

