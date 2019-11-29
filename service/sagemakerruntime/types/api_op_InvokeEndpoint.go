// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
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