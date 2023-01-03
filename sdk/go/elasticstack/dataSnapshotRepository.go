// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the information about the registered snaphosts repositories
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/Tonkean/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myUrlRepoSnapshotRepository, err := elasticstack.NewSnapshotRepository(ctx, "myUrlRepoSnapshotRepository", &elasticstack.SnapshotRepositoryArgs{
//				Url: &elasticstack.SnapshotRepositoryUrlArgs{
//					Url: pulumi.String("https://example.com/repo"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			myFsRepoSnapshotRepository, err := elasticstack.NewSnapshotRepository(ctx, "myFsRepoSnapshotRepository", &elasticstack.SnapshotRepositoryArgs{
//				Fs: &elasticstack.SnapshotRepositoryFsArgs{
//					Location:              pulumi.String("/tmp"),
//					Compress:              pulumi.Bool(true),
//					MaxRestoreBytesPerSec: pulumi.String("10mb"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = elasticstack.DataSnapshotRepositoryOutput(ctx, elasticstack.DataSnapshotRepositoryOutputArgs{
//				Name: myFsRepoSnapshotRepository.Name,
//			}, nil)
//			_ = elasticstack.DataSnapshotRepositoryOutput(ctx, elasticstack.DataSnapshotRepositoryOutputArgs{
//				Name: myUrlRepoSnapshotRepository.Name,
//			}, nil)
//			ctx.Export("repoFsLocation", myFsRepoSnapshotRepository.Fs.ApplyT(func(fs elasticstack.SnapshotRepositoryFs) (*string, error) {
//				return &fs.Location, nil
//			}).(pulumi.StringPtrOutput))
//			ctx.Export("repoUrl", myUrlRepoSnapshotRepository.Url.ApplyT(func(url elasticstack.SnapshotRepositoryUrl) (*string, error) {
//				return &url.Url, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
func DataSnapshotRepository(ctx *pulumi.Context, args *DataSnapshotRepositoryArgs, opts ...pulumi.InvokeOption) (*DataSnapshotRepositoryResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv DataSnapshotRepositoryResult
	err := ctx.Invoke("elasticstack:index/dataSnapshotRepository:DataSnapshotRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking DataSnapshotRepository.
type DataSnapshotRepositoryArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *DataSnapshotRepositoryElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Name of the snapshot repository.
	Name string `pulumi:"name"`
}

// A collection of values returned by DataSnapshotRepository.
type DataSnapshotRepositoryResult struct {
	// Azure Blob storage as a repository. Set only if the type of the fetched repo is `azure`.
	Azures []DataSnapshotRepositoryAzure `pulumi:"azures"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *DataSnapshotRepositoryElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Shared filesystem repository. Set only if the type of the fetched repo is `fs`.
	Fs []DataSnapshotRepositoryF `pulumi:"fs"`
	// Google Cloud Storage service as a repository. Set only if the type of the fetched repo is `gcs`.
	Gcs []DataSnapshotRepositoryGc `pulumi:"gcs"`
	// HDFS File System as a repository. Set only if the type of the fetched repo is `hdfs`.
	Hdfs []DataSnapshotRepositoryHdf `pulumi:"hdfs"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Name of the snapshot repository.
	Name string `pulumi:"name"`
	// AWS S3 as a repository. Set only if the type of the fetched repo is `s3`.
	S3s []DataSnapshotRepositoryS3 `pulumi:"s3s"`
	// Repository type.
	Type string `pulumi:"type"`
	// URL repository. Set only if the type of the fetched repo is `url`.
	Urls []DataSnapshotRepositoryUrl `pulumi:"urls"`
}

func DataSnapshotRepositoryOutput(ctx *pulumi.Context, args DataSnapshotRepositoryOutputArgs, opts ...pulumi.InvokeOption) DataSnapshotRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (DataSnapshotRepositoryResult, error) {
			args := v.(DataSnapshotRepositoryArgs)
			r, err := DataSnapshotRepository(ctx, &args, opts...)
			var s DataSnapshotRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(DataSnapshotRepositoryResultOutput)
}

// A collection of arguments for invoking DataSnapshotRepository.
type DataSnapshotRepositoryOutputArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection DataSnapshotRepositoryElasticsearchConnectionPtrInput `pulumi:"elasticsearchConnection"`
	// Name of the snapshot repository.
	Name pulumi.StringInput `pulumi:"name"`
}

func (DataSnapshotRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSnapshotRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by DataSnapshotRepository.
type DataSnapshotRepositoryResultOutput struct{ *pulumi.OutputState }

func (DataSnapshotRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSnapshotRepositoryResult)(nil)).Elem()
}

func (o DataSnapshotRepositoryResultOutput) ToDataSnapshotRepositoryResultOutput() DataSnapshotRepositoryResultOutput {
	return o
}

func (o DataSnapshotRepositoryResultOutput) ToDataSnapshotRepositoryResultOutputWithContext(ctx context.Context) DataSnapshotRepositoryResultOutput {
	return o
}

// Azure Blob storage as a repository. Set only if the type of the fetched repo is `azure`.
func (o DataSnapshotRepositoryResultOutput) Azures() DataSnapshotRepositoryAzureArrayOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) []DataSnapshotRepositoryAzure { return v.Azures }).(DataSnapshotRepositoryAzureArrayOutput)
}

// Elasticsearch connection configuration block.
func (o DataSnapshotRepositoryResultOutput) ElasticsearchConnection() DataSnapshotRepositoryElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) *DataSnapshotRepositoryElasticsearchConnection {
		return v.ElasticsearchConnection
	}).(DataSnapshotRepositoryElasticsearchConnectionPtrOutput)
}

// Shared filesystem repository. Set only if the type of the fetched repo is `fs`.
func (o DataSnapshotRepositoryResultOutput) Fs() DataSnapshotRepositoryFArrayOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) []DataSnapshotRepositoryF { return v.Fs }).(DataSnapshotRepositoryFArrayOutput)
}

// Google Cloud Storage service as a repository. Set only if the type of the fetched repo is `gcs`.
func (o DataSnapshotRepositoryResultOutput) Gcs() DataSnapshotRepositoryGcArrayOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) []DataSnapshotRepositoryGc { return v.Gcs }).(DataSnapshotRepositoryGcArrayOutput)
}

// HDFS File System as a repository. Set only if the type of the fetched repo is `hdfs`.
func (o DataSnapshotRepositoryResultOutput) Hdfs() DataSnapshotRepositoryHdfArrayOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) []DataSnapshotRepositoryHdf { return v.Hdfs }).(DataSnapshotRepositoryHdfArrayOutput)
}

// Internal identifier of the resource
func (o DataSnapshotRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the snapshot repository.
func (o DataSnapshotRepositoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// AWS S3 as a repository. Set only if the type of the fetched repo is `s3`.
func (o DataSnapshotRepositoryResultOutput) S3s() DataSnapshotRepositoryS3ArrayOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) []DataSnapshotRepositoryS3 { return v.S3s }).(DataSnapshotRepositoryS3ArrayOutput)
}

// Repository type.
func (o DataSnapshotRepositoryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) string { return v.Type }).(pulumi.StringOutput)
}

// URL repository. Set only if the type of the fetched repo is `url`.
func (o DataSnapshotRepositoryResultOutput) Urls() DataSnapshotRepositoryUrlArrayOutput {
	return o.ApplyT(func(v DataSnapshotRepositoryResult) []DataSnapshotRepositoryUrl { return v.Urls }).(DataSnapshotRepositoryUrlArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSnapshotRepositoryResultOutput{})
}
