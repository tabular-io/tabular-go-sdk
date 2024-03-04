# LoadTableStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableLoadId** | Pointer to **string** |  | [optional] 
**WarehouseId** | Pointer to **string** |  | [optional] 
**TableRefId** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**LoadTableState** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**Attempt** | Pointer to **int32** |  | [optional] 
**TotalBytes** | Pointer to **int64** |  | [optional] 
**TotalLoadedBytes** | Pointer to **int64** |  | [optional] 
**TotalFileCount** | Pointer to **int32** |  | [optional] 
**TotalLoadedFileCount** | Pointer to **int32** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpireAt** | Pointer to **time.Time** |  | [optional] 
**TimeoutAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLoadTableStatus

`func NewLoadTableStatus() *LoadTableStatus`

NewLoadTableStatus instantiates a new LoadTableStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadTableStatusWithDefaults

`func NewLoadTableStatusWithDefaults() *LoadTableStatus`

NewLoadTableStatusWithDefaults instantiates a new LoadTableStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableLoadId

`func (o *LoadTableStatus) GetTableLoadId() string`

GetTableLoadId returns the TableLoadId field if non-nil, zero value otherwise.

### GetTableLoadIdOk

`func (o *LoadTableStatus) GetTableLoadIdOk() (*string, bool)`

GetTableLoadIdOk returns a tuple with the TableLoadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableLoadId

`func (o *LoadTableStatus) SetTableLoadId(v string)`

SetTableLoadId sets TableLoadId field to given value.

### HasTableLoadId

`func (o *LoadTableStatus) HasTableLoadId() bool`

HasTableLoadId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *LoadTableStatus) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *LoadTableStatus) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *LoadTableStatus) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *LoadTableStatus) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetTableRefId

`func (o *LoadTableStatus) GetTableRefId() string`

GetTableRefId returns the TableRefId field if non-nil, zero value otherwise.

### GetTableRefIdOk

`func (o *LoadTableStatus) GetTableRefIdOk() (*string, bool)`

GetTableRefIdOk returns a tuple with the TableRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableRefId

`func (o *LoadTableStatus) SetTableRefId(v string)`

SetTableRefId sets TableRefId field to given value.

### HasTableRefId

`func (o *LoadTableStatus) HasTableRefId() bool`

HasTableRefId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *LoadTableStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *LoadTableStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *LoadTableStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *LoadTableStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetState

`func (o *LoadTableStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoadTableStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoadTableStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LoadTableStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLoadTableState

`func (o *LoadTableStatus) GetLoadTableState() string`

GetLoadTableState returns the LoadTableState field if non-nil, zero value otherwise.

### GetLoadTableStateOk

`func (o *LoadTableStatus) GetLoadTableStateOk() (*string, bool)`

GetLoadTableStateOk returns a tuple with the LoadTableState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTableState

`func (o *LoadTableStatus) SetLoadTableState(v string)`

SetLoadTableState sets LoadTableState field to given value.

### HasLoadTableState

`func (o *LoadTableStatus) HasLoadTableState() bool`

HasLoadTableState returns a boolean if a field has been set.

### GetProgress

`func (o *LoadTableStatus) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *LoadTableStatus) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *LoadTableStatus) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *LoadTableStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetAttempt

`func (o *LoadTableStatus) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *LoadTableStatus) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *LoadTableStatus) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *LoadTableStatus) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetTotalBytes

`func (o *LoadTableStatus) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *LoadTableStatus) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *LoadTableStatus) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *LoadTableStatus) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetTotalLoadedBytes

`func (o *LoadTableStatus) GetTotalLoadedBytes() int64`

GetTotalLoadedBytes returns the TotalLoadedBytes field if non-nil, zero value otherwise.

### GetTotalLoadedBytesOk

`func (o *LoadTableStatus) GetTotalLoadedBytesOk() (*int64, bool)`

GetTotalLoadedBytesOk returns a tuple with the TotalLoadedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLoadedBytes

`func (o *LoadTableStatus) SetTotalLoadedBytes(v int64)`

SetTotalLoadedBytes sets TotalLoadedBytes field to given value.

### HasTotalLoadedBytes

`func (o *LoadTableStatus) HasTotalLoadedBytes() bool`

HasTotalLoadedBytes returns a boolean if a field has been set.

### GetTotalFileCount

`func (o *LoadTableStatus) GetTotalFileCount() int32`

GetTotalFileCount returns the TotalFileCount field if non-nil, zero value otherwise.

### GetTotalFileCountOk

`func (o *LoadTableStatus) GetTotalFileCountOk() (*int32, bool)`

GetTotalFileCountOk returns a tuple with the TotalFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFileCount

`func (o *LoadTableStatus) SetTotalFileCount(v int32)`

SetTotalFileCount sets TotalFileCount field to given value.

### HasTotalFileCount

`func (o *LoadTableStatus) HasTotalFileCount() bool`

HasTotalFileCount returns a boolean if a field has been set.

### GetTotalLoadedFileCount

`func (o *LoadTableStatus) GetTotalLoadedFileCount() int32`

GetTotalLoadedFileCount returns the TotalLoadedFileCount field if non-nil, zero value otherwise.

### GetTotalLoadedFileCountOk

`func (o *LoadTableStatus) GetTotalLoadedFileCountOk() (*int32, bool)`

GetTotalLoadedFileCountOk returns a tuple with the TotalLoadedFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLoadedFileCount

`func (o *LoadTableStatus) SetTotalLoadedFileCount(v int32)`

SetTotalLoadedFileCount sets TotalLoadedFileCount field to given value.

### HasTotalLoadedFileCount

`func (o *LoadTableStatus) HasTotalLoadedFileCount() bool`

HasTotalLoadedFileCount returns a boolean if a field has been set.

### GetLastModified

`func (o *LoadTableStatus) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *LoadTableStatus) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *LoadTableStatus) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *LoadTableStatus) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoadTableStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadTableStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadTableStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadTableStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpireAt

`func (o *LoadTableStatus) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *LoadTableStatus) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *LoadTableStatus) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *LoadTableStatus) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetTimeoutAt

`func (o *LoadTableStatus) GetTimeoutAt() time.Time`

GetTimeoutAt returns the TimeoutAt field if non-nil, zero value otherwise.

### GetTimeoutAtOk

`func (o *LoadTableStatus) GetTimeoutAtOk() (*time.Time, bool)`

GetTimeoutAtOk returns a tuple with the TimeoutAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutAt

`func (o *LoadTableStatus) SetTimeoutAt(v time.Time)`

SetTimeoutAt sets TimeoutAt field to given value.

### HasTimeoutAt

`func (o *LoadTableStatus) HasTimeoutAt() bool`

HasTimeoutAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


