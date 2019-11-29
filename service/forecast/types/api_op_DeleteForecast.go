// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteForecastInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the forecast to delete.
	//
	// ForecastArn is a required field
	ForecastArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteForecastInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteForecastInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteForecastInput"}

	if s.ForecastArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ForecastArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteForecastOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteForecastOutput) String() string {
	return awsutil.Prettify(s)
}
