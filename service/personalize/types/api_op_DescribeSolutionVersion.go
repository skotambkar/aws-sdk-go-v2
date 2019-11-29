// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSolutionVersionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the solution version.
	//
	// SolutionVersionArn is a required field
	SolutionVersionArn *string `locationName:"solutionVersionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSolutionVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSolutionVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSolutionVersionInput"}

	if s.SolutionVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SolutionVersionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSolutionVersionOutput struct {
	_ struct{} `type:"structure"`

	// The solution version.
	SolutionVersion *SolutionVersion `locationName:"solutionVersion" type:"structure"`
}

// String returns the string representation
func (s DescribeSolutionVersionOutput) String() string {
	return awsutil.Prettify(s)
}