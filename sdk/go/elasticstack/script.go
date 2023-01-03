// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates or updates a stored script or search template. See https://www.elastic.co/guide/en/elasticsearch/reference/current/create-stored-script-api.html
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
//			_, err := elasticstack.NewScript(ctx, "myScript", &elasticstack.ScriptArgs{
//				ScriptId: pulumi.String("my_script"),
//				Lang:     pulumi.String("painless"),
//				Source:   pulumi.String("Math.log(_score * 2) + params['my_modifier']"),
//				Context:  pulumi.String("score"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"query": map[string]interface{}{
//					"match": map[string]interface{}{
//						"message": "{{query_string}}",
//					},
//				},
//				"from": "{{from}}",
//				"size": "{{size}}",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"query_string": "My query string",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = elasticstack.NewScript(ctx, "mySearchTemplate", &elasticstack.ScriptArgs{
//				ScriptId: pulumi.String("my_search_template"),
//				Lang:     pulumi.String("mustache"),
//				Source:   pulumi.String(json0),
//				Params:   pulumi.String(json1),
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
//	$ pulumi import elasticstack:index/script:Script my_script <cluster_uuid>/<script id>
//
// ```
type Script struct {
	pulumi.CustomResourceState

	// Context in which the script or search template should run.
	Context pulumi.StringPtrOutput `pulumi:"context"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ScriptElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Script language. For search templates, use `mustache`.
	Lang pulumi.StringOutput `pulumi:"lang"`
	// Parameters for the script or search template.
	Params pulumi.StringPtrOutput `pulumi:"params"`
	// Identifier for the stored script. Must be unique within the cluster.
	ScriptId pulumi.StringOutput `pulumi:"scriptId"`
	// For scripts, a string containing the script. For search templates, an object containing the search template.
	Source pulumi.StringOutput `pulumi:"source"`
}

// NewScript registers a new resource with the given unique name, arguments, and options.
func NewScript(ctx *pulumi.Context,
	name string, args *ScriptArgs, opts ...pulumi.ResourceOption) (*Script, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Lang == nil {
		return nil, errors.New("invalid value for required argument 'Lang'")
	}
	if args.ScriptId == nil {
		return nil, errors.New("invalid value for required argument 'ScriptId'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Script
	err := ctx.RegisterResource("elasticstack:index/script:Script", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScript gets an existing Script resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptState, opts ...pulumi.ResourceOption) (*Script, error) {
	var resource Script
	err := ctx.ReadResource("elasticstack:index/script:Script", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Script resources.
type scriptState struct {
	// Context in which the script or search template should run.
	Context *string `pulumi:"context"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ScriptElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Script language. For search templates, use `mustache`.
	Lang *string `pulumi:"lang"`
	// Parameters for the script or search template.
	Params *string `pulumi:"params"`
	// Identifier for the stored script. Must be unique within the cluster.
	ScriptId *string `pulumi:"scriptId"`
	// For scripts, a string containing the script. For search templates, an object containing the search template.
	Source *string `pulumi:"source"`
}

type ScriptState struct {
	// Context in which the script or search template should run.
	Context pulumi.StringPtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ScriptElasticsearchConnectionPtrInput
	// Script language. For search templates, use `mustache`.
	Lang pulumi.StringPtrInput
	// Parameters for the script or search template.
	Params pulumi.StringPtrInput
	// Identifier for the stored script. Must be unique within the cluster.
	ScriptId pulumi.StringPtrInput
	// For scripts, a string containing the script. For search templates, an object containing the search template.
	Source pulumi.StringPtrInput
}

func (ScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptState)(nil)).Elem()
}

