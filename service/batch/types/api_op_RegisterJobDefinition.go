// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/batch/enums"
)

type RegisterJobDefinitionInput struct {
	_ struct{} `type:"structure"`

	// An object with various properties specific to single-node container-based
	// jobs. If the job definition's type parameter is container, then you must
	// specify either containerProperties or nodeProperties.
	ContainerProperties *ContainerProperties `locationName:"containerProperties" type:"structure"`

	// The name of the job definition to register. Up to 128 letters (uppercase
	// and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// JobDefinitionName is a required field
	JobDefinitionName *string `locationName:"jobDefinitionName" type:"string" required:"true"`

	// An object with various properties specific to multi-node parallel jobs. If
	// you specify node properties for a job, it becomes a multi-node parallel job.
	// For more information, see Multi-node Parallel Jobs (https://docs.aws.amazon.com/batch/latest/userguide/multi-node-parallel-jobs.html)
	// in the AWS Batch User Guide. If the job definition's type parameter is container,
	// then you must specify either containerProperties or nodeProperties.
	NodeProperties *NodeProperties `locationName:"nodeProperties" type:"structure"`

	// Default parameter substitution placeholders to set in the job definition.
	// Parameters are specified as a key-value pair mapping. Parameters in a SubmitJob
	// request override any corresponding parameter defaults from the job definition.
	Parameters map[string]string `locationName:"parameters" type:"map"`

	// The retry strategy to use for failed jobs that are submitted with this job
	// definition. Any retry strategy that is specified during a SubmitJob operation
	// overrides the retry strategy defined here. If a job is terminated due to
	// a timeout, it is not retried.
	RetryStrategy *RetryStrategy `locationName:"retryStrategy" type:"structure"`

	// The timeout configuration for jobs that are submitted with this job definition,
	// after which AWS Batch terminates your jobs if they have not finished. If
	// a job is terminated due to a timeout, it is not retried. The minimum value
	// for the timeout is 60 seconds. Any timeout configuration that is specified
	// during a SubmitJob operation overrides the timeout configuration defined
	// here. For more information, see Job Timeouts (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/job_timeouts.html)
	// in the Amazon Elastic Container Service Developer Guide.
	Timeout *JobTimeout `locationName:"timeout" type:"structure"`

	// The type of job definition.
	//
	// Type is a required field
	Type enums.JobDefinitionType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s RegisterJobDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterJobDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterJobDefinitionInput"}

	if s.JobDefinitionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobDefinitionName"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.ContainerProperties != nil {
		if err := s.ContainerProperties.Validate(); err != nil {
			invalidParams.AddNested("ContainerProperties", err.(aws.ErrInvalidParams))
		}
	}
	if s.NodeProperties != nil {
		if err := s.NodeProperties.Validate(); err != nil {
			invalidParams.AddNested("NodeProperties", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterJobDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the job definition.
	//
	// JobDefinitionArn is a required field
	JobDefinitionArn *string `locationName:"jobDefinitionArn" type:"string" required:"true"`

	// The name of the job definition.
	//
	// JobDefinitionName is a required field
	JobDefinitionName *string `locationName:"jobDefinitionName" type:"string" required:"true"`

	// The revision of the job definition.
	//
	// Revision is a required field
	Revision *int64 `locationName:"revision" type:"integer" required:"true"`
}

// String returns the string representation
func (s RegisterJobDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}
