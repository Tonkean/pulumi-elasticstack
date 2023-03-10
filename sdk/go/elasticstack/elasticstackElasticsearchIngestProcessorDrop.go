// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Drops the document without raising any errors. This is useful to prevent the document from getting indexed based on some condition.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/drop-processor.html
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
//			drop, err := elasticstack.ElasticstackElasticsearchIngestProcessorDrop(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorDropArgs{
//				If: pulumi.StringRef("ctx.network_name == 'Guest'"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(drop.Json),
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
func ElasticstackElasticsearchIngestProcessorDrop(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorDropArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorDropResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorDropResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorDrop:ElasticstackElasticsearchIngestProcessorDrop", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorDrop.
type ElasticstackElasticsearchIngestProcessorDropArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorDrop.
type ElasticstackElasticsearchIngestProcessorDropResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
}

func ElasticstackElasticsearchIngestProcessorDropOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorDropOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorDropResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorDropResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorDropArgs)
			r, err := ElasticstackElasticsearchIngestProcessorDrop(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorDropResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorDropResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorDrop.
type ElasticstackElasticsearchIngestProcessorDropOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
}

func (ElasticstackElasticsearchIngestProcessorDropOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorDropArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorDrop.
type ElasticstackElasticsearchIngestProcessorDropResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorDropResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorDropResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) ToElasticstackElasticsearchIngestProcessorDropResultOutput() ElasticstackElasticsearchIngestProcessorDropResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) ToElasticstackElasticsearchIngestProcessorDropResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorDropResultOutput {
	return o
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDropResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Internal identifier of the resource
func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDropResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDropResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDropResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDropResult) string { return v.Json }).(pulumi.StringOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDropResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorDropResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorDropResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorDropResultOutput{})
}
