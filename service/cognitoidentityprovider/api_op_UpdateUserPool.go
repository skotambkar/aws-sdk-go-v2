// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opUpdateUserPool = "UpdateUserPool"

// UpdateUserPoolRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Updates the specified user pool with the specified attributes. If you don't
// provide a value for an attribute, it will be set to the default value. You
// can get a list of the current user pool settings with .
//
//    // Example sending a request using UpdateUserPoolRequest.
//    req := client.UpdateUserPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateUserPool
func (c *Client) UpdateUserPoolRequest(input *types.UpdateUserPoolInput) UpdateUserPoolRequest {
	op := &aws.Operation{
		Name:       opUpdateUserPool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateUserPoolInput{}
	}

	req := c.newRequest(op, input, &types.UpdateUserPoolOutput{})
	return UpdateUserPoolRequest{Request: req, Input: input, Copy: c.UpdateUserPoolRequest}
}

// UpdateUserPoolRequest is the request type for the
// UpdateUserPool API operation.
type UpdateUserPoolRequest struct {
	*aws.Request
	Input *types.UpdateUserPoolInput
	Copy  func(*types.UpdateUserPoolInput) UpdateUserPoolRequest
}

// Send marshals and sends the UpdateUserPool API request.
func (r UpdateUserPoolRequest) Send(ctx context.Context) (*UpdateUserPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserPoolResponse{
		UpdateUserPoolOutput: r.Request.Data.(*types.UpdateUserPoolOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserPoolResponse is the response type for the
// UpdateUserPool API operation.
type UpdateUserPoolResponse struct {
	*types.UpdateUserPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserPool request.
func (r *UpdateUserPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
