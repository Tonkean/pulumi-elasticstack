// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The enrich processor can enrich documents with data from another index. See enrich data section for more information about how to set this up.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-enriching-data.html and https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-processor.html
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
//			enrich, err := elasticstack.ElasticstackElasticsearchIngestProcessorEnrich(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorEnrichArgs{
//				PolicyName:  "users-policy",
//				Field:       "email",
//				TargetField: "user",
//				MaxMatches:  pulumi.IntRef(1),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(enrich.Json),
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
func ElasticstackElasticsearchIngestProcessorEnrich(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorEnrichArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorEnrichResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorEnrichResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorEnrich:ElasticstackElasticsearchIngestProcessorEnrich", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorEnrich.
type ElasticstackElasticsearchIngestProcessorEnrichArgs struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field in the input document that matches the policies matchField used to retrieve the enrichment data.
	Field string `pulumi:"field"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// The maximum number of matched documents to include under the configured target field.
	MaxMatches *int `pulumi:"maxMatches"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// If processor will update fields with pre-existing non-null-valued field.
	Override *bool `pulumi:"override"`
	// The name of the enrich policy to use.
	PolicyName string `pulumi:"policyName"`
	// A spatial relation operator used to match the geoshape of incoming documents to documents in the enrich index.
	ShapeRelation *string `pulumi:"shapeRelation"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// Field added to incoming documents to contain enrich data.
	TargetField string `pulumi:"targetField"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorEnrich.
type ElasticstackElasticsearchIngestProcessorEnrichResult struct {
	// Description of the processor.
	Description *string `pulumi:"description"`
	// The field in the input document that matches the policies matchField used to retrieve the enrichment data.
	Field string `pulumi:"field"`
	// Internal identifier of the resource
	Id string `pulumi:"id"`
	// Conditionally execute the processor
	If *string `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure *bool `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// The maximum number of matched documents to include under the configured target field.
	MaxMatches *int `pulumi:"maxMatches"`
	// Handle failures for the processor.
	OnFailures []string `pulumi:"onFailures"`
	// If processor will update fields with pre-existing non-null-valued field.
	Override *bool `pulumi:"override"`
	// The name of the enrich policy to use.
	PolicyName string `pulumi:"policyName"`
	// A spatial relation operator used to match the geoshape of incoming documents to documents in the enrich index.
	ShapeRelation *string `pulumi:"shapeRelation"`
	// Identifier for the processor.
	Tag *string `pulumi:"tag"`
	// Field added to incoming documents to contain enrich data.
	TargetField string `pulumi:"targetField"`
}

func ElasticstackElasticsearchIngestProcessorEnrichOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorEnrichOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorEnrichResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorEnrichResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorEnrichArgs)
			r, err := ElasticstackElasticsearchIngestProcessorEnrich(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorEnrichResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorEnrichResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorEnrich.
type ElasticstackElasticsearchIngestProcessorEnrichOutputArgs struct {
	// Description of the processor.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The field in the input document that matches the policies matchField used to retrieve the enrichment data.
	Field pulumi.StringInput `pulumi:"field"`
	// Conditionally execute the processor
	If pulumi.StringPtrInput `pulumi:"if"`
	// Ignore failures for the processor.
	IgnoreFailure pulumi.BoolPtrInput `pulumi:"ignoreFailure"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// The maximum number of matched documents to include under the configured target field.
	MaxMatches pulumi.IntPtrInput `pulumi:"maxMatches"`
	// Handle failures for the processor.
	OnFailures pulumi.StringArrayInput `pulumi:"onFailures"`
	// If processor will update fields with pre-existing non-null-valued field.
	Override pulumi.BoolPtrInput `pulumi:"override"`
	// The name of the enrich policy to use.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// A spatial relation operator used to match the geoshape of incoming documents to documents in the enrich index.
	ShapeRelation pulumi.StringPtrInput `pulumi:"shapeRelation"`
	// Identifier for the processor.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// Field added to incoming documents to contain enrich data.
	TargetField pulumi.StringInput `pulumi:"targetField"`
}

func (ElasticstackElasticsearchIngestProcessorEnrichOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorEnrichArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorEnrich.
type ElasticstackElasticsearchIngestProcessorEnrichResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorEnrichResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorEnrichResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) ToElasticstackElasticsearchIngestProcessorEnrichResultOutput() ElasticstackElasticsearchIngestProcessorEnrichResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) ToElasticstackElasticsearchIngestProcessorEnrichResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorEnrichResultOutput {
	return o
}

// Description of the processor.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The field in the input document that matches the policies matchField used to retrieve the enrichment data.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) string { return v.Id }).(pulumi.StringOutput)
}

// Conditionally execute the processor
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) If() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *string { return v.If }).(pulumi.StringPtrOutput)
}

// Ignore failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) IgnoreFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *bool { return v.IgnoreFailure }).(pulumi.BoolPtrOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) string { return v.Json }).(pulumi.StringOutput)
}

// The maximum number of matched documents to include under the configured target field.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) MaxMatches() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *int { return v.MaxMatches }).(pulumi.IntPtrOutput)
}

// Handle failures for the processor.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) OnFailures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) []string { return v.OnFailures }).(pulumi.StringArrayOutput)
}

// If processor will update fields with pre-existing non-null-valued field.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) Override() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *bool { return v.Override }).(pulumi.BoolPtrOutput)
}

// The name of the enrich policy to use.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) string { return v.PolicyName }).(pulumi.StringOutput)
}

// A spatial relation operator used to match the geoshape of incoming documents to documents in the enrich index.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) ShapeRelation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *string { return v.ShapeRelation }).(pulumi.StringPtrOutput)
}

// Identifier for the processor.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// Field added to incoming documents to contain enrich data.
func (o ElasticstackElasticsearchIngestProcessorEnrichResultOutput) TargetField() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorEnrichResult) string { return v.TargetField }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorEnrichResultOutput{})
}
