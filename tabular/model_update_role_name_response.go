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

// checks if the UpdateRoleNameResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleNameResponse{}

// UpdateRoleNameResponse struct for UpdateRoleNameResponse
type UpdateRoleNameResponse struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUpdateRoleNameResponse instantiates a new UpdateRoleNameResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleNameResponse() *UpdateRoleNameResponse {
	this := UpdateRoleNameResponse{}
	return &this
}

// NewUpdateRoleNameResponseWithDefaults instantiates a new UpdateRoleNameResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleNameResponseWithDefaults() *UpdateRoleNameResponse {
	this := UpdateRoleNameResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateRoleNameResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleNameResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateRoleNameResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateRoleNameResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRoleNameResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleNameResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRoleNameResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRoleNameResponse) SetName(v string) {
	o.Name = &v
}

func (o UpdateRoleNameResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleNameResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateRoleNameResponse struct {
	value *UpdateRoleNameResponse
	isSet bool
}

func (v NullableUpdateRoleNameResponse) Get() *UpdateRoleNameResponse {
	return v.value
}

func (v *NullableUpdateRoleNameResponse) Set(val *UpdateRoleNameResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleNameResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleNameResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleNameResponse(val *UpdateRoleNameResponse) *NullableUpdateRoleNameResponse {
	return &NullableUpdateRoleNameResponse{value: val, isSet: true}
}

func (v NullableUpdateRoleNameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleNameResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
