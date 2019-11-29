// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sms/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
)

const opCreateReplicationJob = "CreateReplicationJob"

// CreateReplicationJobRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Creates a replication job. The replication job schedules periodic replication
// runs to replicate your server to AWS. Each replication run creates an Amazon
// Machine Image (AMI).
//
//    // Example sending a request using CreateReplicationJobRequest.
//    req := client.CreateReplicationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/CreateReplicationJob
func (c *Client) CreateReplicationJobRequest(input *types.CreateReplicationJobInput) CreateReplicationJobRequest {
	op := &aws.Operation{
		Name:       opCreateReplicationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateReplicationJobInput{}
	}

	req := c.newRequest(op, input, &types.CreateReplicationJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateReplicationJobMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateReplicationJobRequest{Request: req, Input: input, Copy: c.CreateReplicationJobRequest}
}

// CreateReplicationJobRequest is the request type for the
// CreateReplicationJob API operation.
type CreateReplicationJobRequest struct {
	*aws.Request
	Input *types.CreateReplicationJobInput
	Copy  func(*types.CreateReplicationJobInput) CreateReplicationJobRequest
}

// Send marshals and sends the CreateReplicationJob API request.
func (r CreateReplicationJobRequest) Send(ctx context.Context) (*CreateReplicationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReplicationJobResponse{
		CreateReplicationJobOutput: r.Request.Data.(*types.CreateReplicationJobOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReplicationJobResponse is the response type for the
// CreateReplicationJob API operation.
type CreateReplicationJobResponse struct {
	*types.CreateReplicationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReplicationJob request.
func (r *CreateReplicationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
