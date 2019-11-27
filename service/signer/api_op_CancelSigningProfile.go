// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
)

const opCancelSigningProfile = "CancelSigningProfile"

// CancelSigningProfileRequest returns a request value for making API operation for
// AWS Signer.
//
// Changes the state of an ACTIVE signing profile to CANCELED. A canceled profile
// is still viewable with the ListSigningProfiles operation, but it cannot perform
// new signing jobs, and is deleted two years after cancelation.
//
//    // Example sending a request using CancelSigningProfileRequest.
//    req := client.CancelSigningProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/CancelSigningProfile
func (c *Client) CancelSigningProfileRequest(input *types.CancelSigningProfileInput) CancelSigningProfileRequest {
	op := &aws.Operation{
		Name:       opCancelSigningProfile,
		HTTPMethod: "DELETE",
		HTTPPath:   "/signing-profiles/{profileName}",
	}

	if input == nil {
		input = &types.CancelSigningProfileInput{}
	}

	req := c.newRequest(op, input, &types.CancelSigningProfileOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CancelSigningProfileRequest{Request: req, Input: input, Copy: c.CancelSigningProfileRequest}
}

// CancelSigningProfileRequest is the request type for the
// CancelSigningProfile API operation.
type CancelSigningProfileRequest struct {
	*aws.Request
	Input *types.CancelSigningProfileInput
	Copy  func(*types.CancelSigningProfileInput) CancelSigningProfileRequest
}

// Send marshals and sends the CancelSigningProfile API request.
func (r CancelSigningProfileRequest) Send(ctx context.Context) (*CancelSigningProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelSigningProfileResponse{
		CancelSigningProfileOutput: r.Request.Data.(*types.CancelSigningProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelSigningProfileResponse is the response type for the
// CancelSigningProfile API operation.
type CancelSigningProfileResponse struct {
	*types.CancelSigningProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelSigningProfile request.
func (r *CancelSigningProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
