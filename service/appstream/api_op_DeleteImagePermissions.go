// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opDeleteImagePermissions = "DeleteImagePermissions"

// DeleteImagePermissionsRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Deletes permissions for the specified private image. After you delete permissions
// for an image, AWS accounts to which you previously granted these permissions
// can no longer use the image.
//
//    // Example sending a request using DeleteImagePermissionsRequest.
//    req := client.DeleteImagePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteImagePermissions
func (c *Client) DeleteImagePermissionsRequest(input *types.DeleteImagePermissionsInput) DeleteImagePermissionsRequest {
	op := &aws.Operation{
		Name:       opDeleteImagePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteImagePermissionsInput{}
	}

	req := c.newRequest(op, input, &types.DeleteImagePermissionsOutput{})
	return DeleteImagePermissionsRequest{Request: req, Input: input, Copy: c.DeleteImagePermissionsRequest}
}

// DeleteImagePermissionsRequest is the request type for the
// DeleteImagePermissions API operation.
type DeleteImagePermissionsRequest struct {
	*aws.Request
	Input *types.DeleteImagePermissionsInput
	Copy  func(*types.DeleteImagePermissionsInput) DeleteImagePermissionsRequest
}

// Send marshals and sends the DeleteImagePermissions API request.
func (r DeleteImagePermissionsRequest) Send(ctx context.Context) (*DeleteImagePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImagePermissionsResponse{
		DeleteImagePermissionsOutput: r.Request.Data.(*types.DeleteImagePermissionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImagePermissionsResponse is the response type for the
// DeleteImagePermissions API operation.
type DeleteImagePermissionsResponse struct {
	*types.DeleteImagePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImagePermissions request.
func (r *DeleteImagePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
