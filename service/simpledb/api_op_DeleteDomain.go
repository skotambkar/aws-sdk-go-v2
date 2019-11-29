// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/simpledb/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/simpledb/types"
)

const opDeleteDomain = "DeleteDomain"

// DeleteDomainRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// The DeleteDomain operation deletes a domain. Any items (and their attributes)
// in the domain are deleted as well. The DeleteDomain operation might take
// 10 or more seconds to complete.
//   Running
//    DeleteDomain
//  on a domain that does not exist or running the function multiple times using
//  the same domain name will not result in an error response.
//
//    // Example sending a request using DeleteDomainRequest.
//    req := client.DeleteDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteDomainRequest(input *types.DeleteDomainInput) DeleteDomainRequest {
	op := &aws.Operation{
		Name:       opDeleteDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteDomainInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDomainOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteDomainMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDomainRequest{Request: req, Input: input, Copy: c.DeleteDomainRequest}
}

// DeleteDomainRequest is the request type for the
// DeleteDomain API operation.
type DeleteDomainRequest struct {
	*aws.Request
	Input *types.DeleteDomainInput
	Copy  func(*types.DeleteDomainInput) DeleteDomainRequest
}

// Send marshals and sends the DeleteDomain API request.
func (r DeleteDomainRequest) Send(ctx context.Context) (*DeleteDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDomainResponse{
		DeleteDomainOutput: r.Request.Data.(*types.DeleteDomainOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDomainResponse is the response type for the
// DeleteDomain API operation.
type DeleteDomainResponse struct {
	*types.DeleteDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDomain request.
func (r *DeleteDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
