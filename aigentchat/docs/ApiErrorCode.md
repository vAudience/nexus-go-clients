# ApiErrorCode

## Enum


* `ErrCodeNil` (value: `""`)

* `ErrCodePermissionDenied` (value: `"permission_denied"`)

* `ErrCodeUnknown` (value: `"unknown"`)

* `ErrCodeUnauthorized` (value: `"unauthorized"`)

* `ErrCodeInvalidPayload` (value: `"invalid_payload"`)

* `ErrCodeInvalidEntity` (value: `"invalid_entity"`)

* `ErrCodeInvalidParams` (value: `"invalid_params"`)

* `ErrCodeInvalidDate` (value: `"invalid_date"`)

* `ErrCodeStartDateMustBeBeforeEndDate` (value: `"start_date_must_be_before_end_date"`)

* `ErrCodeAgentNotFound` (value: `"agent_not_found"`)

* `ErrCodeInvalidAgentID` (value: `"invalid_agent_id"`)

* `ErrCodeFailedToCreateAgent` (value: `"failed_to_create_agent"`)

* `ErrCodeFailedToCreateDefaultAgents` (value: `"failed_to_create_default_agents"`)

* `ErrCodeFailedToDeleteAgent` (value: `"failed_to_delete_agent"`)

* `ErrCodeFailedToGetAgent` (value: `"failed_to_get_agent"`)

* `ErrCodeFailedToUpdateAgent` (value: `"failed_to_update_agent"`)

* `ErrCodeAIModelNotFound` (value: `"ai_model_not_found"`)

* `ErrCodeInvalidAIModelID` (value: `"invalid_ai_model_id"`)

* `ErrCodeInvalidAIModelDefinition` (value: `"invalid_ai_model_definition"`)

* `ErrCodeFailedToCreateAIModel` (value: `"failed_to_create_ai_model"`)

* `ErrCodeFailedToUpdateAIModel` (value: `"failed_to_update_ai_model"`)

* `ErrCodeFailedToDeleteAIModel` (value: `"failed_to_delete_ai_model"`)

* `ErrCodeFailedToFetchAIModel` (value: `"failed_to_fetch_ai_model"`)

* `ErrCodeInvalidAIModelServiceID` (value: `"invalid_ai_model_service_id"`)

* `ErrCodeAIModelServiceNotFound` (value: `"ai_model_service_not_found"`)

* `ErrCodeFailedToCreateAIModelService` (value: `"failed_to_create_ai_model_service"`)

* `ErrCodeFailedToUpdateAIModelService` (value: `"failed_to_update_ai_model_service"`)

* `ErrCodeFailedToDeleteAIModelService` (value: `"failed_to_delete_ai_model_service"`)

* `ErrCodeFailedToFetchAIModelService` (value: `"failed_to_fetch_ai_model_service"`)

* `ErrCodeAiServiceBadRequest` (value: `"ai_service_bad_request"`)

* `ErrCodeAiServiceContentFilter` (value: `"ai_service_content_filter"`)

* `ErrCodeAiServiceAuthentication` (value: `"ai_service_authentication"`)

* `ErrCodeAiServicePermissionDenied` (value: `"ai_service_permission_denied"`)

* `ErrCodeAiServiceRequestTooLarge` (value: `"ai_service_request_too_large"`)

* `ErrCodeAiServiceRateLimit` (value: `"ai_service_rate_limit"`)

* `ErrCodeAiServiceInternalServerError` (value: `"ai_service_internal_server_error"`)

* `ErrCodeAiServiceOverload` (value: `"ai_service_overload"`)

* `ErrCodeAiServiceUnprocessableEntity` (value: `"ai_service_unprocessable_entity"`)

* `ErrCodeNoAudioFileSaved` (value: `"no_audio_file_saved"`)

* `ErrCodeCreatingAudioTranscriptionContext` (value: `"failed_to_create_audio_transcription_context"`)

