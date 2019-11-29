// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opCreateRegexMatchSet = "CreateRegexMatchSet"

// CreateRegexMatchSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Creates a RegexMatchSet. You then use UpdateRegexMatchSet to identify the
// part of a web request that you want AWS WAF to inspect, such as the values
// of the User-Agent header or the query string. For example, you can create
// a RegexMatchSet that contains a RegexMatchTuple that looks for any requests
// with User-Agent headers that match a RegexPatternSet with pattern B[a@]dB[o0]t.
// You can then configure AWS WAF to reject those requests.
//
// To create and configure a RegexMatchSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateRegexMatchSet request.
//
// Submit a CreateRegexMatchSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRegexMatchSet request.
//
// Submit an UpdateRegexMatchSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and
// the value, using a RegexPatternSet, that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateRegexMatchSetRequest.
//    req := client.CreateRegexMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateRegexMatchSet
func (c *Client) CreateRegexMatchSetRequest(input *types.CreateRegexMatchSetInput) CreateRegexMatchSetRequest {
	op := &aws.Operation{
		Name:       opCreateRegexMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateRegexMatchSetInput{}
	}

	req := c.newRequest(op, input, &types.CreateRegexMatchSetOutput{})
	return CreateRegexMatchSetRequest{Request: req, Input: input, Copy: c.CreateRegexMatchSetRequest}
}

// CreateRegexMatchSetRequest is the request type for the
// CreateRegexMatchSet API operation.
type CreateRegexMatchSetRequest struct {
	*aws.Request
	Input *types.CreateRegexMatchSetInput
	Copy  func(*types.CreateRegexMatchSetInput) CreateRegexMatchSetRequest
}

// Send marshals and sends the CreateRegexMatchSet API request.
func (r CreateRegexMatchSetRequest) Send(ctx context.Context) (*CreateRegexMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRegexMatchSetResponse{
		CreateRegexMatchSetOutput: r.Request.Data.(*types.CreateRegexMatchSetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRegexMatchSetResponse is the response type for the
// CreateRegexMatchSet API operation.
type CreateRegexMatchSetResponse struct {
	*types.CreateRegexMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRegexMatchSet request.
func (r *CreateRegexMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
