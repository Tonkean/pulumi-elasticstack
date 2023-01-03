// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates or updates an index template. Index templates define settings, mappings, and aliases that can be applied automatically to new indices. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-template.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/Tonkean/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"number_of_shards": "3",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = elasticstack.NewIndexTemplate(ctx, "myTemplate", &elasticstack.IndexTemplateArgs{
//				Priority: pulumi.Int(42),
//				IndexPatterns: pulumi.StringArray{
//					pulumi.String("logstash*"),
//					pulumi.String("filebeat*"),
//				},
//				Template: &elasticstack.IndexTemplateTemplateArgs{
//					Aliases: elasticstack.IndexTemplateTemplateAliasArray{
//						&elasticstack.IndexTemplateTemplateAliasArgs{
//							Name: pulumi.String("my_template_test"),
//						},
//						&elasticstack.IndexTemplateTemplateAliasArgs{
//							Name: pulumi.String("another_test"),
//						},
//					},
//					Settings: pulumi.String(json0),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewIndexTemplate(ctx, "myDataStream", &elasticstack.IndexTemplateArgs{
//				IndexPatterns: pulumi.StringArray{
//					pulumi.String("stream*"),
//				},
//				DataStream: nil,
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
//	$ pulumi import elasticstack:index/indexTemplate:IndexTemplate my_template <cluster_uuid>/<template_name>
//
// ```
type IndexTemplate struct {
	pulumi.CustomResourceState

	// An ordered list of component template names.
	ComposedOfs pulumi.StringArrayOutput `pulumi:"composedOfs"`
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	DataStream IndexTemplateDataStreamPtrOutput `pulumi:"dataStream"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection IndexTemplateElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	IndexPatterns pulumi.StringArrayOutput `pulumi:"indexPatterns"`
	// Optional user metadata about the index template.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// Name of the index template to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority to determine index template precedence when a new data stream or index is created.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template IndexTemplateTemplatePtrOutput `pulumi:"template"`
	// Version number used to manage index templates externally.
	Version pulumi.IntPtrOutput `pulumi:"version"`
}

// NewIndexTemplate registers a new resource with the given unique name, arguments, and options.
func NewIndexTemplate(ctx *pulumi.Context,
	name string, args *IndexTemplateArgs, opts ...pulumi.ResourceOption) (*IndexTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IndexPatterns == nil {
		return nil, errors.New("invalid value for required argument 'IndexPatterns'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IndexTemplate
	err := ctx.RegisterResource("elasticstack:index/indexTemplate:IndexTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndexTemplate gets an existing IndexTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndexTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexTemplateState, opts ...pulumi.ResourceOption) (*IndexTemplate, error) {
	var resource IndexTemplate
	err := ctx.ReadResource("elasticstack:index/indexTemplate:IndexTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IndexTemplate resources.
type indexTemplateState struct {
	// An ordered list of component template names.
	ComposedOfs []string `pulumi:"composedOfs"`
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	DataStream *IndexTemplateDataStream `pulumi:"dataStream"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *IndexTemplateElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	IndexPatterns []string `pulumi:"indexPatterns"`
	// Optional user metadata about the index template.
	Metadata *string `pulumi:"metadata"`
	// Name of the index template to create.
	Name *string `pulumi:"name"`
	// Priority to determine index template precedence when a new data stream or index is created.
	Priority *int `pulumi:"priority"`
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template *IndexTemplateTemplate `pulumi:"template"`
	// Version number used to manage index templates externally.
	Version *int `pulumi:"version"`
}

type IndexTemplateState struct {
	// An ordered list of component template names.
	ComposedOfs pulumi.StringArrayInput
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	DataStream IndexTemplateDataStreamPtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection IndexTemplateElasticsearchConnectionPtrInput
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	IndexPatterns pulumi.StringArrayInput
	// Optional user metadata about the index template.
	Metadata pulumi.StringPtrInput
	// Name of the index template to create.
	Name pulumi.StringPtrInput
	// Priority to determine index template precedence when a new data stream or index is created.
	Priority pulumi.IntPtrInput
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template IndexTemplateTemplatePtrInput
	// Version number used to manage index templates externally.
	Version pulumi.IntPtrInput
}

