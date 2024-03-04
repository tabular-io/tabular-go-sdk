# Label

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AllowedResourceTypes** | Pointer to **[]string** |  | [optional] 
**MaskMode** | Pointer to **string** |  | [optional] 

## Methods

### NewLabel

`func NewLabel() *Label`

NewLabel instantiates a new Label object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelWithDefaults

`func NewLabelWithDefaults() *Label`

NewLabelWithDefaults instantiates a new Label object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Label) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Label) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Label) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Label) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Label) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Label) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Label) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Label) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Label) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Label) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Label) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Label) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllowedResourceTypes

`func (o *Label) GetAllowedResourceTypes() []string`

GetAllowedResourceTypes returns the AllowedResourceTypes field if non-nil, zero value otherwise.

### GetAllowedResourceTypesOk

`func (o *Label) GetAllowedResourceTypesOk() (*[]string, bool)`

GetAllowedResourceTypesOk returns a tuple with the AllowedResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResourceTypes

`func (o *Label) SetAllowedResourceTypes(v []string)`

SetAllowedResourceTypes sets AllowedResourceTypes field to given value.

### HasAllowedResourceTypes

`func (o *Label) HasAllowedResourceTypes() bool`

HasAllowedResourceTypes returns a boolean if a field has been set.

### GetMaskMode

`func (o *Label) GetMaskMode() string`

GetMaskMode returns the MaskMode field if non-nil, zero value otherwise.

### GetMaskModeOk

`func (o *Label) GetMaskModeOk() (*string, bool)`

GetMaskModeOk returns a tuple with the MaskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskMode

`func (o *Label) SetMaskMode(v string)`

SetMaskMode sets MaskMode field to given value.

### HasMaskMode

`func (o *Label) HasMaskMode() bool`

HasMaskMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


