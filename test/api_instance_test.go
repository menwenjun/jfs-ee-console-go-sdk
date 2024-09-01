/*
JuiceFS console API

Testing InstanceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/menwenjun/jfs-ee-console-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_InstanceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InstanceAPIService MetadatasMetadataIDInstancesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var metadataID int32

		resp, httpRes, err := apiClient.InstanceAPI.MetadatasMetadataIDInstancesGet(context.Background(), metadataID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceAPIService MetadatasMetadataIDInstancesInstanceIDGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var metadataID int32
		var instanceID int32

		resp, httpRes, err := apiClient.InstanceAPI.MetadatasMetadataIDInstancesInstanceIDGet(context.Background(), metadataID, instanceID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceAPIService MetadatasMetadataIDInstancesInstanceIDPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var metadataID int32
		var instanceID int32

		resp, httpRes, err := apiClient.InstanceAPI.MetadatasMetadataIDInstancesInstanceIDPatch(context.Background(), metadataID, instanceID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
