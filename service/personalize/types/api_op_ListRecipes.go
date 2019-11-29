// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/personalize/enums"
)

type ListRecipesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of recipes to return.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token returned from the previous call to ListRecipes for getting the next
	// set of recipes (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`

	// The default is SERVICE.
	RecipeProvider enums.RecipeProvider `locationName:"recipeProvider" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListRecipesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRecipesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRecipesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRecipesOutput struct {
	_ struct{} `type:"structure"`

	// A token for getting the next set of recipes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of available recipes.
	Recipes []RecipeSummary `locationName:"recipes" type:"list"`
}

// String returns the string representation
func (s ListRecipesOutput) String() string {
	return awsutil.Prettify(s)
}