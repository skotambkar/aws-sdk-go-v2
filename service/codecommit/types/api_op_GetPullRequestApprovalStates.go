// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetPullRequestApprovalStatesInput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID for the pull request.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The system-generated ID for the pull request revision.
	//
	// RevisionId is a required field
	RevisionId *string `locationName:"revisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPullRequestApprovalStatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPullRequestApprovalStatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPullRequestApprovalStatesInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if s.RevisionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RevisionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetPullRequestApprovalStatesOutput struct {
	_ struct{} `type:"structure"`

	// Information about users who have approved the pull request.
	Approvals []Approval `locationName:"approvals" type:"list"`
}

// String returns the string representation
func (s GetPullRequestApprovalStatesOutput) String() string {
	return awsutil.Prettify(s)
}
