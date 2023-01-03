// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Parses a Uniform Resource Identifier (URI) string and extracts its components as an object. This URI object includes properties for the URI’s domain, path, fragment, port, query, scheme, user info, username, and password.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/uri-parts-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean/pulumi-elasticstack";
 *
 * const parts = elasticstack.IngestProcessorUriParts({
 *     field: "input_field",
 *     targetField: "url",
 *     keepOriginal: true,
 *     removeIfSuccessful: false,
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [parts.then(parts => parts.json)]});
 * ```
 */
export function ingestProcessorUriParts(args: IngestProcessorUriPartsArgs, opts?: pulumi.InvokeOptions): Promise<IngestProcessorUriPartsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/ingestProcessorUriParts:IngestProcessorUriParts", {
        "description": args.description,
        "field": args.field,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "keepOriginal": args.keepOriginal,
        "onFailures": args.onFailures,
        "removeIfSuccessful": args.removeIfSuccessful,
        "tag": args.tag,
        "targetField": args.targetField,
    }, opts);
}

/**
 * A collection of arguments for invoking IngestProcessorUriParts.
 */
export interface IngestProcessorUriPartsArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * Field containing the URI string.
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
     * If true, the processor copies the unparsed URI to `<target_field>.original.`
     */
    keepOriginal?: boolean;
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * If `true`, the processor removes the `field` after parsing the URI string. If parsing fails, the processor does not remove the `field`.
     */
    removeIfSuccessful?: boolean;
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * Output field for the URI object.
     */
    targetField?: string;
}

/**
 * A collection of values returned by IngestProcessorUriParts.
 */
export interface IngestProcessorUriPartsResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * Field containing the URI string.
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
     * If true, the processor copies the unparsed URI to `<target_field>.original.`
     */
    readonly keepOriginal?: boolean;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * If `true`, the processor removes the `field` after parsing the URI string. If parsing fails, the processor does not remove the `field`.
     */
    readonly removeIfSuccessful?: boolean;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * Output field for the URI object.
     */
    readonly targetField?: string;
}
/**
 * Parses a Uniform Resource Identifier (URI) string and extracts its components as an object. This URI object includes properties for the URI’s domain, path, fragment, port, query, scheme, user info, username, and password.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/uri-parts-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean/pulumi-elasticstack";
 *
 * const parts = elasticstack.IngestProcessorUriParts({
 *     field: "input_field",
 *     targetField: "url",
 *     keepOriginal: true,
 *     removeIfSuccessful: false,
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [parts.then(parts => parts.json)]});
 * ```
 */
export function ingestProcessorUriPartsOutput(args: IngestProcessorUriPartsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IngestProcessorUriPartsResult> {
    return pulumi.output(args).apply((a: any) => ingestProcessorUriParts(a, opts))
}

/**
 * A collection of arguments for invoking IngestProcessorUriParts.
 */
export interface IngestProcessorUriPartsOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * Field containing the URI string.
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
     * If true, the processor copies the unparsed URI to `<target_field>.original.`
     */
    keepOriginal?: pulumi.Input<boolean>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If `true`, the processor removes the `field` after parsing the URI string. If parsing fails, the processor does not remove the `field`.
     */
    removeIfSuccessful?: pulumi.Input<boolean>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * Output field for the URI object.
     */
    targetField?: pulumi.Input<string>;
}