// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates or updates a snapshot lifecycle policy. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-put-policy.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			repo, err := elasticstack.NewElasticstackElasticsearchSnapshotRepository(ctx, "repo", &elasticstack.ElasticstackElasticsearchSnapshotRepositoryArgs{
//				Fs: &elasticstack.ElasticstackElasticsearchSnapshotRepositoryFsArgs{
//					Location:              pulumi.String("/tmp/snapshots"),
//					Compress:              pulumi.Bool(true),
//					MaxRestoreBytesPerSec: pulumi.String("20mb"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchSnapshotLifecycle(ctx, "slmPolicy", &elasticstack.ElasticstackElasticsearchSnapshotLifecycleArgs{
//				Schedule:     pulumi.String("0 30 1 * * ?"),
//				SnapshotName: pulumi.String("<daily-snap-{now/d}>"),
//				Repository:   repo.Name,
//				Indices: pulumi.StringArray{
//					pulumi.String("data-*"),
//					pulumi.String("important"),
//				},
//				IgnoreUnavailable:  pulumi.Bool(false),
//				IncludeGlobalState: pulumi.Bool(false),
//				ExpireAfter:        pulumi.String("30d"),
//				MinCount:           pulumi.Int(5),
//				MaxCount:           pulumi.Int(50),
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
//	$ pulumi import elasticstack:index/elasticstackElasticsearchSnapshotLifecycle:ElasticstackElasticsearchSnapshotLifecycle my_policy <cluster_uuid>/<slm policy name>
//
// ```
type ElasticstackElasticsearchSnapshotLifecycle struct {
	pulumi.CustomResourceState

	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards pulumi.StringPtrOutput `pulumi:"expandWildcards"`
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter pulumi.StringPtrOutput `pulumi:"expireAfter"`
	// Feature states to include in the snapshot.
	FeatureStates pulumi.StringArrayOutput `pulumi:"featureStates"`
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable pulumi.BoolPtrOutput `pulumi:"ignoreUnavailable"`
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState pulumi.BoolPtrOutput `pulumi:"includeGlobalState"`
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices pulumi.StringArrayOutput `pulumi:"indices"`
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// Attaches arbitrary metadata to the snapshot.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// ID for the snapshot lifecycle policy you want to create or update.
	Name pulumi.StringOutput `pulumi:"name"`
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial pulumi.BoolPtrOutput `pulumi:"partial"`
	// Repository used to store snapshots created by this policy.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName pulumi.StringPtrOutput `pulumi:"snapshotName"`
}

// NewElasticstackElasticsearchSnapshotLifecycle registers a new resource with the given unique name, arguments, and options.
func NewElasticstackElasticsearchSnapshotLifecycle(ctx *pulumi.Context,
	name string, args *ElasticstackElasticsearchSnapshotLifecycleArgs, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchSnapshotLifecycle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	var resource ElasticstackElasticsearchSnapshotLifecycle
	err := ctx.RegisterResource("elasticstack:index/elasticstackElasticsearchSnapshotLifecycle:ElasticstackElasticsearchSnapshotLifecycle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticstackElasticsearchSnapshotLifecycle gets an existing ElasticstackElasticsearchSnapshotLifecycle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticstackElasticsearchSnapshotLifecycle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticstackElasticsearchSnapshotLifecycleState, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchSnapshotLifecycle, error) {
	var resource ElasticstackElasticsearchSnapshotLifecycle
	err := ctx.ReadResource("elasticstack:index/elasticstackElasticsearchSnapshotLifecycle:ElasticstackElasticsearchSnapshotLifecycle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticstackElasticsearchSnapshotLifecycle resources.
type elasticstackElasticsearchSnapshotLifecycleState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards *string `pulumi:"expandWildcards"`
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter *string `pulumi:"expireAfter"`
	// Feature states to include in the snapshot.
	FeatureStates []string `pulumi:"featureStates"`
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable *bool `pulumi:"ignoreUnavailable"`
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState *bool `pulumi:"includeGlobalState"`
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices []string `pulumi:"indices"`
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount *int `pulumi:"maxCount"`
	// Attaches arbitrary metadata to the snapshot.
	Metadata *string `pulumi:"metadata"`
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount *int `pulumi:"minCount"`
	// ID for the snapshot lifecycle policy you want to create or update.
	Name *string `pulumi:"name"`
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial *bool `pulumi:"partial"`
	// Repository used to store snapshots created by this policy.
	Repository *string `pulumi:"repository"`
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule *string `pulumi:"schedule"`
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName *string `pulumi:"snapshotName"`
}

type ElasticstackElasticsearchSnapshotLifecycleState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnectionPtrInput
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards pulumi.StringPtrInput
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter pulumi.StringPtrInput
	// Feature states to include in the snapshot.
	FeatureStates pulumi.StringArrayInput
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable pulumi.BoolPtrInput
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState pulumi.BoolPtrInput
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices pulumi.StringArrayInput
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount pulumi.IntPtrInput
	// Attaches arbitrary metadata to the snapshot.
	Metadata pulumi.StringPtrInput
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount pulumi.IntPtrInput
	// ID for the snapshot lifecycle policy you want to create or update.
	Name pulumi.StringPtrInput
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial pulumi.BoolPtrInput
	// Repository used to store snapshots created by this policy.
	Repository pulumi.StringPtrInput
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule pulumi.StringPtrInput
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName pulumi.StringPtrInput
}

func (ElasticstackElasticsearchSnapshotLifecycleState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchSnapshotLifecycleState)(nil)).Elem()
}

