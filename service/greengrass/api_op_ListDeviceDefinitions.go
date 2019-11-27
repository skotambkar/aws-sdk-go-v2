// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opListDeviceDefinitions = "ListDeviceDefinitions"

// ListDeviceDefinitionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of device definitions.
//
//    // Example sending a request using ListDeviceDefinitionsRequest.
//    req := client.ListDeviceDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeviceDefinitions
func (c *Client) ListDeviceDefinitionsRequest(input *types.ListDeviceDefinitionsInput) ListDeviceDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListDeviceDefinitions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/devices",
	}

	if input == nil {
		input = &types.ListDeviceDefinitionsInput{}
	}

	req := c.newRequest(op, input, &types.ListDeviceDefinitionsOutput{})
	return ListDeviceDefinitionsRequest{Request: req, Input: input, Copy: c.ListDeviceDefinitionsRequest}
}

// ListDeviceDefinitionsRequest is the request type for the
// ListDeviceDefinitions API operation.
type ListDeviceDefinitionsRequest struct {
	*aws.Request
	Input *types.ListDeviceDefinitionsInput
	Copy  func(*types.ListDeviceDefinitionsInput) ListDeviceDefinitionsRequest
}

// Send marshals and sends the ListDeviceDefinitions API request.
func (r ListDeviceDefinitionsRequest) Send(ctx context.Context) (*ListDeviceDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeviceDefinitionsResponse{
		ListDeviceDefinitionsOutput: r.Request.Data.(*types.ListDeviceDefinitionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDeviceDefinitionsResponse is the response type for the
// ListDeviceDefinitions API operation.
type ListDeviceDefinitionsResponse struct {
	*types.ListDeviceDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeviceDefinitions request.
func (r *ListDeviceDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
