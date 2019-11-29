// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to list information about a model in an existing RestApi resource.
type GetModelInput struct {
	_ struct{} `type:"structure"`

	// A query parameter of a Boolean value to resolve (true) all external model
	// references and returns a flattened model schema or not (false) The default
	// is false.
	Flatten *bool `location:"querystring" locationName:"flatten" type:"boolean"`

	// [Required] The name of the model as an identifier.
	//
	// ModelName is a required field
	ModelName *string `location:"uri" locationName:"model_name" type:"string" required:"true"`

	// [Required] The RestApi identifier under which the Model exists.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetModelInput"}

	if s.ModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelName"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the data structure of a method's request or response payload.
//
// A request model defines the data structure of the client-supplied request
// payload. A response model defines the data structure of the response payload
// returned by the back end. Although not required, models are useful for mapping
// payloads between the front end and back end.
//
// A model is used for generating an API's SDK, validating the input request
// body, and creating a skeletal mapping template.
//
// Method, MethodResponse, Models and Mappings (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html)
type GetModelOutput struct {
	_ struct{} `type:"structure"`

	// The content-type for the model.
	ContentType *string `locationName:"contentType" type:"string"`

	// The description of the model.
	Description *string `locationName:"description" type:"string"`

	// The identifier for the model resource.
	Id *string `locationName:"id" type:"string"`

	// The name of the model. Must be an alphanumeric string.
	Name *string `locationName:"name" type:"string"`

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model.
	// Do not include "\*/" characters in the description of any properties because
	// such "\*/" characters may be interpreted as the closing marker for comments
	// in some languages, such as Java or JavaScript, causing the installation of
	// your API's SDK generated by API Gateway to fail.
	Schema *string `locationName:"schema" type:"string"`
}

// String returns the string representation
func (s GetModelOutput) String() string {
	return awsutil.Prettify(s)
}