// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package s3controliface provides an interface to enable mocking the AWS S3 Control service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package s3controliface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3control"
)

// S3ControlAPI provides an interface to enable mocking the
// s3control.S3Control service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS S3 Control.
//    func myFunc(svc s3controliface.S3ControlAPI) bool {
//        // Make svc.CreateAccessPoint request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := s3control.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockS3ControlClient struct {
//        s3controliface.S3ControlAPI
//    }
//    func (m *mockS3ControlClient) CreateAccessPoint(input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockS3ControlClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type S3ControlAPI interface {
	CreateAccessPoint(*s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error)
	CreateAccessPointWithContext(aws.Context, *s3control.CreateAccessPointInput, ...request.Option) (*s3control.CreateAccessPointOutput, error)
	CreateAccessPointRequest(*s3control.CreateAccessPointInput) (*request.Request, *s3control.CreateAccessPointOutput)

	CreateBucket(*s3control.CreateBucketInput) (*s3control.CreateBucketOutput, error)
	CreateBucketWithContext(aws.Context, *s3control.CreateBucketInput, ...request.Option) (*s3control.CreateBucketOutput, error)
	CreateBucketRequest(*s3control.CreateBucketInput) (*request.Request, *s3control.CreateBucketOutput)

	CreateJob(*s3control.CreateJobInput) (*s3control.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *s3control.CreateJobInput, ...request.Option) (*s3control.CreateJobOutput, error)
	CreateJobRequest(*s3control.CreateJobInput) (*request.Request, *s3control.CreateJobOutput)

	DeleteAccessPoint(*s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error)
	DeleteAccessPointWithContext(aws.Context, *s3control.DeleteAccessPointInput, ...request.Option) (*s3control.DeleteAccessPointOutput, error)
	DeleteAccessPointRequest(*s3control.DeleteAccessPointInput) (*request.Request, *s3control.DeleteAccessPointOutput)

	DeleteAccessPointPolicy(*s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error)
	DeleteAccessPointPolicyWithContext(aws.Context, *s3control.DeleteAccessPointPolicyInput, ...request.Option) (*s3control.DeleteAccessPointPolicyOutput, error)
	DeleteAccessPointPolicyRequest(*s3control.DeleteAccessPointPolicyInput) (*request.Request, *s3control.DeleteAccessPointPolicyOutput)

	DeleteBucket(*s3control.DeleteBucketInput) (*s3control.DeleteBucketOutput, error)
	DeleteBucketWithContext(aws.Context, *s3control.DeleteBucketInput, ...request.Option) (*s3control.DeleteBucketOutput, error)
	DeleteBucketRequest(*s3control.DeleteBucketInput) (*request.Request, *s3control.DeleteBucketOutput)

	DeleteBucketLifecycleConfiguration(*s3control.DeleteBucketLifecycleConfigurationInput) (*s3control.DeleteBucketLifecycleConfigurationOutput, error)
	DeleteBucketLifecycleConfigurationWithContext(aws.Context, *s3control.DeleteBucketLifecycleConfigurationInput, ...request.Option) (*s3control.DeleteBucketLifecycleConfigurationOutput, error)
	DeleteBucketLifecycleConfigurationRequest(*s3control.DeleteBucketLifecycleConfigurationInput) (*request.Request, *s3control.DeleteBucketLifecycleConfigurationOutput)

	DeleteBucketPolicy(*s3control.DeleteBucketPolicyInput) (*s3control.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyWithContext(aws.Context, *s3control.DeleteBucketPolicyInput, ...request.Option) (*s3control.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyRequest(*s3control.DeleteBucketPolicyInput) (*request.Request, *s3control.DeleteBucketPolicyOutput)

	DeleteBucketTagging(*s3control.DeleteBucketTaggingInput) (*s3control.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingWithContext(aws.Context, *s3control.DeleteBucketTaggingInput, ...request.Option) (*s3control.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingRequest(*s3control.DeleteBucketTaggingInput) (*request.Request, *s3control.DeleteBucketTaggingOutput)

	DeleteJobTagging(*s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error)
	DeleteJobTaggingWithContext(aws.Context, *s3control.DeleteJobTaggingInput, ...request.Option) (*s3control.DeleteJobTaggingOutput, error)
	DeleteJobTaggingRequest(*s3control.DeleteJobTaggingInput) (*request.Request, *s3control.DeleteJobTaggingOutput)

	DeletePublicAccessBlock(*s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error)
	DeletePublicAccessBlockWithContext(aws.Context, *s3control.DeletePublicAccessBlockInput, ...request.Option) (*s3control.DeletePublicAccessBlockOutput, error)
	DeletePublicAccessBlockRequest(*s3control.DeletePublicAccessBlockInput) (*request.Request, *s3control.DeletePublicAccessBlockOutput)

	DeleteStorageLensConfiguration(*s3control.DeleteStorageLensConfigurationInput) (*s3control.DeleteStorageLensConfigurationOutput, error)
	DeleteStorageLensConfigurationWithContext(aws.Context, *s3control.DeleteStorageLensConfigurationInput, ...request.Option) (*s3control.DeleteStorageLensConfigurationOutput, error)
	DeleteStorageLensConfigurationRequest(*s3control.DeleteStorageLensConfigurationInput) (*request.Request, *s3control.DeleteStorageLensConfigurationOutput)

	DeleteStorageLensConfigurationTagging(*s3control.DeleteStorageLensConfigurationTaggingInput) (*s3control.DeleteStorageLensConfigurationTaggingOutput, error)
	DeleteStorageLensConfigurationTaggingWithContext(aws.Context, *s3control.DeleteStorageLensConfigurationTaggingInput, ...request.Option) (*s3control.DeleteStorageLensConfigurationTaggingOutput, error)
	DeleteStorageLensConfigurationTaggingRequest(*s3control.DeleteStorageLensConfigurationTaggingInput) (*request.Request, *s3control.DeleteStorageLensConfigurationTaggingOutput)

	DescribeJob(*s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error)
	DescribeJobWithContext(aws.Context, *s3control.DescribeJobInput, ...request.Option) (*s3control.DescribeJobOutput, error)
	DescribeJobRequest(*s3control.DescribeJobInput) (*request.Request, *s3control.DescribeJobOutput)

	GetAccessPoint(*s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error)
	GetAccessPointWithContext(aws.Context, *s3control.GetAccessPointInput, ...request.Option) (*s3control.GetAccessPointOutput, error)
	GetAccessPointRequest(*s3control.GetAccessPointInput) (*request.Request, *s3control.GetAccessPointOutput)

	GetAccessPointPolicy(*s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyWithContext(aws.Context, *s3control.GetAccessPointPolicyInput, ...request.Option) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyRequest(*s3control.GetAccessPointPolicyInput) (*request.Request, *s3control.GetAccessPointPolicyOutput)

	GetAccessPointPolicyStatus(*s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusWithContext(aws.Context, *s3control.GetAccessPointPolicyStatusInput, ...request.Option) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusRequest(*s3control.GetAccessPointPolicyStatusInput) (*request.Request, *s3control.GetAccessPointPolicyStatusOutput)

	GetBucket(*s3control.GetBucketInput) (*s3control.GetBucketOutput, error)
	GetBucketWithContext(aws.Context, *s3control.GetBucketInput, ...request.Option) (*s3control.GetBucketOutput, error)
	GetBucketRequest(*s3control.GetBucketInput) (*request.Request, *s3control.GetBucketOutput)

	GetBucketLifecycleConfiguration(*s3control.GetBucketLifecycleConfigurationInput) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationWithContext(aws.Context, *s3control.GetBucketLifecycleConfigurationInput, ...request.Option) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationRequest(*s3control.GetBucketLifecycleConfigurationInput) (*request.Request, *s3control.GetBucketLifecycleConfigurationOutput)

	GetBucketPolicy(*s3control.GetBucketPolicyInput) (*s3control.GetBucketPolicyOutput, error)
	GetBucketPolicyWithContext(aws.Context, *s3control.GetBucketPolicyInput, ...request.Option) (*s3control.GetBucketPolicyOutput, error)
	GetBucketPolicyRequest(*s3control.GetBucketPolicyInput) (*request.Request, *s3control.GetBucketPolicyOutput)

	GetBucketTagging(*s3control.GetBucketTaggingInput) (*s3control.GetBucketTaggingOutput, error)
	GetBucketTaggingWithContext(aws.Context, *s3control.GetBucketTaggingInput, ...request.Option) (*s3control.GetBucketTaggingOutput, error)
	GetBucketTaggingRequest(*s3control.GetBucketTaggingInput) (*request.Request, *s3control.GetBucketTaggingOutput)

	GetJobTagging(*s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error)
	GetJobTaggingWithContext(aws.Context, *s3control.GetJobTaggingInput, ...request.Option) (*s3control.GetJobTaggingOutput, error)
	GetJobTaggingRequest(*s3control.GetJobTaggingInput) (*request.Request, *s3control.GetJobTaggingOutput)

	GetPublicAccessBlock(*s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error)
	GetPublicAccessBlockWithContext(aws.Context, *s3control.GetPublicAccessBlockInput, ...request.Option) (*s3control.GetPublicAccessBlockOutput, error)
	GetPublicAccessBlockRequest(*s3control.GetPublicAccessBlockInput) (*request.Request, *s3control.GetPublicAccessBlockOutput)

	GetStorageLensConfiguration(*s3control.GetStorageLensConfigurationInput) (*s3control.GetStorageLensConfigurationOutput, error)
	GetStorageLensConfigurationWithContext(aws.Context, *s3control.GetStorageLensConfigurationInput, ...request.Option) (*s3control.GetStorageLensConfigurationOutput, error)
	GetStorageLensConfigurationRequest(*s3control.GetStorageLensConfigurationInput) (*request.Request, *s3control.GetStorageLensConfigurationOutput)

	GetStorageLensConfigurationTagging(*s3control.GetStorageLensConfigurationTaggingInput) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	GetStorageLensConfigurationTaggingWithContext(aws.Context, *s3control.GetStorageLensConfigurationTaggingInput, ...request.Option) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	GetStorageLensConfigurationTaggingRequest(*s3control.GetStorageLensConfigurationTaggingInput) (*request.Request, *s3control.GetStorageLensConfigurationTaggingOutput)

	ListAccessPoints(*s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsWithContext(aws.Context, *s3control.ListAccessPointsInput, ...request.Option) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsRequest(*s3control.ListAccessPointsInput) (*request.Request, *s3control.ListAccessPointsOutput)

	ListAccessPointsPages(*s3control.ListAccessPointsInput, func(*s3control.ListAccessPointsOutput, bool) bool) error
	ListAccessPointsPagesWithContext(aws.Context, *s3control.ListAccessPointsInput, func(*s3control.ListAccessPointsOutput, bool) bool, ...request.Option) error

	ListJobs(*s3control.ListJobsInput) (*s3control.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *s3control.ListJobsInput, ...request.Option) (*s3control.ListJobsOutput, error)
	ListJobsRequest(*s3control.ListJobsInput) (*request.Request, *s3control.ListJobsOutput)

	ListJobsPages(*s3control.ListJobsInput, func(*s3control.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *s3control.ListJobsInput, func(*s3control.ListJobsOutput, bool) bool, ...request.Option) error

	ListRegionalBuckets(*s3control.ListRegionalBucketsInput) (*s3control.ListRegionalBucketsOutput, error)
	ListRegionalBucketsWithContext(aws.Context, *s3control.ListRegionalBucketsInput, ...request.Option) (*s3control.ListRegionalBucketsOutput, error)
	ListRegionalBucketsRequest(*s3control.ListRegionalBucketsInput) (*request.Request, *s3control.ListRegionalBucketsOutput)

	ListRegionalBucketsPages(*s3control.ListRegionalBucketsInput, func(*s3control.ListRegionalBucketsOutput, bool) bool) error
	ListRegionalBucketsPagesWithContext(aws.Context, *s3control.ListRegionalBucketsInput, func(*s3control.ListRegionalBucketsOutput, bool) bool, ...request.Option) error

	ListStorageLensConfigurations(*s3control.ListStorageLensConfigurationsInput) (*s3control.ListStorageLensConfigurationsOutput, error)
	ListStorageLensConfigurationsWithContext(aws.Context, *s3control.ListStorageLensConfigurationsInput, ...request.Option) (*s3control.ListStorageLensConfigurationsOutput, error)
	ListStorageLensConfigurationsRequest(*s3control.ListStorageLensConfigurationsInput) (*request.Request, *s3control.ListStorageLensConfigurationsOutput)

	PutAccessPointPolicy(*s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error)
	PutAccessPointPolicyWithContext(aws.Context, *s3control.PutAccessPointPolicyInput, ...request.Option) (*s3control.PutAccessPointPolicyOutput, error)
	PutAccessPointPolicyRequest(*s3control.PutAccessPointPolicyInput) (*request.Request, *s3control.PutAccessPointPolicyOutput)

	PutBucketLifecycleConfiguration(*s3control.PutBucketLifecycleConfigurationInput) (*s3control.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationWithContext(aws.Context, *s3control.PutBucketLifecycleConfigurationInput, ...request.Option) (*s3control.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationRequest(*s3control.PutBucketLifecycleConfigurationInput) (*request.Request, *s3control.PutBucketLifecycleConfigurationOutput)

	PutBucketPolicy(*s3control.PutBucketPolicyInput) (*s3control.PutBucketPolicyOutput, error)
	PutBucketPolicyWithContext(aws.Context, *s3control.PutBucketPolicyInput, ...request.Option) (*s3control.PutBucketPolicyOutput, error)
	PutBucketPolicyRequest(*s3control.PutBucketPolicyInput) (*request.Request, *s3control.PutBucketPolicyOutput)

	PutBucketTagging(*s3control.PutBucketTaggingInput) (*s3control.PutBucketTaggingOutput, error)
	PutBucketTaggingWithContext(aws.Context, *s3control.PutBucketTaggingInput, ...request.Option) (*s3control.PutBucketTaggingOutput, error)
	PutBucketTaggingRequest(*s3control.PutBucketTaggingInput) (*request.Request, *s3control.PutBucketTaggingOutput)

	PutJobTagging(*s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error)
	PutJobTaggingWithContext(aws.Context, *s3control.PutJobTaggingInput, ...request.Option) (*s3control.PutJobTaggingOutput, error)
	PutJobTaggingRequest(*s3control.PutJobTaggingInput) (*request.Request, *s3control.PutJobTaggingOutput)

	PutPublicAccessBlock(*s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error)
	PutPublicAccessBlockWithContext(aws.Context, *s3control.PutPublicAccessBlockInput, ...request.Option) (*s3control.PutPublicAccessBlockOutput, error)
	PutPublicAccessBlockRequest(*s3control.PutPublicAccessBlockInput) (*request.Request, *s3control.PutPublicAccessBlockOutput)

	PutStorageLensConfiguration(*s3control.PutStorageLensConfigurationInput) (*s3control.PutStorageLensConfigurationOutput, error)
	PutStorageLensConfigurationWithContext(aws.Context, *s3control.PutStorageLensConfigurationInput, ...request.Option) (*s3control.PutStorageLensConfigurationOutput, error)
	PutStorageLensConfigurationRequest(*s3control.PutStorageLensConfigurationInput) (*request.Request, *s3control.PutStorageLensConfigurationOutput)

	PutStorageLensConfigurationTagging(*s3control.PutStorageLensConfigurationTaggingInput) (*s3control.PutStorageLensConfigurationTaggingOutput, error)
	PutStorageLensConfigurationTaggingWithContext(aws.Context, *s3control.PutStorageLensConfigurationTaggingInput, ...request.Option) (*s3control.PutStorageLensConfigurationTaggingOutput, error)
	PutStorageLensConfigurationTaggingRequest(*s3control.PutStorageLensConfigurationTaggingInput) (*request.Request, *s3control.PutStorageLensConfigurationTaggingOutput)

	UpdateJobPriority(*s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error)
	UpdateJobPriorityWithContext(aws.Context, *s3control.UpdateJobPriorityInput, ...request.Option) (*s3control.UpdateJobPriorityOutput, error)
	UpdateJobPriorityRequest(*s3control.UpdateJobPriorityInput) (*request.Request, *s3control.UpdateJobPriorityOutput)

	UpdateJobStatus(*s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error)
	UpdateJobStatusWithContext(aws.Context, *s3control.UpdateJobStatusInput, ...request.Option) (*s3control.UpdateJobStatusOutput, error)
	UpdateJobStatusRequest(*s3control.UpdateJobStatusInput) (*request.Request, *s3control.UpdateJobStatusOutput)
}

var _ S3ControlAPI = (*s3control.S3Control)(nil)