func (IndexTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexTemplateState)(nil)).Elem()
}

type indexTemplateArgs struct {
	// An ordered list of component template names.
	ComposedOfs []string `pulumi:"composedOfs"`
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	DataStream *IndexTemplateDataStream `pulumi:"dataStream"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *IndexTemplateElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	IndexPatterns []string `pulumi:"indexPatterns"`
	// Optional user metadata about the index template.
	Metadata *string `pulumi:"metadata"`
	// Name of the index template to create.
	Name *string `pulumi:"name"`
	// Priority to determine index template precedence when a new data stream or index is created.
	Priority *int `pulumi:"priority"`
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template *IndexTemplateTemplate `pulumi:"template"`
	// Version number used to manage index templates externally.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a IndexTemplate resource.
type IndexTemplateArgs struct {
	// An ordered list of component template names.
	ComposedOfs pulumi.StringArrayInput
	// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
	DataStream IndexTemplateDataStreamPtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection IndexTemplateElasticsearchConnectionPtrInput
	// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
	IndexPatterns pulumi.StringArrayInput
	// Optional user metadata about the index template.
	Metadata pulumi.StringPtrInput
	// Name of the index template to create.
	Name pulumi.StringPtrInput
	// Priority to determine index template precedence when a new data stream or index is created.
	Priority pulumi.IntPtrInput
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template IndexTemplateTemplatePtrInput
	// Version number used to manage index templates externally.
	Version pulumi.IntPtrInput
}

func (IndexTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexTemplateArgs)(nil)).Elem()
}

type IndexTemplateInput interface {
	pulumi.Input

	ToIndexTemplateOutput() IndexTemplateOutput
	ToIndexTemplateOutputWithContext(ctx context.Context) IndexTemplateOutput
}

func (*IndexTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexTemplate)(nil)).Elem()
}

func (i *IndexTemplate) ToIndexTemplateOutput() IndexTemplateOutput {
	return i.ToIndexTemplateOutputWithContext(context.Background())
}

func (i *IndexTemplate) ToIndexTemplateOutputWithContext(ctx context.Context) IndexTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexTemplateOutput)
}

// IndexTemplateArrayInput is an input type that accepts IndexTemplateArray and IndexTemplateArrayOutput values.
// You can construct a concrete instance of `IndexTemplateArrayInput` via:
//
//	IndexTemplateArray{ IndexTemplateArgs{...} }
type IndexTemplateArrayInput interface {
	pulumi.Input

	ToIndexTemplateArrayOutput() IndexTemplateArrayOutput
	ToIndexTemplateArrayOutputWithContext(context.Context) IndexTemplateArrayOutput
}

type IndexTemplateArray []IndexTemplateInput

func (IndexTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IndexTemplate)(nil)).Elem()
}

func (i IndexTemplateArray) ToIndexTemplateArrayOutput() IndexTemplateArrayOutput {
	return i.ToIndexTemplateArrayOutputWithContext(context.Background())
}

func (i IndexTemplateArray) ToIndexTemplateArrayOutputWithContext(ctx context.Context) IndexTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexTemplateArrayOutput)
}

// IndexTemplateMapInput is an input type that accepts IndexTemplateMap and IndexTemplateMapOutput values.
// You can construct a concrete instance of `IndexTemplateMapInput` via:
//
//	IndexTemplateMap{ "key": IndexTemplateArgs{...} }
type IndexTemplateMapInput interface {
	pulumi.Input

	ToIndexTemplateMapOutput() IndexTemplateMapOutput
	ToIndexTemplateMapOutputWithContext(context.Context) IndexTemplateMapOutput
}

type IndexTemplateMap map[string]IndexTemplateInput

func (IndexTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IndexTemplate)(nil)).Elem()
}

func (i IndexTemplateMap) ToIndexTemplateMapOutput() IndexTemplateMapOutput {
	return i.ToIndexTemplateMapOutputWithContext(context.Background())
}

func (i IndexTemplateMap) ToIndexTemplateMapOutputWithContext(ctx context.Context) IndexTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexTemplateMapOutput)
}

type IndexTemplateOutput struct{ *pulumi.OutputState }

func (IndexTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexTemplate)(nil)).Elem()
}

func (o IndexTemplateOutput) ToIndexTemplateOutput() IndexTemplateOutput {
	return o
}

func (o IndexTemplateOutput) ToIndexTemplateOutputWithContext(ctx context.Context) IndexTemplateOutput {
	return o
}

// An ordered list of component template names.
func (o IndexTemplateOutput) ComposedOfs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IndexTemplate) pulumi.StringArrayOutput { return v.ComposedOfs }).(pulumi.StringArrayOutput)
}

// If this object is included, the template is used to create data streams and their backing indices. Supports an empty object.
func (o IndexTemplateOutput) DataStream() IndexTemplateDataStreamPtrOutput {
	return o.ApplyT(func(v *IndexTemplate) IndexTemplateDataStreamPtrOutput { return v.DataStream }).(IndexTemplateDataStreamPtrOutput)
}

// Elasticsearch connection configuration block.
func (o IndexTemplateOutput) ElasticsearchConnection() IndexTemplateElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *IndexTemplate) IndexTemplateElasticsearchConnectionPtrOutput { return v.ElasticsearchConnection }).(IndexTemplateElasticsearchConnectionPtrOutput)
}

// Array of wildcard (*) expressions used to match the names of data streams and indices during creation.
func (o IndexTemplateOutput) IndexPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IndexTemplate) pulumi.StringArrayOutput { return v.IndexPatterns }).(pulumi.StringArrayOutput)
}

// Optional user metadata about the index template.
func (o IndexTemplateOutput) Metadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndexTemplate) pulumi.StringPtrOutput { return v.Metadata }).(pulumi.StringPtrOutput)
}

// Name of the index template to create.
func (o IndexTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IndexTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority to determine index template precedence when a new data stream or index is created.
func (o IndexTemplateOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IndexTemplate) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
func (o IndexTemplateOutput) Template() IndexTemplateTemplatePtrOutput {
	return o.ApplyT(func(v *IndexTemplate) IndexTemplateTemplatePtrOutput { return v.Template }).(IndexTemplateTemplatePtrOutput)
}

// Version number used to manage index templates externally.
func (o IndexTemplateOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IndexTemplate) pulumi.IntPtrOutput { return v.Version }).(pulumi.IntPtrOutput)
}

type IndexTemplateArrayOutput struct{ *pulumi.OutputState }

func (IndexTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IndexTemplate)(nil)).Elem()
}

func (o IndexTemplateArrayOutput) ToIndexTemplateArrayOutput() IndexTemplateArrayOutput {
	return o
}

func (o IndexTemplateArrayOutput) ToIndexTemplateArrayOutputWithContext(ctx context.Context) IndexTemplateArrayOutput {
	return o
}

func (o IndexTemplateArrayOutput) Index(i pulumi.IntInput) IndexTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IndexTemplate {
		return vs[0].([]*IndexTemplate)[vs[1].(int)]
	}).(IndexTemplateOutput)
}

type IndexTemplateMapOutput struct{ *pulumi.OutputState }

func (IndexTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IndexTemplate)(nil)).Elem()
}

func (o IndexTemplateMapOutput) ToIndexTemplateMapOutput() IndexTemplateMapOutput {
	return o
}

func (o IndexTemplateMapOutput) ToIndexTemplateMapOutputWithContext(ctx context.Context) IndexTemplateMapOutput {
	return o
}

func (o IndexTemplateMapOutput) MapIndex(k pulumi.StringInput) IndexTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IndexTemplate {
		return vs[0].(map[string]*IndexTemplate)[vs[1].(string)]
	}).(IndexTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndexTemplateInput)(nil)).Elem(), &IndexTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*IndexTemplateArrayInput)(nil)).Elem(), IndexTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IndexTemplateMapInput)(nil)).Elem(), IndexTemplateMap{})
	pulumi.RegisterOutputType(IndexTemplateOutput{})
	pulumi.RegisterOutputType(IndexTemplateArrayOutput{})
	pulumi.RegisterOutputType(IndexTemplateMapOutput{})
}
