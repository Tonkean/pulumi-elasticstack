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
    'IngestProcessorDateIndexNameResult',
    'AwaitableIngestProcessorDateIndexNameResult',
    'ingest_processor_date_index_name',
    'ingest_processor_date_index_name_output',
]

@pulumi.output_type
class IngestProcessorDateIndexNameResult:
    """
    A collection of values returned by IngestProcessorDateIndexName.
    """
    def __init__(__self__, date_formats=None, date_rounding=None, description=None, field=None, id=None, if_=None, ignore_failure=None, index_name_format=None, index_name_prefix=None, json=None, locale=None, on_failures=None, tag=None, timezone=None):
        if date_formats and not isinstance(date_formats, list):
            raise TypeError("Expected argument 'date_formats' to be a list")
        pulumi.set(__self__, "date_formats", date_formats)
        if date_rounding and not isinstance(date_rounding, str):
            raise TypeError("Expected argument 'date_rounding' to be a str")
        pulumi.set(__self__, "date_rounding", date_rounding)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if field and not isinstance(field, str):
            raise TypeError("Expected argument 'field' to be a str")
        pulumi.set(__self__, "field", field)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if if_ and not isinstance(if_, str):
            raise TypeError("Expected argument 'if_' to be a str")
        pulumi.set(__self__, "if_", if_)
        if ignore_failure and not isinstance(ignore_failure, bool):
            raise TypeError("Expected argument 'ignore_failure' to be a bool")
        pulumi.set(__self__, "ignore_failure", ignore_failure)
        if index_name_format and not isinstance(index_name_format, str):
            raise TypeError("Expected argument 'index_name_format' to be a str")
        pulumi.set(__self__, "index_name_format", index_name_format)
        if index_name_prefix and not isinstance(index_name_prefix, str):
            raise TypeError("Expected argument 'index_name_prefix' to be a str")
        pulumi.set(__self__, "index_name_prefix", index_name_prefix)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if locale and not isinstance(locale, str):
            raise TypeError("Expected argument 'locale' to be a str")
        pulumi.set(__self__, "locale", locale)
        if on_failures and not isinstance(on_failures, list):
            raise TypeError("Expected argument 'on_failures' to be a list")
        pulumi.set(__self__, "on_failures", on_failures)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if timezone and not isinstance(timezone, str):
            raise TypeError("Expected argument 'timezone' to be a str")
        pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter(name="dateFormats")
    def date_formats(self) -> Optional[Sequence[str]]:
        """
        An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
        """
        return pulumi.get(self, "date_formats")

    @property
    @pulumi.getter(name="dateRounding")
    def date_rounding(self) -> str:
        """
        How to round the date when formatting the date into the index name.
        """
        return pulumi.get(self, "date_rounding")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the processor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def field(self) -> str:
        """
        The field to get the date or timestamp from.
        """
        return pulumi.get(self, "field")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Internal identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="if")
    def if_(self) -> Optional[str]:
        """
        Conditionally execute the processor
        """
        return pulumi.get(self, "if_")

    @property
    @pulumi.getter(name="ignoreFailure")
    def ignore_failure(self) -> Optional[bool]:
        """
        Ignore failures for the processor.
        """
        return pulumi.get(self, "ignore_failure")

    @property
    @pulumi.getter(name="indexNameFormat")
    def index_name_format(self) -> Optional[str]:
        """
        The format to be used when printing the parsed date into the index name.
        """
        return pulumi.get(self, "index_name_format")

    @property
    @pulumi.getter(name="indexNamePrefix")
    def index_name_prefix(self) -> Optional[str]:
        """
        A prefix of the index name to be prepended before the printed date.
        """
        return pulumi.get(self, "index_name_prefix")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        JSON representation of this data source.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def locale(self) -> Optional[str]:
        """
        The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
        """
        return pulumi.get(self, "locale")

    @property
    @pulumi.getter(name="onFailures")
    def on_failures(self) -> Optional[Sequence[str]]:
        """
        Handle failures for the processor.
        """
        return pulumi.get(self, "on_failures")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def timezone(self) -> Optional[str]:
        """
        The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
        """
        return pulumi.get(self, "timezone")


