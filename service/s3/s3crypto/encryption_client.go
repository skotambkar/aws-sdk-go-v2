package s3crypto

import (
	"context"
	"encoding/hex"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// DefaultMinFileSize is used to check whether we want to write to a temp file
// or store the data in memory.
const DefaultMinFileSize = 1024 * 512 * 5

// EncryptionClient is an S3 crypto client. By default the SDK will use Authentication mode which
// will use KMS for key wrapping and AES GCM for content encryption.
// AES GCM will load all data into memory. However, the rest of the content algorithms
// do not load the entire contents into memory.
type EncryptionClient struct {
	S3Client             s3iface.ClientAPI
	ContentCipherBuilder ContentCipherBuilder
	// SaveStrategy will dictate where the envelope is saved.
	//
	// Defaults to the object's metadata
	SaveStrategy SaveStrategy
	// TempFolderPath is used to store temp files when calling PutObject.
	// Temporary files are needed to compute the X-Amz-Content-Sha256 header.
	TempFolderPath string
	// MinFileSize is the minimum size for the content to write to a
	// temporary file instead of using memory.
	MinFileSize int64
}

// NewEncryptionClient instantiates a new S3 crypto client
//
// Example:
//	cmkID := "arn:aws:kms:region:000000000000:key/00000000-0000-0000-0000-000000000000"
//  cfg, err := external.LoadDefaultAWSConfig()
//	handler := s3crypto.NewKMSKeyGenerator(kms.New(cfg), cmkID)
//	svc := s3crypto.New(cfg, s3crypto.AESGCMContentCipherBuilder(handler))
func NewEncryptionClient(cfg aws.Config, builder ContentCipherBuilder, options ...func(*EncryptionClient)) *EncryptionClient {
	client := &EncryptionClient{
		S3Client:             s3.New(cfg),
		ContentCipherBuilder: builder,
		SaveStrategy:         HeaderV2SaveStrategy{},
		MinFileSize:          DefaultMinFileSize,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// PutObjectRequest creates a temp file to encrypt the contents into. It then streams
// that data to S3.
//
// Example:
//  cfg, err := external.LoadDefaultAWSConfig()
//	svc := s3crypto.New(cfg, s3crypto.AESGCMContentCipherBuilder(handler))
//	req := svc.PutObjectRequest(&s3.PutObjectInput {
//	  Key: aws.String("testKey"),
//	  Bucket: aws.String("testBucket"),
//	  Body: strings.NewReader("test data"),
//	})
//	err := req.Send(context.Background())
func (c *EncryptionClient) PutObjectRequest(input *types.PutObjectInput) s3.PutObjectRequest {
	req := c.S3Client.PutObjectRequest(input)

	// Get Size of file
	n, err := aws.SeekerLen(input.Body)
	if err != nil {
		req.Error = err
		return req
	}

	dst, err := getWriterStore(req.Request, c.TempFolderPath, n >= c.MinFileSize)
	if err != nil {
		req.Error = err
		return req
	}

	encryptor, err := c.ContentCipherBuilder.ContentCipher()
	req.Handlers.Build.PushFront(func(r *request.Request) {
		if err != nil {
			r.Error = err
			return
		}

		md5 := newMD5Reader(input.Body)
		sha := newSHA256Writer(dst)
		reader, err := encryptor.EncryptContents(md5)
		if err != nil {
			r.Error = err
			return
		}

		_, err = io.Copy(sha, reader)
		if err != nil {
			r.Error = err
			return
		}

		data := encryptor.GetCipherData()
		env, err := encodeMeta(md5, data)
		if err != nil {
			r.Error = err
			return
		}

		shaHex := hex.EncodeToString(sha.GetValue())
		req.HTTPRequest.Header.Set("X-Amz-Content-Sha256", shaHex)

		dst.Seek(0, io.SeekStart)
		input.Body = dst

		err = c.SaveStrategy.Save(env, r)
		r.Error = err
	})

	return req
}

// PutObject is a wrapper for PutObjectRequest
func (c *EncryptionClient) PutObject(input *types.PutObjectInput) (*types.PutObjectOutput, error) {
	req := c.PutObjectRequest(input)
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	return resp.PutObjectOutput, nil
}

// PutObjectWithContext is a wrapper for PutObjectRequest with the additional
// context, and request options support.
//
// PutObjectWithContext is the same as PutObject with the additional support for
// Context input parameters. The Context must not be nil. A nil Context will
// cause a panic. Use the Context to add deadlining, timeouts, ect. In the future
// this may create sub-contexts for individual underlying requests.
func (c *EncryptionClient) PutObjectWithContext(ctx context.Context, input *types.PutObjectInput, opts ...request.Option) (*types.PutObjectOutput, error) {
	req := c.PutObjectRequest(input)
	req.ApplyOptions(opts...)
	resp, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}
	return resp.PutObjectOutput, nil
}
