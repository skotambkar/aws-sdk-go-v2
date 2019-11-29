// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSecurityConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the security configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The security configuration details in JSON format. For JSON parameters and
	// examples, see Use Security Configurations to Set Up Cluster Security (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-security-configurations.html)
	// in the Amazon EMR Management Guide.
	//
	// SecurityConfiguration is a required field
	SecurityConfiguration *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSecurityConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSecurityConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSecurityConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SecurityConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSecurityConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The date and time the security configuration was created.
	//
	// CreationDateTime is a required field
	CreationDateTime *time.Time `type:"timestamp" required:"true"`

	// The name of the security configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSecurityConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
