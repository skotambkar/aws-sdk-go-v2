// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opPutEmailIdentityDkimAttributes = "PutEmailIdentityDkimAttributes"

// PutEmailIdentityDkimAttributesRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Used to enable or disable DKIM authentication for an email identity.
//
//    // Example sending a request using PutEmailIdentityDkimAttributesRequest.
//    req := client.PutEmailIdentityDkimAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutEmailIdentityDkimAttributes
func (c *Client) PutEmailIdentityDkimAttributesRequest(input *types.PutEmailIdentityDkimAttributesInput) PutEmailIdentityDkimAttributesRequest {
	op := &aws.Operation{
		Name:       opPutEmailIdentityDkimAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/identities/{EmailIdentity}/dkim",
	}

	if input == nil {
		input = &types.PutEmailIdentityDkimAttributesInput{}
	}

	req := c.newRequest(op, input, &types.PutEmailIdentityDkimAttributesOutput{})
	return PutEmailIdentityDkimAttributesRequest{Request: req, Input: input, Copy: c.PutEmailIdentityDkimAttributesRequest}
}

// PutEmailIdentityDkimAttributesRequest is the request type for the
// PutEmailIdentityDkimAttributes API operation.
type PutEmailIdentityDkimAttributesRequest struct {
	*aws.Request
	Input *types.PutEmailIdentityDkimAttributesInput
	Copy  func(*types.PutEmailIdentityDkimAttributesInput) PutEmailIdentityDkimAttributesRequest
}

// Send marshals and sends the PutEmailIdentityDkimAttributes API request.
func (r PutEmailIdentityDkimAttributesRequest) Send(ctx context.Context) (*PutEmailIdentityDkimAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEmailIdentityDkimAttributesResponse{
		PutEmailIdentityDkimAttributesOutput: r.Request.Data.(*types.PutEmailIdentityDkimAttributesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEmailIdentityDkimAttributesResponse is the response type for the
// PutEmailIdentityDkimAttributes API operation.
type PutEmailIdentityDkimAttributesResponse struct {
	*types.PutEmailIdentityDkimAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEmailIdentityDkimAttributes request.
func (r *PutEmailIdentityDkimAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
