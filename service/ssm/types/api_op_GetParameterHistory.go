// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetParameterHistoryInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The name of a parameter you want to query.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// Return decrypted values for secure string parameters. This flag is ignored
	// for String and StringList parameter types.
	WithDecryption *bool `type:"boolean"`
}

// String returns the string representation
func (s GetParameterHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetParameterHistoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetParameterHistoryInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetParameterHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`

	// A list of parameters returned by the request.
	Parameters []ParameterHistory `type:"list"`
}

// String returns the string representation
func (s GetParameterHistoryOutput) String() string {
	return awsutil.Prettify(s)
}
