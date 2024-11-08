# OpenIDProviderGetOidcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | [**OpenIDProviderGetOidcProviderClient**](OpenIDProviderGetOidcProviderClient.md) |  | 
**Issuer** | **string** |  | 
**Urls** | [**OpenIDProviderGetOidcProviderUrls**](OpenIDProviderGetOidcProviderUrls.md) |  | 
**GroupScope** | Pointer to **string** |  | [optional] 
**ClaimsMapping** | Pointer to [**OpenIDProviderManualPostBodyOidcProviderClaimsMapping**](OpenIDProviderManualPostBodyOidcProviderClaimsMapping.md) |  | [optional] 

## Methods

### NewOpenIDProviderGetOidcProvider

`func NewOpenIDProviderGetOidcProvider(client OpenIDProviderGetOidcProviderClient, issuer string, urls OpenIDProviderGetOidcProviderUrls, ) *OpenIDProviderGetOidcProvider`

NewOpenIDProviderGetOidcProvider instantiates a new OpenIDProviderGetOidcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderGetOidcProviderWithDefaults

`func NewOpenIDProviderGetOidcProviderWithDefaults() *OpenIDProviderGetOidcProvider`

NewOpenIDProviderGetOidcProviderWithDefaults instantiates a new OpenIDProviderGetOidcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *OpenIDProviderGetOidcProvider) GetClient() OpenIDProviderGetOidcProviderClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OpenIDProviderGetOidcProvider) GetClientOk() (*OpenIDProviderGetOidcProviderClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OpenIDProviderGetOidcProvider) SetClient(v OpenIDProviderGetOidcProviderClient)`

SetClient sets Client field to given value.


### GetIssuer

`func (o *OpenIDProviderGetOidcProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OpenIDProviderGetOidcProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OpenIDProviderGetOidcProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetUrls

`func (o *OpenIDProviderGetOidcProvider) GetUrls() OpenIDProviderGetOidcProviderUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *OpenIDProviderGetOidcProvider) GetUrlsOk() (*OpenIDProviderGetOidcProviderUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *OpenIDProviderGetOidcProvider) SetUrls(v OpenIDProviderGetOidcProviderUrls)`

SetUrls sets Urls field to given value.


### GetGroupScope

`func (o *OpenIDProviderGetOidcProvider) GetGroupScope() string`

GetGroupScope returns the GroupScope field if non-nil, zero value otherwise.

### GetGroupScopeOk

`func (o *OpenIDProviderGetOidcProvider) GetGroupScopeOk() (*string, bool)`

GetGroupScopeOk returns a tuple with the GroupScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScope

`func (o *OpenIDProviderGetOidcProvider) SetGroupScope(v string)`

SetGroupScope sets GroupScope field to given value.

### HasGroupScope

`func (o *OpenIDProviderGetOidcProvider) HasGroupScope() bool`

HasGroupScope returns a boolean if a field has been set.

### GetClaimsMapping

`func (o *OpenIDProviderGetOidcProvider) GetClaimsMapping() OpenIDProviderManualPostBodyOidcProviderClaimsMapping`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *OpenIDProviderGetOidcProvider) GetClaimsMappingOk() (*OpenIDProviderManualPostBodyOidcProviderClaimsMapping, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *OpenIDProviderGetOidcProvider) SetClaimsMapping(v OpenIDProviderManualPostBodyOidcProviderClaimsMapping)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *OpenIDProviderGetOidcProvider) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


