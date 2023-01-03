// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Outputs
{

    [OutputType]
    public sealed class SecurityRoleIndexFieldSecurity
    {
        public readonly ImmutableArray<string> Excepts;
        public readonly ImmutableArray<string> Grants;

        [OutputConstructor]
        private SecurityRoleIndexFieldSecurity(
            ImmutableArray<string> excepts,

            ImmutableArray<string> grants)
        {
            Excepts = excepts;
            Grants = grants;
        }
    }
}