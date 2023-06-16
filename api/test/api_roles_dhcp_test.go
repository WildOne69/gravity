/*
gravity

Testing RolesDhcpApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "beryju.io/gravity/api"
)

func Test_api_RolesDhcpApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RolesDhcpApiService DhcpDeleteLeases", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RolesDhcpApi.DhcpDeleteLeases(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpDeleteScopes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RolesDhcpApi.DhcpDeleteScopes(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpGetLeases", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RolesDhcpApi.DhcpGetLeases(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpGetRoleConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RolesDhcpApi.DhcpGetRoleConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpGetScopes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RolesDhcpApi.DhcpGetScopes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpPutLeases", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RolesDhcpApi.DhcpPutLeases(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpPutRoleConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RolesDhcpApi.DhcpPutRoleConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpPutScopes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RolesDhcpApi.DhcpPutScopes(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesDhcpApiService DhcpWolLeases", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RolesDhcpApi.DhcpWolLeases(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
