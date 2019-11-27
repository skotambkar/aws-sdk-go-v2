// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Result message containing a list of application version descriptions.
type DescribeConfigurationOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the application associated with the configuration template or
	// environment. Only needed if you want to describe the configuration options
	// associated with either the configuration template or environment.
	ApplicationName *string `min:"1" type:"string"`

	// The name of the environment whose configuration options you want to describe.
	EnvironmentName *string `min:"4" type:"string"`

	// If specified, restricts the descriptions to only the specified options.
	Options []OptionSpecification `type:"list"`

	// The ARN of the custom platform.
	PlatformArn *string `type:"string"`

	// The name of the solution stack whose configuration options you want to describe.
	SolutionStackName *string `type:"string"`

	// The name of the configuration template whose configuration options you want
	// to describe.
	TemplateName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeConfigurationOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConfigurationOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConfigurationOptionsInput"}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}
	if s.Options != nil {
		for i, v := range s.Options {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Options", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the settings for a specified configuration set.
type DescribeConfigurationOptionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of ConfigurationOptionDescription.
	Options []ConfigurationOptionDescription `type:"list"`

	// The ARN of the platform.
	PlatformArn *string `type:"string"`

	// The name of the solution stack these configuration options belong to.
	SolutionStackName *string `type:"string"`
}

// String returns the string representation
func (s DescribeConfigurationOptionsOutput) String() string {
	return awsutil.Prettify(s)
}
