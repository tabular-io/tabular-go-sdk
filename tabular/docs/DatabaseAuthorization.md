# DatabaseAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SubjectId** | Pointer to **string** |  | [optional] 
**Privilege** | Pointer to **string** |  | [optional] 
**WithGrant** | Pointer to **bool** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 

## Methods

### NewDatabaseAuthorization

`func NewDatabaseAuthorization() *DatabaseAuthorization`

NewDatabaseAuthorization instantiates a new DatabaseAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseAuthorizationWithDefaults

`func NewDatabaseAuthorizationWithDefaults() *DatabaseAuthorization`

NewDatabaseAuthorizationWithDefaults instantiates a new DatabaseAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseAuthorization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseAuthorization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseAuthorization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatabaseAuthorization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubjectId

`func (o *DatabaseAuthorization) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *DatabaseAuthorization) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *DatabaseAuthorization) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *DatabaseAuthorization) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetPrivilege

`func (o *DatabaseAuthorization) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *DatabaseAuthorization) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *DatabaseAuthorization) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *DatabaseAuthorization) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetWithGrant

`func (o *DatabaseAuthorization) GetWithGrant() bool`

GetWithGrant returns the WithGrant field if non-nil, zero value otherwise.

### GetWithGrantOk

`func (o *DatabaseAuthorization) GetWithGrantOk() (*bool, bool)`

GetWithGrantOk returns a tuple with the WithGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGrant

`func (o *DatabaseAuthorization) SetWithGrant(v bool)`

SetWithGrant sets WithGrant field to given value.

### HasWithGrant

`func (o *DatabaseAuthorization) HasWithGrant() bool`

HasWithGrant returns a boolean if a field has been set.

### GetResource

`func (o *DatabaseAuthorization) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DatabaseAuthorization) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DatabaseAuthorization) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DatabaseAuthorization) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceType

`func (o *DatabaseAuthorization) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DatabaseAuthorization) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DatabaseAuthorization) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DatabaseAuthorization) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


