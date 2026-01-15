from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def hello():
    return 'Hello World!'

@app.route('/user/<name>')
def user(name):
    return f'Hello {name}'

@app.route('/login', methods=['GET', 'POST'])
def login():
    return render_template('login.html')

@app.route('/login_post', methods=['POST'])
def login_post():
    return '<h1>POST</h1>'

@app.route('/login_get', methods=['GET'])
def login_get():
    return '<h1>GET</h1>'

if __name__ == '__main__':
    app.run(debug=True)
