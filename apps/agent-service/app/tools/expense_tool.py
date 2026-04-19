import requests
from app.config import settings
from langchain.tools import tool

@tool 
def get_balance(group_id: str):
    """ Get balance of a group """
    url = f"{settings.BACKEND_URL}/groups/{group_id}/balance"
    response = requests.get(url)
    return response.json()

@tool
def get_settlements(group_id: str):
    """ Get settlements of a group """
    url = f"{settings.BACKEND_URL}/groups/{group_id}/settlements"
    response = requests.get(url)
    return response.json()