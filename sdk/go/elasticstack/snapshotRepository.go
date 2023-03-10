// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Registers or updates a snapshot repository. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/put-snapshot-repo-api.html and https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html
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
//			_, err := elasticstack.NewSnapshotRepository(ctx, "myUrlRepo", &elasticstack.SnapshotRepositoryArgs{
//				Url: &elasticstack.SnapshotRepositoryUrlArgs{
//					Url: pulumi.String("https://example.com/repo"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewSnapshotRepository(ctx, "myFsRepo", &elasticstack.SnapshotRepositoryArgs{
//				Fs: &elasticstack.SnapshotRepositoryFsArgs{
//					Compress:              pulumi.Bool(true),
//					Location:              pulumi.String("/tmp"),
//					MaxRestoreBytesPerSec: pulumi.String("10mb"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import elasticstack:index/snapshotRepository:SnapshotRepository my_repository <cluster_uuid>/<repository name>
//
// ```
type SnapshotRepository struct {
	pulumi.CustomResourceState

	// Support for using Azure Blob storage as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-azure.html
	Azure SnapshotRepositoryAzurePtrOutput `pulumi:"azure"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection SnapshotRepositoryElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Shared filesystem repository. Repositories of this type use a shared filesystem to store snapshots. This filesystem must be accessible to all master and data nodes in the cluster.
	Fs SnapshotRepositoryFsPtrOutput `pulumi:"fs"`
	// Support for using the Google Cloud Storage service as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-gcs.html
	Gcs SnapshotRepositoryGcsPtrOutput `pulumi:"gcs"`
	// Support for using HDFS File System as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-hdfs.html
	Hdfs SnapshotRepositoryHdfsPtrOutput `pulumi:"hdfs"`
	// Name of the snapshot repository to register or update.
	Name pulumi.StringOutput `pulumi:"name"`
	// Support for using AWS S3 as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-repository.html
	S3 SnapshotRepositoryS3PtrOutput `pulumi:"s3"`
	// URL repository. Repositories of this type are read-only for the cluster. This means the cluster can retrieve or restore snapshots from the repository but cannot write or create snapshots in it.
	Url SnapshotRepositoryUrlPtrOutput `pulumi:"url"`
	// If true, the request verifies the repository is functional on all master and data nodes in the cluster.
	Verify pulumi.BoolPtrOutput `pulumi:"verify"`
}

// NewSnapshotRepository registers a new resource with the given unique name, arguments, and options.
func NewSnapshotRepository(ctx *pulumi.Context,
	name string, args *SnapshotRepositoryArgs, opts ...pulumi.ResourceOption) (*SnapshotRepository, error) {
	if args == nil {
		args = &SnapshotRepositoryArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SnapshotRepository
	err := ctx.RegisterResource("elasticstack:index/snapshotRepository:SnapshotRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotRepository gets an existing SnapshotRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotRepositoryState, opts ...pulumi.ResourceOption) (*SnapshotRepository, error) {
	var resource SnapshotRepository
	err := ctx.ReadResource("elasticstack:index/snapshotRepository:SnapshotRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotRepository resources.
type snapshotRepositoryState struct {
	// Support for using Azure Blob storage as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-azure.html
	Azure *SnapshotRepositoryAzure `pulumi:"azure"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *SnapshotRepositoryElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Shared filesystem repository. Repositories of this type use a shared filesystem to store snapshots. This filesystem must be accessible to all master and data nodes in the cluster.
	Fs *SnapshotRepositoryFs `pulumi:"fs"`
	// Support for using the Google Cloud Storage service as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-gcs.html
	Gcs *SnapshotRepositoryGcs `pulumi:"gcs"`
	// Support for using HDFS File System as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-hdfs.html
	Hdfs *SnapshotRepositoryHdfs `pulumi:"hdfs"`
	// Name of the snapshot repository to register or update.
	Name *string `pulumi:"name"`
	// Support for using AWS S3 as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-repository.html
	S3 *SnapshotRepositoryS3 `pulumi:"s3"`
	// URL repository. Repositories of this type are read-only for the cluster. This means the cluster can retrieve or restore snapshots from the repository but cannot write or create snapshots in it.
	Url *SnapshotRepositoryUrl `pulumi:"url"`
	// If true, the request verifies the repository is functional on all master and data nodes in the cluster.
	Verify *bool `pulumi:"verify"`
}

