# The prototype package

This directory contains the prototype server project. Everything in this directory is managed by the Django framework. File and folder names, and folder structures are predefined by Django and cannot be changed, as it expects to find them in the exact same place and with the exact same name.

## Structure

- The [reservations](./reservations) directory contains the main prototype API logic for the application. Views, serializers, models and routers are defined here.
- The [prototype](./prototype) directory contains the Django project settings and configurations.
- The [manage.py](./manage.py) file is the main entry point for the Django project. It is used to run the server, run tests, and perform other administrative tasks.
- The [DjangoPrototype.postman_collection.json](./DjangoPrototype.postman_collection.json) file is a Postman collection that contains all the tests for the prototype API.
