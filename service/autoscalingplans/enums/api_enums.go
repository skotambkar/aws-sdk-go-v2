// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ForecastDataType string

// Enum values for ForecastDataType
const (
	ForecastDataTypeCapacityForecast           ForecastDataType = "CapacityForecast"
	ForecastDataTypeLoadForecast               ForecastDataType = "LoadForecast"
	ForecastDataTypeScheduledActionMinCapacity ForecastDataType = "ScheduledActionMinCapacity"
	ForecastDataTypeScheduledActionMaxCapacity ForecastDataType = "ScheduledActionMaxCapacity"
)

func (enum ForecastDataType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ForecastDataType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LoadMetricType string

// Enum values for LoadMetricType
const (
	LoadMetricTypeAsgtotalCpuutilization     LoadMetricType = "ASGTotalCPUUtilization"
	LoadMetricTypeAsgtotalNetworkIn          LoadMetricType = "ASGTotalNetworkIn"
	LoadMetricTypeAsgtotalNetworkOut         LoadMetricType = "ASGTotalNetworkOut"
	LoadMetricTypeAlbtargetGroupRequestCount LoadMetricType = "ALBTargetGroupRequestCount"
)

func (enum LoadMetricType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LoadMetricType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MetricStatistic string

// Enum values for MetricStatistic
const (
	MetricStatisticAverage     MetricStatistic = "Average"
	MetricStatisticMinimum     MetricStatistic = "Minimum"
	MetricStatisticMaximum     MetricStatistic = "Maximum"
	MetricStatisticSampleCount MetricStatistic = "SampleCount"
	MetricStatisticSum         MetricStatistic = "Sum"
)

func (enum MetricStatistic) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MetricStatistic) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeTargetTrackingScaling PolicyType = "TargetTrackingScaling"
)

