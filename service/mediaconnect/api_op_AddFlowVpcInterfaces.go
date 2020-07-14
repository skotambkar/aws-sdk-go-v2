// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Adds VPC interfaces to an existing flow.
type AddFlowVpcInterfacesInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// A list of VPC interfaces that you want to add.
	//
	// VpcInterfaces is a required field
	VpcInterfaces []VpcInterfaceRequest `locationName:"vpcInterfaces" type:"list" required:"true"`
}

// String returns the string representation
func (s AddFlowVpcInterfacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddFlowVpcInterfacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddFlowVpcInterfacesInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if s.VpcInterfaces == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcInterfaces"))
	}
	if s.VpcInterfaces != nil {
		for i, v := range s.VpcInterfaces {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "VpcInterfaces", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddFlowVpcInterfacesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VpcInterfaces != nil {
		v := s.VpcInterfaces

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "vpcInterfaces", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful AddFlowVpcInterfaces request. The response includes
// the details of the newly added VPC interfaces.
type AddFlowVpcInterfacesOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that these VPC interfaces were added to.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The details of the newly added VPC interfaces.
	VpcInterfaces []VpcInterface `locationName:"vpcInterfaces" type:"list"`
}

// String returns the string representation
func (s AddFlowVpcInterfacesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddFlowVpcInterfacesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VpcInterfaces != nil {
		v := s.VpcInterfaces

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "vpcInterfaces", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opAddFlowVpcInterfaces = "AddFlowVpcInterfaces"

// AddFlowVpcInterfacesRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Adds VPC interfaces to flow
//
//    // Example sending a request using AddFlowVpcInterfacesRequest.
//    req := client.AddFlowVpcInterfacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/AddFlowVpcInterfaces
func (c *Client) AddFlowVpcInterfacesRequest(input *AddFlowVpcInterfacesInput) AddFlowVpcInterfacesRequest {
	op := &aws.Operation{
		Name:       opAddFlowVpcInterfaces,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/flows/{flowArn}/vpcInterfaces",
	}

	if input == nil {
		input = &AddFlowVpcInterfacesInput{}
	}

	req := c.newRequest(op, input, &AddFlowVpcInterfacesOutput{})

	return AddFlowVpcInterfacesRequest{Request: req, Input: input, Copy: c.AddFlowVpcInterfacesRequest}
}

// AddFlowVpcInterfacesRequest is the request type for the
// AddFlowVpcInterfaces API operation.
type AddFlowVpcInterfacesRequest struct {
	*aws.Request
	Input *AddFlowVpcInterfacesInput
	Copy  func(*AddFlowVpcInterfacesInput) AddFlowVpcInterfacesRequest
}

// Send marshals and sends the AddFlowVpcInterfaces API request.
func (r AddFlowVpcInterfacesRequest) Send(ctx context.Context) (*AddFlowVpcInterfacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddFlowVpcInterfacesResponse{
		AddFlowVpcInterfacesOutput: r.Request.Data.(*AddFlowVpcInterfacesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddFlowVpcInterfacesResponse is the response type for the
// AddFlowVpcInterfaces API operation.
type AddFlowVpcInterfacesResponse struct {
	*AddFlowVpcInterfacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddFlowVpcInterfaces request.
func (r *AddFlowVpcInterfacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}