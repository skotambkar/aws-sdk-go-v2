// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// The new amount of storage (in gibibytes) to allocate for the DB instance.
	//
	// Not applicable. Storage is managed by the DB Cluster.
	AllocatedStorage *int64 `type:"integer"`

	// Indicates that major version upgrades are allowed. Changing this parameter
	// doesn't result in an outage and the change is asynchronously applied as soon
	// as possible.
	//
	// Constraints: This parameter must be set to true when specifying a value for
	// the EngineVersion parameter that is a different major version than the DB
	// instance's current version.
	AllowMajorVersionUpgrade *bool `type:"boolean"`

	// Specifies whether the modifications in this request and any pending modifications
	// are asynchronously applied as soon as possible, regardless of the PreferredMaintenanceWindow
	// setting for the DB instance.
	//
	// If this parameter is set to false, changes to the DB instance are applied
	// during the next maintenance window. Some parameter changes can cause an outage
	// and are applied on the next call to RebootDBInstance, or the next failure
	// reboot.
	//
	// Default: false
	ApplyImmediately *bool `type:"boolean"`

	// Indicates that minor version upgrades are applied automatically to the DB
	// instance during the maintenance window. Changing this parameter doesn't result
	// in an outage except in the following case and the change is asynchronously
	// applied as soon as possible. An outage will result if this parameter is set
	// to true during the maintenance window, and a newer minor version is available,
	// and Neptune has enabled auto patching for that engine version.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// Not applicable. The retention period for automated backups is managed by
	// the DB cluster. For more information, see ModifyDBCluster.
	//
	// Default: Uses existing setting
	BackupRetentionPeriod *int64 `type:"integer"`

	// Indicates the certificate that needs to be associated with the instance.
	CACertificateIdentifier *string `type:"string"`

	// The configuration setting for the log types to be enabled for export to CloudWatch
	// Logs for a specific DB instance or DB cluster.
	CloudwatchLogsExportConfiguration *CloudwatchLogsExportConfiguration `type:"structure"`

	// True to copy all tags from the DB instance to snapshots of the DB instance,
	// and otherwise false. The default is false.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The new compute and memory capacity of the DB instance, for example, db.m4.large.
	// Not all DB instance classes are available in all AWS Regions.
	//
	// If you modify the DB instance class, an outage occurs during the change.
	// The change is applied during the next maintenance window, unless ApplyImmediately
	// is specified as true for this request.
	//
	// Default: Uses existing setting
	DBInstanceClass *string `type:"string"`

	// The DB instance identifier. This value is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBInstance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The name of the DB parameter group to apply to the DB instance. Changing
	// this setting doesn't result in an outage. The parameter group name itself
	// is changed immediately, but the actual parameter changes are not applied
	// until you reboot the instance without failover. The db instance will NOT
	// be rebooted automatically and the parameter changes will NOT be applied during
	// the next maintenance window.
	//
	// Default: Uses existing setting
	//
	// Constraints: The DB parameter group must be in the same DB parameter group
	// family as this DB instance.
	DBParameterGroupName *string `type:"string"`

	// The port number on which the database accepts connections.
	//
	// The value of the DBPortNumber parameter must not match any of the port values
	// specified for options in the option group for the DB instance.
	//
	// Your database will restart when you change the DBPortNumber value regardless
	// of the value of the ApplyImmediately parameter.
	//
	// Default: 8182
	DBPortNumber *int64 `type:"integer"`

	// A list of DB security groups to authorize on this DB instance. Changing this
	// setting doesn't result in an outage and the change is asynchronously applied
	// as soon as possible.
	//
	// Constraints:
	//
	//    * If supplied, must match existing DBSecurityGroups.
	DBSecurityGroups []string `locationNameList:"DBSecurityGroupName" type:"list"`

	// The new DB subnet group for the DB instance. You can use this parameter to
	// move your DB instance to a different VPC.
	//
	// Changing the subnet group causes an outage during the change. The change
	// is applied during the next maintenance window, unless you specify true for
	// the ApplyImmediately parameter.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetGroup
	DBSubnetGroupName *string `type:"string"`

	// Not supported.
	Domain *string `type:"string"`

	// Not supported
	DomainIAMRoleName *string `type:"string"`

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false.
	//
	// You can enable IAM database authentication for the following database engines
	//
	// Not applicable. Mapping AWS IAM accounts to database accounts is managed
	// by the DB cluster. For more information, see ModifyDBCluster.
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// Not supported.
	EnablePerformanceInsights *bool `type:"boolean"`

	// The version number of the database engine to upgrade to. Changing this parameter
	// results in an outage and the change is applied during the next maintenance
	// window unless the ApplyImmediately parameter is set to true for this request.
	//
	// For major version upgrades, if a nondefault DB parameter group is currently
	// in use, a new DB parameter group in the DB parameter group family for the
	// new engine version must be specified. The new DB parameter group can be the
	// default for that DB parameter group family.
	EngineVersion *string `type:"string"`

	// The new Provisioned IOPS (I/O operations per second) value for the instance.
	//
	// Changing this setting doesn't result in an outage and the change is applied
	// during the next maintenance window unless the ApplyImmediately parameter
	// is set to true for this request.
	//
	// Default: Uses existing setting
	Iops *int64 `type:"integer"`

	// Not supported.
	LicenseModel *string `type:"string"`

	// Not applicable.
	MasterUserPassword *string `type:"string"`

	// The interval, in seconds, between points when Enhanced Monitoring metrics
	// are collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0.
	//
	// If MonitoringRoleArn is specified, then you must also set MonitoringInterval
	// to a value other than 0.
	//
	// Valid Values: 0, 1, 5, 10, 15, 30, 60
	MonitoringInterval *int64 `type:"integer"`

	// The ARN for the IAM role that permits Neptune to send enhanced monitoring
	// metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess.
	//
	// If MonitoringInterval is set to a value other than 0, then you must supply
	// a MonitoringRoleArn value.
	MonitoringRoleArn *string `type:"string"`

	// Specifies if the DB instance is a Multi-AZ deployment. Changing this parameter
	// doesn't result in an outage and the change is applied during the next maintenance
	// window unless the ApplyImmediately parameter is set to true for this request.
	MultiAZ *bool `type:"boolean"`

	// The new DB instance identifier for the DB instance when renaming a DB instance.
	// When you change the DB instance identifier, an instance reboot will occur
	// immediately if you set Apply Immediately to true, or will occur during the
	// next maintenance window if Apply Immediately to false. This value is stored
	// as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * The first character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: mydbinstance
	NewDBInstanceIdentifier *string `type:"string"`

	// Indicates that the DB instance should be associated with the specified option
	// group. Changing this parameter doesn't result in an outage except in the
	// following case and the change is applied during the next maintenance window
	// unless the ApplyImmediately parameter is set to true for this request. If
	// the parameter change results in an option group that enables OEM, this change
	// can cause a brief (sub-second) period during which new connections are rejected
	// but existing connections are not interrupted.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance once it is associated with a DB instance
	OptionGroupName *string `type:"string"`

	// Not supported.
	PerformanceInsightsKMSKeyId *string `type:"string"`

	// The daily time range during which automated backups are created if automated
	// backups are enabled.
	//
	// Not applicable. The daily time range for creating automated backups is managed
	// by the DB cluster. For more information, see ModifyDBCluster.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi
	//
	//    * Must be in Universal Time Coordinated (UTC)
	//
	//    * Must not conflict with the preferred maintenance window
	//
	//    * Must be at least 30 minutes
	PreferredBackupWindow *string `type:"string"`

	// The weekly time range (in UTC) during which system maintenance can occur,
	// which might result in an outage. Changing this parameter doesn't result in
	// an outage, except in the following situation, and the change is asynchronously
	// applied as soon as possible. If there are pending actions that cause a reboot,
	// and the maintenance window is changed to include the current time, then changing
	// this parameter will cause a reboot of the DB instance. If moving this window
	// to the current time, there must be at least 30 minutes between the current
	// time and end of the window to ensure pending changes are applied.
	//
	// Default: Uses existing setting
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// Valid Days: Mon | Tue | Wed | Thu | Fri | Sat | Sun
	//
	// Constraints: Must be at least 30 minutes
	PreferredMaintenanceWindow *string `type:"string"`

	// A value that specifies the order in which a Read Replica is promoted to the
	// primary instance after a failure of the existing primary instance.
	//
	// Default: 1
	//
	// Valid Values: 0 - 15
	PromotionTier *int64 `type:"integer"`

	// This flag should no longer be used.
	PubliclyAccessible *bool `deprecated:"true" type:"boolean"`

	// Not supported.
	StorageType *string `type:"string"`

	// The ARN from the key store with which to associate the instance for TDE encryption.
	TdeCredentialArn *string `type:"string"`

	// The password for the given ARN from the key store in order to access the
	// device.
	TdeCredentialPassword *string `type:"string"`

	// A list of EC2 VPC security groups to authorize on this DB instance. This
	// change is asynchronously applied as soon as possible.
	//
	// Not applicable. The associated list of EC2 VPC security groups is managed
	// by the DB cluster. For more information, see ModifyDBCluster.
	//
	// Constraints:
	//
	//    * If supplied, must match existing VpcSecurityGroupIds.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s ModifyDBInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBInstanceInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s ModifyDBInstanceOutput) String() string {
	return awsutil.Prettify(s)
}
