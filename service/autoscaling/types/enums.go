// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type InstanceMetadataEndpointState string

// Enum values for InstanceMetadataEndpointState
const (
	InstanceMetadataEndpointStateDisabled InstanceMetadataEndpointState = "disabled"
	InstanceMetadataEndpointStateEnabled  InstanceMetadataEndpointState = "enabled"
)

// Values returns all known values for InstanceMetadataEndpointState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (InstanceMetadataEndpointState) Values() []InstanceMetadataEndpointState {
	return []InstanceMetadataEndpointState{
		"disabled",
		"enabled",
	}
}

type InstanceMetadataHttpTokensState string

// Enum values for InstanceMetadataHttpTokensState
const (
	InstanceMetadataHttpTokensStateOptional InstanceMetadataHttpTokensState = "optional"
	InstanceMetadataHttpTokensStateRequired InstanceMetadataHttpTokensState = "required"
)

// Values returns all known values for InstanceMetadataHttpTokensState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (InstanceMetadataHttpTokensState) Values() []InstanceMetadataHttpTokensState {
	return []InstanceMetadataHttpTokensState{
		"optional",
		"required",
	}
}

type InstanceRefreshStatus string

// Enum values for InstanceRefreshStatus
const (
	InstanceRefreshStatusPending    InstanceRefreshStatus = "Pending"
	InstanceRefreshStatusInProgress InstanceRefreshStatus = "InProgress"
	InstanceRefreshStatusSuccessful InstanceRefreshStatus = "Successful"
	InstanceRefreshStatusFailed     InstanceRefreshStatus = "Failed"
	InstanceRefreshStatusCancelling InstanceRefreshStatus = "Cancelling"
	InstanceRefreshStatusCancelled  InstanceRefreshStatus = "Cancelled"
)

// Values returns all known values for InstanceRefreshStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceRefreshStatus) Values() []InstanceRefreshStatus {
	return []InstanceRefreshStatus{
		"Pending",
		"InProgress",
		"Successful",
		"Failed",
		"Cancelling",
		"Cancelled",
	}
}

type LifecycleState string

// Enum values for LifecycleState
const (
	LifecycleStatePending                  LifecycleState = "Pending"
	LifecycleStatePendingWait              LifecycleState = "Pending:Wait"
	LifecycleStatePendingProceed           LifecycleState = "Pending:Proceed"
	LifecycleStateQuarantined              LifecycleState = "Quarantined"
	LifecycleStateInService                LifecycleState = "InService"
	LifecycleStateTerminating              LifecycleState = "Terminating"
	LifecycleStateTerminatingWait          LifecycleState = "Terminating:Wait"
	LifecycleStateTerminatingProceed       LifecycleState = "Terminating:Proceed"
	LifecycleStateTerminated               LifecycleState = "Terminated"
	LifecycleStateDetaching                LifecycleState = "Detaching"
	LifecycleStateDetached                 LifecycleState = "Detached"
	LifecycleStateEnteringStandby          LifecycleState = "EnteringStandby"
	LifecycleStateStandby                  LifecycleState = "Standby"
	LifecycleStateWarmedPending            LifecycleState = "Warmed:Pending"
	LifecycleStateWarmedPendingWait        LifecycleState = "Warmed:Pending:Wait"
	LifecycleStateWarmedPendingProceed     LifecycleState = "Warmed:Pending:Proceed"
	LifecycleStateWarmedTerminating        LifecycleState = "Warmed:Terminating"
	LifecycleStateWarmedTerminatingWait    LifecycleState = "Warmed:Terminating:Wait"
	LifecycleStateWarmedTerminatingProceed LifecycleState = "Warmed:Terminating:Proceed"
	LifecycleStateWarmedTerminated         LifecycleState = "Warmed:Terminated"
	LifecycleStateWarmedStopped            LifecycleState = "Warmed:Stopped"
	LifecycleStateWarmedRunning            LifecycleState = "Warmed:Running"
)

// Values returns all known values for LifecycleState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LifecycleState) Values() []LifecycleState {
	return []LifecycleState{
		"Pending",
		"Pending:Wait",
		"Pending:Proceed",
		"Quarantined",
		"InService",
		"Terminating",
		"Terminating:Wait",
		"Terminating:Proceed",
		"Terminated",
		"Detaching",
		"Detached",
		"EnteringStandby",
		"Standby",
		"Warmed:Pending",
		"Warmed:Pending:Wait",
		"Warmed:Pending:Proceed",
		"Warmed:Terminating",
		"Warmed:Terminating:Wait",
		"Warmed:Terminating:Proceed",
		"Warmed:Terminated",
		"Warmed:Stopped",
		"Warmed:Running",
	}
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

