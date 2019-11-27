// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opValidateSecurityProfileBehaviors = "ValidateSecurityProfileBehaviors"

// ValidateSecurityProfileBehaviorsRequest returns a request value for making API operation for
// AWS IoT.
//
// Validates a Device Defender security profile behaviors specification.
//
//    // Example sending a request using ValidateSecurityProfileBehaviorsRequest.
//    req := client.ValidateSecurityProfileBehaviorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ValidateSecurityProfileBehaviorsRequest(input *types.ValidateSecurityProfileBehaviorsInput) ValidateSecurityProfileBehaviorsRequest {
	op := &aws.Operation{
		Name:       opValidateSecurityProfileBehaviors,
		HTTPMethod: "POST",
		HTTPPath:   "/security-profile-behaviors/validate",
	}

	if input == nil {
		input = &types.ValidateSecurityProfileBehaviorsInput{}
	}

	req := c.newRequest(op, input, &types.ValidateSecurityProfileBehaviorsOutput{})
	return ValidateSecurityProfileBehaviorsRequest{Request: req, Input: input, Copy: c.ValidateSecurityProfileBehaviorsRequest}
}

// ValidateSecurityProfileBehaviorsRequest is the request type for the
// ValidateSecurityProfileBehaviors API operation.
type ValidateSecurityProfileBehaviorsRequest struct {
	*aws.Request
	Input *types.ValidateSecurityProfileBehaviorsInput
	Copy  func(*types.ValidateSecurityProfileBehaviorsInput) ValidateSecurityProfileBehaviorsRequest
}

// Send marshals and sends the ValidateSecurityProfileBehaviors API request.
func (r ValidateSecurityProfileBehaviorsRequest) Send(ctx context.Context) (*ValidateSecurityProfileBehaviorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ValidateSecurityProfileBehaviorsResponse{
		ValidateSecurityProfileBehaviorsOutput: r.Request.Data.(*types.ValidateSecurityProfileBehaviorsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ValidateSecurityProfileBehaviorsResponse is the response type for the
// ValidateSecurityProfileBehaviors API operation.
type ValidateSecurityProfileBehaviorsResponse struct {
	*types.ValidateSecurityProfileBehaviorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ValidateSecurityProfileBehaviors request.
func (r *ValidateSecurityProfileBehaviorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
