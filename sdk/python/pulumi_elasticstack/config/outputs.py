# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'Elasticsearch',
]

@pulumi.output_type
class Elasticsearch(dict):
    def __init__(__self__, *,
                 api_key: Optional[str] = None,
                 ca_data: Optional[str] = None,
                 ca_file: Optional[str] = None,
                 cert_data: Optional[str] = None,
                 cert_file: Optional[str] = None,
                 endpoints: Optional[Sequence[str]] = None,
                 insecure: Optional[bool] = None,
                 key_data: Optional[str] = None,
                 key_file: Optional[str] = None,
                 password: Optional[str] = None,
                 username: Optional[str] = None):
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if ca_data is not None:
            pulumi.set(__self__, "ca_data", ca_data)
        if ca_file is not None:
            pulumi.set(__self__, "ca_file", ca_file)
        if cert_data is not None:
            pulumi.set(__self__, "cert_data", cert_data)
        if cert_file is not None:
            pulumi.set(__self__, "cert_file", cert_file)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if key_data is not None:
            pulumi.set(__self__, "key_data", key_data)
        if key_file is not None:
            pulumi.set(__self__, "key_file", key_file)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[str]:
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="caData")
    def ca_data(self) -> Optional[str]:
        return pulumi.get(self, "ca_data")

    @property
    @pulumi.getter(name="caFile")
    def ca_file(self) -> Optional[str]:
        return pulumi.get(self, "ca_file")

    @property
    @pulumi.getter(name="certData")
    def cert_data(self) -> Optional[str]:
        return pulumi.get(self, "cert_data")

    @property
    @pulumi.getter(name="certFile")
    def cert_file(self) -> Optional[str]:
        return pulumi.get(self, "cert_file")

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def insecure(self) -> Optional[bool]:
        return pulumi.get(self, "insecure")

    @property
    @pulumi.getter(name="keyData")
    def key_data(self) -> Optional[str]:
        return pulumi.get(self, "key_data")

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> Optional[str]:
        return pulumi.get(self, "key_file")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")


