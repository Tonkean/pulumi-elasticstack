// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Raises an exception. This is useful for when you expect a pipeline to fail and want to relay a specific message to the requester.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/fail-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean/pulumi-elasticstack";
 *
 * const fail = elasticstack.IngestProcessorFail({
 *     "if": "ctx.tags.contains('production') != true",
 *     message: "The production tag is not present, found tags: {{{tags}}}",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [fail.then(fail => fail.json)]});
 * ```
 */
export function ingestProcessorFail(args: IngestProcessorFailArgs, opts?: pulumi.InvokeOptions): Promise<IngestProcessorFailResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/ingestProcessorFail:IngestProcessorFail", {
        "description": args.description,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "message": args.message,
        "onFailures": args.onFailures,
        "tag": args.tag,
    }, opts);
}

/**
 * A collection of arguments for invoking IngestProcessorFail.
 */
export interface IngestProcessorFailArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * Conditionally execute the processor
     */
    if?: string;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: boolean;
    /**
     * The error message thrown by the processor.
     */
    message: string;
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * Identifier for the processor.
     */
    tag?: string;
}

/**
 * A collection of values returned by IngestProcessorFail.
 */
export interface IngestProcessorFailResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * Internal identifier of the resource
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
     * The error message thrown by the processor.
     */
    readonly message: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
}
/**
 * Raises an exception. This is useful for when you expect a pipeline to fail and want to relay a specific message to the requester.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/fail-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean/pulumi-elasticstack";
 *
 * const fail = elasticstack.IngestProcessorFail({
 *     "if": "ctx.tags.contains('production') != true",
 *     message: "The production tag is not present, found tags: {{{tags}}}",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [fail.then(fail => fail.json)]});
 * ```
 */
export function ingestProcessorFailOutput(args: IngestProcessorFailOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IngestProcessorFailResult> {
    return pulumi.output(args).apply((a: any) => ingestProcessorFail(a, opts))
}

/**
 * A collection of arguments for invoking IngestProcessorFail.
 */
export interface IngestProcessorFailOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * Conditionally execute the processor
     */
    if?: pulumi.Input<string>;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: pulumi.Input<boolean>;
    /**
     * The error message thrown by the processor.
     */
    message: pulumi.Input<string>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
}