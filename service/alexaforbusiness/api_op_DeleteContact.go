// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opDeleteContact = "DeleteContact"

// DeleteContactRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes a contact by the contact ARN.
//
//    // Example sending a request using DeleteContactRequest.
//    req := client.DeleteContactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteContact
func (c *Client) DeleteContactRequest(input *types.DeleteContactInput) DeleteContactRequest {
	op := &aws.Operation{
		Name:       opDeleteContact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteContactInput{}
	}

	req := c.newRequest(op, input, &types.DeleteContactOutput{})
	return DeleteContactRequest{Request: req, Input: input, Copy: c.DeleteContactRequest}
}

// DeleteContactRequest is the request type for the
// DeleteContact API operation.
type DeleteContactRequest struct {
	*aws.Request
	Input *types.DeleteContactInput
	Copy  func(*types.DeleteContactInput) DeleteContactRequest
}

// Send marshals and sends the DeleteContact API request.
func (r DeleteContactRequest) Send(ctx context.Context) (*DeleteContactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteContactResponse{
		DeleteContactOutput: r.Request.Data.(*types.DeleteContactOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteContactResponse is the response type for the
// DeleteContact API operation.
type DeleteContactResponse struct {
	*types.DeleteContactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteContact request.
func (r *DeleteContactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
