# from django.shortcuts import render

# Create your views here.

import os
import logging
import pathlib
import magic
import socket

from django.contrib.auth.models import User, Group
from django.core.files.storage import FileSystemStorage
from django.http import JsonResponse
from django.conf import settings

from rest_framework import viewsets
from rest_framework.decorators import api_view, renderer_classes
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.renderers import StaticHTMLRenderer

from replicaServer.api.serializers import UserSerializer, GroupSerializer

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
logger = logging.getLogger(__name__)


@api_view(['GET'])
@renderer_classes([StaticHTMLRenderer])
def index(request):

    htmlFile = os.path.join(BASE_DIR, 'statics/client/') + "index.html"
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

class IPAddr(APIView):
    def get(self, request, format=None):

        ipAddr = {}
        try: 
            ipAddr["host"] = socket.gethostname()
            ipAddr["ip"] = socket.gethostbyname(ipAddr["host"])
        except: 
            print("Unable to get Hostname and IP")

        return JsonResponse(data=ipAddr, status=200)

class FileList(APIView):

    def get(self, request, format=None):
        fileInfos = []
        for file in pathlib.Path(settings.MEDIA_ROOT).iterdir():
            if file.is_file:
                fileInfo = {}
                fileInfo["name"] = file.name
                fileInfo["type"] = magic.from_file( str(file) , mime=True)
                fileInfos.append(fileInfo)
                
        data = {
            "files": fileInfos
        }

        return JsonResponse(data=data, status=200)
    
    # Upload files
    def post(self, request, format=None):
        context = {}
        context['ok'] = True
        context['message'] = "Successfully uploaded the file"
        status = 200
        # return Response("Uploading files")
        try:
            uploaded_file = request.FILES['document']
            fs = FileSystemStorage()
            name = fs.save(uploaded_file.name, uploaded_file)
            context['url'] = fs.url(name)
            pass
        except Exception as e:
            context['ok'] = False
            status = 500
            context['message'] = 'Failed to upload the file'
            logger.error('Failed to upload file: ' + str(e.__traceback__))
            pass

        return JsonResponse(data=context, status=status)


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