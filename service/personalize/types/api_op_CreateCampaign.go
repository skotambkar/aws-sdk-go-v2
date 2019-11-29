// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateCampaignInput struct {
	_ struct{} `type:"structure"`

	// Specifies the requested minimum provisioned transactions (recommendations)
	// per second that Amazon Personalize will support.
	//
	// MinProvisionedTPS is a required field
	MinProvisionedTPS *int64 `locationName:"minProvisionedTPS" min:"1" type:"integer" required:"true"`

	// A name for the new campaign. The campaign name must be unique within your
	// account.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the solution version to deploy.
	//
	// SolutionVersionArn is a required field
	SolutionVersionArn *string `locationName:"solutionVersionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCampaignInput"}

	if s.MinProvisionedTPS == nil {
		invalidParams.Add(aws.NewErrParamRequired("MinProvisionedTPS"))
	}
	if s.MinProvisionedTPS != nil && *s.MinProvisionedTPS < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MinProvisionedTPS", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SolutionVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SolutionVersionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCampaignOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the campaign.
	CampaignArn *string `locationName:"campaignArn" type:"string"`
}

// String returns the string representation
func (s CreateCampaignOutput) String() string {
	return awsutil.Prettify(s)
}