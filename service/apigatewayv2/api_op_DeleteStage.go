// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opDeleteStage = "DeleteStage"

// DeleteStageRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes a Stage.
//
//    // Example sending a request using DeleteStageRequest.
//    req := client.DeleteStageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteStage
func (c *Client) DeleteStageRequest(input *types.DeleteStageInput) DeleteStageRequest {
	op := &aws.Operation{
		Name:       opDeleteStage,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/apis/{apiId}/stages/{stageName}",
	}

	if input == nil {
		input = &types.DeleteStageInput{}
	}

	req := c.newRequest(op, input, &types.DeleteStageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteStageMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteStageRequest{Request: req, Input: input, Copy: c.DeleteStageRequest}
}

// DeleteStageRequest is the request type for the
// DeleteStage API operation.
type DeleteStageRequest struct {
	*aws.Request
	Input *types.DeleteStageInput
	Copy  func(*types.DeleteStageInput) DeleteStageRequest
}

// Send marshals and sends the DeleteStage API request.
func (r DeleteStageRequest) Send(ctx context.Context) (*DeleteStageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteStageResponse{
		DeleteStageOutput: r.Request.Data.(*types.DeleteStageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteStageResponse is the response type for the
// DeleteStage API operation.
type DeleteStageResponse struct {
	*types.DeleteStageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteStage request.
func (r *DeleteStageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
