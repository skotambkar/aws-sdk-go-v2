// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build integration

package applicationautoscaling_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
)

var _ aws.Config
var _ awserr.Error

func TestInteg_00_DescribeScalableTargets(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg := integration.ConfigWithDefaultRegion("us-west-2")
	svc := applicationautoscaling.New(cfg)
	params := &applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: applicationautoscaling.ServiceNamespaceEc2,
	}

	req := svc.DescribeScalableTargetsRequest(params)
	req.Handlers.Validate.Remove(defaults.ValidateParametersHandler)
	_, err := req.Send(ctx)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
