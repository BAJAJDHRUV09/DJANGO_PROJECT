#there are two types of views class-based and function-based
from django.http import HttpResponse

def homepage(request):
    print("Home page requested")
    return HttpResponse("This is home page.")