class AwaitableIngestProcessorDateIndexNameResult(IngestProcessorDateIndexNameResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorDateIndexNameResult(
            date_formats=self.date_formats,
            date_rounding=self.date_rounding,
            description=self.description,
            field=self.field,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            index_name_format=self.index_name_format,
            index_name_prefix=self.index_name_prefix,
            json=self.json,
            locale=self.locale,
            on_failures=self.on_failures,
            tag=self.tag,
            timezone=self.timezone)


def ingest_processor_date_index_name(date_formats: Optional[Sequence[str]] = None,
                                     date_rounding: Optional[str] = None,
                                     description: Optional[str] = None,
                                     field: Optional[str] = None,
                                     if_: Optional[str] = None,
                                     ignore_failure: Optional[bool] = None,
                                     index_name_format: Optional[str] = None,
                                     index_name_prefix: Optional[str] = None,
                                     locale: Optional[str] = None,
                                     on_failures: Optional[Sequence[str]] = None,
                                     tag: Optional[str] = None,
                                     timezone: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorDateIndexNameResult:
    """
    The purpose of this processor is to point documents to the right time based index based on a date or timestamp field in a document by using the date math index name support.

    The processor sets the _index metadata field with a date math index name expression based on the provided index name prefix, a date or timestamp field in the documents being processed and the provided date rounding.

    First, this processor fetches the date or timestamp from a field in the document being processed. Optionally, date formatting can be configured on how the field’s value should be parsed into a date. Then this date, the provided index name prefix and the provided date rounding get formatted into a date math index name expression. Also here optionally date formatting can be specified on how the date should be formatted into a date math index name expression.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-index-name-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    date_index_name = elasticstack.ingest_processor_date_index_name(description="monthly date-time index naming",
        field="date1",
        index_name_prefix="my-index-",
        date_rounding="M")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[date_index_name.json])
    ```


    :param Sequence[str] date_formats: An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
    :param str date_rounding: How to round the date when formatting the date into the index name.
    :param str description: Description of the processor.
    :param str field: The field to get the date or timestamp from.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param str index_name_format: The format to be used when printing the parsed date into the index name.
    :param str index_name_prefix: A prefix of the index name to be prepended before the printed date.
    :param str locale: The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str tag: Identifier for the processor.
    :param str timezone: The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
    """
    __args__ = dict()
    __args__['dateFormats'] = date_formats
    __args__['dateRounding'] = date_rounding
    __args__['description'] = description
    __args__['field'] = field
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['indexNameFormat'] = index_name_format
    __args__['indexNamePrefix'] = index_name_prefix
    __args__['locale'] = locale
    __args__['onFailures'] = on_failures
    __args__['tag'] = tag
    __args__['timezone'] = timezone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorDateIndexName:IngestProcessorDateIndexName', __args__, opts=opts, typ=IngestProcessorDateIndexNameResult).value

    return AwaitableIngestProcessorDateIndexNameResult(
        date_formats=__ret__.date_formats,
        date_rounding=__ret__.date_rounding,
        description=__ret__.description,
        field=__ret__.field,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        index_name_format=__ret__.index_name_format,
        index_name_prefix=__ret__.index_name_prefix,
        json=__ret__.json,
        locale=__ret__.locale,
        on_failures=__ret__.on_failures,
        tag=__ret__.tag,
        timezone=__ret__.timezone)


@_utilities.lift_output_func(ingest_processor_date_index_name)
def ingest_processor_date_index_name_output(date_formats: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                            date_rounding: Optional[pulumi.Input[str]] = None,
                                            description: Optional[pulumi.Input[Optional[str]]] = None,
                                            field: Optional[pulumi.Input[str]] = None,
                                            if_: Optional[pulumi.Input[Optional[str]]] = None,
                                            ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                            index_name_format: Optional[pulumi.Input[Optional[str]]] = None,
                                            index_name_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                                            locale: Optional[pulumi.Input[Optional[str]]] = None,
                                            on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                            tag: Optional[pulumi.Input[Optional[str]]] = None,
                                            timezone: Optional[pulumi.Input[Optional[str]]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorDateIndexNameResult]:
    """
    The purpose of this processor is to point documents to the right time based index based on a date or timestamp field in a document by using the date math index name support.

    The processor sets the _index metadata field with a date math index name expression based on the provided index name prefix, a date or timestamp field in the documents being processed and the provided date rounding.

    First, this processor fetches the date or timestamp from a field in the document being processed. Optionally, date formatting can be configured on how the field’s value should be parsed into a date. Then this date, the provided index name prefix and the provided date rounding get formatted into a date math index name expression. Also here optionally date formatting can be specified on how the date should be formatted into a date math index name expression.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-index-name-processor.html

    ## Example Usage

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    date_index_name = elasticstack.ingest_processor_date_index_name(description="monthly date-time index naming",
        field="date1",
        index_name_prefix="my-index-",
        date_rounding="M")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[date_index_name.json])
    ```


    :param Sequence[str] date_formats: An array of the expected date formats for parsing dates / timestamps in the document being preprocessed.
    :param str date_rounding: How to round the date when formatting the date into the index name.
    :param str description: Description of the processor.
    :param str field: The field to get the date or timestamp from.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param str index_name_format: The format to be used when printing the parsed date into the index name.
    :param str index_name_prefix: A prefix of the index name to be prepended before the printed date.
    :param str locale: The locale to use when parsing the date from the document being preprocessed, relevant when parsing month names or week days.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str tag: Identifier for the processor.
    :param str timezone: The timezone to use when parsing the date and when date math index supports resolves expressions into concrete index names.
    """
    ...