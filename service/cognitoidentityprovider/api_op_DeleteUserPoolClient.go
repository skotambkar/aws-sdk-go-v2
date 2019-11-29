// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opDeleteUserPoolClient = "DeleteUserPoolClient"

// DeleteUserPoolClientRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Allows the developer to delete the user pool client.
//
//    // Example sending a request using DeleteUserPoolClientRequest.
//    req := client.DeleteUserPoolClientRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteUserPoolClient
func (c *Client) DeleteUserPoolClientRequest(input *types.DeleteUserPoolClientInput) DeleteUserPoolClientRequest {
	op := &aws.Operation{
		Name:       opDeleteUserPoolClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteUserPoolClientInput{}
	}

	req := c.newRequest(op, input, &types.DeleteUserPoolClientOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteUserPoolClientRequest{Request: req, Input: input, Copy: c.DeleteUserPoolClientRequest}
}

// DeleteUserPoolClientRequest is the request type for the
// DeleteUserPoolClient API operation.
type DeleteUserPoolClientRequest struct {
	*aws.Request
	Input *types.DeleteUserPoolClientInput
	Copy  func(*types.DeleteUserPoolClientInput) DeleteUserPoolClientRequest
}

// Send marshals and sends the DeleteUserPoolClient API request.
func (r DeleteUserPoolClientRequest) Send(ctx context.Context) (*DeleteUserPoolClientResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserPoolClientResponse{
		DeleteUserPoolClientOutput: r.Request.Data.(*types.DeleteUserPoolClientOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserPoolClientResponse is the response type for the
// DeleteUserPoolClient API operation.
type DeleteUserPoolClientResponse struct {
	*types.DeleteUserPoolClientOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserPoolClient request.
func (r *DeleteUserPoolClientResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
