// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSchemaInput struct {
	_ struct{} `type:"structure"`

	// The name of the schema.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the development schema. For more information,
	// see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSchemaInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The ARN that is associated with the updated schema. For more information,
	// see arns.
	SchemaArn *string `type:"string"`
}

// String returns the string representation
func (s UpdateSchemaOutput) String() string {
	return awsutil.Prettify(s)
}