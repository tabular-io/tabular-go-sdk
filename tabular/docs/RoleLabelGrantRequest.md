# RoleLabelGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | Pointer to **string** |  | [optional] 
**Privilege** | Pointer to **string** |  | [optional] 
**WithGrant** | Pointer to **bool** |  | [optional] 

## Methods

### NewRoleLabelGrantRequest

`func NewRoleLabelGrantRequest() *RoleLabelGrantRequest`

NewRoleLabelGrantRequest instantiates a new RoleLabelGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleLabelGrantRequestWithDefaults

`func NewRoleLabelGrantRequestWithDefaults() *RoleLabelGrantRequest`

NewRoleLabelGrantRequestWithDefaults instantiates a new RoleLabelGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *RoleLabelGrantRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleLabelGrantRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleLabelGrantRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RoleLabelGrantRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetPrivilege

`func (o *RoleLabelGrantRequest) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *RoleLabelGrantRequest) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *RoleLabelGrantRequest) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *RoleLabelGrantRequest) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetWithGrant

`func (o *RoleLabelGrantRequest) GetWithGrant() bool`

GetWithGrant returns the WithGrant field if non-nil, zero value otherwise.

### GetWithGrantOk

`func (o *RoleLabelGrantRequest) GetWithGrantOk() (*bool, bool)`

GetWithGrantOk returns a tuple with the WithGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGrant

`func (o *RoleLabelGrantRequest) SetWithGrant(v bool)`

SetWithGrant sets WithGrant field to given value.

### HasWithGrant

`func (o *RoleLabelGrantRequest) HasWithGrant() bool`

HasWithGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


