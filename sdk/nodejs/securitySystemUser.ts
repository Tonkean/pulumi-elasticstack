// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as elasticstack from "@tonkean-public/pulumi-elasticstack";
 *
 * const kibanaSystem = new elasticstack.SecuritySystemUser("kibanaSystem", {
 *     elasticsearchConnection: {
 *         endpoints: ["http://localhost:9200"],
 *         password: "changeme",
 *         username: "elastic",
 *     },
 *     passwordHash: `$2a$10$rMZe6TdsUwBX/TA8vRDz0OLwKAZeCzXM4jT3tfCjpSTB8HoFuq8xO`,
 *     username: "kibana_system",
 * });
 * ```
 */
export class SecuritySystemUser extends pulumi.CustomResource {
    /**
     * Get an existing SecuritySystemUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecuritySystemUserState, opts?: pulumi.CustomResourceOptions): SecuritySystemUser {
        return new SecuritySystemUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'elasticstack:index/securitySystemUser:SecuritySystemUser';

    /**
     * Returns true if the given object is an instance of SecuritySystemUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecuritySystemUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecuritySystemUser.__pulumiType;
    }

    /**
     * Elasticsearch connection configuration block.
     */
    public readonly elasticsearchConnection!: pulumi.Output<outputs.SecuritySystemUserElasticsearchConnection | undefined>;
    /**
     * Specifies whether the user is enabled. The default value is true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The user???s password. Passwords must be at least 6 characters long.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
     */
    public readonly passwordHash!: pulumi.Output<string | undefined>;
    /**
     * An identifier for the system user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/built-in-users.html).
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a SecuritySystemUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecuritySystemUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecuritySystemUserArgs | SecuritySystemUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecuritySystemUserState | undefined;
            resourceInputs["elasticsearchConnection"] = state ? state.elasticsearchConnection : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordHash"] = state ? state.passwordHash : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as SecuritySystemUserArgs | undefined;
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["elasticsearchConnection"] = args ? args.elasticsearchConnection : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["passwordHash"] = args?.passwordHash ? pulumi.secret(args.passwordHash) : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password", "passwordHash"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecuritySystemUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecuritySystemUser resources.
 */
export interface SecuritySystemUserState {
    /**
     * Elasticsearch connection configuration block.
     */
    elasticsearchConnection?: pulumi.Input<inputs.SecuritySystemUserElasticsearchConnection>;
    /**
     * Specifies whether the user is enabled. The default value is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The user???s password. Passwords must be at least 6 characters long.
     */
    password?: pulumi.Input<string>;
    /**
     * A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
     */
    passwordHash?: pulumi.Input<string>;
    /**
     * An identifier for the system user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/built-in-users.html).
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecuritySystemUser resource.
 */
export interface SecuritySystemUserArgs {
    /**
     * Elasticsearch connection configuration block.
     */
    elasticsearchConnection?: pulumi.Input<inputs.SecuritySystemUserElasticsearchConnection>;
    /**
     * Specifies whether the user is enabled. The default value is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The user???s password. Passwords must be at least 6 characters long.
     */
    password?: pulumi.Input<string>;
    /**
     * A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
     */
    passwordHash?: pulumi.Input<string>;
    /**
     * An identifier for the system user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/built-in-users.html).
     */
    username: pulumi.Input<string>;
}
