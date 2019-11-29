// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opCreateSizeConstraintSet = "CreateSizeConstraintSet"

// CreateSizeConstraintSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Creates a SizeConstraintSet. You then use UpdateSizeConstraintSet to identify
// the part of a web request that you want AWS WAF to check for length, such
// as the length of the User-Agent header or the length of the query string.
// For example, you can create a SizeConstraintSet that matches any requests
// that have a query string that is longer than 100 bytes. You can then configure
// AWS WAF to reject those requests.
//
// To create and configure a SizeConstraintSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateSizeConstraintSet request.
//
// Submit a CreateSizeConstraintSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateSizeConstraintSet request.
//
// Submit an UpdateSizeConstraintSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and
// the value that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateSizeConstraintSetRequest.
//    req := client.CreateSizeConstraintSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/CreateSizeConstraintSet
func (c *Client) CreateSizeConstraintSetRequest(input *types.CreateSizeConstraintSetInput) CreateSizeConstraintSetRequest {
	op := &aws.Operation{
		Name:       opCreateSizeConstraintSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateSizeConstraintSetInput{}
	}

	req := c.newRequest(op, input, &types.CreateSizeConstraintSetOutput{})
	return CreateSizeConstraintSetRequest{Request: req, Input: input, Copy: c.CreateSizeConstraintSetRequest}
}

// CreateSizeConstraintSetRequest is the request type for the
// CreateSizeConstraintSet API operation.
type CreateSizeConstraintSetRequest struct {
	*aws.Request
	Input *types.CreateSizeConstraintSetInput
	Copy  func(*types.CreateSizeConstraintSetInput) CreateSizeConstraintSetRequest
}

// Send marshals and sends the CreateSizeConstraintSet API request.
func (r CreateSizeConstraintSetRequest) Send(ctx context.Context) (*CreateSizeConstraintSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSizeConstraintSetResponse{
		CreateSizeConstraintSetOutput: r.Request.Data.(*types.CreateSizeConstraintSetOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSizeConstraintSetResponse is the response type for the
// CreateSizeConstraintSet API operation.
type CreateSizeConstraintSetResponse struct {
	*types.CreateSizeConstraintSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSizeConstraintSet request.
func (r *CreateSizeConstraintSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
