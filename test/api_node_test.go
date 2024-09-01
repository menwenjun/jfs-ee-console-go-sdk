/*
JuiceFS console API

Testing NodeAPIService

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

func Test_openapi_NodeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NodeAPIService NodesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NodeAPI.NodesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeAPIService NodesNodeIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeID int32

		resp, httpRes, err := apiClient.NodeAPI.NodesNodeIDGet(context.Background(), nodeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeAPIService NodesNodeIDPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeID int32

		resp, httpRes, err := apiClient.NodeAPI.NodesNodeIDPatch(context.Background(), nodeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeAPIService NodesNodeIDPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeID int32

		resp, httpRes, err := apiClient.NodeAPI.NodesNodeIDPut(context.Background(), nodeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
