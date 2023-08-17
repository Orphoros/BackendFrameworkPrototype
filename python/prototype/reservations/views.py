import datetime
import time

from reservations.mixins import PatchModelMixin
from reservations.models import Reservation, Room, RoomStatus, User
from reservations.permissions import (CreateForSelf, UpdateProfile,
                                      UpdateReservation)
from reservations.serializer import (ReservationSerializer, RoomSerializer,
                                     StatusSerializer, UserSerializer)
from rest_framework import mixins, viewsets
from rest_framework.permissions import IsAuthenticated
from rest_framework_simplejwt.authentication import JWTAuthentication
from rest_framework_simplejwt.views import (TokenObtainPairView,
                                            TokenRefreshView)

# Create your views here.


class UserViewSet(mixins.RetrieveModelMixin,
                  mixins.ListModelMixin,
                  mixins.CreateModelMixin,
                  PatchModelMixin,
                  mixins.DestroyModelMixin,
                  viewsets.GenericViewSet
                  ):
    """CRUD endpoint for Users

        Checks if the user can update their own profile when running manipulating operations
    """
    queryset = User.objects.all()
    serializer_class = UserSerializer
    authentication_classes = (JWTAuthentication,)
    permission_classes = (UpdateProfile,)


class LoginViewset(viewsets.ViewSet):
    """POST endpoint for obtaining a JWT via loggin in"""

    def create(self, request):
        """Logs in a user by returning a JWT"""
        return TokenObtainPairView().as_view()(request=request._request)


class RefreshViewset(viewsets.ViewSet):
    """POST endpoint for requesting a new JWT using a refresh token when the original expires"""

    def create(self, request):
        """Refreshes a JWT access token using a refresh token"""
        return TokenRefreshView().as_view()(request=request._request)


# Disable no-member violations in this class because of the dynamic nature of the model
# pylint: disable=no-member
class RoomViewSet(viewsets.ReadOnlyModelViewSet):
    """GET endpoint for Rooms

        Allows querying for all free rooms on a given date as Unix Epoch
    """

    queryset = Room.objects.all()
    serializer_class = RoomSerializer

    def get_queryset(self):
        """GET all rooms"""
        free = self.request.query_params.get('free', None)
        if free:
            return Room.objects.exclude(door_number__in=Reservation.objects.filter(date=free).values_list('room'))
        return super().get_queryset()


class ReservationViewSet(mixins.RetrieveModelMixin,
                         mixins.ListModelMixin,
                         mixins.CreateModelMixin,
                         PatchModelMixin,
                         mixins.DestroyModelMixin,
                         viewsets.GenericViewSet
                         ):
    """CRUD endpoint for Reservations

        Checks if the user is manipulating reservations that belong to them or that they are creating reservations for themselves
    """
    queryset = Reservation.objects.all()
    serializer_class = ReservationSerializer
    authentication_classes = (JWTAuthentication,)
    permission_classes = (UpdateReservation, CreateForSelf)


class StatusViewSet(mixins.ListModelMixin, viewsets.GenericViewSet):
    """GET endpoint for the status of Rooms in a selected period

        Requires the user be authenticated with JWT. 

        Allows querying for room status within a given period as Unix Epoch. 

        The query parameter is added to today's midnight Unix Epoch to display reservations until end of the period
    """
    serializer_class = StatusSerializer
    authentication_classes = (JWTAuthentication,)
    permission_classes = (IsAuthenticated,)

    def get_queryset(self):
        """GET status for all rooms"""
        response = []
        period = self.request.query_params.get('period', None)
        rooms = Room.objects.all().iterator()
        reservations = None
        if period:
            midnight_datetime = datetime.datetime.combine(
                datetime.date.today(), datetime.datetime.min.time())
            midnight_time = int(time.mktime(midnight_datetime.timetuple()))
            end_time = midnight_time+int(period)
            reservations = Reservation.objects.filter(
                date__gte=midnight_time).filter(date__lte=end_time).iterator()
        else:
            reservations = Reservation.objects.all().iterator()

        for room in rooms:
            reserved = False
            for reservation in reservations:
                if reservation.room.pk == room.pk:
                    response.append(RoomStatus(
                        True, room.pk, reservation.date, reservation.user.pk))
                    reserved = True
            if not reserved:
                response.append(RoomStatus(False, room.pk))

        return response
