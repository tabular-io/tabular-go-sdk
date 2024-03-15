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

// checks if the RoleTableGrantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleTableGrantRequest{}

// RoleTableGrantRequest struct for RoleTableGrantRequest
type RoleTableGrantRequest struct {
	RoleId    *string `json:"roleId,omitempty"`
	Privilege *string `json:"privilege,omitempty"`
	WithGrant *bool   `json:"withGrant,omitempty"`
}

// NewRoleTableGrantRequest instantiates a new RoleTableGrantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleTableGrantRequest() *RoleTableGrantRequest {
	this := RoleTableGrantRequest{}
	return &this
}

// NewRoleTableGrantRequestWithDefaults instantiates a new RoleTableGrantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleTableGrantRequestWithDefaults() *RoleTableGrantRequest {
	this := RoleTableGrantRequest{}
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RoleTableGrantRequest) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleTableGrantRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoleTableGrantRequest) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *RoleTableGrantRequest) SetRoleId(v string) {
	o.RoleId = &v
}

// GetPrivilege returns the Privilege field value if set, zero value otherwise.
func (o *RoleTableGrantRequest) GetPrivilege() string {
	if o == nil || IsNil(o.Privilege) {
		var ret string
		return ret
	}
	return *o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleTableGrantRequest) GetPrivilegeOk() (*string, bool) {
	if o == nil || IsNil(o.Privilege) {
		return nil, false
	}
	return o.Privilege, true
}

// HasPrivilege returns a boolean if a field has been set.
func (o *RoleTableGrantRequest) HasPrivilege() bool {
	if o != nil && !IsNil(o.Privilege) {
		return true
	}

	return false
}

// SetPrivilege gets a reference to the given string and assigns it to the Privilege field.
func (o *RoleTableGrantRequest) SetPrivilege(v string) {
	o.Privilege = &v
}

// GetWithGrant returns the WithGrant field value if set, zero value otherwise.
func (o *RoleTableGrantRequest) GetWithGrant() bool {
	if o == nil || IsNil(o.WithGrant) {
		var ret bool
		return ret
	}
	return *o.WithGrant
}

// GetWithGrantOk returns a tuple with the WithGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleTableGrantRequest) GetWithGrantOk() (*bool, bool) {
	if o == nil || IsNil(o.WithGrant) {
		return nil, false
	}
	return o.WithGrant, true
}

// HasWithGrant returns a boolean if a field has been set.
func (o *RoleTableGrantRequest) HasWithGrant() bool {
	if o != nil && !IsNil(o.WithGrant) {
		return true
	}

	return false
}

// SetWithGrant gets a reference to the given bool and assigns it to the WithGrant field.
func (o *RoleTableGrantRequest) SetWithGrant(v bool) {
	o.WithGrant = &v
}

func (o RoleTableGrantRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleTableGrantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !IsNil(o.Privilege) {
		toSerialize["privilege"] = o.Privilege
	}
	if !IsNil(o.WithGrant) {
		toSerialize["withGrant"] = o.WithGrant
	}
	return toSerialize, nil
}

type NullableRoleTableGrantRequest struct {
	value *RoleTableGrantRequest
	isSet bool
}

func (v NullableRoleTableGrantRequest) Get() *RoleTableGrantRequest {
	return v.value
}

func (v *NullableRoleTableGrantRequest) Set(val *RoleTableGrantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleTableGrantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleTableGrantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleTableGrantRequest(val *RoleTableGrantRequest) *NullableRoleTableGrantRequest {
	return &NullableRoleTableGrantRequest{value: val, isSet: true}
}

func (v NullableRoleTableGrantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleTableGrantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
