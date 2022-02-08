from django.db import models


class Rule(models.Model):
    name = models.CharField(max_length=200)
    description = models.CharField(max_length=200)
    is_enabled = models.BooleanField()
    value = models.IntegerField()


class Site(models.Model):
    name = models.CharField(max_length=200)
    site_id = models.CharField(max_length=50)
    address = models.CharField(max_length=200)
    postal_code = models.CharField(max_length=10)
    description = models.CharField(max_length=200)
    capacity = models.IntegerField()
    should_ban = models.BooleanField()
