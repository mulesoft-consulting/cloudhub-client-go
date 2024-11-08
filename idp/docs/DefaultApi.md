# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdIdentityProvidersGet**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersGet) | **Get** /organizations/{orgId}/identityProviders | 
[**OrganizationsOrgIdIdentityProvidersIdpIdDelete**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersIdpIdDelete) | **Delete** /organizations/{orgId}/identityProviders/{idpId} | 
[**OrganizationsOrgIdIdentityProvidersIdpIdGet**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersIdpIdGet) | **Get** /organizations/{orgId}/identityProviders/{idpId} | 
[**OrganizationsOrgIdIdentityProvidersIdpIdPatch**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersIdpIdPatch) | **Patch** /organizations/{orgId}/identityProviders/{idpId} | 
[**OrganizationsOrgIdIdentityProvidersPost**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersPost) | **Post** /organizations/{orgId}/identityProviders | 



## OrganizationsOrgIdIdentityProvidersGet

> OrganizationsOrgIdIdentityProvidersGet200Response OrganizationsOrgIdIdentityProvidersGet(ctx, orgId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/idp"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersGet(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersGet`: OrganizationsOrgIdIdentityProvidersGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationsOrgIdIdentityProvidersGet200Response**](OrganizationsOrgIdIdentityProvidersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdIdentityProvidersIdpIdDelete

> OrganizationsOrgIdIdentityProvidersIdpIdDelete(ctx, orgId, idpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/idp"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    idpId := "idpId_example" // string | The ID of the Identity Provider in GUID format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdDelete(context.Background(), orgId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**idpId** | **string** | The ID of the Identity Provider in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersIdpIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: plain/text

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdIdentityProvidersIdpIdGet

> Idp OrganizationsOrgIdIdentityProvidersIdpIdGet(ctx, orgId, idpId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/idp"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    idpId := "idpId_example" // string | The ID of the Identity Provider in GUID format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdGet(context.Background(), orgId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersIdpIdGet`: Idp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**idpId** | **string** | The ID of the Identity Provider in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersIdpIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Idp**](Idp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: plain/text, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdIdentityProvidersIdpIdPatch

> Idp OrganizationsOrgIdIdentityProvidersIdpIdPatch(ctx, orgId, idpId).IdpPatchBody(idpPatchBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/idp"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    idpId := "idpId_example" // string | The ID of the Identity Provider in GUID format
    idpPatchBody := openapiclient.idpPatchBody{LdapProviderPatch: openapiclient.NewLdapProviderPatch()} // IdpPatchBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdPatch(context.Background(), orgId, idpId).IdpPatchBody(idpPatchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersIdpIdPatch`: Idp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**idpId** | **string** | The ID of the Identity Provider in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersIdpIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idpPatchBody** | [**IdpPatchBody**](IdpPatchBody.md) |  | 

### Return type

[**Idp**](Idp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, plain/text

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdIdentityProvidersPost

> Idp OrganizationsOrgIdIdentityProvidersPost(ctx, orgId).IdpPostBody(idpPostBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/idp"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    idpPostBody := openapiclient.idpPostBody{LdapProviderPostBody: openapiclient.NewLdapProviderPostBody("Name_example", *openapiclient.NewLdapProviderPostBodyType("Name_example"), *openapiclient.NewLdapProviderPostBodyConnection("Binddn_example", float32(123), "Host_example", float32(123), "Password_example", float32(123)), *openapiclient.NewLdapProviderPostBodySearchBases("Group_example", "User_example"), *openapiclient.NewLdapProviderPostBodySearchBases("Group_example", "User_example"), *openapiclient.NewLdapProviderPostBodyFilters("GroupsByUsername_example", "UserByUsername_example"), *openapiclient.NewLdapProviderPostBodyUserMapping("Email_example", "FirstName_example", "Id_example", "LastName_example", "Username_example"), *openapiclient.NewLdapProviderPostBodyGroupMapping("GroupName_example", "Id_example"))} // IdpPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersPost(context.Background(), orgId).IdpPostBody(idpPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersPost`: Idp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpPostBody** | [**IdpPostBody**](IdpPostBody.md) |  | 

### Return type

[**Idp**](Idp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

