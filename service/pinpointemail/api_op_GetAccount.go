// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opGetAccount = "GetAccount"

// GetAccountRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Obtain information about the email-sending status and capabilities of your
// Amazon Pinpoint account in the current AWS Region.
//
//    // Example sending a request using GetAccountRequest.
//    req := client.GetAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetAccount
func (c *Client) GetAccountRequest(input *types.GetAccountInput) GetAccountRequest {
	op := &aws.Operation{
		Name:       opGetAccount,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/account",
	}

	if input == nil {
		input = &types.GetAccountInput{}
	}

	req := c.newRequest(op, input, &types.GetAccountOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetAccountMarshaler{Input: input}.GetNamedBuildHandler())

	return GetAccountRequest{Request: req, Input: input, Copy: c.GetAccountRequest}
}

// GetAccountRequest is the request type for the
// GetAccount API operation.
type GetAccountRequest struct {
	*aws.Request
	Input *types.GetAccountInput
	Copy  func(*types.GetAccountInput) GetAccountRequest
}

// Send marshals and sends the GetAccount API request.
func (r GetAccountRequest) Send(ctx context.Context) (*GetAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountResponse{
		GetAccountOutput: r.Request.Data.(*types.GetAccountOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountResponse is the response type for the
// GetAccount API operation.
type GetAccountResponse struct {
	*types.GetAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccount request.
func (r *GetAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
