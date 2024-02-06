/*
Exchange Assets

Description of the Exchange Assets API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package exchange_assets

import (
	"encoding/json"
	"os"
)

// checks if the AssetSearchResultItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetSearchResultItem{}

// AssetSearchResultItem struct for AssetSearchResultItem
type AssetSearchResultItem struct {
	GroupId *string `json:"groupId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	Version *string `json:"version,omitempty"`
	MinorVersion *string `json:"minorVersion,omitempty"`
	VersionGroup *string `json:"versionGroup,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	Name *string `json:"name,omitempty"`
	ContactName NullableString `json:"contactName,omitempty"`
	ContactEmail NullableString `json:"contactEmail,omitempty"`
	Type *string `json:"type,omitempty"`
	IsSnapshot *bool `json:"isSnapshot,omitempty"`
	Status *string `json:"status,omitempty"`
	ExternalFile *ExternalFile `json:"externalFile,omitempty"`
	CreatedDate *string `json:"createdDate,omitempty"`
	UpdatedDate *string `json:"updatedDate,omitempty"`
	MinMuleVersion NullableString `json:"minMuleVersion,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Categories []Category `json:"categories,omitempty"`
	Files []*os.File `json:"files,omitempty"`
	CustomFields []CustomField `json:"customFields,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
	NumberOfRates *int32 `json:"numberOfRates,omitempty"`
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Id *string `json:"id,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	ModifiedAt *string `json:"modifiedAt,omitempty"`
}

// NewAssetSearchResultItem instantiates a new AssetSearchResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSearchResultItem() *AssetSearchResultItem {
	this := AssetSearchResultItem{}
	return &this
}

// NewAssetSearchResultItemWithDefaults instantiates a new AssetSearchResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSearchResultItemWithDefaults() *AssetSearchResultItem {
	this := AssetSearchResultItem{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AssetSearchResultItem) SetGroupId(v string) {
	o.GroupId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *AssetSearchResultItem) SetAssetId(v string) {
	o.AssetId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AssetSearchResultItem) SetVersion(v string) {
	o.Version = &v
}

// GetMinorVersion returns the MinorVersion field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetMinorVersion() string {
	if o == nil || IsNil(o.MinorVersion) {
		var ret string
		return ret
	}
	return *o.MinorVersion
}

// GetMinorVersionOk returns a tuple with the MinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetMinorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinorVersion) {
		return nil, false
	}
	return o.MinorVersion, true
}

// HasMinorVersion returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasMinorVersion() bool {
	if o != nil && !IsNil(o.MinorVersion) {
		return true
	}

	return false
}

// SetMinorVersion gets a reference to the given string and assigns it to the MinorVersion field.
func (o *AssetSearchResultItem) SetMinorVersion(v string) {
	o.MinorVersion = &v
}

// GetVersionGroup returns the VersionGroup field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetVersionGroup() string {
	if o == nil || IsNil(o.VersionGroup) {
		var ret string
		return ret
	}
	return *o.VersionGroup
}

// GetVersionGroupOk returns a tuple with the VersionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetVersionGroupOk() (*string, bool) {
	if o == nil || IsNil(o.VersionGroup) {
		return nil, false
	}
	return o.VersionGroup, true
}

// HasVersionGroup returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasVersionGroup() bool {
	if o != nil && !IsNil(o.VersionGroup) {
		return true
	}

	return false
}

// SetVersionGroup gets a reference to the given string and assigns it to the VersionGroup field.
func (o *AssetSearchResultItem) SetVersionGroup(v string) {
	o.VersionGroup = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssetSearchResultItem) SetDescription(v string) {
	o.Description = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *AssetSearchResultItem) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetSearchResultItem) SetName(v string) {
	o.Name = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSearchResultItem) GetContactName() string {
	if o == nil || IsNil(o.ContactName.Get()) {
		var ret string
		return ret
	}
	return *o.ContactName.Get()
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSearchResultItem) GetContactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactName.Get(), o.ContactName.IsSet()
}

// HasContactName returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasContactName() bool {
	if o != nil && o.ContactName.IsSet() {
		return true
	}

	return false
}

// SetContactName gets a reference to the given NullableString and assigns it to the ContactName field.
func (o *AssetSearchResultItem) SetContactName(v string) {
	o.ContactName.Set(&v)
}
// SetContactNameNil sets the value for ContactName to be an explicit nil
func (o *AssetSearchResultItem) SetContactNameNil() {
	o.ContactName.Set(nil)
}

// UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
func (o *AssetSearchResultItem) UnsetContactName() {
	o.ContactName.Unset()
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSearchResultItem) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSearchResultItem) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableString and assigns it to the ContactEmail field.
func (o *AssetSearchResultItem) SetContactEmail(v string) {
	o.ContactEmail.Set(&v)
}
// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *AssetSearchResultItem) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *AssetSearchResultItem) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssetSearchResultItem) SetType(v string) {
	o.Type = &v
}

// GetIsSnapshot returns the IsSnapshot field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetIsSnapshot() bool {
	if o == nil || IsNil(o.IsSnapshot) {
		var ret bool
		return ret
	}
	return *o.IsSnapshot
}

// GetIsSnapshotOk returns a tuple with the IsSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetIsSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSnapshot) {
		return nil, false
	}
	return o.IsSnapshot, true
}

// HasIsSnapshot returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasIsSnapshot() bool {
	if o != nil && !IsNil(o.IsSnapshot) {
		return true
	}

	return false
}

// SetIsSnapshot gets a reference to the given bool and assigns it to the IsSnapshot field.
func (o *AssetSearchResultItem) SetIsSnapshot(v bool) {
	o.IsSnapshot = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetSearchResultItem) SetStatus(v string) {
	o.Status = &v
}

// GetExternalFile returns the ExternalFile field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetExternalFile() ExternalFile {
	if o == nil || IsNil(o.ExternalFile) {
		var ret ExternalFile
		return ret
	}
	return *o.ExternalFile
}

// GetExternalFileOk returns a tuple with the ExternalFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetExternalFileOk() (*ExternalFile, bool) {
	if o == nil || IsNil(o.ExternalFile) {
		return nil, false
	}
	return o.ExternalFile, true
}

// HasExternalFile returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasExternalFile() bool {
	if o != nil && !IsNil(o.ExternalFile) {
		return true
	}

	return false
}

// SetExternalFile gets a reference to the given ExternalFile and assigns it to the ExternalFile field.
func (o *AssetSearchResultItem) SetExternalFile(v ExternalFile) {
	o.ExternalFile = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *AssetSearchResultItem) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetUpdatedDate() string {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetUpdatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *AssetSearchResultItem) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetMinMuleVersion returns the MinMuleVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSearchResultItem) GetMinMuleVersion() string {
	if o == nil || IsNil(o.MinMuleVersion.Get()) {
		var ret string
		return ret
	}
	return *o.MinMuleVersion.Get()
}

// GetMinMuleVersionOk returns a tuple with the MinMuleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSearchResultItem) GetMinMuleVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinMuleVersion.Get(), o.MinMuleVersion.IsSet()
}

// HasMinMuleVersion returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasMinMuleVersion() bool {
	if o != nil && o.MinMuleVersion.IsSet() {
		return true
	}

	return false
}

// SetMinMuleVersion gets a reference to the given NullableString and assigns it to the MinMuleVersion field.
func (o *AssetSearchResultItem) SetMinMuleVersion(v string) {
	o.MinMuleVersion.Set(&v)
}
// SetMinMuleVersionNil sets the value for MinMuleVersion to be an explicit nil
func (o *AssetSearchResultItem) SetMinMuleVersionNil() {
	o.MinMuleVersion.Set(nil)
}

// UnsetMinMuleVersion ensures that no value is present for MinMuleVersion, not even an explicit nil
func (o *AssetSearchResultItem) UnsetMinMuleVersion() {
	o.MinMuleVersion.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *AssetSearchResultItem) SetLabels(v []string) {
	o.Labels = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetCategories() []Category {
	if o == nil || IsNil(o.Categories) {
		var ret []Category
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetCategoriesOk() ([]Category, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []Category and assigns it to the Categories field.
func (o *AssetSearchResultItem) SetCategories(v []Category) {
	o.Categories = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetFiles() []*os.File {
	if o == nil || IsNil(o.Files) {
		var ret []*os.File
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetFilesOk() ([]*os.File, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []*os.File and assigns it to the Files field.
func (o *AssetSearchResultItem) SetFiles(v []*os.File) {
	o.Files = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetCustomFields() []CustomField {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomField
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetCustomFieldsOk() ([]CustomField, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomField and assigns it to the CustomFields field.
func (o *AssetSearchResultItem) SetCustomFields(v []CustomField) {
	o.CustomFields = v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *AssetSearchResultItem) SetRating(v int32) {
	o.Rating = &v
}

// GetNumberOfRates returns the NumberOfRates field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetNumberOfRates() int32 {
	if o == nil || IsNil(o.NumberOfRates) {
		var ret int32
		return ret
	}
	return *o.NumberOfRates
}

// GetNumberOfRatesOk returns a tuple with the NumberOfRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetNumberOfRatesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfRates) {
		return nil, false
	}
	return o.NumberOfRates, true
}

// HasNumberOfRates returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasNumberOfRates() bool {
	if o != nil && !IsNil(o.NumberOfRates) {
		return true
	}

	return false
}

// SetNumberOfRates gets a reference to the given int32 and assigns it to the NumberOfRates field.
func (o *AssetSearchResultItem) SetNumberOfRates(v int32) {
	o.NumberOfRates = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetCreatedBy() CreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret CreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetCreatedByOk() (*CreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given CreatedBy and assigns it to the CreatedBy field.
func (o *AssetSearchResultItem) SetCreatedBy(v CreatedBy) {
	o.CreatedBy = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *AssetSearchResultItem) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetSearchResultItem) SetId(v string) {
	o.Id = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSearchResultItem) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSearchResultItem) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *AssetSearchResultItem) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *AssetSearchResultItem) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *AssetSearchResultItem) UnsetIcon() {
	o.Icon.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AssetSearchResultItem) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AssetSearchResultItem) GetModifiedAt() string {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSearchResultItem) GetModifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AssetSearchResultItem) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *AssetSearchResultItem) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

func (o AssetSearchResultItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetSearchResultItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.MinorVersion) {
		toSerialize["minorVersion"] = o.MinorVersion
	}
	if !IsNil(o.VersionGroup) {
		toSerialize["versionGroup"] = o.VersionGroup
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ContactName.IsSet() {
		toSerialize["contactName"] = o.ContactName.Get()
	}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsSnapshot) {
		toSerialize["isSnapshot"] = o.IsSnapshot
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ExternalFile) {
		toSerialize["externalFile"] = o.ExternalFile
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.UpdatedDate) {
		toSerialize["updatedDate"] = o.UpdatedDate
	}
	if o.MinMuleVersion.IsSet() {
		toSerialize["minMuleVersion"] = o.MinMuleVersion.Get()
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.NumberOfRates) {
		toSerialize["numberOfRates"] = o.NumberOfRates
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return toSerialize, nil
}

type NullableAssetSearchResultItem struct {
	value *AssetSearchResultItem
	isSet bool
}

func (v NullableAssetSearchResultItem) Get() *AssetSearchResultItem {
	return v.value
}

func (v *NullableAssetSearchResultItem) Set(val *AssetSearchResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSearchResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSearchResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSearchResultItem(val *AssetSearchResultItem) *NullableAssetSearchResultItem {
	return &NullableAssetSearchResultItem{value: val, isSet: true}
}

func (v NullableAssetSearchResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSearchResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


