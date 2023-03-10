# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'IngestProcessorGeoipResult',
    'AwaitableIngestProcessorGeoipResult',
    'ingest_processor_geoip',
    'ingest_processor_geoip_output',
]

@pulumi.output_type
class IngestProcessorGeoipResult:
    """
    A collection of values returned by IngestProcessorGeoip.
    """
    def __init__(__self__, database_file=None, field=None, first_only=None, id=None, ignore_missing=None, json=None, properties=None, target_field=None):
        if database_file and not isinstance(database_file, str):
            raise TypeError("Expected argument 'database_file' to be a str")
        pulumi.set(__self__, "database_file", database_file)
        if field and not isinstance(field, str):
            raise TypeError("Expected argument 'field' to be a str")
        pulumi.set(__self__, "field", field)
        if first_only and not isinstance(first_only, bool):
            raise TypeError("Expected argument 'first_only' to be a bool")
        pulumi.set(__self__, "first_only", first_only)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_missing and not isinstance(ignore_missing, bool):
            raise TypeError("Expected argument 'ignore_missing' to be a bool")
        pulumi.set(__self__, "ignore_missing", ignore_missing)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if properties and not isinstance(properties, list):
            raise TypeError("Expected argument 'properties' to be a list")
        pulumi.set(__self__, "properties", properties)
        if target_field and not isinstance(target_field, str):
            raise TypeError("Expected argument 'target_field' to be a str")
        pulumi.set(__self__, "target_field", target_field)

    @property
    @pulumi.getter(name="databaseFile")
    def database_file(self) -> Optional[str]:
        """
        The database filename referring to a database the module ships with (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom database in the `ingest-geoip` config directory.
        """
        return pulumi.get(self, "database_file")

    @property
    @pulumi.getter
    def field(self) -> str:
        """
        The field to get the ip address from for the geographical lookup.
        """
        return pulumi.get(self, "field")

    @property
    @pulumi.getter(name="firstOnly")
    def first_only(self) -> Optional[bool]:
        """
        If `true` only first found geoip data will be returned, even if field contains array.
        """
        return pulumi.get(self, "first_only")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Internal identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreMissing")
    def ignore_missing(self) -> Optional[bool]:
        """
        If `true` and `field` does not exist, the processor quietly exits without modifying the document.
        """
        return pulumi.get(self, "ignore_missing")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        JSON representation of this data source.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def properties(self) -> Optional[Sequence[str]]:
        """
        Controls what properties are added to the `target_field` based on the geoip lookup.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="targetField")
    def target_field(self) -> Optional[str]:
        """
        The field that will hold the geographical information looked up from the MaxMind database.
        """
        return pulumi.get(self, "target_field")


class AwaitableIngestProcessorGeoipResult(IngestProcessorGeoipResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorGeoipResult(
            database_file=self.database_file,
            field=self.field,
            first_only=self.first_only,
            id=self.id,
            ignore_missing=self.ignore_missing,
            json=self.json,
            properties=self.properties,
            target_field=self.target_field)


def ingest_processor_geoip(database_file: Optional[str] = None,
                           field: Optional[str] = None,
                           first_only: Optional[bool] = None,
                           ignore_missing: Optional[bool] = None,
                           properties: Optional[Sequence[str]] = None,
                           target_field: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorGeoipResult:
    """
    The geoip processor adds information about the geographical location of an IPv4 or IPv6 address.

    By default, the processor uses the GeoLite2 City, GeoLite2 Country, and GeoLite2 ASN GeoIP2 databases from MaxMind, shared under the CC BY-SA 4.0 license. Elasticsearch automatically downloads updates for these databases from the Elastic GeoIP endpoint: https://geoip.elastic.co/v1/database. To get download statistics for these updates, use the GeoIP stats API.

    If your cluster can???t connect to the Elastic GeoIP endpoint or you want to manage your own updates, [see Manage your own GeoIP2 database updates](https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html#manage-geoip-database-updates).

    If Elasticsearch can???t connect to the endpoint for 30 days all updated databases will become invalid. Elasticsearch will stop enriching documents with geoip data and will add tags: ["_geoip_expired_database"] field instead.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    geoip = elasticstack.ingest_processor_geoip(field="ip")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[geoip.json])
    ```


    :param str database_file: The database filename referring to a database the module ships with (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom database in the `ingest-geoip` config directory.
    :param str field: The field to get the ip address from for the geographical lookup.
    :param bool first_only: If `true` only first found geoip data will be returned, even if field contains array.
    :param bool ignore_missing: If `true` and `field` does not exist, the processor quietly exits without modifying the document.
    :param Sequence[str] properties: Controls what properties are added to the `target_field` based on the geoip lookup.
    :param str target_field: The field that will hold the geographical information looked up from the MaxMind database.
    """
    __args__ = dict()
    __args__['databaseFile'] = database_file
    __args__['field'] = field
    __args__['firstOnly'] = first_only
    __args__['ignoreMissing'] = ignore_missing
    __args__['properties'] = properties
    __args__['targetField'] = target_field
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorGeoip:IngestProcessorGeoip', __args__, opts=opts, typ=IngestProcessorGeoipResult).value

    return AwaitableIngestProcessorGeoipResult(
        database_file=__ret__.database_file,
        field=__ret__.field,
        first_only=__ret__.first_only,
        id=__ret__.id,
        ignore_missing=__ret__.ignore_missing,
        json=__ret__.json,
        properties=__ret__.properties,
        target_field=__ret__.target_field)


@_utilities.lift_output_func(ingest_processor_geoip)
def ingest_processor_geoip_output(database_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  field: Optional[pulumi.Input[str]] = None,
                                  first_only: Optional[pulumi.Input[Optional[bool]]] = None,
                                  ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                                  properties: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  target_field: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorGeoipResult]:
    """
    The geoip processor adds information about the geographical location of an IPv4 or IPv6 address.

    By default, the processor uses the GeoLite2 City, GeoLite2 Country, and GeoLite2 ASN GeoIP2 databases from MaxMind, shared under the CC BY-SA 4.0 license. Elasticsearch automatically downloads updates for these databases from the Elastic GeoIP endpoint: https://geoip.elastic.co/v1/database. To get download statistics for these updates, use the GeoIP stats API.

    If your cluster can???t connect to the Elastic GeoIP endpoint or you want to manage your own updates, [see Manage your own GeoIP2 database updates](https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html#manage-geoip-database-updates).

    If Elasticsearch can???t connect to the endpoint for 30 days all updated databases will become invalid. Elasticsearch will stop enriching documents with geoip data and will add tags: ["_geoip_expired_database"] field instead.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    geoip = elasticstack.ingest_processor_geoip(field="ip")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[geoip.json])
    ```


    :param str database_file: The database filename referring to a database the module ships with (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom database in the `ingest-geoip` config directory.
    :param str field: The field to get the ip address from for the geographical lookup.
    :param bool first_only: If `true` only first found geoip data will be returned, even if field contains array.
    :param bool ignore_missing: If `true` and `field` does not exist, the processor quietly exits without modifying the document.
    :param Sequence[str] properties: Controls what properties are added to the `target_field` based on the geoip lookup.
    :param str target_field: The field that will hold the geographical information looked up from the MaxMind database.
    """
    ...
