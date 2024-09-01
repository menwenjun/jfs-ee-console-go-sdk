/*
JuiceFS console API

Testing VolumeACLAPIService

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

func Test_openapi_VolumeACLAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VolumeACLAPIService VolumesVolumeIDExportsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeACLAPIService VolumesVolumeIDExportsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsPost(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeACLAPIService VolumesVolumeIDExportsVolumeExportIDDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32
		var volumeExportID int32

		httpRes, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDDelete(context.Background(), volumeID, volumeExportID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeACLAPIService VolumesVolumeIDExportsVolumeExportIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32
		var volumeExportID int32

		resp, httpRes, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDGet(context.Background(), volumeID, volumeExportID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeACLAPIService VolumesVolumeIDExportsVolumeExportIDPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32
		var volumeExportID int32

		resp, httpRes, err := apiClient.VolumeACLAPI.VolumesVolumeIDExportsVolumeExportIDPut(context.Background(), volumeID, volumeExportID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}