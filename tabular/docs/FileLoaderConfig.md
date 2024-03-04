# FileLoaderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileFormat** | Pointer to **string** |  | [optional] 
**Delimiter** | Pointer to **string** |  | [optional] 
**HasHeader** | Pointer to **bool** |  | [optional] 
**FileFilter** | Pointer to **string** |  | [optional] 
**TableBranchName** | Pointer to **string** |  | [optional] 
**Escape** | Pointer to **string** |  | [optional] 
**DateFormat** | Pointer to **string** |  | [optional] 
**TimestampFormat** | Pointer to **string** |  | [optional] 
**TimestampNTZFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewFileLoaderConfig

`func NewFileLoaderConfig() *FileLoaderConfig`

NewFileLoaderConfig instantiates a new FileLoaderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLoaderConfigWithDefaults

`func NewFileLoaderConfigWithDefaults() *FileLoaderConfig`

NewFileLoaderConfigWithDefaults instantiates a new FileLoaderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileFormat

`func (o *FileLoaderConfig) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *FileLoaderConfig) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *FileLoaderConfig) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *FileLoaderConfig) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetDelimiter

`func (o *FileLoaderConfig) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *FileLoaderConfig) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *FileLoaderConfig) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *FileLoaderConfig) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetHasHeader

`func (o *FileLoaderConfig) GetHasHeader() bool`

GetHasHeader returns the HasHeader field if non-nil, zero value otherwise.

### GetHasHeaderOk

`func (o *FileLoaderConfig) GetHasHeaderOk() (*bool, bool)`

GetHasHeaderOk returns a tuple with the HasHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHeader

`func (o *FileLoaderConfig) SetHasHeader(v bool)`

SetHasHeader sets HasHeader field to given value.

### HasHasHeader

`func (o *FileLoaderConfig) HasHasHeader() bool`

HasHasHeader returns a boolean if a field has been set.

### GetFileFilter

`func (o *FileLoaderConfig) GetFileFilter() string`

GetFileFilter returns the FileFilter field if non-nil, zero value otherwise.

### GetFileFilterOk

`func (o *FileLoaderConfig) GetFileFilterOk() (*string, bool)`

GetFileFilterOk returns a tuple with the FileFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilter

`func (o *FileLoaderConfig) SetFileFilter(v string)`

SetFileFilter sets FileFilter field to given value.

### HasFileFilter

`func (o *FileLoaderConfig) HasFileFilter() bool`

HasFileFilter returns a boolean if a field has been set.

### GetTableBranchName

`func (o *FileLoaderConfig) GetTableBranchName() string`

GetTableBranchName returns the TableBranchName field if non-nil, zero value otherwise.

### GetTableBranchNameOk

`func (o *FileLoaderConfig) GetTableBranchNameOk() (*string, bool)`

GetTableBranchNameOk returns a tuple with the TableBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableBranchName

`func (o *FileLoaderConfig) SetTableBranchName(v string)`

SetTableBranchName sets TableBranchName field to given value.

### HasTableBranchName

`func (o *FileLoaderConfig) HasTableBranchName() bool`

HasTableBranchName returns a boolean if a field has been set.

### GetEscape

`func (o *FileLoaderConfig) GetEscape() string`

GetEscape returns the Escape field if non-nil, zero value otherwise.

### GetEscapeOk

`func (o *FileLoaderConfig) GetEscapeOk() (*string, bool)`

GetEscapeOk returns a tuple with the Escape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscape

`func (o *FileLoaderConfig) SetEscape(v string)`

SetEscape sets Escape field to given value.

### HasEscape

`func (o *FileLoaderConfig) HasEscape() bool`

HasEscape returns a boolean if a field has been set.

### GetDateFormat

`func (o *FileLoaderConfig) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *FileLoaderConfig) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *FileLoaderConfig) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *FileLoaderConfig) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetTimestampFormat

`func (o *FileLoaderConfig) GetTimestampFormat() string`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *FileLoaderConfig) GetTimestampFormatOk() (*string, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *FileLoaderConfig) SetTimestampFormat(v string)`

SetTimestampFormat sets TimestampFormat field to given value.

### HasTimestampFormat

`func (o *FileLoaderConfig) HasTimestampFormat() bool`

HasTimestampFormat returns a boolean if a field has been set.

### GetTimestampNTZFormat

`func (o *FileLoaderConfig) GetTimestampNTZFormat() string`

GetTimestampNTZFormat returns the TimestampNTZFormat field if non-nil, zero value otherwise.

### GetTimestampNTZFormatOk

`func (o *FileLoaderConfig) GetTimestampNTZFormatOk() (*string, bool)`

GetTimestampNTZFormatOk returns a tuple with the TimestampNTZFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampNTZFormat

`func (o *FileLoaderConfig) SetTimestampNTZFormat(v string)`

SetTimestampNTZFormat sets TimestampNTZFormat field to given value.

### HasTimestampNTZFormat

`func (o *FileLoaderConfig) HasTimestampNTZFormat() bool`

HasTimestampNTZFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


