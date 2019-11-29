// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

const opRemoveListenerCertificates = "RemoveListenerCertificates"

// RemoveListenerCertificatesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Removes the specified certificate from the certificate list for the specified
// HTTPS or TLS listener.
//
// You can't remove the default certificate for a listener. To replace the default
// certificate, call ModifyListener.
//
// To list the certificates for your listener, use DescribeListenerCertificates.
//
//    // Example sending a request using RemoveListenerCertificatesRequest.
//    req := client.RemoveListenerCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/RemoveListenerCertificates
func (c *Client) RemoveListenerCertificatesRequest(input *types.RemoveListenerCertificatesInput) RemoveListenerCertificatesRequest {
	op := &aws.Operation{
		Name:       opRemoveListenerCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RemoveListenerCertificatesInput{}
	}

	req := c.newRequest(op, input, &types.RemoveListenerCertificatesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.RemoveListenerCertificatesMarshaler{Input: input}.GetNamedBuildHandler())

	return RemoveListenerCertificatesRequest{Request: req, Input: input, Copy: c.RemoveListenerCertificatesRequest}
}

// RemoveListenerCertificatesRequest is the request type for the
// RemoveListenerCertificates API operation.
type RemoveListenerCertificatesRequest struct {
	*aws.Request
	Input *types.RemoveListenerCertificatesInput
	Copy  func(*types.RemoveListenerCertificatesInput) RemoveListenerCertificatesRequest
}

// Send marshals and sends the RemoveListenerCertificates API request.
func (r RemoveListenerCertificatesRequest) Send(ctx context.Context) (*RemoveListenerCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveListenerCertificatesResponse{
		RemoveListenerCertificatesOutput: r.Request.Data.(*types.RemoveListenerCertificatesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveListenerCertificatesResponse is the response type for the
// RemoveListenerCertificates API operation.
type RemoveListenerCertificatesResponse struct {
	*types.RemoveListenerCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveListenerCertificates request.
func (r *RemoveListenerCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
