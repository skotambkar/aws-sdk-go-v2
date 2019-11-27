// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetUserDefinedFunctionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the functions to be retrieved are located.
	// If none is provided, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database where the functions are located.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The maximum number of functions to return in one response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`

	// An optional function-name pattern string that filters the function definitions
	// returned.
	//
	// Pattern is a required field
	Pattern *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserDefinedFunctionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserDefinedFunctionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserDefinedFunctionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Pattern == nil {
		invalidParams.Add(aws.NewErrParamRequired("Pattern"))
	}
	if s.Pattern != nil && len(*s.Pattern) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Pattern", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetUserDefinedFunctionsOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if the list of functions returned does not include
	// the last requested function.
	NextToken *string `type:"string"`

	// A list of requested function definitions.
	UserDefinedFunctions []UserDefinedFunction `type:"list"`
}

// String returns the string representation
func (s GetUserDefinedFunctionsOutput) String() string {
	return awsutil.Prettify(s)
}
