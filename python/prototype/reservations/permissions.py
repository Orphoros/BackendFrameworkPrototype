import re

from rest_framework import permissions
from rest_framework.request import Request


class UpdateProfile(permissions.BasePermission):
    """Check if the user updating a profile owns it"""

    def has_object_permission(self, request: Request, _, obj):
        """
        Check if the request can continue with processing

        Args:
            request (Request): Request to check
            obj: Object being requested

        Returns:
            bool: True if request can continue, False otherwise
        """
        if request.method in permissions.SAFE_METHODS:
            return True

        return obj.id == request.user.id


class UpdateReservation(permissions.BasePermission):
    """Check if the user updating a reservation owns it"""

    def has_object_permission(self, request: Request, _, obj):
        """
        Check if the request can continue with processing

        Args:
            request (Request): Request to check
            obj: Object being requested

        Returns:
            bool: True if request can continue, False otherwise
        """
        if request.method in permissions.SAFE_METHODS:
            return True

        return obj.user.id == request.user.id


class CreateForSelf(permissions.BasePermission):
    """Check if the user is creating a reservation for themselves"""

    def has_permission(self, request: Request, _):
        """
        Check if the request can continue with processing

        Args:
            request (Request): Request to check

        Returns:
            bool: True if request can continue, False otherwise
        """
        if request.method in permissions.SAFE_METHODS or 'user' not in request.data:
            return True

        # Pulls the user's UUID from their hyperlink request entry
        target_user = re.search(
            r'[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}', request.data['user']).group(0)
        return target_user == str(request.user.id)
