// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetCodeBindingSourceInput struct {
	_ struct{} `type:"structure"`

	// Language is a required field
	Language *string `location:"uri" locationName:"language" type:"string" required:"true"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// SchemaName is a required field
	SchemaName *string `location:"uri" locationName:"schemaName" type:"string" required:"true"`

	SchemaVersion *string `location:"querystring" locationName:"schemaVersion" type:"string"`
}

// String returns the string representation
func (s GetCodeBindingSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCodeBindingSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCodeBindingSourceInput"}

	if s.Language == nil {
		invalidParams.Add(aws.NewErrParamRequired("Language"))
	}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if s.SchemaName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCodeBindingSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Language != nil {
		v := *s.Language

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "language", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "schemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "schemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetCodeBindingSourceOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	Body []byte `type:"blob"`
}

// String returns the string representation
func (s GetCodeBindingSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCodeBindingSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Body", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opGetCodeBindingSource = "GetCodeBindingSource"

// GetCodeBindingSourceRequest returns a request value for making API operation for
// Schemas.
//
// Get the code binding source URI.
//
//    // Example sending a request using GetCodeBindingSourceRequest.
//    req := client.GetCodeBindingSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/GetCodeBindingSource
func (c *Client) GetCodeBindingSourceRequest(input *GetCodeBindingSourceInput) GetCodeBindingSourceRequest {
	op := &aws.Operation{
		Name:       opGetCodeBindingSource,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas/name/{schemaName}/language/{language}/source",
	}

	if input == nil {
		input = &GetCodeBindingSourceInput{}
	}

	req := c.newRequest(op, input, &GetCodeBindingSourceOutput{})
	return GetCodeBindingSourceRequest{Request: req, Input: input, Copy: c.GetCodeBindingSourceRequest}
}

// GetCodeBindingSourceRequest is the request type for the
// GetCodeBindingSource API operation.
type GetCodeBindingSourceRequest struct {
	*aws.Request
	Input *GetCodeBindingSourceInput
	Copy  func(*GetCodeBindingSourceInput) GetCodeBindingSourceRequest
}

// Send marshals and sends the GetCodeBindingSource API request.
func (r GetCodeBindingSourceRequest) Send(ctx context.Context) (*GetCodeBindingSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCodeBindingSourceResponse{
		GetCodeBindingSourceOutput: r.Request.Data.(*GetCodeBindingSourceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCodeBindingSourceResponse is the response type for the
// GetCodeBindingSource API operation.
type GetCodeBindingSourceResponse struct {
	*GetCodeBindingSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCodeBindingSource request.
func (r *GetCodeBindingSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
