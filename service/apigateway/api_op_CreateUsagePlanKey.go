// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opCreateUsagePlanKey = "CreateUsagePlanKey"

// CreateUsagePlanKeyRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Creates a usage plan key for adding an existing API key to a usage plan.
//
//    // Example sending a request using CreateUsagePlanKeyRequest.
//    req := client.CreateUsagePlanKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateUsagePlanKeyRequest(input *types.CreateUsagePlanKeyInput) CreateUsagePlanKeyRequest {
	op := &aws.Operation{
		Name:       opCreateUsagePlanKey,
		HTTPMethod: "POST",
		HTTPPath:   "/usageplans/{usageplanId}/keys",
	}

	if input == nil {
		input = &types.CreateUsagePlanKeyInput{}
	}

	req := c.newRequest(op, input, &types.CreateUsagePlanKeyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateUsagePlanKeyMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateUsagePlanKeyRequest{Request: req, Input: input, Copy: c.CreateUsagePlanKeyRequest}
}

// CreateUsagePlanKeyRequest is the request type for the
// CreateUsagePlanKey API operation.
type CreateUsagePlanKeyRequest struct {
	*aws.Request
	Input *types.CreateUsagePlanKeyInput
	Copy  func(*types.CreateUsagePlanKeyInput) CreateUsagePlanKeyRequest
}

// Send marshals and sends the CreateUsagePlanKey API request.
func (r CreateUsagePlanKeyRequest) Send(ctx context.Context) (*CreateUsagePlanKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUsagePlanKeyResponse{
		CreateUsagePlanKeyOutput: r.Request.Data.(*types.CreateUsagePlanKeyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUsagePlanKeyResponse is the response type for the
// CreateUsagePlanKey API operation.
type CreateUsagePlanKeyResponse struct {
	*types.CreateUsagePlanKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUsagePlanKey request.
func (r *CreateUsagePlanKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
