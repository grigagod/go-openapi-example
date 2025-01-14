/*
Train Travel API

Testing TripsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package oas

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/grigagod/go-openapi-example/openapitools/gen/client"
)

func Test_oas_TripsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TripsAPIService GetTrips", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TripsAPI.GetTrips(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
