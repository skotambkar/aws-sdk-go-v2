// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
)

const opDeletePackagingConfiguration = "DeletePackagingConfiguration"

// DeletePackagingConfigurationRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Deletes a MediaPackage VOD PackagingConfiguration resource.
//
//    // Example sending a request using DeletePackagingConfigurationRequest.
//    req := client.DeletePackagingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeletePackagingConfiguration
func (c *Client) DeletePackagingConfigurationRequest(input *types.DeletePackagingConfigurationInput) DeletePackagingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeletePackagingConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/packaging_configurations/{id}",
	}

	if input == nil {
		input = &types.DeletePackagingConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DeletePackagingConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeletePackagingConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return DeletePackagingConfigurationRequest{Request: req, Input: input, Copy: c.DeletePackagingConfigurationRequest}
}

// DeletePackagingConfigurationRequest is the request type for the
// DeletePackagingConfiguration API operation.
type DeletePackagingConfigurationRequest struct {
	*aws.Request
	Input *types.DeletePackagingConfigurationInput
	Copy  func(*types.DeletePackagingConfigurationInput) DeletePackagingConfigurationRequest
}

// Send marshals and sends the DeletePackagingConfiguration API request.
func (r DeletePackagingConfigurationRequest) Send(ctx context.Context) (*DeletePackagingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePackagingConfigurationResponse{
		DeletePackagingConfigurationOutput: r.Request.Data.(*types.DeletePackagingConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePackagingConfigurationResponse is the response type for the
// DeletePackagingConfiguration API operation.
type DeletePackagingConfigurationResponse struct {
	*types.DeletePackagingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePackagingConfiguration request.
func (r *DeletePackagingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
