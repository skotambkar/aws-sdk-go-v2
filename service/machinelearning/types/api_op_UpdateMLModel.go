// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateMLModelInput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the MLModel during creation.
	//
	// MLModelId is a required field
	MLModelId *string `min:"1" type:"string" required:"true"`

	// A user-supplied name or description of the MLModel.
	MLModelName *string `type:"string"`

	// The ScoreThreshold used in binary classification MLModel that marks the boundary
	// between a positive prediction and a negative prediction.
	//
	// Output values greater than or equal to the ScoreThreshold receive a positive
	// result from the MLModel, such as true. Output values less than the ScoreThreshold
	// receive a negative response from the MLModel, such as false.
	ScoreThreshold *float64 `type:"float"`
}

// String returns the string representation
func (s UpdateMLModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMLModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMLModelInput"}

	if s.MLModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MLModelId"))
	}
	if s.MLModelId != nil && len(*s.MLModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MLModelId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an UpdateMLModel operation.
//
// You can see the updated content by using the GetMLModel operation.
type UpdateMLModelOutput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the MLModel during creation. This value should be identical
	// to the value of the MLModelID in the request.
	MLModelId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateMLModelOutput) String() string {
	return awsutil.Prettify(s)
}