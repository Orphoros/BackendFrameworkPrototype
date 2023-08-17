from rest_framework.response import Response
from rest_framework.request import Request
from rest_framework.serializers import Serializer


class PatchModelMixin:
    """
    Update a model instance using PATCH.
    """

    def partial_update(self, request: Request, *_, **__) -> Response:
        """
        Updates a single destination object by its ID

        Args:
            request (Request): HTTP Request to process
            pk (str): Destination ID

        Returns:

            Response: A JSON object containing the updated destination object
        """
        partial = True
        instance = self.get_object()
        serializer = self.get_serializer(
            instance, data=request.data, partial=partial)
        serializer.is_valid(raise_exception=True)
        self.perform_update(serializer)

        if getattr(instance, '_prefetched_objects_cache', None):
            # If 'prefetch_related' has been applied to a queryset, we need to
            # forcibly invalidate the prefetch cache on the instance.
            instance._prefetched_objects_cache = {}

        return Response(serializer.data)

    def perform_update(self, serializer: Serializer):
        """
        Commit update in specified serializer

        Args:
            serializer (Serializer): Serializer in which to commit data
        """
        serializer.save()
