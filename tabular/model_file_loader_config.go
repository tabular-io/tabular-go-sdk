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

// checks if the FileLoaderConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileLoaderConfig{}

// FileLoaderConfig struct for FileLoaderConfig
type FileLoaderConfig struct {
	FileFormat         *string `json:"fileFormat,omitempty"`
	Delimiter          *string `json:"delimiter,omitempty"`
	HasHeader          *bool   `json:"hasHeader,omitempty"`
	FileFilter         *string `json:"fileFilter,omitempty"`
	TableBranchName    *string `json:"tableBranchName,omitempty"`
	Escape             *string `json:"escape,omitempty"`
	DateFormat         *string `json:"dateFormat,omitempty"`
	TimestampFormat    *string `json:"timestampFormat,omitempty"`
	TimestampNTZFormat *string `json:"timestampNTZFormat,omitempty"`
}

// NewFileLoaderConfig instantiates a new FileLoaderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLoaderConfig() *FileLoaderConfig {
	this := FileLoaderConfig{}
	return &this
}

// NewFileLoaderConfigWithDefaults instantiates a new FileLoaderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLoaderConfigWithDefaults() *FileLoaderConfig {
	this := FileLoaderConfig{}
	return &this
}

// GetFileFormat returns the FileFormat field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetFileFormat() string {
	if o == nil || IsNil(o.FileFormat) {
		var ret string
		return ret
	}
	return *o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetFileFormatOk() (*string, bool) {
	if o == nil || IsNil(o.FileFormat) {
		return nil, false
	}
	return o.FileFormat, true
}

// HasFileFormat returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasFileFormat() bool {
	if o != nil && !IsNil(o.FileFormat) {
		return true
	}

	return false
}

// SetFileFormat gets a reference to the given string and assigns it to the FileFormat field.
func (o *FileLoaderConfig) SetFileFormat(v string) {
	o.FileFormat = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetDelimiter() string {
	if o == nil || IsNil(o.Delimiter) {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetDelimiterOk() (*string, bool) {
	if o == nil || IsNil(o.Delimiter) {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasDelimiter() bool {
	if o != nil && !IsNil(o.Delimiter) {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *FileLoaderConfig) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetHasHeader returns the HasHeader field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetHasHeader() bool {
	if o == nil || IsNil(o.HasHeader) {
		var ret bool
		return ret
	}
	return *o.HasHeader
}

// GetHasHeaderOk returns a tuple with the HasHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetHasHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.HasHeader) {
		return nil, false
	}
	return o.HasHeader, true
}

// HasHasHeader returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasHasHeader() bool {
	if o != nil && !IsNil(o.HasHeader) {
		return true
	}

	return false
}

// SetHasHeader gets a reference to the given bool and assigns it to the HasHeader field.
func (o *FileLoaderConfig) SetHasHeader(v bool) {
	o.HasHeader = &v
}

// GetFileFilter returns the FileFilter field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetFileFilter() string {
	if o == nil || IsNil(o.FileFilter) {
		var ret string
		return ret
	}
	return *o.FileFilter
}

// GetFileFilterOk returns a tuple with the FileFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetFileFilterOk() (*string, bool) {
	if o == nil || IsNil(o.FileFilter) {
		return nil, false
	}
	return o.FileFilter, true
}

// HasFileFilter returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasFileFilter() bool {
	if o != nil && !IsNil(o.FileFilter) {
		return true
	}

	return false
}

// SetFileFilter gets a reference to the given string and assigns it to the FileFilter field.
func (o *FileLoaderConfig) SetFileFilter(v string) {
	o.FileFilter = &v
}

// GetTableBranchName returns the TableBranchName field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetTableBranchName() string {
	if o == nil || IsNil(o.TableBranchName) {
		var ret string
		return ret
	}
	return *o.TableBranchName
}

// GetTableBranchNameOk returns a tuple with the TableBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetTableBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableBranchName) {
		return nil, false
	}
	return o.TableBranchName, true
}

// HasTableBranchName returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasTableBranchName() bool {
	if o != nil && !IsNil(o.TableBranchName) {
		return true
	}

	return false
}

// SetTableBranchName gets a reference to the given string and assigns it to the TableBranchName field.
func (o *FileLoaderConfig) SetTableBranchName(v string) {
	o.TableBranchName = &v
}

