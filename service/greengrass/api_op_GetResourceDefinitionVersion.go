// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opGetResourceDefinitionVersion = "GetResourceDefinitionVersion"

// GetResourceDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a resource definition version, including which
// resources are included in the version.
//
//    // Example sending a request using GetResourceDefinitionVersionRequest.
//    req := client.GetResourceDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetResourceDefinitionVersion
func (c *Client) GetResourceDefinitionVersionRequest(input *types.GetResourceDefinitionVersionInput) GetResourceDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opGetResourceDefinitionVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/resources/{ResourceDefinitionId}/versions/{ResourceDefinitionVersionId}",
	}

	if input == nil {
		input = &types.GetResourceDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &types.GetResourceDefinitionVersionOutput{})
	return GetResourceDefinitionVersionRequest{Request: req, Input: input, Copy: c.GetResourceDefinitionVersionRequest}
}

// GetResourceDefinitionVersionRequest is the request type for the
// GetResourceDefinitionVersion API operation.
type GetResourceDefinitionVersionRequest struct {
	*aws.Request
	Input *types.GetResourceDefinitionVersionInput
	Copy  func(*types.GetResourceDefinitionVersionInput) GetResourceDefinitionVersionRequest
}

// Send marshals and sends the GetResourceDefinitionVersion API request.
func (r GetResourceDefinitionVersionRequest) Send(ctx context.Context) (*GetResourceDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourceDefinitionVersionResponse{
		GetResourceDefinitionVersionOutput: r.Request.Data.(*types.GetResourceDefinitionVersionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResourceDefinitionVersionResponse is the response type for the
// GetResourceDefinitionVersion API operation.
type GetResourceDefinitionVersionResponse struct {
	*types.GetResourceDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourceDefinitionVersion request.
func (r *GetResourceDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
