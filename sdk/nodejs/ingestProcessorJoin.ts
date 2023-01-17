// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Joins each element of an array into a single string using a separator character between each element. Throws an error when the field is not an array.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/join-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const join = elasticstack.IngestProcessorJoin({
 *     field: "joined_array_field",
 *     separator: "-",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [join.then(join => join.json)]});
 * ```
 */
export function ingestProcessorJoin(args: IngestProcessorJoinArgs, opts?: pulumi.InvokeOptions): Promise<IngestProcessorJoinResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/ingestProcessorJoin:IngestProcessorJoin", {
        "description": args.description,
        "field": args.field,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "onFailures": args.onFailures,
        "separator": args.separator,
        "tag": args.tag,
        "targetField": args.targetField,
    }, opts);
}

/**
 * A collection of arguments for invoking IngestProcessorJoin.
 */
export interface IngestProcessorJoinArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * Field containing array values to join.
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
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * The separator character.
     */
    separator: string;
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * The field to assign the converted value to, by default `field` is updated in-place.
     */
    targetField?: string;
}

/**
 * A collection of values returned by IngestProcessorJoin.
 */
export interface IngestProcessorJoinResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * Field containing array values to join.
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
     * JSON representation of this data source.
     */
    readonly json: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * The separator character.
     */
    readonly separator: string;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * The field to assign the converted value to, by default `field` is updated in-place.
     */
    readonly targetField?: string;
}
/**
 * Joins each element of an array into a single string using a separator character between each element. Throws an error when the field is not an array.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/join-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const join = elasticstack.IngestProcessorJoin({
 *     field: "joined_array_field",
 *     separator: "-",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [join.then(join => join.json)]});
 * ```
 */
export function ingestProcessorJoinOutput(args: IngestProcessorJoinOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IngestProcessorJoinResult> {
    return pulumi.output(args).apply((a: any) => ingestProcessorJoin(a, opts))
}

/**
 * A collection of arguments for invoking IngestProcessorJoin.
 */
export interface IngestProcessorJoinOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * Field containing array values to join.
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
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The separator character.
     */
    separator: pulumi.Input<string>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * The field to assign the converted value to, by default `field` is updated in-place.
     */
    targetField?: pulumi.Input<string>;
}
