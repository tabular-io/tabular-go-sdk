# CreateTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**Prefixes** | Pointer to **[]string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**FileLoaderConfig** | Pointer to [**FileLoaderConfig**](FileLoaderConfig.md) |  | [optional] 
**UseMaps** | Pointer to **bool** |  | [optional] 
**OnCompletionProperties** | Pointer to **map[string]string** |  | [optional] 
**InferPartitionsFromPath** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateTableRequest

`func NewCreateTableRequest() *CreateTableRequest`

NewCreateTableRequest instantiates a new CreateTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTableRequestWithDefaults

`func NewCreateTableRequestWithDefaults() *CreateTableRequest`

NewCreateTableRequestWithDefaults instantiates a new CreateTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *CreateTableRequest) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *CreateTableRequest) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *CreateTableRequest) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *CreateTableRequest) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetBucket

`func (o *CreateTableRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CreateTableRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CreateTableRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CreateTableRequest) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetPrefixes

`func (o *CreateTableRequest) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *CreateTableRequest) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *CreateTableRequest) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *CreateTableRequest) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetMode

`func (o *CreateTableRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateTableRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateTableRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateTableRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetFileLoaderConfig

`func (o *CreateTableRequest) GetFileLoaderConfig() FileLoaderConfig`

GetFileLoaderConfig returns the FileLoaderConfig field if non-nil, zero value otherwise.

### GetFileLoaderConfigOk

`func (o *CreateTableRequest) GetFileLoaderConfigOk() (*FileLoaderConfig, bool)`

GetFileLoaderConfigOk returns a tuple with the FileLoaderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLoaderConfig

`func (o *CreateTableRequest) SetFileLoaderConfig(v FileLoaderConfig)`

SetFileLoaderConfig sets FileLoaderConfig field to given value.

### HasFileLoaderConfig

`func (o *CreateTableRequest) HasFileLoaderConfig() bool`

HasFileLoaderConfig returns a boolean if a field has been set.

### GetUseMaps

`func (o *CreateTableRequest) GetUseMaps() bool`

GetUseMaps returns the UseMaps field if non-nil, zero value otherwise.

### GetUseMapsOk

`func (o *CreateTableRequest) GetUseMapsOk() (*bool, bool)`

GetUseMapsOk returns a tuple with the UseMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaps

`func (o *CreateTableRequest) SetUseMaps(v bool)`

SetUseMaps sets UseMaps field to given value.

### HasUseMaps

`func (o *CreateTableRequest) HasUseMaps() bool`

HasUseMaps returns a boolean if a field has been set.

### GetOnCompletionProperties

`func (o *CreateTableRequest) GetOnCompletionProperties() map[string]string`

GetOnCompletionProperties returns the OnCompletionProperties field if non-nil, zero value otherwise.

### GetOnCompletionPropertiesOk

`func (o *CreateTableRequest) GetOnCompletionPropertiesOk() (*map[string]string, bool)`

GetOnCompletionPropertiesOk returns a tuple with the OnCompletionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCompletionProperties

`func (o *CreateTableRequest) SetOnCompletionProperties(v map[string]string)`

SetOnCompletionProperties sets OnCompletionProperties field to given value.

### HasOnCompletionProperties

`func (o *CreateTableRequest) HasOnCompletionProperties() bool`

HasOnCompletionProperties returns a boolean if a field has been set.

### GetInferPartitionsFromPath

`func (o *CreateTableRequest) GetInferPartitionsFromPath() bool`

GetInferPartitionsFromPath returns the InferPartitionsFromPath field if non-nil, zero value otherwise.

### GetInferPartitionsFromPathOk

`func (o *CreateTableRequest) GetInferPartitionsFromPathOk() (*bool, bool)`

GetInferPartitionsFromPathOk returns a tuple with the InferPartitionsFromPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferPartitionsFromPath

`func (o *CreateTableRequest) SetInferPartitionsFromPath(v bool)`

SetInferPartitionsFromPath sets InferPartitionsFromPath field to given value.

### HasInferPartitionsFromPath

`func (o *CreateTableRequest) HasInferPartitionsFromPath() bool`

HasInferPartitionsFromPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


