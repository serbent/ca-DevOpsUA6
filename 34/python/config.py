import os
bind = os.environ.get('BIND', '0.0.0.0:5000')
workers = os.environ.get('WORKERS', 3)
