/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// ErrorCode the model 'ErrorCode'
type ErrorCode string

// List of ErrorCode
const (
	ErrCodePermissionDenied ErrorCode = "permission_denied"
	ErrCodeRequestBodyInvalid ErrorCode = "request_body_invalid"
	ErrCodeUnknown ErrorCode = "unknown"
	ErrCodeUnauthorized ErrorCode = "unauthorized"
	ErrCodeUserEmailNotInToken ErrorCode = "user_email_not_in_token"
	ErrCodeUserIdNotInToken ErrorCode = "user_id_not_in_token"
	ErrCodeValidationFailed ErrorCode = "validation_failed"
	ErrCodeAssociationIdInvalid ErrorCode = "association_id_invalid"
	ErrCodeAssociationRoleIdInvalid ErrorCode = "association_role_id_invalid"
	ErrCodeCreatedByInvalid ErrorCode = "created_by_invalid"
	ErrCodeMemberIdInvalid ErrorCode = "member_id_invalid"
	ErrCodeOrganizationIdInvalid ErrorCode = "organization_id_invalid"
	ErrCodeOrganizationApiKeyIdInvalid ErrorCode = "organization_api_key_id_invalid"
	ErrCodeOrganizationInviteIdInvalid ErrorCode = "organization_invite_id_invalid"
	ErrCodeOrganizationRoleIdInvalid ErrorCode = "organization_role_id_invalid"
	ErrCodeProductRoleIdInvalid ErrorCode = "product_role_id_invalid"
	ErrCodeRoleIdInvalid ErrorCode = "role_id_invalid"
	ErrCodeTeamIdInvalid ErrorCode = "team_id_invalid"
	ErrCodeUserIdInvalid ErrorCode = "user_id_invalid"
	ErrCodeUserSettingsIdInvalid ErrorCode = "user_settings_id_invalid"
	ErrCodeAssociationMemberNotFound ErrorCode = "association_member_not_found"
	ErrCodeAssociationRoleNotFound ErrorCode = "association_role_not_found"
	ErrCodeOrganizationNotFound ErrorCode = "organization_not_found"
	ErrCodeOrganizationApiKeyNotFound ErrorCode = "organization_api_key_not_found"
	ErrCodeOrganizationInviteNotFound ErrorCode = "organization_invite_not_found"
	ErrCodeOrganizationSubscriptionNotFound ErrorCode = "organization_subscription_not_found"
	ErrCodeProductNotFound ErrorCode = "product_not_found"
	ErrCodeRoleNotFound ErrorCode = "role_not_found"
	ErrCodeTeamNotFound ErrorCode = "team_not_found"
	ErrCodeUserRoleNotFound ErrorCode = "user_role_not_found"
	ErrCodeUserSettingsNotFound ErrorCode = "user_settings_not_found"
	ErrCodeUserNotFound ErrorCode = "user_not_found"
	ErrCodeAuthStateMismatch ErrorCode = "auth_state_mismatch"
	ErrCodeAuthCodeIsRequired ErrorCode = "auth_code_is_required"
	ErrCodeAuthReturnToIsRequired ErrorCode = "auth_return_to_is_required"
	ErrCodeAuthRefreshTokenIsRequired ErrorCode = "auth_refresh_token_is_required"
	ErrCodeAuthVerifierIsRequired ErrorCode = "auth_verifier_is_required"
	ErrCodeAssociationRoleIsDefault ErrorCode = "association_role_is_default"
	ErrCodeAssociationRoleInUse ErrorCode = "association_role_in_use"
	ErrCodeAssociationMemberHasRole ErrorCode = "association_member_has_role"
	ErrCodeAssociationMemberHasNotRole ErrorCode = "association_member_has_not_role"
	ErrCodeOrganizationHasSubscription ErrorCode = "organization_has_subscription"
	ErrCodeOrganizationHasSubscriptionPlan ErrorCode = "organization_has_subscription_plan"
	ErrCodeOrganizationHasSubscriptionSeats ErrorCode = "organization_has_subscription_seats"
	ErrCodeOrganizationHasActiveSubscription ErrorCode = "organization_has_active_subscription"
	ErrCodeOrganizationHasNoActiveSubscription ErrorCode = "organization_has_no_active_subscription"
	ErrCodeOrganizationHasNoCustomer ErrorCode = "organization_has_no_customer"
	ErrCodeOrganizationMaximumSeatsReached ErrorCode = "organization_seats_maximum_reached"
	ErrCodeOrganizationMinimumSeatsRequired ErrorCode = "organization_seats_minimum_required"
	ErrCodeOrganizationUserIsNotMember ErrorCode = "organization_user_is_not_member"
	ErrCodeOrganizationSubscriptionCancelAtPeriodEnd ErrorCode = "organization_subscription_cancel_at_period_end"
	ErrCodeOrganizationSubscriptionCanceled ErrorCode = "organization_subscription_canceled"
	ErrCodeOrganizationSubscriptionInactive ErrorCode = "organization_subscription_inactive"
	ErrCodeOrganizationSubscriptionIncompletePastDue ErrorCode = "organization_subscription_incomplete_past_due"
	ErrCodeOrganizationSubscriptionPeriodEndInvalid ErrorCode = "organization_subscription_period_end_invalid"
	ErrCodeOrganizationSubscriptionRequiresPaymentMethod ErrorCode = "organization_subscription_requires_payment_method"
	ErrCodeOrganizationSubscriptionRequiresAction ErrorCode = "organization_subscription_requires_action"
	ErrCodeOrganizationSubscriptionNotCanceledAt ErrorCode = "organization_subscription_not_canceled_at"
	ErrCodeOrganizationInviteExpired ErrorCode = "organization_invite_expired"
	ErrCodeOrganizationInviteEmailMismatch ErrorCode = "organization_invite_email_mismatch"
	ErrCodeOrganizationInviteOrgIdMismatch ErrorCode = "organization_invite_org_id_mismatch"
	ErrCodeOrganizationInviteSelfInvite ErrorCode = "organization_invite_self_invite"
	ErrCodeProductTypeInvalid ErrorCode = "product_type_invalid"
	ErrCodeStripeSignatureHeaderMissing ErrorCode = "stripe_signature_header_missing"
	ErrCodeStripeEventConstructionFailed ErrorCode = "stripe_event_construction_failed"
	ErrCodeStripeDisabled ErrorCode = "stripe_disabled"
	ErrCodeOrganizationApiKeyExists ErrorCode = "organization_api_key_exists"
	ErrCodeOrganizationMemberExists ErrorCode = "organization_member_exists"
	ErrCodeOrganizationInviteEmailExists ErrorCode = "organization_invite_email_exists"
	ErrCodeOrganizationInviteMemberExists ErrorCode = "organization_invite_member_exists"
	ErrCodeOrganizationInviteAcceptMemberExists ErrorCode = "organization_invite_accept_member_exists"
	ErrCodeOrganizationSettingsExists ErrorCode = "organization_settings_exists"
	ErrCodeUserRoleExists ErrorCode = "user_role_exists"
)

