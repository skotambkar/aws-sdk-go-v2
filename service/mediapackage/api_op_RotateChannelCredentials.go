// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
)

const opRotateChannelCredentials = "RotateChannelCredentials"

// RotateChannelCredentialsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Changes the Channel's first IngestEndpoint's username and password. WARNING
// - This API is deprecated. Please use RotateIngestEndpointCredentials instead
//
//    // Example sending a request using RotateChannelCredentialsRequest.
//    req := client.RotateChannelCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/RotateChannelCredentials
func (c *Client) RotateChannelCredentialsRequest(input *types.RotateChannelCredentialsInput) RotateChannelCredentialsRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, RotateChannelCredentials, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opRotateChannelCredentials,
		HTTPMethod: "PUT",
		HTTPPath:   "/channels/{id}/credentials",
	}

	if input == nil {
		input = &types.RotateChannelCredentialsInput{}
	}

	req := c.newRequest(op, input, &types.RotateChannelCredentialsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.RotateChannelCredentialsMarshaler{Input: input}.GetNamedBuildHandler())

	return RotateChannelCredentialsRequest{Request: req, Input: input, Copy: c.RotateChannelCredentialsRequest}
}

// RotateChannelCredentialsRequest is the request type for the
// RotateChannelCredentials API operation.
type RotateChannelCredentialsRequest struct {
	*aws.Request
	Input *types.RotateChannelCredentialsInput
	Copy  func(*types.RotateChannelCredentialsInput) RotateChannelCredentialsRequest
}

// Send marshals and sends the RotateChannelCredentials API request.
func (r RotateChannelCredentialsRequest) Send(ctx context.Context) (*RotateChannelCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RotateChannelCredentialsResponse{
		RotateChannelCredentialsOutput: r.Request.Data.(*types.RotateChannelCredentialsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RotateChannelCredentialsResponse is the response type for the
// RotateChannelCredentials API operation.
type RotateChannelCredentialsResponse struct {
	*types.RotateChannelCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RotateChannelCredentials request.
func (r *RotateChannelCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
