# OpenIDProviderManualPostBodyOidcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | [**OpenIDProviderManualPostBodyOidcProviderClient**](OpenIDProviderManualPostBodyOidcProviderClient.md) |  | 
**Issuer** | Pointer to **string** |  | [optional] 
**Urls** | [**OpenIDProviderManualPostBodyOidcProviderUrls**](OpenIDProviderManualPostBodyOidcProviderUrls.md) |  | 
**GroupScope** | Pointer to **string** |  | [optional] 
**ClaimsMapping** | Pointer to [**OpenIDProviderManualPostBodyOidcProviderClaimsMapping**](OpenIDProviderManualPostBodyOidcProviderClaimsMapping.md) |  | [optional] 

## Methods

### NewOpenIDProviderManualPostBodyOidcProvider

`func NewOpenIDProviderManualPostBodyOidcProvider(client OpenIDProviderManualPostBodyOidcProviderClient, urls OpenIDProviderManualPostBodyOidcProviderUrls, ) *OpenIDProviderManualPostBodyOidcProvider`

NewOpenIDProviderManualPostBodyOidcProvider instantiates a new OpenIDProviderManualPostBodyOidcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderManualPostBodyOidcProviderWithDefaults

`func NewOpenIDProviderManualPostBodyOidcProviderWithDefaults() *OpenIDProviderManualPostBodyOidcProvider`

NewOpenIDProviderManualPostBodyOidcProviderWithDefaults instantiates a new OpenIDProviderManualPostBodyOidcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetClient() OpenIDProviderManualPostBodyOidcProviderClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetClientOk() (*OpenIDProviderManualPostBodyOidcProviderClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OpenIDProviderManualPostBodyOidcProvider) SetClient(v OpenIDProviderManualPostBodyOidcProviderClient)`

SetClient sets Client field to given value.


### GetIssuer

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OpenIDProviderManualPostBodyOidcProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OpenIDProviderManualPostBodyOidcProvider) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetUrls

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetUrls() OpenIDProviderManualPostBodyOidcProviderUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetUrlsOk() (*OpenIDProviderManualPostBodyOidcProviderUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *OpenIDProviderManualPostBodyOidcProvider) SetUrls(v OpenIDProviderManualPostBodyOidcProviderUrls)`

SetUrls sets Urls field to given value.


### GetGroupScope

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetGroupScope() string`

GetGroupScope returns the GroupScope field if non-nil, zero value otherwise.

### GetGroupScopeOk

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetGroupScopeOk() (*string, bool)`

GetGroupScopeOk returns a tuple with the GroupScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScope

`func (o *OpenIDProviderManualPostBodyOidcProvider) SetGroupScope(v string)`

SetGroupScope sets GroupScope field to given value.

### HasGroupScope

`func (o *OpenIDProviderManualPostBodyOidcProvider) HasGroupScope() bool`

HasGroupScope returns a boolean if a field has been set.

### GetClaimsMapping

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetClaimsMapping() OpenIDProviderManualPostBodyOidcProviderClaimsMapping`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *OpenIDProviderManualPostBodyOidcProvider) GetClaimsMappingOk() (*OpenIDProviderManualPostBodyOidcProviderClaimsMapping, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *OpenIDProviderManualPostBodyOidcProvider) SetClaimsMapping(v OpenIDProviderManualPostBodyOidcProviderClaimsMapping)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *OpenIDProviderManualPostBodyOidcProvider) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


