// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type AgentUpdateStatus string

// Enum values for AgentUpdateStatus
const (
	AgentUpdateStatusPending  AgentUpdateStatus = "PENDING"
	AgentUpdateStatusStaging  AgentUpdateStatus = "STAGING"
	AgentUpdateStatusStaged   AgentUpdateStatus = "STAGED"
	AgentUpdateStatusUpdating AgentUpdateStatus = "UPDATING"
	AgentUpdateStatusUpdated  AgentUpdateStatus = "UPDATED"
	AgentUpdateStatusFailed   AgentUpdateStatus = "FAILED"
)

func (enum AgentUpdateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AgentUpdateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AssignPublicIp string

// Enum values for AssignPublicIp
const (
	AssignPublicIpEnabled  AssignPublicIp = "ENABLED"
	AssignPublicIpDisabled AssignPublicIp = "DISABLED"
)

func (enum AssignPublicIp) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AssignPublicIp) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ClusterField string

// Enum values for ClusterField
const (
	ClusterFieldStatistics ClusterField = "STATISTICS"
	ClusterFieldTags       ClusterField = "TAGS"
)

func (enum ClusterField) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ClusterField) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ClusterSettingName string

// Enum values for ClusterSettingName
const (
	ClusterSettingNameContainerInsights ClusterSettingName = "containerInsights"
)

func (enum ClusterSettingName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ClusterSettingName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Compatibility string

// Enum values for Compatibility
const (
	CompatibilityEc2     Compatibility = "EC2"
	CompatibilityFargate Compatibility = "FARGATE"
)

func (enum Compatibility) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Compatibility) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Connectivity string

// Enum values for Connectivity
const (
	ConnectivityConnected    Connectivity = "CONNECTED"
	ConnectivityDisconnected Connectivity = "DISCONNECTED"
)

func (enum Connectivity) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Connectivity) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContainerCondition string

// Enum values for ContainerCondition
const (
	ContainerConditionStart    ContainerCondition = "START"
	ContainerConditionComplete ContainerCondition = "COMPLETE"
	ContainerConditionSuccess  ContainerCondition = "SUCCESS"
	ContainerConditionHealthy  ContainerCondition = "HEALTHY"
)

func (enum ContainerCondition) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContainerCondition) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContainerInstanceField string

// Enum values for ContainerInstanceField
const (
	ContainerInstanceFieldTags ContainerInstanceField = "TAGS"
)

func (enum ContainerInstanceField) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContainerInstanceField) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContainerInstanceStatus string

// Enum values for ContainerInstanceStatus
const (
	ContainerInstanceStatusActive             ContainerInstanceStatus = "ACTIVE"
	ContainerInstanceStatusDraining           ContainerInstanceStatus = "DRAINING"
	ContainerInstanceStatusRegistering        ContainerInstanceStatus = "REGISTERING"
	ContainerInstanceStatusDeregistering      ContainerInstanceStatus = "DEREGISTERING"
	ContainerInstanceStatusRegistrationFailed ContainerInstanceStatus = "REGISTRATION_FAILED"
)

