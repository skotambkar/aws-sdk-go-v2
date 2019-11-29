// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

const opAssociateKmsKey = "AssociateKmsKey"

// AssociateKmsKeyRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Associates the specified AWS Key Management Service (AWS KMS) customer master
// key (CMK) with the specified log group.
//
// Associating an AWS KMS CMK with a log group overrides any existing associations
// between the log group and a CMK. After a CMK is associated with a log group,
// all newly ingested data for the log group is encrypted using the CMK. This
// association is stored as long as the data encrypted with the CMK is still
// within Amazon CloudWatch Logs. This enables Amazon CloudWatch Logs to decrypt
// this data whenever it is requested.
//
// Note that it can take up to 5 minutes for this operation to take effect.
//
// If you attempt to associate a CMK with a log group but the CMK does not exist
// or the CMK is disabled, you will receive an InvalidParameterException error.
//
//    // Example sending a request using AssociateKmsKeyRequest.
//    req := client.AssociateKmsKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/AssociateKmsKey
func (c *Client) AssociateKmsKeyRequest(input *types.AssociateKmsKeyInput) AssociateKmsKeyRequest {
	op := &aws.Operation{
		Name:       opAssociateKmsKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateKmsKeyInput{}
	}

	req := c.newRequest(op, input, &types.AssociateKmsKeyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AssociateKmsKeyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AssociateKmsKeyRequest{Request: req, Input: input, Copy: c.AssociateKmsKeyRequest}
}

// AssociateKmsKeyRequest is the request type for the
// AssociateKmsKey API operation.
type AssociateKmsKeyRequest struct {
	*aws.Request
	Input *types.AssociateKmsKeyInput
	Copy  func(*types.AssociateKmsKeyInput) AssociateKmsKeyRequest
}

// Send marshals and sends the AssociateKmsKey API request.
func (r AssociateKmsKeyRequest) Send(ctx context.Context) (*AssociateKmsKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateKmsKeyResponse{
		AssociateKmsKeyOutput: r.Request.Data.(*types.AssociateKmsKeyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateKmsKeyResponse is the response type for the
// AssociateKmsKey API operation.
type AssociateKmsKeyResponse struct {
	*types.AssociateKmsKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateKmsKey request.
func (r *AssociateKmsKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
