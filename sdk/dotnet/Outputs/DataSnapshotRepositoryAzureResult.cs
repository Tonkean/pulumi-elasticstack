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
    public sealed class DataSnapshotRepositoryAzureResult
    {
        public readonly string BasePath;
        public readonly string ChunkSize;
        public readonly string Client;
        public readonly bool Compress;
        public readonly string Container;
        public readonly string LocationMode;
        public readonly string MaxRestoreBytesPerSec;
        public readonly string MaxSnapshotBytesPerSec;
        public readonly bool Readonly;

        [OutputConstructor]
        private DataSnapshotRepositoryAzureResult(
            string basePath,

            string chunkSize,

            string client,

            bool compress,

            string container,

            string locationMode,

            string maxRestoreBytesPerSec,

            string maxSnapshotBytesPerSec,

            bool @readonly)
        {
            BasePath = basePath;
            ChunkSize = chunkSize;
            Client = client;
            Compress = compress;
            Container = container;
            LocationMode = locationMode;
            MaxRestoreBytesPerSec = maxRestoreBytesPerSec;
            MaxSnapshotBytesPerSec = maxSnapshotBytesPerSec;
            Readonly = @readonly;
        }
    }
}
