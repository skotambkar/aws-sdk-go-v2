// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSecretInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the secret whose details you want to retrieve. You can
	// specify either the Amazon Resource Name (ARN) or the friendly name of the
	// secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// that end with a hyphen followed by six characters.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSecretInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSecretInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSecretInput"}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSecretOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret.
	ARN *string `min:"20" type:"string"`

	// This value exists if the secret is scheduled for deletion. Some time after
	// the specified date and time, Secrets Manager deletes the secret and all of
	// its versions.
	//
	// If a secret is scheduled for deletion, then its details, including the encrypted
	// secret information, is not accessible. To cancel a scheduled deletion and
	// restore access, use RestoreSecret.
	DeletedDate *time.Time `type:"timestamp"`

	// The user-provided description of the secret.
	Description *string `type:"string"`

	// The ARN or alias of the AWS KMS customer master key (CMK) that's used to
	// encrypt the SecretString or SecretBinary fields in each version of the secret.
	// If you don't provide a key, then Secrets Manager defaults to encrypting the
	// secret fields with the default AWS KMS CMK (the one named awssecretsmanager)
	// for this account.
	KmsKeyId *string `type:"string"`

	// The last date that this secret was accessed. This value is truncated to midnight
	// of the date and therefore shows only the date, not the time.
	LastAccessedDate *time.Time `type:"timestamp"`

	// The last date and time that this secret was modified in any way.
	LastChangedDate *time.Time `type:"timestamp"`

	// The most recent date and time that the Secrets Manager rotation process was
	// successfully completed. This value is null if the secret has never rotated.
	LastRotatedDate *time.Time `type:"timestamp"`

	// The user-provided friendly name of the secret.
	Name *string `min:"1" type:"string"`

	OwningService *string `min:"1" type:"string"`

	// Specifies whether automatic rotation is enabled for this secret.
	//
	// To enable rotation, use RotateSecret with AutomaticallyRotateAfterDays set
	// to a value greater than 0. To disable rotation, use CancelRotateSecret.
	RotationEnabled *bool `type:"boolean"`

	// The ARN of a Lambda function that's invoked by Secrets Manager to rotate
	// the secret either automatically per the schedule or manually by a call to
	// RotateSecret.
	RotationLambdaARN *string `type:"string"`

	// A structure that contains the rotation configuration for this secret.
	RotationRules *RotationRulesType `type:"structure"`

	// The list of user-defined tags that are associated with the secret. To add
	// tags to a secret, use TagResource. To remove tags, use UntagResource.
	Tags []Tag `type:"list"`

	// A list of all of the currently assigned VersionStage staging labels and the
	// VersionId that each is attached to. Staging labels are used to keep track
	// of the different versions during the rotation process.
	//
	// A version that does not have any staging labels attached is considered deprecated
	// and subject to deletion. Such versions are not included in this list.
	VersionIdsToStages map[string][]string `type:"map"`
}

// String returns the string representation
func (s DescribeSecretOutput) String() string {
	return awsutil.Prettify(s)
}
