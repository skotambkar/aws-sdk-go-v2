// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// An object that identifies an item.
//
// The and APIs return a list of PredictedItems.
type PredictedItem struct {
	_ struct{} `type:"structure"`

	// The recommended item ID.
	ItemId *string `locationName:"itemId" type:"string"`
}

// String returns the string representation
func (s PredictedItem) String() string {
	return awsutil.Prettify(s)
}
