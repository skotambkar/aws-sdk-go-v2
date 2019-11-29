// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opListComplianceItems = "ListComplianceItems"

// ListComplianceItemsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// For a specified resource ID, this API action returns a list of compliance
// statuses for different resource types. Currently, you can only specify one
// resource ID per call. List results depend on the criteria specified in the
// filter.
//
//    // Example sending a request using ListComplianceItemsRequest.
//    req := client.ListComplianceItemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListComplianceItems
func (c *Client) ListComplianceItemsRequest(input *types.ListComplianceItemsInput) ListComplianceItemsRequest {
	op := &aws.Operation{
		Name:       opListComplianceItems,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListComplianceItemsInput{}
	}

	req := c.newRequest(op, input, &types.ListComplianceItemsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListComplianceItemsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListComplianceItemsRequest{Request: req, Input: input, Copy: c.ListComplianceItemsRequest}
}

// ListComplianceItemsRequest is the request type for the
// ListComplianceItems API operation.
type ListComplianceItemsRequest struct {
	*aws.Request
	Input *types.ListComplianceItemsInput
	Copy  func(*types.ListComplianceItemsInput) ListComplianceItemsRequest
}

// Send marshals and sends the ListComplianceItems API request.
func (r ListComplianceItemsRequest) Send(ctx context.Context) (*ListComplianceItemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListComplianceItemsResponse{
		ListComplianceItemsOutput: r.Request.Data.(*types.ListComplianceItemsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListComplianceItemsResponse is the response type for the
// ListComplianceItems API operation.
type ListComplianceItemsResponse struct {
	*types.ListComplianceItemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListComplianceItems request.
func (r *ListComplianceItemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
