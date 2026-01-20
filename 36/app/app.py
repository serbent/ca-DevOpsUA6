from flask import Flask
from prometheus_client import Counter, Histogram, generate_latest
import time
import random

app = Flask(__name__)

REQUESTS = Counter('http_requests_total', 'Total HTTP requests')
REQUEST_LATENCY = Histogram('http_request_latency_seconds', 'HTTP request latency')

@app.route('/')
def hello():
    REQUESTS.inc()
    REQUEST_LATENCY.observe(random.random())
    return 'Hello, World!'

@app.route('/metrics')
def metrics():
    return generate_latest(), 200, {'Content-Type': 'text/plain'}

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)


