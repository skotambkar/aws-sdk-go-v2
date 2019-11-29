// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
)

const opPutCorsPolicy = "PutCorsPolicy"

// PutCorsPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Sets the cross-origin resource sharing (CORS) configuration on a container
// so that the container can service cross-origin requests. For example, you
// might want to enable a request whose origin is http://www.example.com to
// access your AWS Elemental MediaStore container at my.example.container.com
// by using the browser's XMLHttpRequest capability.
//
// To enable CORS on a container, you attach a CORS policy to the container.
// In the CORS policy, you configure rules that identify origins and the HTTP
// methods that can be executed on your container. The policy can contain up
// to 398,000 characters. You can add up to 100 rules to a CORS policy. If more
// than one rule applies, the service uses the first applicable rule listed.
//
// To learn more about CORS, see Cross-Origin Resource Sharing (CORS) in AWS
// Elemental MediaStore (https://docs.aws.amazon.com/mediastore/latest/ug/cors-policy.html).
//
//    // Example sending a request using PutCorsPolicyRequest.
//    req := client.PutCorsPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutCorsPolicy
func (c *Client) PutCorsPolicyRequest(input *types.PutCorsPolicyInput) PutCorsPolicyRequest {
	op := &aws.Operation{
		Name:       opPutCorsPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutCorsPolicyInput{}
	}

	req := c.newRequest(op, input, &types.PutCorsPolicyOutput{})
	return PutCorsPolicyRequest{Request: req, Input: input, Copy: c.PutCorsPolicyRequest}
}

// PutCorsPolicyRequest is the request type for the
// PutCorsPolicy API operation.
type PutCorsPolicyRequest struct {
	*aws.Request
	Input *types.PutCorsPolicyInput
	Copy  func(*types.PutCorsPolicyInput) PutCorsPolicyRequest
}

// Send marshals and sends the PutCorsPolicy API request.
func (r PutCorsPolicyRequest) Send(ctx context.Context) (*PutCorsPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutCorsPolicyResponse{
		PutCorsPolicyOutput: r.Request.Data.(*types.PutCorsPolicyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutCorsPolicyResponse is the response type for the
// PutCorsPolicy API operation.
type PutCorsPolicyResponse struct {
	*types.PutCorsPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutCorsPolicy request.
func (r *PutCorsPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
