// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opDeleteRequestValidator = "DeleteRequestValidator"

// DeleteRequestValidatorRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes a RequestValidator of a given RestApi.
//
//    // Example sending a request using DeleteRequestValidatorRequest.
//    req := client.DeleteRequestValidatorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteRequestValidatorRequest(input *types.DeleteRequestValidatorInput) DeleteRequestValidatorRequest {
	op := &aws.Operation{
		Name:       opDeleteRequestValidator,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}/requestvalidators/{requestvalidator_id}",
	}

	if input == nil {
		input = &types.DeleteRequestValidatorInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRequestValidatorOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRequestValidatorRequest{Request: req, Input: input, Copy: c.DeleteRequestValidatorRequest}
}

// DeleteRequestValidatorRequest is the request type for the
// DeleteRequestValidator API operation.
type DeleteRequestValidatorRequest struct {
	*aws.Request
	Input *types.DeleteRequestValidatorInput
	Copy  func(*types.DeleteRequestValidatorInput) DeleteRequestValidatorRequest
}

// Send marshals and sends the DeleteRequestValidator API request.
func (r DeleteRequestValidatorRequest) Send(ctx context.Context) (*DeleteRequestValidatorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRequestValidatorResponse{
		DeleteRequestValidatorOutput: r.Request.Data.(*types.DeleteRequestValidatorOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRequestValidatorResponse is the response type for the
// DeleteRequestValidator API operation.
type DeleteRequestValidatorResponse struct {
	*types.DeleteRequestValidatorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRequestValidator request.
func (r *DeleteRequestValidatorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
