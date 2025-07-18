/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.19.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// ApiErrorCode the model 'ApiErrorCode'
type ApiErrorCode string

// List of ApiErrorCode
const (
	ErrCodeNil ApiErrorCode = ""
	ErrCodePermissionDenied ApiErrorCode = "permission_denied"
	ErrCodeUnknown ApiErrorCode = "unknown"
	ErrCodeUnauthorized ApiErrorCode = "unauthorized"
	ErrCodeInvalidPayload ApiErrorCode = "invalid_payload"
	ErrCodeInvalidEntity ApiErrorCode = "invalid_entity"
	ErrCodeInvalidParams ApiErrorCode = "invalid_params"
	ErrCodeInvalidDate ApiErrorCode = "invalid_date"
	ErrCodeStartDateMustBeBeforeEndDate ApiErrorCode = "start_date_must_be_before_end_date"
	ErrCodeAgentNotFound ApiErrorCode = "agent_not_found"
	ErrCodeInvalidAgentID ApiErrorCode = "invalid_agent_id"
	ErrCodeFailedToCreateAgent ApiErrorCode = "failed_to_create_agent"
	ErrCodeFailedToCreateDefaultAgents ApiErrorCode = "failed_to_create_default_agents"
	ErrCodeFailedToDeleteAgent ApiErrorCode = "failed_to_delete_agent"
	ErrCodeFailedToGetAgent ApiErrorCode = "failed_to_get_agent"
	ErrCodeFailedToUpdateAgent ApiErrorCode = "failed_to_update_agent"
	ErrCodeAIModelNotFound ApiErrorCode = "ai_model_not_found"
	ErrCodeInvalidAIModelID ApiErrorCode = "invalid_ai_model_id"
	ErrCodeInvalidAIModelDefinition ApiErrorCode = "invalid_ai_model_definition"
	ErrCodeFailedToCreateAIModel ApiErrorCode = "failed_to_create_ai_model"
	ErrCodeFailedToUpdateAIModel ApiErrorCode = "failed_to_update_ai_model"
	ErrCodeFailedToDeleteAIModel ApiErrorCode = "failed_to_delete_ai_model"
	ErrCodeFailedToFetchAIModel ApiErrorCode = "failed_to_fetch_ai_model"
	ErrCodeInvalidAIModelServiceID ApiErrorCode = "invalid_ai_model_service_id"
	ErrCodeAIModelServiceNotFound ApiErrorCode = "ai_model_service_not_found"
	ErrCodeFailedToCreateAIModelService ApiErrorCode = "failed_to_create_ai_model_service"
	ErrCodeFailedToUpdateAIModelService ApiErrorCode = "failed_to_update_ai_model_service"
	ErrCodeFailedToDeleteAIModelService ApiErrorCode = "failed_to_delete_ai_model_service"
	ErrCodeFailedToFetchAIModelService ApiErrorCode = "failed_to_fetch_ai_model_service"
	ErrCodeAiServiceBadRequest ApiErrorCode = "ai_service_bad_request"
	ErrCodeAiServiceContentFilter ApiErrorCode = "ai_service_content_filter"
	ErrCodeAiServiceAuthentication ApiErrorCode = "ai_service_authentication"
	ErrCodeAiServicePermissionDenied ApiErrorCode = "ai_service_permission_denied"
	ErrCodeAiServiceRequestTooLarge ApiErrorCode = "ai_service_request_too_large"
	ErrCodeAiServiceRateLimit ApiErrorCode = "ai_service_rate_limit"
	ErrCodeAiServiceInternalServerError ApiErrorCode = "ai_service_internal_server_error"
	ErrCodeAiServiceOverload ApiErrorCode = "ai_service_overload"
	ErrCodeAiServiceUnprocessableEntity ApiErrorCode = "ai_service_unprocessable_entity"
	ErrCodeNoAudioFileSaved ApiErrorCode = "no_audio_file_saved"
	ErrCodeCreatingAudioTranscriptionContext ApiErrorCode = "failed_to_create_audio_transcription_context"
	ErrCodeCreatingAudioGenerationContext ApiErrorCode = "failed_to_create_audio_generation_context"
	ErrCodeFailedToFinalizeAudioStream ApiErrorCode = "failed_to_finalize_audio_stream"
	ErrCodeFailedToGenerateAudio ApiErrorCode = "failed_to_generate_audio"
	ErrCodeChannelNotFound ApiErrorCode = "channel_not_found"
	ErrCodeInvalidChannelID ApiErrorCode = "invalid_channel_id"
	ErrCodeFailedToCreateChannel ApiErrorCode = "failed_to_create_channel"
	ErrCodeFailedToUpdateChannel ApiErrorCode = "failed_to_update_channel"
	ErrCodeFailedToDeleteChannel ApiErrorCode = "failed_to_delete_channel"
	ErrCodeFailedToFetchChannel ApiErrorCode = "failed_to_fetch_channel"
	ErrCodeFailedToFetchChannels ApiErrorCode = "failed_to_fetch_channels"
	ErrCodeFailedToFetchSubscribed ApiErrorCode = "failed_to_fetch_subscribed_channels"
	ErrCodeFailedToUploadChannelFile ApiErrorCode = "failed_to_upload_channel_file"
	ErrCodeFailedToStoreChatCompletionUser ApiErrorCode = "failed_to_store_chat_completion_user"
	ErrCodeFailedToCreateChatContext ApiErrorCode = "failed_to_create_chat_context"
	ErrCodeFailedToAcquireChatCompletionLock ApiErrorCode = "failed_to_acquire_chat_completion_lock"
	ErrCodeChatCompletionIsLockedForChannel ApiErrorCode = "chat_completion_is_locked_for_channel"
	ErrCodeFailedToCheckUserAuthorization ApiErrorCode = "failed_to_check_user_authorization"
	ErrCodeFailedToReleaseChatCompletionLock ApiErrorCode = "failed_to_release_chat_completion_lock"
	ErrCodeCapabilityNotAvailable ApiErrorCode = "capability_not_available"
	ErrCodeNoMessagesFound ApiErrorCode = "no_messages_found"
	ErrCodeInvalidRequestParameters ApiErrorCode = "invalid_request_parameters"
	ErrCodeChatCompletionTooManyTempFiles ApiErrorCode = "chat_completion_too_many_temporary_files"
	ErrCodeChatCompletionExceedsTotalTempFileSize ApiErrorCode = "chat_completion_exceeds_total_temporary_file_size"
	ErrCodeChatCompletionExceedsMaxInputTokens ApiErrorCode = "chat_completion_exceeds_max_input_tokens"
	ErrCodeChatCompletionExceedsMaxToolCalls ApiErrorCode = "chat_completion_exceeds_max_tool_calls"
	ErrCodeFailedToCreateConnectionToken ApiErrorCode = "failed_to_create_connection_token"
	ErrCodeFailedToEmbedText ApiErrorCode = "failed_to_embed_text"
	ErrCodeFailedToSearchExecutionLogs ApiErrorCode = "failed_to_search_execution_logs"
	ErrCodeInvalidFileID ApiErrorCode = "invalid_file_id"
	ErrCodeInvalidStorageUrl ApiErrorCode = "invalid_storage_url"
	ErrCodeFileUploadInfoNotFound ApiErrorCode = "file_upload_info_not_found"
	ErrCodeInvalidFileUploadSize ApiErrorCode = "invalid_file_upload_size"
	ErrCodeInvalidFileUploadMimeType ApiErrorCode = "invalid_file_upload_mime_type"
	ErrCodeProcessingChecksFailed ApiErrorCode = "processing_checks_failed"
	ErrCodeFunctionCallNotFound ApiErrorCode = "functioncall_not_found"
	ErrCodeFailedToGenerateImage ApiErrorCode = "failed_to_generate_image"
	ErrCodeInvalidImageID ApiErrorCode = "invalid_image_id"
	ErrCodeImageNotFound ApiErrorCode = "image_not_found"
	ErrCodeFailedToDeleteImage ApiErrorCode = "failed_to_delete_image"
	ErrCodeFailedToGetImage ApiErrorCode = "failed_to_get_image"
	ErrCodeFailedToSearchImages ApiErrorCode = "failed_to_search_images"
	ErrCodeFailedToStoreImages ApiErrorCode = "failed_to_store_images"
	ErrCodeInvalidMessageID ApiErrorCode = "invalid_message_id"
	ErrCodeMessageNotFound ApiErrorCode = "message_not_found"
	ErrCodeFailedToUpdateMessage ApiErrorCode = "failed_to_update_message"
	ErrCodeFailedToDeleteMessage ApiErrorCode = "failed_to_delete_message"
	ErrCodeFailedToFetchMessage ApiErrorCode = "failed_to_fetch_message"
	ErrCodeFailedToSearchMessage ApiErrorCode = "failed_to_search_messages"
	ErrCodeFailedToCreateMessage ApiErrorCode = "failed_to_create_message"
	ErrCodeFailedToRetrieveChannelMessages ApiErrorCode = "failed_to_retrieve_channel_messages"
	ErrCodeFailedToRetrieveMessagesFromList ApiErrorCode = "failed_to_retrieve_messages_from_list"
	ErrCodeFailedToFormUserMessage ApiErrorCode = "failed_to_form_user_message"
	ErrCodeMissionNotFound ApiErrorCode = "mission_not_found"
	ErrCodeMissionsNotFound ApiErrorCode = "missions_not_found"
	ErrCodeFailedToGetMission ApiErrorCode = "failed_to_get_mission"
	ErrCodeFailedToCreateMission ApiErrorCode = "failed_to_create_mission"
	ErrCodeFailedToDeleteMission ApiErrorCode = "failed_to_delete_mission"
	ErrCodeInvalidMissionID ApiErrorCode = "invalid_mission_id"
	ErrCodeInvalidMissionExecutorID ApiErrorCode = "invalid mission executor ID"
	ErrCodeMissionNotDone ApiErrorCode = "mission_not_done"
	ErrCodeInvalidExecutorID ApiErrorCode = "invalid_executor_id"
	ErrCodeInvalidOrgID ApiErrorCode = "invalid_organization_id"
	ErrCodeOrgCostBudgetNotFound ApiErrorCode = "organization_cost_budget_not_found"
	ErrCodeFailedToCreateBudget ApiErrorCode = "failed_to_create_organization_cost_budget"
	ErrCodeFailedToUpdateBudget ApiErrorCode = "failed_to_update_organization_cost_budget"
	ErrCodeFailedToGetOrgCostBudget ApiErrorCode = "failed_to_get_organization_cost_budget"
	ErrCodeInsufficientBudget ApiErrorCode = "insufficient_budget"
	ErrCodeOrgCostBudgetExists ApiErrorCode = "organization_cost_budget_exists"
	ErrCodeInvalidToolID ApiErrorCode = "invalid_tool_id"
	ErrCodeUnknownToolID ApiErrorCode = "unknown_tool_id"
	ErrCodeFailedToCreateTempApiKey ApiErrorCode = "failed_to_create_temp_api_key"
	ErrCodeInvalidPromptID ApiErrorCode = "invalid_prompt_id"
	ErrCodePromptNotFound ApiErrorCode = "prompt_not_found"
	ErrCodeFailedToGetPrompt ApiErrorCode = "failed_to_get_prompt"
	ErrCodeFailedToUpdatePrompt ApiErrorCode = "failed_to_update_prompt"
	ErrCodeFailedToDeletePrompt ApiErrorCode = "failed_to_delete_prompt"
	ErrCodeInvalidPromptData ApiErrorCode = "invalid_prompt_data"
	ErrCodePromptAlreadyExists ApiErrorCode = "prompt_already_exists"
	ErrCodeInvalidVersion ApiErrorCode = "invalid_prompt_version"
	ErrCodeVariableNotFound ApiErrorCode = "variable_not_found"
	ErrCodeInvalidVariableValue ApiErrorCode = "invalid_variable_value"
)

