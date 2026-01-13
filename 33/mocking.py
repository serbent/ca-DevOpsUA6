import requests
def get_todo_item(item_id):
    # response = requests.get(f"http://localhost:8000/todos/{item_id}")
    response = requests.get(f"https://jsonplaceholder.typicode.com/todos/{item_id}")
    response.raise_for_status()
    if response.status_code == 200:
        return response.json()["title"]  
    return None

# print(get_todo_item(1))
