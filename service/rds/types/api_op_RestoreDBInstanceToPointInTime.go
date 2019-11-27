// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreDBInstanceToPointInTimeInput struct {
	_ struct{} `type:"structure"`

	// A value that indicates whether minor version upgrades are applied automatically
	// to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The Availability Zone (AZ) where the DB instance will be created.
	//
	// Default: A random, system-chosen Availability Zone.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	//
	// Example: us-east-1a
	AvailabilityZone *string `type:"string"`

	// A value that indicates whether to copy all tags from the restored DB instance
	// to snapshots of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The compute and memory capacity of the Amazon RDS DB instance, for example,
	// db.m4.large. Not all DB instance classes are available in all AWS Regions,
	// or for all database engines. For the full list of DB instance classes, and
	// availability for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide.
	//
	// Default: The same DBInstanceClass as the original DB instance.
	DBInstanceClass *string `type:"string"`

	// The database name for the restored DB instance.
	//
	// This parameter is not used for the MySQL or MariaDB engines.
	DBName *string `type:"string"`

	// The name of the DB parameter group to associate with this DB instance.
	//
	// If you do not specify a value for DBParameterGroupName, then the default
	// DBParameterGroup for the specified DB engine is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBParameterGroup.
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	DBParameterGroupName *string `type:"string"`

	// The DB subnet group name to use for the new instance.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see Deleting a DB
	// Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool `type:"boolean"`

	// Specify the Active Directory directory ID to restore the DB instance in.
	// The domain must be created prior to this operation. Currently, only Microsoft
	// SQL Server and Oracle DB instances can be created in an Active Directory
	// Domain.
	//
	// For Microsoft SQL Server DB instances, Amazon RDS can use Windows Authentication
	// to authenticate users that connect to the DB instance. For more information,
	// see Using Windows Authentication with an Amazon RDS DB Instance Running Microsoft
	// SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_SQLServerWinAuth.html)
	// in the Amazon RDS User Guide.
	//
	// For Oracle DB instances, Amazon RDS can use Kerberos Authentication to authenticate
	// users that connect to the DB instance. For more information, see Using Kerberos
	// Authentication with Amazon RDS for Oracle (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos.html)
	// in the Amazon RDS User Guide.
	Domain *string `type:"string"`

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string `type:"string"`

	// The list of logs that the restored DB instance is to export to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For information about the supported DB engines, see CreateDBInstance.
	//
	// For more information about IAM database authentication, see IAM Database
	// Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide.
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The database engine to use for the new instance.
	//
	// Default: The same as source
	//
	// Constraint: Must be compatible with the engine of the source
	//
	// Valid Values:
	//
	//    * mariadb
	//
	//    * mysql
	//
	//    * oracle-ee
	//
	//    * oracle-se2
	//
	//    * oracle-se1
	//
	//    * oracle-se
	//
	//    * postgres
	//
	//    * sqlserver-ee
	//
	//    * sqlserver-se
	//
	//    * sqlserver-ex
	//
	//    * sqlserver-web
	Engine *string `type:"string"`

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance.
	//
	// Constraints: Must be an integer greater than 1000.
	//
	// SQL Server
	//
	// Setting the IOPS value for the SQL Server database engine is not supported.
	Iops *int64 `type:"integer"`

	// License model information for the restored DB instance.
	//
	// Default: Same as source.
	//
	// Valid values: license-included | bring-your-own-license | general-public-license
	LicenseModel *string `type:"string"`

	// A value that indicates whether the DB instance is a Multi-AZ deployment.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	MultiAZ *bool `type:"boolean"`

	// The name of the option group to be used for the restored DB instance.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance once it is associated with a DB instance
	OptionGroupName *string `type:"string"`

	// The port number on which the database accepts connections.
	//
	// Constraints: Value must be 1150-65535
	//
	// Default: The same port as the original DB instance.
	Port *int64 `type:"integer"`

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []ProcessorFeature `locationNameList:"ProcessorFeature" type:"list"`

	// A value that indicates whether the DB instance is publicly accessible. When
	// the DB instance is publicly accessible, it is an Internet-facing instance
	// with a publicly resolvable DNS name, which resolves to a public IP address.
	// When the DB instance is not publicly accessible, it is an internal instance
	// with a DNS name that resolves to a private IP address. For more information,
	// see CreateDBInstance.
	PubliclyAccessible *bool `type:"boolean"`

	// The date and time to restore from.
	//
	// Valid Values: Value must be a time in Universal Coordinated Time (UTC) format
	//
	// Constraints:
	//
	//    * Must be before the latest restorable time for the DB instance
	//
	//    * Can't be specified if the UseLatestRestorableTime parameter is enabled
	//
	// Example: 2009-09-07T23:45:00Z
	RestoreTime *time.Time `type:"timestamp"`

	// The identifier of the source DB instance from which to restore.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DB instance.
	SourceDBInstanceIdentifier *string `type:"string"`

	// The resource ID of the source DB instance from which to restore.
	SourceDbiResourceId *string `type:"string"`

	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: standard | gp2 | io1
	//
	// If you specify io1, you must also include a value for the Iops parameter.
	//
	// Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string `type:"string"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// The name of the new DB instance to be created.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// TargetDBInstanceIdentifier is a required field
	TargetDBInstanceIdentifier *string `type:"string" required:"true"`

	// The ARN from the key store with which to associate the instance for TDE encryption.
	TdeCredentialArn *string `type:"string"`

	// The password for the given ARN from the key store in order to access the
	// device.
	TdeCredentialPassword *string `type:"string"`

	// A value that indicates whether the DB instance class of the DB instance uses
	// its default processor features.
	UseDefaultProcessorFeatures *bool `type:"boolean"`

	// A value that indicates whether the DB instance is restored from the latest
	// backup time. By default, the DB instance is not restored from the latest
	// backup time.
	//
	// Constraints: Can't be specified if the RestoreTime parameter is provided.
	UseLatestRestorableTime *bool `type:"boolean"`

	// A list of EC2 VPC security groups to associate with this DB instance.
	//
	// Default: The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBInstanceToPointInTimeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBInstanceToPointInTimeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBInstanceToPointInTimeInput"}

	if s.TargetDBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetDBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreDBInstanceToPointInTimeOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s RestoreDBInstanceToPointInTimeOutput) String() string {
	return awsutil.Prettify(s)
}
