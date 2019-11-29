// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmailmessageflow

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow/types"
)

const opGetRawMessageContent = "GetRawMessageContent"

// GetRawMessageContentRequest returns a request value for making API operation for
// Amazon WorkMail Message Flow.
//
// Retrieves the raw content of an in-transit email message, in MIME format.
//
//    // Example sending a request using GetRawMessageContentRequest.
//    req := client.GetRawMessageContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/GetRawMessageContent
func (c *Client) GetRawMessageContentRequest(input *types.GetRawMessageContentInput) GetRawMessageContentRequest {
	op := &aws.Operation{
		Name:       opGetRawMessageContent,
		HTTPMethod: "GET",
		HTTPPath:   "/messages/{messageId}",
	}

	if input == nil {
		input = &types.GetRawMessageContentInput{}
	}

	req := c.newRequest(op, input, &types.GetRawMessageContentOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetRawMessageContentMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRawMessageContentRequest{Request: req, Input: input, Copy: c.GetRawMessageContentRequest}
}

// GetRawMessageContentRequest is the request type for the
// GetRawMessageContent API operation.
type GetRawMessageContentRequest struct {
	*aws.Request
	Input *types.GetRawMessageContentInput
	Copy  func(*types.GetRawMessageContentInput) GetRawMessageContentRequest
}

// Send marshals and sends the GetRawMessageContent API request.
func (r GetRawMessageContentRequest) Send(ctx context.Context) (*GetRawMessageContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRawMessageContentResponse{
		GetRawMessageContentOutput: r.Request.Data.(*types.GetRawMessageContentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRawMessageContentResponse is the response type for the
// GetRawMessageContent API operation.
type GetRawMessageContentResponse struct {
	*types.GetRawMessageContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRawMessageContent request.
func (r *GetRawMessageContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