// All allowed values of ApiErrorCode enum
var AllowedApiErrorCodeEnumValues = []ApiErrorCode{
	"",
	"permission_denied",
	"unknown",
	"unauthorized",
	"invalid_payload",
	"invalid_entity",
	"invalid_params",
	"invalid_date",
	"start_date_must_be_before_end_date",
	"agent_not_found",
	"invalid_agent_id",
	"failed_to_create_agent",
	"failed_to_create_default_agents",
	"failed_to_delete_agent",
	"failed_to_get_agent",
	"failed_to_update_agent",
	"ai_model_not_found",
	"invalid_ai_model_id",
	"invalid_ai_model_definition",
	"failed_to_create_ai_model",
	"failed_to_update_ai_model",
	"failed_to_delete_ai_model",
	"failed_to_fetch_ai_model",
	"invalid_ai_model_service_id",
	"ai_model_service_not_found",
	"failed_to_create_ai_model_service",
	"failed_to_update_ai_model_service",
	"failed_to_delete_ai_model_service",
	"failed_to_fetch_ai_model_service",
	"ai_service_bad_request",
	"ai_service_content_filter",
	"ai_service_authentication",
	"ai_service_permission_denied",
	"ai_service_request_too_large",
	"ai_service_rate_limit",
	"ai_service_internal_server_error",
	"ai_service_overload",
	"ai_service_unprocessable_entity",
	"no_audio_file_saved",
	"failed_to_create_audio_transcription_context",
	"failed_to_create_audio_generation_context",
	"failed_to_finalize_audio_stream",
	"failed_to_generate_audio",
	"channel_not_found",
	"invalid_channel_id",
	"failed_to_create_channel",
	"failed_to_update_channel",
	"failed_to_delete_channel",
	"failed_to_fetch_channel",
	"failed_to_fetch_channels",
	"failed_to_fetch_subscribed_channels",
	"failed_to_upload_channel_file",
	"failed_to_store_chat_completion_user",
	"failed_to_create_chat_context",
	"failed_to_acquire_chat_completion_lock",
	"chat_completion_is_locked_for_channel",
	"failed_to_check_user_authorization",
	"failed_to_release_chat_completion_lock",
	"capability_not_available",
	"no_messages_found",
	"invalid_request_parameters",
	"chat_completion_too_many_temporary_files",
	"chat_completion_exceeds_total_temporary_file_size",
	"chat_completion_exceeds_max_input_tokens",
	"chat_completion_exceeds_max_tool_calls",
	"failed_to_create_connection_token",
	"failed_to_embed_text",
	"failed_to_search_execution_logs",
	"invalid_file_id",
	"invalid_storage_url",
	"file_upload_info_not_found",
	"invalid_file_upload_size",
	"invalid_file_upload_mime_type",
	"processing_checks_failed",
	"functioncall_not_found",
	"failed_to_generate_image",
	"invalid_image_id",
	"image_not_found",
	"failed_to_delete_image",
	"failed_to_get_image",
	"failed_to_search_images",
	"failed_to_store_images",
	"invalid_message_id",
	"message_not_found",
	"failed_to_update_message",
	"failed_to_delete_message",
	"failed_to_fetch_message",
	"failed_to_search_messages",
	"failed_to_create_message",
	"failed_to_retrieve_channel_messages",
	"failed_to_retrieve_messages_from_list",
	"failed_to_form_user_message",
	"mission_not_found",
	"missions_not_found",
	"failed_to_get_mission",
	"failed_to_create_mission",
	"failed_to_delete_mission",
	"invalid_mission_id",
	"invalid mission executor ID",
	"mission_not_done",
	"invalid_executor_id",
	"invalid_organization_id",
	"organization_cost_budget_not_found",
	"failed_to_create_organization_cost_budget",
	"failed_to_update_organization_cost_budget",
	"failed_to_get_organization_cost_budget",
	"insufficient_budget",
	"organization_cost_budget_exists",
	"invalid_tool_id",
	"unknown_tool_id",
	"failed_to_create_temp_api_key",
	"invalid_prompt_id",
	"prompt_not_found",
	"failed_to_get_prompt",
	"failed_to_update_prompt",
	"failed_to_delete_prompt",
	"invalid_prompt_data",
	"prompt_already_exists",
	"invalid_prompt_version",
	"variable_not_found",
	"invalid_variable_value",
}

func (v *ApiErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiErrorCode(value)
	for _, existing := range AllowedApiErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiErrorCode", value)
}

// NewApiErrorCodeFromValue returns a pointer to a valid ApiErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiErrorCodeFromValue(v string) (*ApiErrorCode, error) {
	ev := ApiErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiErrorCode: valid values are %v", v, AllowedApiErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiErrorCode) IsValid() bool {
	for _, existing := range AllowedApiErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiErrorCode value
func (v ApiErrorCode) Ptr() *ApiErrorCode {
	return &v
}

type NullableApiErrorCode struct {
	value *ApiErrorCode
	isSet bool
}

func (v NullableApiErrorCode) Get() *ApiErrorCode {
	return v.value
}

func (v *NullableApiErrorCode) Set(val *ApiErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorCode(val *ApiErrorCode) *NullableApiErrorCode {
	return &NullableApiErrorCode{value: val, isSet: true}
}

func (v NullableApiErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

