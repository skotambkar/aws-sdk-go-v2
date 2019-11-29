// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opListProvisionedProductPlans = "ListProvisionedProductPlans"

// ListProvisionedProductPlansRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists the plans for the specified provisioned product or all plans to which
// the user has access.
//
//    // Example sending a request using ListProvisionedProductPlansRequest.
//    req := client.ListProvisionedProductPlansRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListProvisionedProductPlans
func (c *Client) ListProvisionedProductPlansRequest(input *types.ListProvisionedProductPlansInput) ListProvisionedProductPlansRequest {
	op := &aws.Operation{
		Name:       opListProvisionedProductPlans,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListProvisionedProductPlansInput{}
	}

	req := c.newRequest(op, input, &types.ListProvisionedProductPlansOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListProvisionedProductPlansMarshaler{Input: input}.GetNamedBuildHandler())

	return ListProvisionedProductPlansRequest{Request: req, Input: input, Copy: c.ListProvisionedProductPlansRequest}
}

// ListProvisionedProductPlansRequest is the request type for the
// ListProvisionedProductPlans API operation.
type ListProvisionedProductPlansRequest struct {
	*aws.Request
	Input *types.ListProvisionedProductPlansInput
	Copy  func(*types.ListProvisionedProductPlansInput) ListProvisionedProductPlansRequest
}

// Send marshals and sends the ListProvisionedProductPlans API request.
func (r ListProvisionedProductPlansRequest) Send(ctx context.Context) (*ListProvisionedProductPlansResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProvisionedProductPlansResponse{
		ListProvisionedProductPlansOutput: r.Request.Data.(*types.ListProvisionedProductPlansOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListProvisionedProductPlansResponse is the response type for the
// ListProvisionedProductPlans API operation.
type ListProvisionedProductPlansResponse struct {
	*types.ListProvisionedProductPlansOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProvisionedProductPlans request.
func (r *ListProvisionedProductPlansResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
