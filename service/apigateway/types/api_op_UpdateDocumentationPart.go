// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Updates an existing documentation part of a given API.
type UpdateDocumentationPartInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the to-be-updated documentation part.
	//
	// DocumentationPartId is a required field
	DocumentationPartId *string `location:"uri" locationName:"part_id" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDocumentationPartInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDocumentationPartInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDocumentationPartInput"}

	if s.DocumentationPartId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentationPartId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A documentation part for a targeted API entity.
//
// A documentation part consists of a content map (properties) and a target
// (location). The target specifies an API entity to which the documentation
// content applies. The supported API entity types are API, AUTHORIZER, MODEL,
// RESOURCE, METHOD, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY,
// RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. Valid location fields depend
// on the API entity type. All valid fields are not required.
//
// The content map is a JSON string of API-specific key-value pairs. Although
// an API can use any shape for the content map, only the OpenAPI-compliant
// documentation fields will be injected into the associated API entity definition
// in the exported OpenAPI definition file.
//
// Documenting an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationParts
type UpdateDocumentationPartOutput struct {
	_ struct{} `type:"structure"`

	// The DocumentationPart identifier, generated by API Gateway when the DocumentationPart
	// is created.
	Id *string `locationName:"id" type:"string"`

	// The location of the API entity to which the documentation applies. Valid
	// fields depend on the targeted API entity type. All the valid location fields
	// are not required. If not explicitly specified, a valid location field is
	// treated as a wildcard and associated documentation content may be inherited
	// by matching entities, unless overridden.
	Location *DocumentationPartLocation `locationName:"location" type:"structure"`

	// A content map of API-specific key-value pairs describing the targeted API
	// entity. The map must be encoded as a JSON string, e.g., "{ \"description\":
	// \"The API does ...\" }". Only OpenAPI-compliant documentation-related fields
	// from the properties map are exported and, hence, published as part of the
	// API entity definitions, while the original documentation parts are exported
	// in a OpenAPI extension of x-amazon-apigateway-documentation.
	Properties *string `locationName:"properties" type:"string"`
}

// String returns the string representation
func (s UpdateDocumentationPartOutput) String() string {
	return awsutil.Prettify(s)
}
