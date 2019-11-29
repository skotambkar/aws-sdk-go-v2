// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/emr/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
)

const opAddInstanceFleet = "AddInstanceFleet"

// AddInstanceFleetRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Adds an instance fleet to a running cluster.
//
// The instance fleet configuration is available only in Amazon EMR versions
// 4.8.0 and later, excluding 5.0.x.
//
//    // Example sending a request using AddInstanceFleetRequest.
//    req := client.AddInstanceFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/AddInstanceFleet
func (c *Client) AddInstanceFleetRequest(input *types.AddInstanceFleetInput) AddInstanceFleetRequest {
	op := &aws.Operation{
		Name:       opAddInstanceFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AddInstanceFleetInput{}
	}

	req := c.newRequest(op, input, &types.AddInstanceFleetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AddInstanceFleetMarshaler{Input: input}.GetNamedBuildHandler())

	return AddInstanceFleetRequest{Request: req, Input: input, Copy: c.AddInstanceFleetRequest}
}

// AddInstanceFleetRequest is the request type for the
// AddInstanceFleet API operation.
type AddInstanceFleetRequest struct {
	*aws.Request
	Input *types.AddInstanceFleetInput
	Copy  func(*types.AddInstanceFleetInput) AddInstanceFleetRequest
}

// Send marshals and sends the AddInstanceFleet API request.
func (r AddInstanceFleetRequest) Send(ctx context.Context) (*AddInstanceFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddInstanceFleetResponse{
		AddInstanceFleetOutput: r.Request.Data.(*types.AddInstanceFleetOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddInstanceFleetResponse is the response type for the
// AddInstanceFleet API operation.
type AddInstanceFleetResponse struct {
	*types.AddInstanceFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddInstanceFleet request.
func (r *AddInstanceFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
