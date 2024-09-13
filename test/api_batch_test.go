/*
Shipping APIs

Testing BatchAPIService

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

func Test_shipping_BatchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BatchAPIService BulkImportAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchAPI.BulkImportAPI(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService BulkImportAPIERR", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchAPI.BulkImportAPIERR(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService CreateBulkShipmentsAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchAPI.CreateBulkShipmentsAPI(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService CreateBulkShipmentsAPIERR", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BatchAPI.CreateBulkShipmentsAPIERR(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService GetBatchStatusAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var batchId string

		resp, httpRes, err := apiClient.BatchAPI.GetBatchStatusAPI(context.Background(), batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService GetShipmentDetailsForBatchAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var batchId string

		resp, httpRes, err := apiClient.BatchAPI.GetShipmentDetailsForBatchAPI(context.Background(), batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService ProcessBatchAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var batchId string

		resp, httpRes, err := apiClient.BatchAPI.ProcessBatchAPI(context.Background(), batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService VoidShippingLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var batchId string

		resp, httpRes, err := apiClient.BatchAPI.VoidShippingLabel(context.Background(), batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
