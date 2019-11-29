// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateUserDefinedFunctionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which to create the function. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database in which to create the function.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A FunctionInput object that defines the function to create in the Data Catalog.
	//
	// FunctionInput is a required field
	FunctionInput *UserDefinedFunctionInput `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateUserDefinedFunctionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserDefinedFunctionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserDefinedFunctionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.FunctionInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionInput"))
	}
	if s.FunctionInput != nil {
		if err := s.FunctionInput.Validate(); err != nil {
			invalidParams.AddNested("FunctionInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateUserDefinedFunctionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateUserDefinedFunctionOutput) String() string {
	return awsutil.Prettify(s)
}