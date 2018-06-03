#!/usr/bin/python3
'''
Frontend web application for containers on demand.
'''
from flask import Flask, render_template, url_for
from flask_cors import CORS, cross_origin
import os


app = Flask(__name__)
app.url_map.strict_slashes = False
host, port = '0.0.0.0', 5000
cors = CORS(app, resources={r'/api/v1/*': {'origins': '*'}})


@app.route('/')
def hello():
    return render_template('index.html')

if __name__ == '__main__':
    app.run(host=host, port=port, debug=True)
