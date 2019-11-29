// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDirectoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the Directory. Should be unique per account, per region.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the published schema that will be copied
	// into the data Directory. For more information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDirectoryInput"}

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

type CreateDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the published schema in the Directory. Once a published schema
	// is copied into the directory, it has its own ARN, which is referred to applied
	// schema ARN. For more information, see arns.
	//
	// AppliedSchemaArn is a required field
	AppliedSchemaArn *string `type:"string" required:"true"`

	// The ARN that is associated with the Directory. For more information, see
	// arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `type:"string" required:"true"`

	// The name of the Directory.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The root object node of the created directory.
	//
	// ObjectIdentifier is a required field
	ObjectIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}