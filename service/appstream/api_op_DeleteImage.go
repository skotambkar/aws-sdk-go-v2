// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opDeleteImage = "DeleteImage"

// DeleteImageRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Deletes the specified image. You cannot delete an image when it is in use.
// After you delete an image, you cannot provision new capacity using the image.
//
//    // Example sending a request using DeleteImageRequest.
//    req := client.DeleteImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteImage
func (c *Client) DeleteImageRequest(input *types.DeleteImageInput) DeleteImageRequest {
	op := &aws.Operation{
		Name:       opDeleteImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteImageInput{}
	}

	req := c.newRequest(op, input, &types.DeleteImageOutput{})
	return DeleteImageRequest{Request: req, Input: input, Copy: c.DeleteImageRequest}
}

// DeleteImageRequest is the request type for the
// DeleteImage API operation.
type DeleteImageRequest struct {
	*aws.Request
	Input *types.DeleteImageInput
	Copy  func(*types.DeleteImageInput) DeleteImageRequest
}

// Send marshals and sends the DeleteImage API request.
func (r DeleteImageRequest) Send(ctx context.Context) (*DeleteImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImageResponse{
		DeleteImageOutput: r.Request.Data.(*types.DeleteImageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImageResponse is the response type for the
// DeleteImage API operation.
type DeleteImageResponse struct {
	*types.DeleteImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImage request.
func (r *DeleteImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