type elasticstackElasticsearchSnapshotLifecycleArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards *string `pulumi:"expandWildcards"`
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter *string `pulumi:"expireAfter"`
	// Feature states to include in the snapshot.
	FeatureStates []string `pulumi:"featureStates"`
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable *bool `pulumi:"ignoreUnavailable"`
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState *bool `pulumi:"includeGlobalState"`
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices []string `pulumi:"indices"`
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount *int `pulumi:"maxCount"`
	// Attaches arbitrary metadata to the snapshot.
	Metadata *string `pulumi:"metadata"`
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount *int `pulumi:"minCount"`
	// ID for the snapshot lifecycle policy you want to create or update.
	Name *string `pulumi:"name"`
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial *bool `pulumi:"partial"`
	// Repository used to store snapshots created by this policy.
	Repository string `pulumi:"repository"`
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule string `pulumi:"schedule"`
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName *string `pulumi:"snapshotName"`
}

// The set of arguments for constructing a ElasticstackElasticsearchSnapshotLifecycle resource.
type ElasticstackElasticsearchSnapshotLifecycleArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnectionPtrInput
	// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
	ExpandWildcards pulumi.StringPtrInput
	// Time period after which a snapshot is considered expired and eligible for deletion.
	ExpireAfter pulumi.StringPtrInput
	// Feature states to include in the snapshot.
	FeatureStates pulumi.StringArrayInput
	// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
	IgnoreUnavailable pulumi.BoolPtrInput
	// If `true`, include the cluster state in the snapshot.
	IncludeGlobalState pulumi.BoolPtrInput
	// Comma-separated list of data streams and indices to include in the snapshot.
	Indices pulumi.StringArrayInput
	// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
	MaxCount pulumi.IntPtrInput
	// Attaches arbitrary metadata to the snapshot.
	Metadata pulumi.StringPtrInput
	// Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount pulumi.IntPtrInput
	// ID for the snapshot lifecycle policy you want to create or update.
	Name pulumi.StringPtrInput
	// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
	Partial pulumi.BoolPtrInput
	// Repository used to store snapshots created by this policy.
	Repository pulumi.StringInput
	// Periodic or absolute schedule at which the policy creates snapshots.
	Schedule pulumi.StringInput
	// Name automatically assigned to each snapshot created by the policy.
	SnapshotName pulumi.StringPtrInput
}

func (ElasticstackElasticsearchSnapshotLifecycleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchSnapshotLifecycleArgs)(nil)).Elem()
}

type ElasticstackElasticsearchSnapshotLifecycleInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSnapshotLifecycleOutput() ElasticstackElasticsearchSnapshotLifecycleOutput
	ToElasticstackElasticsearchSnapshotLifecycleOutputWithContext(ctx context.Context) ElasticstackElasticsearchSnapshotLifecycleOutput
}

func (*ElasticstackElasticsearchSnapshotLifecycle) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchSnapshotLifecycle)(nil)).Elem()
}

func (i *ElasticstackElasticsearchSnapshotLifecycle) ToElasticstackElasticsearchSnapshotLifecycleOutput() ElasticstackElasticsearchSnapshotLifecycleOutput {
	return i.ToElasticstackElasticsearchSnapshotLifecycleOutputWithContext(context.Background())
}

