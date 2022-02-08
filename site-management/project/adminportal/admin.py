from django.contrib import admin

from .models import Site
from .models import Rule

admin.site.register(Site)
admin.site.register(Rule)
