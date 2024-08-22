# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/runtimefabric/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFabrics**](DefaultApi.md#DeleteFabrics) | **Delete** /organizations/{orgId}/fabrics/{fabricsId} | Delete a specific Runtime Fabrics Instance by id.
[**GetAllFabrics**](DefaultApi.md#GetAllFabrics) | **Get** /organizations/{orgId}/fabrics | Retrieves all available Runtime Fabrics.
[**GetFabrics**](DefaultApi.md#GetFabrics) | **Get** /organizations/{orgId}/fabrics/{fabricsId} | Retrieves a specific Runtime Fabrics Instance by id.
[**GetFabricsAssociations**](DefaultApi.md#GetFabricsAssociations) | **Get** /organizations/{orgId}/fabrics/{fabricsId}/associations | Get Runtime Fabrics instance Environment Associations.
[**GetFabricsHealth**](DefaultApi.md#GetFabricsHealth) | **Get** /organizations/{orgId}/fabrics/{fabricsId}/health | Retrieves health for a Runtime Fabrics Instance by id.
[**GetFabricsHelmRepoProps**](DefaultApi.md#GetFabricsHelmRepoProps) | **Get** /organizations/{orgId}/helmrepoproperties | Retrieves Helm repo properties for a Runtime Fabrics Instance by id.
[**PostFabrics**](DefaultApi.md#PostFabrics) | **Post** /organizations/{orgId}/fabrics | Creates a Runtime Fabrics
[**PostFabricsAssociations**](DefaultApi.md#PostFabricsAssociations) | **Post** /organizations/{orgId}/fabrics/{fabricsId}/associations | Creates Runtime Fabrics instance Environment Associations.



## DeleteFabrics

> DeleteFabrics(ctx, orgId, fabricsId).Execute()

Delete a specific Runtime Fabrics Instance by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    fabricsId := "fabricsId_example" // string | The Runtime Fabrics Instance id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteFabrics(context.Background(), orgId, fabricsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFabrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**fabricsId** | **string** | The Runtime Fabrics Instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFabricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFabrics

> []Fabrics GetAllFabrics(ctx, orgId).Execute()

Retrieves all available Runtime Fabrics.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllFabrics(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllFabrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllFabrics`: []Fabrics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllFabrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFabricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Fabrics**](Fabrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFabrics

> Fabrics GetFabrics(ctx, orgId, fabricsId).Execute()

Retrieves a specific Runtime Fabrics Instance by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    fabricsId := "fabricsId_example" // string | The Runtime Fabrics Instance id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFabrics(context.Background(), orgId, fabricsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFabrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFabrics`: Fabrics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFabrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**fabricsId** | **string** | The Runtime Fabrics Instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFabricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Fabrics**](Fabrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFabricsAssociations

> []FabricsAssociationsInner GetFabricsAssociations(ctx, orgId, fabricsId).Execute()

Get Runtime Fabrics instance Environment Associations.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    fabricsId := "fabricsId_example" // string | The Runtime Fabrics Instance id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFabricsAssociations(context.Background(), orgId, fabricsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFabricsAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFabricsAssociations`: []FabricsAssociationsInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFabricsAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**fabricsId** | **string** | The Runtime Fabrics Instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFabricsAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]FabricsAssociationsInner**](FabricsAssociationsInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFabricsHealth

> FabricsHealth GetFabricsHealth(ctx, orgId, fabricsId).Execute()

Retrieves health for a Runtime Fabrics Instance by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    fabricsId := "fabricsId_example" // string | The Runtime Fabrics Instance id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFabricsHealth(context.Background(), orgId, fabricsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFabricsHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFabricsHealth`: FabricsHealth
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFabricsHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**fabricsId** | **string** | The Runtime Fabrics Instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFabricsHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FabricsHealth**](FabricsHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFabricsHelmRepoProps

> FabricsHelmRepoProps GetFabricsHelmRepoProps(ctx, orgId).Execute()

Retrieves Helm repo properties for a Runtime Fabrics Instance by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFabricsHelmRepoProps(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFabricsHelmRepoProps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFabricsHelmRepoProps`: FabricsHelmRepoProps
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFabricsHelmRepoProps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFabricsHelmRepoPropsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FabricsHelmRepoProps**](FabricsHelmRepoProps.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFabrics

> Fabrics PostFabrics(ctx, orgId).FabricsPostBody(fabricsPostBody).Execute()

Creates a Runtime Fabrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    fabricsPostBody := *openapiclient.NewFabricsPostBody() // FabricsPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostFabrics(context.Background(), orgId).FabricsPostBody(fabricsPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostFabrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFabrics`: Fabrics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostFabrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFabricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fabricsPostBody** | [**FabricsPostBody**](FabricsPostBody.md) |  | 

### Return type

[**Fabrics**](Fabrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFabricsAssociations

> []FabricsAssociationsInner PostFabricsAssociations(ctx, orgId, fabricsId).FabricsAssociationsPostBody(fabricsAssociationsPostBody).Execute()

Creates Runtime Fabrics instance Environment Associations.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rtf"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    fabricsId := "fabricsId_example" // string | The Runtime Fabrics Instance id
    fabricsAssociationsPostBody := *openapiclient.NewFabricsAssociationsPostBody() // FabricsAssociationsPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostFabricsAssociations(context.Background(), orgId, fabricsId).FabricsAssociationsPostBody(fabricsAssociationsPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostFabricsAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFabricsAssociations`: []FabricsAssociationsInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostFabricsAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**fabricsId** | **string** | The Runtime Fabrics Instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFabricsAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fabricsAssociationsPostBody** | [**FabricsAssociationsPostBody**](FabricsAssociationsPostBody.md) |  | 

### Return type

[**[]FabricsAssociationsInner**](FabricsAssociationsInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

