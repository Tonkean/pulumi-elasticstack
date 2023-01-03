// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document.
 * By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `targetField` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html
 *
 * ## Example Usage
 *
 * Here is an example that adds the parsed date to the `timestamp` field based on the `initialDate` field:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean/pulumi-elasticstack";
 *
 * const date = elasticstack.IngestProcessorDate({
 *     field: "initial_date",
 *     targetField: "timestamp",
 *     formats: ["dd/MM/yyyy HH:mm:ss"],
 *     timezone: "Europe/Amsterdam",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [date.then(date => date.json)]});
 * ```
 */
export function ingestProcessorDate(args: IngestProcessorDateArgs, opts?: pulumi.InvokeOptions): Promise<IngestProcessorDateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("elasticstack:index/ingestProcessorDate:IngestProcessorDate", {
        "description": args.description,
        "field": args.field,
        "formats": args.formats,
        "if": args.if,
        "ignoreFailure": args.ignoreFailure,
        "locale": args.locale,
        "onFailures": args.onFailures,
        "outputFormat": args.outputFormat,
        "tag": args.tag,
        "targetField": args.targetField,
        "timezone": args.timezone,
    }, opts);
}

/**
 * A collection of arguments for invoking IngestProcessorDate.
 */
export interface IngestProcessorDateArgs {
    /**
     * Description of the processor.
     */
    description?: string;
    /**
     * The field to get the date from.
     */
    field: string;
    /**
     * An array of the expected date formats.
     */
    formats: string[];
    /**
     * Conditionally execute the processor
     */
    if?: string;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: boolean;
    /**
     * The locale to use when parsing the date, relevant when parsing month names or week days.
     */
    locale?: string;
    /**
     * Handle failures for the processor.
     */
    onFailures?: string[];
    /**
     * The format to use when writing the date to `targetField`.
     */
    outputFormat?: string;
    /**
     * Identifier for the processor.
     */
    tag?: string;
    /**
     * The field that will hold the parsed date.
     */
    targetField?: string;
    /**
     * The timezone to use when parsing the date.
     */
    timezone?: string;
}

/**
 * A collection of values returned by IngestProcessorDate.
 */
export interface IngestProcessorDateResult {
    /**
     * Description of the processor.
     */
    readonly description?: string;
    /**
     * The field to get the date from.
     */
    readonly field: string;
    /**
     * An array of the expected date formats.
     */
    readonly formats: string[];
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
     * The locale to use when parsing the date, relevant when parsing month names or week days.
     */
    readonly locale?: string;
    /**
     * Handle failures for the processor.
     */
    readonly onFailures?: string[];
    /**
     * The format to use when writing the date to `targetField`.
     */
    readonly outputFormat?: string;
    /**
     * Identifier for the processor.
     */
    readonly tag?: string;
    /**
     * The field that will hold the parsed date.
     */
    readonly targetField?: string;
    /**
     * The timezone to use when parsing the date.
     */
    readonly timezone?: string;
}
/**
 * Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document.
 * By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `targetField` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.
 *
 * See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html
 *
 * ## Example Usage
 *
 * Here is an example that adds the parsed date to the `timestamp` field based on the `initialDate` field:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@pulumi/elasticstack";
 * import * as elasticstack from "@tonkean/pulumi-elasticstack";
 *
 * const date = elasticstack.IngestProcessorDate({
 *     field: "initial_date",
 *     targetField: "timestamp",
 *     formats: ["dd/MM/yyyy HH:mm:ss"],
 *     timezone: "Europe/Amsterdam",
 * });
 * const myIngestPipeline = new elasticstack.IngestPipeline("myIngestPipeline", {processors: [date.then(date => date.json)]});
 * ```
 */
export function ingestProcessorDateOutput(args: IngestProcessorDateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IngestProcessorDateResult> {
    return pulumi.output(args).apply((a: any) => ingestProcessorDate(a, opts))
}

/**
 * A collection of arguments for invoking IngestProcessorDate.
 */
export interface IngestProcessorDateOutputArgs {
    /**
     * Description of the processor.
     */
    description?: pulumi.Input<string>;
    /**
     * The field to get the date from.
     */
    field: pulumi.Input<string>;
    /**
     * An array of the expected date formats.
     */
    formats: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Conditionally execute the processor
     */
    if?: pulumi.Input<string>;
    /**
     * Ignore failures for the processor.
     */
    ignoreFailure?: pulumi.Input<boolean>;
    /**
     * The locale to use when parsing the date, relevant when parsing month names or week days.
     */
    locale?: pulumi.Input<string>;
    /**
     * Handle failures for the processor.
     */
    onFailures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The format to use when writing the date to `targetField`.
     */
    outputFormat?: pulumi.Input<string>;
    /**
     * Identifier for the processor.
     */
    tag?: pulumi.Input<string>;
    /**
     * The field that will hold the parsed date.
     */
    targetField?: pulumi.Input<string>;
    /**
     * The timezone to use when parsing the date.
     */
    timezone?: pulumi.Input<string>;
}
