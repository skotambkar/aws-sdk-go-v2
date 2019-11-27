// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDeleteEndpoint = "DeleteEndpoint"

// DeleteEndpointRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Deletes the specified endpoint.
//
// All tasks associated with the endpoint must be deleted before you can delete
// the endpoint.
//
//    // Example sending a request using DeleteEndpointRequest.
//    req := client.DeleteEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteEndpoint
func (c *Client) DeleteEndpointRequest(input *types.DeleteEndpointInput) DeleteEndpointRequest {
	op := &aws.Operation{
		Name:       opDeleteEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteEndpointInput{}
	}

	req := c.newRequest(op, input, &types.DeleteEndpointOutput{})
	return DeleteEndpointRequest{Request: req, Input: input, Copy: c.DeleteEndpointRequest}
}

// DeleteEndpointRequest is the request type for the
// DeleteEndpoint API operation.
type DeleteEndpointRequest struct {
	*aws.Request
	Input *types.DeleteEndpointInput
	Copy  func(*types.DeleteEndpointInput) DeleteEndpointRequest
}

// Send marshals and sends the DeleteEndpoint API request.
func (r DeleteEndpointRequest) Send(ctx context.Context) (*DeleteEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEndpointResponse{
		DeleteEndpointOutput: r.Request.Data.(*types.DeleteEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEndpointResponse is the response type for the
// DeleteEndpoint API operation.
type DeleteEndpointResponse struct {
	*types.DeleteEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEndpoint request.
func (r *DeleteEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
