// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateEvaluationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the DataSource for the evaluation. The schema of the DataSource
	// must match the schema used to create the MLModel.
	//
	// EvaluationDataSourceId is a required field
	EvaluationDataSourceId *string `min:"1" type:"string" required:"true"`

	// A user-supplied ID that uniquely identifies the Evaluation.
	//
	// EvaluationId is a required field
	EvaluationId *string `min:"1" type:"string" required:"true"`

	// A user-supplied name or description of the Evaluation.
	EvaluationName *string `type:"string"`

	// The ID of the MLModel to evaluate.
	//
	// The schema used in creating the MLModel must match the schema of the DataSource
	// used in the Evaluation.
	//
	// MLModelId is a required field
	MLModelId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateEvaluationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEvaluationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEvaluationInput"}

	if s.EvaluationDataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EvaluationDataSourceId"))
	}
	if s.EvaluationDataSourceId != nil && len(*s.EvaluationDataSourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EvaluationDataSourceId", 1))
	}

	if s.EvaluationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EvaluationId"))
	}
	if s.EvaluationId != nil && len(*s.EvaluationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EvaluationId", 1))
	}

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

// Represents the output of a CreateEvaluation operation, and is an acknowledgement
// that Amazon ML received the request.
//
// CreateEvaluation operation is asynchronous. You can poll for status updates
// by using the GetEvcaluation operation and checking the Status parameter.
type CreateEvaluationOutput struct {
	_ struct{} `type:"structure"`

	// The user-supplied ID that uniquely identifies the Evaluation. This value
	// should be identical to the value of the EvaluationId in the request.
	EvaluationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateEvaluationOutput) String() string {
	return awsutil.Prettify(s)
}
