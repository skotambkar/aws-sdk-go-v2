// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type AbortAction string

// Enum values for AbortAction
const (
	AbortActionCancel AbortAction = "CANCEL"
)

func (enum AbortAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AbortAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ActionType string

// Enum values for ActionType
const (
	ActionTypePublish   ActionType = "PUBLISH"
	ActionTypeSubscribe ActionType = "SUBSCRIBE"
	ActionTypeReceive   ActionType = "RECEIVE"
	ActionTypeConnect   ActionType = "CONNECT"
)

func (enum ActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of alert target: one of "SNS".
type AlertTargetType string

// Enum values for AlertTargetType
const (
	AlertTargetTypeSns AlertTargetType = "SNS"
)

func (enum AlertTargetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AlertTargetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditCheckRunStatus string

// Enum values for AuditCheckRunStatus
const (
	AuditCheckRunStatusInProgress               AuditCheckRunStatus = "IN_PROGRESS"
	AuditCheckRunStatusWaitingForDataCollection AuditCheckRunStatus = "WAITING_FOR_DATA_COLLECTION"
	AuditCheckRunStatusCanceled                 AuditCheckRunStatus = "CANCELED"
	AuditCheckRunStatusCompletedCompliant       AuditCheckRunStatus = "COMPLETED_COMPLIANT"
	AuditCheckRunStatusCompletedNonCompliant    AuditCheckRunStatus = "COMPLETED_NON_COMPLIANT"
	AuditCheckRunStatusFailed                   AuditCheckRunStatus = "FAILED"
)

func (enum AuditCheckRunStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditCheckRunStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditFindingSeverity string

// Enum values for AuditFindingSeverity
const (
	AuditFindingSeverityCritical AuditFindingSeverity = "CRITICAL"
	AuditFindingSeverityHigh     AuditFindingSeverity = "HIGH"
	AuditFindingSeverityMedium   AuditFindingSeverity = "MEDIUM"
	AuditFindingSeverityLow      AuditFindingSeverity = "LOW"
)

func (enum AuditFindingSeverity) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditFindingSeverity) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditFrequency string

// Enum values for AuditFrequency
const (
	AuditFrequencyDaily    AuditFrequency = "DAILY"
	AuditFrequencyWeekly   AuditFrequency = "WEEKLY"
	AuditFrequencyBiweekly AuditFrequency = "BIWEEKLY"
	AuditFrequencyMonthly  AuditFrequency = "MONTHLY"
)

func (enum AuditFrequency) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditFrequency) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditMitigationActionsExecutionStatus string

// Enum values for AuditMitigationActionsExecutionStatus
const (
	AuditMitigationActionsExecutionStatusInProgress AuditMitigationActionsExecutionStatus = "IN_PROGRESS"
	AuditMitigationActionsExecutionStatusCompleted  AuditMitigationActionsExecutionStatus = "COMPLETED"
	AuditMitigationActionsExecutionStatusFailed     AuditMitigationActionsExecutionStatus = "FAILED"
	AuditMitigationActionsExecutionStatusCanceled   AuditMitigationActionsExecutionStatus = "CANCELED"
	AuditMitigationActionsExecutionStatusSkipped    AuditMitigationActionsExecutionStatus = "SKIPPED"
	AuditMitigationActionsExecutionStatusPending    AuditMitigationActionsExecutionStatus = "PENDING"
)

func (enum AuditMitigationActionsExecutionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditMitigationActionsExecutionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditMitigationActionsTaskStatus string

// Enum values for AuditMitigationActionsTaskStatus
const (
	AuditMitigationActionsTaskStatusInProgress AuditMitigationActionsTaskStatus = "IN_PROGRESS"
	AuditMitigationActionsTaskStatusCompleted  AuditMitigationActionsTaskStatus = "COMPLETED"
	AuditMitigationActionsTaskStatusFailed     AuditMitigationActionsTaskStatus = "FAILED"
	AuditMitigationActionsTaskStatusCanceled   AuditMitigationActionsTaskStatus = "CANCELED"
)

func (enum AuditMitigationActionsTaskStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditMitigationActionsTaskStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditNotificationType string

// Enum values for AuditNotificationType
const (
	AuditNotificationTypeSns AuditNotificationType = "SNS"
)

func (enum AuditNotificationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditNotificationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditTaskStatus string

// Enum values for AuditTaskStatus
const (
	AuditTaskStatusInProgress AuditTaskStatus = "IN_PROGRESS"
	AuditTaskStatusCompleted  AuditTaskStatus = "COMPLETED"
	AuditTaskStatusFailed     AuditTaskStatus = "FAILED"
	AuditTaskStatusCanceled   AuditTaskStatus = "CANCELED"
)

func (enum AuditTaskStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditTaskStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuditTaskType string

// Enum values for AuditTaskType
const (
	AuditTaskTypeOnDemandAuditTask  AuditTaskType = "ON_DEMAND_AUDIT_TASK"
	AuditTaskTypeScheduledAuditTask AuditTaskType = "SCHEDULED_AUDIT_TASK"
)

func (enum AuditTaskType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuditTaskType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuthDecision string

// Enum values for AuthDecision
const (
	AuthDecisionAllowed      AuthDecision = "ALLOWED"
	AuthDecisionExplicitDeny AuthDecision = "EXPLICIT_DENY"
	AuthDecisionImplicitDeny AuthDecision = "IMPLICIT_DENY"
)

func (enum AuthDecision) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuthDecision) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AuthorizerStatus string

// Enum values for AuthorizerStatus
const (
	AuthorizerStatusActive   AuthorizerStatus = "ACTIVE"
	AuthorizerStatusInactive AuthorizerStatus = "INACTIVE"
)

func (enum AuthorizerStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AuthorizerStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AutoRegistrationStatus string

// Enum values for AutoRegistrationStatus
const (
	AutoRegistrationStatusEnable  AutoRegistrationStatus = "ENABLE"
	AutoRegistrationStatusDisable AutoRegistrationStatus = "DISABLE"
)

func (enum AutoRegistrationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AutoRegistrationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CACertificateStatus string

// Enum values for CACertificateStatus
const (
	CACertificateStatusActive   CACertificateStatus = "ACTIVE"
	CACertificateStatusInactive CACertificateStatus = "INACTIVE"
)

func (enum CACertificateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CACertificateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CACertificateUpdateAction string

// Enum values for CACertificateUpdateAction
const (
	CACertificateUpdateActionDeactivate CACertificateUpdateAction = "DEACTIVATE"
)

func (enum CACertificateUpdateAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CACertificateUpdateAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CannedAccessControlList string

// Enum values for CannedAccessControlList
const (
	CannedAccessControlListPrivate                CannedAccessControlList = "private"
	CannedAccessControlListPublicRead             CannedAccessControlList = "public-read"
	CannedAccessControlListPublicReadWrite        CannedAccessControlList = "public-read-write"
	CannedAccessControlListAwsExecRead            CannedAccessControlList = "aws-exec-read"
	CannedAccessControlListAuthenticatedRead      CannedAccessControlList = "authenticated-read"
	CannedAccessControlListBucketOwnerRead        CannedAccessControlList = "bucket-owner-read"
	CannedAccessControlListBucketOwnerFullControl CannedAccessControlList = "bucket-owner-full-control"
	CannedAccessControlListLogDeliveryWrite       CannedAccessControlList = "log-delivery-write"
)

func (enum CannedAccessControlList) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CannedAccessControlList) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CertificateStatus string

// Enum values for CertificateStatus
const (
	CertificateStatusActive            CertificateStatus = "ACTIVE"
	CertificateStatusInactive          CertificateStatus = "INACTIVE"
	CertificateStatusRevoked           CertificateStatus = "REVOKED"
	CertificateStatusPendingTransfer   CertificateStatus = "PENDING_TRANSFER"
	CertificateStatusRegisterInactive  CertificateStatus = "REGISTER_INACTIVE"
	CertificateStatusPendingActivation CertificateStatus = "PENDING_ACTIVATION"
)

func (enum CertificateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CertificateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorLessThan          ComparisonOperator = "less-than"
	ComparisonOperatorLessThanEquals    ComparisonOperator = "less-than-equals"
	ComparisonOperatorGreaterThan       ComparisonOperator = "greater-than"
	ComparisonOperatorGreaterThanEquals ComparisonOperator = "greater-than-equals"
	ComparisonOperatorInCidrSet         ComparisonOperator = "in-cidr-set"
	ComparisonOperatorNotInCidrSet      ComparisonOperator = "not-in-cidr-set"
	ComparisonOperatorInPortSet         ComparisonOperator = "in-port-set"
	ComparisonOperatorNotInPortSet      ComparisonOperator = "not-in-port-set"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DayOfWeek string

// Enum values for DayOfWeek
const (
	DayOfWeekSun DayOfWeek = "SUN"
	DayOfWeekMon DayOfWeek = "MON"
	DayOfWeekTue DayOfWeek = "TUE"
	DayOfWeekWed DayOfWeek = "WED"
	DayOfWeekThu DayOfWeek = "THU"
	DayOfWeekFri DayOfWeek = "FRI"
	DayOfWeekSat DayOfWeek = "SAT"
)

func (enum DayOfWeek) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DayOfWeek) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceCertificateUpdateAction string

// Enum values for DeviceCertificateUpdateAction
const (
	DeviceCertificateUpdateActionDeactivate DeviceCertificateUpdateAction = "DEACTIVATE"
)

func (enum DeviceCertificateUpdateAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceCertificateUpdateAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DynamicGroupStatus string

// Enum values for DynamicGroupStatus
const (
	DynamicGroupStatusActive     DynamicGroupStatus = "ACTIVE"
	DynamicGroupStatusBuilding   DynamicGroupStatus = "BUILDING"
	DynamicGroupStatusRebuilding DynamicGroupStatus = "REBUILDING"
)

func (enum DynamicGroupStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DynamicGroupStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DynamoKeyType string

// Enum values for DynamoKeyType
const (
	DynamoKeyTypeString DynamoKeyType = "STRING"
	DynamoKeyTypeNumber DynamoKeyType = "NUMBER"
)

func (enum DynamoKeyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DynamoKeyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EventType string

// Enum values for EventType
const (
	EventTypeThing                EventType = "THING"
	EventTypeThingGroup           EventType = "THING_GROUP"
	EventTypeThingType            EventType = "THING_TYPE"
	EventTypeThingGroupMembership EventType = "THING_GROUP_MEMBERSHIP"
	EventTypeThingGroupHierarchy  EventType = "THING_GROUP_HIERARCHY"
	EventTypeThingTypeAssociation EventType = "THING_TYPE_ASSOCIATION"
	EventTypeJob                  EventType = "JOB"
	EventTypeJobExecution         EventType = "JOB_EXECUTION"
	EventTypePolicy               EventType = "POLICY"
	EventTypeCertificate          EventType = "CERTIFICATE"
	EventTypeCaCertificate        EventType = "CA_CERTIFICATE"
)

func (enum EventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FieldType string

// Enum values for FieldType
const (
	FieldTypeNumber  FieldType = "Number"
	FieldTypeString  FieldType = "String"
	FieldTypeBoolean FieldType = "Boolean"
)

func (enum FieldType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FieldType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IndexStatus string

// Enum values for IndexStatus
const (
	IndexStatusActive     IndexStatus = "ACTIVE"
	IndexStatusBuilding   IndexStatus = "BUILDING"
	IndexStatusRebuilding IndexStatus = "REBUILDING"
)

func (enum IndexStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IndexStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobExecutionFailureType string

// Enum values for JobExecutionFailureType
const (
	JobExecutionFailureTypeFailed   JobExecutionFailureType = "FAILED"
	JobExecutionFailureTypeRejected JobExecutionFailureType = "REJECTED"
	JobExecutionFailureTypeTimedOut JobExecutionFailureType = "TIMED_OUT"
	JobExecutionFailureTypeAll      JobExecutionFailureType = "ALL"
)

func (enum JobExecutionFailureType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobExecutionFailureType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobExecutionStatus string

// Enum values for JobExecutionStatus
const (
	JobExecutionStatusQueued     JobExecutionStatus = "QUEUED"
	JobExecutionStatusInProgress JobExecutionStatus = "IN_PROGRESS"
	JobExecutionStatusSucceeded  JobExecutionStatus = "SUCCEEDED"
	JobExecutionStatusFailed     JobExecutionStatus = "FAILED"
	JobExecutionStatusTimedOut   JobExecutionStatus = "TIMED_OUT"
	JobExecutionStatusRejected   JobExecutionStatus = "REJECTED"
	JobExecutionStatusRemoved    JobExecutionStatus = "REMOVED"
	JobExecutionStatusCanceled   JobExecutionStatus = "CANCELED"
)

func (enum JobExecutionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobExecutionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusInProgress         JobStatus = "IN_PROGRESS"
	JobStatusCanceled           JobStatus = "CANCELED"
	JobStatusCompleted          JobStatus = "COMPLETED"
	JobStatusDeletionInProgress JobStatus = "DELETION_IN_PROGRESS"
)

func (enum JobStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelDebug    LogLevel = "DEBUG"
	LogLevelInfo     LogLevel = "INFO"
	LogLevelError    LogLevel = "ERROR"
	LogLevelWarn     LogLevel = "WARN"
	LogLevelDisabled LogLevel = "DISABLED"
)

func (enum LogLevel) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogLevel) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogTargetType string

// Enum values for LogTargetType
const (
	LogTargetTypeDefault    LogTargetType = "DEFAULT"
	LogTargetTypeThingGroup LogTargetType = "THING_GROUP"
)

func (enum LogTargetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogTargetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MessageFormat string

// Enum values for MessageFormat
const (
	MessageFormatRaw  MessageFormat = "RAW"
	MessageFormatJson MessageFormat = "JSON"
)

func (enum MessageFormat) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MessageFormat) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MitigationActionType string

// Enum values for MitigationActionType
const (
	MitigationActionTypeUpdateDeviceCertificate     MitigationActionType = "UPDATE_DEVICE_CERTIFICATE"
	MitigationActionTypeUpdateCaCertificate         MitigationActionType = "UPDATE_CA_CERTIFICATE"
	MitigationActionTypeAddThingsToThingGroup       MitigationActionType = "ADD_THINGS_TO_THING_GROUP"
	MitigationActionTypeReplaceDefaultPolicyVersion MitigationActionType = "REPLACE_DEFAULT_POLICY_VERSION"
	MitigationActionTypeEnableIotLogging            MitigationActionType = "ENABLE_IOT_LOGGING"
	MitigationActionTypePublishFindingToSns         MitigationActionType = "PUBLISH_FINDING_TO_SNS"
)

func (enum MitigationActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MitigationActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OTAUpdateStatus string

// Enum values for OTAUpdateStatus
const (
	OTAUpdateStatusCreatePending    OTAUpdateStatus = "CREATE_PENDING"
	OTAUpdateStatusCreateInProgress OTAUpdateStatus = "CREATE_IN_PROGRESS"
	OTAUpdateStatusCreateComplete   OTAUpdateStatus = "CREATE_COMPLETE"
	OTAUpdateStatusCreateFailed     OTAUpdateStatus = "CREATE_FAILED"
)

func (enum OTAUpdateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OTAUpdateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyTemplateName string

// Enum values for PolicyTemplateName
const (
	PolicyTemplateNameBlankPolicy PolicyTemplateName = "BLANK_POLICY"
)

func (enum PolicyTemplateName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyTemplateName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReportType string

// Enum values for ReportType
const (
	ReportTypeErrors  ReportType = "ERRORS"
	ReportTypeResults ReportType = "RESULTS"
)

func (enum ReportType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReportType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeDeviceCertificate   ResourceType = "DEVICE_CERTIFICATE"
	ResourceTypeCaCertificate       ResourceType = "CA_CERTIFICATE"
	ResourceTypeIotPolicy           ResourceType = "IOT_POLICY"
	ResourceTypeCognitoIdentityPool ResourceType = "COGNITO_IDENTITY_POOL"
	ResourceTypeClientId            ResourceType = "CLIENT_ID"
	ResourceTypeAccountSettings     ResourceType = "ACCOUNT_SETTINGS"
)

func (enum ResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Status string

// Enum values for Status
const (
	StatusInProgress Status = "InProgress"
	StatusCompleted  Status = "Completed"
	StatusFailed     Status = "Failed"
	StatusCancelled  Status = "Cancelled"
	StatusCancelling Status = "Cancelling"
)

func (enum Status) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Status) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TargetSelection string

// Enum values for TargetSelection
const (
	TargetSelectionContinuous TargetSelection = "CONTINUOUS"
	TargetSelectionSnapshot   TargetSelection = "SNAPSHOT"
)

func (enum TargetSelection) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TargetSelection) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ThingConnectivityIndexingMode string

// Enum values for ThingConnectivityIndexingMode
const (
	ThingConnectivityIndexingModeOff    ThingConnectivityIndexingMode = "OFF"
	ThingConnectivityIndexingModeStatus ThingConnectivityIndexingMode = "STATUS"
)

func (enum ThingConnectivityIndexingMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThingConnectivityIndexingMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ThingGroupIndexingMode string

// Enum values for ThingGroupIndexingMode
const (
	ThingGroupIndexingModeOff ThingGroupIndexingMode = "OFF"
	ThingGroupIndexingModeOn  ThingGroupIndexingMode = "ON"
)

func (enum ThingGroupIndexingMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThingGroupIndexingMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ThingIndexingMode string

// Enum values for ThingIndexingMode
const (
	ThingIndexingModeOff               ThingIndexingMode = "OFF"
	ThingIndexingModeRegistry          ThingIndexingMode = "REGISTRY"
	ThingIndexingModeRegistryAndShadow ThingIndexingMode = "REGISTRY_AND_SHADOW"
)

func (enum ThingIndexingMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThingIndexingMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TopicRuleDestinationStatus string

// Enum values for TopicRuleDestinationStatus
const (
	TopicRuleDestinationStatusEnabled    TopicRuleDestinationStatus = "ENABLED"
	TopicRuleDestinationStatusInProgress TopicRuleDestinationStatus = "IN_PROGRESS"
	TopicRuleDestinationStatusDisabled   TopicRuleDestinationStatus = "DISABLED"
	TopicRuleDestinationStatusError      TopicRuleDestinationStatus = "ERROR"
)

func (enum TopicRuleDestinationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TopicRuleDestinationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ViolationEventType string

// Enum values for ViolationEventType
const (
	ViolationEventTypeInAlarm          ViolationEventType = "in-alarm"
	ViolationEventTypeAlarmCleared     ViolationEventType = "alarm-cleared"
	ViolationEventTypeAlarmInvalidated ViolationEventType = "alarm-invalidated"
)

func (enum ViolationEventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ViolationEventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
