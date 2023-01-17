// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Trims whitespace from field. If the field is an array of strings, all members of the array will be trimmed.
 *
 * **NOTE:** This only works on leading and trailing whitespace.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/trim-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const trim = elasticstack.IngestProcessorTrim({
 *     field: "foo",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [trim.then(trim => trim.json)]});
 * ```
 */
export function ingestProcessorTrim(args: IngestProcessorTrimArgs, opts?: pulumi.InvokeOptions): Promise<IngestProcessorTrimResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/ingestProcessorTrim:IngestProcessorTrim", {
        "description": args.description,
        "field": args.field,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "ignoreMissing": args.ignoreMissing,
        "onFailures": args.onFailures,
        "tag": args.tag,
        "targetField": args.targetField,
    }, opts);
}

/**
 * A collection of arguments for invoking IngestProcessorTrim.
 */
export interface IngestProcessorTrimArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * The string-valued field to trim whitespace from.
     */
    field: string;
    /**
     * Conditionally execute the processor
     */
    if?: string;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: boolean;
    /**
     * If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
     */
    ignoreMissing?: boolean;
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * The field to assign the trimmed value to, by default `field` is updated in-place.
     */
    targetField?: string;
}

/**
 * A collection of values returned by IngestProcessorTrim.
 */
export interface IngestProcessorTrimResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * The string-valued field to trim whitespace from.
     */
    readonly field: string;
    /**
     * Internal identifier of the resource.
     */
    readonly id: string;
    /**
     * Conditionally execute the processor
     */
    readonly if?: string;
    /**
     * Ignore failures for the processor.
     */
    readonly ignoreFailure?: boolean;
    /**
     * If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
     */
    readonly ignoreMissing?: boolean;
    /**
     * JSON representation of this data source.
     */
    readonly json: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * The field to assign the trimmed value to, by default `field` is updated in-place.
     */
    readonly targetField?: string;
}
/**
 * Trims whitespace from field. If the field is an array of strings, all members of the array will be trimmed.
 *
 * **NOTE:** This only works on leading and trailing whitespace.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/trim-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const trim = elasticstack.IngestProcessorTrim({
 *     field: "foo",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [trim.then(trim => trim.json)]});
 * ```
 */
export function ingestProcessorTrimOutput(args: IngestProcessorTrimOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IngestProcessorTrimResult> {
    return pulumi.output(args).apply((a: any) => ingestProcessorTrim(a, opts))
}

/**
 * A collection of arguments for invoking IngestProcessorTrim.
 */
export interface IngestProcessorTrimOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * The string-valued field to trim whitespace from.
     */
    field: pulumi.Input<string>;
    /**
     * Conditionally execute the processor
     */
    if?: pulumi.Input<string>;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: pulumi.Input<boolean>;
    /**
     * If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
     */
    ignoreMissing?: pulumi.Input<boolean>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * The field to assign the trimmed value to, by default `field` is updated in-place.
     */
    targetField?: pulumi.Input<string>;
}
