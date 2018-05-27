#!/usr/bin/python3
'''
Frontend web application for containers on demand.
'''
from flask import Flask, render_template, url_for


app = Flask(__name__)
app.url_map.strict_slashes = False
host, port = '0.0.0.0', 5000


@app.route('/')
def hello():
    return render_template('index.html')

if __name__ == '__main__':
    app.run(host=host, port=port)
