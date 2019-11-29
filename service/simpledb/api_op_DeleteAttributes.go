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

const opDeleteAttributes = "DeleteAttributes"

// DeleteAttributesRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// Deletes one or more attributes associated with an item. If all attributes
// of the item are deleted, the item is deleted.
//   If
//    DeleteAttributes
//  is called without being passed any attributes or values specified, all the
//  attributes for the item are deleted.
// DeleteAttributes is an idempotent operation; running it multiple times on
// the same item or attribute does not result in an error response.
//
// Because Amazon SimpleDB makes multiple copies of item data and uses an eventual
// consistency update model, performing a GetAttributes or Select operation
// (read) immediately after a DeleteAttributes or PutAttributes operation (write)
// might not return updated item data.
//
//    // Example sending a request using DeleteAttributesRequest.
//    req := client.DeleteAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteAttributesRequest(input *types.DeleteAttributesInput) DeleteAttributesRequest {
	op := &aws.Operation{
		Name:       opDeleteAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAttributesInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAttributesRequest{Request: req, Input: input, Copy: c.DeleteAttributesRequest}
}

// DeleteAttributesRequest is the request type for the
// DeleteAttributes API operation.
type DeleteAttributesRequest struct {
	*aws.Request
	Input *types.DeleteAttributesInput
	Copy  func(*types.DeleteAttributesInput) DeleteAttributesRequest
}

// Send marshals and sends the DeleteAttributes API request.
func (r DeleteAttributesRequest) Send(ctx context.Context) (*DeleteAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAttributesResponse{
		DeleteAttributesOutput: r.Request.Data.(*types.DeleteAttributesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAttributesResponse is the response type for the
// DeleteAttributes API operation.
type DeleteAttributesResponse struct {
	*types.DeleteAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAttributes request.
func (r *DeleteAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
