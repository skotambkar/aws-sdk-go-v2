// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/enums"
)

type DescribeLabelingJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the labeling job to return information for.
	//
	// LabelingJobName is a required field
	LabelingJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLabelingJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLabelingJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLabelingJobInput"}

	if s.LabelingJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LabelingJobName"))
	}
	if s.LabelingJobName != nil && len(*s.LabelingJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LabelingJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLabelingJobOutput struct {
	_ struct{} `type:"structure"`

	// The date and time that the labeling job was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// If the job failed, the reason that it failed.
	FailureReason *string `type:"string"`

	// Configuration information required for human workers to complete a labeling
	// task.
	//
	// HumanTaskConfig is a required field
	HumanTaskConfig *HumanTaskConfig `type:"structure" required:"true"`

	// Input configuration information for the labeling job, such as the Amazon
	// S3 location of the data objects and the location of the manifest file that
	// describes the data objects.
	//
	// InputConfig is a required field
	InputConfig *LabelingJobInputConfig `type:"structure" required:"true"`

	// A unique identifier for work done as part of a labeling job.
	//
	// JobReferenceCode is a required field
	JobReferenceCode *string `min:"1" type:"string" required:"true"`

	// The attribute used as the label in the output manifest file.
	LabelAttributeName *string `min:"1" type:"string"`

	// The S3 location of the JSON file that defines the categories used to label
	// data objects.
	//
	// The file is a JSON structure in the following format:
	//
	// {
	//
	// "document-version": "2018-11-28"
	//
	// "labels": [
	//
	// {
	//
	// "label": "label 1"
	//
	// },
	//
	// {
	//
	// "label": "label 2"
	//
	// },
	//
	// ...
	//
	// {
	//
	// "label": "label n"
	//
	// }
	//
	// ]
	//
	// }
	LabelCategoryConfigS3Uri *string `type:"string"`

	// Provides a breakdown of the number of data objects labeled by humans, the
	// number of objects labeled by machine, the number of objects than couldn't
	// be labeled, and the total number of objects labeled.
	//
	// LabelCounters is a required field
	LabelCounters *LabelCounters `type:"structure" required:"true"`

	// Configuration information for automated data labeling.
	LabelingJobAlgorithmsConfig *LabelingJobAlgorithmsConfig `type:"structure"`

	// The Amazon Resource Name (ARN) of the labeling job.
	//
	// LabelingJobArn is a required field
	LabelingJobArn *string `type:"string" required:"true"`

	// The name assigned to the labeling job when it was created.
	//
	// LabelingJobName is a required field
	LabelingJobName *string `min:"1" type:"string" required:"true"`

	// The location of the output produced by the labeling job.
	LabelingJobOutput *LabelingJobOutput `type:"structure"`

	// The processing status of the labeling job.
	//
	// LabelingJobStatus is a required field
	LabelingJobStatus enums.LabelingJobStatus `type:"string" required:"true" enum:"true"`

	// The date and time that the labeling job was last updated.
	//
	// LastModifiedTime is a required field
	LastModifiedTime *time.Time `type:"timestamp" required:"true"`

	// The location of the job's output data and the AWS Key Management Service
	// key ID for the key used to encrypt the output data, if any.
	//
	// OutputConfig is a required field
	OutputConfig *LabelingJobOutputConfig `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) that Amazon SageMaker assumes to perform tasks
	// on your behalf during data labeling.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// A set of conditions for stopping a labeling job. If any of the conditions
	// are met, the job is automatically stopped.
	StoppingConditions *LabelingJobStoppingConditions `type:"structure"`

	// An array of key/value pairs. For more information, see Using Cost Allocation
	// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s DescribeLabelingJobOutput) String() string {
	return awsutil.Prettify(s)
}
