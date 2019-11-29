// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
)

const opListConfigurations = "ListConfigurations"

// ListConfigurationsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Retrieves a list of configuration items as specified by the value passed
// to the required paramater configurationType. Optional filtering may be applied
// to refine search results.
//
//    // Example sending a request using ListConfigurationsRequest.
//    req := client.ListConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ListConfigurations
func (c *Client) ListConfigurationsRequest(input *types.ListConfigurationsInput) ListConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListConfigurationsInput{}
	}

	req := c.newRequest(op, input, &types.ListConfigurationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListConfigurationsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListConfigurationsRequest{Request: req, Input: input, Copy: c.ListConfigurationsRequest}
}

// ListConfigurationsRequest is the request type for the
// ListConfigurations API operation.
type ListConfigurationsRequest struct {
	*aws.Request
	Input *types.ListConfigurationsInput
	Copy  func(*types.ListConfigurationsInput) ListConfigurationsRequest
}

// Send marshals and sends the ListConfigurations API request.
func (r ListConfigurationsRequest) Send(ctx context.Context) (*ListConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConfigurationsResponse{
		ListConfigurationsOutput: r.Request.Data.(*types.ListConfigurationsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListConfigurationsResponse is the response type for the
// ListConfigurations API operation.
type ListConfigurationsResponse struct {
	*types.ListConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConfigurations request.
func (r *ListConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
