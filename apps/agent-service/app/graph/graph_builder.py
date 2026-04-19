from langgraph.graph import StateGraph
from app.graph.state import State
from app.graph.node import parse_query, execute_tool, format_response



def build_graph():
    graph = StateGraph(State)

    graph.add_node("parse", parse_query)
    graph.add_node("execute", execute_tool)
    graph.add_node("format", format_response)

    graph.set_entry_point("parse")

    graph.add_edge("parse", "execute")
    graph.add_edge("execute", "format")

    return graph.compile()