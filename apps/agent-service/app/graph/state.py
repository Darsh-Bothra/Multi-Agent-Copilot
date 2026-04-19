from typing import TypedDict, Optional

class State(TypedDict):
    query: str
    group_id: Optional[str]
    result: Optional[dict]
    response: Optional[str]