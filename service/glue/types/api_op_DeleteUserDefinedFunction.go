// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteUserDefinedFunctionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the function to be deleted is located. If
	// none is supplied, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database where the function is located.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The name of the function definition to be deleted.
	//
	// FunctionName is a required field
	FunctionName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserDefinedFunctionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserDefinedFunctionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserDefinedFunctionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserDefinedFunctionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserDefinedFunctionOutput) String() string {
	return awsutil.Prettify(s)
}