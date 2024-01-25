# Go API client for secretgroup_tlscontext

Secret Group TLS Context API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import secretgroup_tlscontext "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_tlscontext"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), secretgroup_tlscontext.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), secretgroup_tlscontext.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), secretgroup_tlscontext.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), secretgroup_tlscontext.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetSecretGroupTlsContextDetails**](docs/DefaultApi.md#getsecretgrouptlscontextdetails) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts/{secretId} | Retrieve tls-context details
*DefaultApi* | [**GetSecretGroupTlsContexts**](docs/DefaultApi.md#getsecretgrouptlscontexts) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts | Retrieves a secret-groups&#39; collection of tls-contexts.
*DefaultApi* | [**PatchSecretGroupTlsContext**](docs/DefaultApi.md#patchsecretgrouptlscontext) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts/{secretId} | Update a given secret-group tls-context
*DefaultApi* | [**PostSecretGroupTlsContext**](docs/DefaultApi.md#postsecretgrouptlscontext) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts | Create a secret-groups&#39; tls-context.
*DefaultApi* | [**PutSecretGroupTlsContext**](docs/DefaultApi.md#putsecretgrouptlscontext) | **Put** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts/{secretId} | Update a given secret-group tls-context


## Documentation For Models

 - [AcceptableCipherSuites](docs/AcceptableCipherSuites.md)
 - [AcceptableTlsVersions](docs/AcceptableTlsVersions.md)
 - [AuthenticationOverrides](docs/AuthenticationOverrides.md)
 - [CertificatePinning](docs/CertificatePinning.md)
 - [ErrorsResponse](docs/ErrorsResponse.md)
 - [ErrorsResponseErrorsInner](docs/ErrorsResponseErrorsInner.md)
 - [GetSecretGroupTlsContextDetails404Response](docs/GetSecretGroupTlsContextDetails404Response.md)
 - [Meta](docs/Meta.md)
 - [MutualAuthentication](docs/MutualAuthentication.md)
 - [PostSecretGroupTlsContext201Response](docs/PostSecretGroupTlsContext201Response.md)
 - [PutSecretGroupTlsContext200Response](docs/PutSecretGroupTlsContext200Response.md)
 - [SecretPath](docs/SecretPath.md)
 - [TlsContextDetails](docs/TlsContextDetails.md)
 - [TlsContextDetailsInboundSettings](docs/TlsContextDetailsInboundSettings.md)
 - [TlsContextDetailsOutboundSettings](docs/TlsContextDetailsOutboundSettings.md)
 - [TlsContextFlexGatewayBody](docs/TlsContextFlexGatewayBody.md)
 - [TlsContextFlexGatewayBodyInboundSettings](docs/TlsContextFlexGatewayBodyInboundSettings.md)
 - [TlsContextFlexGatewayBodyOutboundSettings](docs/TlsContextFlexGatewayBodyOutboundSettings.md)
 - [TlsContextMuleBody](docs/TlsContextMuleBody.md)
 - [TlsContextPostBody](docs/TlsContextPostBody.md)
 - [TlsContextPutBody](docs/TlsContextPutBody.md)
 - [TlsContextSfBody](docs/TlsContextSfBody.md)
 - [TlsContextSummary](docs/TlsContextSummary.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


