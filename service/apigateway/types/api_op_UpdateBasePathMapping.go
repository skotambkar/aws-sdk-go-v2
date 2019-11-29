// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to change information about the BasePathMapping resource.
type UpdateBasePathMappingInput struct {
	_ struct{} `type:"structure"`

	// [Required] The base path of the BasePathMapping resource to change.
	//
	// To specify an empty base path, set this parameter to '(none)'.
	//
	// BasePath is a required field
	BasePath *string `location:"uri" locationName:"base_path" type:"string" required:"true"`

	// [Required] The domain name of the BasePathMapping resource to change.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domain_name" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`
}

// String returns the string representation
func (s UpdateBasePathMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBasePathMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBasePathMappingInput"}

	if s.BasePath == nil {
		invalidParams.Add(aws.NewErrParamRequired("BasePath"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the base path that callers of the API must provide as part of
// the URL after the domain name.
//
// A custom domain name plus a BasePathMapping specification identifies a deployed
// RestApi in a given stage of the owner Account.
//
// Use Custom Domain Names (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type UpdateBasePathMappingOutput struct {
	_ struct{} `type:"structure"`

	// The base path name that callers of the API must provide as part of the URL
	// after the domain name.
	BasePath *string `locationName:"basePath" type:"string"`

	// The string identifier of the associated RestApi.
	RestApiId *string `locationName:"restApiId" type:"string"`

	// The name of the associated stage.
	Stage *string `locationName:"stage" type:"string"`
}

// String returns the string representation
func (s UpdateBasePathMappingOutput) String() string {
	return awsutil.Prettify(s)
}