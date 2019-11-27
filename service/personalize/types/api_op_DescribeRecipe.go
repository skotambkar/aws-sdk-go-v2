// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRecipeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the recipe to describe.
	//
	// RecipeArn is a required field
	RecipeArn *string `locationName:"recipeArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRecipeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRecipeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRecipeInput"}

	if s.RecipeArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RecipeArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeRecipeOutput struct {
	_ struct{} `type:"structure"`

	// An object that describes the recipe.
	Recipe *Recipe `locationName:"recipe" type:"structure"`
}

// String returns the string representation
func (s DescribeRecipeOutput) String() string {
	return awsutil.Prettify(s)
}