func (enum PolicyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PredictiveScalingMaxCapacityBehavior string

// Enum values for PredictiveScalingMaxCapacityBehavior
const (
	PredictiveScalingMaxCapacityBehaviorSetForecastCapacityToMaxCapacity    PredictiveScalingMaxCapacityBehavior = "SetForecastCapacityToMaxCapacity"
	PredictiveScalingMaxCapacityBehaviorSetMaxCapacityToForecastCapacity    PredictiveScalingMaxCapacityBehavior = "SetMaxCapacityToForecastCapacity"
	PredictiveScalingMaxCapacityBehaviorSetMaxCapacityAboveForecastCapacity PredictiveScalingMaxCapacityBehavior = "SetMaxCapacityAboveForecastCapacity"
)

func (enum PredictiveScalingMaxCapacityBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PredictiveScalingMaxCapacityBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PredictiveScalingMode string

// Enum values for PredictiveScalingMode
const (
	PredictiveScalingModeForecastAndScale PredictiveScalingMode = "ForecastAndScale"
	PredictiveScalingModeForecastOnly     PredictiveScalingMode = "ForecastOnly"
)

func (enum PredictiveScalingMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PredictiveScalingMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalableDimension string

// Enum values for ScalableDimension
const (
	ScalableDimensionAutoscalingAutoScalingGroupDesiredCapacity ScalableDimension = "autoscaling:autoScalingGroup:DesiredCapacity"
	ScalableDimensionEcsServiceDesiredCount                     ScalableDimension = "ecs:service:DesiredCount"
	ScalableDimensionEc2SpotFleetRequestTargetCapacity          ScalableDimension = "ec2:spot-fleet-request:TargetCapacity"
	ScalableDimensionRdsClusterReadReplicaCount                 ScalableDimension = "rds:cluster:ReadReplicaCount"
	ScalableDimensionDynamodbTableReadCapacityUnits             ScalableDimension = "dynamodb:table:ReadCapacityUnits"
	ScalableDimensionDynamodbTableWriteCapacityUnits            ScalableDimension = "dynamodb:table:WriteCapacityUnits"
	ScalableDimensionDynamodbIndexReadCapacityUnits             ScalableDimension = "dynamodb:index:ReadCapacityUnits"
	ScalableDimensionDynamodbIndexWriteCapacityUnits            ScalableDimension = "dynamodb:index:WriteCapacityUnits"
)

func (enum ScalableDimension) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalableDimension) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalingMetricType string

// Enum values for ScalingMetricType
const (
	ScalingMetricTypeAsgaverageCpuutilization                 ScalingMetricType = "ASGAverageCPUUtilization"
	ScalingMetricTypeAsgaverageNetworkIn                      ScalingMetricType = "ASGAverageNetworkIn"
	ScalingMetricTypeAsgaverageNetworkOut                     ScalingMetricType = "ASGAverageNetworkOut"
	ScalingMetricTypeDynamoDbreadCapacityUtilization          ScalingMetricType = "DynamoDBReadCapacityUtilization"
	ScalingMetricTypeDynamoDbwriteCapacityUtilization         ScalingMetricType = "DynamoDBWriteCapacityUtilization"
	ScalingMetricTypeEcsserviceAverageCpuutilization          ScalingMetricType = "ECSServiceAverageCPUUtilization"
	ScalingMetricTypeEcsserviceAverageMemoryUtilization       ScalingMetricType = "ECSServiceAverageMemoryUtilization"
	ScalingMetricTypeAlbrequestCountPerTarget                 ScalingMetricType = "ALBRequestCountPerTarget"
	ScalingMetricTypeRdsreaderAverageCpuutilization           ScalingMetricType = "RDSReaderAverageCPUUtilization"
	ScalingMetricTypeRdsreaderAverageDatabaseConnections      ScalingMetricType = "RDSReaderAverageDatabaseConnections"
	ScalingMetricTypeEc2spotFleetRequestAverageCpuutilization ScalingMetricType = "EC2SpotFleetRequestAverageCPUUtilization"
	ScalingMetricTypeEc2spotFleetRequestAverageNetworkIn      ScalingMetricType = "EC2SpotFleetRequestAverageNetworkIn"
	ScalingMetricTypeEc2spotFleetRequestAverageNetworkOut     ScalingMetricType = "EC2SpotFleetRequestAverageNetworkOut"
)

func (enum ScalingMetricType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalingMetricType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalingPlanStatusCode string

// Enum values for ScalingPlanStatusCode
const (
	ScalingPlanStatusCodeActive             ScalingPlanStatusCode = "Active"
	ScalingPlanStatusCodeActiveWithProblems ScalingPlanStatusCode = "ActiveWithProblems"
	ScalingPlanStatusCodeCreationInProgress ScalingPlanStatusCode = "CreationInProgress"
	ScalingPlanStatusCodeCreationFailed     ScalingPlanStatusCode = "CreationFailed"
	ScalingPlanStatusCodeDeletionInProgress ScalingPlanStatusCode = "DeletionInProgress"
	ScalingPlanStatusCodeDeletionFailed     ScalingPlanStatusCode = "DeletionFailed"
	ScalingPlanStatusCodeUpdateInProgress   ScalingPlanStatusCode = "UpdateInProgress"
	ScalingPlanStatusCodeUpdateFailed       ScalingPlanStatusCode = "UpdateFailed"
)

func (enum ScalingPlanStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalingPlanStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalingPolicyUpdateBehavior string

// Enum values for ScalingPolicyUpdateBehavior
const (
	ScalingPolicyUpdateBehaviorKeepExternalPolicies    ScalingPolicyUpdateBehavior = "KeepExternalPolicies"
	ScalingPolicyUpdateBehaviorReplaceExternalPolicies ScalingPolicyUpdateBehavior = "ReplaceExternalPolicies"
)

func (enum ScalingPolicyUpdateBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalingPolicyUpdateBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalingStatusCode string

// Enum values for ScalingStatusCode
const (
	ScalingStatusCodeInactive        ScalingStatusCode = "Inactive"
	ScalingStatusCodePartiallyActive ScalingStatusCode = "PartiallyActive"
	ScalingStatusCodeActive          ScalingStatusCode = "Active"
)

func (enum ScalingStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalingStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServiceNamespace string

// Enum values for ServiceNamespace
const (
	ServiceNamespaceAutoscaling ServiceNamespace = "autoscaling"
	ServiceNamespaceEcs         ServiceNamespace = "ecs"
	ServiceNamespaceEc2         ServiceNamespace = "ec2"
	ServiceNamespaceRds         ServiceNamespace = "rds"
	ServiceNamespaceDynamodb    ServiceNamespace = "dynamodb"
)

func (enum ServiceNamespace) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceNamespace) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}