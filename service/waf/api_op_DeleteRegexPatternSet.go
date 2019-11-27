// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opDeleteRegexPatternSet = "DeleteRegexPatternSet"

// DeleteRegexPatternSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Permanently deletes a RegexPatternSet. You can't delete a RegexPatternSet
// if it's still used in any RegexMatchSet or if the RegexPatternSet is not
// empty.
//
//    // Example sending a request using DeleteRegexPatternSetRequest.
//    req := client.DeleteRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteRegexPatternSet
func (c *Client) DeleteRegexPatternSetRequest(input *types.DeleteRegexPatternSetInput) DeleteRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opDeleteRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRegexPatternSetOutput{})
	return DeleteRegexPatternSetRequest{Request: req, Input: input, Copy: c.DeleteRegexPatternSetRequest}
}

// DeleteRegexPatternSetRequest is the request type for the
// DeleteRegexPatternSet API operation.
type DeleteRegexPatternSetRequest struct {
	*aws.Request
	Input *types.DeleteRegexPatternSetInput
	Copy  func(*types.DeleteRegexPatternSetInput) DeleteRegexPatternSetRequest
}

// Send marshals and sends the DeleteRegexPatternSet API request.
func (r DeleteRegexPatternSetRequest) Send(ctx context.Context) (*DeleteRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRegexPatternSetResponse{
		DeleteRegexPatternSetOutput: r.Request.Data.(*types.DeleteRegexPatternSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRegexPatternSetResponse is the response type for the
// DeleteRegexPatternSet API operation.
type DeleteRegexPatternSetResponse struct {
	*types.DeleteRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRegexPatternSet request.
func (r *DeleteRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
