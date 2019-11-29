// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
)

const opIndexDocuments = "IndexDocuments"

// IndexDocumentsRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Tells the search domain to start indexing its documents using the latest
// indexing options. This operation must be invoked to activate options whose
// OptionStatus is RequiresIndexDocuments.
//
//    // Example sending a request using IndexDocumentsRequest.
//    req := client.IndexDocumentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) IndexDocumentsRequest(input *types.IndexDocumentsInput) IndexDocumentsRequest {
	op := &aws.Operation{
		Name:       opIndexDocuments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.IndexDocumentsInput{}
	}

	req := c.newRequest(op, input, &types.IndexDocumentsOutput{})
	return IndexDocumentsRequest{Request: req, Input: input, Copy: c.IndexDocumentsRequest}
}

// IndexDocumentsRequest is the request type for the
// IndexDocuments API operation.
type IndexDocumentsRequest struct {
	*aws.Request
	Input *types.IndexDocumentsInput
	Copy  func(*types.IndexDocumentsInput) IndexDocumentsRequest
}

// Send marshals and sends the IndexDocuments API request.
func (r IndexDocumentsRequest) Send(ctx context.Context) (*IndexDocumentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &IndexDocumentsResponse{
		IndexDocumentsOutput: r.Request.Data.(*types.IndexDocumentsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// IndexDocumentsResponse is the response type for the
// IndexDocuments API operation.
type IndexDocumentsResponse struct {
	*types.IndexDocumentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// IndexDocuments request.
func (r *IndexDocumentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
