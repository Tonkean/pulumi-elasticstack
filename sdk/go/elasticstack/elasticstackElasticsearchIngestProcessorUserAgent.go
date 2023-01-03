// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `userAgent` processor extracts details from the user agent string a browser sends with its web requests. This processor adds this information by default under the `userAgent` field.
//
// The ingest-user-agent module ships by default with the regexes.yaml made available by uap-java with an Apache 2.0 license. For more details see https://github.com/ua-parser/uap-core.
//
// See: https://www.elastic.co/guide/en/elasticsearch/reference/current/user-agent-processor.html
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
//			agent, err := elasticstack.ElasticstackElasticsearchIngestProcessorUserAgent(ctx, &elasticstack.ElasticstackElasticsearchIngestProcessorUserAgentArgs{
//				Field: "agent",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticstack.NewElasticstackElasticsearchIngestPipeline(ctx, "myIngestPipeline", &elasticstack.ElasticstackElasticsearchIngestPipelineArgs{
//				Processors: pulumi.StringArray{
//					*pulumi.String(agent.Json),
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
func ElasticstackElasticsearchIngestProcessorUserAgent(ctx *pulumi.Context, args *ElasticstackElasticsearchIngestProcessorUserAgentArgs, opts ...pulumi.InvokeOption) (*ElasticstackElasticsearchIngestProcessorUserAgentResult, error) {
	var rv ElasticstackElasticsearchIngestProcessorUserAgentResult
	err := ctx.Invoke("elasticstack:index/elasticstackElasticsearchIngestProcessorUserAgent:ElasticstackElasticsearchIngestProcessorUserAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorUserAgent.
type ElasticstackElasticsearchIngestProcessorUserAgentArgs struct {
	// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
	ExtractDeviceType *bool `pulumi:"extractDeviceType"`
	// The field containing the user agent string.
	Field string `pulumi:"field"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// Controls what properties are added to `targetField`.
	Properties []string `pulumi:"properties"`
	// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
	RegexFile *string `pulumi:"regexFile"`
	// The field that will be filled with the user agent details.
	TargetField *string `pulumi:"targetField"`
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorUserAgent.
type ElasticstackElasticsearchIngestProcessorUserAgentResult struct {
	// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
	ExtractDeviceType *bool `pulumi:"extractDeviceType"`
	// The field containing the user agent string.
	Field string `pulumi:"field"`
	// Internal identifier of the resource.
	Id string `pulumi:"id"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// JSON representation of this data source.
	Json string `pulumi:"json"`
	// Controls what properties are added to `targetField`.
	Properties []string `pulumi:"properties"`
	// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
	RegexFile *string `pulumi:"regexFile"`
	// The field that will be filled with the user agent details.
	TargetField *string `pulumi:"targetField"`
}

func ElasticstackElasticsearchIngestProcessorUserAgentOutput(ctx *pulumi.Context, args ElasticstackElasticsearchIngestProcessorUserAgentOutputArgs, opts ...pulumi.InvokeOption) ElasticstackElasticsearchIngestProcessorUserAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ElasticstackElasticsearchIngestProcessorUserAgentResult, error) {
			args := v.(ElasticstackElasticsearchIngestProcessorUserAgentArgs)
			r, err := ElasticstackElasticsearchIngestProcessorUserAgent(ctx, &args, opts...)
			var s ElasticstackElasticsearchIngestProcessorUserAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ElasticstackElasticsearchIngestProcessorUserAgentResultOutput)
}

// A collection of arguments for invoking ElasticstackElasticsearchIngestProcessorUserAgent.
type ElasticstackElasticsearchIngestProcessorUserAgentOutputArgs struct {
	// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
	ExtractDeviceType pulumi.BoolPtrInput `pulumi:"extractDeviceType"`
	// The field containing the user agent string.
	Field pulumi.StringInput `pulumi:"field"`
	// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// Controls what properties are added to `targetField`.
	Properties pulumi.StringArrayInput `pulumi:"properties"`
	// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
	RegexFile pulumi.StringPtrInput `pulumi:"regexFile"`
	// The field that will be filled with the user agent details.
	TargetField pulumi.StringPtrInput `pulumi:"targetField"`
}

func (ElasticstackElasticsearchIngestProcessorUserAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorUserAgentArgs)(nil)).Elem()
}

// A collection of values returned by ElasticstackElasticsearchIngestProcessorUserAgent.
type ElasticstackElasticsearchIngestProcessorUserAgentResultOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticstackElasticsearchIngestProcessorUserAgentResult)(nil)).Elem()
}

func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) ToElasticstackElasticsearchIngestProcessorUserAgentResultOutput() ElasticstackElasticsearchIngestProcessorUserAgentResultOutput {
	return o
}

func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) ToElasticstackElasticsearchIngestProcessorUserAgentResultOutputWithContext(ctx context.Context) ElasticstackElasticsearchIngestProcessorUserAgentResultOutput {
	return o
}

// Extracts device type from the user agent string on a best-effort basis. Supported only starting from Elasticsearch version **8.0**
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) ExtractDeviceType() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) *bool { return v.ExtractDeviceType }).(pulumi.BoolPtrOutput)
}

// The field containing the user agent string.
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) string { return v.Field }).(pulumi.StringOutput)
}

// Internal identifier of the resource.
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// JSON representation of this data source.
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) string { return v.Json }).(pulumi.StringOutput)
}

// Controls what properties are added to `targetField`.
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) Properties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) []string { return v.Properties }).(pulumi.StringArrayOutput)
}

// The name of the file in the `config/ingest-user-agent` directory containing the regular expressions for parsing the user agent string.
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) RegexFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) *string { return v.RegexFile }).(pulumi.StringPtrOutput)
}

// The field that will be filled with the user agent details.
func (o ElasticstackElasticsearchIngestProcessorUserAgentResultOutput) TargetField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ElasticstackElasticsearchIngestProcessorUserAgentResult) *string { return v.TargetField }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticstackElasticsearchIngestProcessorUserAgentResultOutput{})
}