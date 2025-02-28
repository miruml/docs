from typing import TypeVar
from config import Config

T = TypeVar("T")

class Client:
    def __init__(
        self,
        sdk_key: str = None,
        client_id: str = None,
        ca_file: str = None,
        cert_file: str = None,
        key_file: str = None,
    ):
        pass

    @classmethod
    def from_sdk_key(cls, sdk_key: str, client_id: str):
        pass

    @classmethod
    def from_cert(cls, ca_file: str, cert_file: str, key_file: str):
        pass

    def init_config(
        self,
        slug: str,
        schema_file: str,
        data: T,
        development: bool = False,
        development_file: str | None = None,
    ) -> Config:
        pass