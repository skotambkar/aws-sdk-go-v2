// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/acm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
)

const opListTagsForCertificate = "ListTagsForCertificate"

// ListTagsForCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Lists the tags that have been applied to the ACM certificate. Use the certificate's
// Amazon Resource Name (ARN) to specify the certificate. To add a tag to an
// ACM certificate, use the AddTagsToCertificate action. To delete a tag, use
// the RemoveTagsFromCertificate action.
//
//    // Example sending a request using ListTagsForCertificateRequest.
//    req := client.ListTagsForCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ListTagsForCertificate
func (c *Client) ListTagsForCertificateRequest(input *types.ListTagsForCertificateInput) ListTagsForCertificateRequest {
	op := &aws.Operation{
		Name:       opListTagsForCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListTagsForCertificateInput{}
	}

	req := c.newRequest(op, input, &types.ListTagsForCertificateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListTagsForCertificateMarshaler{Input: input}.GetNamedBuildHandler())

	return ListTagsForCertificateRequest{Request: req, Input: input, Copy: c.ListTagsForCertificateRequest}
}

// ListTagsForCertificateRequest is the request type for the
// ListTagsForCertificate API operation.
type ListTagsForCertificateRequest struct {
	*aws.Request
	Input *types.ListTagsForCertificateInput
	Copy  func(*types.ListTagsForCertificateInput) ListTagsForCertificateRequest
}

// Send marshals and sends the ListTagsForCertificate API request.
func (r ListTagsForCertificateRequest) Send(ctx context.Context) (*ListTagsForCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForCertificateResponse{
		ListTagsForCertificateOutput: r.Request.Data.(*types.ListTagsForCertificateOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForCertificateResponse is the response type for the
// ListTagsForCertificate API operation.
type ListTagsForCertificateResponse struct {
	*types.ListTagsForCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForCertificate request.
func (r *ListTagsForCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
