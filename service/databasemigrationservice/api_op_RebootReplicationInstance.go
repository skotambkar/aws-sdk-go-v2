// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opRebootReplicationInstance = "RebootReplicationInstance"

// RebootReplicationInstanceRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Reboots a replication instance. Rebooting results in a momentary outage,
// until the replication instance becomes available again.
//
//    // Example sending a request using RebootReplicationInstanceRequest.
//    req := client.RebootReplicationInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/RebootReplicationInstance
func (c *Client) RebootReplicationInstanceRequest(input *types.RebootReplicationInstanceInput) RebootReplicationInstanceRequest {
	op := &aws.Operation{
		Name:       opRebootReplicationInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RebootReplicationInstanceInput{}
	}

	req := c.newRequest(op, input, &types.RebootReplicationInstanceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RebootReplicationInstanceMarshaler{Input: input}.GetNamedBuildHandler())

	return RebootReplicationInstanceRequest{Request: req, Input: input, Copy: c.RebootReplicationInstanceRequest}
}

// RebootReplicationInstanceRequest is the request type for the
// RebootReplicationInstance API operation.
type RebootReplicationInstanceRequest struct {
	*aws.Request
	Input *types.RebootReplicationInstanceInput
	Copy  func(*types.RebootReplicationInstanceInput) RebootReplicationInstanceRequest
}

// Send marshals and sends the RebootReplicationInstance API request.
func (r RebootReplicationInstanceRequest) Send(ctx context.Context) (*RebootReplicationInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootReplicationInstanceResponse{
		RebootReplicationInstanceOutput: r.Request.Data.(*types.RebootReplicationInstanceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootReplicationInstanceResponse is the response type for the
// RebootReplicationInstance API operation.
type RebootReplicationInstanceResponse struct {
	*types.RebootReplicationInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootReplicationInstance request.
func (r *RebootReplicationInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
