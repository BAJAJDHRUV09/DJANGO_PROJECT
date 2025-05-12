from django.db import models
from django.contrib.auth.models import User

# Create your models here.
class LearningModule(models.Model):
    subject = models.CharField(max_length=80)
    grade = models.CharField(max_length=20)
    chapter = models.CharField(max_length=100)
    content = models.JSONField()
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    created_by = models.ForeignKey(User, on_delete=models.CASCADE)

    class Meta:
        ordering = ['-created_at']

    def __str__(self):
        return f"{self.subject} - Grade {self.grade} - Chapter {self.chapter}"



    
