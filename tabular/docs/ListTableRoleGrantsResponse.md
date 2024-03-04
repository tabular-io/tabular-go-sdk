# ListTableRoleGrantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | Pointer to [**[]RoleGrantDetail**](RoleGrantDetail.md) |  | [optional] 

## Methods

### NewListTableRoleGrantsResponse

`func NewListTableRoleGrantsResponse() *ListTableRoleGrantsResponse`

NewListTableRoleGrantsResponse instantiates a new ListTableRoleGrantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTableRoleGrantsResponseWithDefaults

`func NewListTableRoleGrantsResponseWithDefaults() *ListTableRoleGrantsResponse`

NewListTableRoleGrantsResponseWithDefaults instantiates a new ListTableRoleGrantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *ListTableRoleGrantsResponse) GetGrants() []RoleGrantDetail`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *ListTableRoleGrantsResponse) GetGrantsOk() (*[]RoleGrantDetail, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *ListTableRoleGrantsResponse) SetGrants(v []RoleGrantDetail)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *ListTableRoleGrantsResponse) HasGrants() bool`

HasGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


