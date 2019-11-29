// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
)

const opDeleteChannel = "DeleteChannel"

// DeleteChannelRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Deletes an existing Channel.
//
//    // Example sending a request using DeleteChannelRequest.
//    req := client.DeleteChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DeleteChannel
func (c *Client) DeleteChannelRequest(input *types.DeleteChannelInput) DeleteChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/channels/{id}",
	}

	if input == nil {
		input = &types.DeleteChannelInput{}
	}

	req := c.newRequest(op, input, &types.DeleteChannelOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteChannelMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteChannelRequest{Request: req, Input: input, Copy: c.DeleteChannelRequest}
}

// DeleteChannelRequest is the request type for the
// DeleteChannel API operation.
type DeleteChannelRequest struct {
	*aws.Request
	Input *types.DeleteChannelInput
	Copy  func(*types.DeleteChannelInput) DeleteChannelRequest
}

// Send marshals and sends the DeleteChannel API request.
func (r DeleteChannelRequest) Send(ctx context.Context) (*DeleteChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteChannelResponse{
		DeleteChannelOutput: r.Request.Data.(*types.DeleteChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteChannelResponse is the response type for the
// DeleteChannel API operation.
type DeleteChannelResponse struct {
	*types.DeleteChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteChannel request.
func (r *DeleteChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
