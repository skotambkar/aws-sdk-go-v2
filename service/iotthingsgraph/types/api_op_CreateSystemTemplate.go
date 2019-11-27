// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSystemTemplateInput struct {
	_ struct{} `type:"structure"`

	// The namespace version in which the system is to be created.
	//
	// If no value is specified, the latest version is used by default.
	CompatibleNamespaceVersion *int64 `locationName:"compatibleNamespaceVersion" type:"long"`

	// The DefinitionDocument used to create the system.
	//
	// Definition is a required field
	Definition *DefinitionDocument `locationName:"definition" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateSystemTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSystemTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSystemTemplateInput"}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			invalidParams.AddNested("Definition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSystemTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The summary object that describes the created system.
	Summary *SystemTemplateSummary `locationName:"summary" type:"structure"`
}

// String returns the string representation
func (s CreateSystemTemplateOutput) String() string {
	return awsutil.Prettify(s)
}