// Values returns all known values for MetricStatistic. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MetricStatistic) Values() []MetricStatistic {
	return []MetricStatistic{
		"Average",
		"Minimum",
		"Maximum",
		"SampleCount",
		"Sum",
	}
}

type MetricType string

// Enum values for MetricType
const (
	MetricTypeASGAverageCPUUtilization MetricType = "ASGAverageCPUUtilization"
	MetricTypeASGAverageNetworkIn      MetricType = "ASGAverageNetworkIn"
	MetricTypeASGAverageNetworkOut     MetricType = "ASGAverageNetworkOut"
	MetricTypeALBRequestCountPerTarget MetricType = "ALBRequestCountPerTarget"
)

// Values returns all known values for MetricType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MetricType) Values() []MetricType {
	return []MetricType{
		"ASGAverageCPUUtilization",
		"ASGAverageNetworkIn",
		"ASGAverageNetworkOut",
		"ALBRequestCountPerTarget",
	}
}

type PredefinedLoadMetricType string

// Enum values for PredefinedLoadMetricType
const (
	PredefinedLoadMetricTypeASGTotalCPUUtilization     PredefinedLoadMetricType = "ASGTotalCPUUtilization"
	PredefinedLoadMetricTypeASGTotalNetworkIn          PredefinedLoadMetricType = "ASGTotalNetworkIn"
	PredefinedLoadMetricTypeASGTotalNetworkOut         PredefinedLoadMetricType = "ASGTotalNetworkOut"
	PredefinedLoadMetricTypeALBTargetGroupRequestCount PredefinedLoadMetricType = "ALBTargetGroupRequestCount"
)

// Values returns all known values for PredefinedLoadMetricType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PredefinedLoadMetricType) Values() []PredefinedLoadMetricType {
	return []PredefinedLoadMetricType{
		"ASGTotalCPUUtilization",
		"ASGTotalNetworkIn",
		"ASGTotalNetworkOut",
		"ALBTargetGroupRequestCount",
	}
}

type PredefinedMetricPairType string

// Enum values for PredefinedMetricPairType
const (
	PredefinedMetricPairTypeASGCPUUtilization PredefinedMetricPairType = "ASGCPUUtilization"
	PredefinedMetricPairTypeASGNetworkIn      PredefinedMetricPairType = "ASGNetworkIn"
	PredefinedMetricPairTypeASGNetworkOut     PredefinedMetricPairType = "ASGNetworkOut"
	PredefinedMetricPairTypeALBRequestCount   PredefinedMetricPairType = "ALBRequestCount"
)

// Values returns all known values for PredefinedMetricPairType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PredefinedMetricPairType) Values() []PredefinedMetricPairType {
	return []PredefinedMetricPairType{
		"ASGCPUUtilization",
		"ASGNetworkIn",
		"ASGNetworkOut",
		"ALBRequestCount",
	}
}

type PredefinedScalingMetricType string

// Enum values for PredefinedScalingMetricType
const (
	PredefinedScalingMetricTypeASGAverageCPUUtilization PredefinedScalingMetricType = "ASGAverageCPUUtilization"
	PredefinedScalingMetricTypeASGAverageNetworkIn      PredefinedScalingMetricType = "ASGAverageNetworkIn"
	PredefinedScalingMetricTypeASGAverageNetworkOut     PredefinedScalingMetricType = "ASGAverageNetworkOut"
	PredefinedScalingMetricTypeALBRequestCountPerTarget PredefinedScalingMetricType = "ALBRequestCountPerTarget"
)

// Values returns all known values for PredefinedScalingMetricType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PredefinedScalingMetricType) Values() []PredefinedScalingMetricType {
	return []PredefinedScalingMetricType{
		"ASGAverageCPUUtilization",
		"ASGAverageNetworkIn",
		"ASGAverageNetworkOut",
		"ALBRequestCountPerTarget",
	}
}

type PredictiveScalingMaxCapacityBreachBehavior string

