// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Elasticstack.Inputs
{

    public sealed class DataStreamIndexArgs : global::Pulumi.ResourceArgs
    {
        [Input("indexName")]
        public Input<string>? IndexName { get; set; }

        [Input("indexUuid")]
        public Input<string>? IndexUuid { get; set; }

        public DataStreamIndexArgs()
        {
        }
        public static new DataStreamIndexArgs Empty => new DataStreamIndexArgs();
    }
}
