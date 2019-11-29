// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PublishSchemaInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the development schema.
	// For more information, see arns.
	//
	// DevelopmentSchemaArn is a required field
	DevelopmentSchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The minor version under which the schema will be published. This parameter
	// is recommended. Schemas have both a major and minor version associated with
	// them.
	MinorVersion *string `min:"1" type:"string"`

	// The new name under which the schema will be published. If this is not provided,
	// the development schema is considered.
	Name *string `min:"1" type:"string"`

	// The major version under which the schema will be published. Schemas have
	// both a major and minor version associated with them.
	//
	// Version is a required field
	Version *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PublishSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PublishSchemaInput"}

	if s.DevelopmentSchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DevelopmentSchemaArn"))
	}
	if s.MinorVersion != nil && len(*s.MinorVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MinorVersion", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PublishSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The ARN that is associated with the published schema. For more information,
	// see arns.
	PublishedSchemaArn *string `type:"string"`
}

// String returns the string representation
func (s PublishSchemaOutput) String() string {
	return awsutil.Prettify(s)
}