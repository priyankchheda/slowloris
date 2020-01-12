from django.contrib import admin
from django.urls import path, re_path
from django.conf import settings
from django.conf.urls.static import static
from adoptions import views

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', views.home, name='home'),
    re_path(r'adoptions/(\d+)/', views.pet_detail, name='pet_detail'),
]

urlpatterns += static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)