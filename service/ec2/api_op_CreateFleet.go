// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateFleet = "CreateFleet"

// CreateFleetRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Launches an EC2 Fleet.
//
// You can create a single EC2 Fleet that includes multiple launch specifications
// that vary by instance type, AMI, Availability Zone, or subnet.
//
// For more information, see Launching an EC2 Fleet (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using CreateFleetRequest.
//    req := client.CreateFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateFleet
func (c *Client) CreateFleetRequest(input *types.CreateFleetInput) CreateFleetRequest {
	op := &aws.Operation{
		Name:       opCreateFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateFleetInput{}
	}

	req := c.newRequest(op, input, &types.CreateFleetOutput{})
	return CreateFleetRequest{Request: req, Input: input, Copy: c.CreateFleetRequest}
}

// CreateFleetRequest is the request type for the
// CreateFleet API operation.
type CreateFleetRequest struct {
	*aws.Request
	Input *types.CreateFleetInput
	Copy  func(*types.CreateFleetInput) CreateFleetRequest
}

// Send marshals and sends the CreateFleet API request.
func (r CreateFleetRequest) Send(ctx context.Context) (*CreateFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFleetResponse{
		CreateFleetOutput: r.Request.Data.(*types.CreateFleetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFleetResponse is the response type for the
// CreateFleet API operation.
type CreateFleetResponse struct {
	*types.CreateFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFleet request.
func (r *CreateFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