// GetEscape returns the Escape field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetEscape() string {
	if o == nil || IsNil(o.Escape) {
		var ret string
		return ret
	}
	return *o.Escape
}

// GetEscapeOk returns a tuple with the Escape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetEscapeOk() (*string, bool) {
	if o == nil || IsNil(o.Escape) {
		return nil, false
	}
	return o.Escape, true
}

// HasEscape returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasEscape() bool {
	if o != nil && !IsNil(o.Escape) {
		return true
	}

	return false
}

// SetEscape gets a reference to the given string and assigns it to the Escape field.
func (o *FileLoaderConfig) SetEscape(v string) {
	o.Escape = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetDateFormat() string {
	if o == nil || IsNil(o.DateFormat) {
		var ret string
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DateFormat) {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasDateFormat() bool {
	if o != nil && !IsNil(o.DateFormat) {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.
func (o *FileLoaderConfig) SetDateFormat(v string) {
	o.DateFormat = &v
}

// GetTimestampFormat returns the TimestampFormat field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetTimestampFormat() string {
	if o == nil || IsNil(o.TimestampFormat) {
		var ret string
		return ret
	}
	return *o.TimestampFormat
}

// GetTimestampFormatOk returns a tuple with the TimestampFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetTimestampFormatOk() (*string, bool) {
	if o == nil || IsNil(o.TimestampFormat) {
		return nil, false
	}
	return o.TimestampFormat, true
}

// HasTimestampFormat returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasTimestampFormat() bool {
	if o != nil && !IsNil(o.TimestampFormat) {
		return true
	}

	return false
}

// SetTimestampFormat gets a reference to the given string and assigns it to the TimestampFormat field.
func (o *FileLoaderConfig) SetTimestampFormat(v string) {
	o.TimestampFormat = &v
}

// GetTimestampNTZFormat returns the TimestampNTZFormat field value if set, zero value otherwise.
func (o *FileLoaderConfig) GetTimestampNTZFormat() string {
	if o == nil || IsNil(o.TimestampNTZFormat) {
		var ret string
		return ret
	}
	return *o.TimestampNTZFormat
}

// GetTimestampNTZFormatOk returns a tuple with the TimestampNTZFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLoaderConfig) GetTimestampNTZFormatOk() (*string, bool) {
	if o == nil || IsNil(o.TimestampNTZFormat) {
		return nil, false
	}
	return o.TimestampNTZFormat, true
}

// HasTimestampNTZFormat returns a boolean if a field has been set.
func (o *FileLoaderConfig) HasTimestampNTZFormat() bool {
	if o != nil && !IsNil(o.TimestampNTZFormat) {
		return true
	}

	return false
}

// SetTimestampNTZFormat gets a reference to the given string and assigns it to the TimestampNTZFormat field.
func (o *FileLoaderConfig) SetTimestampNTZFormat(v string) {
	o.TimestampNTZFormat = &v
}

func (o FileLoaderConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLoaderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileFormat) {
		toSerialize["fileFormat"] = o.FileFormat
	}
	if !IsNil(o.Delimiter) {
		toSerialize["delimiter"] = o.Delimiter
	}
	if !IsNil(o.HasHeader) {
		toSerialize["hasHeader"] = o.HasHeader
	}
	if !IsNil(o.FileFilter) {
		toSerialize["fileFilter"] = o.FileFilter
	}
	if !IsNil(o.TableBranchName) {
		toSerialize["tableBranchName"] = o.TableBranchName
	}
	if !IsNil(o.Escape) {
		toSerialize["escape"] = o.Escape
	}
	if !IsNil(o.DateFormat) {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if !IsNil(o.TimestampFormat) {
		toSerialize["timestampFormat"] = o.TimestampFormat
	}
	if !IsNil(o.TimestampNTZFormat) {
		toSerialize["timestampNTZFormat"] = o.TimestampNTZFormat
	}
	return toSerialize, nil
}

type NullableFileLoaderConfig struct {
	value *FileLoaderConfig
	isSet bool
}

func (v NullableFileLoaderConfig) Get() *FileLoaderConfig {
	return v.value
}

func (v *NullableFileLoaderConfig) Set(val *FileLoaderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLoaderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLoaderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLoaderConfig(val *FileLoaderConfig) *NullableFileLoaderConfig {
	return &NullableFileLoaderConfig{value: val, isSet: true}
}

func (v NullableFileLoaderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLoaderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
