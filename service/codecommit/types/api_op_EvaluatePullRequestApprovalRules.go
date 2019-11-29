// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EvaluatePullRequestApprovalRulesInput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the pull request you want to evaluate.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The system-generated ID for the pull request revision. To retrieve the most
	// recent revision ID for a pull request, use GetPullRequest.
	//
	// RevisionId is a required field
	RevisionId *string `locationName:"revisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s EvaluatePullRequestApprovalRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EvaluatePullRequestApprovalRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EvaluatePullRequestApprovalRulesInput"}

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

type EvaluatePullRequestApprovalRulesOutput struct {
	_ struct{} `type:"structure"`

	// The result of the evaluation, including the names of the rules whose conditions
	// have been met (if any), the names of the rules whose conditions have not
	// been met (if any), whether the pull request is in the approved state, and
	// whether the pull request approval rule has been set aside by an override.
	//
	// Evaluation is a required field
	Evaluation *Evaluation `locationName:"evaluation" type:"structure" required:"true"`
}

// String returns the string representation
func (s EvaluatePullRequestApprovalRulesOutput) String() string {
	return awsutil.Prettify(s)
}