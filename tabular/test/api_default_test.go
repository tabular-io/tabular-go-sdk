/*
Tabular API

Testing DefaultAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tabular

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/tabular-io/tabular-sdk-go/tabular"
	"testing"
)

func Test_tabular_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService AddChildToRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultAPI.AddChildToRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService AddRoleMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultAPI.AddRoleMembers(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateCustomerIdentityProviderCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateCustomerIdentityProviderCredential(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateDatabase(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateIamRoleMapping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateIamRoleMapping(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateLabel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateLabel(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateMemberCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateMemberCredential(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateOIDCIntegration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateOIDCIntegration(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateRole(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateServiceAccountCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateServiceAccountCredential(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateStorageProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateStorageProfile(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateTableFromFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateTableFromFiles(context.Background(), organizationId, warehouseId, databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.CreateWarehouse(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string

		httpRes, err := apiClient.DefaultAPI.DeleteDatabase(context.Background(), organizationId, warehouseId, databaseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteLabel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string

		httpRes, err := apiClient.DefaultAPI.DeleteLabel(context.Background(), organizationId, labelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteMemberCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var credentialKey string

		httpRes, err := apiClient.DefaultAPI.DeleteMemberCredential(context.Background(), organizationId, credentialKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultAPI.DeleteRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteServiceAccountCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var credentialKey string

		httpRes, err := apiClient.DefaultAPI.DeleteServiceAccountCredential(context.Background(), organizationId, credentialKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteStorageProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var storageProfileId string

		httpRes, err := apiClient.DefaultAPI.DeleteStorageProfile(context.Background(), organizationId, storageProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		httpRes, err := apiClient.DefaultAPI.DeleteWarehouse(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var credentialKey string

		resp, httpRes, err := apiClient.DefaultAPI.GetCredential(context.Background(), organizationId, credentialKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var database string

		resp, httpRes, err := apiClient.DefaultAPI.GetDatabase(context.Background(), organizationId, warehouseId, database).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetLabel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string

		resp, httpRes, err := apiClient.DefaultAPI.GetLabel(context.Background(), organizationId, labelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetLoadTableDataStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var tableId string
		var tableLoadId string

		resp, httpRes, err := apiClient.DefaultAPI.GetLoadTableDataStatus(context.Background(), organizationId, warehouseId, databaseId, tableId, tableLoadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		resp, httpRes, err := apiClient.DefaultAPI.GetRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetStorageProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var storageProfileId string

		resp, httpRes, err := apiClient.DefaultAPI.GetStorageProfile(context.Background(), organizationId, storageProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetTable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var table string

		resp, httpRes, err := apiClient.DefaultAPI.GetTable(context.Background(), organizationId, warehouseId, databaseId, table).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		resp, httpRes, err := apiClient.DefaultAPI.GetWarehouse(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GrantPrivilegesOnDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string

		httpRes, err := apiClient.DefaultAPI.GrantPrivilegesOnDatabase(context.Background(), organizationId, warehouseId, databaseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GrantPrivilegesOnLabel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string

		httpRes, err := apiClient.DefaultAPI.GrantPrivilegesOnLabel(context.Background(), organizationId, labelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GrantPrivilegesOnTable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var tableId string

		httpRes, err := apiClient.DefaultAPI.GrantPrivilegesOnTable(context.Background(), organizationId, warehouseId, databaseId, tableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GrantPrivilegesOnWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		httpRes, err := apiClient.DefaultAPI.GrantPrivilegesOnWarehouse(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListDatabaseRoleGrants", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string

		resp, httpRes, err := apiClient.DefaultAPI.ListDatabaseRoleGrants(context.Background(), organizationId, warehouseId, databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListDatabaseRoleGrantsForRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var roleId string

		resp, httpRes, err := apiClient.DefaultAPI.ListDatabaseRoleGrantsForRole(context.Background(), organizationId, warehouseId, databaseId, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListLabelRoleGrants", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string

		resp, httpRes, err := apiClient.DefaultAPI.ListLabelRoleGrants(context.Background(), organizationId, labelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListLabelRoleGrantsForRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string
		var roleId string

		resp, httpRes, err := apiClient.DefaultAPI.ListLabelRoleGrantsForRole(context.Background(), organizationId, labelId, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListLabels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.ListLabels(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListOrganizationMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.ListOrganizationMembers(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListTableRoleGrants", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var tableId string

		resp, httpRes, err := apiClient.DefaultAPI.ListTableRoleGrants(context.Background(), organizationId, warehouseId, databaseId, tableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListTableRoleGrantsForRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var tableId string
		var roleId string

		resp, httpRes, err := apiClient.DefaultAPI.ListTableRoleGrantsForRole(context.Background(), organizationId, warehouseId, databaseId, tableId, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListWarehouseRoleGrantsForRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var roleId string

		resp, httpRes, err := apiClient.DefaultAPI.ListWarehouseRoleGrantsForRole(context.Background(), organizationId, warehouseId, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListWarehouses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultAPI.ListWarehouses(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService LoadTableData", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var tableId string

		resp, httpRes, err := apiClient.DefaultAPI.LoadTableData(context.Background(), organizationId, warehouseId, databaseId, tableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RemoveChildFromRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultAPI.RemoveChildFromRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RemoveRoleMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultAPI.RemoveRoleMembers(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RevokePrivilegesOnDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string

		httpRes, err := apiClient.DefaultAPI.RevokePrivilegesOnDatabase(context.Background(), organizationId, warehouseId, databaseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RevokePrivilegesOnLabel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string

		httpRes, err := apiClient.DefaultAPI.RevokePrivilegesOnLabel(context.Background(), organizationId, labelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RevokePrivilegesOnTable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var databaseId string
		var tableId string

		httpRes, err := apiClient.DefaultAPI.RevokePrivilegesOnTable(context.Background(), organizationId, warehouseId, databaseId, tableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RevokePrivilegesOnWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		httpRes, err := apiClient.DefaultAPI.RevokePrivilegesOnWarehouse(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UpdateLabel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string

		resp, httpRes, err := apiClient.DefaultAPI.UpdateLabel(context.Background(), organizationId, labelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UpdateLabelFieldMaskingMode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var labelId string

		resp, httpRes, err := apiClient.DefaultAPI.UpdateLabelFieldMaskingMode(context.Background(), organizationId, labelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UpdateRoleName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		resp, httpRes, err := apiClient.DefaultAPI.UpdateRoleName(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
