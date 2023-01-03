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
//	"github.com/pulumi/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mapping, err := elasticstack.ElasticstackElasticsearchDataSecurityRoleMapping(ctx, &elasticstack.ElasticstackElasticsearchDataSecurityRoleMappingArgs{
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
func ElasticstackElasticsearchDataSecurityRoleMapping(ctx *pulumi.Context, args *ElasticstackElasticsearchDataSecurityRoleMappingArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchDataSecurityRoleMappingResult, error) {
	var rv ElasticstackElasticsearchDataSecurityRoleMappingResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchDataSecurityRoleMapping:ElasticstackElasticsearchDataSecurityRoleMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchDataSecurityRoleMapping.
type ElasticstackElasticsearchDataSecurityRoleMappingArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchDataSecurityRoleMappingElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name string `pulumi:"name"`
}

// A collection of values returned by ElasticstackElasticsearchDataSecurityRoleMapping.
type ElasticstackElasticsearchDataSecurityRoleMappingResult struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchDataSecurityRoleMappingElasticsearchConnection `pulumi:"elasticsearchConnection"`
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

func ElasticstackElasticsearchDataSecurityRoleMappingOutput(ctx *pulumi.Context, args ElasticstackElasticsearchDataSecurityRoleMappingOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchDataSecurityRoleMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchDataSecurityRoleMappingResult, error) {
			args := v.(ElasticstackElasticsearchDataSecurityRoleMappingArgs)
			r, err := ElasticstackElasticsearchDataSecurityRoleMapping(ctx, &args, opts...)
			var s ElasticstackElasticsearchDataSecurityRoleMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchDataSecurityRoleMappingResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchDataSecurityRoleMapping.
type ElasticstackElasticsearchDataSecurityRoleMappingOutputArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchDataSecurityRoleMappingElasticsearchConnectionPtrInput `pulumi:"elasticsearchConnection"`
	// The distinct name that identifies the role mapping, used solely as an identifier.
	Name pulumi.StringInput `pulumi:"name"`
}

func (ElasticstackElasticsearchDataSecurityRoleMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchDataSecurityRoleMappingArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchDataSecurityRoleMapping.
type ElasticstackElasticsearchDataSecurityRoleMappingResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchDataSecurityRoleMappingResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) ToElasticstackElasticsearchDataSecurityRoleMappingResultOutput() ElasticstackElasticsearchDataSecurityRoleMappingResultOutput {
	return o
}

func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) ToElasticstackElasticsearchDataSecurityRoleMappingResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchDataSecurityRoleMappingResultOutput {
	return o
}

// Elasticsearch connection configuration block.
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) ElasticsearchConnection() ElasticstackElasticsearchDataSecurityRoleMappingElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) *ElasticstackElasticsearchDataSecurityRoleMappingElasticsearchConnection {
		return v.ElasticsearchConnection
	}).(ElasticstackElasticsearchDataSecurityRoleMappingElasticsearchConnectionPtrOutput)
}

// Mappings that have `enabled` set to `false` are ignored when role mapping is performed.
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Internal identifier of the resource
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Additional metadata that helps define which roles are assigned to each user. Keys beginning with `_` are reserved for system usage.
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) string { return v.Metadata }).(pulumi.StringOutput)
}

// The distinct name that identifies the role mapping, used solely as an identifier.
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of mustache templates that will be evaluated to determine the roles names that should granted to the users that match the role mapping rules.
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) RoleTemplates() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) string { return v.RoleTemplates }).(pulumi.StringOutput)
}

// A list of role names that are granted to the users that match the role mapping rules.
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// The rules that determine which users should be matched by the mapping. A rule is a logical condition that is expressed by using a JSON DSL.
func (o ElasticstackElasticsearchDataSecurityRoleMappingResultOutput) Rules() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchDataSecurityRoleMappingResult) string { return v.Rules }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchDataSecurityRoleMappingResultOutput{})
}