// All allowed values of ErrorCode enum
var AllowedErrorCodeEnumValues = []ErrorCode{
	"permission_denied",
	"request_body_invalid",
	"unknown",
	"unauthorized",
	"user_email_not_in_token",
	"user_id_not_in_token",
	"validation_failed",
	"association_id_invalid",
	"association_role_id_invalid",
	"created_by_invalid",
	"member_id_invalid",
	"organization_id_invalid",
	"organization_api_key_id_invalid",
	"organization_invite_id_invalid",
	"organization_role_id_invalid",
	"product_role_id_invalid",
	"role_id_invalid",
	"team_id_invalid",
	"user_id_invalid",
	"user_settings_id_invalid",
	"association_member_not_found",
	"association_role_not_found",
	"organization_not_found",
	"organization_api_key_not_found",
	"organization_invite_not_found",
	"organization_subscription_not_found",
	"product_not_found",
	"role_not_found",
	"team_not_found",
	"user_role_not_found",
	"user_settings_not_found",
	"user_not_found",
	"auth_state_mismatch",
	"auth_code_is_required",
	"auth_return_to_is_required",
	"auth_refresh_token_is_required",
	"auth_verifier_is_required",
	"association_role_is_default",
	"association_role_in_use",
	"association_member_has_role",
	"association_member_has_not_role",
	"organization_has_subscription",
	"organization_has_subscription_plan",
	"organization_has_subscription_seats",
	"organization_has_active_subscription",
	"organization_has_no_active_subscription",
	"organization_has_no_customer",
	"organization_seats_maximum_reached",
	"organization_seats_minimum_required",
	"organization_user_is_not_member",
	"organization_subscription_cancel_at_period_end",
	"organization_subscription_canceled",
	"organization_subscription_inactive",
	"organization_subscription_incomplete_past_due",
	"organization_subscription_period_end_invalid",
	"organization_subscription_requires_payment_method",
	"organization_subscription_requires_action",
	"organization_subscription_not_canceled_at",
	"organization_invite_expired",
	"organization_invite_email_mismatch",
	"organization_invite_org_id_mismatch",
	"organization_invite_self_invite",
	"product_type_invalid",
	"stripe_signature_header_missing",
	"stripe_event_construction_failed",
	"stripe_disabled",
	"organization_api_key_exists",
	"organization_member_exists",
	"organization_invite_email_exists",
	"organization_invite_member_exists",
	"organization_invite_accept_member_exists",
	"organization_settings_exists",
	"user_role_exists",
}

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// NewErrorCodeFromValue returns a pointer to a valid ErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodeFromValue(v string) (*ErrorCode, error) {
	ev := ErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCode: valid values are %v", v, AllowedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCode) IsValid() bool {
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

