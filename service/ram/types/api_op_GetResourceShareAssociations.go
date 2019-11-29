// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ram/enums"
)

type GetResourceShareAssociationsInput struct {
	_ struct{} `type:"structure"`

	// The association status.
	AssociationStatus enums.ResourceShareAssociationStatus `locationName:"associationStatus" type:"string" enum:"true"`

	// The association type.
	//
	// AssociationType is a required field
	AssociationType enums.ResourceShareAssociationType `locationName:"associationType" type:"string" required:"true" enum:"true"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The principal. You cannot specify this parameter if the association type
	// is RESOURCE.
	Principal *string `locationName:"principal" type:"string"`

	// The Amazon Resource Name (ARN) of the resource. You cannot specify this parameter
	// if the association type is PRINCIPAL.
	ResourceArn *string `locationName:"resourceArn" type:"string"`

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string `locationName:"resourceShareArns" type:"list"`
}

// String returns the string representation
func (s GetResourceShareAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceShareAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourceShareAssociationsInput"}
	if len(s.AssociationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AssociationType"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetResourceShareAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the associations.
	ResourceShareAssociations []ResourceShareAssociation `locationName:"resourceShareAssociations" type:"list"`
}

// String returns the string representation
func (s GetResourceShareAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}