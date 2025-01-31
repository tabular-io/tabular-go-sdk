/*
Tabular API

Tabular OpenAPI Definition

API version: 1.56.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tabular

import (
	"encoding/json"
)

// checks if the CreateS3StorageProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateS3StorageProfileResponse{}

// CreateS3StorageProfileResponse struct for CreateS3StorageProfileResponse
type CreateS3StorageProfileResponse struct {
	Id             *string                `json:"id,omitempty"`
	OrganizationId *string                `json:"organizationId,omitempty"`
	AccountId      *string                `json:"accountId,omitempty"`
	Region         *string                `json:"region,omitempty"`
	Bucket         *string                `json:"bucket,omitempty"`
	RoleArn        *string                `json:"roleArn,omitempty"`
	ExternalId     *string                `json:"externalId,omitempty"`
	Properties     map[string]interface{} `json:"properties,omitempty"`
}

// NewCreateS3StorageProfileResponse instantiates a new CreateS3StorageProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateS3StorageProfileResponse() *CreateS3StorageProfileResponse {
	this := CreateS3StorageProfileResponse{}
	return &this
}

// NewCreateS3StorageProfileResponseWithDefaults instantiates a new CreateS3StorageProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateS3StorageProfileResponseWithDefaults() *CreateS3StorageProfileResponse {
	this := CreateS3StorageProfileResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateS3StorageProfileResponse) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *CreateS3StorageProfileResponse) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CreateS3StorageProfileResponse) SetAccountId(v string) {
	o.AccountId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateS3StorageProfileResponse) SetRegion(v string) {
	o.Region = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *CreateS3StorageProfileResponse) SetBucket(v string) {
	o.Bucket = &v
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetRoleArn() string {
	if o == nil || IsNil(o.RoleArn) {
		var ret string
		return ret
	}
	return *o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.RoleArn) {
		return nil, false
	}
	return o.RoleArn, true
}

// HasRoleArn returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasRoleArn() bool {
	if o != nil && !IsNil(o.RoleArn) {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given string and assigns it to the RoleArn field.
func (o *CreateS3StorageProfileResponse) SetRoleArn(v string) {
	o.RoleArn = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CreateS3StorageProfileResponse) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateS3StorageProfileResponse) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3StorageProfileResponse) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateS3StorageProfileResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *CreateS3StorageProfileResponse) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o CreateS3StorageProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateS3StorageProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.RoleArn) {
		toSerialize["roleArn"] = o.RoleArn
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableCreateS3StorageProfileResponse struct {
	value *CreateS3StorageProfileResponse
	isSet bool
}

func (v NullableCreateS3StorageProfileResponse) Get() *CreateS3StorageProfileResponse {
	return v.value
}

func (v *NullableCreateS3StorageProfileResponse) Set(val *CreateS3StorageProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateS3StorageProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateS3StorageProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateS3StorageProfileResponse(val *CreateS3StorageProfileResponse) *NullableCreateS3StorageProfileResponse {
	return &NullableCreateS3StorageProfileResponse{value: val, isSet: true}
}

func (v NullableCreateS3StorageProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateS3StorageProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
