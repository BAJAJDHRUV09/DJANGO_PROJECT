from rest_framework import serializers
from .models import LearningModule

class LearningModuleSerializer(serializers.ModelSerializer):
    created_by = serializers.ReadOnlyField(source='created_by.username')

    class Meta:
        model = LearningModule
        fields = ['id', 'subject', 'grade', 'chapter', 'content', 'created_at', 'updated_at', 'created_by']
        read_only_fields = ['created_at', 'updated_at', 'created_by']