// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListIncomingTypedLinks = "ListIncomingTypedLinks"

// ListIncomingTypedLinksRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns a paginated list of all the incoming TypedLinkSpecifier information
// for an object. It also supports filtering by typed link facet and identity
// attributes. For more information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using ListIncomingTypedLinksRequest.
//    req := client.ListIncomingTypedLinksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListIncomingTypedLinks
func (c *Client) ListIncomingTypedLinksRequest(input *types.ListIncomingTypedLinksInput) ListIncomingTypedLinksRequest {
	op := &aws.Operation{
		Name:       opListIncomingTypedLinks,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/incoming",
	}

	if input == nil {
		input = &types.ListIncomingTypedLinksInput{}
	}

	req := c.newRequest(op, input, &types.ListIncomingTypedLinksOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListIncomingTypedLinksMarshaler{Input: input}.GetNamedBuildHandler())

	return ListIncomingTypedLinksRequest{Request: req, Input: input, Copy: c.ListIncomingTypedLinksRequest}
}

// ListIncomingTypedLinksRequest is the request type for the
// ListIncomingTypedLinks API operation.
type ListIncomingTypedLinksRequest struct {
	*aws.Request
	Input *types.ListIncomingTypedLinksInput
	Copy  func(*types.ListIncomingTypedLinksInput) ListIncomingTypedLinksRequest
}

// Send marshals and sends the ListIncomingTypedLinks API request.
func (r ListIncomingTypedLinksRequest) Send(ctx context.Context) (*ListIncomingTypedLinksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIncomingTypedLinksResponse{
		ListIncomingTypedLinksOutput: r.Request.Data.(*types.ListIncomingTypedLinksOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIncomingTypedLinksResponse is the response type for the
// ListIncomingTypedLinks API operation.
type ListIncomingTypedLinksResponse struct {
	*types.ListIncomingTypedLinksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIncomingTypedLinks request.
func (r *ListIncomingTypedLinksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
