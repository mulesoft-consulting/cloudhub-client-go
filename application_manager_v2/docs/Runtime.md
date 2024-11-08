# Runtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | On deployment operations it can be set to:   * a full image version with tag (i.e \&quot;4.6.0:40e-java17\&quot;),   * a base version with a partial tag not indicating the java version (i.e. \&quot;4.6.0:40\&quot;)   * or only a base version (i.e. \&quot;4.6.0\&quot;). Defaults to the latest image version. This field has precedence over the legacy &#39;target.deploymentSettings.runtimeVersion&#39;  | [optional] 
**ReleaseChannel** | Pointer to **string** | On deployment operations it can be set to one of:   * \&quot;LTS\&quot;   * \&quot;EDGE\&quot;   * \&quot;LEGACY\&quot;. Defaults to \&quot;EDGE\&quot;. This field has precedence over the legacy &#39;target.deploymentSettings.runtimeReleaseChannel&#39;. Learn more on release channels at https://docs.mulesoft.com/release-notes/mule-runtime/lts-edge-release-cadence.  | [optional] 
**Java** | Pointer to **string** | On deployment operations it can be set to one of:   * \&quot;8\&quot;   * \&quot;17\&quot; Defaults to \&quot;8\&quot;. Learn more about Java support at https://docs.mulesoft.com/general/java-support.  | [optional] 

## Methods

### NewRuntime

`func NewRuntime() *Runtime`

NewRuntime instantiates a new Runtime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeWithDefaults

`func NewRuntimeWithDefaults() *Runtime`

NewRuntimeWithDefaults instantiates a new Runtime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Runtime) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Runtime) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Runtime) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Runtime) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetReleaseChannel

`func (o *Runtime) GetReleaseChannel() string`

GetReleaseChannel returns the ReleaseChannel field if non-nil, zero value otherwise.

### GetReleaseChannelOk

`func (o *Runtime) GetReleaseChannelOk() (*string, bool)`

GetReleaseChannelOk returns a tuple with the ReleaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannel

`func (o *Runtime) SetReleaseChannel(v string)`

SetReleaseChannel sets ReleaseChannel field to given value.

### HasReleaseChannel

`func (o *Runtime) HasReleaseChannel() bool`

HasReleaseChannel returns a boolean if a field has been set.

### GetJava

`func (o *Runtime) GetJava() string`

GetJava returns the Java field if non-nil, zero value otherwise.

### GetJavaOk

`func (o *Runtime) GetJavaOk() (*string, bool)`

GetJavaOk returns a tuple with the Java field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJava

`func (o *Runtime) SetJava(v string)`

SetJava sets Java field to given value.

### HasJava

`func (o *Runtime) HasJava() bool`

HasJava returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


