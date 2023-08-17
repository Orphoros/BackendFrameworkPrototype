# Django reservation API prototype

## Getting Started

> Important: Depending on the OS, you may need to run `python3` instead of `python`

To get started with the project, make sure you meet all the prerequisites and follow the instructions below.

### Prerequisites

Make sure you have all the following prerequisites installed on your development machine:

- Install [Python](https://www.python.org/downloads/) 3.10 or above
- Install [MySQL](https://dev.mysql.com/doc/mysql-getting-started/en/#mysql-getting-started-installing) 8 or above, and make sure that the database server is running

## Setup

The following steps are required to run the server locally after the prerequisites are installed.

### Create a virtual environment

1. Create the virtual environment:

```sh
python -m venv env_name
```

2. Activate the environment:

> The command depends on the OS and terminal. Run this command every time before running the server

```sh
source env_name/bin/activate
```

Or

```sh
env_name/Scripts/Activate.ps1
```

Or

```sh
env_name/Scripts/activate.bat
```

### Install dependencies

```bash
cd prototype
pip install -r requirements.txt
```

### MySQL

- Create a `reservation` schema in MySQL
- Create an `api_user` user in MySQL with the password `Test12345!` and grant all privileges on the `reservation` schema

The backend expects the MySQL server to be running on `localhost:3306` with the `api_user` user and the `Test12345!` password.

After MySQL is set up, run the following commands to create the database tables:

```bash
cd prototype
python manage.py makemigrations
python manage.py migrate
```

After the migrations complete, create a superuser account for Django management:

```sh
python manage.py createsuperuser
```

## Running the server

```bash
python manage.py runserver
```

## Running the tests

### Internal endpoint authentication tests

`python manage.py test reservations`

### External endpoint tests

Import and run the Postman collection `prototype/DjangoPrototype.postman_collection.json`
