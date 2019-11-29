// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartCrawlerInput struct {
	_ struct{} `type:"structure"`

	// Name of the crawler to start.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartCrawlerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartCrawlerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartCrawlerInput"}

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

type StartCrawlerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartCrawlerOutput) String() string {
	return awsutil.Prettify(s)
}