import requests
import box

# url = "https://jsonplaceholder.typicode.com/posts"
url = "http://pihole.lan/api.php"

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
    "Authorization": "Bearer YOUR_ACCESS_TOKEN",
    "Accept": "application/json"
}


response = requests.get(url, headers=headers)

box.boxprint(response.json())

box.boxprint(response.headers)

box.boxprint(response.text)

box.boxprint(response.status_code)