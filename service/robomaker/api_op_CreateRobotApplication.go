// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
)

const opCreateRobotApplication = "CreateRobotApplication"

// CreateRobotApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Creates a robot application.
//
//    // Example sending a request using CreateRobotApplicationRequest.
//    req := client.CreateRobotApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateRobotApplication
func (c *Client) CreateRobotApplicationRequest(input *types.CreateRobotApplicationInput) CreateRobotApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateRobotApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/createRobotApplication",
	}

	if input == nil {
		input = &types.CreateRobotApplicationInput{}
	}

	req := c.newRequest(op, input, &types.CreateRobotApplicationOutput{})
	return CreateRobotApplicationRequest{Request: req, Input: input, Copy: c.CreateRobotApplicationRequest}
}

// CreateRobotApplicationRequest is the request type for the
// CreateRobotApplication API operation.
type CreateRobotApplicationRequest struct {
	*aws.Request
	Input *types.CreateRobotApplicationInput
	Copy  func(*types.CreateRobotApplicationInput) CreateRobotApplicationRequest
}

// Send marshals and sends the CreateRobotApplication API request.
func (r CreateRobotApplicationRequest) Send(ctx context.Context) (*CreateRobotApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRobotApplicationResponse{
		CreateRobotApplicationOutput: r.Request.Data.(*types.CreateRobotApplicationOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRobotApplicationResponse is the response type for the
// CreateRobotApplication API operation.
type CreateRobotApplicationResponse struct {
	*types.CreateRobotApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRobotApplication request.
func (r *CreateRobotApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
