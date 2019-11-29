// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type AdjustmentType string

// Enum values for AdjustmentType
const (
	AdjustmentTypeChangeInCapacity        AdjustmentType = "ChangeInCapacity"
	AdjustmentTypePercentChangeInCapacity AdjustmentType = "PercentChangeInCapacity"
	AdjustmentTypeExactCapacity           AdjustmentType = "ExactCapacity"
)

func (enum AdjustmentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AdjustmentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MetricAggregationType string

// Enum values for MetricAggregationType
const (
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
)

func (enum MetricAggregationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MetricAggregationType) MarshalValueBuf(b []byte) ([]byte, error) {
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

type MetricType string

// Enum values for MetricType
const (
	MetricTypeDynamoDbreadCapacityUtilization          MetricType = "DynamoDBReadCapacityUtilization"
	MetricTypeDynamoDbwriteCapacityUtilization         MetricType = "DynamoDBWriteCapacityUtilization"
	MetricTypeAlbrequestCountPerTarget                 MetricType = "ALBRequestCountPerTarget"
	MetricTypeRdsreaderAverageCpuutilization           MetricType = "RDSReaderAverageCPUUtilization"
	MetricTypeRdsreaderAverageDatabaseConnections      MetricType = "RDSReaderAverageDatabaseConnections"
	MetricTypeEc2spotFleetRequestAverageCpuutilization MetricType = "EC2SpotFleetRequestAverageCPUUtilization"
	MetricTypeEc2spotFleetRequestAverageNetworkIn      MetricType = "EC2SpotFleetRequestAverageNetworkIn"
	MetricTypeEc2spotFleetRequestAverageNetworkOut     MetricType = "EC2SpotFleetRequestAverageNetworkOut"
	MetricTypeSageMakerVariantInvocationsPerInstance   MetricType = "SageMakerVariantInvocationsPerInstance"
	MetricTypeEcsserviceAverageCpuutilization          MetricType = "ECSServiceAverageCPUUtilization"
	MetricTypeEcsserviceAverageMemoryUtilization       MetricType = "ECSServiceAverageMemoryUtilization"
)

func (enum MetricType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MetricType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeStepScaling           PolicyType = "StepScaling"
	PolicyTypeTargetTrackingScaling PolicyType = "TargetTrackingScaling"
)

func (enum PolicyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalableDimension string

// Enum values for ScalableDimension
const (
	ScalableDimensionEcsServiceDesiredCount                     ScalableDimension = "ecs:service:DesiredCount"
	ScalableDimensionEc2SpotFleetRequestTargetCapacity          ScalableDimension = "ec2:spot-fleet-request:TargetCapacity"
	ScalableDimensionElasticmapreduceInstancegroupInstanceCount ScalableDimension = "elasticmapreduce:instancegroup:InstanceCount"
	ScalableDimensionAppstreamFleetDesiredCapacity              ScalableDimension = "appstream:fleet:DesiredCapacity"
	ScalableDimensionDynamodbTableReadCapacityUnits             ScalableDimension = "dynamodb:table:ReadCapacityUnits"
	ScalableDimensionDynamodbTableWriteCapacityUnits            ScalableDimension = "dynamodb:table:WriteCapacityUnits"
	ScalableDimensionDynamodbIndexReadCapacityUnits             ScalableDimension = "dynamodb:index:ReadCapacityUnits"
	ScalableDimensionDynamodbIndexWriteCapacityUnits            ScalableDimension = "dynamodb:index:WriteCapacityUnits"
	ScalableDimensionRdsClusterReadReplicaCount                 ScalableDimension = "rds:cluster:ReadReplicaCount"
	ScalableDimensionSagemakerVariantDesiredInstanceCount       ScalableDimension = "sagemaker:variant:DesiredInstanceCount"
	ScalableDimensionCustomResourceResourceTypeProperty         ScalableDimension = "custom-resource:ResourceType:Property"
)

func (enum ScalableDimension) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalableDimension) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScalingActivityStatusCode string

// Enum values for ScalingActivityStatusCode
const (
	ScalingActivityStatusCodePending     ScalingActivityStatusCode = "Pending"
	ScalingActivityStatusCodeInProgress  ScalingActivityStatusCode = "InProgress"
	ScalingActivityStatusCodeSuccessful  ScalingActivityStatusCode = "Successful"
	ScalingActivityStatusCodeOverridden  ScalingActivityStatusCode = "Overridden"
	ScalingActivityStatusCodeUnfulfilled ScalingActivityStatusCode = "Unfulfilled"
	ScalingActivityStatusCodeFailed      ScalingActivityStatusCode = "Failed"
)

func (enum ScalingActivityStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScalingActivityStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServiceNamespace string

// Enum values for ServiceNamespace
const (
	ServiceNamespaceEcs              ServiceNamespace = "ecs"
	ServiceNamespaceElasticmapreduce ServiceNamespace = "elasticmapreduce"
	ServiceNamespaceEc2              ServiceNamespace = "ec2"
	ServiceNamespaceAppstream        ServiceNamespace = "appstream"
	ServiceNamespaceDynamodb         ServiceNamespace = "dynamodb"
	ServiceNamespaceRds              ServiceNamespace = "rds"
	ServiceNamespaceSagemaker        ServiceNamespace = "sagemaker"
	ServiceNamespaceCustomResource   ServiceNamespace = "custom-resource"
)

func (enum ServiceNamespace) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceNamespace) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
