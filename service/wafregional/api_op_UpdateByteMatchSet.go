// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opUpdateByteMatchSet = "UpdateByteMatchSet"

// UpdateByteMatchSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Inserts or deletes ByteMatchTuple objects (filters) in a ByteMatchSet. For
// each ByteMatchTuple object, you specify the following values:
//
//    * Whether to insert or delete the object from the array. If you want to
//    change a ByteMatchSetUpdate object, you delete the existing object and
//    add a new one.
//
//    * The part of a web request that you want AWS WAF to inspect, such as
//    a query string or the value of the User-Agent header.
//
//    * The bytes (typically a string that corresponds with ASCII characters)
//    that you want AWS WAF to look for. For more information, including how
//    you specify the values for the AWS WAF API and the AWS CLI or SDKs, see
//    TargetString in the ByteMatchTuple data type.
//
//    * Where to look, such as at the beginning or the end of a query string.
//
//    * Whether to perform any conversions on the request, such as converting
//    it to lowercase, before inspecting it for the specified string.
//
// For example, you can add a ByteMatchSetUpdate object that matches web requests
// in which User-Agent headers contain the string BadBot. You can then configure
// AWS WAF to block those requests.
//
// To create and configure a ByteMatchSet, perform the following steps:
//
// Create a ByteMatchSet. For more information, see CreateByteMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateByteMatchSet request.
//
// Submit an UpdateByteMatchSet request to specify the part of the request that
// you want AWS WAF to inspect (for example, the header or the URI) and the
// value that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateByteMatchSetRequest.
//    req := client.UpdateByteMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/UpdateByteMatchSet
func (c *Client) UpdateByteMatchSetRequest(input *types.UpdateByteMatchSetInput) UpdateByteMatchSetRequest {
	op := &aws.Operation{
		Name:       opUpdateByteMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateByteMatchSetInput{}
	}

	req := c.newRequest(op, input, &types.UpdateByteMatchSetOutput{})
	return UpdateByteMatchSetRequest{Request: req, Input: input, Copy: c.UpdateByteMatchSetRequest}
}

// UpdateByteMatchSetRequest is the request type for the
// UpdateByteMatchSet API operation.
type UpdateByteMatchSetRequest struct {
	*aws.Request
	Input *types.UpdateByteMatchSetInput
	Copy  func(*types.UpdateByteMatchSetInput) UpdateByteMatchSetRequest
}

// Send marshals and sends the UpdateByteMatchSet API request.
func (r UpdateByteMatchSetRequest) Send(ctx context.Context) (*UpdateByteMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateByteMatchSetResponse{
		UpdateByteMatchSetOutput: r.Request.Data.(*types.UpdateByteMatchSetOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateByteMatchSetResponse is the response type for the
// UpdateByteMatchSet API operation.
type UpdateByteMatchSetResponse struct {
	*types.UpdateByteMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateByteMatchSet request.
func (r *UpdateByteMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
