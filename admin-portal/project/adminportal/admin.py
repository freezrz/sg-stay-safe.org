from django.contrib import admin

from .models import Site
from .models import Rule
from .models import SafeAmbassador
from .models import Region
from .models import User

admin.site.register(Site)
admin.site.register(Rule)
admin.site.register(SafeAmbassador)
admin.site.register(Region)
admin.site.register(User)
