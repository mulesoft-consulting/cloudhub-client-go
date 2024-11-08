# OpenIDProviderDynamicPostBodyOidcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | [**OpenIDProviderDynamicPostBodyOidcProviderClient**](OpenIDProviderDynamicPostBodyOidcProviderClient.md) |  | 
**Issuer** | Pointer to **string** |  | [optional] 
**Urls** | [**OpenIDProviderManualPostBodyOidcProviderUrls**](OpenIDProviderManualPostBodyOidcProviderUrls.md) |  | 
**GroupScope** | Pointer to **string** |  | [optional] 
**ClaimsMapping** | Pointer to [**OpenIDProviderManualPostBodyOidcProviderClaimsMapping**](OpenIDProviderManualPostBodyOidcProviderClaimsMapping.md) |  | [optional] 

## Methods

### NewOpenIDProviderDynamicPostBodyOidcProvider

`func NewOpenIDProviderDynamicPostBodyOidcProvider(client OpenIDProviderDynamicPostBodyOidcProviderClient, urls OpenIDProviderManualPostBodyOidcProviderUrls, ) *OpenIDProviderDynamicPostBodyOidcProvider`

NewOpenIDProviderDynamicPostBodyOidcProvider instantiates a new OpenIDProviderDynamicPostBodyOidcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderDynamicPostBodyOidcProviderWithDefaults

`func NewOpenIDProviderDynamicPostBodyOidcProviderWithDefaults() *OpenIDProviderDynamicPostBodyOidcProvider`

NewOpenIDProviderDynamicPostBodyOidcProviderWithDefaults instantiates a new OpenIDProviderDynamicPostBodyOidcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClient() OpenIDProviderDynamicPostBodyOidcProviderClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClientOk() (*OpenIDProviderDynamicPostBodyOidcProviderClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetClient(v OpenIDProviderDynamicPostBodyOidcProviderClient)`

SetClient sets Client field to given value.


### GetIssuer

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetUrls

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetUrls() OpenIDProviderManualPostBodyOidcProviderUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetUrlsOk() (*OpenIDProviderManualPostBodyOidcProviderUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetUrls(v OpenIDProviderManualPostBodyOidcProviderUrls)`

SetUrls sets Urls field to given value.


### GetGroupScope

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetGroupScope() string`

GetGroupScope returns the GroupScope field if non-nil, zero value otherwise.

### GetGroupScopeOk

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetGroupScopeOk() (*string, bool)`

GetGroupScopeOk returns a tuple with the GroupScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScope

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetGroupScope(v string)`

SetGroupScope sets GroupScope field to given value.

### HasGroupScope

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) HasGroupScope() bool`

HasGroupScope returns a boolean if a field has been set.

### GetClaimsMapping

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClaimsMapping() OpenIDProviderManualPostBodyOidcProviderClaimsMapping`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClaimsMappingOk() (*OpenIDProviderManualPostBodyOidcProviderClaimsMapping, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetClaimsMapping(v OpenIDProviderManualPostBodyOidcProviderClaimsMapping)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *OpenIDProviderDynamicPostBodyOidcProvider) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


