from langchain.chat_models import ChatOpenAI
from app.config import settings

def get_llm():
    return ChatOpenAI(
        model="gpt-4o-mini",
        temperature=0,
        api_key=settings.OPENAI_API_KEY
    )