// Enum values for PredictiveScalingMaxCapacityBreachBehavior
const (
	PredictiveScalingMaxCapacityBreachBehaviorHonorMaxCapacity    PredictiveScalingMaxCapacityBreachBehavior = "HonorMaxCapacity"
	PredictiveScalingMaxCapacityBreachBehaviorIncreaseMaxCapacity PredictiveScalingMaxCapacityBreachBehavior = "IncreaseMaxCapacity"
)

// Values returns all known values for PredictiveScalingMaxCapacityBreachBehavior.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (PredictiveScalingMaxCapacityBreachBehavior) Values() []PredictiveScalingMaxCapacityBreachBehavior {
	return []PredictiveScalingMaxCapacityBreachBehavior{
		"HonorMaxCapacity",
		"IncreaseMaxCapacity",
	}
}

type PredictiveScalingMode string

// Enum values for PredictiveScalingMode
const (
	PredictiveScalingModeForecastAndScale PredictiveScalingMode = "ForecastAndScale"
	PredictiveScalingModeForecastOnly     PredictiveScalingMode = "ForecastOnly"
)

// Values returns all known values for PredictiveScalingMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PredictiveScalingMode) Values() []PredictiveScalingMode {
	return []PredictiveScalingMode{
		"ForecastAndScale",
		"ForecastOnly",
	}
}

type RefreshStrategy string

// Enum values for RefreshStrategy
const (
	RefreshStrategyRolling RefreshStrategy = "Rolling"
)

// Values returns all known values for RefreshStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RefreshStrategy) Values() []RefreshStrategy {
	return []RefreshStrategy{
		"Rolling",
	}
}

type ScalingActivityStatusCode string

// Enum values for ScalingActivityStatusCode
const (
	ScalingActivityStatusCodePendingSpotBidPlacement         ScalingActivityStatusCode = "PendingSpotBidPlacement"
	ScalingActivityStatusCodeWaitingForSpotInstanceRequestId ScalingActivityStatusCode = "WaitingForSpotInstanceRequestId"
	ScalingActivityStatusCodeWaitingForSpotInstanceId        ScalingActivityStatusCode = "WaitingForSpotInstanceId"
	ScalingActivityStatusCodeWaitingForInstanceId            ScalingActivityStatusCode = "WaitingForInstanceId"
	ScalingActivityStatusCodePreInService                    ScalingActivityStatusCode = "PreInService"
	ScalingActivityStatusCodeInProgress                      ScalingActivityStatusCode = "InProgress"
	ScalingActivityStatusCodeWaitingForELBConnectionDraining ScalingActivityStatusCode = "WaitingForELBConnectionDraining"
	ScalingActivityStatusCodeMidLifecycleAction              ScalingActivityStatusCode = "MidLifecycleAction"
	ScalingActivityStatusCodeWaitingForInstanceWarmup        ScalingActivityStatusCode = "WaitingForInstanceWarmup"
	ScalingActivityStatusCodeSuccessful                      ScalingActivityStatusCode = "Successful"
	ScalingActivityStatusCodeFailed                          ScalingActivityStatusCode = "Failed"
	ScalingActivityStatusCodeCancelled                       ScalingActivityStatusCode = "Cancelled"
)

// Values returns all known values for ScalingActivityStatusCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScalingActivityStatusCode) Values() []ScalingActivityStatusCode {
	return []ScalingActivityStatusCode{
		"PendingSpotBidPlacement",
		"WaitingForSpotInstanceRequestId",
		"WaitingForSpotInstanceId",
		"WaitingForInstanceId",
		"PreInService",
		"InProgress",
		"WaitingForELBConnectionDraining",
		"MidLifecycleAction",
		"WaitingForInstanceWarmup",
		"Successful",
		"Failed",
		"Cancelled",
	}
}

type WarmPoolState string

// Enum values for WarmPoolState
const (
	WarmPoolStateStopped WarmPoolState = "Stopped"
	WarmPoolStateRunning WarmPoolState = "Running"
)

// Values returns all known values for WarmPoolState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WarmPoolState) Values() []WarmPoolState {
	return []WarmPoolState{
		"Stopped",
		"Running",
	}
}

type WarmPoolStatus string

// Enum values for WarmPoolStatus
const (
	WarmPoolStatusPendingDelete WarmPoolStatus = "PendingDelete"
)

// Values returns all known values for WarmPoolStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WarmPoolStatus) Values() []WarmPoolStatus {
	return []WarmPoolStatus{
		"PendingDelete",
	}
}
