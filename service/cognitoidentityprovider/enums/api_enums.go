// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type AccountTakeoverEventActionType string

// Enum values for AccountTakeoverEventActionType
const (
	AccountTakeoverEventActionTypeBlock           AccountTakeoverEventActionType = "BLOCK"
	AccountTakeoverEventActionTypeMfaIfConfigured AccountTakeoverEventActionType = "MFA_IF_CONFIGURED"
	AccountTakeoverEventActionTypeMfaRequired     AccountTakeoverEventActionType = "MFA_REQUIRED"
	AccountTakeoverEventActionTypeNoAction        AccountTakeoverEventActionType = "NO_ACTION"
)

func (enum AccountTakeoverEventActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AccountTakeoverEventActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AdvancedSecurityModeType string

// Enum values for AdvancedSecurityModeType
const (
	AdvancedSecurityModeTypeOff      AdvancedSecurityModeType = "OFF"
	AdvancedSecurityModeTypeAudit    AdvancedSecurityModeType = "AUDIT"
	AdvancedSecurityModeTypeEnforced AdvancedSecurityModeType = "ENFORCED"
)

func (enum AdvancedSecurityModeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AdvancedSecurityModeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AliasAttributeType string

// Enum values for AliasAttributeType
const (
	AliasAttributeTypePhoneNumber       AliasAttributeType = "phone_number"
	AliasAttributeTypeEmail             AliasAttributeType = "email"
	AliasAttributeTypePreferredUsername AliasAttributeType = "preferred_username"
)

func (enum AliasAttributeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AliasAttributeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AttributeDataType string

// Enum values for AttributeDataType
const (
	AttributeDataTypeString   AttributeDataType = "String"
	AttributeDataTypeNumber   AttributeDataType = "Number"
	AttributeDataTypeDateTime AttributeDataType = "DateTime"
	AttributeDataTypeBoolean  AttributeDataType = "Boolean"
)

func (enum AttributeDataType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AttributeDataType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuthFlowType string

// Enum values for AuthFlowType
const (
	AuthFlowTypeUserSrpAuth      AuthFlowType = "USER_SRP_AUTH"
	AuthFlowTypeRefreshTokenAuth AuthFlowType = "REFRESH_TOKEN_AUTH"
	AuthFlowTypeRefreshToken     AuthFlowType = "REFRESH_TOKEN"
	AuthFlowTypeCustomAuth       AuthFlowType = "CUSTOM_AUTH"
	AuthFlowTypeAdminNoSrpAuth   AuthFlowType = "ADMIN_NO_SRP_AUTH"
	AuthFlowTypeUserPasswordAuth AuthFlowType = "USER_PASSWORD_AUTH"
)

func (enum AuthFlowType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuthFlowType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChallengeName string

// Enum values for ChallengeName
const (
	ChallengeNamePassword ChallengeName = "Password"
	ChallengeNameMfa      ChallengeName = "Mfa"
)

func (enum ChallengeName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChallengeName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChallengeNameType string

// Enum values for ChallengeNameType
const (
	ChallengeNameTypeSmsMfa                 ChallengeNameType = "SMS_MFA"
	ChallengeNameTypeSoftwareTokenMfa       ChallengeNameType = "SOFTWARE_TOKEN_MFA"
	ChallengeNameTypeSelectMfaType          ChallengeNameType = "SELECT_MFA_TYPE"
	ChallengeNameTypeMfaSetup               ChallengeNameType = "MFA_SETUP"
	ChallengeNameTypePasswordVerifier       ChallengeNameType = "PASSWORD_VERIFIER"
	ChallengeNameTypeCustomChallenge        ChallengeNameType = "CUSTOM_CHALLENGE"
	ChallengeNameTypeDeviceSrpAuth          ChallengeNameType = "DEVICE_SRP_AUTH"
	ChallengeNameTypeDevicePasswordVerifier ChallengeNameType = "DEVICE_PASSWORD_VERIFIER"
	ChallengeNameTypeAdminNoSrpAuth         ChallengeNameType = "ADMIN_NO_SRP_AUTH"
	ChallengeNameTypeNewPasswordRequired    ChallengeNameType = "NEW_PASSWORD_REQUIRED"
)

func (enum ChallengeNameType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChallengeNameType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChallengeResponse string

// Enum values for ChallengeResponse
const (
	ChallengeResponseSuccess ChallengeResponse = "Success"
	ChallengeResponseFailure ChallengeResponse = "Failure"
)

func (enum ChallengeResponse) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChallengeResponse) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CompromisedCredentialsEventActionType string

// Enum values for CompromisedCredentialsEventActionType
const (
	CompromisedCredentialsEventActionTypeBlock    CompromisedCredentialsEventActionType = "BLOCK"
	CompromisedCredentialsEventActionTypeNoAction CompromisedCredentialsEventActionType = "NO_ACTION"
)

func (enum CompromisedCredentialsEventActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CompromisedCredentialsEventActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DefaultEmailOptionType string

// Enum values for DefaultEmailOptionType
const (
	DefaultEmailOptionTypeConfirmWithLink DefaultEmailOptionType = "CONFIRM_WITH_LINK"
	DefaultEmailOptionTypeConfirmWithCode DefaultEmailOptionType = "CONFIRM_WITH_CODE"
)

func (enum DefaultEmailOptionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DefaultEmailOptionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeliveryMediumType string

// Enum values for DeliveryMediumType
const (
	DeliveryMediumTypeSms   DeliveryMediumType = "SMS"
	DeliveryMediumTypeEmail DeliveryMediumType = "EMAIL"
)

func (enum DeliveryMediumType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliveryMediumType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceRememberedStatusType string

// Enum values for DeviceRememberedStatusType
const (
	DeviceRememberedStatusTypeRemembered    DeviceRememberedStatusType = "remembered"
	DeviceRememberedStatusTypeNotRemembered DeviceRememberedStatusType = "not_remembered"
)

func (enum DeviceRememberedStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceRememberedStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DomainStatusType string

// Enum values for DomainStatusType
const (
	DomainStatusTypeCreating DomainStatusType = "CREATING"
	DomainStatusTypeDeleting DomainStatusType = "DELETING"
	DomainStatusTypeUpdating DomainStatusType = "UPDATING"
	DomainStatusTypeActive   DomainStatusType = "ACTIVE"
	DomainStatusTypeFailed   DomainStatusType = "FAILED"
)

func (enum DomainStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DomainStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EmailSendingAccountType string

// Enum values for EmailSendingAccountType
const (
	EmailSendingAccountTypeCognitoDefault EmailSendingAccountType = "COGNITO_DEFAULT"
	EmailSendingAccountTypeDeveloper      EmailSendingAccountType = "DEVELOPER"
)

func (enum EmailSendingAccountType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EmailSendingAccountType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EventFilterType string

// Enum values for EventFilterType
const (
	EventFilterTypeSignIn         EventFilterType = "SIGN_IN"
	EventFilterTypePasswordChange EventFilterType = "PASSWORD_CHANGE"
	EventFilterTypeSignUp         EventFilterType = "SIGN_UP"
)

func (enum EventFilterType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventFilterType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EventResponseType string

// Enum values for EventResponseType
const (
	EventResponseTypeSuccess EventResponseType = "Success"
	EventResponseTypeFailure EventResponseType = "Failure"
)

func (enum EventResponseType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventResponseType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EventType string

// Enum values for EventType
const (
	EventTypeSignIn         EventType = "SignIn"
	EventTypeSignUp         EventType = "SignUp"
	EventTypeForgotPassword EventType = "ForgotPassword"
)

func (enum EventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExplicitAuthFlowsType string

// Enum values for ExplicitAuthFlowsType
const (
	ExplicitAuthFlowsTypeAdminNoSrpAuth     ExplicitAuthFlowsType = "ADMIN_NO_SRP_AUTH"
	ExplicitAuthFlowsTypeCustomAuthFlowOnly ExplicitAuthFlowsType = "CUSTOM_AUTH_FLOW_ONLY"
	ExplicitAuthFlowsTypeUserPasswordAuth   ExplicitAuthFlowsType = "USER_PASSWORD_AUTH"
)

func (enum ExplicitAuthFlowsType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExplicitAuthFlowsType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FeedbackValueType string

// Enum values for FeedbackValueType
const (
	FeedbackValueTypeValid   FeedbackValueType = "Valid"
	FeedbackValueTypeInvalid FeedbackValueType = "Invalid"
)

func (enum FeedbackValueType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FeedbackValueType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IdentityProviderTypeType string

// Enum values for IdentityProviderTypeType
const (
	IdentityProviderTypeTypeSaml            IdentityProviderTypeType = "SAML"
	IdentityProviderTypeTypeFacebook        IdentityProviderTypeType = "Facebook"
	IdentityProviderTypeTypeGoogle          IdentityProviderTypeType = "Google"
	IdentityProviderTypeTypeLoginWithAmazon IdentityProviderTypeType = "LoginWithAmazon"
	IdentityProviderTypeTypeOidc            IdentityProviderTypeType = "OIDC"
)

func (enum IdentityProviderTypeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IdentityProviderTypeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MessageActionType string

// Enum values for MessageActionType
const (
	MessageActionTypeResend   MessageActionType = "RESEND"
	MessageActionTypeSuppress MessageActionType = "SUPPRESS"
)

func (enum MessageActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MessageActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OAuthFlowType string

// Enum values for OAuthFlowType
const (
	OAuthFlowTypeCode              OAuthFlowType = "code"
	OAuthFlowTypeImplicit          OAuthFlowType = "implicit"
	OAuthFlowTypeClientCredentials OAuthFlowType = "client_credentials"
)

func (enum OAuthFlowType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OAuthFlowType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RiskDecisionType string

// Enum values for RiskDecisionType
const (
	RiskDecisionTypeNoRisk          RiskDecisionType = "NoRisk"
	RiskDecisionTypeAccountTakeover RiskDecisionType = "AccountTakeover"
	RiskDecisionTypeBlock           RiskDecisionType = "Block"
)

func (enum RiskDecisionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RiskDecisionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RiskLevelType string

// Enum values for RiskLevelType
const (
	RiskLevelTypeLow    RiskLevelType = "Low"
	RiskLevelTypeMedium RiskLevelType = "Medium"
	RiskLevelTypeHigh   RiskLevelType = "High"
)

func (enum RiskLevelType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RiskLevelType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StatusType string

// Enum values for StatusType
const (
	StatusTypeEnabled  StatusType = "Enabled"
	StatusTypeDisabled StatusType = "Disabled"
)

func (enum StatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UserImportJobStatusType string

// Enum values for UserImportJobStatusType
const (
	UserImportJobStatusTypeCreated    UserImportJobStatusType = "Created"
	UserImportJobStatusTypePending    UserImportJobStatusType = "Pending"
	UserImportJobStatusTypeInProgress UserImportJobStatusType = "InProgress"
	UserImportJobStatusTypeStopping   UserImportJobStatusType = "Stopping"
	UserImportJobStatusTypeExpired    UserImportJobStatusType = "Expired"
	UserImportJobStatusTypeStopped    UserImportJobStatusType = "Stopped"
	UserImportJobStatusTypeFailed     UserImportJobStatusType = "Failed"
	UserImportJobStatusTypeSucceeded  UserImportJobStatusType = "Succeeded"
)

func (enum UserImportJobStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UserImportJobStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UserPoolMfaType string

// Enum values for UserPoolMfaType
const (
	UserPoolMfaTypeOff      UserPoolMfaType = "OFF"
	UserPoolMfaTypeOn       UserPoolMfaType = "ON"
	UserPoolMfaTypeOptional UserPoolMfaType = "OPTIONAL"
)

func (enum UserPoolMfaType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UserPoolMfaType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UserStatusType string

// Enum values for UserStatusType
const (
	UserStatusTypeUnconfirmed         UserStatusType = "UNCONFIRMED"
	UserStatusTypeConfirmed           UserStatusType = "CONFIRMED"
	UserStatusTypeArchived            UserStatusType = "ARCHIVED"
	UserStatusTypeCompromised         UserStatusType = "COMPROMISED"
	UserStatusTypeUnknown             UserStatusType = "UNKNOWN"
	UserStatusTypeResetRequired       UserStatusType = "RESET_REQUIRED"
	UserStatusTypeForceChangePassword UserStatusType = "FORCE_CHANGE_PASSWORD"
)

func (enum UserStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UserStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UsernameAttributeType string

// Enum values for UsernameAttributeType
const (
	UsernameAttributeTypePhoneNumber UsernameAttributeType = "phone_number"
	UsernameAttributeTypeEmail       UsernameAttributeType = "email"
)

func (enum UsernameAttributeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UsernameAttributeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type VerifiedAttributeType string

// Enum values for VerifiedAttributeType
const (
	VerifiedAttributeTypePhoneNumber VerifiedAttributeType = "phone_number"
	VerifiedAttributeTypeEmail       VerifiedAttributeType = "email"
)

func (enum VerifiedAttributeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VerifiedAttributeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type VerifySoftwareTokenResponseType string

// Enum values for VerifySoftwareTokenResponseType
const (
	VerifySoftwareTokenResponseTypeSuccess VerifySoftwareTokenResponseType = "SUCCESS"
	VerifySoftwareTokenResponseTypeError   VerifySoftwareTokenResponseType = "ERROR"
)

func (enum VerifySoftwareTokenResponseType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VerifySoftwareTokenResponseType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
