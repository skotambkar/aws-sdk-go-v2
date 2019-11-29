// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateLogGroupInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	// For more information, see Amazon Resource Names - AWS Key Management Service
	// (AWS KMS) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms).
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The key-value pairs to use for the tags.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateLogGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLogGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLogGroupInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLogGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLogGroupOutput) String() string {
	return awsutil.Prettify(s)
}
