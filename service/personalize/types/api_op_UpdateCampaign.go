// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateCampaignInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the campaign.
	//
	// CampaignArn is a required field
	CampaignArn *string `locationName:"campaignArn" type:"string" required:"true"`

	// Specifies the requested minimum provisioned transactions (recommendations)
	// per second that Amazon Personalize will support.
	MinProvisionedTPS *int64 `locationName:"minProvisionedTPS" min:"1" type:"integer"`

	// The ARN of a new solution version to deploy.
	SolutionVersionArn *string `locationName:"solutionVersionArn" type:"string"`
}

// String returns the string representation
func (s UpdateCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCampaignInput"}

	if s.CampaignArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignArn"))
	}
	if s.MinProvisionedTPS != nil && *s.MinProvisionedTPS < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MinProvisionedTPS", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCampaignOutput struct {
	_ struct{} `type:"structure"`

	// The same campaign ARN as given in the request.
	CampaignArn *string `locationName:"campaignArn" type:"string"`
}

// String returns the string representation
func (s UpdateCampaignOutput) String() string {
	return awsutil.Prettify(s)
}