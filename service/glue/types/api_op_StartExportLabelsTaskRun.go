// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartExportLabelsTaskRunInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 path where you export the labels.
	//
	// OutputS3Path is a required field
	OutputS3Path *string `type:"string" required:"true"`

	// The unique identifier of the machine learning transform.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartExportLabelsTaskRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartExportLabelsTaskRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartExportLabelsTaskRunInput"}

	if s.OutputS3Path == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputS3Path"))
	}

	if s.TransformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransformId"))
	}
	if s.TransformId != nil && len(*s.TransformId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransformId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartExportLabelsTaskRunOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the task run.
	TaskRunId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartExportLabelsTaskRunOutput) String() string {
	return awsutil.Prettify(s)
}