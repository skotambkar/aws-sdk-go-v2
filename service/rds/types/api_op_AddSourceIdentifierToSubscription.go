// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddSourceIdentifierToSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the event source to be added.
	//
	// Constraints:
	//
	//    * If the source type is a DB instance, then a DBInstanceIdentifier must
	//    be supplied.
	//
	//    * If the source type is a DB security group, a DBSecurityGroupName must
	//    be supplied.
	//
	//    * If the source type is a DB parameter group, a DBParameterGroupName must
	//    be supplied.
	//
	//    * If the source type is a DB snapshot, a DBSnapshotIdentifier must be
	//    supplied.
	//
	// SourceIdentifier is a required field
	SourceIdentifier *string `type:"string" required:"true"`

	// The name of the RDS event notification subscription you want to add a source
	// identifier to.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddSourceIdentifierToSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddSourceIdentifierToSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddSourceIdentifierToSubscriptionInput"}

	if s.SourceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceIdentifier"))
	}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddSourceIdentifierToSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful invocation of the DescribeEventSubscriptions
	// action.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s AddSourceIdentifierToSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}
