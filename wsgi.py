#!/usr/bin/python3
"""
imports Flask instance for gunicorn configurations
gunicorn --bind 127.0.0.1:8001 wsgi:cod.app
"""

web = __import__('views.cod', globals(), locals(), ['*'])

if __name__ == "__main__":
    """runs the main flask app"""
    web.app.run()
