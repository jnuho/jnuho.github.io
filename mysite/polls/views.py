from django.http import HttpResponse

# Create your views here.
def index(request):
    return HttpResponse("Hello, world. You're at the polls index.")

def test(request):
    return HttpResponse("Hello, world. You're at the test index.")
