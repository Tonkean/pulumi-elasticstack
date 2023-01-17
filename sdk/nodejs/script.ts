// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates or updates a stored script or search template. See https://www.elastic.co/guide/en/elasticsearch/reference/current/create-stored-script-api.html
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const myScript = new elasticstack.Script("myScript", {
 *     scriptId: "my_script",
 *     lang: "painless",
 *     source: "Math.log(_score * 2) + params['my_modifier']",
 *     context: "score",
 * });
 * const mySearchTemplate = new elasticstack.Script("mySearchTemplate", {
 *     scriptId: "my_search_template",
 *     lang: "mustache",
 *     source: JSON.stringify({
 *         query: {
 *             match: {
 *                 message: "{{query_string}}",
 *             },
 *         },
 *         from: "{{from}}",
 *         size: "{{size}}",
 *     }),
 *     params: JSON.stringify({
 *         query_string: "My query string",
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import elasticstack:index/script:Script my_script <cluster_uuid>/<script id>
 * ```
 */
export class Script extends pulumi.CustomResource {
    /**
     * Get an existing Script resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScriptState, opts?: pulumi.CustomResourceOptions): Script {
        return new Script(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'elasticstack:index/script:Script';

    /**
     * Returns true if the given object is an instance of Script.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Script {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Script.__pulumiType;
    }

    /**
     * Context in which the script or search template should run.
     */
    public readonly context!: pulumi.Output<string | undefined>;
    /**
     * Elasticsearch connection configuration block.
     */
    public readonly elasticsearchConnection!: pulumi.Output<outputs.ScriptElasticsearchConnection | undefined>;
    /**
     * Script language. For search templates, use `mustache`.
     */
    public readonly lang!: pulumi.Output<string>;
    /**
     * Parameters for the script or search template.
     */
    public readonly params!: pulumi.Output<string | undefined>;
    /**
     * Identifier for the stored script. Must be unique within the cluster.
     */
    public readonly scriptId!: pulumi.Output<string>;
    /**
     * For scripts, a string containing the script. For search templates, an object containing the search template.
     */
    public readonly source!: pulumi.Output<string>;

    /**
     * Create a Script resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScriptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScriptArgs | ScriptState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScriptState | undefined;
            resourceInputs["context"] = state ? state.context : undefined;
            resourceInputs["elasticsearchConnection"] = state ? state.elasticsearchConnection : undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
            resourceInputs["params"] = state ? state.params : undefined;
            resourceInputs["scriptId"] = state ? state.scriptId : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
        } else {
            const args = argsOrState as ScriptArgs | undefined;
            if ((!args || args.lang === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lang'");
            }
            if ((!args || args.scriptId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scriptId'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["context"] = args ? args.context : undefined;
            resourceInputs["elasticsearchConnection"] = args ? args.elasticsearchConnection : undefined;
            resourceInputs["lang"] = args ? args.lang : undefined;
            resourceInputs["params"] = args ? args.params : undefined;
            resourceInputs["scriptId"] = args ? args.scriptId : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Script.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Script resources.
 */
export interface ScriptState {
    /**
     * Context in which the script or search template should run.
     */
    context?: pulumi.Input<string>;
    /**
     * Elasticsearch connection configuration block.
     */
    elasticsearchConnection?: pulumi.Input<inputs.ScriptElasticsearchConnection>;
    /**
     * Script language. For search templates, use `mustache`.
     */
    lang?: pulumi.Input<string>;
    /**
     * Parameters for the script or search template.
     */
    params?: pulumi.Input<string>;
    /**
     * Identifier for the stored script. Must be unique within the cluster.
     */
    scriptId?: pulumi.Input<string>;
    /**
     * For scripts, a string containing the script. For search templates, an object containing the search template.
     */
    source?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Script resource.
 */
export interface ScriptArgs {
    /**
     * Context in which the script or search template should run.
     */
    context?: pulumi.Input<string>;
    /**
     * Elasticsearch connection configuration block.
     */
    elasticsearchConnection?: pulumi.Input<inputs.ScriptElasticsearchConnection>;
    /**
     * Script language. For search templates, use `mustache`.
     */
    lang: pulumi.Input<string>;
    /**
     * Parameters for the script or search template.
     */
    params?: pulumi.Input<string>;
    /**
     * Identifier for the stored script. Must be unique within the cluster.
     */
    scriptId: pulumi.Input<string>;
    /**
     * For scripts, a string containing the script. For search templates, an object containing the search template.
     */
    source: pulumi.Input<string>;
}
