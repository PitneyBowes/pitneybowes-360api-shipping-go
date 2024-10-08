/*
Shipping APIs

Testing DefaultsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package shipping

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/PitneyBowes/pitneybowes-360api-shipping-go"
)

func Test_shipping_DefaultsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultsAPIService CreateDefaults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultsAPI.CreateDefaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultsAPIService DeleteDefaultsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var defaultID string

		httpRes, err := apiClient.DefaultsAPI.DeleteDefaultsById(context.Background(), defaultID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultsAPIService GetAllDefaults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultsAPI.GetAllDefaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultsAPIService GetDefaultsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var defaultID string

		resp, httpRes, err := apiClient.DefaultsAPI.GetDefaultsById(context.Background(), defaultID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultsAPIService PutDefaultsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var defaultID string

		httpRes, err := apiClient.DefaultsAPI.PutDefaultsById(context.Background(), defaultID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
