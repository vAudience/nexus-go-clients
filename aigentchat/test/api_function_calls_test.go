/*
vAudience AIgentChat API

Testing FunctionCallsAPIService

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

func Test_aigentchat_FunctionCallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FunctionCallsAPIService ExecuteFunctionCall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var name string

		resp, httpRes, err := apiClient.FunctionCallsAPI.ExecuteFunctionCall(context.Background(), orgId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionCallsAPIService GetFunctionCall", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var name string

		resp, httpRes, err := apiClient.FunctionCallsAPI.GetFunctionCall(context.Background(), orgId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FunctionCallsAPIService GetFunctionCalls", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.FunctionCallsAPI.GetFunctionCalls(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
