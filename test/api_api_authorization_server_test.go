/*
OneLogin API

Testing APIAuthorizationServerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package onelogin

import (
	"context"
	"testing"

	openapiclient "github.com/onelogin/onelogin-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_onelogin_APIAuthorizationServerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test APIAuthorizationServerApiService CreateAuthServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.APIAuthorizationServerApi.CreateAuthServer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIAuthorizationServerApiService DeleteAuthServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiAuthId string

		httpRes, err := apiClient.APIAuthorizationServerApi.DeleteAuthServer(context.Background(), apiAuthId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIAuthorizationServerApiService GetAuthServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiAuthId string

		resp, httpRes, err := apiClient.APIAuthorizationServerApi.GetAuthServer(context.Background(), apiAuthId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIAuthorizationServerApiService ListAuthServers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.APIAuthorizationServerApi.ListAuthServers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIAuthorizationServerApiService UpdateAuthServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiAuthId string

		resp, httpRes, err := apiClient.APIAuthorizationServerApi.UpdateAuthServer(context.Background(), apiAuthId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
