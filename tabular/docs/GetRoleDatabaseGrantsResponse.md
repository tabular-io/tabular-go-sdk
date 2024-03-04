# GetRoleDatabaseGrantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorizations** | Pointer to [**[]DatabaseAuthorization**](DatabaseAuthorization.md) |  | [optional] 

## Methods

### NewGetRoleDatabaseGrantsResponse

`func NewGetRoleDatabaseGrantsResponse() *GetRoleDatabaseGrantsResponse`

NewGetRoleDatabaseGrantsResponse instantiates a new GetRoleDatabaseGrantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleDatabaseGrantsResponseWithDefaults

`func NewGetRoleDatabaseGrantsResponseWithDefaults() *GetRoleDatabaseGrantsResponse`

NewGetRoleDatabaseGrantsResponseWithDefaults instantiates a new GetRoleDatabaseGrantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizations

`func (o *GetRoleDatabaseGrantsResponse) GetAuthorizations() []DatabaseAuthorization`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *GetRoleDatabaseGrantsResponse) GetAuthorizationsOk() (*[]DatabaseAuthorization, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *GetRoleDatabaseGrantsResponse) SetAuthorizations(v []DatabaseAuthorization)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *GetRoleDatabaseGrantsResponse) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


