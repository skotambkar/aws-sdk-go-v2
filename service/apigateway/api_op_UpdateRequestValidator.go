// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opUpdateRequestValidator = "UpdateRequestValidator"

// UpdateRequestValidatorRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Updates a RequestValidator of a given RestApi.
//
//    // Example sending a request using UpdateRequestValidatorRequest.
//    req := client.UpdateRequestValidatorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateRequestValidatorRequest(input *types.UpdateRequestValidatorInput) UpdateRequestValidatorRequest {
	op := &aws.Operation{
		Name:       opUpdateRequestValidator,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/requestvalidators/{requestvalidator_id}",
	}

	if input == nil {
		input = &types.UpdateRequestValidatorInput{}
	}

	req := c.newRequest(op, input, &types.UpdateRequestValidatorOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateRequestValidatorMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateRequestValidatorRequest{Request: req, Input: input, Copy: c.UpdateRequestValidatorRequest}
}

// UpdateRequestValidatorRequest is the request type for the
// UpdateRequestValidator API operation.
type UpdateRequestValidatorRequest struct {
	*aws.Request
	Input *types.UpdateRequestValidatorInput
	Copy  func(*types.UpdateRequestValidatorInput) UpdateRequestValidatorRequest
}

// Send marshals and sends the UpdateRequestValidator API request.
func (r UpdateRequestValidatorRequest) Send(ctx context.Context) (*UpdateRequestValidatorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRequestValidatorResponse{
		UpdateRequestValidatorOutput: r.Request.Data.(*types.UpdateRequestValidatorOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRequestValidatorResponse is the response type for the
// UpdateRequestValidator API operation.
type UpdateRequestValidatorResponse struct {
	*types.UpdateRequestValidatorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRequestValidator request.
func (r *UpdateRequestValidatorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
