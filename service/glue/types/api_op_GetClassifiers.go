// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetClassifiersInput struct {
	_ struct{} `type:"structure"`

	// The size of the list to return (optional).
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional continuation token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetClassifiersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetClassifiersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetClassifiersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetClassifiersOutput struct {
	_ struct{} `type:"structure"`

	// The requested list of classifier objects.
	Classifiers []Classifier `type:"list"`

	// A continuation token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetClassifiersOutput) String() string {
	return awsutil.Prettify(s)
}