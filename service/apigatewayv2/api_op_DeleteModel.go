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

const opDeleteModel = "DeleteModel"

// DeleteModelRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes a Model.
//
//    // Example sending a request using DeleteModelRequest.
//    req := client.DeleteModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteModel
func (c *Client) DeleteModelRequest(input *types.DeleteModelInput) DeleteModelRequest {
	op := &aws.Operation{
		Name:       opDeleteModel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/apis/{apiId}/models/{modelId}",
	}

	if input == nil {
		input = &types.DeleteModelInput{}
	}

	req := c.newRequest(op, input, &types.DeleteModelOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteModelMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteModelRequest{Request: req, Input: input, Copy: c.DeleteModelRequest}
}

// DeleteModelRequest is the request type for the
// DeleteModel API operation.
type DeleteModelRequest struct {
	*aws.Request
	Input *types.DeleteModelInput
	Copy  func(*types.DeleteModelInput) DeleteModelRequest
}

// Send marshals and sends the DeleteModel API request.
func (r DeleteModelRequest) Send(ctx context.Context) (*DeleteModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteModelResponse{
		DeleteModelOutput: r.Request.Data.(*types.DeleteModelOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteModelResponse is the response type for the
// DeleteModel API operation.
type DeleteModelResponse struct {
	*types.DeleteModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteModel request.
func (r *DeleteModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
