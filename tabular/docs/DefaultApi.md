# \DefaultAPI

All URIs are relative to *https://api.tabular.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChildToRole**](DefaultAPI.md#AddChildToRole) | **Put** /v1/organizations/{organizationId}/roles/{roleName}/children | Add child to role
[**AddRoleMembers**](DefaultAPI.md#AddRoleMembers) | **Put** /v1/organizations/{organizationId}/roles/{roleName}/members | Add members to a role
[**CreateCustomerIdentityProviderCredential**](DefaultAPI.md#CreateCustomerIdentityProviderCredential) | **Post** /v1/organizations/{organizationId}/iam/credentials/custom-idp | Create a custom identity provider credential
[**CreateDatabase**](DefaultAPI.md#CreateDatabase) | **Post** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/ | Creates database with extended information
[**CreateIamRoleMapping**](DefaultAPI.md#CreateIamRoleMapping) | **Post** /v1/organizations/{organizationId}/iam/credentials/aws | Create an IAM Role Mapping
[**CreateLabel**](DefaultAPI.md#CreateLabel) | **Post** /v1/organizations/{organizationId}/labels/ | Add a label to an organization
[**CreateMemberCredential**](DefaultAPI.md#CreateMemberCredential) | **Post** /v1/organizations/{organizationId}/iam/credentials/member | Create a member credential
[**CreateOIDCIntegration**](DefaultAPI.md#CreateOIDCIntegration) | **Post** /v1/organizations/{organizationId}/iam/credentials/oidc | Create an OIDC Integration
[**CreateRole**](DefaultAPI.md#CreateRole) | **Post** /v1/organizations/{organizationId}/roles/ | Create role
[**CreateServiceAccountCredential**](DefaultAPI.md#CreateServiceAccountCredential) | **Post** /v1/organizations/{organizationId}/iam/credentials/service-account | Create a service credential
[**CreateStorageProfile**](DefaultAPI.md#CreateStorageProfile) | **Post** /v1/organizations/{organizationId}/storage-profiles/ | Create a storage profile
[**CreateTableFromFiles**](DefaultAPI.md#CreateTableFromFiles) | **Post** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables | Create and load a table with the data found in a bucket.
[**CreateWarehouse**](DefaultAPI.md#CreateWarehouse) | **Post** /v1/organizations/{organizationId}/warehouses/ | Create a warehouse
[**DeleteDatabase**](DefaultAPI.md#DeleteDatabase) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId} | Drops a database
[**DeleteLabel**](DefaultAPI.md#DeleteLabel) | **Delete** /v1/organizations/{organizationId}/labels/{labelId} | Remove a label from an organization.
[**DeleteMemberCredential**](DefaultAPI.md#DeleteMemberCredential) | **Delete** /v1/organizations/{organizationId}/iam/credentials/{credentialKey} | Delete a credential
[**DeleteRole**](DefaultAPI.md#DeleteRole) | **Delete** /v1/organizations/{organizationId}/roles/{roleName} | Delete Role
[**DeleteStorageProfile**](DefaultAPI.md#DeleteStorageProfile) | **Delete** /v1/organizations/{organizationId}/storage-profiles/{storageProfileId} | Delete a storage profile by ID
[**DeleteWarehouse**](DefaultAPI.md#DeleteWarehouse) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId} | Delete warehouse by id
[**GetCredential**](DefaultAPI.md#GetCredential) | **Get** /v1/organizations/{organizationId}/iam/credentials/{credentialKey} | Fetch information about a credential
[**GetDatabase**](DefaultAPI.md#GetDatabase) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{database} | Gets extended information on a database
[**GetLabel**](DefaultAPI.md#GetLabel) | **Get** /v1/organizations/{organizationId}/labels/{labelId} | Get organization label
[**GetLoadTableDataStatus**](DefaultAPI.md#GetLoadTableDataStatus) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables/{tableId}/loads/{tableLoadId} | One time batch load a table with the data found in a bucket.
[**GetRole**](DefaultAPI.md#GetRole) | **Get** /v1/organizations/{organizationId}/roles/{roleName} | Get role
[**GetStorageProfile**](DefaultAPI.md#GetStorageProfile) | **Get** /v1/organizations/{organizationId}/storage-profiles/{storageProfileId} | Get a storage profile by ID
[**GetTable**](DefaultAPI.md#GetTable) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables/{table} | Get table information
[**GetWarehouse**](DefaultAPI.md#GetWarehouse) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId} | Get a warehouse by id
[**GrantPrivilegesOnDatabase**](DefaultAPI.md#GrantPrivilegesOnDatabase) | **Put** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/grants | Grant privileges on database
[**GrantPrivilegesOnLabel**](DefaultAPI.md#GrantPrivilegesOnLabel) | **Put** /v1/organizations/{organizationId}/labels/{labelId}/grants | Grant privileges on a label.
[**GrantPrivilegesOnTable**](DefaultAPI.md#GrantPrivilegesOnTable) | **Put** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables/{tableId}/grants | Grant privileges on table
[**GrantPrivilegesOnWarehouse**](DefaultAPI.md#GrantPrivilegesOnWarehouse) | **Put** /v1/organizations/{organizationId}/warehouses/{warehouseId}/grants | Grant privileges on a warehouse
[**ListDatabaseRoleGrants**](DefaultAPI.md#ListDatabaseRoleGrants) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/grants | List all grants for database
[**ListDatabaseRoleGrantsForRole**](DefaultAPI.md#ListDatabaseRoleGrantsForRole) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/grants/roles/{roleId} | Get database grants for role id
[**ListLabelRoleGrants**](DefaultAPI.md#ListLabelRoleGrants) | **Get** /v1/organizations/{organizationId}/labels/{labelId}/grants | List all grants for label
[**ListLabelRoleGrantsForRole**](DefaultAPI.md#ListLabelRoleGrantsForRole) | **Get** /v1/organizations/{organizationId}/labels/{labelId}/grants/roles/{roleId} | Get label grants for role id.
[**ListLabels**](DefaultAPI.md#ListLabels) | **Get** /v1/organizations/{organizationId}/labels/ | Get organization labels
[**ListOrganizationMembers**](DefaultAPI.md#ListOrganizationMembers) | **Get** /v1/organizations/{organizationId}/members/ | Get organization members
[**ListTableRoleGrants**](DefaultAPI.md#ListTableRoleGrants) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables/{tableId}/grants | List all grants for table
[**ListTableRoleGrantsForRole**](DefaultAPI.md#ListTableRoleGrantsForRole) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables/{tableId}/grants/roles/{roleId} | Get table grants for role id
[**ListWarehouseRoleGrantsForRole**](DefaultAPI.md#ListWarehouseRoleGrantsForRole) | **Get** /v1/organizations/{organizationId}/warehouses/{warehouseId}/grants/roles/{roleId} | Get warehouse grants for role id
[**ListWarehouses**](DefaultAPI.md#ListWarehouses) | **Get** /v1/organizations/{organizationId}/warehouses/ | List all warehouses
[**LoadTableData**](DefaultAPI.md#LoadTableData) | **Post** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables/{tableId}/loads | One time batch load a table with the data found in a bucket.
[**RemoveChildFromRole**](DefaultAPI.md#RemoveChildFromRole) | **Delete** /v1/organizations/{organizationId}/roles/{roleName}/children | Remove child from role
[**RemoveRoleMembers**](DefaultAPI.md#RemoveRoleMembers) | **Delete** /v1/organizations/{organizationId}/roles/{roleName}/members | Remove members from a role
[**RevokePrivilegesOnDatabase**](DefaultAPI.md#RevokePrivilegesOnDatabase) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/grants | Revoke privileges on database
[**RevokePrivilegesOnLabel**](DefaultAPI.md#RevokePrivilegesOnLabel) | **Delete** /v1/organizations/{organizationId}/labels/{labelId}/grants | Revoke privileges on a label.
[**RevokePrivilegesOnTable**](DefaultAPI.md#RevokePrivilegesOnTable) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId}/databases/{databaseId}/tables/{tableId}/grants | Revoke privileges on table
[**RevokePrivilegesOnWarehouse**](DefaultAPI.md#RevokePrivilegesOnWarehouse) | **Delete** /v1/organizations/{organizationId}/warehouses/{warehouseId}/grants | Revoke privileges on a warehouse
[**UpdateLabel**](DefaultAPI.md#UpdateLabel) | **Patch** /v1/organizations/{organizationId}/labels/{labelId} | Update organization label.
[**UpdateLabelFieldMaskingMode**](DefaultAPI.md#UpdateLabelFieldMaskingMode) | **Put** /v1/organizations/{organizationId}/labels/{labelId}/field/masking | Change label field masking mode.
[**UpdateRoleName**](DefaultAPI.md#UpdateRoleName) | **Put** /v1/organizations/{organizationId}/roles/{roleName} | Update role



## AddChildToRole

> AddChildToRole(ctx, organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()

Add child to role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleName := "roleName_example" // string | 
	updateRoleRequest := *openapiclient.NewUpdateRoleRequest() // UpdateRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AddChildToRole(context.Background(), organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddChildToRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddChildToRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleMembers

> AddRoleMembers(ctx, organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()

Add members to a role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleName := "roleName_example" // string | 
	updateRoleMemberRequest := []openapiclient.UpdateRoleMemberRequest{*openapiclient.NewUpdateRoleMemberRequest()} // []UpdateRoleMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AddRoleMembers(context.Background(), organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddRoleMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleMemberRequest** | [**[]UpdateRoleMemberRequest**](UpdateRoleMemberRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerIdentityProviderCredential

> CreateCredentialResponse CreateCustomerIdentityProviderCredential(ctx, organizationId).CreateCustomIdpCredentialRequest(createCustomIdpCredentialRequest).Execute()

Create a custom identity provider credential

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createCustomIdpCredentialRequest := *openapiclient.NewCreateCustomIdpCredentialRequest() // CreateCustomIdpCredentialRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCustomerIdentityProviderCredential(context.Background(), organizationId).CreateCustomIdpCredentialRequest(createCustomIdpCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCustomerIdentityProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerIdentityProviderCredential`: CreateCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCustomerIdentityProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerIdentityProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomIdpCredentialRequest** | [**CreateCustomIdpCredentialRequest**](CreateCustomIdpCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabase

> CreateDatabaseResponse CreateDatabase(ctx, organizationId, warehouseId).CreateDatabaseRequest(createDatabaseRequest).Execute()

Creates database with extended information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createDatabaseRequest := *openapiclient.NewCreateDatabaseRequest() // CreateDatabaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateDatabase(context.Background(), organizationId, warehouseId).CreateDatabaseRequest(createDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabase`: CreateDatabaseResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDatabaseRequest** | [**CreateDatabaseRequest**](CreateDatabaseRequest.md) |  | 

### Return type

[**CreateDatabaseResponse**](CreateDatabaseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIamRoleMapping

> CreateCredentialResponse CreateIamRoleMapping(ctx, organizationId).CreateIamRoleMappingRequest(createIamRoleMappingRequest).Execute()

Create an IAM Role Mapping

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createIamRoleMappingRequest := *openapiclient.NewCreateIamRoleMappingRequest() // CreateIamRoleMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateIamRoleMapping(context.Background(), organizationId).CreateIamRoleMappingRequest(createIamRoleMappingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateIamRoleMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIamRoleMapping`: CreateCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateIamRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIamRoleMappingRequest** | [**CreateIamRoleMappingRequest**](CreateIamRoleMappingRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLabel

> Label CreateLabel(ctx, organizationId).CreateLabelRequest(createLabelRequest).Execute()

Add a label to an organization

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createLabelRequest := *openapiclient.NewCreateLabelRequest() // CreateLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateLabel(context.Background(), organizationId).CreateLabelRequest(createLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLabel`: Label
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLabelRequest** | [**CreateLabelRequest**](CreateLabelRequest.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMemberCredential

> CreateCredentialResponse CreateMemberCredential(ctx, organizationId).CreateMemberCredentialRequest(createMemberCredentialRequest).Execute()

Create a member credential

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createMemberCredentialRequest := *openapiclient.NewCreateMemberCredentialRequest() // CreateMemberCredentialRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateMemberCredential(context.Background(), organizationId).CreateMemberCredentialRequest(createMemberCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateMemberCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMemberCredential`: CreateCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateMemberCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemberCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMemberCredentialRequest** | [**CreateMemberCredentialRequest**](CreateMemberCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOIDCIntegration

> CreateCredentialResponse CreateOIDCIntegration(ctx, organizationId).CreateOidcCredentialRequest(createOidcCredentialRequest).Execute()

Create an OIDC Integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createOidcCredentialRequest := *openapiclient.NewCreateOidcCredentialRequest() // CreateOidcCredentialRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateOIDCIntegration(context.Background(), organizationId).CreateOidcCredentialRequest(createOidcCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateOIDCIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOIDCIntegration`: CreateCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateOIDCIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOIDCIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOidcCredentialRequest** | [**CreateOidcCredentialRequest**](CreateOidcCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> CreateRoleResponse CreateRole(ctx, organizationId).CreateRoleRequest(createRoleRequest).Execute()

Create role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createRoleRequest := *openapiclient.NewCreateRoleRequest() // CreateRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateRole(context.Background(), organizationId).CreateRoleRequest(createRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: CreateRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRoleRequest** | [**CreateRoleRequest**](CreateRoleRequest.md) |  | 

### Return type

[**CreateRoleResponse**](CreateRoleResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccountCredential

> CreateCredentialResponse CreateServiceAccountCredential(ctx, organizationId).CreateServiceAccountCredentialRequest(createServiceAccountCredentialRequest).Execute()

Create a service credential

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createServiceAccountCredentialRequest := *openapiclient.NewCreateServiceAccountCredentialRequest() // CreateServiceAccountCredentialRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateServiceAccountCredential(context.Background(), organizationId).CreateServiceAccountCredentialRequest(createServiceAccountCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateServiceAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccountCredential`: CreateCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateServiceAccountCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServiceAccountCredentialRequest** | [**CreateServiceAccountCredentialRequest**](CreateServiceAccountCredentialRequest.md) |  | 

### Return type

[**CreateCredentialResponse**](CreateCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorageProfile

> CreateS3StorageProfileResponse CreateStorageProfile(ctx, organizationId).CreateS3StorageProfileRequest(createS3StorageProfileRequest).CreatorRoleId(creatorRoleId).Execute()

Create a storage profile

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createS3StorageProfileRequest := *openapiclient.NewCreateS3StorageProfileRequest() // CreateS3StorageProfileRequest | 
	creatorRoleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateStorageProfile(context.Background(), organizationId).CreateS3StorageProfileRequest(createS3StorageProfileRequest).CreatorRoleId(creatorRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateStorageProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStorageProfile`: CreateS3StorageProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateStorageProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createS3StorageProfileRequest** | [**CreateS3StorageProfileRequest**](CreateS3StorageProfileRequest.md) |  | 
 **creatorRoleId** | **string** |  | 

### Return type

[**CreateS3StorageProfileResponse**](CreateS3StorageProfileResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTableFromFiles

> CreateTableResponse CreateTableFromFiles(ctx, organizationId, warehouseId, databaseId).CreateTableRequest(createTableRequest).Execute()

Create and load a table with the data found in a bucket.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createTableRequest := *openapiclient.NewCreateTableRequest() // CreateTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTableFromFiles(context.Background(), organizationId, warehouseId, databaseId).CreateTableRequest(createTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTableFromFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTableFromFiles`: CreateTableResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTableFromFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableFromFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createTableRequest** | [**CreateTableRequest**](CreateTableRequest.md) |  | 

### Return type

[**CreateTableResponse**](CreateTableResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWarehouse

> CreateWarehouseResponse CreateWarehouse(ctx, organizationId).CreateWarehouseRequest(createWarehouseRequest).CreatorRoleId(creatorRoleId).Execute()

Create a warehouse

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "organizationId_example" // string | 
	createWarehouseRequest := *openapiclient.NewCreateWarehouseRequest() // CreateWarehouseRequest | 
	creatorRoleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateWarehouse(context.Background(), organizationId).CreateWarehouseRequest(createWarehouseRequest).CreatorRoleId(creatorRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarehouse`: CreateWarehouseResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWarehouseRequest** | [**CreateWarehouseRequest**](CreateWarehouseRequest.md) |  | 
 **creatorRoleId** | **string** |  | 

### Return type

[**CreateWarehouseResponse**](CreateWarehouseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> DeleteDatabase(ctx, organizationId, warehouseId, databaseId).Execute()

Drops a database

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteDatabase(context.Background(), organizationId, warehouseId, databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLabel

> DeleteLabel(ctx, organizationId, labelId).Execute()

Remove a label from an organization.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteLabel(context.Background(), organizationId, labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemberCredential

> DeleteMemberCredential(ctx, organizationId, credentialKey).Execute()

Delete a credential

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	credentialKey := "credentialKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteMemberCredential(context.Background(), organizationId, credentialKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteMemberCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**credentialKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, organizationId, roleName).Force(force).Execute()

Delete Role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleName := "roleName_example" // string | 
	force := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteRole(context.Background(), organizationId, roleName).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorageProfile

> DeleteStorageProfile(ctx, organizationId, storageProfileId).Execute()

Delete a storage profile by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	storageProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteStorageProfile(context.Background(), organizationId, storageProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteStorageProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**storageProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWarehouse

> DeleteWarehouse(ctx, organizationId, warehouseId).Execute()

Delete warehouse by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteWarehouse(context.Background(), organizationId, warehouseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredential

> GetCredentialResponse GetCredential(ctx, organizationId, credentialKey).Execute()

Fetch information about a credential

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	credentialKey := "credentialKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCredential(context.Background(), organizationId, credentialKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: GetCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**credentialKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCredentialResponse**](GetCredentialResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabase

> GetDatabaseResponse GetDatabase(ctx, organizationId, warehouseId, database).Type_(type_).Execute()

Gets extended information on a database

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	database := "database_example" // string | 
	type_ := "type__example" // string |  (optional) (default to "name")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDatabase(context.Background(), organizationId, warehouseId, database).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabase`: GetDatabaseResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** |  | [default to &quot;name&quot;]

### Return type

[**GetDatabaseResponse**](GetDatabaseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabel

> GetLabelResponse GetLabel(ctx, organizationId, labelId).Execute()

Get organization label

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetLabel(context.Background(), organizationId, labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabel`: GetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLabelResponse**](GetLabelResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadTableDataStatus

> LoadTableStatus GetLoadTableDataStatus(ctx, organizationId, warehouseId, databaseId, tableId, tableLoadId).Execute()

One time batch load a table with the data found in a bucket.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableLoadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetLoadTableDataStatus(context.Background(), organizationId, warehouseId, databaseId, tableId, tableLoadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLoadTableDataStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadTableDataStatus`: LoadTableStatus
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLoadTableDataStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**tableId** | **string** |  | 
**tableLoadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadTableDataStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**LoadTableStatus**](LoadTableStatus.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> GetRoleResponse GetRole(ctx, organizationId, roleName).Execute()

Get role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRole(context.Background(), organizationId, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: GetRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRoleResponse**](GetRoleResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageProfile

> GetStorageProfileResponse GetStorageProfile(ctx, organizationId, storageProfileId).Execute()

Get a storage profile by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	storageProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetStorageProfile(context.Background(), organizationId, storageProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetStorageProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageProfile`: GetStorageProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetStorageProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**storageProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStorageProfileResponse**](GetStorageProfileResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTable

> GetTableResponse GetTable(ctx, organizationId, warehouseId, databaseId, table).Execute()

Get table information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	table := "table_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTable(context.Background(), organizationId, warehouseId, databaseId, table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTable`: GetTableResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**table** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetTableResponse**](GetTableResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarehouse

> GetWarehouseResponse GetWarehouse(ctx, organizationId, warehouseId).Execute()

Get a warehouse by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetWarehouse(context.Background(), organizationId, warehouseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarehouse`: GetWarehouseResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetWarehouseResponse**](GetWarehouseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPrivilegesOnDatabase

> GrantPrivilegesOnDatabase(ctx, organizationId, warehouseId, databaseId).RoleDatabaseGrantRequest(roleDatabaseGrantRequest).Execute()

Grant privileges on database

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleDatabaseGrantRequest := []openapiclient.RoleDatabaseGrantRequest{*openapiclient.NewRoleDatabaseGrantRequest()} // []RoleDatabaseGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GrantPrivilegesOnDatabase(context.Background(), organizationId, warehouseId, databaseId).RoleDatabaseGrantRequest(roleDatabaseGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GrantPrivilegesOnDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPrivilegesOnDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleDatabaseGrantRequest** | [**[]RoleDatabaseGrantRequest**](RoleDatabaseGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPrivilegesOnLabel

> GrantPrivilegesOnLabel(ctx, organizationId, labelId).RoleLabelGrantRequest(roleLabelGrantRequest).Execute()

Grant privileges on a label.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleLabelGrantRequest := []openapiclient.RoleLabelGrantRequest{*openapiclient.NewRoleLabelGrantRequest()} // []RoleLabelGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GrantPrivilegesOnLabel(context.Background(), organizationId, labelId).RoleLabelGrantRequest(roleLabelGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GrantPrivilegesOnLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPrivilegesOnLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleLabelGrantRequest** | [**[]RoleLabelGrantRequest**](RoleLabelGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPrivilegesOnTable

> GrantPrivilegesOnTable(ctx, organizationId, warehouseId, databaseId, tableId).RoleTableGrantRequest(roleTableGrantRequest).Execute()

Grant privileges on table

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleTableGrantRequest := []openapiclient.RoleTableGrantRequest{*openapiclient.NewRoleTableGrantRequest()} // []RoleTableGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GrantPrivilegesOnTable(context.Background(), organizationId, warehouseId, databaseId, tableId).RoleTableGrantRequest(roleTableGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GrantPrivilegesOnTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**tableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPrivilegesOnTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **roleTableGrantRequest** | [**[]RoleTableGrantRequest**](RoleTableGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPrivilegesOnWarehouse

> GrantPrivilegesOnWarehouse(ctx, organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()

Grant privileges on a warehouse

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleWarehouseGrantRequest := []openapiclient.RoleWarehouseGrantRequest{*openapiclient.NewRoleWarehouseGrantRequest()} // []RoleWarehouseGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GrantPrivilegesOnWarehouse(context.Background(), organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GrantPrivilegesOnWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPrivilegesOnWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleWarehouseGrantRequest** | [**[]RoleWarehouseGrantRequest**](RoleWarehouseGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseRoleGrants

> ListDatabaseRoleGrantsResponse ListDatabaseRoleGrants(ctx, organizationId, warehouseId, databaseId).Execute()

List all grants for database

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListDatabaseRoleGrants(context.Background(), organizationId, warehouseId, databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListDatabaseRoleGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabaseRoleGrants`: ListDatabaseRoleGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListDatabaseRoleGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseRoleGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListDatabaseRoleGrantsResponse**](ListDatabaseRoleGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseRoleGrantsForRole

> GetRoleDatabaseGrantsResponse ListDatabaseRoleGrantsForRole(ctx, organizationId, warehouseId, databaseId, roleId).Execute()

Get database grants for role id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListDatabaseRoleGrantsForRole(context.Background(), organizationId, warehouseId, databaseId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListDatabaseRoleGrantsForRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabaseRoleGrantsForRole`: GetRoleDatabaseGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListDatabaseRoleGrantsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseRoleGrantsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetRoleDatabaseGrantsResponse**](GetRoleDatabaseGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabelRoleGrants

> ListRoleLabelGrantsResponse ListLabelRoleGrants(ctx, organizationId, labelId).Execute()

List all grants for label

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListLabelRoleGrants(context.Background(), organizationId, labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLabelRoleGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLabelRoleGrants`: ListRoleLabelGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListLabelRoleGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelRoleGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListRoleLabelGrantsResponse**](ListRoleLabelGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabelRoleGrantsForRole

> ListRoleLabelGrantsResponse ListLabelRoleGrantsForRole(ctx, organizationId, labelId, roleId).Execute()

Get label grants for role id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListLabelRoleGrantsForRole(context.Background(), organizationId, labelId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLabelRoleGrantsForRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLabelRoleGrantsForRole`: ListRoleLabelGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListLabelRoleGrantsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelRoleGrantsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListRoleLabelGrantsResponse**](ListRoleLabelGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabels

> ListLabelsResponse ListLabels(ctx, organizationId).Execute()

Get organization labels

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListLabels(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLabels`: ListLabelsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLabelsResponse**](ListLabelsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationMembers

> ListMembersResponse ListOrganizationMembers(ctx, organizationId).Execute()

Get organization members

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListOrganizationMembers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListOrganizationMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationMembers`: ListMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListOrganizationMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMembersResponse**](ListMembersResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTableRoleGrants

> ListTableRoleGrantsResponse ListTableRoleGrants(ctx, organizationId, warehouseId, databaseId, tableId).Execute()

List all grants for table

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTableRoleGrants(context.Background(), organizationId, warehouseId, databaseId, tableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTableRoleGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTableRoleGrants`: ListTableRoleGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTableRoleGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**tableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTableRoleGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ListTableRoleGrantsResponse**](ListTableRoleGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTableRoleGrantsForRole

> ListTableRoleGrantsResponse ListTableRoleGrantsForRole(ctx, organizationId, warehouseId, databaseId, tableId, roleId).Execute()

Get table grants for role id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTableRoleGrantsForRole(context.Background(), organizationId, warehouseId, databaseId, tableId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTableRoleGrantsForRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTableRoleGrantsForRole`: ListTableRoleGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTableRoleGrantsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**tableId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTableRoleGrantsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**ListTableRoleGrantsResponse**](ListTableRoleGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWarehouseRoleGrantsForRole

> GetRoleWarehouseGrantsResponse ListWarehouseRoleGrantsForRole(ctx, organizationId, warehouseId, roleId).Execute()

Get warehouse grants for role id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWarehouseRoleGrantsForRole(context.Background(), organizationId, warehouseId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWarehouseRoleGrantsForRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWarehouseRoleGrantsForRole`: GetRoleWarehouseGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWarehouseRoleGrantsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWarehouseRoleGrantsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetRoleWarehouseGrantsResponse**](GetRoleWarehouseGrantsResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWarehouses

> ListWarehouseResponse ListWarehouses(ctx, organizationId).Execute()

List all warehouses

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWarehouses(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWarehouses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWarehouses`: ListWarehouseResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWarehouses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListWarehouseResponse**](ListWarehouseResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadTableData

> LoadTableResponse LoadTableData(ctx, organizationId, warehouseId, databaseId, tableId).LoadTableRequest(loadTableRequest).Execute()

One time batch load a table with the data found in a bucket.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	loadTableRequest := *openapiclient.NewLoadTableRequest() // LoadTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LoadTableData(context.Background(), organizationId, warehouseId, databaseId, tableId).LoadTableRequest(loadTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LoadTableData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTableData`: LoadTableResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LoadTableData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**tableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTableDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **loadTableRequest** | [**LoadTableRequest**](LoadTableRequest.md) |  | 

### Return type

[**LoadTableResponse**](LoadTableResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveChildFromRole

> RemoveChildFromRole(ctx, organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()

Remove child from role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleName := "roleName_example" // string | 
	updateRoleRequest := *openapiclient.NewUpdateRoleRequest() // UpdateRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RemoveChildFromRole(context.Background(), organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RemoveChildFromRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveChildFromRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleMembers

> RemoveRoleMembers(ctx, organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()

Remove members from a role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleName := "roleName_example" // string | 
	updateRoleMemberRequest := []openapiclient.UpdateRoleMemberRequest{*openapiclient.NewUpdateRoleMemberRequest()} // []UpdateRoleMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RemoveRoleMembers(context.Background(), organizationId, roleName).UpdateRoleMemberRequest(updateRoleMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RemoveRoleMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleMemberRequest** | [**[]UpdateRoleMemberRequest**](UpdateRoleMemberRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePrivilegesOnDatabase

> RevokePrivilegesOnDatabase(ctx, organizationId, warehouseId, databaseId).RoleDatabaseGrantRequest(roleDatabaseGrantRequest).Execute()

Revoke privileges on database

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleDatabaseGrantRequest := []openapiclient.RoleDatabaseGrantRequest{*openapiclient.NewRoleDatabaseGrantRequest()} // []RoleDatabaseGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RevokePrivilegesOnDatabase(context.Background(), organizationId, warehouseId, databaseId).RoleDatabaseGrantRequest(roleDatabaseGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokePrivilegesOnDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrivilegesOnDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleDatabaseGrantRequest** | [**[]RoleDatabaseGrantRequest**](RoleDatabaseGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePrivilegesOnLabel

> RevokePrivilegesOnLabel(ctx, organizationId, labelId).RoleLabelGrantRequest(roleLabelGrantRequest).Execute()

Revoke privileges on a label.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleLabelGrantRequest := []openapiclient.RoleLabelGrantRequest{*openapiclient.NewRoleLabelGrantRequest()} // []RoleLabelGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RevokePrivilegesOnLabel(context.Background(), organizationId, labelId).RoleLabelGrantRequest(roleLabelGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokePrivilegesOnLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrivilegesOnLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleLabelGrantRequest** | [**[]RoleLabelGrantRequest**](RoleLabelGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePrivilegesOnTable

> RevokePrivilegesOnTable(ctx, organizationId, warehouseId, databaseId, tableId).RoleTableGrantRequest(roleTableGrantRequest).Execute()

Revoke privileges on table

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleTableGrantRequest := []openapiclient.RoleTableGrantRequest{*openapiclient.NewRoleTableGrantRequest()} // []RoleTableGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RevokePrivilegesOnTable(context.Background(), organizationId, warehouseId, databaseId, tableId).RoleTableGrantRequest(roleTableGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokePrivilegesOnTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 
**databaseId** | **string** |  | 
**tableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrivilegesOnTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **roleTableGrantRequest** | [**[]RoleTableGrantRequest**](RoleTableGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePrivilegesOnWarehouse

> RevokePrivilegesOnWarehouse(ctx, organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()

Revoke privileges on a warehouse

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleWarehouseGrantRequest := []openapiclient.RoleWarehouseGrantRequest{*openapiclient.NewRoleWarehouseGrantRequest()} // []RoleWarehouseGrantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RevokePrivilegesOnWarehouse(context.Background(), organizationId, warehouseId).RoleWarehouseGrantRequest(roleWarehouseGrantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RevokePrivilegesOnWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrivilegesOnWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleWarehouseGrantRequest** | [**[]RoleWarehouseGrantRequest**](RoleWarehouseGrantRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabel

> Label UpdateLabel(ctx, organizationId, labelId).UpdateLabelRequest(updateLabelRequest).Execute()

Update organization label.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateLabelRequest := *openapiclient.NewUpdateLabelRequest() // UpdateLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateLabel(context.Background(), organizationId, labelId).UpdateLabelRequest(updateLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLabel`: Label
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateLabelRequest** | [**UpdateLabelRequest**](UpdateLabelRequest.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabelFieldMaskingMode

> Label UpdateLabelFieldMaskingMode(ctx, organizationId, labelId).UpdateLabelFieldMaskingRequest(updateLabelFieldMaskingRequest).Execute()

Change label field masking mode.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateLabelFieldMaskingRequest := *openapiclient.NewUpdateLabelFieldMaskingRequest() // UpdateLabelFieldMaskingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateLabelFieldMaskingMode(context.Background(), organizationId, labelId).UpdateLabelFieldMaskingRequest(updateLabelFieldMaskingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateLabelFieldMaskingMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLabelFieldMaskingMode`: Label
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateLabelFieldMaskingMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelFieldMaskingModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateLabelFieldMaskingRequest** | [**UpdateLabelFieldMaskingRequest**](UpdateLabelFieldMaskingRequest.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleName

> UpdateRoleNameResponse UpdateRoleName(ctx, organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()

Update role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleName := "roleName_example" // string | 
	updateRoleRequest := *openapiclient.NewUpdateRoleRequest() // UpdateRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateRoleName(context.Background(), organizationId, roleName).UpdateRoleRequest(updateRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateRoleName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleName`: UpdateRoleNameResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateRoleName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

[**UpdateRoleNameResponse**](UpdateRoleNameResponse.md)

### Authorization

[TabularJWTAuth](../README.md#TabularJWTAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

