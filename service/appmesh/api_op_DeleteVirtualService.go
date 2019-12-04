// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteVirtualServiceInput struct {
	_ struct{} `type:"structure"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// VirtualServiceName is a required field
	VirtualServiceName *string `location:"uri" locationName:"virtualServiceName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVirtualServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVirtualServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVirtualServiceInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.VirtualServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualServiceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVirtualServiceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualServiceName != nil {
		v := *s.VirtualServiceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualServiceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVirtualServiceOutput struct {
	_ struct{} `type:"structure" payload:"VirtualService"`

	// An object that represents a virtual service returned by a describe operation.
	//
	// VirtualService is a required field
	VirtualService *VirtualServiceData `locationName:"virtualService" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteVirtualServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVirtualServiceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualService != nil {
		v := s.VirtualService

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualService", v, metadata)
	}
	return nil
}

const opDeleteVirtualService = "DeleteVirtualService"

// DeleteVirtualServiceRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Deletes an existing virtual service.
//
//    // Example sending a request using DeleteVirtualServiceRequest.
//    req := client.DeleteVirtualServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DeleteVirtualService
func (c *Client) DeleteVirtualServiceRequest(input *DeleteVirtualServiceInput) DeleteVirtualServiceRequest {
	op := &aws.Operation{
		Name:       opDeleteVirtualService,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualServices/{virtualServiceName}",
	}

	if input == nil {
		input = &DeleteVirtualServiceInput{}
	}

	req := c.newRequest(op, input, &DeleteVirtualServiceOutput{})
	return DeleteVirtualServiceRequest{Request: req, Input: input, Copy: c.DeleteVirtualServiceRequest}
}

// DeleteVirtualServiceRequest is the request type for the
// DeleteVirtualService API operation.
type DeleteVirtualServiceRequest struct {
	*aws.Request
	Input *DeleteVirtualServiceInput
	Copy  func(*DeleteVirtualServiceInput) DeleteVirtualServiceRequest
}

// Send marshals and sends the DeleteVirtualService API request.
func (r DeleteVirtualServiceRequest) Send(ctx context.Context) (*DeleteVirtualServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVirtualServiceResponse{
		DeleteVirtualServiceOutput: r.Request.Data.(*DeleteVirtualServiceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVirtualServiceResponse is the response type for the
// DeleteVirtualService API operation.
type DeleteVirtualServiceResponse struct {
	*DeleteVirtualServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVirtualService request.
func (r *DeleteVirtualServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
