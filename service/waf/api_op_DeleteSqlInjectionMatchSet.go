// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opDeleteSqlInjectionMatchSet = "DeleteSqlInjectionMatchSet"

// DeleteSqlInjectionMatchSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Permanently deletes a SqlInjectionMatchSet. You can't delete a SqlInjectionMatchSet
// if it's still used in any Rules or if it still contains any SqlInjectionMatchTuple
// objects.
//
// If you just want to remove a SqlInjectionMatchSet from a Rule, use UpdateRule.
//
// To permanently delete a SqlInjectionMatchSet from AWS WAF, perform the following
// steps:
//
// Update the SqlInjectionMatchSet to remove filters, if any. For more information,
// see UpdateSqlInjectionMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteSqlInjectionMatchSet request.
//
// Submit a DeleteSqlInjectionMatchSet request.
//
//    // Example sending a request using DeleteSqlInjectionMatchSetRequest.
//    req := client.DeleteSqlInjectionMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteSqlInjectionMatchSet
func (c *Client) DeleteSqlInjectionMatchSetRequest(input *types.DeleteSqlInjectionMatchSetInput) DeleteSqlInjectionMatchSetRequest {
	op := &aws.Operation{
		Name:       opDeleteSqlInjectionMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteSqlInjectionMatchSetInput{}
	}

	req := c.newRequest(op, input, &types.DeleteSqlInjectionMatchSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteSqlInjectionMatchSetMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteSqlInjectionMatchSetRequest{Request: req, Input: input, Copy: c.DeleteSqlInjectionMatchSetRequest}
}

// DeleteSqlInjectionMatchSetRequest is the request type for the
// DeleteSqlInjectionMatchSet API operation.
type DeleteSqlInjectionMatchSetRequest struct {
	*aws.Request
	Input *types.DeleteSqlInjectionMatchSetInput
	Copy  func(*types.DeleteSqlInjectionMatchSetInput) DeleteSqlInjectionMatchSetRequest
}

// Send marshals and sends the DeleteSqlInjectionMatchSet API request.
func (r DeleteSqlInjectionMatchSetRequest) Send(ctx context.Context) (*DeleteSqlInjectionMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSqlInjectionMatchSetResponse{
		DeleteSqlInjectionMatchSetOutput: r.Request.Data.(*types.DeleteSqlInjectionMatchSetOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSqlInjectionMatchSetResponse is the response type for the
// DeleteSqlInjectionMatchSet API operation.
type DeleteSqlInjectionMatchSetResponse struct {
	*types.DeleteSqlInjectionMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSqlInjectionMatchSet request.
func (r *DeleteSqlInjectionMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
