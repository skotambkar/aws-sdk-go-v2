package route53

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/internal/integrationtest"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"

	"github.com/awslabs/smithy-go"
)

func TestInteg_00_ListHostedZones(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := route53.NewFromConfig(cfg)
	params := &route53.ListHostedZonesInput{}
	_, err = client.ListHostedZones(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_GetHostedZone(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := route53.NewFromConfig(cfg)
	params := &route53.GetHostedZoneInput{
		Id: aws.String("fake-zone"),
	}
	_, err = client.GetHostedZone(ctx, params)
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

func TestInteg_02_ChangeResourceRecordSets(t *testing.T) {
	cases := map[string]struct {
		action types.ChangeAction
	}{
		"Create": {
			action: types.ChangeActionCreate,
		},
		"Upsert": {
			action: types.ChangeActionUpsert,
		},
		"Delete": {
			action: types.ChangeActionDelete,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelFn()

			cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
			if err != nil {
				t.Fatalf("failed to load config, %v", err)
			}

			client := route53.NewFromConfig(cfg)

			zoneID := "Z0985815JOOYDBSVBJT6"
			params := &route53.ChangeResourceRecordSetsInput{
				ChangeBatch: &types.ChangeBatch{
					Changes: []*types.Change{
						{
							Action: c.action,
							ResourceRecordSet: &types.ResourceRecordSet{
								Name: aws.String("mockname"),
								ResourceRecords: []*types.ResourceRecord{
									{
										Value: aws.String("127.0.0.1"),
									},
								},
								TTL:  aws.Int64(60),
								Type: types.RRTypeA,
							},
						},
					},
					Comment: aws.String("Register instance"),
				},
				HostedZoneId: aws.String(zoneID),
			}

			resp, err := client.ChangeResourceRecordSets(ctx, params)
			if err != nil {
				t.Errorf("expect no error, got %v", err)
			}

			t.Fatalf("%v", resp)
		})
	}
}
