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

// checks if the CreateWarehouseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWarehouseResponse{}

// CreateWarehouseResponse struct for CreateWarehouseResponse
type CreateWarehouseResponse struct {
	Id             *string            `json:"id,omitempty"`
	Name           *string            `json:"name,omitempty"`
	Region         *string            `json:"region,omitempty"`
	OrganizationId *string            `json:"organizationId,omitempty"`
	StorageProfile *string            `json:"storageProfile,omitempty"`
	Properties     *map[string]string `json:"properties,omitempty"`
}

// NewCreateWarehouseResponse instantiates a new CreateWarehouseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWarehouseResponse() *CreateWarehouseResponse {
	this := CreateWarehouseResponse{}
	return &this
}

// NewCreateWarehouseResponseWithDefaults instantiates a new CreateWarehouseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWarehouseResponseWithDefaults() *CreateWarehouseResponse {
	this := CreateWarehouseResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateWarehouseResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateWarehouseResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateWarehouseResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateWarehouseResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWarehouseResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateWarehouseResponse) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateWarehouseResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateWarehouseResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateWarehouseResponse) SetRegion(v string) {
	o.Region = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *CreateWarehouseResponse) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateWarehouseResponse) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *CreateWarehouseResponse) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetStorageProfile returns the StorageProfile field value if set, zero value otherwise.
func (o *CreateWarehouseResponse) GetStorageProfile() string {
	if o == nil || IsNil(o.StorageProfile) {
		var ret string
		return ret
	}
	return *o.StorageProfile
}

// GetStorageProfileOk returns a tuple with the StorageProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseResponse) GetStorageProfileOk() (*string, bool) {
	if o == nil || IsNil(o.StorageProfile) {
		return nil, false
	}
	return o.StorageProfile, true
}

// HasStorageProfile returns a boolean if a field has been set.
func (o *CreateWarehouseResponse) HasStorageProfile() bool {
	if o != nil && !IsNil(o.StorageProfile) {
		return true
	}

	return false
}

// SetStorageProfile gets a reference to the given string and assigns it to the StorageProfile field.
func (o *CreateWarehouseResponse) SetStorageProfile(v string) {
	o.StorageProfile = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateWarehouseResponse) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWarehouseResponse) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateWarehouseResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *CreateWarehouseResponse) SetProperties(v map[string]string) {
	o.Properties = &v
}

func (o CreateWarehouseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWarehouseResponse) ToMap() (map[string]interface{}, error) {
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

type NullableCreateWarehouseResponse struct {
	value *CreateWarehouseResponse
	isSet bool
}

func (v NullableCreateWarehouseResponse) Get() *CreateWarehouseResponse {
	return v.value
}

func (v *NullableCreateWarehouseResponse) Set(val *CreateWarehouseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWarehouseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWarehouseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWarehouseResponse(val *CreateWarehouseResponse) *NullableCreateWarehouseResponse {
	return &NullableCreateWarehouseResponse{value: val, isSet: true}
}

func (v NullableCreateWarehouseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWarehouseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
