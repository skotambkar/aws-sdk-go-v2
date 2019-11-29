// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateWebhookInput struct {
	_ struct{} `type:"structure"`

	// A regular expression used to determine which repository branches are built
	// when a webhook is triggered. If the name of a branch matches the regular
	// expression, then it is built. If branchFilter is empty, then all branches
	// are built.
	//
	// It is recommended that you use filterGroups instead of branchFilter.
	BranchFilter *string `locationName:"branchFilter" type:"string"`

	// An array of arrays of WebhookFilter objects used to determine which webhooks
	// are triggered. At least one WebhookFilter in the array must specify EVENT
	// as its type.
	//
	// For a build to be triggered, at least one filter group in the filterGroups
	// array must pass. For a filter group to pass, each of its filters must pass.
	FilterGroups [][]WebhookFilter `locationName:"filterGroups" type:"list"`

	// The name of the AWS CodeBuild project.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateWebhookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWebhookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWebhookInput"}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateWebhookOutput struct {
	_ struct{} `type:"structure"`

	// Information about a webhook that connects repository events to a build project
	// in AWS CodeBuild.
	Webhook *Webhook `locationName:"webhook" type:"structure"`
}

// String returns the string representation
func (s CreateWebhookOutput) String() string {
	return awsutil.Prettify(s)
}