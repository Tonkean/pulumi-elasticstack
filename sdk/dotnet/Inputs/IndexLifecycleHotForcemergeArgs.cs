// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class IndexLifecycleHotForcemergeArgs : global::Pulumi.ResourceArgs
    {
        [Input("indexCodec")]
        public Input<string>? IndexCodec { get; set; }

        [Input("maxNumSegments", required: true)]
        public Input<int> MaxNumSegments { get; set; } = null!;

        public IndexLifecycleHotForcemergeArgs()
        {
        }
        public static new IndexLifecycleHotForcemergeArgs Empty => new IndexLifecycleHotForcemergeArgs();
    }
}
