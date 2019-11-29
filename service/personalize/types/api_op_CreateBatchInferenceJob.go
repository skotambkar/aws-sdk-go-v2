// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateBatchInferenceJobInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 path that leads to the input file to base your recommendations
	// on. The input material must be in JSON format.
	//
	// JobInput is a required field
	JobInput *BatchInferenceJobInput `locationName:"jobInput" type:"structure" required:"true"`

	// The name of the batch inference job to create.
	//
	// JobName is a required field
	JobName *string `locationName:"jobName" min:"1" type:"string" required:"true"`

	// The path to the Amazon S3 bucket where the job's output will be stored.
	//
	// JobOutput is a required field
	JobOutput *BatchInferenceJobOutput `locationName:"jobOutput" type:"structure" required:"true"`

	// The number of recommendations to retreive.
	NumResults *int64 `locationName:"numResults" type:"integer"`

	// The ARN of the Amazon Identity and Access Management role that has permissions
	// to read and write to your input and out Amazon S3 buckets respectively.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the solution version that will be used
	// to generate the batch inference recommendations.
	//
	// SolutionVersionArn is a required field
	SolutionVersionArn *string `locationName:"solutionVersionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateBatchInferenceJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBatchInferenceJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBatchInferenceJobInput"}

	if s.JobInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobInput"))
	}

	if s.JobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobName"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}

	if s.JobOutput == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobOutput"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if s.SolutionVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SolutionVersionArn"))
	}
	if s.JobInput != nil {
		if err := s.JobInput.Validate(); err != nil {
			invalidParams.AddNested("JobInput", err.(aws.ErrInvalidParams))
		}
	}
	if s.JobOutput != nil {
		if err := s.JobOutput.Validate(); err != nil {
			invalidParams.AddNested("JobOutput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateBatchInferenceJobOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the batch inference job.
	BatchInferenceJobArn *string `locationName:"batchInferenceJobArn" type:"string"`
}

// String returns the string representation
func (s CreateBatchInferenceJobOutput) String() string {
	return awsutil.Prettify(s)
}