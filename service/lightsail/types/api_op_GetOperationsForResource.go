// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetOperationsForResourceInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your get operations
	// for resource request.
	PageToken *string `locationName:"pageToken" type:"string"`

	// The name of the resource for which you are requesting information.
	//
	// ResourceName is a required field
	ResourceName *string `locationName:"resourceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetOperationsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOperationsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOperationsForResourceInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetOperationsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// (Deprecated) Returns the number of pages of results that remain.
	//
	// In releases prior to June 12, 2017, this parameter returned null by the API.
	// It is now deprecated, and the API returns the next page token parameter instead.
	NextPageCount *string `locationName:"nextPageCount" deprecated:"true" type:"string"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// An array of key-value pairs containing information about the results of your
	// get operations for resource request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s GetOperationsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}
