#!/usr/bin/python3
'''
Frontend web application for containers on demand.
'''
from flask import Flask, render_template, url_for
from flask_cors import CORS, cross_origin
import os


application = Flask(__name__)
application.url_map.strict_slashes = False
host, port = '0.0.0.0', 5000
cors = CORS(application, resources={r'/api/v1/*': {'origins': '*'}})


@application.route('/')
def hello():
    return render_template('index.html')

if __name__ == '__main__':
    application.run(host=host, port=port, debug=True)
