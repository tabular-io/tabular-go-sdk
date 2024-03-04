/*
Tabular API

Tabular OpenAPI Definition

API version: 1.54.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tabular

import (
	"encoding/json"
)

// checks if the Warehouse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Warehouse{}

// Warehouse struct for Warehouse
type Warehouse struct {
	Id             *string            `json:"id,omitempty"`
	Name           *string            `json:"name,omitempty"`
	Region         *string            `json:"region,omitempty"`
	OrganizationId *string            `json:"organizationId,omitempty"`
	StorageProfile *string            `json:"storageProfile,omitempty"`
	Properties     *map[string]string `json:"properties,omitempty"`
}

// NewWarehouse instantiates a new Warehouse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouse() *Warehouse {
	this := Warehouse{}
	return &this
}

// NewWarehouseWithDefaults instantiates a new Warehouse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseWithDefaults() *Warehouse {
	this := Warehouse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Warehouse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Warehouse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Warehouse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Warehouse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Warehouse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Warehouse) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Warehouse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Warehouse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Warehouse) SetRegion(v string) {
	o.Region = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Warehouse) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Warehouse) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Warehouse) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetStorageProfile returns the StorageProfile field value if set, zero value otherwise.
func (o *Warehouse) GetStorageProfile() string {
	if o == nil || IsNil(o.StorageProfile) {
		var ret string
		return ret
	}
	return *o.StorageProfile
}

// GetStorageProfileOk returns a tuple with the StorageProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetStorageProfileOk() (*string, bool) {
	if o == nil || IsNil(o.StorageProfile) {
		return nil, false
	}
	return o.StorageProfile, true
}

// HasStorageProfile returns a boolean if a field has been set.
func (o *Warehouse) HasStorageProfile() bool {
	if o != nil && !IsNil(o.StorageProfile) {
		return true
	}

	return false
}

// SetStorageProfile gets a reference to the given string and assigns it to the StorageProfile field.
func (o *Warehouse) SetStorageProfile(v string) {
	o.StorageProfile = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Warehouse) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Warehouse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *Warehouse) SetProperties(v map[string]string) {
	o.Properties = &v
}

func (o Warehouse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Warehouse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.StorageProfile) {
		toSerialize["storageProfile"] = o.StorageProfile
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableWarehouse struct {
	value *Warehouse
	isSet bool
}

func (v NullableWarehouse) Get() *Warehouse {
	return v.value
}

func (v *NullableWarehouse) Set(val *Warehouse) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouse) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouse(val *Warehouse) *NullableWarehouse {
	return &NullableWarehouse{value: val, isSet: true}
}

func (v NullableWarehouse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
