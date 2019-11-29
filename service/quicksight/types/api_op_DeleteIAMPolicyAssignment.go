// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteIAMPolicyAssignmentInput struct {
	_ struct{} `type:"structure"`

	// The name of the assignment.
	//
	// AssignmentName is a required field
	AssignmentName *string `location:"uri" locationName:"AssignmentName" min:"1" type:"string" required:"true"`

	// The AWS account ID where you want to delete an IAM policy assignment.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The namespace that contains the assignment.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIAMPolicyAssignmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIAMPolicyAssignmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIAMPolicyAssignmentInput"}

	if s.AssignmentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentName"))
	}
	if s.AssignmentName != nil && len(*s.AssignmentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssignmentName", 1))
	}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteIAMPolicyAssignmentOutput struct {
	_ struct{} `type:"structure"`

	// The name of the assignment.
	AssignmentName *string `min:"1" type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DeleteIAMPolicyAssignmentOutput) String() string {
	return awsutil.Prettify(s)
}