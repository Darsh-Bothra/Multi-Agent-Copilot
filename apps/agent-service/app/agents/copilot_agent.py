from langchain.agents import initialize_agent, AgentType

from app.services.llm_service import get_llm
from app.tools.expense_tool import get_balance, get_settlements

def get_copilot_agent():
    llm = get_llm()
    tools = [get_balance, get_settlements]
    agent = initialize_agent(
        tools,
        llm,
        agent=AgentType.ZERO_SHOT_REACT_DESCRIPTION,
        verbose=True,
    )
    return agent