type scriptArgs struct {
	// Context in which the script or search template should run.
	Context *string `pulumi:"context"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ScriptElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Script language. For search templates, use `mustache`.
	Lang string `pulumi:"lang"`
	// Parameters for the script or search template.
	Params *string `pulumi:"params"`
	// Identifier for the stored script. Must be unique within the cluster.
	ScriptId string `pulumi:"scriptId"`
	// For scripts, a string containing the script. For search templates, an object containing the search template.
	Source string `pulumi:"source"`
}

// The set of arguments for constructing a Script resource.
type ScriptArgs struct {
	// Context in which the script or search template should run.
	Context pulumi.StringPtrInput
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ScriptElasticsearchConnectionPtrInput
	// Script language. For search templates, use `mustache`.
	Lang pulumi.StringInput
	// Parameters for the script or search template.
	Params pulumi.StringPtrInput
	// Identifier for the stored script. Must be unique within the cluster.
	ScriptId pulumi.StringInput
	// For scripts, a string containing the script. For search templates, an object containing the search template.
	Source pulumi.StringInput
}

func (ScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptArgs)(nil)).Elem()
}

type ScriptInput interface {
	pulumi.Input

	ToScriptOutput() ScriptOutput
	ToScriptOutputWithContext(ctx context.Context) ScriptOutput
}

func (*Script) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (i *Script) ToScriptOutput() ScriptOutput {
	return i.ToScriptOutputWithContext(context.Background())
}

func (i *Script) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptOutput)
}

// ScriptArrayInput is an input type that accepts ScriptArray and ScriptArrayOutput values.
// You can construct a concrete instance of `ScriptArrayInput` via:
//
//	ScriptArray{ ScriptArgs{...} }
type ScriptArrayInput interface {
	pulumi.Input

	ToScriptArrayOutput() ScriptArrayOutput
	ToScriptArrayOutputWithContext(context.Context) ScriptArrayOutput
}

type ScriptArray []ScriptInput

func (ScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Script)(nil)).Elem()
}

func (i ScriptArray) ToScriptArrayOutput() ScriptArrayOutput {
	return i.ToScriptArrayOutputWithContext(context.Background())
}

func (i ScriptArray) ToScriptArrayOutputWithContext(ctx context.Context) ScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptArrayOutput)
}

// ScriptMapInput is an input type that accepts ScriptMap and ScriptMapOutput values.
// You can construct a concrete instance of `ScriptMapInput` via:
//
//	ScriptMap{ "key": ScriptArgs{...} }
type ScriptMapInput interface {
	pulumi.Input

	ToScriptMapOutput() ScriptMapOutput
	ToScriptMapOutputWithContext(context.Context) ScriptMapOutput
}

type ScriptMap map[string]ScriptInput

func (ScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Script)(nil)).Elem()
}

func (i ScriptMap) ToScriptMapOutput() ScriptMapOutput {
	return i.ToScriptMapOutputWithContext(context.Background())
}

func (i ScriptMap) ToScriptMapOutputWithContext(ctx context.Context) ScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptMapOutput)
}

type ScriptOutput struct{ *pulumi.OutputState }

func (ScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (o ScriptOutput) ToScriptOutput() ScriptOutput {
	return o
}

func (o ScriptOutput) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return o
}

// Context in which the script or search template should run.
func (o ScriptOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.StringPtrOutput { return v.Context }).(pulumi.StringPtrOutput)
}

// Elasticsearch connection configuration block.
func (o ScriptOutput) ElasticsearchConnection() ScriptElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *Script) ScriptElasticsearchConnectionPtrOutput { return v.ElasticsearchConnection }).(ScriptElasticsearchConnectionPtrOutput)
}

// Script language. For search templates, use `mustache`.
func (o ScriptOutput) Lang() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Lang }).(pulumi.StringOutput)
}

// Parameters for the script or search template.
func (o ScriptOutput) Params() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.StringPtrOutput { return v.Params }).(pulumi.StringPtrOutput)
}

// Identifier for the stored script. Must be unique within the cluster.
func (o ScriptOutput) ScriptId() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.ScriptId }).(pulumi.StringOutput)
}

// For scripts, a string containing the script. For search templates, an object containing the search template.
func (o ScriptOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

type ScriptArrayOutput struct{ *pulumi.OutputState }

func (ScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Script)(nil)).Elem()
}

func (o ScriptArrayOutput) ToScriptArrayOutput() ScriptArrayOutput {
	return o
}

func (o ScriptArrayOutput) ToScriptArrayOutputWithContext(ctx context.Context) ScriptArrayOutput {
	return o
}

func (o ScriptArrayOutput) Index(i pulumi.IntInput) ScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Script {
		return vs[0].([]*Script)[vs[1].(int)]
	}).(ScriptOutput)
}

type ScriptMapOutput struct{ *pulumi.OutputState }

func (ScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Script)(nil)).Elem()
}

func (o ScriptMapOutput) ToScriptMapOutput() ScriptMapOutput {
	return o
}

func (o ScriptMapOutput) ToScriptMapOutputWithContext(ctx context.Context) ScriptMapOutput {
	return o
}

func (o ScriptMapOutput) MapIndex(k pulumi.StringInput) ScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Script {
		return vs[0].(map[string]*Script)[vs[1].(string)]
	}).(ScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptInput)(nil)).Elem(), &Script{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptArrayInput)(nil)).Elem(), ScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptMapInput)(nil)).Elem(), ScriptMap{})
	pulumi.RegisterOutputType(ScriptOutput{})
	pulumi.RegisterOutputType(ScriptArrayOutput{})
	pulumi.RegisterOutputType(ScriptMapOutput{})
}
