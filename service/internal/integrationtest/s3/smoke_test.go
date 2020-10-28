package s3

import (
	"context"
	"testing"
	"time"

	"github.com/awslabs/smithy-go/ptr"

	"github.com/aws/aws-sdk-go-v2/service/internal/integrationtest"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func TestInteg_00_ListBuckets(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	params := &s3.ListBucketsInput{}
	op, err := client.ListBuckets(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	t.Fatalf("%v", *op.Buckets[0].Name)
}

func TestInteg__01__PutBucketAccelerateConfiguration(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	params := &s3.PutBucketAccelerateConfigurationInput{
		Bucket: ptr.String("mockbucket-01"),
		AccelerateConfiguration: &types.AccelerateConfiguration{
			Status: types.BucketAccelerateStatusSuspended,
		},
	}
	op, err := client.PutBucketAccelerateConfiguration(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	t.Fatalf("%v", op)
}
