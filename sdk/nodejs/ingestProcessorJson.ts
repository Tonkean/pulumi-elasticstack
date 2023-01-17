// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Converts a JSON string into a structured JSON object.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/json-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const jsonProc = elasticstack.IngestProcessorJson({
 *     field: "string_source",
 *     targetField: "json_target",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [jsonProc.then(jsonProc => jsonProc.json)]});
 * ```
 */
export function ingestProcessorJson(args: IngestProcessorJsonArgs, opts?: pulumi.InvokeOptions): Promise<IngestProcessorJsonResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/ingestProcessorJson:IngestProcessorJson", {
        "addToRoot": args.addToRoot,
        "addToRootConflictStrategy": args.addToRootConflictStrategy,
        "allowDuplicateKeys": args.allowDuplicateKeys,
        "description": args.description,
        "field": args.field,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "onFailures": args.onFailures,
        "tag": args.tag,
        "targetField": args.targetField,
    }, opts);
}

/**
 * A collection of arguments for invoking IngestProcessorJson.
 */
export interface IngestProcessorJsonArgs {
    /**
     * Flag that forces the parsed JSON to be added at the top level of the document. `targetField` must not be set when this option is chosen.
     */
    addToRoot?: boolean;
    /**
     * When set to `replace`, root fields that conflict with fields from the parsed JSON will be overridden. When set to `merge`, conflicting fields will be merged. Only applicable if `addToRoot` is set to `true`.
     */
    addToRootConflictStrategy?: string;
    /**
     * When set to `true`, the JSON parser will not fail if the JSON contains duplicate keys. Instead, the last encountered value for any duplicate key wins.
     */
    allowDuplicateKeys?: boolean;
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * The field to be parsed.
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
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * The field that the converted structured object will be written into. Any existing content in this field will be overwritten.
     */
    targetField?: string;
}

/**
 * A collection of values returned by IngestProcessorJson.
 */
export interface IngestProcessorJsonResult {
    /**
     * Flag that forces the parsed JSON to be added at the top level of the document. `targetField` must not be set when this option is chosen.
     */
    readonly addToRoot?: boolean;
    /**
     * When set to `replace`, root fields that conflict with fields from the parsed JSON will be overridden. When set to `merge`, conflicting fields will be merged. Only applicable if `addToRoot` is set to `true`.
     */
    readonly addToRootConflictStrategy?: string;
    /**
     * When set to `true`, the JSON parser will not fail if the JSON contains duplicate keys. Instead, the last encountered value for any duplicate key wins.
     */
    readonly allowDuplicateKeys?: boolean;
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * The field to be parsed.
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
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * The field that the converted structured object will be written into. Any existing content in this field will be overwritten.
     */
    readonly targetField?: string;
}
/**
 * Converts a JSON string into a structured JSON object.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/json-processor.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const jsonProc = elasticstack.IngestProcessorJson({
 *     field: "string_source",
 *     targetField: "json_target",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [jsonProc.then(jsonProc => jsonProc.json)]});
 * ```
 */
export function ingestProcessorJsonOutput(args: IngestProcessorJsonOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IngestProcessorJsonResult> {
    return pulumi.output(args).apply((a: any) => ingestProcessorJson(a, opts))
}

/**
 * A collection of arguments for invoking IngestProcessorJson.
 */
export interface IngestProcessorJsonOutputArgs {
    /**
     * Flag that forces the parsed JSON to be added at the top level of the document. `targetField` must not be set when this option is chosen.
     */
    addToRoot?: pulumi.Input<boolean>;
    /**
     * When set to `replace`, root fields that conflict with fields from the parsed JSON will be overridden. When set to `merge`, conflicting fields will be merged. Only applicable if `addToRoot` is set to `true`.
     */
    addToRootConflictStrategy?: pulumi.Input<string>;
    /**
     * When set to `true`, the JSON parser will not fail if the JSON contains duplicate keys. Instead, the last encountered value for any duplicate key wins.
     */
    allowDuplicateKeys?: pulumi.Input<boolean>;
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * The field to be parsed.
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
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * The field that the converted structured object will be written into. Any existing content in this field will be overwritten.
     */
    targetField?: pulumi.Input<string>;
}
