// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetSdk = "GetSdk"

// GetSdkRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Generates a client SDK for a RestApi and Stage.
//
//    // Example sending a request using GetSdkRequest.
//    req := client.GetSdkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetSdkRequest(input *types.GetSdkInput) GetSdkRequest {
	op := &aws.Operation{
		Name:       opGetSdk,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/stages/{stage_name}/sdks/{sdk_type}",
	}

	if input == nil {
		input = &types.GetSdkInput{}
	}

	req := c.newRequest(op, input, &types.GetSdkOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetSdkMarshaler{Input: input}.GetNamedBuildHandler())

	return GetSdkRequest{Request: req, Input: input, Copy: c.GetSdkRequest}
}

// GetSdkRequest is the request type for the
// GetSdk API operation.
type GetSdkRequest struct {
	*aws.Request
	Input *types.GetSdkInput
	Copy  func(*types.GetSdkInput) GetSdkRequest
}

// Send marshals and sends the GetSdk API request.
func (r GetSdkRequest) Send(ctx context.Context) (*GetSdkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSdkResponse{
		GetSdkOutput: r.Request.Data.(*types.GetSdkOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSdkResponse is the response type for the
// GetSdk API operation.
type GetSdkResponse struct {
	*types.GetSdkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSdk request.
func (r *GetSdkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
