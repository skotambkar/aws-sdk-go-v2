// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to update an existing model in an existing RestApi resource.
type UpdateModelInput struct {
	_ struct{} `type:"structure"`

	// [Required] The name of the model to update.
	//
	// ModelName is a required field
	ModelName *string `location:"uri" locationName:"model_name" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateModelInput"}

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
type UpdateModelOutput struct {
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
func (s UpdateModelOutput) String() string {
	return awsutil.Prettify(s)
}