# CreateLabelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AllowedResourceTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateLabelRequest

`func NewCreateLabelRequest() *CreateLabelRequest`

NewCreateLabelRequest instantiates a new CreateLabelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLabelRequestWithDefaults

`func NewCreateLabelRequestWithDefaults() *CreateLabelRequest`

NewCreateLabelRequestWithDefaults instantiates a new CreateLabelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelName

`func (o *CreateLabelRequest) GetLabelName() string`

GetLabelName returns the LabelName field if non-nil, zero value otherwise.

### GetLabelNameOk

`func (o *CreateLabelRequest) GetLabelNameOk() (*string, bool)`

GetLabelNameOk returns a tuple with the LabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelName

`func (o *CreateLabelRequest) SetLabelName(v string)`

SetLabelName sets LabelName field to given value.

### HasLabelName

`func (o *CreateLabelRequest) HasLabelName() bool`

HasLabelName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLabelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLabelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLabelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLabelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllowedResourceTypes

`func (o *CreateLabelRequest) GetAllowedResourceTypes() []string`

GetAllowedResourceTypes returns the AllowedResourceTypes field if non-nil, zero value otherwise.

### GetAllowedResourceTypesOk

`func (o *CreateLabelRequest) GetAllowedResourceTypesOk() (*[]string, bool)`

GetAllowedResourceTypesOk returns a tuple with the AllowedResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResourceTypes

`func (o *CreateLabelRequest) SetAllowedResourceTypes(v []string)`

SetAllowedResourceTypes sets AllowedResourceTypes field to given value.

### HasAllowedResourceTypes

`func (o *CreateLabelRequest) HasAllowedResourceTypes() bool`

HasAllowedResourceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


