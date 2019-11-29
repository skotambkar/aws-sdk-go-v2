// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which the connection resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// A ConnectionInput object that redefines the connection in question.
	//
	// ConnectionInput is a required field
	ConnectionInput *ConnectionInput `type:"structure" required:"true"`

	// The name of the connection definition to update.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConnectionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.ConnectionInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionInput"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.ConnectionInput != nil {
		if err := s.ConnectionInput.Validate(); err != nil {
			invalidParams.AddNested("ConnectionInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConnectionOutput) String() string {
	return awsutil.Prettify(s)
}