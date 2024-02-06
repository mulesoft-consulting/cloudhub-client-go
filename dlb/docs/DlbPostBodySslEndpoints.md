# DlbPostBodySslEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyLabel** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PublicKeyLabel** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**VerifyClientMode** | Pointer to **string** |  | [optional] [default to "off"]
**ClientCertLabel** | Pointer to **string** |  | [optional] 
**ClientCert** | Pointer to **string** |  | [optional] 
**RevocationListLabel** | Pointer to **string** |  | [optional] 
**RevocationList** | Pointer to **string** |  | [optional] 
**Mappings** | Pointer to [**[]DlbPostBodyMappings**](DlbPostBodyMappings.md) |  | [optional] 

## Methods

### NewDlbPostBodySslEndpoints

`func NewDlbPostBodySslEndpoints() *DlbPostBodySslEndpoints`

NewDlbPostBodySslEndpoints instantiates a new DlbPostBodySslEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbPostBodySslEndpointsWithDefaults

`func NewDlbPostBodySslEndpointsWithDefaults() *DlbPostBodySslEndpoints`

NewDlbPostBodySslEndpointsWithDefaults instantiates a new DlbPostBodySslEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyLabel

`func (o *DlbPostBodySslEndpoints) GetPrivateKeyLabel() string`

GetPrivateKeyLabel returns the PrivateKeyLabel field if non-nil, zero value otherwise.

### GetPrivateKeyLabelOk

`func (o *DlbPostBodySslEndpoints) GetPrivateKeyLabelOk() (*string, bool)`

GetPrivateKeyLabelOk returns a tuple with the PrivateKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyLabel

`func (o *DlbPostBodySslEndpoints) SetPrivateKeyLabel(v string)`

SetPrivateKeyLabel sets PrivateKeyLabel field to given value.

### HasPrivateKeyLabel

`func (o *DlbPostBodySslEndpoints) HasPrivateKeyLabel() bool`

HasPrivateKeyLabel returns a boolean if a field has been set.

### GetPrivateKey

`func (o *DlbPostBodySslEndpoints) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *DlbPostBodySslEndpoints) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *DlbPostBodySslEndpoints) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *DlbPostBodySslEndpoints) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKeyLabel

`func (o *DlbPostBodySslEndpoints) GetPublicKeyLabel() string`

GetPublicKeyLabel returns the PublicKeyLabel field if non-nil, zero value otherwise.

### GetPublicKeyLabelOk

`func (o *DlbPostBodySslEndpoints) GetPublicKeyLabelOk() (*string, bool)`

GetPublicKeyLabelOk returns a tuple with the PublicKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyLabel

`func (o *DlbPostBodySslEndpoints) SetPublicKeyLabel(v string)`

SetPublicKeyLabel sets PublicKeyLabel field to given value.

### HasPublicKeyLabel

`func (o *DlbPostBodySslEndpoints) HasPublicKeyLabel() bool`

HasPublicKeyLabel returns a boolean if a field has been set.

### GetPublicKey

`func (o *DlbPostBodySslEndpoints) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DlbPostBodySslEndpoints) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DlbPostBodySslEndpoints) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DlbPostBodySslEndpoints) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetVerifyClientMode

`func (o *DlbPostBodySslEndpoints) GetVerifyClientMode() string`

GetVerifyClientMode returns the VerifyClientMode field if non-nil, zero value otherwise.

### GetVerifyClientModeOk

`func (o *DlbPostBodySslEndpoints) GetVerifyClientModeOk() (*string, bool)`

GetVerifyClientModeOk returns a tuple with the VerifyClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyClientMode

`func (o *DlbPostBodySslEndpoints) SetVerifyClientMode(v string)`

SetVerifyClientMode sets VerifyClientMode field to given value.

### HasVerifyClientMode

`func (o *DlbPostBodySslEndpoints) HasVerifyClientMode() bool`

HasVerifyClientMode returns a boolean if a field has been set.

### GetClientCertLabel

`func (o *DlbPostBodySslEndpoints) GetClientCertLabel() string`

GetClientCertLabel returns the ClientCertLabel field if non-nil, zero value otherwise.

### GetClientCertLabelOk

`func (o *DlbPostBodySslEndpoints) GetClientCertLabelOk() (*string, bool)`

GetClientCertLabelOk returns a tuple with the ClientCertLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertLabel

`func (o *DlbPostBodySslEndpoints) SetClientCertLabel(v string)`

SetClientCertLabel sets ClientCertLabel field to given value.

### HasClientCertLabel

`func (o *DlbPostBodySslEndpoints) HasClientCertLabel() bool`

HasClientCertLabel returns a boolean if a field has been set.

### GetClientCert

`func (o *DlbPostBodySslEndpoints) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *DlbPostBodySslEndpoints) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *DlbPostBodySslEndpoints) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *DlbPostBodySslEndpoints) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetRevocationListLabel

`func (o *DlbPostBodySslEndpoints) GetRevocationListLabel() string`

GetRevocationListLabel returns the RevocationListLabel field if non-nil, zero value otherwise.

### GetRevocationListLabelOk

`func (o *DlbPostBodySslEndpoints) GetRevocationListLabelOk() (*string, bool)`

GetRevocationListLabelOk returns a tuple with the RevocationListLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationListLabel

`func (o *DlbPostBodySslEndpoints) SetRevocationListLabel(v string)`

SetRevocationListLabel sets RevocationListLabel field to given value.

### HasRevocationListLabel

`func (o *DlbPostBodySslEndpoints) HasRevocationListLabel() bool`

HasRevocationListLabel returns a boolean if a field has been set.

### GetRevocationList

`func (o *DlbPostBodySslEndpoints) GetRevocationList() string`

GetRevocationList returns the RevocationList field if non-nil, zero value otherwise.

### GetRevocationListOk

`func (o *DlbPostBodySslEndpoints) GetRevocationListOk() (*string, bool)`

GetRevocationListOk returns a tuple with the RevocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationList

`func (o *DlbPostBodySslEndpoints) SetRevocationList(v string)`

SetRevocationList sets RevocationList field to given value.

### HasRevocationList

`func (o *DlbPostBodySslEndpoints) HasRevocationList() bool`

HasRevocationList returns a boolean if a field has been set.

### GetMappings

`func (o *DlbPostBodySslEndpoints) GetMappings() []DlbPostBodyMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *DlbPostBodySslEndpoints) GetMappingsOk() (*[]DlbPostBodyMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *DlbPostBodySslEndpoints) SetMappings(v []DlbPostBodyMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *DlbPostBodySslEndpoints) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


