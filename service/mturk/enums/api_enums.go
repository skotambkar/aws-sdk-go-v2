// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type AssignmentStatus string

// Enum values for AssignmentStatus
const (
	AssignmentStatusSubmitted AssignmentStatus = "Submitted"
	AssignmentStatusApproved  AssignmentStatus = "Approved"
	AssignmentStatusRejected  AssignmentStatus = "Rejected"
)

func (enum AssignmentStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AssignmentStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Comparator string

// Enum values for Comparator
const (
	ComparatorLessThan             Comparator = "LessThan"
	ComparatorLessThanOrEqualTo    Comparator = "LessThanOrEqualTo"
	ComparatorGreaterThan          Comparator = "GreaterThan"
	ComparatorGreaterThanOrEqualTo Comparator = "GreaterThanOrEqualTo"
	ComparatorEqualTo              Comparator = "EqualTo"
	ComparatorNotEqualTo           Comparator = "NotEqualTo"
	ComparatorExists               Comparator = "Exists"
	ComparatorDoesNotExist         Comparator = "DoesNotExist"
	ComparatorIn                   Comparator = "In"
	ComparatorNotIn                Comparator = "NotIn"
)

func (enum Comparator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Comparator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EventType string

// Enum values for EventType
const (
	EventTypeAssignmentAccepted  EventType = "AssignmentAccepted"
	EventTypeAssignmentAbandoned EventType = "AssignmentAbandoned"
	EventTypeAssignmentReturned  EventType = "AssignmentReturned"
	EventTypeAssignmentSubmitted EventType = "AssignmentSubmitted"
	EventTypeAssignmentRejected  EventType = "AssignmentRejected"
	EventTypeAssignmentApproved  EventType = "AssignmentApproved"
	EventTypeHitcreated          EventType = "HITCreated"
	EventTypeHitexpired          EventType = "HITExpired"
	EventTypeHitreviewable       EventType = "HITReviewable"
	EventTypeHitextended         EventType = "HITExtended"
	EventTypeHitdisposed         EventType = "HITDisposed"
	EventTypePing                EventType = "Ping"
)

func (enum EventType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EventType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HITAccessActions string

// Enum values for HITAccessActions
const (
	HITAccessActionsAccept                   HITAccessActions = "Accept"
	HITAccessActionsPreviewAndAccept         HITAccessActions = "PreviewAndAccept"
	HITAccessActionsDiscoverPreviewAndAccept HITAccessActions = "DiscoverPreviewAndAccept"
)

func (enum HITAccessActions) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HITAccessActions) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HITReviewStatus string

// Enum values for HITReviewStatus
const (
	HITReviewStatusNotReviewed           HITReviewStatus = "NotReviewed"
	HITReviewStatusMarkedForReview       HITReviewStatus = "MarkedForReview"
	HITReviewStatusReviewedAppropriate   HITReviewStatus = "ReviewedAppropriate"
	HITReviewStatusReviewedInappropriate HITReviewStatus = "ReviewedInappropriate"
)

func (enum HITReviewStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HITReviewStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HITStatus string

// Enum values for HITStatus
const (
	HITStatusAssignable   HITStatus = "Assignable"
	HITStatusUnassignable HITStatus = "Unassignable"
	HITStatusReviewable   HITStatus = "Reviewable"
	HITStatusReviewing    HITStatus = "Reviewing"
	HITStatusDisposed     HITStatus = "Disposed"
)

func (enum HITStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HITStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NotificationTransport string

// Enum values for NotificationTransport
const (
	NotificationTransportEmail NotificationTransport = "Email"
	NotificationTransportSqs   NotificationTransport = "SQS"
	NotificationTransportSns   NotificationTransport = "SNS"
)

func (enum NotificationTransport) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NotificationTransport) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NotifyWorkersFailureCode string

// Enum values for NotifyWorkersFailureCode
const (
	NotifyWorkersFailureCodeSoftFailure NotifyWorkersFailureCode = "SoftFailure"
	NotifyWorkersFailureCodeHardFailure NotifyWorkersFailureCode = "HardFailure"
)

func (enum NotifyWorkersFailureCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NotifyWorkersFailureCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type QualificationStatus string

// Enum values for QualificationStatus
const (
	QualificationStatusGranted QualificationStatus = "Granted"
	QualificationStatusRevoked QualificationStatus = "Revoked"
)

func (enum QualificationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum QualificationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type QualificationTypeStatus string

// Enum values for QualificationTypeStatus
const (
	QualificationTypeStatusActive   QualificationTypeStatus = "Active"
	QualificationTypeStatusInactive QualificationTypeStatus = "Inactive"
)

func (enum QualificationTypeStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum QualificationTypeStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReviewActionStatus string

// Enum values for ReviewActionStatus
const (
	ReviewActionStatusIntended  ReviewActionStatus = "Intended"
	ReviewActionStatusSucceeded ReviewActionStatus = "Succeeded"
	ReviewActionStatusFailed    ReviewActionStatus = "Failed"
	ReviewActionStatusCancelled ReviewActionStatus = "Cancelled"
)

func (enum ReviewActionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReviewActionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReviewPolicyLevel string

// Enum values for ReviewPolicyLevel
const (
	ReviewPolicyLevelAssignment ReviewPolicyLevel = "Assignment"
	ReviewPolicyLevelHit        ReviewPolicyLevel = "HIT"
)

func (enum ReviewPolicyLevel) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReviewPolicyLevel) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReviewableHITStatus string

// Enum values for ReviewableHITStatus
const (
	ReviewableHITStatusReviewable ReviewableHITStatus = "Reviewable"
	ReviewableHITStatusReviewing  ReviewableHITStatus = "Reviewing"
)

func (enum ReviewableHITStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReviewableHITStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
