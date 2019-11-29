// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Specifies the settings for each trail.
type CreateTrailInput struct {
	_ struct{} `type:"structure"`

	// Specifies a log group name using an Amazon Resource Name (ARN), a unique
	// identifier that represents the log group to which CloudTrail logs will be
	// delivered. Not required unless you specify CloudWatchLogsRoleArn.
	CloudWatchLogsLogGroupArn *string `type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user's log group.
	CloudWatchLogsRoleArn *string `type:"string"`

	// Specifies whether log file integrity validation is enabled. The default is
	// false.
	//
	// When you disable log file integrity validation, the chain of digest files
	// is broken after one hour. CloudTrail will not create digest files for log
	// files that were delivered during a period in which log file integrity validation
	// was disabled. For example, if you enable log file integrity validation at
	// noon on January 1, disable it at noon on January 2, and re-enable it at noon
	// on January 10, digest files will not be created for the log files delivered
	// from noon on January 2 to noon on January 10. The same applies whenever you
	// stop CloudTrail logging or delete a trail.
	EnableLogFileValidation *bool `type:"boolean"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies whether the trail is created in the current region or in all regions.
	// The default is false, which creates a trail only in the region where you
	// are signed in. As a best practice, consider creating trails that log events
	// in all regions.
	IsMultiRegionTrail *bool `type:"boolean"`

	// Specifies whether the trail is created for all accounts in an organization
	// in AWS Organizations, or only for the current AWS account. The default is
	// false, and cannot be true unless the call is made on behalf of an AWS account
	// that is the master account for an organization in AWS Organizations.
	IsOrganizationTrail *bool `type:"boolean"`

	// Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail.
	// The value can be an alias name prefixed by "alias/", a fully specified ARN
	// to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// Examples:
	//
	//    * alias/MyAliasName
	//
	//    * arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	//
	//    * arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//    * 12345678-1234-1234-1234-123456789012
	KmsKeyId *string `type:"string"`

	// Specifies the name of the trail. The name must meet the following requirements:
	//
	//    * Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
	//    (_), or dashes (-)
	//
	//    * Start with a letter or number, and end with a letter or number
	//
	//    * Be between 3 and 128 characters
	//
	//    * Have no adjacent periods, underscores or dashes. Names like my-_namespace
	//    and my--namespace are invalid.
	//
	//    * Not be in IP address format (for example, 192.168.5.4)
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files. See Amazon S3 Bucket Naming Requirements (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).
	//
	// S3BucketName is a required field
	S3BucketName *string `type:"string" required:"true"`

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket
	// you have designated for log file delivery. For more information, see Finding
	// Your CloudTrail Log Files (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).
	// The maximum length is 200 characters.
	S3KeyPrefix *string `type:"string"`

	// Specifies the name of the Amazon SNS topic defined for notification of log
	// file delivery. The maximum length is 256 characters.
	SnsTopicName *string `type:"string"`

	// A list of tags.
	TagsList []Tag `type:"list"`
}

// String returns the string representation
func (s CreateTrailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTrailInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.S3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3BucketName"))
	}
	if s.TagsList != nil {
		for i, v := range s.TagsList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagsList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type CreateTrailOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the Amazon Resource Name (ARN) of the log group to which CloudTrail
	// logs will be delivered.
	CloudWatchLogsLogGroupArn *string `type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user's log group.
	CloudWatchLogsRoleArn *string `type:"string"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies whether the trail exists in one region or in all regions.
	IsMultiRegionTrail *bool `type:"boolean"`

	// Specifies whether the trail is an organization trail.
	IsOrganizationTrail *bool `type:"boolean"`

	// Specifies the KMS key ID that encrypts the logs delivered by CloudTrail.
	// The value is a fully specified ARN to a KMS key in the format:
	//
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string `type:"string"`

	// Specifies whether log file integrity validation is enabled.
	LogFileValidationEnabled *bool `type:"boolean"`

	// Specifies the name of the trail.
	Name *string `type:"string"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files.
	S3BucketName *string `type:"string"`

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket
	// you have designated for log file delivery. For more information, see Finding
	// Your CloudTrail Log Files (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).
	S3KeyPrefix *string `type:"string"`

	// Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send notifications
	// when log files are delivered. The format of a topic ARN is:
	//
	// arn:aws:sns:us-east-2:123456789012:MyTopic
	SnsTopicARN *string `type:"string"`

	// This field is no longer in use. Use SnsTopicARN.
	SnsTopicName *string `deprecated:"true" type:"string"`

	// Specifies the ARN of the trail that was created. The format of a trail ARN
	// is:
	//
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailARN *string `type:"string"`
}

// String returns the string representation
func (s CreateTrailOutput) String() string {
	return awsutil.Prettify(s)
}
