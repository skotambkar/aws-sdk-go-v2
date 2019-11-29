// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opDeleteProfile = "DeleteProfile"

// DeleteProfileRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes a room profile by the profile ARN.
//
//    // Example sending a request using DeleteProfileRequest.
//    req := client.DeleteProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteProfile
func (c *Client) DeleteProfileRequest(input *types.DeleteProfileInput) DeleteProfileRequest {
	op := &aws.Operation{
		Name:       opDeleteProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteProfileInput{}
	}

	req := c.newRequest(op, input, &types.DeleteProfileOutput{})
	return DeleteProfileRequest{Request: req, Input: input, Copy: c.DeleteProfileRequest}
}

// DeleteProfileRequest is the request type for the
// DeleteProfile API operation.
type DeleteProfileRequest struct {
	*aws.Request
	Input *types.DeleteProfileInput
	Copy  func(*types.DeleteProfileInput) DeleteProfileRequest
}

// Send marshals and sends the DeleteProfile API request.
func (r DeleteProfileRequest) Send(ctx context.Context) (*DeleteProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProfileResponse{
		DeleteProfileOutput: r.Request.Data.(*types.DeleteProfileOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProfileResponse is the response type for the
// DeleteProfile API operation.
type DeleteProfileResponse struct {
	*types.DeleteProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProfile request.
func (r *DeleteProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
