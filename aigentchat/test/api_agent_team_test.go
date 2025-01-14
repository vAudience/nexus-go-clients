/*
vAudience AIgentChat API

Testing AgentTeamAPIService

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

func Test_aigentchat_AgentTeamAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AgentTeamAPIService AddAgentToTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string
		var agentId string

		resp, httpRes, err := apiClient.AgentTeamAPI.AddAgentToTeam(context.Background(), orgId, teamId, agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService CreateAgentTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AgentTeamAPI.CreateAgentTeam(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService DeleteAgentTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.AgentTeamAPI.DeleteAgentTeam(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService GetAgentTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.AgentTeamAPI.GetAgentTeam(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService GetFullAgentTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.AgentTeamAPI.GetFullAgentTeam(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService ListAgentTeams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.AgentTeamAPI.ListAgentTeams(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService RemoveAgentFromTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string
		var agentId string

		resp, httpRes, err := apiClient.AgentTeamAPI.RemoveAgentFromTeam(context.Background(), orgId, teamId, agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService UpdateAgentTeam", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.AgentTeamAPI.UpdateAgentTeam(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService UpdateSystemMessages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.AgentTeamAPI.UpdateSystemMessages(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentTeamAPIService UpdateUserMessages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.AgentTeamAPI.UpdateUserMessages(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
