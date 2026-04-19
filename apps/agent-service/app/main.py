from fastapi import FastAPI
from app.graph.graph_builder import build_graph

app = FastAPI()

graph = build_graph()


@app.post("/chat")
def chat(query: str):

    result = graph.invoke({
        "query": query
    })

    return result