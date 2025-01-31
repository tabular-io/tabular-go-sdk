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

// checks if the CreateTableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTableRequest{}

// CreateTableRequest struct for CreateTableRequest
type CreateTableRequest struct {
	TableName               *string            `json:"tableName,omitempty"`
	Bucket                  *string            `json:"bucket,omitempty"`
	Prefixes                []string           `json:"prefixes,omitempty"`
	Mode                    *string            `json:"mode,omitempty"`
	FileLoaderConfig        *FileLoaderConfig  `json:"fileLoaderConfig,omitempty"`
	UseMaps                 *bool              `json:"useMaps,omitempty"`
	OnCompletionProperties  *map[string]string `json:"onCompletionProperties,omitempty"`
	InferPartitionsFromPath *bool              `json:"inferPartitionsFromPath,omitempty"`
}

// NewCreateTableRequest instantiates a new CreateTableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTableRequest() *CreateTableRequest {
	this := CreateTableRequest{}
	return &this
}

// NewCreateTableRequestWithDefaults instantiates a new CreateTableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTableRequestWithDefaults() *CreateTableRequest {
	this := CreateTableRequest{}
	return &this
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *CreateTableRequest) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *CreateTableRequest) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *CreateTableRequest) SetTableName(v string) {
	o.TableName = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *CreateTableRequest) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *CreateTableRequest) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *CreateTableRequest) SetBucket(v string) {
	o.Bucket = &v
}

// GetPrefixes returns the Prefixes field value if set, zero value otherwise.
func (o *CreateTableRequest) GetPrefixes() []string {
	if o == nil || IsNil(o.Prefixes) {
		var ret []string
		return ret
	}
	return o.Prefixes
}

// GetPrefixesOk returns a tuple with the Prefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetPrefixesOk() ([]string, bool) {
	if o == nil || IsNil(o.Prefixes) {
		return nil, false
	}
	return o.Prefixes, true
}

// HasPrefixes returns a boolean if a field has been set.
func (o *CreateTableRequest) HasPrefixes() bool {
	if o != nil && !IsNil(o.Prefixes) {
		return true
	}

	return false
}

// SetPrefixes gets a reference to the given []string and assigns it to the Prefixes field.
func (o *CreateTableRequest) SetPrefixes(v []string) {
	o.Prefixes = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CreateTableRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CreateTableRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *CreateTableRequest) SetMode(v string) {
	o.Mode = &v
}

// GetFileLoaderConfig returns the FileLoaderConfig field value if set, zero value otherwise.
func (o *CreateTableRequest) GetFileLoaderConfig() FileLoaderConfig {
	if o == nil || IsNil(o.FileLoaderConfig) {
		var ret FileLoaderConfig
		return ret
	}
	return *o.FileLoaderConfig
}

// GetFileLoaderConfigOk returns a tuple with the FileLoaderConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetFileLoaderConfigOk() (*FileLoaderConfig, bool) {
	if o == nil || IsNil(o.FileLoaderConfig) {
		return nil, false
	}
	return o.FileLoaderConfig, true
}

// HasFileLoaderConfig returns a boolean if a field has been set.
func (o *CreateTableRequest) HasFileLoaderConfig() bool {
	if o != nil && !IsNil(o.FileLoaderConfig) {
		return true
	}

	return false
}

// SetFileLoaderConfig gets a reference to the given FileLoaderConfig and assigns it to the FileLoaderConfig field.
func (o *CreateTableRequest) SetFileLoaderConfig(v FileLoaderConfig) {
	o.FileLoaderConfig = &v
}

// GetUseMaps returns the UseMaps field value if set, zero value otherwise.
func (o *CreateTableRequest) GetUseMaps() bool {
	if o == nil || IsNil(o.UseMaps) {
		var ret bool
		return ret
	}
	return *o.UseMaps
}

// GetUseMapsOk returns a tuple with the UseMaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetUseMapsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseMaps) {
		return nil, false
	}
	return o.UseMaps, true
}

// HasUseMaps returns a boolean if a field has been set.
func (o *CreateTableRequest) HasUseMaps() bool {
	if o != nil && !IsNil(o.UseMaps) {
		return true
	}

	return false
}

// SetUseMaps gets a reference to the given bool and assigns it to the UseMaps field.
func (o *CreateTableRequest) SetUseMaps(v bool) {
	o.UseMaps = &v
}

// GetOnCompletionProperties returns the OnCompletionProperties field value if set, zero value otherwise.
func (o *CreateTableRequest) GetOnCompletionProperties() map[string]string {
	if o == nil || IsNil(o.OnCompletionProperties) {
		var ret map[string]string
		return ret
	}
	return *o.OnCompletionProperties
}

// GetOnCompletionPropertiesOk returns a tuple with the OnCompletionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetOnCompletionPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.OnCompletionProperties) {
		return nil, false
	}
	return o.OnCompletionProperties, true
}

// HasOnCompletionProperties returns a boolean if a field has been set.
func (o *CreateTableRequest) HasOnCompletionProperties() bool {
	if o != nil && !IsNil(o.OnCompletionProperties) {
		return true
	}

	return false
}

// SetOnCompletionProperties gets a reference to the given map[string]string and assigns it to the OnCompletionProperties field.
func (o *CreateTableRequest) SetOnCompletionProperties(v map[string]string) {
	o.OnCompletionProperties = &v
}

// GetInferPartitionsFromPath returns the InferPartitionsFromPath field value if set, zero value otherwise.
func (o *CreateTableRequest) GetInferPartitionsFromPath() bool {
	if o == nil || IsNil(o.InferPartitionsFromPath) {
		var ret bool
		return ret
	}
	return *o.InferPartitionsFromPath
}

// GetInferPartitionsFromPathOk returns a tuple with the InferPartitionsFromPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequest) GetInferPartitionsFromPathOk() (*bool, bool) {
	if o == nil || IsNil(o.InferPartitionsFromPath) {
		return nil, false
	}
	return o.InferPartitionsFromPath, true
}

// HasInferPartitionsFromPath returns a boolean if a field has been set.
func (o *CreateTableRequest) HasInferPartitionsFromPath() bool {
	if o != nil && !IsNil(o.InferPartitionsFromPath) {
		return true
	}

	return false
}

// SetInferPartitionsFromPath gets a reference to the given bool and assigns it to the InferPartitionsFromPath field.
func (o *CreateTableRequest) SetInferPartitionsFromPath(v bool) {
	o.InferPartitionsFromPath = &v
}

func (o CreateTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TableName) {
		toSerialize["tableName"] = o.TableName
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Prefixes) {
		toSerialize["prefixes"] = o.Prefixes
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.FileLoaderConfig) {
		toSerialize["fileLoaderConfig"] = o.FileLoaderConfig
	}
	if !IsNil(o.UseMaps) {
		toSerialize["useMaps"] = o.UseMaps
	}
	if !IsNil(o.OnCompletionProperties) {
		toSerialize["onCompletionProperties"] = o.OnCompletionProperties
	}
	if !IsNil(o.InferPartitionsFromPath) {
		toSerialize["inferPartitionsFromPath"] = o.InferPartitionsFromPath
	}
	return toSerialize, nil
}

type NullableCreateTableRequest struct {
	value *CreateTableRequest
	isSet bool
}

func (v NullableCreateTableRequest) Get() *CreateTableRequest {
	return v.value
}

func (v *NullableCreateTableRequest) Set(val *CreateTableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTableRequest(val *CreateTableRequest) *NullableCreateTableRequest {
	return &NullableCreateTableRequest{value: val, isSet: true}
}

func (v NullableCreateTableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
