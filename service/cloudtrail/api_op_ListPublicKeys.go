// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

const opListPublicKeys = "ListPublicKeys"

// ListPublicKeysRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Returns all public keys whose private keys were used to sign the digest files
// within the specified time range. The public key is needed to validate digest
// files that were signed with its corresponding private key.
//
// CloudTrail uses different private/public key pairs per region. Each digest
// file is signed with a private key unique to its region. Therefore, when you
// validate a digest file from a particular region, you must look in the same
// region for its corresponding public key.
//
//    // Example sending a request using ListPublicKeysRequest.
//    req := client.ListPublicKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/ListPublicKeys
func (c *Client) ListPublicKeysRequest(input *types.ListPublicKeysInput) ListPublicKeysRequest {
	op := &aws.Operation{
		Name:       opListPublicKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPublicKeysInput{}
	}

	req := c.newRequest(op, input, &types.ListPublicKeysOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListPublicKeysMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPublicKeysRequest{Request: req, Input: input, Copy: c.ListPublicKeysRequest}
}

// ListPublicKeysRequest is the request type for the
// ListPublicKeys API operation.
type ListPublicKeysRequest struct {
	*aws.Request
	Input *types.ListPublicKeysInput
	Copy  func(*types.ListPublicKeysInput) ListPublicKeysRequest
}

// Send marshals and sends the ListPublicKeys API request.
func (r ListPublicKeysRequest) Send(ctx context.Context) (*ListPublicKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPublicKeysResponse{
		ListPublicKeysOutput: r.Request.Data.(*types.ListPublicKeysOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPublicKeysRequestPaginator returns a paginator for ListPublicKeys.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPublicKeysRequest(input)
//   p := cloudtrail.NewListPublicKeysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPublicKeysPaginator(req ListPublicKeysRequest) ListPublicKeysPaginator {
	return ListPublicKeysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPublicKeysInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListPublicKeysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPublicKeysPaginator struct {
	aws.Pager
}

func (p *ListPublicKeysPaginator) CurrentPage() *types.ListPublicKeysOutput {
	return p.Pager.CurrentPage().(*types.ListPublicKeysOutput)
}

// ListPublicKeysResponse is the response type for the
// ListPublicKeys API operation.
type ListPublicKeysResponse struct {
	*types.ListPublicKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPublicKeys request.
func (r *ListPublicKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
