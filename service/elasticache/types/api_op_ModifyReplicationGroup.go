// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/enums"
)

// Represents the input of a ModifyReplicationGroups operation.
type ModifyReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// If true, this parameter causes the modifications in this request and any
	// pending modifications to be applied, asynchronously and as soon as possible,
	// regardless of the PreferredMaintenanceWindow setting for the replication
	// group.
	//
	// If false, changes to the nodes in the replication group are applied on the
	// next maintenance reboot, or the next failure reboot, whichever occurs first.
	//
	// Valid values: true | false
	//
	// Default: false
	ApplyImmediately *bool `type:"boolean"`

	// Reserved parameter. The password used to access a password protected server.
	// This parameter must be specified with the auth-token-update-strategy parameter.
	// Password constraints:
	//
	//    * Must be only printable ASCII characters
	//
	//    * Must be at least 16 characters and no more than 128 characters in length
	//
	//    * Cannot contain any of the following characters: '/', '"', or '@', '%'
	//
	// For more information, see AUTH password at AUTH (http://redis.io/commands/AUTH).
	AuthToken *string `type:"string"`

	// Specifies the strategy to use to update the AUTH token. This parameter must
	// be specified with the auth-token parameter. Possible values:
	//
	//    * Rotate
	//
	//    * Set
	//
	// For more information, see Authenticating Users with Redis AUTH (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/auth.html)
	AuthTokenUpdateStrategy enums.AuthTokenUpdateStrategyType `type:"string" enum:"true"`

	// This parameter is currently disabled.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// Determines whether a read replica is automatically promoted to read/write
	// primary if the existing primary encounters a failure.
	//
	// Valid values: true | false
	//
	// Amazon ElastiCache for Redis does not support Multi-AZ with automatic failover
	// on:
	//
	//    * Redis versions earlier than 2.8.6.
	//
	//    * Redis (cluster mode disabled): T1 node types.
	//
	//    * Redis (cluster mode enabled): T1 node types.
	AutomaticFailoverEnabled *bool `type:"boolean"`

	// A valid cache node type that you want to scale this replication group to.
	CacheNodeType *string `type:"string"`

	// The name of the cache parameter group to apply to all of the clusters in
	// this replication group. This change is asynchronously applied as soon as
	// possible for parameters when the ApplyImmediately parameter is specified
	// as true for this request.
	CacheParameterGroupName *string `type:"string"`

	// A list of cache security group names to authorize for the clusters in this
	// replication group. This change is asynchronously applied as soon as possible.
	//
	// This parameter can be used only with replication group containing clusters
	// running outside of an Amazon Virtual Private Cloud (Amazon VPC).
	//
	// Constraints: Must contain no more than 255 alphanumeric characters. Must
	// not be Default.
	CacheSecurityGroupNames []string `locationNameList:"CacheSecurityGroupName" type:"list"`

	// The upgraded version of the cache engine to be run on the clusters in the
	// replication group.
	//
	// Important: You can upgrade to a newer engine version (see Selecting a Cache
	// Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)),
	// but you cannot downgrade to an earlier engine version. If you want to use
	// an earlier engine version, you must delete the existing replication group
	// and create it anew with the earlier engine version.
	EngineVersion *string `type:"string"`

	// Deprecated. This parameter is not used.
	NodeGroupId *string `deprecated:"true" type:"string"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications
	// are sent.
	//
	// The Amazon SNS topic owner must be same as the replication group owner.
	NotificationTopicArn *string `type:"string"`

	// The status of the Amazon SNS notification topic for the replication group.
	// Notifications are sent only if the status is active.
	//
	// Valid values: active | inactive
	NotificationTopicStatus *string `type:"string"`

	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// Valid values for ddd are:
	//
	//    * sun
	//
	//    * mon
	//
	//    * tue
	//
	//    * wed
	//
	//    * thu
	//
	//    * fri
	//
	//    * sat
	//
	// Example: sun:23:00-mon:01:30
	PreferredMaintenanceWindow *string `type:"string"`

	// For replication groups with a single primary, if this parameter is specified,
	// ElastiCache promotes the specified cluster in the specified replication group
	// to the primary role. The nodes of all other clusters in the replication group
	// are read replicas.
	PrimaryClusterId *string `type:"string"`

	// A description for the replication group. Maximum length is 255 characters.
	ReplicationGroupDescription *string `type:"string"`

	// The identifier of the replication group to modify.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`

	// Specifies the VPC Security Groups associated with the clusters in the replication
	// group.
	//
	// This parameter can be used only with replication group containing clusters
	// running in an Amazon Virtual Private Cloud (Amazon VPC).
	SecurityGroupIds []string `locationNameList:"SecurityGroupId" type:"list"`

	// The number of days for which ElastiCache retains automatic node group (shard)
	// snapshots before deleting them. For example, if you set SnapshotRetentionLimit
	// to 5, a snapshot that was taken today is retained for 5 days before being
	// deleted.
	//
	// Important If the value of SnapshotRetentionLimit is set to zero (0), backups
	// are turned off.
	SnapshotRetentionLimit *int64 `type:"integer"`

	// The daily time range (in UTC) during which ElastiCache begins taking a daily
	// snapshot of the node group (shard) specified by SnapshottingClusterId.
	//
	// Example: 05:00-09:00
	//
	// If you do not specify this parameter, ElastiCache automatically chooses an
	// appropriate time range.
	SnapshotWindow *string `type:"string"`

	// The cluster ID that is used as the daily snapshot source for the replication
	// group. This parameter cannot be set for Redis (cluster mode enabled) replication
	// groups.
	SnapshottingClusterId *string `type:"string"`
}

// String returns the string representation
func (s ModifyReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyReplicationGroupInput"}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyReplicationGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s ModifyReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}
