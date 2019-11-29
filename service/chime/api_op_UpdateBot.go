// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opUpdateBot = "UpdateBot"

// UpdateBotRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates the status of the specified bot, such as starting or stopping the
// bot from running in your Amazon Chime Enterprise account.
//
//    // Example sending a request using UpdateBotRequest.
//    req := client.UpdateBotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateBot
func (c *Client) UpdateBotRequest(input *types.UpdateBotInput) UpdateBotRequest {
	op := &aws.Operation{
		Name:       opUpdateBot,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/bots/{botId}",
	}

	if input == nil {
		input = &types.UpdateBotInput{}
	}

	req := c.newRequest(op, input, &types.UpdateBotOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateBotMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateBotRequest{Request: req, Input: input, Copy: c.UpdateBotRequest}
}

// UpdateBotRequest is the request type for the
// UpdateBot API operation.
type UpdateBotRequest struct {
	*aws.Request
	Input *types.UpdateBotInput
	Copy  func(*types.UpdateBotInput) UpdateBotRequest
}

// Send marshals and sends the UpdateBot API request.
func (r UpdateBotRequest) Send(ctx context.Context) (*UpdateBotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBotResponse{
		UpdateBotOutput: r.Request.Data.(*types.UpdateBotOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBotResponse is the response type for the
// UpdateBot API operation.
type UpdateBotResponse struct {
	*types.UpdateBotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBot request.
func (r *UpdateBotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
