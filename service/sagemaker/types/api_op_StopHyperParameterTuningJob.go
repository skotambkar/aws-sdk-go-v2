// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopHyperParameterTuningJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the tuning job to stop.
	//
	// HyperParameterTuningJobName is a required field
	HyperParameterTuningJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopHyperParameterTuningJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopHyperParameterTuningJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopHyperParameterTuningJobInput"}

	if s.HyperParameterTuningJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("HyperParameterTuningJobName"))
	}
	if s.HyperParameterTuningJobName != nil && len(*s.HyperParameterTuningJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HyperParameterTuningJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopHyperParameterTuningJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopHyperParameterTuningJobOutput) String() string {
	return awsutil.Prettify(s)
}