func (i *ElasticstackElasticsearchSnapshotLifecycle) ToElasticstackElasticsearchSnapshotLifecycleOutputWithContext(ctx context.Context) ElasticstackElasticsearchSnapshotLifecycleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSnapshotLifecycleOutput)
}

// ElasticstackElasticsearchSnapshotLifecycleArrayInput is an input type that accepts ElasticstackElasticsearchSnapshotLifecycleArray and ElasticstackElasticsearchSnapshotLifecycleArrayOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchSnapshotLifecycleArrayInput` via:
//
//	ElasticstackElasticsearchSnapshotLifecycleArray{ ElasticstackElasticsearchSnapshotLifecycleArgs{...} }
type ElasticstackElasticsearchSnapshotLifecycleArrayInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSnapshotLifecycleArrayOutput() ElasticstackElasticsearchSnapshotLifecycleArrayOutput
	ToElasticstackElasticsearchSnapshotLifecycleArrayOutputWithContext(context.Context) ElasticstackElasticsearchSnapshotLifecycleArrayOutput
}

type ElasticstackElasticsearchSnapshotLifecycleArray []ElasticstackElasticsearchSnapshotLifecycleInput

func (ElasticstackElasticsearchSnapshotLifecycleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchSnapshotLifecycle)(nil)).Elem()
}

func (i ElasticstackElasticsearchSnapshotLifecycleArray) ToElasticstackElasticsearchSnapshotLifecycleArrayOutput() ElasticstackElasticsearchSnapshotLifecycleArrayOutput {
	return i.ToElasticstackElasticsearchSnapshotLifecycleArrayOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchSnapshotLifecycleArray) ToElasticstackElasticsearchSnapshotLifecycleArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchSnapshotLifecycleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSnapshotLifecycleArrayOutput)
}

// ElasticstackElasticsearchSnapshotLifecycleMapInput is an input type that accepts ElasticstackElasticsearchSnapshotLifecycleMap and ElasticstackElasticsearchSnapshotLifecycleMapOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchSnapshotLifecycleMapInput` via:
//
//	ElasticstackElasticsearchSnapshotLifecycleMap{ "key": ElasticstackElasticsearchSnapshotLifecycleArgs{...} }
type ElasticstackElasticsearchSnapshotLifecycleMapInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSnapshotLifecycleMapOutput() ElasticstackElasticsearchSnapshotLifecycleMapOutput
	ToElasticstackElasticsearchSnapshotLifecycleMapOutputWithContext(context.Context) ElasticstackElasticsearchSnapshotLifecycleMapOutput
}

type ElasticstackElasticsearchSnapshotLifecycleMap map[string]ElasticstackElasticsearchSnapshotLifecycleInput

func (ElasticstackElasticsearchSnapshotLifecycleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchSnapshotLifecycle)(nil)).Elem()
}

func (i ElasticstackElasticsearchSnapshotLifecycleMap) ToElasticstackElasticsearchSnapshotLifecycleMapOutput() ElasticstackElasticsearchSnapshotLifecycleMapOutput {
	return i.ToElasticstackElasticsearchSnapshotLifecycleMapOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchSnapshotLifecycleMap) ToElasticstackElasticsearchSnapshotLifecycleMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchSnapshotLifecycleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSnapshotLifecycleMapOutput)
}

type ElasticstackElasticsearchSnapshotLifecycleOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSnapshotLifecycleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchSnapshotLifecycle)(nil)).Elem()
}

func (o ElasticstackElasticsearchSnapshotLifecycleOutput) ToElasticstackElasticsearchSnapshotLifecycleOutput() ElasticstackElasticsearchSnapshotLifecycleOutput {
	return o
}

func (o ElasticstackElasticsearchSnapshotLifecycleOutput) ToElasticstackElasticsearchSnapshotLifecycleOutputWithContext(ctx context.Context) ElasticstackElasticsearchSnapshotLifecycleOutput {
	return o
}

// Elasticsearch connection configuration block.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) ElasticsearchConnection() ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnectionPtrOutput {
		return v.ElasticsearchConnection
	}).(ElasticstackElasticsearchSnapshotLifecycleElasticsearchConnectionPtrOutput)
}

// Determines how wildcard patterns in the `indices` parameter match data streams and indices. Supports comma-separated values, such as `closed,hidden`.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) ExpandWildcards() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringPtrOutput { return v.ExpandWildcards }).(pulumi.StringPtrOutput)
}

