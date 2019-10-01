# from django.shortcuts import render

# Create your views here.

from django.contrib.auth.models import User, Group
from rest_framework import viewsets
# from rest_framework.decorators import api_view
from rest_framework.views import APIView
from rest_framework.response import Response
from replicaServer.api.serializers import UserSerializer, GroupSerializer

# @api_view(['GET'])
# def index(request):
#     return Response("API Works")

class Index(APIView):
    def get(self, request, format=None):
        return Response("API Works")

class File(APIView):
    def get(self, request, format=None):
        return Response("File Works")

class UserViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows users to be viewed or edited.
    """
    queryset = User.objects.all().order_by('-date_joined')
    serializer_class = UserSerializer


class GroupViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows groups to be viewed or edited.
    """
    queryset = Group.objects.all()
    serializer_class = GroupSerializer