* `ErrCodeCreatingAudioGenerationContext` (value: `"failed_to_create_audio_generation_context"`)

* `ErrCodeFailedToFinalizeAudioStream` (value: `"failed_to_finalize_audio_stream"`)

* `ErrCodeFailedToGenerateAudio` (value: `"failed_to_generate_audio"`)

* `ErrCodeChannelNotFound` (value: `"channel_not_found"`)

* `ErrCodeInvalidChannelID` (value: `"invalid_channel_id"`)

* `ErrCodeFailedToCreateChannel` (value: `"failed_to_create_channel"`)

* `ErrCodeFailedToUpdateChannel` (value: `"failed_to_update_channel"`)

* `ErrCodeFailedToDeleteChannel` (value: `"failed_to_delete_channel"`)

* `ErrCodeFailedToFetchChannel` (value: `"failed_to_fetch_channel"`)

* `ErrCodeFailedToFetchChannels` (value: `"failed_to_fetch_channels"`)

* `ErrCodeFailedToFetchSubscribed` (value: `"failed_to_fetch_subscribed_channels"`)

* `ErrCodeFailedToUploadChannelFile` (value: `"failed_to_upload_channel_file"`)

* `ErrCodeFailedToStoreChatCompletionUser` (value: `"failed_to_store_chat_completion_user"`)

* `ErrCodeFailedToCreateChatContext` (value: `"failed_to_create_chat_context"`)

* `ErrCodeFailedToAcquireChatCompletionLock` (value: `"failed_to_acquire_chat_completion_lock"`)

* `ErrCodeChatCompletionIsLockedForChannel` (value: `"chat_completion_is_locked_for_channel"`)

* `ErrCodeFailedToCheckUserAuthorization` (value: `"failed_to_check_user_authorization"`)

* `ErrCodeFailedToReleaseChatCompletionLock` (value: `"failed_to_release_chat_completion_lock"`)

* `ErrCodeCapabilityNotAvailable` (value: `"capability_not_available"`)

* `ErrCodeNoMessagesFound` (value: `"no_messages_found"`)

* `ErrCodeInvalidRequestParameters` (value: `"invalid_request_parameters"`)

* `ErrCodeChatCompletionTooManyTempFiles` (value: `"chat_completion_too_many_temporary_files"`)

* `ErrCodeChatCompletionExceedsTotalTempFileSize` (value: `"chat_completion_exceeds_total_temporary_file_size"`)

* `ErrCodeChatCompletionExceedsMaxInputTokens` (value: `"chat_completion_exceeds_max_input_tokens"`)

* `ErrCodeChatCompletionExceedsMaxToolCalls` (value: `"chat_completion_exceeds_max_tool_calls"`)

* `ErrCodeFailedToCreateConnectionToken` (value: `"failed_to_create_connection_token"`)

* `ErrCodeFailedToEmbedText` (value: `"failed_to_embed_text"`)

* `ErrCodeFailedToSearchExecutionLogs` (value: `"failed_to_search_execution_logs"`)

* `ErrCodeInvalidFileID` (value: `"invalid_file_id"`)

* `ErrCodeInvalidStorageUrl` (value: `"invalid_storage_url"`)

* `ErrCodeFileUploadInfoNotFound` (value: `"file_upload_info_not_found"`)

* `ErrCodeInvalidFileUploadSize` (value: `"invalid_file_upload_size"`)

* `ErrCodeInvalidFileUploadMimeType` (value: `"invalid_file_upload_mime_type"`)

* `ErrCodeProcessingChecksFailed` (value: `"processing_checks_failed"`)

* `ErrCodeFunctionCallNotFound` (value: `"functioncall_not_found"`)

* `ErrCodeFailedToGenerateImage` (value: `"failed_to_generate_image"`)

* `ErrCodeInvalidImageID` (value: `"invalid_image_id"`)

* `ErrCodeImageNotFound` (value: `"image_not_found"`)

* `ErrCodeFailedToDeleteImage` (value: `"failed_to_delete_image"`)

