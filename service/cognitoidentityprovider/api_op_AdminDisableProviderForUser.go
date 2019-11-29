// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opAdminDisableProviderForUser = "AdminDisableProviderForUser"

// AdminDisableProviderForUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Disables the user from signing in with the specified external (SAML or social)
// identity provider. If the user to disable is a Cognito User Pools native
// username + password user, they are not permitted to use their password to
// sign-in. If the user to disable is a linked external IdP user, any link between
// that user and an existing user is removed. The next time the external user
// (no longer attached to the previously linked DestinationUser) signs in, they
// must create a new user account. See .
//
// This action is enabled only for admin access and requires developer credentials.
//
// The ProviderName must match the value specified when creating an IdP for
// the pool.
//
// To disable a native username + password user, the ProviderName value must
// be Cognito and the ProviderAttributeName must be Cognito_Subject, with the
// ProviderAttributeValue being the name that is used in the user pool for the
// user.
//
// The ProviderAttributeName must always be Cognito_Subject for social identity
// providers. The ProviderAttributeValue must always be the exact subject that
// was used when the user was originally linked as a source user.
//
// For de-linking a SAML identity, there are two scenarios. If the linked identity
// has not yet been used to sign-in, the ProviderAttributeName and ProviderAttributeValue
// must be the same values that were used for the SourceUser when the identities
// were originally linked in the call. (If the linking was done with ProviderAttributeName
// set to Cognito_Subject, the same applies here). However, if the user has
// already signed in, the ProviderAttributeName must be Cognito_Subject and
// ProviderAttributeValue must be the subject of the SAML assertion.
//
//    // Example sending a request using AdminDisableProviderForUserRequest.
//    req := client.AdminDisableProviderForUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminDisableProviderForUser
func (c *Client) AdminDisableProviderForUserRequest(input *types.AdminDisableProviderForUserInput) AdminDisableProviderForUserRequest {
	op := &aws.Operation{
		Name:       opAdminDisableProviderForUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AdminDisableProviderForUserInput{}
	}

	req := c.newRequest(op, input, &types.AdminDisableProviderForUserOutput{})
	return AdminDisableProviderForUserRequest{Request: req, Input: input, Copy: c.AdminDisableProviderForUserRequest}
}

// AdminDisableProviderForUserRequest is the request type for the
// AdminDisableProviderForUser API operation.
type AdminDisableProviderForUserRequest struct {
	*aws.Request
	Input *types.AdminDisableProviderForUserInput
	Copy  func(*types.AdminDisableProviderForUserInput) AdminDisableProviderForUserRequest
}

// Send marshals and sends the AdminDisableProviderForUser API request.
func (r AdminDisableProviderForUserRequest) Send(ctx context.Context) (*AdminDisableProviderForUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminDisableProviderForUserResponse{
		AdminDisableProviderForUserOutput: r.Request.Data.(*types.AdminDisableProviderForUserOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminDisableProviderForUserResponse is the response type for the
// AdminDisableProviderForUser API operation.
type AdminDisableProviderForUserResponse struct {
	*types.AdminDisableProviderForUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminDisableProviderForUser request.
func (r *AdminDisableProviderForUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
