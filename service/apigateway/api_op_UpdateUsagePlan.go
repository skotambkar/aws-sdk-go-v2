// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opUpdateUsagePlan = "UpdateUsagePlan"

// UpdateUsagePlanRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Updates a usage plan of a given plan Id.
//
//    // Example sending a request using UpdateUsagePlanRequest.
//    req := client.UpdateUsagePlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateUsagePlanRequest(input *types.UpdateUsagePlanInput) UpdateUsagePlanRequest {
	op := &aws.Operation{
		Name:       opUpdateUsagePlan,
		HTTPMethod: "PATCH",
		HTTPPath:   "/usageplans/{usageplanId}",
	}

	if input == nil {
		input = &types.UpdateUsagePlanInput{}
	}

	req := c.newRequest(op, input, &types.UpdateUsagePlanOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateUsagePlanMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateUsagePlanRequest{Request: req, Input: input, Copy: c.UpdateUsagePlanRequest}
}

// UpdateUsagePlanRequest is the request type for the
// UpdateUsagePlan API operation.
type UpdateUsagePlanRequest struct {
	*aws.Request
	Input *types.UpdateUsagePlanInput
	Copy  func(*types.UpdateUsagePlanInput) UpdateUsagePlanRequest
}

// Send marshals and sends the UpdateUsagePlan API request.
func (r UpdateUsagePlanRequest) Send(ctx context.Context) (*UpdateUsagePlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUsagePlanResponse{
		UpdateUsagePlanOutput: r.Request.Data.(*types.UpdateUsagePlanOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUsagePlanResponse is the response type for the
// UpdateUsagePlan API operation.
type UpdateUsagePlanResponse struct {
	*types.UpdateUsagePlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUsagePlan request.
func (r *UpdateUsagePlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
