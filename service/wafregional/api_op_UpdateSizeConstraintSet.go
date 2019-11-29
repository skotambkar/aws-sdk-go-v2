// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
)

const opUpdateSizeConstraintSet = "UpdateSizeConstraintSet"

// UpdateSizeConstraintSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Inserts or deletes SizeConstraint objects (filters) in a SizeConstraintSet.
// For each SizeConstraint object, you specify the following values:
//
//    * Whether to insert or delete the object from the array. If you want to
//    change a SizeConstraintSetUpdate object, you delete the existing object
//    and add a new one.
//
//    * The part of a web request that you want AWS WAF to evaluate, such as
//    the length of a query string or the length of the User-Agent header.
//
//    * Whether to perform any transformations on the request, such as converting
//    it to lowercase, before checking its length. Note that transformations
//    of the request body are not supported because the AWS resource forwards
//    only the first 8192 bytes of your request to AWS WAF. You can only specify
//    a single type of TextTransformation.
//
//    * A ComparisonOperator used for evaluating the selected part of the request
//    against the specified Size, such as equals, greater than, less than, and
//    so on.
//
//    * The length, in bytes, that you want AWS WAF to watch for in selected
//    part of the request. The length is computed after applying the transformation.
//
// For example, you can add a SizeConstraintSetUpdate object that matches web
// requests in which the length of the User-Agent header is greater than 100
// bytes. You can then configure AWS WAF to block those requests.
//
// To create and configure a SizeConstraintSet, perform the following steps:
//
// Create a SizeConstraintSet. For more information, see CreateSizeConstraintSet.
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
//    // Example sending a request using UpdateSizeConstraintSetRequest.
//    req := client.UpdateSizeConstraintSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/UpdateSizeConstraintSet
func (c *Client) UpdateSizeConstraintSetRequest(input *types.UpdateSizeConstraintSetInput) UpdateSizeConstraintSetRequest {
	op := &aws.Operation{
		Name:       opUpdateSizeConstraintSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateSizeConstraintSetInput{}
	}

	req := c.newRequest(op, input, &types.UpdateSizeConstraintSetOutput{})
	return UpdateSizeConstraintSetRequest{Request: req, Input: input, Copy: c.UpdateSizeConstraintSetRequest}
}

// UpdateSizeConstraintSetRequest is the request type for the
// UpdateSizeConstraintSet API operation.
type UpdateSizeConstraintSetRequest struct {
	*aws.Request
	Input *types.UpdateSizeConstraintSetInput
	Copy  func(*types.UpdateSizeConstraintSetInput) UpdateSizeConstraintSetRequest
}

// Send marshals and sends the UpdateSizeConstraintSet API request.
func (r UpdateSizeConstraintSetRequest) Send(ctx context.Context) (*UpdateSizeConstraintSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSizeConstraintSetResponse{
		UpdateSizeConstraintSetOutput: r.Request.Data.(*types.UpdateSizeConstraintSetOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSizeConstraintSetResponse is the response type for the
// UpdateSizeConstraintSet API operation.
type UpdateSizeConstraintSetResponse struct {
	*types.UpdateSizeConstraintSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSizeConstraintSet request.
func (r *UpdateSizeConstraintSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
