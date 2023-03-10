// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates or updates a component template. Component templates are building blocks for constructing index templates that specify index mappings, settings, and aliases. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
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
//	"github.com/pulumi/pulumi-elasticstack/sdk/go/elasticstack"
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
//			myTemplateElasticstackElasticsearchComponentTemplate, err := elasticstack.NewElasticstackElasticsearchComponentTemplate(ctx, "myTemplateElasticstackElasticsearchComponentTemplate", &elasticstack.ElasticstackElasticsearchComponentTemplateArgs{
//				Template: &elasticstack.ElasticstackElasticsearchComponentTemplateTemplateArgs{
//					Aliases: elasticstack.ElasticstackElasticsearchComponentTemplateTemplateAliasArray{
//						&elasticstack.ElasticstackElasticsearchComponentTemplateTemplateAliasArgs{
//							Name: pulumi.String("my_template_test"),
//						},
//					},
//					Settings: pulumi.String(json0),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIndexTemplate(ctx, "myTemplateElasticstackElasticsearchIndexTemplate", &elasticstack.ElasticstackElasticsearchIndexTemplateArgs{
//				IndexPatterns: pulumi.StringArray{
//					pulumi.String("stream*"),
//				},
//				ComposedOfs: pulumi.StringArray{
//					myTemplateElasticstackElasticsearchComponentTemplate.Name,
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
//	$ pulumi import elasticstack:index/elasticstackElasticsearchComponentTemplate:ElasticstackElasticsearchComponentTemplate my_template <cluster_uuid>/<component_name>
//
// ```
type ElasticstackElasticsearchComponentTemplate struct {
	pulumi.CustomResourceState

	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchComponentTemplateElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Optional user metadata about the component template.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// Name of the component template to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template ElasticstackElasticsearchComponentTemplateTemplateOutput `pulumi:"template"`
	// Version number used to manage component templates externally.
	Version pulumi.IntPtrOutput `pulumi:"version"`
}

// NewElasticstackElasticsearchComponentTemplate registers a new resource with the given unique name, arguments, and options.
func NewElasticstackElasticsearchComponentTemplate(ctx *pulumi.Context,
	name string, args *ElasticstackElasticsearchComponentTemplateArgs, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchComponentTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	var resource ElasticstackElasticsearchComponentTemplate
	err := ctx.RegisterResource("elasticstack:index/elasticstackElasticsearchComponentTemplate:ElasticstackElasticsearchComponentTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticstackElasticsearchComponentTemplate gets an existing ElasticstackElasticsearchComponentTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticstackElasticsearchComponentTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticstackElasticsearchComponentTemplateState, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchComponentTemplate, error) {
	var resource ElasticstackElasticsearchComponentTemplate
	err := ctx.ReadResource("elasticstack:index/elasticstackElasticsearchComponentTemplate:ElasticstackElasticsearchComponentTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticstackElasticsearchComponentTemplate resources.
type elasticstackElasticsearchComponentTemplateState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchComponentTemplateElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Optional user metadata about the component template.
	Metadata *string `pulumi:"metadata"`
	// Name of the component template to create.
	Name *string `pulumi:"name"`
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template *ElasticstackElasticsearchComponentTemplateTemplate `pulumi:"template"`
	// Version number used to manage component templates externally.
	Version *int `pulumi:"version"`
}

type ElasticstackElasticsearchComponentTemplateState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchComponentTemplateElasticsearchConnectionPtrInput
	// Optional user metadata about the component template.
	Metadata pulumi.StringPtrInput
	// Name of the component template to create.
	Name pulumi.StringPtrInput
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template ElasticstackElasticsearchComponentTemplateTemplatePtrInput
	// Version number used to manage component templates externally.
	Version pulumi.IntPtrInput
}

func (ElasticstackElasticsearchComponentTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchComponentTemplateState)(nil)).Elem()
}

type elasticstackElasticsearchComponentTemplateArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchComponentTemplateElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Optional user metadata about the component template.
	Metadata *string `pulumi:"metadata"`
	// Name of the component template to create.
	Name *string `pulumi:"name"`
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template ElasticstackElasticsearchComponentTemplateTemplate `pulumi:"template"`
	// Version number used to manage component templates externally.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a ElasticstackElasticsearchComponentTemplate resource.
type ElasticstackElasticsearchComponentTemplateArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchComponentTemplateElasticsearchConnectionPtrInput
	// Optional user metadata about the component template.
	Metadata pulumi.StringPtrInput
	// Name of the component template to create.
	Name pulumi.StringPtrInput
	// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
	Template ElasticstackElasticsearchComponentTemplateTemplateInput
	// Version number used to manage component templates externally.
	Version pulumi.IntPtrInput
}

func (ElasticstackElasticsearchComponentTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchComponentTemplateArgs)(nil)).Elem()
}

type ElasticstackElasticsearchComponentTemplateInput interface {
	pulumi.Input

	ToElasticstackElasticsearchComponentTemplateOutput() ElasticstackElasticsearchComponentTemplateOutput
	ToElasticstackElasticsearchComponentTemplateOutputWithContext(ctx context.Context) ElasticstackElasticsearchComponentTemplateOutput
}

func (*ElasticstackElasticsearchComponentTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchComponentTemplate)(nil)).Elem()
}

func (i *ElasticstackElasticsearchComponentTemplate) ToElasticstackElasticsearchComponentTemplateOutput() ElasticstackElasticsearchComponentTemplateOutput {
	return i.ToElasticstackElasticsearchComponentTemplateOutputWithContext(context.Background())
}

func (i *ElasticstackElasticsearchComponentTemplate) ToElasticstackElasticsearchComponentTemplateOutputWithContext(ctx context.Context) ElasticstackElasticsearchComponentTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchComponentTemplateOutput)
}

// ElasticstackElasticsearchComponentTemplateArrayInput is an input type that accepts ElasticstackElasticsearchComponentTemplateArray and ElasticstackElasticsearchComponentTemplateArrayOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchComponentTemplateArrayInput` via:
//
//	ElasticstackElasticsearchComponentTemplateArray{ ElasticstackElasticsearchComponentTemplateArgs{...} }
type ElasticstackElasticsearchComponentTemplateArrayInput interface {
	pulumi.Input

	ToElasticstackElasticsearchComponentTemplateArrayOutput() ElasticstackElasticsearchComponentTemplateArrayOutput
	ToElasticstackElasticsearchComponentTemplateArrayOutputWithContext(context.Context) ElasticstackElasticsearchComponentTemplateArrayOutput
}

type ElasticstackElasticsearchComponentTemplateArray []ElasticstackElasticsearchComponentTemplateInput

func (ElasticstackElasticsearchComponentTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchComponentTemplate)(nil)).Elem()
}

func (i ElasticstackElasticsearchComponentTemplateArray) ToElasticstackElasticsearchComponentTemplateArrayOutput() ElasticstackElasticsearchComponentTemplateArrayOutput {
	return i.ToElasticstackElasticsearchComponentTemplateArrayOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchComponentTemplateArray) ToElasticstackElasticsearchComponentTemplateArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchComponentTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchComponentTemplateArrayOutput)
}

// ElasticstackElasticsearchComponentTemplateMapInput is an input type that accepts ElasticstackElasticsearchComponentTemplateMap and ElasticstackElasticsearchComponentTemplateMapOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchComponentTemplateMapInput` via:
//
//	ElasticstackElasticsearchComponentTemplateMap{ "key": ElasticstackElasticsearchComponentTemplateArgs{...} }
type ElasticstackElasticsearchComponentTemplateMapInput interface {
	pulumi.Input

	ToElasticstackElasticsearchComponentTemplateMapOutput() ElasticstackElasticsearchComponentTemplateMapOutput
	ToElasticstackElasticsearchComponentTemplateMapOutputWithContext(context.Context) ElasticstackElasticsearchComponentTemplateMapOutput
}

type ElasticstackElasticsearchComponentTemplateMap map[string]ElasticstackElasticsearchComponentTemplateInput

func (ElasticstackElasticsearchComponentTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchComponentTemplate)(nil)).Elem()
}

func (i ElasticstackElasticsearchComponentTemplateMap) ToElasticstackElasticsearchComponentTemplateMapOutput() ElasticstackElasticsearchComponentTemplateMapOutput {
	return i.ToElasticstackElasticsearchComponentTemplateMapOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchComponentTemplateMap) ToElasticstackElasticsearchComponentTemplateMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchComponentTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchComponentTemplateMapOutput)
}

type ElasticstackElasticsearchComponentTemplateOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchComponentTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchComponentTemplate)(nil)).Elem()
}

func (o ElasticstackElasticsearchComponentTemplateOutput) ToElasticstackElasticsearchComponentTemplateOutput() ElasticstackElasticsearchComponentTemplateOutput {
	return o
}

func (o ElasticstackElasticsearchComponentTemplateOutput) ToElasticstackElasticsearchComponentTemplateOutputWithContext(ctx context.Context) ElasticstackElasticsearchComponentTemplateOutput {
	return o
}

// Elasticsearch connection configuration block.
func (o ElasticstackElasticsearchComponentTemplateOutput) ElasticsearchConnection() ElasticstackElasticsearchComponentTemplateElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchComponentTemplate) ElasticstackElasticsearchComponentTemplateElasticsearchConnectionPtrOutput {
		return v.ElasticsearchConnection
	}).(ElasticstackElasticsearchComponentTemplateElasticsearchConnectionPtrOutput)
}

// Optional user metadata about the component template.
func (o ElasticstackElasticsearchComponentTemplateOutput) Metadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchComponentTemplate) pulumi.StringPtrOutput { return v.Metadata }).(pulumi.StringPtrOutput)
}

// Name of the component template to create.
func (o ElasticstackElasticsearchComponentTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchComponentTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Template to be applied. It may optionally include an aliases, mappings, or settings configuration.
func (o ElasticstackElasticsearchComponentTemplateOutput) Template() ElasticstackElasticsearchComponentTemplateTemplateOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchComponentTemplate) ElasticstackElasticsearchComponentTemplateTemplateOutput {
		return v.Template
	}).(ElasticstackElasticsearchComponentTemplateTemplateOutput)
}

// Version number used to manage component templates externally.
func (o ElasticstackElasticsearchComponentTemplateOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchComponentTemplate) pulumi.IntPtrOutput { return v.Version }).(pulumi.IntPtrOutput)
}

type ElasticstackElasticsearchComponentTemplateArrayOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchComponentTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchComponentTemplate)(nil)).Elem()
}

func (o ElasticstackElasticsearchComponentTemplateArrayOutput) ToElasticstackElasticsearchComponentTemplateArrayOutput() ElasticstackElasticsearchComponentTemplateArrayOutput {
	return o
}

func (o ElasticstackElasticsearchComponentTemplateArrayOutput) ToElasticstackElasticsearchComponentTemplateArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchComponentTemplateArrayOutput {
	return o
}

func (o ElasticstackElasticsearchComponentTemplateArrayOutput) Index(i pulumi.IntInput) ElasticstackElasticsearchComponentTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchComponentTemplate {
		return vs[0].([]*ElasticstackElasticsearchComponentTemplate)[vs[1].(int)]
	}).(ElasticstackElasticsearchComponentTemplateOutput)
}

type ElasticstackElasticsearchComponentTemplateMapOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchComponentTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchComponentTemplate)(nil)).Elem()
}

func (o ElasticstackElasticsearchComponentTemplateMapOutput) ToElasticstackElasticsearchComponentTemplateMapOutput() ElasticstackElasticsearchComponentTemplateMapOutput {
	return o
}

func (o ElasticstackElasticsearchComponentTemplateMapOutput) ToElasticstackElasticsearchComponentTemplateMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchComponentTemplateMapOutput {
	return o
}

func (o ElasticstackElasticsearchComponentTemplateMapOutput) MapIndex(k pulumi.StringInput) ElasticstackElasticsearchComponentTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchComponentTemplate {
		return vs[0].(map[string]*ElasticstackElasticsearchComponentTemplate)[vs[1].(string)]
	}).(ElasticstackElasticsearchComponentTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchComponentTemplateInput)(nil)).Elem(), &ElasticstackElasticsearchComponentTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchComponentTemplateArrayInput)(nil)).Elem(), ElasticstackElasticsearchComponentTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchComponentTemplateMapInput)(nil)).Elem(), ElasticstackElasticsearchComponentTemplateMap{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchComponentTemplateOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchComponentTemplateArrayOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchComponentTemplateMapOutput{})
}