type SnapshotRepositoryState struct {
	// Support for using Azure Blob storage as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-azure.html
	Azure SnapshotRepositoryAzurePtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection SnapshotRepositoryElasticsearchConnectionPtrInput
	// Shared filesystem repository. Repositories of this type use a shared filesystem to store snapshots. This filesystem must be accessible to all master and data nodes in the cluster.
	Fs SnapshotRepositoryFsPtrInput
	// Support for using the Google Cloud Storage service as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-gcs.html
	Gcs SnapshotRepositoryGcsPtrInput
	// Support for using HDFS File System as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-hdfs.html
	Hdfs SnapshotRepositoryHdfsPtrInput
	// Name of the snapshot repository to register or update.
	Name pulumi.StringPtrInput
	// Support for using AWS S3 as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-repository.html
	S3 SnapshotRepositoryS3PtrInput
	// URL repository. Repositories of this type are read-only for the cluster. This means the cluster can retrieve or restore snapshots from the repository but cannot write or create snapshots in it.
	Url SnapshotRepositoryUrlPtrInput
	// If true, the request verifies the repository is functional on all master and data nodes in the cluster.
	Verify pulumi.BoolPtrInput
}

func (SnapshotRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotRepositoryState)(nil)).Elem()
}

type snapshotRepositoryArgs struct {
	// Support for using Azure Blob storage as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-azure.html
	Azure *SnapshotRepositoryAzure `pulumi:"azure"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *SnapshotRepositoryElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Shared filesystem repository. Repositories of this type use a shared filesystem to store snapshots. This filesystem must be accessible to all master and data nodes in the cluster.
	Fs *SnapshotRepositoryFs `pulumi:"fs"`
	// Support for using the Google Cloud Storage service as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-gcs.html
	Gcs *SnapshotRepositoryGcs `pulumi:"gcs"`
	// Support for using HDFS File System as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-hdfs.html
	Hdfs *SnapshotRepositoryHdfs `pulumi:"hdfs"`
	// Name of the snapshot repository to register or update.
	Name *string `pulumi:"name"`
	// Support for using AWS S3 as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-repository.html
	S3 *SnapshotRepositoryS3 `pulumi:"s3"`
	// URL repository. Repositories of this type are read-only for the cluster. This means the cluster can retrieve or restore snapshots from the repository but cannot write or create snapshots in it.
	Url *SnapshotRepositoryUrl `pulumi:"url"`
	// If true, the request verifies the repository is functional on all master and data nodes in the cluster.
	Verify *bool `pulumi:"verify"`
}

// The set of arguments for constructing a SnapshotRepository resource.
type SnapshotRepositoryArgs struct {
	// Support for using Azure Blob storage as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-azure.html
	Azure SnapshotRepositoryAzurePtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection SnapshotRepositoryElasticsearchConnectionPtrInput
	// Shared filesystem repository. Repositories of this type use a shared filesystem to store snapshots. This filesystem must be accessible to all master and data nodes in the cluster.
	Fs SnapshotRepositoryFsPtrInput
	// Support for using the Google Cloud Storage service as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-gcs.html
	Gcs SnapshotRepositoryGcsPtrInput
	// Support for using HDFS File System as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-hdfs.html
	Hdfs SnapshotRepositoryHdfsPtrInput
	// Name of the snapshot repository to register or update.
	Name pulumi.StringPtrInput
	// Support for using AWS S3 as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-repository.html
	S3 SnapshotRepositoryS3PtrInput
	// URL repository. Repositories of this type are read-only for the cluster. This means the cluster can retrieve or restore snapshots from the repository but cannot write or create snapshots in it.
	Url SnapshotRepositoryUrlPtrInput
	// If true, the request verifies the repository is functional on all master and data nodes in the cluster.
	Verify pulumi.BoolPtrInput
}

func (SnapshotRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotRepositoryArgs)(nil)).Elem()
}

type SnapshotRepositoryInput interface {
	pulumi.Input

	ToSnapshotRepositoryOutput() SnapshotRepositoryOutput
	ToSnapshotRepositoryOutputWithContext(ctx context.Context) SnapshotRepositoryOutput
}

func (*SnapshotRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotRepository)(nil)).Elem()
}

func (i *SnapshotRepository) ToSnapshotRepositoryOutput() SnapshotRepositoryOutput {
	return i.ToSnapshotRepositoryOutputWithContext(context.Background())
}

func (i *SnapshotRepository) ToSnapshotRepositoryOutputWithContext(ctx context.Context) SnapshotRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotRepositoryOutput)
}

// SnapshotRepositoryArrayInput is an input type that accepts SnapshotRepositoryArray and SnapshotRepositoryArrayOutput values.
// You can construct a concrete instance of `SnapshotRepositoryArrayInput` via:
//
//	SnapshotRepositoryArray{ SnapshotRepositoryArgs{...} }
type SnapshotRepositoryArrayInput interface {
	pulumi.Input

	ToSnapshotRepositoryArrayOutput() SnapshotRepositoryArrayOutput
	ToSnapshotRepositoryArrayOutputWithContext(context.Context) SnapshotRepositoryArrayOutput
}

type SnapshotRepositoryArray []SnapshotRepositoryInput

func (SnapshotRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotRepository)(nil)).Elem()
}

func (i SnapshotRepositoryArray) ToSnapshotRepositoryArrayOutput() SnapshotRepositoryArrayOutput {
	return i.ToSnapshotRepositoryArrayOutputWithContext(context.Background())
}

func (i SnapshotRepositoryArray) ToSnapshotRepositoryArrayOutputWithContext(ctx context.Context) SnapshotRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotRepositoryArrayOutput)
}

// SnapshotRepositoryMapInput is an input type that accepts SnapshotRepositoryMap and SnapshotRepositoryMapOutput values.
// You can construct a concrete instance of `SnapshotRepositoryMapInput` via:
//
//	SnapshotRepositoryMap{ "key": SnapshotRepositoryArgs{...} }
type SnapshotRepositoryMapInput interface {
	pulumi.Input

	ToSnapshotRepositoryMapOutput() SnapshotRepositoryMapOutput
	ToSnapshotRepositoryMapOutputWithContext(context.Context) SnapshotRepositoryMapOutput
}

type SnapshotRepositoryMap map[string]SnapshotRepositoryInput

func (SnapshotRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotRepository)(nil)).Elem()
}

func (i SnapshotRepositoryMap) ToSnapshotRepositoryMapOutput() SnapshotRepositoryMapOutput {
	return i.ToSnapshotRepositoryMapOutputWithContext(context.Background())
}

func (i SnapshotRepositoryMap) ToSnapshotRepositoryMapOutputWithContext(ctx context.Context) SnapshotRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotRepositoryMapOutput)
}

type SnapshotRepositoryOutput struct{ *pulumi.OutputState }

func (SnapshotRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotRepository)(nil)).Elem()
}

func (o SnapshotRepositoryOutput) ToSnapshotRepositoryOutput() SnapshotRepositoryOutput {
	return o
}

func (o SnapshotRepositoryOutput) ToSnapshotRepositoryOutputWithContext(ctx context.Context) SnapshotRepositoryOutput {
	return o
}

// Support for using Azure Blob storage as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-azure.html
func (o SnapshotRepositoryOutput) Azure() SnapshotRepositoryAzurePtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) SnapshotRepositoryAzurePtrOutput { return v.Azure }).(SnapshotRepositoryAzurePtrOutput)
}

// Elasticsearch connection configuration block.
func (o SnapshotRepositoryOutput) ElasticsearchConnection() SnapshotRepositoryElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) SnapshotRepositoryElasticsearchConnectionPtrOutput {
		return v.ElasticsearchConnection
	}).(SnapshotRepositoryElasticsearchConnectionPtrOutput)
}

// Shared filesystem repository. Repositories of this type use a shared filesystem to store snapshots. This filesystem must be accessible to all master and data nodes in the cluster.
func (o SnapshotRepositoryOutput) Fs() SnapshotRepositoryFsPtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) SnapshotRepositoryFsPtrOutput { return v.Fs }).(SnapshotRepositoryFsPtrOutput)
}

// Support for using the Google Cloud Storage service as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-gcs.html
func (o SnapshotRepositoryOutput) Gcs() SnapshotRepositoryGcsPtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) SnapshotRepositoryGcsPtrOutput { return v.Gcs }).(SnapshotRepositoryGcsPtrOutput)
}

// Support for using HDFS File System as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-hdfs.html
func (o SnapshotRepositoryOutput) Hdfs() SnapshotRepositoryHdfsPtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) SnapshotRepositoryHdfsPtrOutput { return v.Hdfs }).(SnapshotRepositoryHdfsPtrOutput)
}

// Name of the snapshot repository to register or update.
func (o SnapshotRepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotRepository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Support for using AWS S3 as a repository for Snapshot/Restore. See: https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-repository.html
func (o SnapshotRepositoryOutput) S3() SnapshotRepositoryS3PtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) SnapshotRepositoryS3PtrOutput { return v.S3 }).(SnapshotRepositoryS3PtrOutput)
}

// URL repository. Repositories of this type are read-only for the cluster. This means the cluster can retrieve or restore snapshots from the repository but cannot write or create snapshots in it.
func (o SnapshotRepositoryOutput) Url() SnapshotRepositoryUrlPtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) SnapshotRepositoryUrlPtrOutput { return v.Url }).(SnapshotRepositoryUrlPtrOutput)
}

// If true, the request verifies the repository is functional on all master and data nodes in the cluster.
func (o SnapshotRepositoryOutput) Verify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotRepository) pulumi.BoolPtrOutput { return v.Verify }).(pulumi.BoolPtrOutput)
}

type SnapshotRepositoryArrayOutput struct{ *pulumi.OutputState }

func (SnapshotRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotRepository)(nil)).Elem()
}

func (o SnapshotRepositoryArrayOutput) ToSnapshotRepositoryArrayOutput() SnapshotRepositoryArrayOutput {
	return o
}

func (o SnapshotRepositoryArrayOutput) ToSnapshotRepositoryArrayOutputWithContext(ctx context.Context) SnapshotRepositoryArrayOutput {
	return o
}

func (o SnapshotRepositoryArrayOutput) Index(i pulumi.IntInput) SnapshotRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotRepository {
		return vs[0].([]*SnapshotRepository)[vs[1].(int)]
	}).(SnapshotRepositoryOutput)
}

type SnapshotRepositoryMapOutput struct{ *pulumi.OutputState }

func (SnapshotRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotRepository)(nil)).Elem()
}

func (o SnapshotRepositoryMapOutput) ToSnapshotRepositoryMapOutput() SnapshotRepositoryMapOutput {
	return o
}

func (o SnapshotRepositoryMapOutput) ToSnapshotRepositoryMapOutputWithContext(ctx context.Context) SnapshotRepositoryMapOutput {
	return o
}

func (o SnapshotRepositoryMapOutput) MapIndex(k pulumi.StringInput) SnapshotRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotRepository {
		return vs[0].(map[string]*SnapshotRepository)[vs[1].(string)]
	}).(SnapshotRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotRepositoryInput)(nil)).Elem(), &SnapshotRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotRepositoryArrayInput)(nil)).Elem(), SnapshotRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotRepositoryMapInput)(nil)).Elem(), SnapshotRepositoryMap{})
	pulumi.RegisterOutputType(SnapshotRepositoryOutput{})
	pulumi.RegisterOutputType(SnapshotRepositoryArrayOutput{})
	pulumi.RegisterOutputType(SnapshotRepositoryMapOutput{})
}
