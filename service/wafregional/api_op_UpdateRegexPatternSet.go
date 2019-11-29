// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opUpdateRegexPatternSet = "UpdateRegexPatternSet"

// UpdateRegexPatternSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Inserts or deletes RegexPatternString objects in a RegexPatternSet. For each
// RegexPatternString object, you specify the following values:
//
//    * Whether to insert or delete the RegexPatternString.
//
//    * The regular expression pattern that you want to insert or delete. For
//    more information, see RegexPatternSet.
//
// For example, you can create a RegexPatternString such as B[a@]dB[o0]t. AWS
// WAF will match this RegexPatternString to:
//
//    * BadBot
//
//    * BadB0t
//
//    * B@dBot
//
//    * B@dB0t
//
// To create and configure a RegexPatternSet, perform the following steps:
//
// Create a RegexPatternSet. For more information, see CreateRegexPatternSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRegexPatternSet request.
//
// Submit an UpdateRegexPatternSet request to specify the regular expression
// pattern that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateRegexPatternSetRequest.
//    req := client.UpdateRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/UpdateRegexPatternSet
func (c *Client) UpdateRegexPatternSetRequest(input *types.UpdateRegexPatternSetInput) UpdateRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opUpdateRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &types.UpdateRegexPatternSetOutput{})
	return UpdateRegexPatternSetRequest{Request: req, Input: input, Copy: c.UpdateRegexPatternSetRequest}
}

// UpdateRegexPatternSetRequest is the request type for the
// UpdateRegexPatternSet API operation.
type UpdateRegexPatternSetRequest struct {
	*aws.Request
	Input *types.UpdateRegexPatternSetInput
	Copy  func(*types.UpdateRegexPatternSetInput) UpdateRegexPatternSetRequest
}

// Send marshals and sends the UpdateRegexPatternSet API request.
func (r UpdateRegexPatternSetRequest) Send(ctx context.Context) (*UpdateRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRegexPatternSetResponse{
		UpdateRegexPatternSetOutput: r.Request.Data.(*types.UpdateRegexPatternSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRegexPatternSetResponse is the response type for the
// UpdateRegexPatternSet API operation.
type UpdateRegexPatternSetResponse struct {
	*types.UpdateRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRegexPatternSet request.
func (r *UpdateRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
