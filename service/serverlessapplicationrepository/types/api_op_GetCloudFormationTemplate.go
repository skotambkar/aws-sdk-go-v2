// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/enums"
)

type GetCloudFormationTemplateInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"templateId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCloudFormationTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCloudFormationTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCloudFormationTemplateInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCloudFormationTemplateOutput struct {
	_ struct{} `type:"structure"`

	ApplicationId *string `locationName:"applicationId" type:"string"`

	CreationTime *string `locationName:"creationTime" type:"string"`

	ExpirationTime *string `locationName:"expirationTime" type:"string"`

	SemanticVersion *string `locationName:"semanticVersion" type:"string"`

	Status enums.Status `locationName:"status" type:"string" enum:"true"`

	TemplateId *string `locationName:"templateId" type:"string"`

	TemplateUrl *string `locationName:"templateUrl" type:"string"`
}

// String returns the string representation
func (s GetCloudFormationTemplateOutput) String() string {
	return awsutil.Prettify(s)
}