// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opGetIdentityDkimAttributes = "GetIdentityDkimAttributes"

// GetIdentityDkimAttributesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns the current status of Easy DKIM signing for an entity. For domain
// name identities, this operation also returns the DKIM tokens that are required
// for Easy DKIM signing, and whether Amazon SES has successfully verified that
// these tokens have been published.
//
// This operation takes a list of identities as input and returns the following
// information for each:
//
//    * Whether Easy DKIM signing is enabled or disabled.
//
//    * A set of DKIM tokens that represent the identity. If the identity is
//    an email address, the tokens represent the domain of that address.
//
//    * Whether Amazon SES has successfully verified the DKIM tokens published
//    in the domain's DNS. This information is only returned for domain name
//    identities, not for email addresses.
//
// This operation is throttled at one request per second and can only get DKIM
// attributes for up to 100 identities at a time.
//
// For more information about creating DNS records using DKIM tokens, go to
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
//
//    // Example sending a request using GetIdentityDkimAttributesRequest.
//    req := client.GetIdentityDkimAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetIdentityDkimAttributes
func (c *Client) GetIdentityDkimAttributesRequest(input *types.GetIdentityDkimAttributesInput) GetIdentityDkimAttributesRequest {
	op := &aws.Operation{
		Name:       opGetIdentityDkimAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetIdentityDkimAttributesInput{}
	}

	req := c.newRequest(op, input, &types.GetIdentityDkimAttributesOutput{})
	return GetIdentityDkimAttributesRequest{Request: req, Input: input, Copy: c.GetIdentityDkimAttributesRequest}
}

// GetIdentityDkimAttributesRequest is the request type for the
// GetIdentityDkimAttributes API operation.
type GetIdentityDkimAttributesRequest struct {
	*aws.Request
	Input *types.GetIdentityDkimAttributesInput
	Copy  func(*types.GetIdentityDkimAttributesInput) GetIdentityDkimAttributesRequest
}

// Send marshals and sends the GetIdentityDkimAttributes API request.
func (r GetIdentityDkimAttributesRequest) Send(ctx context.Context) (*GetIdentityDkimAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIdentityDkimAttributesResponse{
		GetIdentityDkimAttributesOutput: r.Request.Data.(*types.GetIdentityDkimAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIdentityDkimAttributesResponse is the response type for the
// GetIdentityDkimAttributes API operation.
type GetIdentityDkimAttributesResponse struct {
	*types.GetIdentityDkimAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIdentityDkimAttributes request.
func (r *GetIdentityDkimAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
