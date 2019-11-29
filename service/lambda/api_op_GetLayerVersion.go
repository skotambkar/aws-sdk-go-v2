// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opGetLayerVersion = "GetLayerVersion"

// GetLayerVersionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns information about a version of an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html),
// with a link to download the layer archive that's valid for 10 minutes.
//
//    // Example sending a request using GetLayerVersionRequest.
//    req := client.GetLayerVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetLayerVersion
func (c *Client) GetLayerVersionRequest(input *types.GetLayerVersionInput) GetLayerVersionRequest {
	op := &aws.Operation{
		Name:       opGetLayerVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/2018-10-31/layers/{LayerName}/versions/{VersionNumber}",
	}

	if input == nil {
		input = &types.GetLayerVersionInput{}
	}

	req := c.newRequest(op, input, &types.GetLayerVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetLayerVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return GetLayerVersionRequest{Request: req, Input: input, Copy: c.GetLayerVersionRequest}
}

// GetLayerVersionRequest is the request type for the
// GetLayerVersion API operation.
type GetLayerVersionRequest struct {
	*aws.Request
	Input *types.GetLayerVersionInput
	Copy  func(*types.GetLayerVersionInput) GetLayerVersionRequest
}

// Send marshals and sends the GetLayerVersion API request.
func (r GetLayerVersionRequest) Send(ctx context.Context) (*GetLayerVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLayerVersionResponse{
		GetLayerVersionOutput: r.Request.Data.(*types.GetLayerVersionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLayerVersionResponse is the response type for the
// GetLayerVersion API operation.
type GetLayerVersionResponse struct {
	*types.GetLayerVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLayerVersion request.
func (r *GetLayerVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
