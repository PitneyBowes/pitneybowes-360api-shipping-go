/*
Shipping APIs

Testing MultipieceAPIService

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

func Test_shipping_MultipieceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MultipieceAPIService MultipieceRates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MultipieceAPI.MultipieceRates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultipieceAPIService MultipieceShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MultipieceAPI.MultipieceShipment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultipieceAPIService MultipieceShipmentCancel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipmentId string

		resp, httpRes, err := apiClient.MultipieceAPI.MultipieceShipmentCancel(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultipieceAPIService MultipieceShipmentReprint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shipmentId string

		resp, httpRes, err := apiClient.MultipieceAPI.MultipieceShipmentReprint(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
