from rest_framework import serializers
from reservations.models import User, Room, Reservation


class UserSerializer(serializers.HyperlinkedModelSerializer):
    """
    A User database model serializer
    """
    password = serializers.CharField(
        style={'input_type': 'password'}, write_only=True)

    class Meta:
        """
        Configuration values for the Serializer
        """
        model = User
        fields = ['url', 'name', 'email', 'password']

        extra_kwargs = {
            'email': {'write_only': True},
        }

    def create(self, validated_data):
        """Creates the user with specified password"""
        user = super().create(validated_data)
        user.set_password(validated_data['password'])
        user.save()
        return user

    def update(self, instance, validated_data):
        """Updates the user's password if they are patching it"""
        user = super().update(instance, validated_data)

        try:
            user.set_password(validated_data['password'])
            user.save()
        except KeyError:
            pass

        return user


class RoomSerializer(serializers.HyperlinkedModelSerializer):
    """A serializer for the Room database object"""
    class Meta:
        """
        Configuration values for the Serializer
        """
        model = Room
        fields = '__all__'


class ReservationSerializer(serializers.HyperlinkedModelSerializer):
    """A serializer for the Reservation database object"""
    # Disable no-member violations in this class because of the dynamic nature of the model
    # pylint: disable=no-member
    class Meta:
        """
        Configuration values for the Serializer
        """
        model = Reservation
        fields = '__all__'

    def create(self, validated_data):
        """Checks that a room isn't already reserved for that period"""
        if Reservation.objects.filter(date=validated_data['date']).exists():
            raise serializers.ValidationError(
                'Room already reserved for selected date'
            )

        return super().create(validated_data)


class StatusSerializer(serializers.Serializer):
    """A serializer for the virtual room status model"""
    is_reserved = serializers.BooleanField()
    date = serializers.IntegerField()
    room = serializers.CharField()
    user = serializers.CharField()

    def create(self, validated_data):
        """A dummy implementation"""

    def update(self, instance, validated_data):
        """A dummy implementation"""
