/*
Tabular API

Tabular OpenAPI Definition

API version: 1.54.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tabular

import (
	"encoding/json"
	"time"
)

// checks if the CreateDatabaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDatabaseResponse{}

// CreateDatabaseResponse struct for CreateDatabaseResponse
type CreateDatabaseResponse struct {
	Id             *string            `json:"id,omitempty"`
	WarehouseId    *string            `json:"warehouseId,omitempty"`
	Name           *string            `json:"name,omitempty"`
	CreatedBy      *string            `json:"createdBy,omitempty"`
	CreatedAt      *time.Time         `json:"createdAt,omitempty"`
	LastModifiedBy *string            `json:"lastModifiedBy,omitempty"`
	LastModified   *time.Time         `json:"lastModified,omitempty"`
	Properties     *map[string]string `json:"properties,omitempty"`
}

// NewCreateDatabaseResponse instantiates a new CreateDatabaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabaseResponse() *CreateDatabaseResponse {
	this := CreateDatabaseResponse{}
	return &this
}

// NewCreateDatabaseResponseWithDefaults instantiates a new CreateDatabaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabaseResponseWithDefaults() *CreateDatabaseResponse {
	this := CreateDatabaseResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateDatabaseResponse) SetId(v string) {
	o.Id = &v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetWarehouseId() string {
	if o == nil || IsNil(o.WarehouseId) {
		var ret string
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given string and assigns it to the WarehouseId field.
func (o *CreateDatabaseResponse) SetWarehouseId(v string) {
	o.WarehouseId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDatabaseResponse) SetName(v string) {
	o.Name = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *CreateDatabaseResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateDatabaseResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *CreateDatabaseResponse) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *CreateDatabaseResponse) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *CreateDatabaseResponse) SetProperties(v map[string]string) {
	o.Properties = &v
}

func (o CreateDatabaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDatabaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouseId"] = o.WarehouseId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableCreateDatabaseResponse struct {
	value *CreateDatabaseResponse
	isSet bool
}

func (v NullableCreateDatabaseResponse) Get() *CreateDatabaseResponse {
	return v.value
}

func (v *NullableCreateDatabaseResponse) Set(val *CreateDatabaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabaseResponse(val *CreateDatabaseResponse) *NullableCreateDatabaseResponse {
	return &NullableCreateDatabaseResponse{value: val, isSet: true}
}

func (v NullableCreateDatabaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
