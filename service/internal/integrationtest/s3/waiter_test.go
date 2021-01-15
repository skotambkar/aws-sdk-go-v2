// +build integration

package s3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/internal/integrationtest"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"testing"
	"time"
)

func TestInteg_00_WaituntilExists(t *testing.T) {

	key := integrationtest.UniqueID()

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	objectExistsWaiter := s3.NewObjectExistsWaiter(client)
	objectNotExistsWaiter := s3.NewObjectNotExistsWaiter(client)
	bucketExistsWaiter := s3.NewBucketExistsWaiter(client)
	bucketNotExistsWaiter := s3.NewBucketNotExistsWaiter(client)

	// construct a put object
	putObjectInput := &s3.PutObjectInput{
		Bucket: &setupMetadata.Buckets.Source.Name,
		Key:    &key,
		Body:   bytes.NewReader([]byte("hello world")),
	}

	err = objectExistsWaiter.Wait(ctx, &s3.HeadObjectInput{
		Key:    &key,
		Bucket: &setupMetadata.Buckets.Source.Name,
	}, 5*time.Second)
	if err != nil {
		t.Logf("object exists waiter failed as expected, %v", err)
	}

	_, err = client.PutObject(ctx, putObjectInput)
	if err != nil {
		t.Fatalf("failed to put object %v", err)
	}

	err = objectExistsWaiter.Wait(ctx, &s3.HeadObjectInput{
		Key:    &key,
		Bucket: &setupMetadata.Buckets.Source.Name,
	}, 120*time.Second)
	if err != nil {
		t.Fatalf("object exists waiter failed to verify, %v", err)
	}

	err = objectNotExistsWaiter.Wait(ctx, &s3.HeadObjectInput{
		Key:    &key,
		Bucket: &setupMetadata.Buckets.Source.Name,
	}, 5*time.Second)
	if err != nil {
		t.Logf("Object Not exists waiter failed as expected, %v", err)
	}

	_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &setupMetadata.Buckets.Source.Name,
		Key:    &key,
	})
	if err != nil {
		t.Fatalf("failed to delete object %v", err)
	}

	err = objectNotExistsWaiter.Wait(ctx, &s3.HeadObjectInput{
		Key:    &key,
		Bucket: &setupMetadata.Buckets.Source.Name,
	}, 120*time.Second)
	if err != nil {
		t.Fatalf("object not exists waiter failed to verify, %v", err)
	}

	err = bucketExistsWaiter.Wait(ctx, &s3.HeadBucketInput{
		Bucket: aws.String("newmock"),
	}, 5*time.Second)
	if err != nil {
		t.Logf("Bucket exists waiter failed as expected, %v", err)
	}

	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String("newmock"),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintUsWest2,
		},
	})
	if err != nil {
		t.Fatalf("create bucket failed, %v", err)
	}

	err = bucketExistsWaiter.Wait(ctx, &s3.HeadBucketInput{
		Bucket: aws.String("newmock"),
	}, 20*time.Second)
	if err != nil {
		t.Fatalf("Bucket exists waiter failed, %v", err)
	}

	err = bucketNotExistsWaiter.Wait(ctx, &s3.HeadBucketInput{
		Bucket: aws.String("newmock"),
	}, 5*time.Second)
	if err != nil {
		t.Logf("Bucket not exists waiter failed as expected, %v", err)
	}

	_, err = client.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: aws.String("newmock"),
	})
	if err != nil {
		t.Fatalf("failed to delete bucket, %v", err)
	}

	err = bucketNotExistsWaiter.Wait(ctx, &s3.HeadBucketInput{
		Bucket: aws.String("newmock"),
	}, 20*time.Second)
	if err != nil {
		t.Fatalf("Bucket not exists waiter failed, %v", err)
	}

}
