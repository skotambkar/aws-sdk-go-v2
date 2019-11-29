// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

const opDetachInstances = "DetachInstances"

// DetachInstancesRequest returns a request value for making API operation for
// Auto Scaling.
//
// Removes one or more instances from the specified Auto Scaling group.
//
// After the instances are detached, you can manage them independent of the
// Auto Scaling group.
//
// If you do not specify the option to decrement the desired capacity, Amazon
// EC2 Auto Scaling launches instances to replace the ones that are detached.
//
// If there is a Classic Load Balancer attached to the Auto Scaling group, the
// instances are deregistered from the load balancer. If there are target groups
// attached to the Auto Scaling group, the instances are deregistered from the
// target groups.
//
// For more information, see Detach EC2 Instances from Your Auto Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/detach-instance-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using DetachInstancesRequest.
//    req := client.DetachInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DetachInstances
func (c *Client) DetachInstancesRequest(input *types.DetachInstancesInput) DetachInstancesRequest {
	op := &aws.Operation{
		Name:       opDetachInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DetachInstancesInput{}
	}

	req := c.newRequest(op, input, &types.DetachInstancesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DetachInstancesMarshaler{Input: input}.GetNamedBuildHandler())

	return DetachInstancesRequest{Request: req, Input: input, Copy: c.DetachInstancesRequest}
}

// DetachInstancesRequest is the request type for the
// DetachInstances API operation.
type DetachInstancesRequest struct {
	*aws.Request
	Input *types.DetachInstancesInput
	Copy  func(*types.DetachInstancesInput) DetachInstancesRequest
}

// Send marshals and sends the DetachInstances API request.
func (r DetachInstancesRequest) Send(ctx context.Context) (*DetachInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachInstancesResponse{
		DetachInstancesOutput: r.Request.Data.(*types.DetachInstancesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachInstancesResponse is the response type for the
// DetachInstances API operation.
type DetachInstancesResponse struct {
	*types.DetachInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachInstances request.
func (r *DetachInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
