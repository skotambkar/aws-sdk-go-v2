// +build integration

package ec2

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/service/internal/integrationtest"
)

func TestInteg_00_DescribeRegions(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := ec2.NewFromConfig(cfg)
	params := &ec2.DescribeRegionsInput{}
	_, err = client.DescribeRegions(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_DescribeInstances(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := ec2.NewFromConfig(cfg)
	params := &ec2.DescribeInstancesInput{
		InstanceIds: []string{
			"i-12345678",
		},
	}
	_, err = client.DescribeInstances(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}

	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expect error to be API error, was not, %v", err)
	}
	if len(apiErr.ErrorCode()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if len(apiErr.ErrorMessage()) == 0 {
		t.Errorf("expect non-empty error message")
	}
}

func TestInteg_01_DescribeVolumes(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	az := "us-west-2d"
	log.Println("Create a new volume in", az)
	volumeOpts := &ec2.CreateVolumeInput{
		AvailabilityZone: &az,
		VolumeType:       types.VolumeTypeStandard,
		Size:             1,
		TagSpecifications: []types.TagSpecification{
			{
				Tags: []types.Tag{
					{Key: aws.String("Name"), Value: aws.String("TO REMOVE")},

					{Key: aws.String("Owner"), Value: aws.String("shaftoe")},
				},
				ResourceType: types.ResourceTypeVolume,
			},
		},
	}

	client := ec2.NewFromConfig(cfg)

	volume, err := client.CreateVolume(ctx, volumeOpts)
	if err != nil {
		t.Fatalf("Failure creating new volume: %v", err)
	}
	log.Println("Volume", *volume.VolumeId, "creation request sent")

	timeout := time.Minute * 5
	log.Println("Waiting for creation process to complete in less than", timeout)
	err = ec2.NewVolumeAvailableWaiter(client).Wait(
		context.TODO(),
		&ec2.DescribeVolumesInput{
			VolumeIds: []string{*volume.VolumeId},
		},
		timeout,
	)
	if err != nil {
		t.Fatalf("Waiter error: %v", err)
	}

	log.Println("Volume creation completed")

	_, err = client.DeleteVolume(context.TODO(), &ec2.DeleteVolumeInput{
		VolumeId: volume.VolumeId,
		DryRun:   false,
	})
	if err != nil {
		t.Fatalf("error deleting volumn with id %v: err %v", *(volume.VolumeId), err)
	}
	log.Println("Volume deletion completed")
}
