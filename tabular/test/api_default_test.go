/*
OpenAPI definition

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tabular

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/tabular-io/tabular-sdk-go"
	"testing"
)

func Test_tabular_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService AddChildToRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultApi.AddChildToRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService AddRoleMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultApi.AddRoleMembers(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateCustomerIdentityProviderCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateCustomerIdentityProviderCredential(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		resp, httpRes, err := apiClient.DefaultApi.CreateDatabase(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateIamRoleMapping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateIamRoleMapping(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateMemberCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateMemberCredential(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateOIDCIntegration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateOIDCIntegration(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateRole(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateServiceAccountCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateServiceAccountCredential(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateStorageProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateStorageProfile(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.CreateWarehouse(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var database string

		httpRes, err := apiClient.DefaultApi.DeleteDatabase(context.Background(), organizationId, warehouseId, database).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteMemberCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var credentialKey string

		httpRes, err := apiClient.DefaultApi.DeleteMemberCredential(context.Background(), organizationId, credentialKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultApi.DeleteRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteStorageProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var storageProfileId string

		httpRes, err := apiClient.DefaultApi.DeleteStorageProfile(context.Background(), organizationId, storageProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		httpRes, err := apiClient.DefaultApi.DeleteWarehouse(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var credentialKey string

		resp, httpRes, err := apiClient.DefaultApi.GetCredential(context.Background(), organizationId, credentialKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var database string

		resp, httpRes, err := apiClient.DefaultApi.GetDatabase(context.Background(), organizationId, warehouseId, database).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetDatabaseRoleGrants", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var database string

		resp, httpRes, err := apiClient.DefaultApi.GetDatabaseRoleGrants(context.Background(), organizationId, warehouseId, database).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetOrganizationMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.GetOrganizationMembers(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		resp, httpRes, err := apiClient.DefaultApi.GetRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetStorageProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var storageProfileId string

		resp, httpRes, err := apiClient.DefaultApi.GetStorageProfile(context.Background(), organizationId, storageProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetWarehouse", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string

		resp, httpRes, err := apiClient.DefaultApi.GetWarehouse(context.Background(), organizationId, warehouseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GrantPrivilegesOnDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var database string

		httpRes, err := apiClient.DefaultApi.GrantPrivilegesOnDatabase(context.Background(), organizationId, warehouseId, database).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService ListWarehouses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DefaultApi.ListWarehouses(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService RemoveChildFromRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultApi.RemoveChildFromRole(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService RemoveRoleMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		httpRes, err := apiClient.DefaultApi.RemoveRoleMembers(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService RevokePrivilegesOnDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var warehouseId string
		var database string

		httpRes, err := apiClient.DefaultApi.RevokePrivilegesOnDatabase(context.Background(), organizationId, warehouseId, database).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService UpdateRoleName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var roleName string

		resp, httpRes, err := apiClient.DefaultApi.UpdateRoleName(context.Background(), organizationId, roleName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
