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
    'IngestProcessorDateResult',
    'AwaitableIngestProcessorDateResult',
    'ingest_processor_date',
    'ingest_processor_date_output',
]

@pulumi.output_type
class IngestProcessorDateResult:
    """
    A collection of values returned by IngestProcessorDate.
    """
    def __init__(__self__, description=None, field=None, formats=None, id=None, if_=None, ignore_failure=None, json=None, locale=None, on_failures=None, output_format=None, tag=None, target_field=None, timezone=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if field and not isinstance(field, str):
            raise TypeError("Expected argument 'field' to be a str")
        pulumi.set(__self__, "field", field)
        if formats and not isinstance(formats, list):
            raise TypeError("Expected argument 'formats' to be a list")
        pulumi.set(__self__, "formats", formats)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if if_ and not isinstance(if_, str):
            raise TypeError("Expected argument 'if_' to be a str")
        pulumi.set(__self__, "if_", if_)
        if ignore_failure and not isinstance(ignore_failure, bool):
            raise TypeError("Expected argument 'ignore_failure' to be a bool")
        pulumi.set(__self__, "ignore_failure", ignore_failure)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if locale and not isinstance(locale, str):
            raise TypeError("Expected argument 'locale' to be a str")
        pulumi.set(__self__, "locale", locale)
        if on_failures and not isinstance(on_failures, list):
            raise TypeError("Expected argument 'on_failures' to be a list")
        pulumi.set(__self__, "on_failures", on_failures)
        if output_format and not isinstance(output_format, str):
            raise TypeError("Expected argument 'output_format' to be a str")
        pulumi.set(__self__, "output_format", output_format)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if target_field and not isinstance(target_field, str):
            raise TypeError("Expected argument 'target_field' to be a str")
        pulumi.set(__self__, "target_field", target_field)
        if timezone and not isinstance(timezone, str):
            raise TypeError("Expected argument 'timezone' to be a str")
        pulumi.set(__self__, "timezone", timezone)

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
        The field to get the date from.
        """
        return pulumi.get(self, "field")

    @property
    @pulumi.getter
    def formats(self) -> Sequence[str]:
        """
        An array of the expected date formats.
        """
        return pulumi.get(self, "formats")

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
        The locale to use when parsing the date, relevant when parsing month names or week days.
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
    @pulumi.getter(name="outputFormat")
    def output_format(self) -> Optional[str]:
        """
        The format to use when writing the date to `target_field`.
        """
        return pulumi.get(self, "output_format")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        Identifier for the processor.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="targetField")
    def target_field(self) -> Optional[str]:
        """
        The field that will hold the parsed date.
        """
        return pulumi.get(self, "target_field")

    @property
    @pulumi.getter
    def timezone(self) -> Optional[str]:
        """
        The timezone to use when parsing the date.
        """
        return pulumi.get(self, "timezone")


class AwaitableIngestProcessorDateResult(IngestProcessorDateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IngestProcessorDateResult(
            description=self.description,
            field=self.field,
            formats=self.formats,
            id=self.id,
            if_=self.if_,
            ignore_failure=self.ignore_failure,
            json=self.json,
            locale=self.locale,
            on_failures=self.on_failures,
            output_format=self.output_format,
            tag=self.tag,
            target_field=self.target_field,
            timezone=self.timezone)


def ingest_processor_date(description: Optional[str] = None,
                          field: Optional[str] = None,
                          formats: Optional[Sequence[str]] = None,
                          if_: Optional[str] = None,
                          ignore_failure: Optional[bool] = None,
                          locale: Optional[str] = None,
                          on_failures: Optional[Sequence[str]] = None,
                          output_format: Optional[str] = None,
                          tag: Optional[str] = None,
                          target_field: Optional[str] = None,
                          timezone: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIngestProcessorDateResult:
    """
    Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document.
    By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `target_field` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html

    ## Example Usage

    Here is an example that adds the parsed date to the `timestamp` field based on the `initial_date` field:

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    date = elasticstack.ingest_processor_date(field="initial_date",
        target_field="timestamp",
        formats=["dd/MM/yyyy HH:mm:ss"],
        timezone="Europe/Amsterdam")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[date.json])
    ```


    :param str description: Description of the processor.
    :param str field: The field to get the date from.
    :param Sequence[str] formats: An array of the expected date formats.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param str locale: The locale to use when parsing the date, relevant when parsing month names or week days.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str output_format: The format to use when writing the date to `target_field`.
    :param str tag: Identifier for the processor.
    :param str target_field: The field that will hold the parsed date.
    :param str timezone: The timezone to use when parsing the date.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['field'] = field
    __args__['formats'] = formats
    __args__['if'] = if_
    __args__['ignoreFailure'] = ignore_failure
    __args__['locale'] = locale
    __args__['onFailures'] = on_failures
    __args__['outputFormat'] = output_format
    __args__['tag'] = tag
    __args__['targetField'] = target_field
    __args__['timezone'] = timezone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('elasticstack:index/ingestProcessorDate:IngestProcessorDate', __args__, opts=opts, typ=IngestProcessorDateResult).value

    return AwaitableIngestProcessorDateResult(
        description=__ret__.description,
        field=__ret__.field,
        formats=__ret__.formats,
        id=__ret__.id,
        if_=__ret__.if_,
        ignore_failure=__ret__.ignore_failure,
        json=__ret__.json,
        locale=__ret__.locale,
        on_failures=__ret__.on_failures,
        output_format=__ret__.output_format,
        tag=__ret__.tag,
        target_field=__ret__.target_field,
        timezone=__ret__.timezone)


@_utilities.lift_output_func(ingest_processor_date)
def ingest_processor_date_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                                 field: Optional[pulumi.Input[str]] = None,
                                 formats: Optional[pulumi.Input[Sequence[str]]] = None,
                                 if_: Optional[pulumi.Input[Optional[str]]] = None,
                                 ignore_failure: Optional[pulumi.Input[Optional[bool]]] = None,
                                 locale: Optional[pulumi.Input[Optional[str]]] = None,
                                 on_failures: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 output_format: Optional[pulumi.Input[Optional[str]]] = None,
                                 tag: Optional[pulumi.Input[Optional[str]]] = None,
                                 target_field: Optional[pulumi.Input[Optional[str]]] = None,
                                 timezone: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IngestProcessorDateResult]:
    """
    Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document.
    By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `target_field` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.

    See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html

    ## Example Usage

    Here is an example that adds the parsed date to the `timestamp` field based on the `initial_date` field:

    ```python
    import pulumi
    import pulumi_elasticstack as elasticstack

    date = elasticstack.ingest_processor_date(field="initial_date",
        target_field="timestamp",
        formats=["dd/MM/yyyy HH:mm:ss"],
        timezone="Europe/Amsterdam")
    my_ingest_pipeline = elasticstack.IngestPipeline("myIngestPipeline", processors=[date.json])
    ```


    :param str description: Description of the processor.
    :param str field: The field to get the date from.
    :param Sequence[str] formats: An array of the expected date formats.
    :param str if_: Conditionally execute the processor
    :param bool ignore_failure: Ignore failures for the processor.
    :param str locale: The locale to use when parsing the date, relevant when parsing month names or week days.
    :param Sequence[str] on_failures: Handle failures for the processor.
    :param str output_format: The format to use when writing the date to `target_field`.
    :param str tag: Identifier for the processor.
    :param str target_field: The field that will hold the parsed date.
    :param str timezone: The timezone to use when parsing the date.
    """
    ...
