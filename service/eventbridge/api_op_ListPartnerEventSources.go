// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

const opListPartnerEventSources = "ListPartnerEventSources"

// ListPartnerEventSourcesRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// An SaaS partner can use this operation to list all the partner event source
// names that they have created.
//
// This operation is not used by AWS customers.
//
//    // Example sending a request using ListPartnerEventSourcesRequest.
//    req := client.ListPartnerEventSourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/ListPartnerEventSources
func (c *Client) ListPartnerEventSourcesRequest(input *types.ListPartnerEventSourcesInput) ListPartnerEventSourcesRequest {
	op := &aws.Operation{
		Name:       opListPartnerEventSources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListPartnerEventSourcesInput{}
	}

	req := c.newRequest(op, input, &types.ListPartnerEventSourcesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListPartnerEventSourcesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPartnerEventSourcesRequest{Request: req, Input: input, Copy: c.ListPartnerEventSourcesRequest}
}

// ListPartnerEventSourcesRequest is the request type for the
// ListPartnerEventSources API operation.
type ListPartnerEventSourcesRequest struct {
	*aws.Request
	Input *types.ListPartnerEventSourcesInput
	Copy  func(*types.ListPartnerEventSourcesInput) ListPartnerEventSourcesRequest
}

// Send marshals and sends the ListPartnerEventSources API request.
func (r ListPartnerEventSourcesRequest) Send(ctx context.Context) (*ListPartnerEventSourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPartnerEventSourcesResponse{
		ListPartnerEventSourcesOutput: r.Request.Data.(*types.ListPartnerEventSourcesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPartnerEventSourcesResponse is the response type for the
// ListPartnerEventSources API operation.
type ListPartnerEventSourcesResponse struct {
	*types.ListPartnerEventSourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPartnerEventSources request.
func (r *ListPartnerEventSourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
