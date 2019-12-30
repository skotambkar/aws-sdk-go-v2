// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for get backend environment request.
type GetBackendEnvironmentInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name for the backend environment.
	//
	// EnvironmentName is a required field
	EnvironmentName *string `location:"uri" locationName:"environmentName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBackendEnvironmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBackendEnvironmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBackendEnvironmentInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.EnvironmentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentName"))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackendEnvironmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnvironmentName != nil {
		v := *s.EnvironmentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "environmentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for get backend environment result.
type GetBackendEnvironmentOutput struct {
	_ struct{} `type:"structure"`

	// Backend environment structure for an an Amplify App.
	//
	// BackendEnvironment is a required field
	BackendEnvironment *BackendEnvironment `locationName:"backendEnvironment" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetBackendEnvironmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackendEnvironmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackendEnvironment != nil {
		v := s.BackendEnvironment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "backendEnvironment", v, metadata)
	}
	return nil
}

const opGetBackendEnvironment = "GetBackendEnvironment"

// GetBackendEnvironmentRequest returns a request value for making API operation for
// AWS Amplify.
//
// Retrieves a backend environment for an Amplify App.
//
//    // Example sending a request using GetBackendEnvironmentRequest.
//    req := client.GetBackendEnvironmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetBackendEnvironment
func (c *Client) GetBackendEnvironmentRequest(input *GetBackendEnvironmentInput) GetBackendEnvironmentRequest {
	op := &aws.Operation{
		Name:       opGetBackendEnvironment,
		HTTPMethod: "GET",
		HTTPPath:   "/apps/{appId}/backendenvironments/{environmentName}",
	}

	if input == nil {
		input = &GetBackendEnvironmentInput{}
	}

	req := c.newRequest(op, input, &GetBackendEnvironmentOutput{})
	return GetBackendEnvironmentRequest{Request: req, Input: input, Copy: c.GetBackendEnvironmentRequest}
}

// GetBackendEnvironmentRequest is the request type for the
// GetBackendEnvironment API operation.
type GetBackendEnvironmentRequest struct {
	*aws.Request
	Input *GetBackendEnvironmentInput
	Copy  func(*GetBackendEnvironmentInput) GetBackendEnvironmentRequest
}

// Send marshals and sends the GetBackendEnvironment API request.
func (r GetBackendEnvironmentRequest) Send(ctx context.Context) (*GetBackendEnvironmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackendEnvironmentResponse{
		GetBackendEnvironmentOutput: r.Request.Data.(*GetBackendEnvironmentOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackendEnvironmentResponse is the response type for the
// GetBackendEnvironment API operation.
type GetBackendEnvironmentResponse struct {
	*GetBackendEnvironmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackendEnvironment request.
func (r *GetBackendEnvironmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
