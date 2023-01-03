// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves role mappings. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role-mapping.html
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
//			mapping, err := elasticstack.DataSecurityRoleMapping(ctx, &elasticstack.DataSecurityRoleMappingArgs{
//				Name: "my_mapping",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("user", mapping.Name)
//			return nil
//		})
//	}
//
// ```
func DataSecurityRoleMapping(ctx *pulumi.Context, args *DataSecurityRoleMappingArgs, opts ...pulumi.InvokeOption) (*DataSecurityRoleMappingResult, error) {
	var rv DataSecurityRoleMappingResult
	err := ctx.Invoke("elasticstack:index/dataSecurityRoleMapping:DataSecurityRoleMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking DataSecurityRoleMapping.
type DataSecurityRoleMappingArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *DataSecurityRoleMappingElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name string `pulumi:"name"`
}

// A collection of values returned by DataSecurityRoleMapping.
type DataSecurityRoleMappingResult struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *DataSecurityRoleMappingElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
	Enabled bool `pulumi:"enabled"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
	Metadata string `pulumi:"metadata"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name string `pulumi:"name"`
	// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
	RoleTemplates string `pulumi:"roleTemplates"`
	// A list of role names that are granted to the users that match the role mapping rules.
	Roles []string `pulumi:"roles"`
	// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
	Rules string `pulumi:"rules"`
}

func DataSecurityRoleMappingOutput(ctx *pulumi.Context, args DataSecurityRoleMappingOutputArgs, opts ...pulumi.InvokeOption) DataSecurityRoleMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (DataSecurityRoleMappingResult, error) {
			args := v.(DataSecurityRoleMappingArgs)
			r, err := DataSecurityRoleMapping(ctx, &args, opts...)
			var s DataSecurityRoleMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(DataSecurityRoleMappingResultOutput)
}

// A collection of arguments for invoking DataSecurityRoleMapping.
type DataSecurityRoleMappingOutputArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection DataSecurityRoleMappingElasticsearchConnectionPtrInput `pulumi:"elasticsearchConnection"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name pulumi.StringInput `pulumi:"name"`
}

func (DataSecurityRoleMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSecurityRoleMappingArgs)(nil)).Elem()
}

// A collection of values returned by DataSecurityRoleMapping.
type DataSecurityRoleMappingResultOutput struct{ *pulumi.OutputState }

func (DataSecurityRoleMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSecurityRoleMappingResult)(nil)).Elem()
}

func (o DataSecurityRoleMappingResultOutput) ToDataSecurityRoleMappingResultOutput() DataSecurityRoleMappingResultOutput {
	return o
}

func (o DataSecurityRoleMappingResultOutput) ToDataSecurityRoleMappingResultOutputWithContext(ctx context.Context) DataSecurityRoleMappingResultOutput {
	return o
}

// Elasticsearch connection configuration block.
func (o DataSecurityRoleMappingResultOutput) ElasticsearchConnection() DataSecurityRoleMappingElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) *DataSecurityRoleMappingElasticsearchConnection {
		return v.ElasticsearchConnection
	}).(DataSecurityRoleMappingElasticsearchConnectionPtrOutput)
}

// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
func (o DataSecurityRoleMappingResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Internal identifier of the resource
func (o DataSecurityRoleMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
func (o DataSecurityRoleMappingResultOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) string { return v.Metadata }).(pulumi.StringOutput)
}

// The distinct name that identifies the role mapping, used solely as an identifier.
func (o DataSecurityRoleMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
func (o DataSecurityRoleMappingResultOutput) RoleTemplates() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) string { return v.RoleTemplates }).(pulumi.StringOutput)
}

// A list of role names that are granted to the users that match the role mapping rules.
func (o DataSecurityRoleMappingResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
func (o DataSecurityRoleMappingResultOutput) Rules() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleMappingResult) string { return v.Rules }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSecurityRoleMappingResultOutput{})
}
