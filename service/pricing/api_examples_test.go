// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/enums"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To retrieve a list of services and service codes
//

func ExampleClient_DescribeServicesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := pricing.New(cfg)
	input := &types.DescribeServicesInput{
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(1),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	req := svc.DescribeServicesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case pricing.ErrCodeInternalErrorException:
				fmt.Println(pricing.ErrCodeInternalErrorException, aerr.Error())
			case pricing.ErrCodeInvalidParameterException:
				fmt.Println(pricing.ErrCodeInvalidParameterException, aerr.Error())
			case pricing.ErrCodeNotFoundException:
				fmt.Println(pricing.ErrCodeNotFoundException, aerr.Error())
			case pricing.ErrCodeInvalidNextTokenException:
				fmt.Println(pricing.ErrCodeInvalidNextTokenException, aerr.Error())
			case pricing.ErrCodeExpiredNextTokenException:
				fmt.Println(pricing.ErrCodeExpiredNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve a list of attribute values
//
// This operation returns a list of values available for the given attribute.
func ExampleClient_GetAttributeValuesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := pricing.New(cfg)
	input := &types.GetAttributeValuesInput{
		AttributeName: aws.String("volumeType"),
		MaxResults:    aws.Int64(2),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	req := svc.GetAttributeValuesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case pricing.ErrCodeInternalErrorException:
				fmt.Println(pricing.ErrCodeInternalErrorException, aerr.Error())
			case pricing.ErrCodeInvalidParameterException:
				fmt.Println(pricing.ErrCodeInvalidParameterException, aerr.Error())
			case pricing.ErrCodeNotFoundException:
				fmt.Println(pricing.ErrCodeNotFoundException, aerr.Error())
			case pricing.ErrCodeInvalidNextTokenException:
				fmt.Println(pricing.ErrCodeInvalidNextTokenException, aerr.Error())
			case pricing.ErrCodeExpiredNextTokenException:
				fmt.Println(pricing.ErrCodeExpiredNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve a list of products
//
// This operation returns a list of products that match the given criteria.
func ExampleClient_GetProductsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := pricing.New(cfg)
	input := &types.GetProductsInput{
		Filters: []types.Filter{
			{
				Field: aws.String("ServiceCode"),
				Type:  enums.FilterTypeTermMatch,
				Value: aws.String("AmazonEC2"),
			},
			{
				Field: aws.String("volumeType"),
				Type:  enums.FilterTypeTermMatch,
				Value: aws.String("Provisioned IOPS"),
			},
		},
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(1),
	}

	req := svc.GetProductsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case pricing.ErrCodeInternalErrorException:
				fmt.Println(pricing.ErrCodeInternalErrorException, aerr.Error())
			case pricing.ErrCodeInvalidParameterException:
				fmt.Println(pricing.ErrCodeInvalidParameterException, aerr.Error())
			case pricing.ErrCodeNotFoundException:
				fmt.Println(pricing.ErrCodeNotFoundException, aerr.Error())
			case pricing.ErrCodeInvalidNextTokenException:
				fmt.Println(pricing.ErrCodeInvalidNextTokenException, aerr.Error())
			case pricing.ErrCodeExpiredNextTokenException:
				fmt.Println(pricing.ErrCodeExpiredNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
