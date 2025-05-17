from typing import TypeVar

T = TypeVar("T")

class Config:
    def get(
        self,
        obj: T | None,
        default: T | None = None,
    ) -> T:
        pass

    def write_to_file(self, dest_file: str):
        pass

    def set(self, obj: T) -> T:
        pass

    def data(self) -> T:
        pass