/*
OneLogin API

Testing EventsApiService

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

func Test_onelogin_EventsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EventsApiService GetEventById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var eventId int32

		resp, httpRes, err := apiClient.EventsApi.GetEventById(context.Background(), eventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GetEventTypes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EventsApi.GetEventTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GetEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EventsApi.GetEvents(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
