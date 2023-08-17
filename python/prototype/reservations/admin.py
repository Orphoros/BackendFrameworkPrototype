from django.contrib import admin
from reservations.models import User, Room, Reservation

# Register your models here.
admin.site.register(User)
admin.site.register(Room)
admin.site.register(Reservation)
