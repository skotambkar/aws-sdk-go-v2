// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
)

const opDeleteOriginEndpoint = "DeleteOriginEndpoint"

// DeleteOriginEndpointRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Deletes an existing OriginEndpoint.
//
//    // Example sending a request using DeleteOriginEndpointRequest.
//    req := client.DeleteOriginEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DeleteOriginEndpoint
func (c *Client) DeleteOriginEndpointRequest(input *types.DeleteOriginEndpointInput) DeleteOriginEndpointRequest {
	op := &aws.Operation{
		Name:       opDeleteOriginEndpoint,
		HTTPMethod: "DELETE",
		HTTPPath:   "/origin_endpoints/{id}",
	}

	if input == nil {
		input = &types.DeleteOriginEndpointInput{}
	}

	req := c.newRequest(op, input, &types.DeleteOriginEndpointOutput{})
	return DeleteOriginEndpointRequest{Request: req, Input: input, Copy: c.DeleteOriginEndpointRequest}
}

// DeleteOriginEndpointRequest is the request type for the
// DeleteOriginEndpoint API operation.
type DeleteOriginEndpointRequest struct {
	*aws.Request
	Input *types.DeleteOriginEndpointInput
	Copy  func(*types.DeleteOriginEndpointInput) DeleteOriginEndpointRequest
}

// Send marshals and sends the DeleteOriginEndpoint API request.
func (r DeleteOriginEndpointRequest) Send(ctx context.Context) (*DeleteOriginEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOriginEndpointResponse{
		DeleteOriginEndpointOutput: r.Request.Data.(*types.DeleteOriginEndpointOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOriginEndpointResponse is the response type for the
// DeleteOriginEndpoint API operation.
type DeleteOriginEndpointResponse struct {
	*types.DeleteOriginEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOriginEndpoint request.
func (r *DeleteOriginEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