func (enum ContainerInstanceStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContainerInstanceStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeploymentControllerType string

// Enum values for DeploymentControllerType
const (
	DeploymentControllerTypeEcs        DeploymentControllerType = "ECS"
	DeploymentControllerTypeCodeDeploy DeploymentControllerType = "CODE_DEPLOY"
	DeploymentControllerTypeExternal   DeploymentControllerType = "EXTERNAL"
)

func (enum DeploymentControllerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeploymentControllerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DesiredStatus string

// Enum values for DesiredStatus
const (
	DesiredStatusRunning DesiredStatus = "RUNNING"
	DesiredStatusPending DesiredStatus = "PENDING"
	DesiredStatusStopped DesiredStatus = "STOPPED"
)

func (enum DesiredStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DesiredStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceCgroupPermission string

// Enum values for DeviceCgroupPermission
const (
	DeviceCgroupPermissionRead  DeviceCgroupPermission = "read"
	DeviceCgroupPermissionWrite DeviceCgroupPermission = "write"
	DeviceCgroupPermissionMknod DeviceCgroupPermission = "mknod"
)

func (enum DeviceCgroupPermission) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceCgroupPermission) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FirelensConfigurationType string

// Enum values for FirelensConfigurationType
const (
	FirelensConfigurationTypeFluentd   FirelensConfigurationType = "fluentd"
	FirelensConfigurationTypeFluentbit FirelensConfigurationType = "fluentbit"
)

func (enum FirelensConfigurationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FirelensConfigurationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HealthStatus string

// Enum values for HealthStatus
const (
	HealthStatusHealthy   HealthStatus = "HEALTHY"
	HealthStatusUnhealthy HealthStatus = "UNHEALTHY"
	HealthStatusUnknown   HealthStatus = "UNKNOWN"
)

func (enum HealthStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HealthStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IpcMode string

// Enum values for IpcMode
const (
	IpcModeHost IpcMode = "host"
	IpcModeTask IpcMode = "task"
	IpcModeNone IpcMode = "none"
)

func (enum IpcMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IpcMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LaunchType string

// Enum values for LaunchType
const (
	LaunchTypeEc2     LaunchType = "EC2"
	LaunchTypeFargate LaunchType = "FARGATE"
)

func (enum LaunchType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LaunchType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogDriver string

// Enum values for LogDriver
const (
	LogDriverJsonFile    LogDriver = "json-file"
	LogDriverSyslog      LogDriver = "syslog"
	LogDriverJournald    LogDriver = "journald"
	LogDriverGelf        LogDriver = "gelf"
	LogDriverFluentd     LogDriver = "fluentd"
	LogDriverAwslogs     LogDriver = "awslogs"
	LogDriverSplunk      LogDriver = "splunk"
	LogDriverAwsfirelens LogDriver = "awsfirelens"
)

func (enum LogDriver) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogDriver) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NetworkMode string

// Enum values for NetworkMode
const (
	NetworkModeBridge NetworkMode = "bridge"
	NetworkModeHost   NetworkMode = "host"
	NetworkModeAwsvpc NetworkMode = "awsvpc"
	NetworkModeNone   NetworkMode = "none"
)

func (enum NetworkMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NetworkMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PidMode string

// Enum values for PidMode
const (
	PidModeHost PidMode = "host"
	PidModeTask PidMode = "task"
)

func (enum PidMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PidMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PlacementConstraintType string

// Enum values for PlacementConstraintType
const (
	PlacementConstraintTypeDistinctInstance PlacementConstraintType = "distinctInstance"
	PlacementConstraintTypeMemberOf         PlacementConstraintType = "memberOf"
)

func (enum PlacementConstraintType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PlacementConstraintType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PlacementStrategyType string

// Enum values for PlacementStrategyType
const (
	PlacementStrategyTypeRandom  PlacementStrategyType = "random"
	PlacementStrategyTypeSpread  PlacementStrategyType = "spread"
	PlacementStrategyTypeBinpack PlacementStrategyType = "binpack"
)

func (enum PlacementStrategyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PlacementStrategyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PlatformDeviceType string

// Enum values for PlatformDeviceType
const (
	PlatformDeviceTypeGpu PlatformDeviceType = "GPU"
)

func (enum PlatformDeviceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PlatformDeviceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PropagateTags string

// Enum values for PropagateTags
const (
	PropagateTagsTaskDefinition PropagateTags = "TASK_DEFINITION"
	PropagateTagsService        PropagateTags = "SERVICE"
)

func (enum PropagateTags) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PropagateTags) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProxyConfigurationType string

// Enum values for ProxyConfigurationType
const (
	ProxyConfigurationTypeAppmesh ProxyConfigurationType = "APPMESH"
)

func (enum ProxyConfigurationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProxyConfigurationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeGpu                  ResourceType = "GPU"
	ResourceTypeInferenceAccelerator ResourceType = "InferenceAccelerator"
)

func (enum ResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScaleUnit string

// Enum values for ScaleUnit
const (
	ScaleUnitPercent ScaleUnit = "PERCENT"
)

func (enum ScaleUnit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScaleUnit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SchedulingStrategy string

// Enum values for SchedulingStrategy
const (
	SchedulingStrategyReplica SchedulingStrategy = "REPLICA"
	SchedulingStrategyDaemon  SchedulingStrategy = "DAEMON"
)

func (enum SchedulingStrategy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SchedulingStrategy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Scope string

// Enum values for Scope
const (
	ScopeTask   Scope = "task"
	ScopeShared Scope = "shared"
)

func (enum Scope) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Scope) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServiceField string

// Enum values for ServiceField
const (
	ServiceFieldTags ServiceField = "TAGS"
)

func (enum ServiceField) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceField) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SettingName string

// Enum values for SettingName
const (
	SettingNameServiceLongArnFormat           SettingName = "serviceLongArnFormat"
	SettingNameTaskLongArnFormat              SettingName = "taskLongArnFormat"
	SettingNameContainerInstanceLongArnFormat SettingName = "containerInstanceLongArnFormat"
	SettingNameAwsvpcTrunking                 SettingName = "awsvpcTrunking"
	SettingNameContainerInsights              SettingName = "containerInsights"
)

func (enum SettingName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SettingName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

func (enum SortOrder) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortOrder) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StabilityStatus string

// Enum values for StabilityStatus
const (
	StabilityStatusSteadyState StabilityStatus = "STEADY_STATE"
	StabilityStatusStabilizing StabilityStatus = "STABILIZING"
)

func (enum StabilityStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StabilityStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeContainerInstance TargetType = "container-instance"
)

func (enum TargetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TargetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskDefinitionFamilyStatus string

// Enum values for TaskDefinitionFamilyStatus
const (
	TaskDefinitionFamilyStatusActive   TaskDefinitionFamilyStatus = "ACTIVE"
	TaskDefinitionFamilyStatusInactive TaskDefinitionFamilyStatus = "INACTIVE"
	TaskDefinitionFamilyStatusAll      TaskDefinitionFamilyStatus = "ALL"
)

func (enum TaskDefinitionFamilyStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskDefinitionFamilyStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskDefinitionField string

// Enum values for TaskDefinitionField
const (
	TaskDefinitionFieldTags TaskDefinitionField = "TAGS"
)

func (enum TaskDefinitionField) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskDefinitionField) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskDefinitionPlacementConstraintType string

// Enum values for TaskDefinitionPlacementConstraintType
const (
	TaskDefinitionPlacementConstraintTypeMemberOf TaskDefinitionPlacementConstraintType = "memberOf"
)

func (enum TaskDefinitionPlacementConstraintType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskDefinitionPlacementConstraintType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskDefinitionStatus string

// Enum values for TaskDefinitionStatus
const (
	TaskDefinitionStatusActive   TaskDefinitionStatus = "ACTIVE"
	TaskDefinitionStatusInactive TaskDefinitionStatus = "INACTIVE"
)

func (enum TaskDefinitionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskDefinitionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskField string

// Enum values for TaskField
const (
	TaskFieldTags TaskField = "TAGS"
)

func (enum TaskField) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskField) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TaskStopCode string

// Enum values for TaskStopCode
const (
	TaskStopCodeTaskFailedToStart        TaskStopCode = "TaskFailedToStart"
	TaskStopCodeEssentialContainerExited TaskStopCode = "EssentialContainerExited"
	TaskStopCodeUserInitiated            TaskStopCode = "UserInitiated"
)

func (enum TaskStopCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TaskStopCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TransportProtocol string

// Enum values for TransportProtocol
const (
	TransportProtocolTcp TransportProtocol = "tcp"
	TransportProtocolUdp TransportProtocol = "udp"
)

func (enum TransportProtocol) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TransportProtocol) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UlimitName string

// Enum values for UlimitName
const (
	UlimitNameCore       UlimitName = "core"
	UlimitNameCpu        UlimitName = "cpu"
	UlimitNameData       UlimitName = "data"
	UlimitNameFsize      UlimitName = "fsize"
	UlimitNameLocks      UlimitName = "locks"
	UlimitNameMemlock    UlimitName = "memlock"
	UlimitNameMsgqueue   UlimitName = "msgqueue"
	UlimitNameNice       UlimitName = "nice"
	UlimitNameNofile     UlimitName = "nofile"
	UlimitNameNproc      UlimitName = "nproc"
	UlimitNameRss        UlimitName = "rss"
	UlimitNameRtprio     UlimitName = "rtprio"
	UlimitNameRttime     UlimitName = "rttime"
	UlimitNameSigpending UlimitName = "sigpending"
	UlimitNameStack      UlimitName = "stack"
)

func (enum UlimitName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UlimitName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
