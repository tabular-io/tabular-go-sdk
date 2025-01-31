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

// checks if the UpdateLabelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLabelRequest{}

// UpdateLabelRequest struct for UpdateLabelRequest
type UpdateLabelRequest struct {
	LabelName            *string  `json:"labelName,omitempty"`
	Description          *string  `json:"description,omitempty"`
	AllowedResourceTypes []string `json:"allowedResourceTypes,omitempty"`
}

// NewUpdateLabelRequest instantiates a new UpdateLabelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLabelRequest() *UpdateLabelRequest {
	this := UpdateLabelRequest{}
	return &this
}

// NewUpdateLabelRequestWithDefaults instantiates a new UpdateLabelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLabelRequestWithDefaults() *UpdateLabelRequest {
	this := UpdateLabelRequest{}
	return &this
}

// GetLabelName returns the LabelName field value if set, zero value otherwise.
func (o *UpdateLabelRequest) GetLabelName() string {
	if o == nil || IsNil(o.LabelName) {
		var ret string
		return ret
	}
	return *o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLabelRequest) GetLabelNameOk() (*string, bool) {
	if o == nil || IsNil(o.LabelName) {
		return nil, false
	}
	return o.LabelName, true
}

// HasLabelName returns a boolean if a field has been set.
func (o *UpdateLabelRequest) HasLabelName() bool {
	if o != nil && !IsNil(o.LabelName) {
		return true
	}

	return false
}

// SetLabelName gets a reference to the given string and assigns it to the LabelName field.
func (o *UpdateLabelRequest) SetLabelName(v string) {
	o.LabelName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateLabelRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLabelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateLabelRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateLabelRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAllowedResourceTypes returns the AllowedResourceTypes field value if set, zero value otherwise.
func (o *UpdateLabelRequest) GetAllowedResourceTypes() []string {
	if o == nil || IsNil(o.AllowedResourceTypes) {
		var ret []string
		return ret
	}
	return o.AllowedResourceTypes
}

// GetAllowedResourceTypesOk returns a tuple with the AllowedResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLabelRequest) GetAllowedResourceTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedResourceTypes) {
		return nil, false
	}
	return o.AllowedResourceTypes, true
}

// HasAllowedResourceTypes returns a boolean if a field has been set.
func (o *UpdateLabelRequest) HasAllowedResourceTypes() bool {
	if o != nil && !IsNil(o.AllowedResourceTypes) {
		return true
	}

	return false
}

// SetAllowedResourceTypes gets a reference to the given []string and assigns it to the AllowedResourceTypes field.
func (o *UpdateLabelRequest) SetAllowedResourceTypes(v []string) {
	o.AllowedResourceTypes = v
}

func (o UpdateLabelRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLabelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelName) {
		toSerialize["labelName"] = o.LabelName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AllowedResourceTypes) {
		toSerialize["allowedResourceTypes"] = o.AllowedResourceTypes
	}
	return toSerialize, nil
}

type NullableUpdateLabelRequest struct {
	value *UpdateLabelRequest
	isSet bool
}

func (v NullableUpdateLabelRequest) Get() *UpdateLabelRequest {
	return v.value
}

func (v *NullableUpdateLabelRequest) Set(val *UpdateLabelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLabelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLabelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLabelRequest(val *UpdateLabelRequest) *NullableUpdateLabelRequest {
	return &NullableUpdateLabelRequest{value: val, isSet: true}
}

func (v NullableUpdateLabelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLabelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
