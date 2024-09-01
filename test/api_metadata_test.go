/*
JuiceFS console API

Testing MetadataAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_MetadataAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetadataAPIService MetadataDnsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetadataAPI.MetadataDnsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataAPIService MetadatasGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetadataAPI.MetadatasGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataAPIService MetadatasMetadataIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metadataID int32

		resp, httpRes, err := apiClient.MetadataAPI.MetadatasMetadataIDGet(context.Background(), metadataID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataAPIService MetadatasMetadataIDLicenseGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metadataID int32

		resp, httpRes, err := apiClient.MetadataAPI.MetadatasMetadataIDLicenseGet(context.Background(), metadataID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetadataAPIService MetadatasMetadataIDLicensePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metadataID int32

		resp, httpRes, err := apiClient.MetadataAPI.MetadatasMetadataIDLicensePut(context.Background(), metadataID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}