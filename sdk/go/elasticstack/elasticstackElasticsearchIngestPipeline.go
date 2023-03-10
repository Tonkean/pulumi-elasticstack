// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// ```sh
//
//	$ pulumi import elasticstack:index/elasticstackElasticsearchIngestPipeline:ElasticstackElasticsearchIngestPipeline my_ingest_pipeline <cluster_uuid>/<ingest pipeline name>
//
// ```
type ElasticstackElasticsearchIngestPipeline struct {
	pulumi.CustomResourceState

	// Description of the ingest pipeline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchIngestPipelineElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Optional user metadata about the index template.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// The name of the ingest pipeline.
	Name pulumi.StringOutput `pulumi:"name"`
	// Processors to run immediately after a processor failure. Each processor supports a processor-level `onFailure` value. If a processor without an `onFailure` value fails, Elasticsearch uses this pipeline-level parameter as a fallback. The processors in this parameter run sequentially in the order specified. Elasticsearch will not attempt to run the pipeline’s remaining processors. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document
	OnFailures pulumi.StringArrayOutput `pulumi:"onFailures"`
	// Processors used to perform transformations on documents before indexing. Processors run sequentially in the order specified. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document.
	Processors pulumi.StringArrayOutput `pulumi:"processors"`
}

// NewElasticstackElasticsearchIngestPipeline registers a new resource with the given unique name, arguments, and options.
func NewElasticstackElasticsearchIngestPipeline(ctx *pulumi.Context,
	name string, args *ElasticstackElasticsearchIngestPipelineArgs, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchIngestPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Processors == nil {
		return nil, errors.New("invalid value for required argument 'Processors'")
	}
	var resource ElasticstackElasticsearchIngestPipeline
	err := ctx.RegisterResource("elasticstack:index/elasticstackElasticsearchIngestPipeline:ElasticstackElasticsearchIngestPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticstackElasticsearchIngestPipeline gets an existing ElasticstackElasticsearchIngestPipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticstackElasticsearchIngestPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticstackElasticsearchIngestPipelineState, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchIngestPipeline, error) {
	var resource ElasticstackElasticsearchIngestPipeline
	err := ctx.ReadResource("elasticstack:index/elasticstackElasticsearchIngestPipeline:ElasticstackElasticsearchIngestPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticstackElasticsearchIngestPipeline resources.
type elasticstackElasticsearchIngestPipelineState struct {
	// Description of the ingest pipeline.
	Description *string `pulumi:"description"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchIngestPipelineElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Optional user metadata about the index template.
	Metadata *string `pulumi:"metadata"`
	// The name of the ingest pipeline.
	Name *string `pulumi:"name"`
	// Processors to run immediately after a processor failure. Each processor supports a processor-level `onFailure` value. If a processor without an `onFailure` value fails, Elasticsearch uses this pipeline-level parameter as a fallback. The processors in this parameter run sequentially in the order specified. Elasticsearch will not attempt to run the pipeline’s remaining processors. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document
	OnFailures []string `pulumi:"onFailures"`
	// Processors used to perform transformations on documents before indexing. Processors run sequentially in the order specified. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document.
	Processors []string `pulumi:"processors"`
}

type ElasticstackElasticsearchIngestPipelineState struct {
	// Description of the ingest pipeline.
	Description pulumi.StringPtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchIngestPipelineElasticsearchConnectionPtrInput
	// Optional user metadata about the index template.
	Metadata pulumi.StringPtrInput
	// The name of the ingest pipeline.
	Name pulumi.StringPtrInput
	// Processors to run immediately after a processor failure. Each processor supports a processor-level `onFailure` value. If a processor without an `onFailure` value fails, Elasticsearch uses this pipeline-level parameter as a fallback. The processors in this parameter run sequentially in the order specified. Elasticsearch will not attempt to run the pipeline’s remaining processors. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document
	OnFailures pulumi.StringArrayInput
	// Processors used to perform transformations on documents before indexing. Processors run sequentially in the order specified. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document.
	Processors pulumi.StringArrayInput
}

func (ElasticstackElasticsearchIngestPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchIngestPipelineState)(nil)).Elem()
}

type elasticstackElasticsearchIngestPipelineArgs struct {
	// Description of the ingest pipeline.
	Description *string `pulumi:"description"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchIngestPipelineElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Optional user metadata about the index template.
	Metadata *string `pulumi:"metadata"`
	// The name of the ingest pipeline.
	Name *string `pulumi:"name"`
	// Processors to run immediately after a processor failure. Each processor supports a processor-level `onFailure` value. If a processor without an `onFailure` value fails, Elasticsearch uses this pipeline-level parameter as a fallback. The processors in this parameter run sequentially in the order specified. Elasticsearch will not attempt to run the pipeline’s remaining processors. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document
	OnFailures []string `pulumi:"onFailures"`
	// Processors used to perform transformations on documents before indexing. Processors run sequentially in the order specified. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document.
	Processors []string `pulumi:"processors"`
}

// The set of arguments for constructing a ElasticstackElasticsearchIngestPipeline resource.
type ElasticstackElasticsearchIngestPipelineArgs struct {
	// Description of the ingest pipeline.
	Description pulumi.StringPtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchIngestPipelineElasticsearchConnectionPtrInput
	// Optional user metadata about the index template.
	Metadata pulumi.StringPtrInput
	// The name of the ingest pipeline.
	Name pulumi.StringPtrInput
	// Processors to run immediately after a processor failure. Each processor supports a processor-level `onFailure` value. If a processor without an `onFailure` value fails, Elasticsearch uses this pipeline-level parameter as a fallback. The processors in this parameter run sequentially in the order specified. Elasticsearch will not attempt to run the pipeline’s remaining processors. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document
	OnFailures pulumi.StringArrayInput
	// Processors used to perform transformations on documents before indexing. Processors run sequentially in the order specified. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document.
	Processors pulumi.StringArrayInput
}

func (ElasticstackElasticsearchIngestPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchIngestPipelineArgs)(nil)).Elem()
}

type ElasticstackElasticsearchIngestPipelineInput interface {
	pulumi.Input

	ToElasticstackElasticsearchIngestPipelineOutput() ElasticstackElasticsearchIngestPipelineOutput
	ToElasticstackElasticsearchIngestPipelineOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestPipelineOutput
}

func (*ElasticstackElasticsearchIngestPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchIngestPipeline)(nil)).Elem()
}

