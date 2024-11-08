# OpenIDProviderPatchOidcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**OpenIDProviderPatchOidcProviderClient**](OpenIDProviderPatchOidcProviderClient.md) |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Urls** | Pointer to [**OpenIDProviderPatchOidcProviderUrls**](OpenIDProviderPatchOidcProviderUrls.md) |  | [optional] 
**GroupScope** | Pointer to **string** |  | [optional] 
**ClaimsMapping** | Pointer to [**OpenIDProviderManualPostBodyOidcProviderClaimsMapping**](OpenIDProviderManualPostBodyOidcProviderClaimsMapping.md) |  | [optional] 

## Methods

### NewOpenIDProviderPatchOidcProvider

`func NewOpenIDProviderPatchOidcProvider() *OpenIDProviderPatchOidcProvider`

NewOpenIDProviderPatchOidcProvider instantiates a new OpenIDProviderPatchOidcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderPatchOidcProviderWithDefaults

`func NewOpenIDProviderPatchOidcProviderWithDefaults() *OpenIDProviderPatchOidcProvider`

NewOpenIDProviderPatchOidcProviderWithDefaults instantiates a new OpenIDProviderPatchOidcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *OpenIDProviderPatchOidcProvider) GetClient() OpenIDProviderPatchOidcProviderClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OpenIDProviderPatchOidcProvider) GetClientOk() (*OpenIDProviderPatchOidcProviderClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OpenIDProviderPatchOidcProvider) SetClient(v OpenIDProviderPatchOidcProviderClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *OpenIDProviderPatchOidcProvider) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetIssuer

`func (o *OpenIDProviderPatchOidcProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OpenIDProviderPatchOidcProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OpenIDProviderPatchOidcProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OpenIDProviderPatchOidcProvider) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetUrls

`func (o *OpenIDProviderPatchOidcProvider) GetUrls() OpenIDProviderPatchOidcProviderUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *OpenIDProviderPatchOidcProvider) GetUrlsOk() (*OpenIDProviderPatchOidcProviderUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *OpenIDProviderPatchOidcProvider) SetUrls(v OpenIDProviderPatchOidcProviderUrls)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *OpenIDProviderPatchOidcProvider) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetGroupScope

`func (o *OpenIDProviderPatchOidcProvider) GetGroupScope() string`

GetGroupScope returns the GroupScope field if non-nil, zero value otherwise.

### GetGroupScopeOk

`func (o *OpenIDProviderPatchOidcProvider) GetGroupScopeOk() (*string, bool)`

GetGroupScopeOk returns a tuple with the GroupScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScope

`func (o *OpenIDProviderPatchOidcProvider) SetGroupScope(v string)`

SetGroupScope sets GroupScope field to given value.

### HasGroupScope

`func (o *OpenIDProviderPatchOidcProvider) HasGroupScope() bool`

HasGroupScope returns a boolean if a field has been set.

### GetClaimsMapping

`func (o *OpenIDProviderPatchOidcProvider) GetClaimsMapping() OpenIDProviderManualPostBodyOidcProviderClaimsMapping`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *OpenIDProviderPatchOidcProvider) GetClaimsMappingOk() (*OpenIDProviderManualPostBodyOidcProviderClaimsMapping, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *OpenIDProviderPatchOidcProvider) SetClaimsMapping(v OpenIDProviderManualPostBodyOidcProviderClaimsMapping)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *OpenIDProviderPatchOidcProvider) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


