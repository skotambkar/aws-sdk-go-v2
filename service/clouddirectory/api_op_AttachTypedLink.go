// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opAttachTypedLink = "AttachTypedLink"

// AttachTypedLinkRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Attaches a typed link to a specified source and target object. For more information,
// see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using AttachTypedLinkRequest.
//    req := client.AttachTypedLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachTypedLink
func (c *Client) AttachTypedLinkRequest(input *types.AttachTypedLinkInput) AttachTypedLinkRequest {
	op := &aws.Operation{
		Name:       opAttachTypedLink,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/attach",
	}

	if input == nil {
		input = &types.AttachTypedLinkInput{}
	}

	req := c.newRequest(op, input, &types.AttachTypedLinkOutput{})
	return AttachTypedLinkRequest{Request: req, Input: input, Copy: c.AttachTypedLinkRequest}
}

// AttachTypedLinkRequest is the request type for the
// AttachTypedLink API operation.
type AttachTypedLinkRequest struct {
	*aws.Request
	Input *types.AttachTypedLinkInput
	Copy  func(*types.AttachTypedLinkInput) AttachTypedLinkRequest
}

// Send marshals and sends the AttachTypedLink API request.
func (r AttachTypedLinkRequest) Send(ctx context.Context) (*AttachTypedLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachTypedLinkResponse{
		AttachTypedLinkOutput: r.Request.Data.(*types.AttachTypedLinkOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachTypedLinkResponse is the response type for the
// AttachTypedLink API operation.
type AttachTypedLinkResponse struct {
	*types.AttachTypedLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachTypedLink request.
func (r *AttachTypedLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