func (i *ElasticstackElasticsearchIngestPipeline) ToElasticstackElasticsearchIngestPipelineOutput() ElasticstackElasticsearchIngestPipelineOutput {
	return i.ToElasticstackElasticsearchIngestPipelineOutputWithContext(context.Background())
}

func (i *ElasticstackElasticsearchIngestPipeline) ToElasticstackElasticsearchIngestPipelineOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchIngestPipelineOutput)
}

// ElasticstackElasticsearchIngestPipelineArrayInput is an input type that accepts ElasticstackElasticsearchIngestPipelineArray and ElasticstackElasticsearchIngestPipelineArrayOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchIngestPipelineArrayInput` via:
//
//	ElasticstackElasticsearchIngestPipelineArray{ ElasticstackElasticsearchIngestPipelineArgs{...} }
type ElasticstackElasticsearchIngestPipelineArrayInput interface {
	pulumi.Input

	ToElasticstackElasticsearchIngestPipelineArrayOutput() ElasticstackElasticsearchIngestPipelineArrayOutput
	ToElasticstackElasticsearchIngestPipelineArrayOutputWithContext(context.Context) ElasticstackElasticsearchIngestPipelineArrayOutput
}

type ElasticstackElasticsearchIngestPipelineArray []ElasticstackElasticsearchIngestPipelineInput

func (ElasticstackElasticsearchIngestPipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchIngestPipeline)(nil)).Elem()
}

func (i ElasticstackElasticsearchIngestPipelineArray) ToElasticstackElasticsearchIngestPipelineArrayOutput() ElasticstackElasticsearchIngestPipelineArrayOutput {
	return i.ToElasticstackElasticsearchIngestPipelineArrayOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchIngestPipelineArray) ToElasticstackElasticsearchIngestPipelineArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestPipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchIngestPipelineArrayOutput)
}

// ElasticstackElasticsearchIngestPipelineMapInput is an input type that accepts ElasticstackElasticsearchIngestPipelineMap and ElasticstackElasticsearchIngestPipelineMapOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchIngestPipelineMapInput` via:
//
//	ElasticstackElasticsearchIngestPipelineMap{ "key": ElasticstackElasticsearchIngestPipelineArgs{...} }
type ElasticstackElasticsearchIngestPipelineMapInput interface {
	pulumi.Input

	ToElasticstackElasticsearchIngestPipelineMapOutput() ElasticstackElasticsearchIngestPipelineMapOutput
	ToElasticstackElasticsearchIngestPipelineMapOutputWithContext(context.Context) ElasticstackElasticsearchIngestPipelineMapOutput
}

type ElasticstackElasticsearchIngestPipelineMap map[string]ElasticstackElasticsearchIngestPipelineInput

func (ElasticstackElasticsearchIngestPipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchIngestPipeline)(nil)).Elem()
}

func (i ElasticstackElasticsearchIngestPipelineMap) ToElasticstackElasticsearchIngestPipelineMapOutput() ElasticstackElasticsearchIngestPipelineMapOutput {
	return i.ToElasticstackElasticsearchIngestPipelineMapOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchIngestPipelineMap) ToElasticstackElasticsearchIngestPipelineMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestPipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchIngestPipelineMapOutput)
}

type ElasticstackElasticsearchIngestPipelineOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchIngestPipeline)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestPipelineOutput) ToElasticstackElasticsearchIngestPipelineOutput() ElasticstackElasticsearchIngestPipelineOutput {
	return o
}

func (o ElasticstackElasticsearchIngestPipelineOutput) ToElasticstackElasticsearchIngestPipelineOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestPipelineOutput {
	return o
}

// Description of the ingest pipeline.
func (o ElasticstackElasticsearchIngestPipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchIngestPipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Elasticsearch connection configuration block.
func (o ElasticstackElasticsearchIngestPipelineOutput) ElasticsearchConnection() ElasticstackElasticsearchIngestPipelineElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchIngestPipeline) ElasticstackElasticsearchIngestPipelineElasticsearchConnectionPtrOutput {
		return v.ElasticsearchConnection
	}).(ElasticstackElasticsearchIngestPipelineElasticsearchConnectionPtrOutput)
}

// Optional user metadata about the index template.
func (o ElasticstackElasticsearchIngestPipelineOutput) Metadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchIngestPipeline) pulumi.StringPtrOutput { return v.Metadata }).(pulumi.StringPtrOutput)
}

// The name of the ingest pipeline.
func (o ElasticstackElasticsearchIngestPipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchIngestPipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Processors to run immediately after a processor failure. Each processor supports a processor-level `onFailure` value. If a processor without an `onFailure` value fails, Elasticsearch uses this pipeline-level parameter as a fallback. The processors in this parameter run sequentially in the order specified. Elasticsearch will not attempt to run the pipeline’s remaining processors. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document
func (o ElasticstackElasticsearchIngestPipelineOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchIngestPipeline) pulumi.StringArrayOutput { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Processors used to perform transformations on documents before indexing. Processors run sequentially in the order specified. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html. Each record must be a valid JSON document.
func (o ElasticstackElasticsearchIngestPipelineOutput) Processors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchIngestPipeline) pulumi.StringArrayOutput { return v.Processors }).(pulumi.StringArrayOutput)
}

type ElasticstackElasticsearchIngestPipelineArrayOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestPipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchIngestPipeline)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestPipelineArrayOutput) ToElasticstackElasticsearchIngestPipelineArrayOutput() ElasticstackElasticsearchIngestPipelineArrayOutput {
	return o
}

func (o ElasticstackElasticsearchIngestPipelineArrayOutput) ToElasticstackElasticsearchIngestPipelineArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestPipelineArrayOutput {
	return o
}

func (o ElasticstackElasticsearchIngestPipelineArrayOutput) Index(i pulumi.IntInput) ElasticstackElasticsearchIngestPipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchIngestPipeline {
		return vs[0].([]*ElasticstackElasticsearchIngestPipeline)[vs[1].(int)]
	}).(ElasticstackElasticsearchIngestPipelineOutput)
}

type ElasticstackElasticsearchIngestPipelineMapOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestPipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchIngestPipeline)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestPipelineMapOutput) ToElasticstackElasticsearchIngestPipelineMapOutput() ElasticstackElasticsearchIngestPipelineMapOutput {
	return o
}

func (o ElasticstackElasticsearchIngestPipelineMapOutput) ToElasticstackElasticsearchIngestPipelineMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestPipelineMapOutput {
	return o
}

func (o ElasticstackElasticsearchIngestPipelineMapOutput) MapIndex(k pulumi.StringInput) ElasticstackElasticsearchIngestPipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchIngestPipeline {
		return vs[0].(map[string]*ElasticstackElasticsearchIngestPipeline)[vs[1].(string)]
	}).(ElasticstackElasticsearchIngestPipelineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchIngestPipelineInput)(nil)).Elem(), &ElasticstackElasticsearchIngestPipeline{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchIngestPipelineArrayInput)(nil)).Elem(), ElasticstackElasticsearchIngestPipelineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchIngestPipelineMapInput)(nil)).Elem(), ElasticstackElasticsearchIngestPipelineMap{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestPipelineOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestPipelineArrayOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestPipelineMapOutput{})
}
