// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartDataSourceSyncJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the data source to synchronize.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The identifier of the index that contains the data source.
	//
	// IndexId is a required field
	IndexId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s StartDataSourceSyncJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDataSourceSyncJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDataSourceSyncJobInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.IndexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexId"))
	}
	if s.IndexId != nil && len(*s.IndexId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartDataSourceSyncJobOutput struct {
	_ struct{} `type:"structure"`

	// Identifies a particular synchronization job.
	ExecutionId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartDataSourceSyncJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartDataSourceSyncJob = "StartDataSourceSyncJob"

// StartDataSourceSyncJobRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Starts a synchronization job for a data source. If a synchronization job
// is already in progress, Amazon Kendra returns a ResourceInUseException exception.
//
//    // Example sending a request using StartDataSourceSyncJobRequest.
//    req := client.StartDataSourceSyncJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/StartDataSourceSyncJob
func (c *Client) StartDataSourceSyncJobRequest(input *StartDataSourceSyncJobInput) StartDataSourceSyncJobRequest {
	op := &aws.Operation{
		Name:       opStartDataSourceSyncJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDataSourceSyncJobInput{}
	}

	req := c.newRequest(op, input, &StartDataSourceSyncJobOutput{})
	return StartDataSourceSyncJobRequest{Request: req, Input: input, Copy: c.StartDataSourceSyncJobRequest}
}

// StartDataSourceSyncJobRequest is the request type for the
// StartDataSourceSyncJob API operation.
type StartDataSourceSyncJobRequest struct {
	*aws.Request
	Input *StartDataSourceSyncJobInput
	Copy  func(*StartDataSourceSyncJobInput) StartDataSourceSyncJobRequest
}

// Send marshals and sends the StartDataSourceSyncJob API request.
func (r StartDataSourceSyncJobRequest) Send(ctx context.Context) (*StartDataSourceSyncJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDataSourceSyncJobResponse{
		StartDataSourceSyncJobOutput: r.Request.Data.(*StartDataSourceSyncJobOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDataSourceSyncJobResponse is the response type for the
// StartDataSourceSyncJob API operation.
type StartDataSourceSyncJobResponse struct {
	*StartDataSourceSyncJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDataSourceSyncJob request.
func (r *StartDataSourceSyncJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
