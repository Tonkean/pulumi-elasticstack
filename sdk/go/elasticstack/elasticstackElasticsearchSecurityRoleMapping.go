// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage role mappings. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role-mapping.html
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
//				"any": []interface{}{
//					map[string]interface{}{
//						"field": map[string]interface{}{
//							"username": "esadmin",
//						},
//					},
//					map[string]interface{}{
//						"field": map[string]interface{}{
//							"groups": "cn=admins,dc=example,dc=com",
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			example, err := elasticstack.NewElasticstackElasticsearchSecurityRoleMapping(ctx, "example", &elasticstack.ElasticstackElasticsearchSecurityRoleMappingArgs{
//				Enabled: pulumi.Bool(true),
//				Roles: pulumi.StringArray{
//					pulumi.String("admin"),
//				},
//				Rules: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("role", example.Name)
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
//	$ pulumi import elasticstack:index/elasticstackElasticsearchSecurityRoleMapping:ElasticstackElasticsearchSecurityRoleMapping my_role_mapping <cluster_uuid>/<role mapping name>
//
// ```
type ElasticstackElasticsearchSecurityRoleMapping struct {
	pulumi.CustomResourceState

	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
	RoleTemplates pulumi.StringPtrOutput `pulumi:"roleTemplates"`
	// A list of role names that are granted to the users that match the role mapping rules.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
	Rules pulumi.StringOutput `pulumi:"rules"`
}

// NewElasticstackElasticsearchSecurityRoleMapping registers a new resource with the given unique name, arguments, and options.
func NewElasticstackElasticsearchSecurityRoleMapping(ctx *pulumi.Context,
	name string, args *ElasticstackElasticsearchSecurityRoleMappingArgs, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchSecurityRoleMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource ElasticstackElasticsearchSecurityRoleMapping
	err := ctx.RegisterResource("elasticstack:index/elasticstackElasticsearchSecurityRoleMapping:ElasticstackElasticsearchSecurityRoleMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticstackElasticsearchSecurityRoleMapping gets an existing ElasticstackElasticsearchSecurityRoleMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticstackElasticsearchSecurityRoleMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticstackElasticsearchSecurityRoleMappingState, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchSecurityRoleMapping, error) {
	var resource ElasticstackElasticsearchSecurityRoleMapping
	err := ctx.ReadResource("elasticstack:index/elasticstackElasticsearchSecurityRoleMapping:ElasticstackElasticsearchSecurityRoleMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticstackElasticsearchSecurityRoleMapping resources.
type elasticstackElasticsearchSecurityRoleMappingState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
	Enabled *bool `pulumi:"enabled"`
	// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
	Metadata *string `pulumi:"metadata"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name *string `pulumi:"name"`
	// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
	RoleTemplates *string `pulumi:"roleTemplates"`
	// A list of role names that are granted to the users that match the role mapping rules.
	Roles []string `pulumi:"roles"`
	// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
	Rules *string `pulumi:"rules"`
}

type ElasticstackElasticsearchSecurityRoleMappingState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnectionPtrInput
	// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
	Enabled pulumi.BoolPtrInput
	// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
	Metadata pulumi.StringPtrInput
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name pulumi.StringPtrInput
	// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
	RoleTemplates pulumi.StringPtrInput
	// A list of role names that are granted to the users that match the role mapping rules.
	Roles pulumi.StringArrayInput
	// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
	Rules pulumi.StringPtrInput
}

func (ElasticstackElasticsearchSecurityRoleMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchSecurityRoleMappingState)(nil)).Elem()
}

type elasticstackElasticsearchSecurityRoleMappingArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
	Enabled *bool `pulumi:"enabled"`
	// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
	Metadata *string `pulumi:"metadata"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name *string `pulumi:"name"`
	// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
	RoleTemplates *string `pulumi:"roleTemplates"`
	// A list of role names that are granted to the users that match the role mapping rules.
	Roles []string `pulumi:"roles"`
	// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
	Rules string `pulumi:"rules"`
}

// The set of arguments for constructing a ElasticstackElasticsearchSecurityRoleMapping resource.
type ElasticstackElasticsearchSecurityRoleMappingArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnectionPtrInput
	// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
	Enabled pulumi.BoolPtrInput
	// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
	Metadata pulumi.StringPtrInput
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name pulumi.StringPtrInput
	// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
	RoleTemplates pulumi.StringPtrInput
	// A list of role names that are granted to the users that match the role mapping rules.
	Roles pulumi.StringArrayInput
	// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
	Rules pulumi.StringInput
}

func (ElasticstackElasticsearchSecurityRoleMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchSecurityRoleMappingArgs)(nil)).Elem()
}

type ElasticstackElasticsearchSecurityRoleMappingInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSecurityRoleMappingOutput() ElasticstackElasticsearchSecurityRoleMappingOutput
	ToElasticstackElasticsearchSecurityRoleMappingOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityRoleMappingOutput
}

