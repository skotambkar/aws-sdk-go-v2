// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteEvaluationInput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the Evaluation to delete.
	//
	// EvaluationId is a required field
	EvaluationId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEvaluationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEvaluationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEvaluationInput"}

	if s.EvaluationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EvaluationId"))
	}
	if s.EvaluationId != nil && len(*s.EvaluationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EvaluationId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DeleteEvaluation operation. The output indicates
// that Amazon Machine Learning (Amazon ML) received the request.
//
// You can use the GetEvaluation operation and check the value of the Status
// parameter to see whether an Evaluation is marked as DELETED.
type DeleteEvaluationOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the Evaluation. This value should
	// be identical to the value of the EvaluationId in the request.
	EvaluationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteEvaluationOutput) String() string {
	return awsutil.Prettify(s)
}
