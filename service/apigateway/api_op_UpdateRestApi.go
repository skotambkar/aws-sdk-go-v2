// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opUpdateRestApi = "UpdateRestApi"

// UpdateRestApiRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Changes information about the specified API.
//
//    // Example sending a request using UpdateRestApiRequest.
//    req := client.UpdateRestApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateRestApiRequest(input *types.UpdateRestApiInput) UpdateRestApiRequest {
	op := &aws.Operation{
		Name:       opUpdateRestApi,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}",
	}

	if input == nil {
		input = &types.UpdateRestApiInput{}
	}

	req := c.newRequest(op, input, &types.UpdateRestApiOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateRestApiMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateRestApiRequest{Request: req, Input: input, Copy: c.UpdateRestApiRequest}
}

// UpdateRestApiRequest is the request type for the
// UpdateRestApi API operation.
type UpdateRestApiRequest struct {
	*aws.Request
	Input *types.UpdateRestApiInput
	Copy  func(*types.UpdateRestApiInput) UpdateRestApiRequest
}

// Send marshals and sends the UpdateRestApi API request.
func (r UpdateRestApiRequest) Send(ctx context.Context) (*UpdateRestApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRestApiResponse{
		UpdateRestApiOutput: r.Request.Data.(*types.UpdateRestApiOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRestApiResponse is the response type for the
// UpdateRestApi API operation.
type UpdateRestApiResponse struct {
	*types.UpdateRestApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRestApi request.
func (r *UpdateRestApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
