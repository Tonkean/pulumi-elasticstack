// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an existing Elasticsearch role. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role.html
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
//			roleDataSecurityRole, err := elasticstack.DataSecurityRole(ctx, &elasticstack.DataSecurityRoleArgs{
//				Name: "testrole",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("role", roleDataSecurityRole.Name)
//			return nil
//		})
//	}
//
// ```
func DataSecurityRole(ctx *pulumi.Context, args *DataSecurityRoleArgs, opts ...pulumi.InvokeOption) (*DataSecurityRoleResult, error) {
	var rv DataSecurityRoleResult
	err := ctx.Invoke("elasticstack:index/dataSecurityRole:DataSecurityRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking DataSecurityRole.
type DataSecurityRoleArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *DataSecurityRoleElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// The name of the role.
	Name string `pulumi:"name"`
	// A list of users that the owners of this role can impersonate.
	RunAs []string `pulumi:"runAs"`
}

// A collection of values returned by DataSecurityRole.
type DataSecurityRoleResult struct {
	// A list of application privilege entries.
	Applications []DataSecurityRoleApplication `pulumi:"applications"`
	// A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
	Clusters []string `pulumi:"clusters"`
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *DataSecurityRoleElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// An object defining global privileges.
	Global string `pulumi:"global"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// A list of indices permissions entries.
	Indices []DataSecurityRoleIndex `pulumi:"indices"`
	// Optional meta-data.
	Metadata string `pulumi:"metadata"`
	// The name of the role.
	Name string `pulumi:"name"`
	// A list of users that the owners of this role can impersonate.
	RunAs []string `pulumi:"runAs"`
}

func DataSecurityRoleOutput(ctx *pulumi.Context, args DataSecurityRoleOutputArgs, opts ...pulumi.InvokeOption) DataSecurityRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (DataSecurityRoleResult, error) {
			args := v.(DataSecurityRoleArgs)
			r, err := DataSecurityRole(ctx, &args, opts...)
			var s DataSecurityRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(DataSecurityRoleResultOutput)
}

// A collection of arguments for invoking DataSecurityRole.
type DataSecurityRoleOutputArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection DataSecurityRoleElasticsearchConnectionPtrInput `pulumi:"elasticsearchConnection"`
	// The name of the role.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of users that the owners of this role can impersonate.
	RunAs pulumi.StringArrayInput `pulumi:"runAs"`
}

func (DataSecurityRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSecurityRoleArgs)(nil)).Elem()
}

// A collection of values returned by DataSecurityRole.
type DataSecurityRoleResultOutput struct{ *pulumi.OutputState }

func (DataSecurityRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSecurityRoleResult)(nil)).Elem()
}

func (o DataSecurityRoleResultOutput) ToDataSecurityRoleResultOutput() DataSecurityRoleResultOutput {
	return o
}

func (o DataSecurityRoleResultOutput) ToDataSecurityRoleResultOutputWithContext(ctx context.Context) DataSecurityRoleResultOutput {
	return o
}

// A list of application privilege entries.
func (o DataSecurityRoleResultOutput) Applications() DataSecurityRoleApplicationArrayOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) []DataSecurityRoleApplication { return v.Applications }).(DataSecurityRoleApplicationArrayOutput)
}

// A list of cluster privileges. These privileges define the cluster level actions that users with this role are able to execute.
func (o DataSecurityRoleResultOutput) Clusters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) []string { return v.Clusters }).(pulumi.StringArrayOutput)
}

// Elasticsearch connection configuration block.
func (o DataSecurityRoleResultOutput) ElasticsearchConnection() DataSecurityRoleElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) *DataSecurityRoleElasticsearchConnection {
		return v.ElasticsearchConnection
	}).(DataSecurityRoleElasticsearchConnectionPtrOutput)
}

// An object defining global privileges.
func (o DataSecurityRoleResultOutput) Global() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) string { return v.Global }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o DataSecurityRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of indices permissions entries.
func (o DataSecurityRoleResultOutput) Indices() DataSecurityRoleIndexArrayOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) []DataSecurityRoleIndex { return v.Indices }).(DataSecurityRoleIndexArrayOutput)
}

// Optional meta-data.
func (o DataSecurityRoleResultOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) string { return v.Metadata }).(pulumi.StringOutput)
}

// The name of the role.
func (o DataSecurityRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of users that the owners of this role can impersonate.
func (o DataSecurityRoleResultOutput) RunAs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataSecurityRoleResult) []string { return v.RunAs }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSecurityRoleResultOutput{})
}