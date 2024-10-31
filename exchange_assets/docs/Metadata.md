# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagsInner**](TagsInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *Metadata) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *Metadata) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *Metadata) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *Metadata) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetTags

`func (o *Metadata) GetTags() []TagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Metadata) GetTagsOk() (*[]TagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Metadata) SetTags(v []TagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Metadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *Metadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Metadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Metadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Metadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Metadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Metadata) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


