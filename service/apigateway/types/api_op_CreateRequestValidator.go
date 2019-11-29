// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Creates a RequestValidator of a given RestApi.
type CreateRequestValidatorInput struct {
	_ struct{} `type:"structure"`

	// The name of the to-be-created RequestValidator.
	Name *string `locationName:"name" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// A Boolean flag to indicate whether to validate request body according to
	// the configured model schema for the method (true) or not (false).
	ValidateRequestBody *bool `locationName:"validateRequestBody" type:"boolean"`

	// A Boolean flag to indicate whether to validate request parameters, true,
	// or not false.
	ValidateRequestParameters *bool `locationName:"validateRequestParameters" type:"boolean"`
}

// String returns the string representation
func (s CreateRequestValidatorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRequestValidatorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRequestValidatorInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A set of validation rules for incoming Method requests.
//
// In OpenAPI, a RequestValidator of an API is defined by the x-amazon-apigateway-request-validators.requestValidator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validators.requestValidator.html)
// object. It the referenced using the x-amazon-apigateway-request-validator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validator)
// property.
//
// Enable Basic Request Validation in API Gateway (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html)
type CreateRequestValidatorOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of this RequestValidator.
	Id *string `locationName:"id" type:"string"`

	// The name of this RequestValidator
	Name *string `locationName:"name" type:"string"`

	// A Boolean flag to indicate whether to validate a request body according to
	// the configured Model schema.
	ValidateRequestBody *bool `locationName:"validateRequestBody" type:"boolean"`

	// A Boolean flag to indicate whether to validate request parameters (true)
	// or not (false).
	ValidateRequestParameters *bool `locationName:"validateRequestParameters" type:"boolean"`
}

// String returns the string representation
func (s CreateRequestValidatorOutput) String() string {
	return awsutil.Prettify(s)
}