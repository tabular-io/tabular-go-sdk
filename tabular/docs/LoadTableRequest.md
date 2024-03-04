# LoadTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**Prefixes** | Pointer to **[]string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**FileLoaderConfig** | Pointer to [**FileLoaderConfig**](FileLoaderConfig.md) |  | [optional] 
**CancelRunningLoad** | Pointer to **bool** |  | [optional] 
**OnCompletionProperties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLoadTableRequest

`func NewLoadTableRequest() *LoadTableRequest`

NewLoadTableRequest instantiates a new LoadTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadTableRequestWithDefaults

`func NewLoadTableRequestWithDefaults() *LoadTableRequest`

NewLoadTableRequestWithDefaults instantiates a new LoadTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *LoadTableRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *LoadTableRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *LoadTableRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *LoadTableRequest) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetPrefixes

`func (o *LoadTableRequest) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *LoadTableRequest) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *LoadTableRequest) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *LoadTableRequest) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetMode

`func (o *LoadTableRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LoadTableRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LoadTableRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *LoadTableRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetFileLoaderConfig

`func (o *LoadTableRequest) GetFileLoaderConfig() FileLoaderConfig`

GetFileLoaderConfig returns the FileLoaderConfig field if non-nil, zero value otherwise.

### GetFileLoaderConfigOk

`func (o *LoadTableRequest) GetFileLoaderConfigOk() (*FileLoaderConfig, bool)`

GetFileLoaderConfigOk returns a tuple with the FileLoaderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLoaderConfig

`func (o *LoadTableRequest) SetFileLoaderConfig(v FileLoaderConfig)`

SetFileLoaderConfig sets FileLoaderConfig field to given value.

### HasFileLoaderConfig

`func (o *LoadTableRequest) HasFileLoaderConfig() bool`

HasFileLoaderConfig returns a boolean if a field has been set.

### GetCancelRunningLoad

`func (o *LoadTableRequest) GetCancelRunningLoad() bool`

GetCancelRunningLoad returns the CancelRunningLoad field if non-nil, zero value otherwise.

### GetCancelRunningLoadOk

`func (o *LoadTableRequest) GetCancelRunningLoadOk() (*bool, bool)`

GetCancelRunningLoadOk returns a tuple with the CancelRunningLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRunningLoad

`func (o *LoadTableRequest) SetCancelRunningLoad(v bool)`

SetCancelRunningLoad sets CancelRunningLoad field to given value.

### HasCancelRunningLoad

`func (o *LoadTableRequest) HasCancelRunningLoad() bool`

HasCancelRunningLoad returns a boolean if a field has been set.

### GetOnCompletionProperties

`func (o *LoadTableRequest) GetOnCompletionProperties() map[string]string`

GetOnCompletionProperties returns the OnCompletionProperties field if non-nil, zero value otherwise.

### GetOnCompletionPropertiesOk

`func (o *LoadTableRequest) GetOnCompletionPropertiesOk() (*map[string]string, bool)`

GetOnCompletionPropertiesOk returns a tuple with the OnCompletionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCompletionProperties

`func (o *LoadTableRequest) SetOnCompletionProperties(v map[string]string)`

SetOnCompletionProperties sets OnCompletionProperties field to given value.

### HasOnCompletionProperties

`func (o *LoadTableRequest) HasOnCompletionProperties() bool`

HasOnCompletionProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