func (*ElasticstackElasticsearchSecurityRoleMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchSecurityRoleMapping)(nil)).Elem()
}

func (i *ElasticstackElasticsearchSecurityRoleMapping) ToElasticstackElasticsearchSecurityRoleMappingOutput() ElasticstackElasticsearchSecurityRoleMappingOutput {
	return i.ToElasticstackElasticsearchSecurityRoleMappingOutputWithContext(context.Background())
}

func (i *ElasticstackElasticsearchSecurityRoleMapping) ToElasticstackElasticsearchSecurityRoleMappingOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityRoleMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSecurityRoleMappingOutput)
}

// ElasticstackElasticsearchSecurityRoleMappingArrayInput is an input type that accepts ElasticstackElasticsearchSecurityRoleMappingArray and ElasticstackElasticsearchSecurityRoleMappingArrayOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchSecurityRoleMappingArrayInput` via:
//
//	ElasticstackElasticsearchSecurityRoleMappingArray{ ElasticstackElasticsearchSecurityRoleMappingArgs{...} }
type ElasticstackElasticsearchSecurityRoleMappingArrayInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSecurityRoleMappingArrayOutput() ElasticstackElasticsearchSecurityRoleMappingArrayOutput
	ToElasticstackElasticsearchSecurityRoleMappingArrayOutputWithContext(context.Context) ElasticstackElasticsearchSecurityRoleMappingArrayOutput
}

type ElasticstackElasticsearchSecurityRoleMappingArray []ElasticstackElasticsearchSecurityRoleMappingInput

func (ElasticstackElasticsearchSecurityRoleMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchSecurityRoleMapping)(nil)).Elem()
}

func (i ElasticstackElasticsearchSecurityRoleMappingArray) ToElasticstackElasticsearchSecurityRoleMappingArrayOutput() ElasticstackElasticsearchSecurityRoleMappingArrayOutput {
	return i.ToElasticstackElasticsearchSecurityRoleMappingArrayOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchSecurityRoleMappingArray) ToElasticstackElasticsearchSecurityRoleMappingArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityRoleMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSecurityRoleMappingArrayOutput)
}

// ElasticstackElasticsearchSecurityRoleMappingMapInput is an input type that accepts ElasticstackElasticsearchSecurityRoleMappingMap and ElasticstackElasticsearchSecurityRoleMappingMapOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchSecurityRoleMappingMapInput` via:
//
//	ElasticstackElasticsearchSecurityRoleMappingMap{ "key": ElasticstackElasticsearchSecurityRoleMappingArgs{...} }
type ElasticstackElasticsearchSecurityRoleMappingMapInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSecurityRoleMappingMapOutput() ElasticstackElasticsearchSecurityRoleMappingMapOutput
	ToElasticstackElasticsearchSecurityRoleMappingMapOutputWithContext(context.Context) ElasticstackElasticsearchSecurityRoleMappingMapOutput
}

type ElasticstackElasticsearchSecurityRoleMappingMap map[string]ElasticstackElasticsearchSecurityRoleMappingInput

func (ElasticstackElasticsearchSecurityRoleMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchSecurityRoleMapping)(nil)).Elem()
}

func (i ElasticstackElasticsearchSecurityRoleMappingMap) ToElasticstackElasticsearchSecurityRoleMappingMapOutput() ElasticstackElasticsearchSecurityRoleMappingMapOutput {
	return i.ToElasticstackElasticsearchSecurityRoleMappingMapOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchSecurityRoleMappingMap) ToElasticstackElasticsearchSecurityRoleMappingMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityRoleMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSecurityRoleMappingMapOutput)
}

