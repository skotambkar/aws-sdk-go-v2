// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateIpGroupsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory.
	//
	// DirectoryId is a required field
	DirectoryId *string `min:"10" type:"string" required:"true"`

	// The identifiers of one or more IP access control groups.
	//
	// GroupIds is a required field
	GroupIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DisassociateIpGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateIpGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateIpGroupsInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}
	if s.DirectoryId != nil && len(*s.DirectoryId) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("DirectoryId", 10))
	}

	if s.GroupIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateIpGroupsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateIpGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateIpGroups = "DisassociateIpGroups"

// DisassociateIpGroupsRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Disassociates the specified IP access control group from the specified directory.
//
//    // Example sending a request using DisassociateIpGroupsRequest.
//    req := client.DisassociateIpGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DisassociateIpGroups
func (c *Client) DisassociateIpGroupsRequest(input *DisassociateIpGroupsInput) DisassociateIpGroupsRequest {
	op := &aws.Operation{
		Name:       opDisassociateIpGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateIpGroupsInput{}
	}

	req := c.newRequest(op, input, &DisassociateIpGroupsOutput{})
	return DisassociateIpGroupsRequest{Request: req, Input: input, Copy: c.DisassociateIpGroupsRequest}
}

// DisassociateIpGroupsRequest is the request type for the
// DisassociateIpGroups API operation.
type DisassociateIpGroupsRequest struct {
	*aws.Request
	Input *DisassociateIpGroupsInput
	Copy  func(*DisassociateIpGroupsInput) DisassociateIpGroupsRequest
}

// Send marshals and sends the DisassociateIpGroups API request.
func (r DisassociateIpGroupsRequest) Send(ctx context.Context) (*DisassociateIpGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateIpGroupsResponse{
		DisassociateIpGroupsOutput: r.Request.Data.(*DisassociateIpGroupsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateIpGroupsResponse is the response type for the
// DisassociateIpGroups API operation.
type DisassociateIpGroupsResponse struct {
	*DisassociateIpGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateIpGroups request.
func (r *DisassociateIpGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
