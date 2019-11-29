// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartImportTaskInput struct {
	_ struct{} `type:"structure"`

	// Optional. A unique token that you can provide to prevent the same import
	// request from occurring more than once. If you don't provide a token, a token
	// is automatically generated.
	//
	// Sending more than one StartImportTask request with the same client request
	// token will return information about the original import task with that client
	// request token.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string" idempotencyToken:"true"`

	// The URL for your import file that you've uploaded to Amazon S3.
	//
	// If you're using the AWS CLI, this URL is structured as follows: s3://BucketName/ImportFileName.CSV
	//
	// ImportUrl is a required field
	ImportUrl *string `locationName:"importUrl" min:"1" type:"string" required:"true"`

	// A descriptive name for this request. You can use this name to filter future
	// requests related to this import task, such as identifying applications and
	// servers that were included in this import task. We recommend that you use
	// a meaningful name for each import task.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartImportTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartImportTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartImportTaskInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.ImportUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImportUrl"))
	}
	if s.ImportUrl != nil && len(*s.ImportUrl) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImportUrl", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartImportTaskOutput struct {
	_ struct{} `type:"structure"`

	// An array of information related to the import task request including status
	// information, times, IDs, the Amazon S3 Object URL for the import file, and
	// more.
	Task *ImportTask `locationName:"task" type:"structure"`
}

// String returns the string representation
func (s StartImportTaskOutput) String() string {
	return awsutil.Prettify(s)
}