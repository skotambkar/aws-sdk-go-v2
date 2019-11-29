// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opAssociateBudgetWithResource = "AssociateBudgetWithResource"

// AssociateBudgetWithResourceRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Associates the specified budget with the specified resource.
//
//    // Example sending a request using AssociateBudgetWithResourceRequest.
//    req := client.AssociateBudgetWithResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/AssociateBudgetWithResource
func (c *Client) AssociateBudgetWithResourceRequest(input *types.AssociateBudgetWithResourceInput) AssociateBudgetWithResourceRequest {
	op := &aws.Operation{
		Name:       opAssociateBudgetWithResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateBudgetWithResourceInput{}
	}

	req := c.newRequest(op, input, &types.AssociateBudgetWithResourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AssociateBudgetWithResourceMarshaler{Input: input}.GetNamedBuildHandler())

	return AssociateBudgetWithResourceRequest{Request: req, Input: input, Copy: c.AssociateBudgetWithResourceRequest}
}

// AssociateBudgetWithResourceRequest is the request type for the
// AssociateBudgetWithResource API operation.
type AssociateBudgetWithResourceRequest struct {
	*aws.Request
	Input *types.AssociateBudgetWithResourceInput
	Copy  func(*types.AssociateBudgetWithResourceInput) AssociateBudgetWithResourceRequest
}

// Send marshals and sends the AssociateBudgetWithResource API request.
func (r AssociateBudgetWithResourceRequest) Send(ctx context.Context) (*AssociateBudgetWithResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateBudgetWithResourceResponse{
		AssociateBudgetWithResourceOutput: r.Request.Data.(*types.AssociateBudgetWithResourceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateBudgetWithResourceResponse is the response type for the
// AssociateBudgetWithResource API operation.
type AssociateBudgetWithResourceResponse struct {
	*types.AssociateBudgetWithResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateBudgetWithResource request.
func (r *AssociateBudgetWithResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
