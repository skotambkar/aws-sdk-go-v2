// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/enums"
)

type StartEntitiesDetectionV2JobInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the request. If you don't set the client request
	// token, Amazon Comprehend Medical generates one.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) role that grants Amazon Comprehend Medical read access to your input
	// data. For more information, see Role-Based Permissions Required for Asynchronous
	// Operations (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions-med.html#auth-role-permissions-med).
	//
	// DataAccessRoleArn is a required field
	DataAccessRoleArn *string `min:"20" type:"string" required:"true"`

	// Specifies the format and location of the input data for the job.
	//
	// InputDataConfig is a required field
	InputDataConfig *InputDataConfig `type:"structure" required:"true"`

	// The identifier of the job.
	JobName *string `min:"1" type:"string"`

	// An AWS Key Management Service key to encrypt your output files. If you do
	// not specify a key, the files are written in plain text.
	KMSKey *string `min:"1" type:"string"`

	// The language of the input documents. All documents must be in the same language.
	//
	// LanguageCode is a required field
	LanguageCode enums.LanguageCode `type:"string" required:"true" enum:"true"`

	// Specifies where to send the output files.
	//
	// OutputDataConfig is a required field
	OutputDataConfig *OutputDataConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartEntitiesDetectionV2JobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartEntitiesDetectionV2JobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartEntitiesDetectionV2JobInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DataAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataAccessRoleArn"))
	}
	if s.DataAccessRoleArn != nil && len(*s.DataAccessRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("DataAccessRoleArn", 20))
	}

	if s.InputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDataConfig"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}
	if s.KMSKey != nil && len(*s.KMSKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KMSKey", 1))
	}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.OutputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputDataConfig"))
	}
	if s.InputDataConfig != nil {
		if err := s.InputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputDataConfig != nil {
		if err := s.OutputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartEntitiesDetectionV2JobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the DescribeEntitiesDetectionV2Job operation.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartEntitiesDetectionV2JobOutput) String() string {
	return awsutil.Prettify(s)
}