type ElasticstackElasticsearchSecurityRoleMappingOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSecurityRoleMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchSecurityRoleMapping)(nil)).Elem()
}

func (o ElasticstackElasticsearchSecurityRoleMappingOutput) ToElasticstackElasticsearchSecurityRoleMappingOutput() ElasticstackElasticsearchSecurityRoleMappingOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityRoleMappingOutput) ToElasticstackElasticsearchSecurityRoleMappingOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityRoleMappingOutput {
	return o
}

// Elasticsearch connection configuration block.
func (o ElasticstackElasticsearchSecurityRoleMappingOutput) ElasticsearchConnection() ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityRoleMapping) ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnectionPtrOutput {
		return v.ElasticsearchConnection
	}).(ElasticstackElasticsearchSecurityRoleMappingElasticsearchConnectionPtrOutput)
}

// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
func (o ElasticstackElasticsearchSecurityRoleMappingOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityRoleMapping) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
func (o ElasticstackElasticsearchSecurityRoleMappingOutput) Metadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityRoleMapping) pulumi.StringPtrOutput { return v.Metadata }).(pulumi.StringPtrOutput)
}

// The distinct name that identifies the role mapping, used solely as an identifier.
func (o ElasticstackElasticsearchSecurityRoleMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityRoleMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
func (o ElasticstackElasticsearchSecurityRoleMappingOutput) RoleTemplates() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityRoleMapping) pulumi.StringPtrOutput { return v.RoleTemplates }).(pulumi.StringPtrOutput)
}

// A list of role names that are granted to the users that match the role mapping rules.
func (o ElasticstackElasticsearchSecurityRoleMappingOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityRoleMapping) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
func (o ElasticstackElasticsearchSecurityRoleMappingOutput) Rules() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityRoleMapping) pulumi.StringOutput { return v.Rules }).(pulumi.StringOutput)
}

type ElasticstackElasticsearchSecurityRoleMappingArrayOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSecurityRoleMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchSecurityRoleMapping)(nil)).Elem()
}

func (o ElasticstackElasticsearchSecurityRoleMappingArrayOutput) ToElasticstackElasticsearchSecurityRoleMappingArrayOutput() ElasticstackElasticsearchSecurityRoleMappingArrayOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityRoleMappingArrayOutput) ToElasticstackElasticsearchSecurityRoleMappingArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityRoleMappingArrayOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityRoleMappingArrayOutput) Index(i pulumi.IntInput) ElasticstackElasticsearchSecurityRoleMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchSecurityRoleMapping {
		return vs[0].([]*ElasticstackElasticsearchSecurityRoleMapping)[vs[1].(int)]
	}).(ElasticstackElasticsearchSecurityRoleMappingOutput)
}

type ElasticstackElasticsearchSecurityRoleMappingMapOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSecurityRoleMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchSecurityRoleMapping)(nil)).Elem()
}

func (o ElasticstackElasticsearchSecurityRoleMappingMapOutput) ToElasticstackElasticsearchSecurityRoleMappingMapOutput() ElasticstackElasticsearchSecurityRoleMappingMapOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityRoleMappingMapOutput) ToElasticstackElasticsearchSecurityRoleMappingMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityRoleMappingMapOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityRoleMappingMapOutput) MapIndex(k pulumi.StringInput) ElasticstackElasticsearchSecurityRoleMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchSecurityRoleMapping {
		return vs[0].(map[string]*ElasticstackElasticsearchSecurityRoleMapping)[vs[1].(string)]
	}).(ElasticstackElasticsearchSecurityRoleMappingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSecurityRoleMappingInput)(nil)).Elem(), &ElasticstackElasticsearchSecurityRoleMapping{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSecurityRoleMappingArrayInput)(nil)).Elem(), ElasticstackElasticsearchSecurityRoleMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSecurityRoleMappingMapInput)(nil)).Elem(), ElasticstackElasticsearchSecurityRoleMappingMap{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSecurityRoleMappingOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSecurityRoleMappingArrayOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSecurityRoleMappingMapOutput{})
}
