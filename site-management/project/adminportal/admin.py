from django.contrib import admin

from .models import Site
from .models import Rule
from .models import SafeAmbassador
from .models import Region

admin.site.register(Site)
admin.site.register(Rule)
admin.site.register(SafeAmbassador)
admin.site.register(Region)
