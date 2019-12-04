// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerruntime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type InvokeEndpointInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The desired MIME type of the inference in the response.
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model.
	//
	// For information about the format of the request body, see Common Data Formats—Inference
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true" sensitive:"true"`

	// The MIME type of the input data in the request body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Provides additional information about a request for an inference submitted
	// to a model hosted at an Amazon SageMaker endpoint. The information is an
	// opaque value that is forwarded verbatim. You could use this value, for example,
	// to provide an ID that you can use to track a request or to provide other
	// metadata that a service endpoint was programmed to process. The value must
	// consist of no more than 1024 visible US-ASCII characters as specified in
	// Section 3.3.6. Field Value Components (https://tools.ietf.org/html/rfc7230#section-3.2.6)
	// of the Hypertext Transfer Protocol (HTTP/1.1). This feature is currently
	// supported in the AWS SDKs but not in the Amazon SageMaker Python SDK.
	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string" sensitive:"true"`

	// The name of the endpoint that you specified when you created the endpoint
	// using the CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html)
	// API.
	//
	// EndpointName is a required field
	EndpointName *string `location:"uri" locationName:"EndpointName" type:"string" required:"true"`

	// Specifies the model to be requested for an inference when invoking a multi-model
	// endpoint.
	TargetModel *string `location:"header" locationName:"X-Amzn-SageMaker-Target-Model" min:"1" type:"string"`
}

// String returns the string representation
func (s InvokeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InvokeEndpointInput"}

	if s.Body == nil {
		invalidParams.Add(aws.NewErrParamRequired("Body"))
	}

	if s.EndpointName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointName"))
	}
	if s.TargetModel != nil && len(*s.TargetModel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetModel", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InvokeEndpointInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Accept != nil {
		v := *s.Accept

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Accept", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CustomAttributes != nil {
		v := *s.CustomAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-SageMaker-Custom-Attributes", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TargetModel != nil {
		v := *s.TargetModel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-SageMaker-Target-Model", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointName != nil {
		v := *s.EndpointName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EndpointName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Body", protocol.BytesStream(v), metadata)
	}
	return nil
}

type InvokeEndpointOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Includes the inference provided by the model.
	//
	// For information about the format of the response body, see Common Data Formats—Inference
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true" sensitive:"true"`

	// The MIME type of the inference returned in the response body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Provides additional information in the response about the inference returned
	// by a model hosted at an Amazon SageMaker endpoint. The information is an
	// opaque value that is forwarded verbatim. You could use this value, for example,
	// to return an ID received in the CustomAttributes header of a request or other
	// metadata that a service endpoint was programmed to produce. The value must
	// consist of no more than 1024 visible US-ASCII characters as specified in
	// Section 3.3.6. Field Value Components (https://tools.ietf.org/html/rfc7230#section-3.2.6)
	// of the Hypertext Transfer Protocol (HTTP/1.1). If the customer wants the
	// custom attribute returned, the model must set the custom attribute to be
	// included on the way back.
	//
	// This feature is currently supported in the AWS SDKs but not in the Amazon
	// SageMaker Python SDK.
	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string" sensitive:"true"`

	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string `location:"header" locationName:"x-Amzn-Invoked-Production-Variant" type:"string"`
}

// String returns the string representation
func (s InvokeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InvokeEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CustomAttributes != nil {
		v := *s.CustomAttributes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-SageMaker-Custom-Attributes", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InvokedProductionVariant != nil {
		v := *s.InvokedProductionVariant

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-Amzn-Invoked-Production-Variant", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Body", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opInvokeEndpoint = "InvokeEndpoint"

// InvokeEndpointRequest returns a request value for making API operation for
// Amazon SageMaker Runtime.
//
// After you deploy a model into production using Amazon SageMaker hosting services,
// your client applications use this API to get inferences from the model hosted
// at the specified endpoint.
//
// For an overview of Amazon SageMaker, see How It Works (https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html).
//
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpoint are authenticated by using AWS Signature Version
// 4. For information, see Authenticating Requests (AWS Signature Version 4)
// (http://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon S3 API Reference.
//
// A customer's model containers must respond to requests within 60 seconds.
// The model itself can have a maximum processing time of 60 seconds before
// responding to the /invocations. If your model is going to take 50-60 seconds
// of processing time, the SDK socket timeout should be set to be 70 seconds.
//
// Endpoints are scoped to an individual account, and are not public. The URL
// does not contain the account ID, but Amazon SageMaker determines the account
// ID from the authentication token that is supplied by the caller.
//
//    // Example sending a request using InvokeEndpointRequest.
//    req := client.InvokeEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpoint
func (c *Client) InvokeEndpointRequest(input *InvokeEndpointInput) InvokeEndpointRequest {
	op := &aws.Operation{
		Name:       opInvokeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/endpoints/{EndpointName}/invocations",
	}

	if input == nil {
		input = &InvokeEndpointInput{}
	}

	req := c.newRequest(op, input, &InvokeEndpointOutput{})
	return InvokeEndpointRequest{Request: req, Input: input, Copy: c.InvokeEndpointRequest}
}

// InvokeEndpointRequest is the request type for the
// InvokeEndpoint API operation.
type InvokeEndpointRequest struct {
	*aws.Request
	Input *InvokeEndpointInput
	Copy  func(*InvokeEndpointInput) InvokeEndpointRequest
}

// Send marshals and sends the InvokeEndpoint API request.
func (r InvokeEndpointRequest) Send(ctx context.Context) (*InvokeEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InvokeEndpointResponse{
		InvokeEndpointOutput: r.Request.Data.(*InvokeEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InvokeEndpointResponse is the response type for the
// InvokeEndpoint API operation.
type InvokeEndpointResponse struct {
	*InvokeEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InvokeEndpoint request.
func (r *InvokeEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
