# The Prototype Settings package

This directory contains the project settings for the Django framework. This directory has the same name as the root directory, since the root one defines the project and the nested one defines its settings. All files here are required by Django to run the application and the names as to be exact so that Django can find them.

## Structure

- The [\_\_init\_\_.py](./__init__.py) file is required by Pylint to recognize this directory as a Python module, so it can recursively lint the files
- The [asgi.py](./asgi.py) and the [wsgi.py](./wsgi.py) files are required by Django to configure web servers that will be hosting the application
- The [settings.py](./settings.py) file contains the settings for the Django project
- The [urls.py](./urls.py) file contains the URL patterns for the Django project
