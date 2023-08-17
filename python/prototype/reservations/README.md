# The Reservations package

This directory contains the main API logic for the prototype Django application. Server models, views, serializers, and routers are defined here.

## Structure

- The [\_\_init\_\_.py](./__init__.py) file is required by Pylint to recognize this directory as a Python module, so it can recursively lint the files
- The [admin.py](./admin.py) file contains the admin site (admin database management) configuration for the application
- The [apps.py](./apps.py) file contains the application configuration for the application
- The [models.py](./models.py) file contains the database models for the application
- The [serializers.py](./serializers.py) file contains the serializers (JSON serializers, etc.) for the application models
- The [tests.py](./tests.py) file contains the unit tests for the application
- The [urls.py](./urls.py) file contains the URL patterns (URL access points for the views) for the application
- The [views.py](./views.py) file contains the views (API endpoints) for the application
- The [migrations](./migrations) directory contains the database migrations for the application. This directory is automatically generated and should not be modified manually!
- The [permissions.py](./permissions.py) file contains the middleware configuration for the application.
- The [mixins.py](./mixins.py) file contains custom overrides for the supported API methods
