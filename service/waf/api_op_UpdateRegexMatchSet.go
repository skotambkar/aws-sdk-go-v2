// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opUpdateRegexMatchSet = "UpdateRegexMatchSet"

// UpdateRegexMatchSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Inserts or deletes RegexMatchTuple objects (filters) in a RegexMatchSet.
// For each RegexMatchSetUpdate object, you specify the following values:
//
//    * Whether to insert or delete the object from the array. If you want to
//    change a RegexMatchSetUpdate object, you delete the existing object and
//    add a new one.
//
//    * The part of a web request that you want AWS WAF to inspectupdate, such
//    as a query string or the value of the User-Agent header.
//
//    * The identifier of the pattern (a regular expression) that you want AWS
//    WAF to look for. For more information, see RegexPatternSet.
//
//    * Whether to perform any conversions on the request, such as converting
//    it to lowercase, before inspecting it for the specified string.
//
// For example, you can create a RegexPatternSet that matches any requests with
// User-Agent headers that contain the string B[a@]dB[o0]t. You can then configure
// AWS WAF to reject those requests.
//
// To create and configure a RegexMatchSet, perform the following steps:
//
// Create a RegexMatchSet. For more information, see CreateRegexMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRegexMatchSet request.
//
// Submit an UpdateRegexMatchSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and
// the identifier of the RegexPatternSet that contain the regular expression
// patters you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateRegexMatchSetRequest.
//    req := client.UpdateRegexMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateRegexMatchSet
func (c *Client) UpdateRegexMatchSetRequest(input *types.UpdateRegexMatchSetInput) UpdateRegexMatchSetRequest {
	op := &aws.Operation{
		Name:       opUpdateRegexMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateRegexMatchSetInput{}
	}

	req := c.newRequest(op, input, &types.UpdateRegexMatchSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateRegexMatchSetMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateRegexMatchSetRequest{Request: req, Input: input, Copy: c.UpdateRegexMatchSetRequest}
}

// UpdateRegexMatchSetRequest is the request type for the
// UpdateRegexMatchSet API operation.
type UpdateRegexMatchSetRequest struct {
	*aws.Request
	Input *types.UpdateRegexMatchSetInput
	Copy  func(*types.UpdateRegexMatchSetInput) UpdateRegexMatchSetRequest
}

// Send marshals and sends the UpdateRegexMatchSet API request.
func (r UpdateRegexMatchSetRequest) Send(ctx context.Context) (*UpdateRegexMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRegexMatchSetResponse{
		UpdateRegexMatchSetOutput: r.Request.Data.(*types.UpdateRegexMatchSetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRegexMatchSetResponse is the response type for the
// UpdateRegexMatchSet API operation.
type UpdateRegexMatchSetResponse struct {
	*types.UpdateRegexMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRegexMatchSet request.
func (r *UpdateRegexMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
