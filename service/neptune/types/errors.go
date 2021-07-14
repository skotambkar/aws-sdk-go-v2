// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Specified CIDRIP or EC2 security group is not authorized for the specified DB
// security group. Neptune may not also be authorized via IAM to perform necessary
// actions on your behalf.
type AuthorizationNotFoundFault struct {
	Message *string
}

func (e *AuthorizationNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthorizationNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthorizationNotFoundFault) ErrorCode() string             { return "AuthorizationNotFound" }
func (e *AuthorizationNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// CertificateIdentifier does not refer to an existing certificate.
type CertificateNotFoundFault struct {
	Message *string
}

func (e *CertificateNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateNotFoundFault) ErrorCode() string             { return "CertificateNotFound" }
func (e *CertificateNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// User already has a DB cluster with the given identifier.
type DBClusterAlreadyExistsFault struct {
	Message *string
}

func (e *DBClusterAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterAlreadyExistsFault) ErrorCode() string             { return "DBClusterAlreadyExistsFault" }
func (e *DBClusterAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified custom endpoint cannot be created because it already exists.
type DBClusterEndpointAlreadyExistsFault struct {
	Message *string
}

func (e *DBClusterEndpointAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterEndpointAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterEndpointAlreadyExistsFault) ErrorCode() string {
	return "DBClusterEndpointAlreadyExistsFault"
}
func (e *DBClusterEndpointAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified custom endpoint doesn't exist.
type DBClusterEndpointNotFoundFault struct {
	Message *string
}

func (e *DBClusterEndpointNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterEndpointNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterEndpointNotFoundFault) ErrorCode() string             { return "DBClusterEndpointNotFoundFault" }
func (e *DBClusterEndpointNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The cluster already has the maximum number of custom endpoints.
type DBClusterEndpointQuotaExceededFault struct {
	Message *string
}

func (e *DBClusterEndpointQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterEndpointQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterEndpointQuotaExceededFault) ErrorCode() string {
	return "DBClusterEndpointQuotaExceededFault"
}
func (e *DBClusterEndpointQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBClusterIdentifier does not refer to an existing DB cluster.
type DBClusterNotFoundFault struct {
	Message *string
}

func (e *DBClusterNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterNotFoundFault) ErrorCode() string             { return "DBClusterNotFoundFault" }
func (e *DBClusterNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBClusterParameterGroupName does not refer to an existing DB Cluster parameter
// group.
type DBClusterParameterGroupNotFoundFault struct {
	Message *string
}

func (e *DBClusterParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorCode() string {
	return "DBClusterParameterGroupNotFound"
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// User attempted to create a new DB cluster and the user has already reached the
// maximum allowed DB cluster quota.
type DBClusterQuotaExceededFault struct {
	Message *string
}

func (e *DBClusterQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterQuotaExceededFault) ErrorCode() string             { return "DBClusterQuotaExceededFault" }
func (e *DBClusterQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified IAM role Amazon Resource Name (ARN) is already associated with the
// specified DB cluster.
type DBClusterRoleAlreadyExistsFault struct {
	Message *string
}

func (e *DBClusterRoleAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterRoleAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterRoleAlreadyExistsFault) ErrorCode() string             { return "DBClusterRoleAlreadyExists" }
func (e *DBClusterRoleAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified IAM role Amazon Resource Name (ARN) is not associated with the
// specified DB cluster.
type DBClusterRoleNotFoundFault struct {
	Message *string
}

func (e *DBClusterRoleNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterRoleNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterRoleNotFoundFault) ErrorCode() string             { return "DBClusterRoleNotFound" }
func (e *DBClusterRoleNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the maximum number of IAM roles that can be associated with
// the specified DB cluster.
type DBClusterRoleQuotaExceededFault struct {
	Message *string
}

func (e *DBClusterRoleQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterRoleQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterRoleQuotaExceededFault) ErrorCode() string             { return "DBClusterRoleQuotaExceeded" }
func (e *DBClusterRoleQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// User already has a DB cluster snapshot with the given identifier.
type DBClusterSnapshotAlreadyExistsFault struct {
	Message *string
}

func (e *DBClusterSnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorCode() string {
	return "DBClusterSnapshotAlreadyExistsFault"
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBClusterSnapshotIdentifier does not refer to an existing DB cluster snapshot.
type DBClusterSnapshotNotFoundFault struct {
	Message *string
}

func (e *DBClusterSnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterSnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterSnapshotNotFoundFault) ErrorCode() string             { return "DBClusterSnapshotNotFoundFault" }
func (e *DBClusterSnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// User already has a DB instance with the given identifier.
type DBInstanceAlreadyExistsFault struct {
	Message *string
}

func (e *DBInstanceAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBInstanceAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBInstanceAlreadyExistsFault) ErrorCode() string             { return "DBInstanceAlreadyExists" }
func (e *DBInstanceAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBInstanceIdentifier does not refer to an existing DB instance.
type DBInstanceNotFoundFault struct {
	Message *string
}

func (e *DBInstanceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBInstanceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBInstanceNotFoundFault) ErrorCode() string             { return "DBInstanceNotFound" }
func (e *DBInstanceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A DB parameter group with the same name exists.
type DBParameterGroupAlreadyExistsFault struct {
	Message *string
}

func (e *DBParameterGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorCode() string {
	return "DBParameterGroupAlreadyExists"
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBParameterGroupName does not refer to an existing DB parameter group.
type DBParameterGroupNotFoundFault struct {
	Message *string
}

func (e *DBParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupNotFoundFault) ErrorCode() string             { return "DBParameterGroupNotFound" }
func (e *DBParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request would result in user exceeding the allowed number of DB parameter
// groups.
type DBParameterGroupQuotaExceededFault struct {
	Message *string
}

func (e *DBParameterGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupQuotaExceededFault) ErrorCode() string {
	return "DBParameterGroupQuotaExceeded"
}
func (e *DBParameterGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBSecurityGroupName does not refer to an existing DB security group.
type DBSecurityGroupNotFoundFault struct {
	Message *string
}

func (e *DBSecurityGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSecurityGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSecurityGroupNotFoundFault) ErrorCode() string             { return "DBSecurityGroupNotFound" }
func (e *DBSecurityGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSnapshotIdentifier is already used by an existing snapshot.
type DBSnapshotAlreadyExistsFault struct {
	Message *string
}

func (e *DBSnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSnapshotAlreadyExistsFault) ErrorCode() string             { return "DBSnapshotAlreadyExists" }
func (e *DBSnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSnapshotIdentifier does not refer to an existing DB snapshot.
type DBSnapshotNotFoundFault struct {
	Message *string
}

func (e *DBSnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSnapshotNotFoundFault) ErrorCode() string             { return "DBSnapshotNotFound" }
func (e *DBSnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSubnetGroupName is already used by an existing DB subnet group.
type DBSubnetGroupAlreadyExistsFault struct {
	Message *string
}

func (e *DBSubnetGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupAlreadyExistsFault) ErrorCode() string             { return "DBSubnetGroupAlreadyExists" }
func (e *DBSubnetGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Subnets in the DB subnet group should cover at least two Availability Zones
// unless there is only one Availability Zone.
type DBSubnetGroupDoesNotCoverEnoughAZs struct {
	Message *string
}

func (e *DBSubnetGroupDoesNotCoverEnoughAZs) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorCode() string {
	return "DBSubnetGroupDoesNotCoverEnoughAZs"
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBSubnetGroupName does not refer to an existing DB subnet group.
type DBSubnetGroupNotFoundFault struct {
	Message *string
}

func (e *DBSubnetGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupNotFoundFault) ErrorCode() string             { return "DBSubnetGroupNotFoundFault" }
func (e *DBSubnetGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request would result in user exceeding the allowed number of DB subnet groups.
type DBSubnetGroupQuotaExceededFault struct {
	Message *string
}

func (e *DBSubnetGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupQuotaExceededFault) ErrorCode() string             { return "DBSubnetGroupQuotaExceeded" }
func (e *DBSubnetGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request would result in user exceeding the allowed number of subnets in a DB
// subnet groups.
type DBSubnetQuotaExceededFault struct {
	Message *string
}

func (e *DBSubnetQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetQuotaExceededFault) ErrorCode() string             { return "DBSubnetQuotaExceededFault" }
func (e *DBSubnetQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The DB upgrade failed because a resource the DB depends on could not be
// modified.
type DBUpgradeDependencyFailureFault struct {
	Message *string
}

func (e *DBUpgradeDependencyFailureFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBUpgradeDependencyFailureFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBUpgradeDependencyFailureFault) ErrorCode() string             { return "DBUpgradeDependencyFailure" }
func (e *DBUpgradeDependencyFailureFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Domain does not refer to an existing Active Directory Domain.
type DomainNotFoundFault struct {
	Message *string
}

func (e *DomainNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DomainNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DomainNotFoundFault) ErrorCode() string             { return "DomainNotFoundFault" }
func (e *DomainNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the number of events you can subscribe to.
type EventSubscriptionQuotaExceededFault struct {
	Message *string
}

func (e *EventSubscriptionQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EventSubscriptionQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EventSubscriptionQuotaExceededFault) ErrorCode() string {
	return "EventSubscriptionQuotaExceeded"
}
func (e *EventSubscriptionQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Request would result in user exceeding the allowed number of DB instances.
type InstanceQuotaExceededFault struct {
	Message *string
}

func (e *InstanceQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InstanceQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InstanceQuotaExceededFault) ErrorCode() string             { return "InstanceQuotaExceeded" }
func (e *InstanceQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The DB cluster does not have enough capacity for the current operation.
type InsufficientDBClusterCapacityFault struct {
	Message *string
}

func (e *InsufficientDBClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientDBClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientDBClusterCapacityFault) ErrorCode() string {
	return "InsufficientDBClusterCapacityFault"
}
func (e *InsufficientDBClusterCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Specified DB instance class is not available in the specified Availability Zone.
type InsufficientDBInstanceCapacityFault struct {
	Message *string
}

func (e *InsufficientDBInstanceCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientDBInstanceCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientDBInstanceCapacityFault) ErrorCode() string {
	return "InsufficientDBInstanceCapacity"
}
func (e *InsufficientDBInstanceCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// There is insufficient storage available for the current action. You may be able
// to resolve this error by updating your subnet group to use different
// Availability Zones that have more storage available.
type InsufficientStorageClusterCapacityFault struct {
	Message *string
}

func (e *InsufficientStorageClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientStorageClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientStorageClusterCapacityFault) ErrorCode() string {
	return "InsufficientStorageClusterCapacity"
}
func (e *InsufficientStorageClusterCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested operation cannot be performed on the endpoint while the endpoint
// is in this state.
type InvalidDBClusterEndpointStateFault struct {
	Message *string
}

func (e *InvalidDBClusterEndpointStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBClusterEndpointStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBClusterEndpointStateFault) ErrorCode() string {
	return "InvalidDBClusterEndpointStateFault"
}
func (e *InvalidDBClusterEndpointStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The supplied value is not a valid DB cluster snapshot state.
type InvalidDBClusterSnapshotStateFault struct {
	Message *string
}

func (e *InvalidDBClusterSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorCode() string {
	return "InvalidDBClusterSnapshotStateFault"
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The DB cluster is not in a valid state.
type InvalidDBClusterStateFault struct {
	Message *string
}

func (e *InvalidDBClusterStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBClusterStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBClusterStateFault) ErrorCode() string             { return "InvalidDBClusterStateFault" }
func (e *InvalidDBClusterStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified DB instance is not in the available state.
type InvalidDBInstanceStateFault struct {
	Message *string
}

func (e *InvalidDBInstanceStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBInstanceStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBInstanceStateFault) ErrorCode() string             { return "InvalidDBInstanceState" }
func (e *InvalidDBInstanceStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The DB parameter group is in use or is in an invalid state. If you are
// attempting to delete the parameter group, you cannot delete it when the
// parameter group is in this state.
type InvalidDBParameterGroupStateFault struct {
	Message *string
}

func (e *InvalidDBParameterGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBParameterGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBParameterGroupStateFault) ErrorCode() string             { return "InvalidDBParameterGroupState" }
func (e *InvalidDBParameterGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The state of the DB security group does not allow deletion.
type InvalidDBSecurityGroupStateFault struct {
	Message *string
}

func (e *InvalidDBSecurityGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSecurityGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSecurityGroupStateFault) ErrorCode() string             { return "InvalidDBSecurityGroupState" }
func (e *InvalidDBSecurityGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The state of the DB snapshot does not allow deletion.
type InvalidDBSnapshotStateFault struct {
	Message *string
}

func (e *InvalidDBSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSnapshotStateFault) ErrorCode() string             { return "InvalidDBSnapshotState" }
func (e *InvalidDBSnapshotStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The DB subnet group cannot be deleted because it is in use.
type InvalidDBSubnetGroupStateFault struct {
	Message *string
}

func (e *InvalidDBSubnetGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSubnetGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSubnetGroupStateFault) ErrorCode() string             { return "InvalidDBSubnetGroupStateFault" }
func (e *InvalidDBSubnetGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The DB subnet is not in the available state.
type InvalidDBSubnetStateFault struct {
	Message *string
}

func (e *InvalidDBSubnetStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSubnetStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSubnetStateFault) ErrorCode() string             { return "InvalidDBSubnetStateFault" }
func (e *InvalidDBSubnetStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The event subscription is in an invalid state.
type InvalidEventSubscriptionStateFault struct {
	Message *string
}

func (e *InvalidEventSubscriptionStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEventSubscriptionStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEventSubscriptionStateFault) ErrorCode() string {
	return "InvalidEventSubscriptionState"
}
func (e *InvalidEventSubscriptionStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Cannot restore from vpc backup to non-vpc DB instance.
type InvalidRestoreFault struct {
	Message *string
}

func (e *InvalidRestoreFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRestoreFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRestoreFault) ErrorCode() string             { return "InvalidRestoreFault" }
func (e *InvalidRestoreFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested subnet is invalid, or multiple subnets were requested that are not
// all in a common VPC.
type InvalidSubnet struct {
	Message *string
}

func (e *InvalidSubnet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnet) ErrorCode() string             { return "InvalidSubnet" }
func (e *InvalidSubnet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DB subnet group does not cover all Availability Zones after it is created
// because users' change.
type InvalidVPCNetworkStateFault struct {
	Message *string
}

func (e *InvalidVPCNetworkStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidVPCNetworkStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidVPCNetworkStateFault) ErrorCode() string             { return "InvalidVPCNetworkStateFault" }
func (e *InvalidVPCNetworkStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Error accessing KMS key.
type KMSKeyNotAccessibleFault struct {
	Message *string
}

func (e *KMSKeyNotAccessibleFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSKeyNotAccessibleFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSKeyNotAccessibleFault) ErrorCode() string             { return "KMSKeyNotAccessibleFault" }
func (e *KMSKeyNotAccessibleFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The designated option group could not be found.
type OptionGroupNotFoundFault struct {
	Message *string
}

func (e *OptionGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OptionGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OptionGroupNotFoundFault) ErrorCode() string             { return "OptionGroupNotFoundFault" }
func (e *OptionGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Provisioned IOPS not available in the specified Availability Zone.
type ProvisionedIopsNotAvailableInAZFault struct {
	Message *string
}

func (e *ProvisionedIopsNotAvailableInAZFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProvisionedIopsNotAvailableInAZFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProvisionedIopsNotAvailableInAZFault) ErrorCode() string {
	return "ProvisionedIopsNotAvailableInAZFault"
}
func (e *ProvisionedIopsNotAvailableInAZFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified resource ID was not found.
type ResourceNotFoundFault struct {
	Message *string
}

func (e *ResourceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundFault) ErrorCode() string             { return "ResourceNotFoundFault" }
func (e *ResourceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the maximum number of accounts that you can share a manual DB
// snapshot with.
type SharedSnapshotQuotaExceededFault struct {
	Message *string
}

func (e *SharedSnapshotQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SharedSnapshotQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SharedSnapshotQuotaExceededFault) ErrorCode() string             { return "SharedSnapshotQuotaExceeded" }
func (e *SharedSnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request would result in user exceeding the allowed number of DB snapshots.
type SnapshotQuotaExceededFault struct {
	Message *string
}

func (e *SnapshotQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotQuotaExceededFault) ErrorCode() string             { return "SnapshotQuotaExceeded" }
func (e *SnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The SNS topic is invalid.
type SNSInvalidTopicFault struct {
	Message *string
}

func (e *SNSInvalidTopicFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSInvalidTopicFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSInvalidTopicFault) ErrorCode() string             { return "SNSInvalidTopic" }
func (e *SNSInvalidTopicFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is no SNS authorization.
type SNSNoAuthorizationFault struct {
	Message *string
}

func (e *SNSNoAuthorizationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSNoAuthorizationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSNoAuthorizationFault) ErrorCode() string             { return "SNSNoAuthorization" }
func (e *SNSNoAuthorizationFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The ARN of the SNS topic could not be found.
type SNSTopicArnNotFoundFault struct {
	Message *string
}

func (e *SNSTopicArnNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSTopicArnNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSTopicArnNotFoundFault) ErrorCode() string             { return "SNSTopicArnNotFound" }
func (e *SNSTopicArnNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The source could not be found.
type SourceNotFoundFault struct {
	Message *string
}

func (e *SourceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SourceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SourceNotFoundFault) ErrorCode() string             { return "SourceNotFound" }
func (e *SourceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request would result in user exceeding the allowed amount of storage available
// across all DB instances.
type StorageQuotaExceededFault struct {
	Message *string
}

func (e *StorageQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageQuotaExceededFault) ErrorCode() string             { return "StorageQuotaExceeded" }
func (e *StorageQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// StorageType specified cannot be associated with the DB Instance.
type StorageTypeNotSupportedFault struct {
	Message *string
}

func (e *StorageTypeNotSupportedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageTypeNotSupportedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageTypeNotSupportedFault) ErrorCode() string             { return "StorageTypeNotSupported" }
func (e *StorageTypeNotSupportedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The DB subnet is already in use in the Availability Zone.
type SubnetAlreadyInUse struct {
	Message *string
}

func (e *SubnetAlreadyInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetAlreadyInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetAlreadyInUse) ErrorCode() string             { return "SubnetAlreadyInUse" }
func (e *SubnetAlreadyInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This subscription already exists.
type SubscriptionAlreadyExistFault struct {
	Message *string
}

func (e *SubscriptionAlreadyExistFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubscriptionAlreadyExistFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubscriptionAlreadyExistFault) ErrorCode() string             { return "SubscriptionAlreadyExist" }
func (e *SubscriptionAlreadyExistFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The designated subscription category could not be found.
type SubscriptionCategoryNotFoundFault struct {
	Message *string
}

func (e *SubscriptionCategoryNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubscriptionCategoryNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubscriptionCategoryNotFoundFault) ErrorCode() string             { return "SubscriptionCategoryNotFound" }
func (e *SubscriptionCategoryNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The designated subscription could not be found.
type SubscriptionNotFoundFault struct {
	Message *string
}

func (e *SubscriptionNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubscriptionNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubscriptionNotFoundFault) ErrorCode() string             { return "SubscriptionNotFound" }
func (e *SubscriptionNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
