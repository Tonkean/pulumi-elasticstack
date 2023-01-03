# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'DataSnapshotRepositoryResult',
    'AwaitableDataSnapshotRepositoryResult',
    'data_snapshot_repository',
    'data_snapshot_repository_output',
]

@pulumi.output_type
class DataSnapshotRepositoryResult:
    """
    A collection of values returned by DataSnapshotRepository.
    """
    def __init__(__self__, azures=None, elasticsearch_connection=None, fs=None, gcs=None, hdfs=None, id=None, name=None, s3s=None, type=None, urls=None):
        if azures and not isinstance(azures, list):
            raise TypeError("Expected argument 'azures' to be a list")
        pulumi.set(__self__, "azures", azures)
        if elasticsearch_connection and not isinstance(elasticsearch_connection, dict):
            raise TypeError("Expected argument 'elasticsearch_connection' to be a dict")
        pulumi.set(__self__, "elasticsearch_connection", elasticsearch_connection)
        if fs and not isinstance(fs, list):
            raise TypeError("Expected argument 'fs' to be a list")
        pulumi.set(__self__, "fs", fs)
        if gcs and not isinstance(gcs, list):
            raise TypeError("Expected argument 'gcs' to be a list")
        pulumi.set(__self__, "gcs", gcs)
        if hdfs and not isinstance(hdfs, list):
            raise TypeError("Expected argument 'hdfs' to be a list")
        pulumi.set(__self__, "hdfs", hdfs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if s3s and not isinstance(s3s, list):
            raise TypeError("Expected argument 's3s' to be a list")
        pulumi.set(__self__, "s3s", s3s)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if urls and not isinstance(urls, list):
            raise TypeError("Expected argument 'urls' to be a list")
        pulumi.set(__self__, "urls", urls)

    @property
    @pulumi.getter
    def azures(self) -> Sequence['outputs.DataSnapshotRepositoryAzureResult']:
        """
        Azure Blob storage as a repository. Set only if the type of the fetched repo is `azure`.
        """
        return pulumi.get(self, "azures")

    @property
    @pulumi.getter(name="elasticsearchConnection")
    def elasticsearch_connection(self) -> Optional['outputs.DataSnapshotRepositoryElasticsearchConnectionResult']:
        """
        Elasticsearch connection configuration block.
        """
        return pulumi.get(self, "elasticsearch_connection")

    @property
    @pulumi.getter
    def fs(self) -> Sequence['outputs.DataSnapshotRepositoryFResult']:
        """
        Shared filesystem repository. Set only if the type of the fetched repo is `fs`.
        """
        return pulumi.get(self, "fs")

    @property
    @pulumi.getter
    def gcs(self) -> Sequence['outputs.DataSnapshotRepositoryGcResult']:
        """
        Google Cloud Storage service as a repository. Set only if the type of the fetched repo is `gcs`.
        """
        return pulumi.get(self, "gcs")

    @property
    @pulumi.getter
    def hdfs(self) -> Sequence['outputs.DataSnapshotRepositoryHdfResult']:
        """
        HDFS File System as a repository. Set only if the type of the fetched repo is `hdfs`.
        """
        return pulumi.get(self, "hdfs")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Internal identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the snapshot repository.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def s3s(self) -> Sequence['outputs.DataSnapshotRepositoryS3Result']:
        """
        AWS S3 as a repository. Set only if the type of the fetched repo is `s3`.
        """
        return pulumi.get(self, "s3s")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Repository type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def urls(self) -> Sequence['outputs.DataSnapshotRepositoryUrlResult']:
        """
        URL repository. Set only if the type of the fetched repo is `url`.
        """
        return pulumi.get(self, "urls")


class AwaitableDataSnapshotRepositoryResult(DataSnapshotRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return DataSnapshotRepositoryResult(
            azures=self.azures,
            elasticsearch_connection=self.elasticsearch_connection,
            fs=self.fs,
            gcs=self.gcs,
            hdfs=self.hdfs,
            id=self.id,
            name=self.name,
            s3s=self.s3s,
            type=self.type,
            urls=self.urls)


def data_snapshot_repository(elasticsearch_connection: Optional[pulumi.InputType['DataSnapshotRepositoryElasticsearchConnectionArgs']] = None,
                             name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableDataSnapshotRepositoryResult:
    """
    This data source provides the information about the registered snaphosts repositories

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    my_url_repo_snapshot_repository = elasticstack.SnapshotRepository("myUrlRepoSnapshotRepository", url=elasticstack.SnapshotRepositoryUrlArgs(
        url="https://example.com/repo",
    ))
    my_fs_repo_snapshot_repository = elasticstack.SnapshotRepository("myFsRepoSnapshotRepository", fs=elasticstack.SnapshotRepositoryFsArgs(
        location="/tmp",
        compress=True,
        max_restore_bytes_per_sec="10mb",
    ))
    my_fs_repo_data_snapshot_repository = elasticstack.data_snapshot_repository_output(name=my_fs_repo_snapshot_repository.name)
    my_url_repo_data_snapshot_repository = elasticstack.data_snapshot_repository_output(name=my_url_repo_snapshot_repository.name)
    pulumi.export("repoFsLocation", my_fs_repo_snapshot_repository.fs.location)
    pulumi.export("repoUrl", my_url_repo_snapshot_repository.url.url)
    ```


    :param pulumi.InputType['DataSnapshotRepositoryElasticsearchConnectionArgs'] elasticsearch_connection: Elasticsearch connection configuration block.
    :param str name: Name of the snapshot repository.
    """
    __args__ = dict()
    __args__['elasticsearchConnection'] = elasticsearch_connection
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/dataSnapshotRepository:DataSnapshotRepository', __args__, opts=opts, typ=DataSnapshotRepositoryResult).value

    return AwaitableDataSnapshotRepositoryResult(
        azures=__ret__.azures,
        elasticsearch_connection=__ret__.elasticsearch_connection,
        fs=__ret__.fs,
        gcs=__ret__.gcs,
        hdfs=__ret__.hdfs,
        id=__ret__.id,
        name=__ret__.name,
        s3s=__ret__.s3s,
        type=__ret__.type,
        urls=__ret__.urls)


@_utilities.lift_output_func(data_snapshot_repository)
def data_snapshot_repository_output(elasticsearch_connection: Optional[pulumi.Input[Optional[pulumi.InputType['DataSnapshotRepositoryElasticsearchConnectionArgs']]]] = None,
                                    name: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[DataSnapshotRepositoryResult]:
    """
    This data source provides the information about the registered snaphosts repositories

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    my_url_repo_snapshot_repository = elasticstack.SnapshotRepository("myUrlRepoSnapshotRepository", url=elasticstack.SnapshotRepositoryUrlArgs(
        url="https://example.com/repo",
    ))
    my_fs_repo_snapshot_repository = elasticstack.SnapshotRepository("myFsRepoSnapshotRepository", fs=elasticstack.SnapshotRepositoryFsArgs(
        location="/tmp",
        compress=True,
        max_restore_bytes_per_sec="10mb",
    ))
    my_fs_repo_data_snapshot_repository = elasticstack.data_snapshot_repository_output(name=my_fs_repo_snapshot_repository.name)
    my_url_repo_data_snapshot_repository = elasticstack.data_snapshot_repository_output(name=my_url_repo_snapshot_repository.name)
    pulumi.export("repoFsLocation", my_fs_repo_snapshot_repository.fs.location)
    pulumi.export("repoUrl", my_url_repo_snapshot_repository.url.url)
    ```


    :param pulumi.InputType['DataSnapshotRepositoryElasticsearchConnectionArgs'] elasticsearch_connection: Elasticsearch connection configuration block.
    :param str name: Name of the snapshot repository.
    """
    ...