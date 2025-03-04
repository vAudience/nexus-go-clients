/*
vAudience AIgentChat API

Testing ChannelsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package aigentchat

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/vaudience/nexus-go-clients/aigentchat"
)

func Test_aigentchat_ChannelsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChannelsAPIService CreateChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ChannelsAPI.CreateChannel(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService CreateChannelFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ChannelsAPI.CreateChannelFile(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService DeleteChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var id string

		resp, httpRes, err := apiClient.ChannelsAPI.DeleteChannel(context.Background(), orgId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService DeleteChannelsByOwnerId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ChannelsAPI.DeleteChannelsByOwnerId(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService GetChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var id string

		resp, httpRes, err := apiClient.ChannelsAPI.GetChannel(context.Background(), orgId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService GetChannelFileSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ChannelsAPI.GetChannelFileSettings(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService ListChannelsByOrgId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ChannelsAPI.ListChannelsByOrgId(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService ListChannelsByOwnerId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ChannelsAPI.ListChannelsByOwnerId(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsAPIService UpdateChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var id string

		resp, httpRes, err := apiClient.ChannelsAPI.UpdateChannel(context.Background(), orgId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
