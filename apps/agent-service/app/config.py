import os
from dotenv import load_dotenv

load_dotenv()

class Settings:
    OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
    BACKEND_URL = os.getenv("BACKEND_URL", "http://localhost:8080")

settings = Settings()