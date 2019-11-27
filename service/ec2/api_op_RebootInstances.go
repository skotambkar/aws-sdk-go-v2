// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opRebootInstances = "RebootInstances"

// RebootInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Requests a reboot of the specified instances. This operation is asynchronous;
// it only queues a request to reboot the specified instances. The operation
// succeeds if the instances are valid and belong to you. Requests to reboot
// terminated instances are ignored.
//
// If an instance does not cleanly shut down within four minutes, Amazon EC2
// performs a hard reboot.
//
// For more information about troubleshooting, see Getting Console Output and
// Rebooting Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-console.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using RebootInstancesRequest.
//    req := client.RebootInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RebootInstances
func (c *Client) RebootInstancesRequest(input *types.RebootInstancesInput) RebootInstancesRequest {
	op := &aws.Operation{
		Name:       opRebootInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RebootInstancesInput{}
	}

	req := c.newRequest(op, input, &types.RebootInstancesOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RebootInstancesRequest{Request: req, Input: input, Copy: c.RebootInstancesRequest}
}

// RebootInstancesRequest is the request type for the
// RebootInstances API operation.
type RebootInstancesRequest struct {
	*aws.Request
	Input *types.RebootInstancesInput
	Copy  func(*types.RebootInstancesInput) RebootInstancesRequest
}

// Send marshals and sends the RebootInstances API request.
func (r RebootInstancesRequest) Send(ctx context.Context) (*RebootInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootInstancesResponse{
		RebootInstancesOutput: r.Request.Data.(*types.RebootInstancesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootInstancesResponse is the response type for the
// RebootInstances API operation.
type RebootInstancesResponse struct {
	*types.RebootInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootInstances request.
func (r *RebootInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
