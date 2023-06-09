from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MediaRequest(_message.Message):
    __slots__ = ["contentType", "data"]
    CONTENTTYPE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    contentType: str
    data: bytes
    def __init__(self, data: _Optional[bytes] = ..., contentType: _Optional[str] = ...) -> None: ...

class MediaResponse(_message.Message):
    __slots__ = ["contentType", "data", "description"]
    CONTENTTYPE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    contentType: str
    data: bytes
    description: str
    def __init__(self, data: _Optional[bytes] = ..., contentType: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class TellRequest(_message.Message):
    __slots__ = ["mediaTellReqArr"]
    MEDIATELLREQARR_FIELD_NUMBER: _ClassVar[int]
    mediaTellReqArr: _containers.RepeatedCompositeFieldContainer[MediaRequest]
    def __init__(self, mediaTellReqArr: _Optional[_Iterable[_Union[MediaRequest, _Mapping]]] = ...) -> None: ...

class TellResponse(_message.Message):
    __slots__ = ["descriptionAll", "mediaTellRespArr"]
    DESCRIPTIONALL_FIELD_NUMBER: _ClassVar[int]
    MEDIATELLRESPARR_FIELD_NUMBER: _ClassVar[int]
    descriptionAll: str
    mediaTellRespArr: _containers.RepeatedCompositeFieldContainer[MediaResponse]
    def __init__(self, mediaTellRespArr: _Optional[_Iterable[_Union[MediaResponse, _Mapping]]] = ..., descriptionAll: _Optional[str] = ...) -> None: ...
