// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type Edition string

// Enum values for Edition
const (
	EditionStarter  Edition = "STARTER"
	EditionStandard Edition = "STANDARD"
)

func (enum Edition) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Edition) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Framework string

// Enum values for Framework
const (
	FrameworkHyperledgerFabric Framework = "HYPERLEDGER_FABRIC"
)

func (enum Framework) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Framework) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InvitationStatus string

// Enum values for InvitationStatus
const (
	InvitationStatusPending   InvitationStatus = "PENDING"
	InvitationStatusAccepted  InvitationStatus = "ACCEPTED"
	InvitationStatusAccepting InvitationStatus = "ACCEPTING"
	InvitationStatusRejected  InvitationStatus = "REJECTED"
	InvitationStatusExpired   InvitationStatus = "EXPIRED"
)

func (enum InvitationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InvitationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MemberStatus string

// Enum values for MemberStatus
const (
	MemberStatusCreating     MemberStatus = "CREATING"
	MemberStatusAvailable    MemberStatus = "AVAILABLE"
	MemberStatusCreateFailed MemberStatus = "CREATE_FAILED"
	MemberStatusDeleting     MemberStatus = "DELETING"
	MemberStatusDeleted      MemberStatus = "DELETED"
)

func (enum MemberStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MemberStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NetworkStatus string

// Enum values for NetworkStatus
const (
	NetworkStatusCreating     NetworkStatus = "CREATING"
	NetworkStatusAvailable    NetworkStatus = "AVAILABLE"
	NetworkStatusCreateFailed NetworkStatus = "CREATE_FAILED"
	NetworkStatusDeleting     NetworkStatus = "DELETING"
	NetworkStatusDeleted      NetworkStatus = "DELETED"
)

func (enum NetworkStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NetworkStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NodeStatus string

// Enum values for NodeStatus
const (
	NodeStatusCreating     NodeStatus = "CREATING"
	NodeStatusAvailable    NodeStatus = "AVAILABLE"
	NodeStatusCreateFailed NodeStatus = "CREATE_FAILED"
	NodeStatusDeleting     NodeStatus = "DELETING"
	NodeStatusDeleted      NodeStatus = "DELETED"
	NodeStatusFailed       NodeStatus = "FAILED"
)

func (enum NodeStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NodeStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProposalStatus string

// Enum values for ProposalStatus
const (
	ProposalStatusInProgress   ProposalStatus = "IN_PROGRESS"
	ProposalStatusApproved     ProposalStatus = "APPROVED"
	ProposalStatusRejected     ProposalStatus = "REJECTED"
	ProposalStatusExpired      ProposalStatus = "EXPIRED"
	ProposalStatusActionFailed ProposalStatus = "ACTION_FAILED"
)

func (enum ProposalStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProposalStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ThresholdComparator string

// Enum values for ThresholdComparator
const (
	ThresholdComparatorGreaterThan          ThresholdComparator = "GREATER_THAN"
	ThresholdComparatorGreaterThanOrEqualTo ThresholdComparator = "GREATER_THAN_OR_EQUAL_TO"
)

func (enum ThresholdComparator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ThresholdComparator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type VoteValue string

// Enum values for VoteValue
const (
	VoteValueYes VoteValue = "YES"
	VoteValueNo  VoteValue = "NO"
)

func (enum VoteValue) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VoteValue) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
