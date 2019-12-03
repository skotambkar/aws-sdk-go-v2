// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request a new export of a RestApi for a particular Stage.
type GetExportInput struct {
	_ struct{} `type:"structure"`

	// The content-type of the export, for example application/json. Currently application/json
	// and application/yaml are supported for exportType ofoas30 and swagger. This
	// should be specified in the Accept header for direct API requests.
	Accepts *string `location:"header" locationName:"Accept" type:"string"`

	// [Required] The type of export. Acceptable values are 'oas30' for OpenAPI
	// 3.0.x and 'swagger' for Swagger/OpenAPI 2.0.
	//
	// ExportType is a required field
	ExportType *string `location:"uri" locationName:"export_type" type:"string" required:"true"`

	// A key-value map of query string parameters that specify properties of the
	// export, depending on the requested exportType. For exportType oas30 and swagger,
	// any combination of the following parameters are supported: extensions='integrations'
	// or extensions='apigateway' will export the API with x-amazon-apigateway-integration
	// extensions. extensions='authorizers' will export the API with x-amazon-apigateway-authorizer
	// extensions. postman will export the API with Postman extensions, allowing
	// for import to the Postman tool
	Parameters map[string]string `location:"querystring" locationName:"parameters" type:"map"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The name of the Stage that will be exported.
	//
	// StageName is a required field
	StageName *string `location:"uri" locationName:"stage_name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetExportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetExportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetExportInput"}

	if s.ExportType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExportType"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The binary blob response to GetExport, which contains the generated SDK.
type GetExportOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The binary blob response to GetExport, which contains the export.
	Body []byte `locationName:"body" type:"blob"`

	// The content-disposition header value in the HTTP response.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// The content-type header value in the HTTP response. This will correspond
	// to a valid 'accept' type in the request.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`
}

// String returns the string representation
func (s GetExportOutput) String() string {
	return awsutil.Prettify(s)
}
