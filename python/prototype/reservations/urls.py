from django.urls import path, include
from rest_framework import routers
from . import views

router = routers.DefaultRouter(trailing_slash=False)
router.register(r'users', views.UserViewSet)
router.register(r'auth/login', views.LoginViewset, basename='auth/login')
router.register(r'auth/refresh', views.RefreshViewset, basename='auth/refresh')
router.register(r'rooms', views.RoomViewSet)
router.register(r'status', views.StatusViewSet, basename='status')
router.register(r'reservations', views.ReservationViewSet)

urlpatterns = [
    path('', include(router.urls)),
    path('api-auth/', include('rest_framework.urls', namespace='rest_framework'))
]
