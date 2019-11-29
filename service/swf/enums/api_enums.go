// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ActivityTaskTimeoutType string

// Enum values for ActivityTaskTimeoutType
const (
	ActivityTaskTimeoutTypeStartToClose    ActivityTaskTimeoutType = "START_TO_CLOSE"
	ActivityTaskTimeoutTypeScheduleToStart ActivityTaskTimeoutType = "SCHEDULE_TO_START"
	ActivityTaskTimeoutTypeScheduleToClose ActivityTaskTimeoutType = "SCHEDULE_TO_CLOSE"
	ActivityTaskTimeoutTypeHeartbeat       ActivityTaskTimeoutType = "HEARTBEAT"
)

func (enum ActivityTaskTimeoutType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ActivityTaskTimeoutType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CancelTimerFailedCause string

// Enum values for CancelTimerFailedCause
const (
	CancelTimerFailedCauseTimerIdUnknown        CancelTimerFailedCause = "TIMER_ID_UNKNOWN"
	CancelTimerFailedCauseOperationNotPermitted CancelTimerFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum CancelTimerFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CancelTimerFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CancelWorkflowExecutionFailedCause string

// Enum values for CancelWorkflowExecutionFailedCause
const (
	CancelWorkflowExecutionFailedCauseUnhandledDecision     CancelWorkflowExecutionFailedCause = "UNHANDLED_DECISION"
	CancelWorkflowExecutionFailedCauseOperationNotPermitted CancelWorkflowExecutionFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum CancelWorkflowExecutionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CancelWorkflowExecutionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChildPolicy string

// Enum values for ChildPolicy
const (
	ChildPolicyTerminate     ChildPolicy = "TERMINATE"
	ChildPolicyRequestCancel ChildPolicy = "REQUEST_CANCEL"
	ChildPolicyAbandon       ChildPolicy = "ABANDON"
)

func (enum ChildPolicy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChildPolicy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CloseStatus string

// Enum values for CloseStatus
const (
	CloseStatusCompleted      CloseStatus = "COMPLETED"
	CloseStatusFailed         CloseStatus = "FAILED"
	CloseStatusCanceled       CloseStatus = "CANCELED"
	CloseStatusTerminated     CloseStatus = "TERMINATED"
	CloseStatusContinuedAsNew CloseStatus = "CONTINUED_AS_NEW"
	CloseStatusTimedOut       CloseStatus = "TIMED_OUT"
)

func (enum CloseStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CloseStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CompleteWorkflowExecutionFailedCause string

// Enum values for CompleteWorkflowExecutionFailedCause
const (
	CompleteWorkflowExecutionFailedCauseUnhandledDecision     CompleteWorkflowExecutionFailedCause = "UNHANDLED_DECISION"
	CompleteWorkflowExecutionFailedCauseOperationNotPermitted CompleteWorkflowExecutionFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum CompleteWorkflowExecutionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CompleteWorkflowExecutionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContinueAsNewWorkflowExecutionFailedCause string

// Enum values for ContinueAsNewWorkflowExecutionFailedCause
const (
	ContinueAsNewWorkflowExecutionFailedCauseUnhandledDecision                            ContinueAsNewWorkflowExecutionFailedCause = "UNHANDLED_DECISION"
	ContinueAsNewWorkflowExecutionFailedCauseWorkflowTypeDeprecated                       ContinueAsNewWorkflowExecutionFailedCause = "WORKFLOW_TYPE_DEPRECATED"
	ContinueAsNewWorkflowExecutionFailedCauseWorkflowTypeDoesNotExist                     ContinueAsNewWorkflowExecutionFailedCause = "WORKFLOW_TYPE_DOES_NOT_EXIST"
	ContinueAsNewWorkflowExecutionFailedCauseDefaultExecutionStartToCloseTimeoutUndefined ContinueAsNewWorkflowExecutionFailedCause = "DEFAULT_EXECUTION_START_TO_CLOSE_TIMEOUT_UNDEFINED"
	ContinueAsNewWorkflowExecutionFailedCauseDefaultTaskStartToCloseTimeoutUndefined      ContinueAsNewWorkflowExecutionFailedCause = "DEFAULT_TASK_START_TO_CLOSE_TIMEOUT_UNDEFINED"
	ContinueAsNewWorkflowExecutionFailedCauseDefaultTaskListUndefined                     ContinueAsNewWorkflowExecutionFailedCause = "DEFAULT_TASK_LIST_UNDEFINED"
	ContinueAsNewWorkflowExecutionFailedCauseDefaultChildPolicyUndefined                  ContinueAsNewWorkflowExecutionFailedCause = "DEFAULT_CHILD_POLICY_UNDEFINED"
	ContinueAsNewWorkflowExecutionFailedCauseContinueAsNewWorkflowExecutionRateExceeded   ContinueAsNewWorkflowExecutionFailedCause = "CONTINUE_AS_NEW_WORKFLOW_EXECUTION_RATE_EXCEEDED"
	ContinueAsNewWorkflowExecutionFailedCauseOperationNotPermitted                        ContinueAsNewWorkflowExecutionFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum ContinueAsNewWorkflowExecutionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContinueAsNewWorkflowExecutionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DecisionTaskTimeoutType string

// Enum values for DecisionTaskTimeoutType
const (
	DecisionTaskTimeoutTypeStartToClose DecisionTaskTimeoutType = "START_TO_CLOSE"
)

func (enum DecisionTaskTimeoutType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DecisionTaskTimeoutType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DecisionType string

// Enum values for DecisionType
const (
	DecisionTypeScheduleActivityTask                   DecisionType = "ScheduleActivityTask"
	DecisionTypeRequestCancelActivityTask              DecisionType = "RequestCancelActivityTask"
	DecisionTypeCompleteWorkflowExecution              DecisionType = "CompleteWorkflowExecution"
	DecisionTypeFailWorkflowExecution                  DecisionType = "FailWorkflowExecution"
	DecisionTypeCancelWorkflowExecution                DecisionType = "CancelWorkflowExecution"
	DecisionTypeContinueAsNewWorkflowExecution         DecisionType = "ContinueAsNewWorkflowExecution"
	DecisionTypeRecordMarker                           DecisionType = "RecordMarker"
	DecisionTypeStartTimer                             DecisionType = "StartTimer"
	DecisionTypeCancelTimer                            DecisionType = "CancelTimer"
	DecisionTypeSignalExternalWorkflowExecution        DecisionType = "SignalExternalWorkflowExecution"
	DecisionTypeRequestCancelExternalWorkflowExecution DecisionType = "RequestCancelExternalWorkflowExecution"
	DecisionTypeStartChildWorkflowExecution            DecisionType = "StartChildWorkflowExecution"
	DecisionTypeScheduleLambdaFunction                 DecisionType = "ScheduleLambdaFunction"
)

func (enum DecisionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DecisionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EventType string

// Enum values for EventType
const (
	EventTypeWorkflowExecutionStarted                        EventType = "WorkflowExecutionStarted"
	EventTypeWorkflowExecutionCancelRequested                EventType = "WorkflowExecutionCancelRequested"
	EventTypeWorkflowExecutionCompleted                      EventType = "WorkflowExecutionCompleted"
	EventTypeCompleteWorkflowExecutionFailed                 EventType = "CompleteWorkflowExecutionFailed"
	EventTypeWorkflowExecutionFailed                         EventType = "WorkflowExecutionFailed"
	EventTypeFailWorkflowExecutionFailed                     EventType = "FailWorkflowExecutionFailed"
	EventTypeWorkflowExecutionTimedOut                       EventType = "WorkflowExecutionTimedOut"
	EventTypeWorkflowExecutionCanceled                       EventType = "WorkflowExecutionCanceled"
	EventTypeCancelWorkflowExecutionFailed                   EventType = "CancelWorkflowExecutionFailed"
	EventTypeWorkflowExecutionContinuedAsNew                 EventType = "WorkflowExecutionContinuedAsNew"
	EventTypeContinueAsNewWorkflowExecutionFailed            EventType = "ContinueAsNewWorkflowExecutionFailed"
	EventTypeWorkflowExecutionTerminated                     EventType = "WorkflowExecutionTerminated"
	EventTypeDecisionTaskScheduled                           EventType = "DecisionTaskScheduled"
	EventTypeDecisionTaskStarted                             EventType = "DecisionTaskStarted"
	EventTypeDecisionTaskCompleted                           EventType = "DecisionTaskCompleted"
	EventTypeDecisionTaskTimedOut                            EventType = "DecisionTaskTimedOut"
	EventTypeActivityTaskScheduled                           EventType = "ActivityTaskScheduled"
	EventTypeScheduleActivityTaskFailed                      EventType = "ScheduleActivityTaskFailed"
	EventTypeActivityTaskStarted                             EventType = "ActivityTaskStarted"
	EventTypeActivityTaskCompleted                           EventType = "ActivityTaskCompleted"
	EventTypeActivityTaskFailed                              EventType = "ActivityTaskFailed"
	EventTypeActivityTaskTimedOut                            EventType = "ActivityTaskTimedOut"
	EventTypeActivityTaskCanceled                            EventType = "ActivityTaskCanceled"
	EventTypeActivityTaskCancelRequested                     EventType = "ActivityTaskCancelRequested"
	EventTypeRequestCancelActivityTaskFailed                 EventType = "RequestCancelActivityTaskFailed"
	EventTypeWorkflowExecutionSignaled                       EventType = "WorkflowExecutionSignaled"
	EventTypeMarkerRecorded                                  EventType = "MarkerRecorded"
	EventTypeRecordMarkerFailed                              EventType = "RecordMarkerFailed"
	EventTypeTimerStarted                                    EventType = "TimerStarted"
	EventTypeStartTimerFailed                                EventType = "StartTimerFailed"
	EventTypeTimerFired                                      EventType = "TimerFired"
	EventTypeTimerCanceled                                   EventType = "TimerCanceled"
	EventTypeCancelTimerFailed                               EventType = "CancelTimerFailed"
	EventTypeStartChildWorkflowExecutionInitiated            EventType = "StartChildWorkflowExecutionInitiated"
	EventTypeStartChildWorkflowExecutionFailed               EventType = "StartChildWorkflowExecutionFailed"
	EventTypeChildWorkflowExecutionStarted                   EventType = "ChildWorkflowExecutionStarted"
	EventTypeChildWorkflowExecutionCompleted                 EventType = "ChildWorkflowExecutionCompleted"
	EventTypeChildWorkflowExecutionFailed                    EventType = "ChildWorkflowExecutionFailed"
	EventTypeChildWorkflowExecutionTimedOut                  EventType = "ChildWorkflowExecutionTimedOut"
	EventTypeChildWorkflowExecutionCanceled                  EventType = "ChildWorkflowExecutionCanceled"
	EventTypeChildWorkflowExecutionTerminated                EventType = "ChildWorkflowExecutionTerminated"
	EventTypeSignalExternalWorkflowExecutionInitiated        EventType = "SignalExternalWorkflowExecutionInitiated"
	EventTypeSignalExternalWorkflowExecutionFailed           EventType = "SignalExternalWorkflowExecutionFailed"
	EventTypeExternalWorkflowExecutionSignaled               EventType = "ExternalWorkflowExecutionSignaled"
	EventTypeRequestCancelExternalWorkflowExecutionInitiated EventType = "RequestCancelExternalWorkflowExecutionInitiated"
	EventTypeRequestCancelExternalWorkflowExecutionFailed    EventType = "RequestCancelExternalWorkflowExecutionFailed"
	EventTypeExternalWorkflowExecutionCancelRequested        EventType = "ExternalWorkflowExecutionCancelRequested"
	EventTypeLambdaFunctionScheduled                         EventType = "LambdaFunctionScheduled"
	EventTypeLambdaFunctionStarted                           EventType = "LambdaFunctionStarted"
	EventTypeLambdaFunctionCompleted                         EventType = "LambdaFunctionCompleted"
	EventTypeLambdaFunctionFailed                            EventType = "LambdaFunctionFailed"
	EventTypeLambdaFunctionTimedOut                          EventType = "LambdaFunctionTimedOut"
	EventTypeScheduleLambdaFunctionFailed                    EventType = "ScheduleLambdaFunctionFailed"
	EventTypeStartLambdaFunctionFailed                       EventType = "StartLambdaFunctionFailed"
)

func (enum EventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExecutionStatus string

// Enum values for ExecutionStatus
const (
	ExecutionStatusOpen   ExecutionStatus = "OPEN"
	ExecutionStatusClosed ExecutionStatus = "CLOSED"
)

func (enum ExecutionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExecutionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FailWorkflowExecutionFailedCause string

// Enum values for FailWorkflowExecutionFailedCause
const (
	FailWorkflowExecutionFailedCauseUnhandledDecision     FailWorkflowExecutionFailedCause = "UNHANDLED_DECISION"
	FailWorkflowExecutionFailedCauseOperationNotPermitted FailWorkflowExecutionFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum FailWorkflowExecutionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FailWorkflowExecutionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LambdaFunctionTimeoutType string

// Enum values for LambdaFunctionTimeoutType
const (
	LambdaFunctionTimeoutTypeStartToClose LambdaFunctionTimeoutType = "START_TO_CLOSE"
)

func (enum LambdaFunctionTimeoutType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LambdaFunctionTimeoutType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecordMarkerFailedCause string

// Enum values for RecordMarkerFailedCause
const (
	RecordMarkerFailedCauseOperationNotPermitted RecordMarkerFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum RecordMarkerFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecordMarkerFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RegistrationStatus string

// Enum values for RegistrationStatus
const (
	RegistrationStatusRegistered RegistrationStatus = "REGISTERED"
	RegistrationStatusDeprecated RegistrationStatus = "DEPRECATED"
)

func (enum RegistrationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RegistrationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RequestCancelActivityTaskFailedCause string

// Enum values for RequestCancelActivityTaskFailedCause
const (
	RequestCancelActivityTaskFailedCauseActivityIdUnknown     RequestCancelActivityTaskFailedCause = "ACTIVITY_ID_UNKNOWN"
	RequestCancelActivityTaskFailedCauseOperationNotPermitted RequestCancelActivityTaskFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum RequestCancelActivityTaskFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RequestCancelActivityTaskFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RequestCancelExternalWorkflowExecutionFailedCause string

// Enum values for RequestCancelExternalWorkflowExecutionFailedCause
const (
	RequestCancelExternalWorkflowExecutionFailedCauseUnknownExternalWorkflowExecution                   RequestCancelExternalWorkflowExecutionFailedCause = "UNKNOWN_EXTERNAL_WORKFLOW_EXECUTION"
	RequestCancelExternalWorkflowExecutionFailedCauseRequestCancelExternalWorkflowExecutionRateExceeded RequestCancelExternalWorkflowExecutionFailedCause = "REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_RATE_EXCEEDED"
	RequestCancelExternalWorkflowExecutionFailedCauseOperationNotPermitted                              RequestCancelExternalWorkflowExecutionFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum RequestCancelExternalWorkflowExecutionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RequestCancelExternalWorkflowExecutionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScheduleActivityTaskFailedCause string

// Enum values for ScheduleActivityTaskFailedCause
const (
	ScheduleActivityTaskFailedCauseActivityTypeDeprecated                 ScheduleActivityTaskFailedCause = "ACTIVITY_TYPE_DEPRECATED"
	ScheduleActivityTaskFailedCauseActivityTypeDoesNotExist               ScheduleActivityTaskFailedCause = "ACTIVITY_TYPE_DOES_NOT_EXIST"
	ScheduleActivityTaskFailedCauseActivityIdAlreadyInUse                 ScheduleActivityTaskFailedCause = "ACTIVITY_ID_ALREADY_IN_USE"
	ScheduleActivityTaskFailedCauseOpenActivitiesLimitExceeded            ScheduleActivityTaskFailedCause = "OPEN_ACTIVITIES_LIMIT_EXCEEDED"
	ScheduleActivityTaskFailedCauseActivityCreationRateExceeded           ScheduleActivityTaskFailedCause = "ACTIVITY_CREATION_RATE_EXCEEDED"
	ScheduleActivityTaskFailedCauseDefaultScheduleToCloseTimeoutUndefined ScheduleActivityTaskFailedCause = "DEFAULT_SCHEDULE_TO_CLOSE_TIMEOUT_UNDEFINED"
	ScheduleActivityTaskFailedCauseDefaultTaskListUndefined               ScheduleActivityTaskFailedCause = "DEFAULT_TASK_LIST_UNDEFINED"
	ScheduleActivityTaskFailedCauseDefaultScheduleToStartTimeoutUndefined ScheduleActivityTaskFailedCause = "DEFAULT_SCHEDULE_TO_START_TIMEOUT_UNDEFINED"
	ScheduleActivityTaskFailedCauseDefaultStartToCloseTimeoutUndefined    ScheduleActivityTaskFailedCause = "DEFAULT_START_TO_CLOSE_TIMEOUT_UNDEFINED"
	ScheduleActivityTaskFailedCauseDefaultHeartbeatTimeoutUndefined       ScheduleActivityTaskFailedCause = "DEFAULT_HEARTBEAT_TIMEOUT_UNDEFINED"
	ScheduleActivityTaskFailedCauseOperationNotPermitted                  ScheduleActivityTaskFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum ScheduleActivityTaskFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScheduleActivityTaskFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ScheduleLambdaFunctionFailedCause string

// Enum values for ScheduleLambdaFunctionFailedCause
const (
	ScheduleLambdaFunctionFailedCauseIdAlreadyInUse                     ScheduleLambdaFunctionFailedCause = "ID_ALREADY_IN_USE"
	ScheduleLambdaFunctionFailedCauseOpenLambdaFunctionsLimitExceeded   ScheduleLambdaFunctionFailedCause = "OPEN_LAMBDA_FUNCTIONS_LIMIT_EXCEEDED"
	ScheduleLambdaFunctionFailedCauseLambdaFunctionCreationRateExceeded ScheduleLambdaFunctionFailedCause = "LAMBDA_FUNCTION_CREATION_RATE_EXCEEDED"
	ScheduleLambdaFunctionFailedCauseLambdaServiceNotAvailableInRegion  ScheduleLambdaFunctionFailedCause = "LAMBDA_SERVICE_NOT_AVAILABLE_IN_REGION"
)

func (enum ScheduleLambdaFunctionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ScheduleLambdaFunctionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SignalExternalWorkflowExecutionFailedCause string

// Enum values for SignalExternalWorkflowExecutionFailedCause
const (
	SignalExternalWorkflowExecutionFailedCauseUnknownExternalWorkflowExecution            SignalExternalWorkflowExecutionFailedCause = "UNKNOWN_EXTERNAL_WORKFLOW_EXECUTION"
	SignalExternalWorkflowExecutionFailedCauseSignalExternalWorkflowExecutionRateExceeded SignalExternalWorkflowExecutionFailedCause = "SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_RATE_EXCEEDED"
	SignalExternalWorkflowExecutionFailedCauseOperationNotPermitted                       SignalExternalWorkflowExecutionFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum SignalExternalWorkflowExecutionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SignalExternalWorkflowExecutionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StartChildWorkflowExecutionFailedCause string

// Enum values for StartChildWorkflowExecutionFailedCause
const (
	StartChildWorkflowExecutionFailedCauseWorkflowTypeDoesNotExist                     StartChildWorkflowExecutionFailedCause = "WORKFLOW_TYPE_DOES_NOT_EXIST"
	StartChildWorkflowExecutionFailedCauseWorkflowTypeDeprecated                       StartChildWorkflowExecutionFailedCause = "WORKFLOW_TYPE_DEPRECATED"
	StartChildWorkflowExecutionFailedCauseOpenChildrenLimitExceeded                    StartChildWorkflowExecutionFailedCause = "OPEN_CHILDREN_LIMIT_EXCEEDED"
	StartChildWorkflowExecutionFailedCauseOpenWorkflowsLimitExceeded                   StartChildWorkflowExecutionFailedCause = "OPEN_WORKFLOWS_LIMIT_EXCEEDED"
	StartChildWorkflowExecutionFailedCauseChildCreationRateExceeded                    StartChildWorkflowExecutionFailedCause = "CHILD_CREATION_RATE_EXCEEDED"
	StartChildWorkflowExecutionFailedCauseWorkflowAlreadyRunning                       StartChildWorkflowExecutionFailedCause = "WORKFLOW_ALREADY_RUNNING"
	StartChildWorkflowExecutionFailedCauseDefaultExecutionStartToCloseTimeoutUndefined StartChildWorkflowExecutionFailedCause = "DEFAULT_EXECUTION_START_TO_CLOSE_TIMEOUT_UNDEFINED"
	StartChildWorkflowExecutionFailedCauseDefaultTaskListUndefined                     StartChildWorkflowExecutionFailedCause = "DEFAULT_TASK_LIST_UNDEFINED"
	StartChildWorkflowExecutionFailedCauseDefaultTaskStartToCloseTimeoutUndefined      StartChildWorkflowExecutionFailedCause = "DEFAULT_TASK_START_TO_CLOSE_TIMEOUT_UNDEFINED"
	StartChildWorkflowExecutionFailedCauseDefaultChildPolicyUndefined                  StartChildWorkflowExecutionFailedCause = "DEFAULT_CHILD_POLICY_UNDEFINED"
	StartChildWorkflowExecutionFailedCauseOperationNotPermitted                        StartChildWorkflowExecutionFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum StartChildWorkflowExecutionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StartChildWorkflowExecutionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StartLambdaFunctionFailedCause string

// Enum values for StartLambdaFunctionFailedCause
const (
	StartLambdaFunctionFailedCauseAssumeRoleFailed StartLambdaFunctionFailedCause = "ASSUME_ROLE_FAILED"
)

func (enum StartLambdaFunctionFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StartLambdaFunctionFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StartTimerFailedCause string

// Enum values for StartTimerFailedCause
const (
	StartTimerFailedCauseTimerIdAlreadyInUse       StartTimerFailedCause = "TIMER_ID_ALREADY_IN_USE"
	StartTimerFailedCauseOpenTimersLimitExceeded   StartTimerFailedCause = "OPEN_TIMERS_LIMIT_EXCEEDED"
	StartTimerFailedCauseTimerCreationRateExceeded StartTimerFailedCause = "TIMER_CREATION_RATE_EXCEEDED"
	StartTimerFailedCauseOperationNotPermitted     StartTimerFailedCause = "OPERATION_NOT_PERMITTED"
)

func (enum StartTimerFailedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StartTimerFailedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkflowExecutionCancelRequestedCause string

// Enum values for WorkflowExecutionCancelRequestedCause
const (
	WorkflowExecutionCancelRequestedCauseChildPolicyApplied WorkflowExecutionCancelRequestedCause = "CHILD_POLICY_APPLIED"
)

func (enum WorkflowExecutionCancelRequestedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkflowExecutionCancelRequestedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkflowExecutionTerminatedCause string

// Enum values for WorkflowExecutionTerminatedCause
const (
	WorkflowExecutionTerminatedCauseChildPolicyApplied WorkflowExecutionTerminatedCause = "CHILD_POLICY_APPLIED"
	WorkflowExecutionTerminatedCauseEventLimitExceeded WorkflowExecutionTerminatedCause = "EVENT_LIMIT_EXCEEDED"
	WorkflowExecutionTerminatedCauseOperatorInitiated  WorkflowExecutionTerminatedCause = "OPERATOR_INITIATED"
)

func (enum WorkflowExecutionTerminatedCause) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkflowExecutionTerminatedCause) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WorkflowExecutionTimeoutType string

// Enum values for WorkflowExecutionTimeoutType
const (
	WorkflowExecutionTimeoutTypeStartToClose WorkflowExecutionTimeoutType = "START_TO_CLOSE"
)

func (enum WorkflowExecutionTimeoutType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WorkflowExecutionTimeoutType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}