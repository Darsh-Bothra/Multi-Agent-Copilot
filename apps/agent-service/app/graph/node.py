from app.tools.expense_tool import get_balance, get_settlements

def parse_query(state):
    """ Parse the query to get the group id """
    query = state["query"]

    # simple logic (later: LLM)
    if "balance" in query:
        return {"group_id": "123", "action": "balances"}
    elif "settlement" in query:
        return {"group_id": "123", "action": "settlements"}

    return {"action": "unknown"}

def execute_tool(state):

    if state["action"] == "balances":
        result = get_balance.invoke({"group_id": state["group_id"]})

    elif state["action"] == "settlements":
        result = get_settlements.invoke({"group_id": state["group_id"]})

    else:
        result = {"error": "unknown action"}

    return {"result": result}


def format_response(state):
    return {"response": f"Result: {state['result']}"}