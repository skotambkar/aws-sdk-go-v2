// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
)

const opDeleteRevision = "DeleteRevision"

// DeleteRevisionRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation deletes a revision.
//
//    // Example sending a request using DeleteRevisionRequest.
//    req := client.DeleteRevisionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/DeleteRevision
func (c *Client) DeleteRevisionRequest(input *types.DeleteRevisionInput) DeleteRevisionRequest {
	op := &aws.Operation{
		Name:       opDeleteRevision,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/data-sets/{DataSetId}/revisions/{RevisionId}",
	}

	if input == nil {
		input = &types.DeleteRevisionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRevisionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteRevisionMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRevisionRequest{Request: req, Input: input, Copy: c.DeleteRevisionRequest}
}

// DeleteRevisionRequest is the request type for the
// DeleteRevision API operation.
type DeleteRevisionRequest struct {
	*aws.Request
	Input *types.DeleteRevisionInput
	Copy  func(*types.DeleteRevisionInput) DeleteRevisionRequest
}

// Send marshals and sends the DeleteRevision API request.
func (r DeleteRevisionRequest) Send(ctx context.Context) (*DeleteRevisionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRevisionResponse{
		DeleteRevisionOutput: r.Request.Data.(*types.DeleteRevisionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRevisionResponse is the response type for the
// DeleteRevision API operation.
type DeleteRevisionResponse struct {
	*types.DeleteRevisionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRevision request.
func (r *DeleteRevisionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
