# from django.shortcuts import render

# Create your views here.

import os

from django.contrib.auth.models import User, Group
from rest_framework import viewsets
from rest_framework.decorators import api_view, renderer_classes
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.renderers import StaticHTMLRenderer
from replicaServer.api.serializers import UserSerializer, GroupSerializer

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

@api_view(['GET'])
@renderer_classes([StaticHTMLRenderer])
def index(request):

    htmlFile = os.path.join(BASE_DIR, 'statics/') + "index.html"
    # data = '<html><body><h1>Hello, world</h1></body></html>'
    html = open(htmlFile)
    data = html.read()
    return Response(data)

# class Index(APIView):
#     def get(self, request, format=None):
#         return Response("Index Works")

class API(APIView):
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