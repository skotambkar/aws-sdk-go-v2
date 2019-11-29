// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecr/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

const opPutImageScanningConfiguration = "PutImageScanningConfiguration"

// PutImageScanningConfigurationRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Updates the image scanning configuration for a repository.
//
//    // Example sending a request using PutImageScanningConfigurationRequest.
//    req := client.PutImageScanningConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutImageScanningConfiguration
func (c *Client) PutImageScanningConfigurationRequest(input *types.PutImageScanningConfigurationInput) PutImageScanningConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutImageScanningConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutImageScanningConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.PutImageScanningConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.PutImageScanningConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return PutImageScanningConfigurationRequest{Request: req, Input: input, Copy: c.PutImageScanningConfigurationRequest}
}

// PutImageScanningConfigurationRequest is the request type for the
// PutImageScanningConfiguration API operation.
type PutImageScanningConfigurationRequest struct {
	*aws.Request
	Input *types.PutImageScanningConfigurationInput
	Copy  func(*types.PutImageScanningConfigurationInput) PutImageScanningConfigurationRequest
}

// Send marshals and sends the PutImageScanningConfiguration API request.
func (r PutImageScanningConfigurationRequest) Send(ctx context.Context) (*PutImageScanningConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutImageScanningConfigurationResponse{
		PutImageScanningConfigurationOutput: r.Request.Data.(*types.PutImageScanningConfigurationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutImageScanningConfigurationResponse is the response type for the
// PutImageScanningConfiguration API operation.
type PutImageScanningConfigurationResponse struct {
	*types.PutImageScanningConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutImageScanningConfiguration request.
func (r *PutImageScanningConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
