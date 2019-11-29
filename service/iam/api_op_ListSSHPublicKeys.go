// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opListSSHPublicKeys = "ListSSHPublicKeys"

// ListSSHPublicKeysRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Returns information about the SSH public keys associated with the specified
// IAM user. If none exists, the operation returns an empty list.
//
// The SSH public keys returned by this operation are used only for authenticating
// the IAM user to an AWS CodeCommit repository. For more information about
// using SSH keys to authenticate to an AWS CodeCommit repository, see Set up
// AWS CodeCommit for SSH Connections (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html)
// in the AWS CodeCommit User Guide.
//
// Although each user is limited to a small number of keys, you can still paginate
// the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using ListSSHPublicKeysRequest.
//    req := client.ListSSHPublicKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListSSHPublicKeys
func (c *Client) ListSSHPublicKeysRequest(input *types.ListSSHPublicKeysInput) ListSSHPublicKeysRequest {
	op := &aws.Operation{
		Name:       opListSSHPublicKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &types.ListSSHPublicKeysInput{}
	}

	req := c.newRequest(op, input, &types.ListSSHPublicKeysOutput{})
	return ListSSHPublicKeysRequest{Request: req, Input: input, Copy: c.ListSSHPublicKeysRequest}
}

// ListSSHPublicKeysRequest is the request type for the
// ListSSHPublicKeys API operation.
type ListSSHPublicKeysRequest struct {
	*aws.Request
	Input *types.ListSSHPublicKeysInput
	Copy  func(*types.ListSSHPublicKeysInput) ListSSHPublicKeysRequest
}

// Send marshals and sends the ListSSHPublicKeys API request.
func (r ListSSHPublicKeysRequest) Send(ctx context.Context) (*ListSSHPublicKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSSHPublicKeysResponse{
		ListSSHPublicKeysOutput: r.Request.Data.(*types.ListSSHPublicKeysOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSSHPublicKeysRequestPaginator returns a paginator for ListSSHPublicKeys.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSSHPublicKeysRequest(input)
//   p := iam.NewListSSHPublicKeysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSSHPublicKeysPaginator(req ListSSHPublicKeysRequest) ListSSHPublicKeysPaginator {
	return ListSSHPublicKeysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListSSHPublicKeysInput
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

// ListSSHPublicKeysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSSHPublicKeysPaginator struct {
	aws.Pager
}

func (p *ListSSHPublicKeysPaginator) CurrentPage() *types.ListSSHPublicKeysOutput {
	return p.Pager.CurrentPage().(*types.ListSSHPublicKeysOutput)
}

// ListSSHPublicKeysResponse is the response type for the
// ListSSHPublicKeys API operation.
type ListSSHPublicKeysResponse struct {
	*types.ListSSHPublicKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSSHPublicKeys request.
func (r *ListSSHPublicKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
