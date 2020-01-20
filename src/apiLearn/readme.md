```python

import requests

url = "localhost:3000/users"

payload = "{\"name\":\"tom\",\"email\":\"tom@163.com\",\"password\":\"123456\"}"
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data = payload)

print(response.text.encode('utf8'))

```