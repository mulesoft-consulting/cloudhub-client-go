# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectedApp**](DefaultApi.md#CreateConnectedApp) | **Post** /organizations/{orgId}/connectedApplications | 
[**DeleteConnectedApp**](DefaultApi.md#DeleteConnectedApp) | **Delete** /organizations/{orgId}/connectedApplications/{connAppId} | 
[**GetAllConnectedApps**](DefaultApi.md#GetAllConnectedApps) | **Get** /connectedApplications | 
[**GetConnectedApp**](DefaultApi.md#GetConnectedApp) | **Get** /organizations/{orgId}/connectedApplications/{connAppId} | 
[**GetConnectedAppByIdOnly**](DefaultApi.md#GetConnectedAppByIdOnly) | **Get** /connectedApplications/{connAppId} | 
[**GetConnectedAppScopes**](DefaultApi.md#GetConnectedAppScopes) | **Get** /organizations/{orgId}/connectedApplications/{connAppId}/scopes | 
[**UpdateConnectedApp**](DefaultApi.md#UpdateConnectedApp) | **Patch** /organizations/{orgId}/connectedApplications/{connAppId} | 
[**UpdateConnectedAppScopes**](DefaultApi.md#UpdateConnectedAppScopes) | **Put** /organizations/{orgId}/connectedApplications/{connAppId}/scopes | 



## CreateConnectedApp

> ConnectedAppRespExt CreateConnectedApp(ctx, orgId).ConnectedAppCore(connectedAppCore).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization
    connectedAppCore := *openapiclient.NewConnectedAppCore() // ConnectedAppCore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateConnectedApp(context.Background(), orgId).ConnectedAppCore(connectedAppCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConnectedApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectedApp`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConnectedApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectedAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectedAppCore** | [**ConnectedAppCore**](ConnectedAppCore.md) |  | 

### Return type

[**ConnectedAppRespExt**](ConnectedAppRespExt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectedApp

> DeleteConnectedApp(ctx, orgId, connAppId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization
    connAppId := "connAppId_example" // string | The ID of the connected app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteConnectedApp(context.Background(), orgId, connAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConnectedApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization | 
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectedAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConnectedApps

> GetAllConnectedApps200Response GetAllConnectedApps(ctx).IncludeUsage(includeUsage).OrgId(orgId).Offset(offset).Limit(limit).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    includeUsage := true // bool | flag to indicate whether to return usage statistics (optional)
    orgId := "orgId_example" // string | Provide an orgId to get all clients from other organization (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)
    limit := int32(56) // int32 | Maximum records to retrieve per request. (optional)
    search := "search_example" // string | A search string to use for case-insensitive partial matches on all object properties. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllConnectedApps(context.Background()).IncludeUsage(includeUsage).OrgId(orgId).Offset(offset).Limit(limit).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllConnectedApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConnectedApps`: GetAllConnectedApps200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllConnectedApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConnectedAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUsage** | **bool** | flag to indicate whether to return usage statistics | 
 **orgId** | **string** | Provide an orgId to get all clients from other organization | 
 **offset** | **int32** | The number of records to omit from the response. | 
 **limit** | **int32** | Maximum records to retrieve per request. | 
 **search** | **string** | A search string to use for case-insensitive partial matches on all object properties. | 

### Return type

[**GetAllConnectedApps200Response**](GetAllConnectedApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedApp

> ConnectedAppRespExt GetConnectedApp(ctx, orgId, connAppId).IncludeUsage(includeUsage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization
    connAppId := "connAppId_example" // string | The ID of the connected app
    includeUsage := true // bool | flag to indicate whether to return usage statistics (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetConnectedApp(context.Background(), orgId, connAppId).IncludeUsage(includeUsage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConnectedApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectedApp`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConnectedApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization | 
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeUsage** | **bool** | flag to indicate whether to return usage statistics | 

### Return type

[**ConnectedAppRespExt**](ConnectedAppRespExt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedAppByIdOnly

> ConnectedAppRespExt GetConnectedAppByIdOnly(ctx, connAppId).IncludeUsage(includeUsage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    connAppId := "connAppId_example" // string | The ID of the connected app
    includeUsage := true // bool | flag to indicate whether to return usage statistics (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetConnectedAppByIdOnly(context.Background(), connAppId).IncludeUsage(includeUsage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConnectedAppByIdOnly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectedAppByIdOnly`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConnectedAppByIdOnly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedAppByIdOnlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUsage** | **bool** | flag to indicate whether to return usage statistics | 

### Return type

[**ConnectedAppRespExt**](ConnectedAppRespExt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedAppScopes

> GetConnectedAppScopes200Response GetConnectedAppScopes(ctx, orgId, connAppId).Offset(offset).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization
    connAppId := "connAppId_example" // string | The ID of the connected app
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)
    limit := int32(56) // int32 | Maximum records to retrieve per request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetConnectedAppScopes(context.Background(), orgId, connAppId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConnectedAppScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectedAppScopes`: GetConnectedAppScopes200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConnectedAppScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization | 
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedAppScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | The number of records to omit from the response. | 
 **limit** | **int32** | Maximum records to retrieve per request. | 

### Return type

[**GetConnectedAppScopes200Response**](GetConnectedAppScopes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectedApp

> ConnectedAppRespExt UpdateConnectedApp(ctx, orgId, connAppId).ResetSecret(resetSecret).ConnectedAppPatchExt(connectedAppPatchExt).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization
    connAppId := "connAppId_example" // string | The ID of the connected app
    resetSecret := true // bool | Asks service to reset secret as part of this operation (optional)
    connectedAppPatchExt := *openapiclient.NewConnectedAppPatchExt() // ConnectedAppPatchExt |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateConnectedApp(context.Background(), orgId, connAppId).ResetSecret(resetSecret).ConnectedAppPatchExt(connectedAppPatchExt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConnectedApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectedApp`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConnectedApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization | 
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectedAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resetSecret** | **bool** | Asks service to reset secret as part of this operation | 
 **connectedAppPatchExt** | [**ConnectedAppPatchExt**](ConnectedAppPatchExt.md) |  | 

### Return type

[**ConnectedAppRespExt**](ConnectedAppRespExt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectedAppScopes

> UpdateConnectedAppScopes(ctx, orgId, connAppId).ConnectedAppScopesPutBody(connectedAppScopesPutBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization
    connAppId := "connAppId_example" // string | The ID of the connected app
    connectedAppScopesPutBody := *openapiclient.NewConnectedAppScopesPutBody() // ConnectedAppScopesPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateConnectedAppScopes(context.Background(), orgId, connAppId).ConnectedAppScopesPutBody(connectedAppScopesPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConnectedAppScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization | 
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectedAppScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectedAppScopesPutBody** | [**ConnectedAppScopesPutBody**](ConnectedAppScopesPutBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

