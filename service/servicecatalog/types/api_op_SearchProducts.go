// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/enums"
)

type SearchProductsInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The search filters. If no search filters are specified, the output includes
	// all products to which the caller has access.
	Filters map[string][]string `type:"map"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The sort field. If no value is specified, the results are not sorted.
	SortBy enums.ProductViewSortBy `type:"string" enum:"true"`

	// The sort order. If no value is specified, the results are not sorted.
	SortOrder enums.SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s SearchProductsInput) String() string {
	return awsutil.Prettify(s)
}

type SearchProductsOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// The product view aggregations.
	ProductViewAggregations map[string][]ProductViewAggregationValue `type:"map"`

	// Information about the product views.
	ProductViewSummaries []ProductViewSummary `type:"list"`
}

// String returns the string representation
func (s SearchProductsOutput) String() string {
	return awsutil.Prettify(s)
}
