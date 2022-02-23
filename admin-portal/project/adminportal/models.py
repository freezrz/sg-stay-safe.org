from django.db import models


class Rule(models.Model):
    name = models.CharField(max_length=200)
    description = models.CharField(max_length=200)
    is_enabled = models.BooleanField()
    value = models.IntegerField()

    def __str__(self):
        return self.name

class Region(models.Model):
    name = models.CharField(max_length=50)
    description = models.CharField(max_length=200)

    def __str__(self):
        return self.name

class Site(models.Model):
    name = models.CharField(max_length=200)
    site_id = models.CharField(max_length=50)
    owner = models.CharField(max_length=50)
    phone = models.CharField(max_length=50)
    email = models.CharField(max_length=50)
    address = models.CharField(max_length=200)
    postal_code = models.CharField(max_length=10)
    description = models.CharField(max_length=200)
    capacity = models.IntegerField()
    should_ban = models.BooleanField()
    region = models.ForeignKey(Region, on_delete=models.DO_NOTHING, default=None, blank=True)

    def __str__(self):
        return self.name

class SafeAmbassador(models.Model):
    name = models.CharField(max_length=200)
    phone = models.CharField(max_length=50)
    email = models.CharField(max_length=50)
    region = models.ForeignKey(Region, on_delete=models.DO_NOTHING, default=None, blank=True)
    description = models.CharField(max_length=200)

    def __str__(self):
        return self.name