// Time period after which a snapshot is considered expired and eligible for deletion.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) ExpireAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringPtrOutput { return v.ExpireAfter }).(pulumi.StringPtrOutput)
}

// Feature states to include in the snapshot.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) FeatureStates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringArrayOutput { return v.FeatureStates }).(pulumi.StringArrayOutput)
}

// If `false`, the snapshot fails if any data stream or index in indices is missing or closed. If `true`, the snapshot ignores missing or closed data streams and indices.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) IgnoreUnavailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.BoolPtrOutput { return v.IgnoreUnavailable }).(pulumi.BoolPtrOutput)
}

// If `true`, include the cluster state in the snapshot.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) IncludeGlobalState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.BoolPtrOutput { return v.IncludeGlobalState }).(pulumi.BoolPtrOutput)
}

// Comma-separated list of data streams and indices to include in the snapshot.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) Indices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringArrayOutput { return v.Indices }).(pulumi.StringArrayOutput)
}

// Maximum number of snapshots to retain, even if the snapshots have not yet expired.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// Attaches arbitrary metadata to the snapshot.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringOutput { return v.Metadata }).(pulumi.StringOutput)
}

// Minimum number of snapshots to retain, even if the snapshots have expired.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

// ID for the snapshot lifecycle policy you want to create or update.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If `false`, the entire snapshot will fail if one or more indices included in the snapshot do not have all primary shards available.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) Partial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.BoolPtrOutput { return v.Partial }).(pulumi.BoolPtrOutput)
}

// Repository used to store snapshots created by this policy.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// Periodic or absolute schedule at which the policy creates snapshots.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// Name automatically assigned to each snapshot created by the policy.
func (o ElasticstackElasticsearchSnapshotLifecycleOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSnapshotLifecycle) pulumi.StringPtrOutput { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

type ElasticstackElasticsearchSnapshotLifecycleArrayOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSnapshotLifecycleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchSnapshotLifecycle)(nil)).Elem()
}

func (o ElasticstackElasticsearchSnapshotLifecycleArrayOutput) ToElasticstackElasticsearchSnapshotLifecycleArrayOutput() ElasticstackElasticsearchSnapshotLifecycleArrayOutput {
	return o
}

func (o ElasticstackElasticsearchSnapshotLifecycleArrayOutput) ToElasticstackElasticsearchSnapshotLifecycleArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchSnapshotLifecycleArrayOutput {
	return o
}

func (o ElasticstackElasticsearchSnapshotLifecycleArrayOutput) Index(i pulumi.IntInput) ElasticstackElasticsearchSnapshotLifecycleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchSnapshotLifecycle {
		return vs[0].([]*ElasticstackElasticsearchSnapshotLifecycle)[vs[1].(int)]
	}).(ElasticstackElasticsearchSnapshotLifecycleOutput)
}

type ElasticstackElasticsearchSnapshotLifecycleMapOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSnapshotLifecycleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchSnapshotLifecycle)(nil)).Elem()
}

func (o ElasticstackElasticsearchSnapshotLifecycleMapOutput) ToElasticstackElasticsearchSnapshotLifecycleMapOutput() ElasticstackElasticsearchSnapshotLifecycleMapOutput {
	return o
}

func (o ElasticstackElasticsearchSnapshotLifecycleMapOutput) ToElasticstackElasticsearchSnapshotLifecycleMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchSnapshotLifecycleMapOutput {
	return o
}

func (o ElasticstackElasticsearchSnapshotLifecycleMapOutput) MapIndex(k pulumi.StringInput) ElasticstackElasticsearchSnapshotLifecycleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchSnapshotLifecycle {
		return vs[0].(map[string]*ElasticstackElasticsearchSnapshotLifecycle)[vs[1].(string)]
	}).(ElasticstackElasticsearchSnapshotLifecycleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSnapshotLifecycleInput)(nil)).Elem(), &ElasticstackElasticsearchSnapshotLifecycle{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSnapshotLifecycleArrayInput)(nil)).Elem(), ElasticstackElasticsearchSnapshotLifecycleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSnapshotLifecycleMapInput)(nil)).Elem(), ElasticstackElasticsearchSnapshotLifecycleMap{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSnapshotLifecycleOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSnapshotLifecycleArrayOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSnapshotLifecycleMapOutput{})
}