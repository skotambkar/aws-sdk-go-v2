// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to add a MethodResponse to an existing Method resource.
type PutMethodResponseInput struct {
	_ struct{} `type:"structure"`

	// [Required] The HTTP verb of the Method resource.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// [Required] The Resource identifier for the Method resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// Specifies the Model resources used for the response's content type. Response
	// models are represented as a key/value map, with a content type as the key
	// and a Model name as the value.
	ResponseModels map[string]string `locationName:"responseModels" type:"map"`

	// A key-value map specifying required or optional response parameters that
	// API Gateway can send back to the caller. A key defines a method response
	// header name and the associated value is a Boolean flag indicating whether
	// the method response parameter is required or not. The method response header
	// names must match the pattern of method.response.header.{name}, where name
	// is a valid and unique header name. The response parameter names defined here
	// are available in the integration response to be mapped from an integration
	// response header expressed in integration.response.header.{name}, a static
	// value enclosed within a pair of single quotes (e.g., 'application/json'),
	// or a JSON expression from the back-end response payload in the form of integration.response.body.{JSON-expression},
	// where JSON-expression is a valid JSON expression without the $ prefix.)
	ResponseParameters map[string]bool `locationName:"responseParameters" type:"map"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The method response's status code.
	//
	// StatusCode is a required field
	StatusCode *string `location:"uri" locationName:"status_code" type:"string" required:"true"`
}

// String returns the string representation
func (s PutMethodResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutMethodResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutMethodResponseInput"}

	if s.HttpMethod == nil {
		invalidParams.Add(aws.NewErrParamRequired("HttpMethod"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if s.StatusCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatusCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a method response of a given HTTP status code returned to the
// client. The method response is passed from the back end through the associated
// integration response that can be transformed using a mapping template.
//
// Example: A MethodResponse instance of an API
//
// Request
//
// The example request retrieves a MethodResponse of the 200 status code.
//  GET /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200
//  HTTP/1.1 Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
//  X-Amz-Date: 20160603T222952Z Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160603/us-east-1/apigateway/aws4_request,
//  SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
// Response
//
// The successful response returns 200 OK status and a payload as follows:
//  { "_links": { "curies": { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
//  "name": "methodresponse", "templated": true }, "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
//  "title": "200" }, "methodresponse:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
//  }, "methodresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
//  } }, "responseModels": { "application/json": "Empty" }, "responseParameters":
//  { "method.response.header.Content-Type": false }, "statusCode": "200" }
//
// Method, IntegrationResponse, Integration Creating an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type PutMethodResponseOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the Model resources used for the response's content-type. Response
	// models are represented as a key/value map, with a content-type as the key
	// and a Model name as the value.
	ResponseModels map[string]string `locationName:"responseModels" type:"map"`

	// A key-value map specifying required or optional response parameters that
	// API Gateway can send back to the caller. A key defines a method response
	// header and the value specifies whether the associated method response header
	// is required or not. The expression of the key must match the pattern method.response.header.{name},
	// where name is a valid and unique header name. API Gateway passes certain
	// integration response data to the method response headers specified here according
	// to the mapping you prescribe in the API's IntegrationResponse. The integration
	// response data that can be mapped include an integration response header expressed
	// in integration.response.header.{name}, a static value enclosed within a pair
	// of single quotes (e.g., 'application/json'), or a JSON expression from the
	// back-end response payload in the form of integration.response.body.{JSON-expression},
	// where JSON-expression is a valid JSON expression without the $ prefix.)
	ResponseParameters map[string]bool `locationName:"responseParameters" type:"map"`

	// The method response's status code.
	StatusCode *string `locationName:"statusCode" type:"string"`
}

// String returns the string representation
func (s PutMethodResponseOutput) String() string {
	return awsutil.Prettify(s)
}