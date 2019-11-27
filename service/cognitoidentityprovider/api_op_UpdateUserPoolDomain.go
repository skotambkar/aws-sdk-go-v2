// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opUpdateUserPoolDomain = "UpdateUserPoolDomain"

// UpdateUserPoolDomainRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Updates the Secure Sockets Layer (SSL) certificate for the custom domain
// for your user pool.
//
// You can use this operation to provide the Amazon Resource Name (ARN) of a
// new certificate to Amazon Cognito. You cannot use it to change the domain
// for a user pool.
//
// A custom domain is used to host the Amazon Cognito hosted UI, which provides
// sign-up and sign-in pages for your application. When you set up a custom
// domain, you provide a certificate that you manage with AWS Certificate Manager
// (ACM). When necessary, you can use this operation to change the certificate
// that you applied to your custom domain.
//
// Usually, this is unnecessary following routine certificate renewal with ACM.
// When you renew your existing certificate in ACM, the ARN for your certificate
// remains the same, and your custom domain uses the new certificate automatically.
//
// However, if you replace your existing certificate with a new one, ACM gives
// the new certificate a new ARN. To apply the new certificate to your custom
// domain, you must provide this ARN to Amazon Cognito.
//
// When you add your new certificate in ACM, you must choose US East (N. Virginia)
// as the AWS Region.
//
// After you submit your request, Amazon Cognito requires up to 1 hour to distribute
// your new certificate to your custom domain.
//
// For more information about adding a custom domain to your user pool, see
// Using Your Own Domain for the Hosted UI (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html).
//
//    // Example sending a request using UpdateUserPoolDomainRequest.
//    req := client.UpdateUserPoolDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateUserPoolDomain
func (c *Client) UpdateUserPoolDomainRequest(input *types.UpdateUserPoolDomainInput) UpdateUserPoolDomainRequest {
	op := &aws.Operation{
		Name:       opUpdateUserPoolDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateUserPoolDomainInput{}
	}

	req := c.newRequest(op, input, &types.UpdateUserPoolDomainOutput{})
	return UpdateUserPoolDomainRequest{Request: req, Input: input, Copy: c.UpdateUserPoolDomainRequest}
}

// UpdateUserPoolDomainRequest is the request type for the
// UpdateUserPoolDomain API operation.
type UpdateUserPoolDomainRequest struct {
	*aws.Request
	Input *types.UpdateUserPoolDomainInput
	Copy  func(*types.UpdateUserPoolDomainInput) UpdateUserPoolDomainRequest
}

// Send marshals and sends the UpdateUserPoolDomain API request.
func (r UpdateUserPoolDomainRequest) Send(ctx context.Context) (*UpdateUserPoolDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserPoolDomainResponse{
		UpdateUserPoolDomainOutput: r.Request.Data.(*types.UpdateUserPoolDomainOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserPoolDomainResponse is the response type for the
// UpdateUserPoolDomain API operation.
type UpdateUserPoolDomainResponse struct {
	*types.UpdateUserPoolDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserPoolDomain request.
func (r *UpdateUserPoolDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
