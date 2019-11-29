// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opDeleteRegexMatchSet = "DeleteRegexMatchSet"

// DeleteRegexMatchSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Permanently deletes a RegexMatchSet. You can't delete a RegexMatchSet if
// it's still used in any Rules or if it still includes any RegexMatchTuples
// objects (any filters).
//
// If you just want to remove a RegexMatchSet from a Rule, use UpdateRule.
//
// To permanently delete a RegexMatchSet, perform the following steps:
//
// Update the RegexMatchSet to remove filters, if any. For more information,
// see UpdateRegexMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteRegexMatchSet request.
//
// Submit a DeleteRegexMatchSet request.
//
//    // Example sending a request using DeleteRegexMatchSetRequest.
//    req := client.DeleteRegexMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteRegexMatchSet
func (c *Client) DeleteRegexMatchSetRequest(input *types.DeleteRegexMatchSetInput) DeleteRegexMatchSetRequest {
	op := &aws.Operation{
		Name:       opDeleteRegexMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteRegexMatchSetInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRegexMatchSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteRegexMatchSetMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteRegexMatchSetRequest{Request: req, Input: input, Copy: c.DeleteRegexMatchSetRequest}
}

// DeleteRegexMatchSetRequest is the request type for the
// DeleteRegexMatchSet API operation.
type DeleteRegexMatchSetRequest struct {
	*aws.Request
	Input *types.DeleteRegexMatchSetInput
	Copy  func(*types.DeleteRegexMatchSetInput) DeleteRegexMatchSetRequest
}

// Send marshals and sends the DeleteRegexMatchSet API request.
func (r DeleteRegexMatchSetRequest) Send(ctx context.Context) (*DeleteRegexMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRegexMatchSetResponse{
		DeleteRegexMatchSetOutput: r.Request.Data.(*types.DeleteRegexMatchSetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRegexMatchSetResponse is the response type for the
// DeleteRegexMatchSet API operation.
type DeleteRegexMatchSetResponse struct {
	*types.DeleteRegexMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRegexMatchSet request.
func (r *DeleteRegexMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