* `ErrCodeFailedToGetImage` (value: `"failed_to_get_image"`)

* `ErrCodeFailedToSearchImages` (value: `"failed_to_search_images"`)

* `ErrCodeFailedToStoreImages` (value: `"failed_to_store_images"`)

* `ErrCodeInvalidMessageID` (value: `"invalid_message_id"`)

* `ErrCodeMessageNotFound` (value: `"message_not_found"`)

* `ErrCodeFailedToUpdateMessage` (value: `"failed_to_update_message"`)

* `ErrCodeFailedToDeleteMessage` (value: `"failed_to_delete_message"`)

* `ErrCodeFailedToFetchMessage` (value: `"failed_to_fetch_message"`)

* `ErrCodeFailedToSearchMessage` (value: `"failed_to_search_messages"`)

* `ErrCodeFailedToCreateMessage` (value: `"failed_to_create_message"`)

* `ErrCodeFailedToRetrieveChannelMessages` (value: `"failed_to_retrieve_channel_messages"`)

* `ErrCodeFailedToRetrieveMessagesFromList` (value: `"failed_to_retrieve_messages_from_list"`)

* `ErrCodeFailedToFormUserMessage` (value: `"failed_to_form_user_message"`)

* `ErrCodeMissionNotFound` (value: `"mission_not_found"`)

* `ErrCodeMissionsNotFound` (value: `"missions_not_found"`)

* `ErrCodeFailedToGetMission` (value: `"failed_to_get_mission"`)

* `ErrCodeFailedToCreateMission` (value: `"failed_to_create_mission"`)

* `ErrCodeFailedToDeleteMission` (value: `"failed_to_delete_mission"`)

* `ErrCodeInvalidMissionID` (value: `"invalid_mission_id"`)

* `ErrCodeInvalidMissionExecutorID` (value: `"invalid mission executor ID"`)

* `ErrCodeMissionNotDone` (value: `"mission_not_done"`)

* `ErrCodeInvalidExecutorID` (value: `"invalid_executor_id"`)

* `ErrCodeInvalidOrgID` (value: `"invalid_organization_id"`)

* `ErrCodeOrgCostBudgetNotFound` (value: `"organization_cost_budget_not_found"`)

* `ErrCodeFailedToCreateBudget` (value: `"failed_to_create_organization_cost_budget"`)

* `ErrCodeFailedToUpdateBudget` (value: `"failed_to_update_organization_cost_budget"`)

* `ErrCodeFailedToGetOrgCostBudget` (value: `"failed_to_get_organization_cost_budget"`)

* `ErrCodeInsufficientBudget` (value: `"insufficient_budget"`)

* `ErrCodeOrgCostBudgetExists` (value: `"organization_cost_budget_exists"`)

* `ErrCodeInvalidToolID` (value: `"invalid_tool_id"`)

* `ErrCodeUnknownToolID` (value: `"unknown_tool_id"`)

* `ErrCodeFailedToCreateTempApiKey` (value: `"failed_to_create_temp_api_key"`)

* `ErrCodeInvalidPromptID` (value: `"invalid_prompt_id"`)

* `ErrCodePromptNotFound` (value: `"prompt_not_found"`)

* `ErrCodeFailedToGetPrompt` (value: `"failed_to_get_prompt"`)

* `ErrCodeFailedToUpdatePrompt` (value: `"failed_to_update_prompt"`)

* `ErrCodeFailedToDeletePrompt` (value: `"failed_to_delete_prompt"`)

* `ErrCodeInvalidPromptData` (value: `"invalid_prompt_data"`)

* `ErrCodePromptAlreadyExists` (value: `"prompt_already_exists"`)

* `ErrCodeInvalidVersion` (value: `"invalid_prompt_version"`)

* `ErrCodeVariableNotFound` (value: `"variable_not_found"`)

* `ErrCodeInvalidVariableValue` (value: `"invalid_variable_value"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


