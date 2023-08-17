import uuid

from django.contrib.auth.models import (AbstractBaseUser, BaseUserManager,
                                        PermissionsMixin)
from django.db import models


class UserManager(BaseUserManager):
    """
    A Django authentication user account manager
    """

    def create_user(self, name: str, email: str, password: str):
        """
        Create a user

        Args:
            name (str): Name of the user
            email (str): Email of the user
            password (str): Password of the user
        """

        if not email:
            raise ValueError('Users must have an email address')

        user = self.model(
            email=self.normalize_email(email),
            name=name
        )

        user.set_password(password)
        user.save(using=self._db)
        return user

    def create_superuser(self, name: str, email: str, password: str):
        """
        Create a super user

        Args:
            name (str): Name of the user
            email (str): Email of the user
            password (str): Password of the user
        """
        user = self.create_user(name, email, password)

        user.is_superuser = True
        user.is_staff = True

        user.save(using=self._db)
        return user


class User(AbstractBaseUser, PermissionsMixin):
    """
    A user database object model
    """
    name = models.TextField()
    email = models.EmailField(unique=True)
    id = models.UUIDField(
        primary_key=True, default=uuid.uuid4, editable=False, max_length=36)
    password = models.TextField()
    is_active = models.BooleanField(default=True)
    is_staff = models.BooleanField(default=False)
    objects = UserManager()

    USERNAME_FIELD = 'email'
    REQUIRED_FIELDS = ['name']

    def get_full_name(self) -> str:
        """A method to retrieve the user full name. Used by the Django Admin panel"""
        return self.name

    def get_short_name(self) -> str:
        """A method to retrieve the user short name. Used by the Django Admin panel"""
        return self.name

    # Disable str violation because why even?
    # pylint: disable=invalid-str-returned
    def __str__(self) -> str:
        """A method to retrieve the user email. Used by the Django Admin panel"""
        return self.email


class Room(models.Model):
    """
    A model that represents a room, reservable by a user
    """
    floor = models.PositiveIntegerField()
    door_number = models.PositiveIntegerField(primary_key=True)
    description = models.TextField()


class Reservation(models.Model):
    """
    A model that represents a reservation for a specific room by a specific user on a specific date
    """
    id = models.UUIDField(
        primary_key=True, default=uuid.uuid4, editable=False, max_length=36)
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    room = models.ForeignKey(Room, on_delete=models.CASCADE)
    date = models.BigIntegerField()


class RoomStatus:
    """A virtual model that represents the status of a room. This model is not committed to the database"""

    def __init__(self, is_reserved: bool, room: str, date: int | None = None, user: str | None = None) -> None:
        self.is_reserved = is_reserved
        self.date = date
        self.room = room
        self.user = user
