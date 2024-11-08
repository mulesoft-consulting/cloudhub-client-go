# OpenIDProviderGetOidcProviderClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**OpenIDProviderGetOidcProviderClientCredentials**](OpenIDProviderGetOidcProviderClientCredentials.md) |  | 
**Urls** | Pointer to [**OpenIDProviderDynamicPostBodyOidcProviderClientUrls**](OpenIDProviderDynamicPostBodyOidcProviderClientUrls.md) |  | [optional] 
**Registration** | Pointer to [**OpenIDProviderPatchOidcProviderClientRegistration**](OpenIDProviderPatchOidcProviderClientRegistration.md) |  | [optional] 
**Metadata** | Pointer to [**OpenIDProviderDynamicPostBodyOidcProviderClientMetadata**](OpenIDProviderDynamicPostBodyOidcProviderClientMetadata.md) |  | [optional] 
**TokenEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOpenIDProviderGetOidcProviderClient

`func NewOpenIDProviderGetOidcProviderClient(credentials OpenIDProviderGetOidcProviderClientCredentials, ) *OpenIDProviderGetOidcProviderClient`

NewOpenIDProviderGetOidcProviderClient instantiates a new OpenIDProviderGetOidcProviderClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIDProviderGetOidcProviderClientWithDefaults

`func NewOpenIDProviderGetOidcProviderClientWithDefaults() *OpenIDProviderGetOidcProviderClient`

NewOpenIDProviderGetOidcProviderClientWithDefaults instantiates a new OpenIDProviderGetOidcProviderClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *OpenIDProviderGetOidcProviderClient) GetCredentials() OpenIDProviderGetOidcProviderClientCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OpenIDProviderGetOidcProviderClient) GetCredentialsOk() (*OpenIDProviderGetOidcProviderClientCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OpenIDProviderGetOidcProviderClient) SetCredentials(v OpenIDProviderGetOidcProviderClientCredentials)`

SetCredentials sets Credentials field to given value.


### GetUrls

`func (o *OpenIDProviderGetOidcProviderClient) GetUrls() OpenIDProviderDynamicPostBodyOidcProviderClientUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *OpenIDProviderGetOidcProviderClient) GetUrlsOk() (*OpenIDProviderDynamicPostBodyOidcProviderClientUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *OpenIDProviderGetOidcProviderClient) SetUrls(v OpenIDProviderDynamicPostBodyOidcProviderClientUrls)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *OpenIDProviderGetOidcProviderClient) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetRegistration

`func (o *OpenIDProviderGetOidcProviderClient) GetRegistration() OpenIDProviderPatchOidcProviderClientRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *OpenIDProviderGetOidcProviderClient) GetRegistrationOk() (*OpenIDProviderPatchOidcProviderClientRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *OpenIDProviderGetOidcProviderClient) SetRegistration(v OpenIDProviderPatchOidcProviderClientRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *OpenIDProviderGetOidcProviderClient) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetMetadata

`func (o *OpenIDProviderGetOidcProviderClient) GetMetadata() OpenIDProviderDynamicPostBodyOidcProviderClientMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OpenIDProviderGetOidcProviderClient) GetMetadataOk() (*OpenIDProviderDynamicPostBodyOidcProviderClientMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OpenIDProviderGetOidcProviderClient) SetMetadata(v OpenIDProviderDynamicPostBodyOidcProviderClientMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OpenIDProviderGetOidcProviderClient) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTokenEndpointAuthMethodsSupported

`func (o *OpenIDProviderGetOidcProviderClient) GetTokenEndpointAuthMethodsSupported() []string`

GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodsSupportedOk

`func (o *OpenIDProviderGetOidcProviderClient) GetTokenEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethodsSupported

`func (o *OpenIDProviderGetOidcProviderClient) SetTokenEndpointAuthMethodsSupported(v []string)`

SetTokenEndpointAuthMethodsSupported sets TokenEndpointAuthMethodsSupported field to given value.

### HasTokenEndpointAuthMethodsSupported

`func (o *OpenIDProviderGetOidcProviderClient) HasTokenEndpointAuthMethodsSupported() bool`

HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


