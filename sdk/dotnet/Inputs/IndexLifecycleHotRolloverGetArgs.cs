// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class IndexLifecycleHotRolloverGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxAge")]
        public Input<string>? MaxAge { get; set; }

        [Input("maxDocs")]
        public Input<int>? MaxDocs { get; set; }

        [Input("maxPrimaryShardSize")]
        public Input<string>? MaxPrimaryShardSize { get; set; }

        [Input("maxSize")]
        public Input<string>? MaxSize { get; set; }

        public IndexLifecycleHotRolloverGetArgs()
        {
        }
        public static new IndexLifecycleHotRolloverGetArgs Empty => new IndexLifecycleHotRolloverGetArgs();
    }
}
