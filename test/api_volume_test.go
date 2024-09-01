/*
JuiceFS console API

Testing VolumeAPIService

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

func Test_openapi_VolumeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VolumeAPIService VolumesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VolumeAPI.VolumesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VolumeAPI.VolumesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDClientsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDClientsGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDDelete(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDIsReadyGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDIsReadyGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDKerberosGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDKerberosGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDKerberosPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDKerberosPatch(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDListDirGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDListDirGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDMetricsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDMetricsGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDMirrorsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDMirrorsGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDMirrorsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDMirrorsPost(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDRangerGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDRangerGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDRangerPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDRangerPatch(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDSettingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDSettingGet(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VolumeAPIService VolumesVolumeIDSettingPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volumeID int32

		resp, httpRes, err := apiClient.VolumeAPI.VolumesVolumeIDSettingPost(context.Background(), volumeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
