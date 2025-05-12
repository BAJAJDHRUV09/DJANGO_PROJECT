from django.test import TestCase
from django.contrib.auth.models import User
from rest_framework.test import APIClient
from rest_framework import status
from .models import LearningModule
import json

class LearningModuleTests(TestCase):
    def setUp(self):
        self.client = APIClient()
        self.user = User.objects.create_user(username='MYSELFTESTUSER', password='testing987')
        self.client.force_authenticate(user=self.user)
        
        self.module_data = {
            'subject': 'MTH111',
            'grade': 'BTECH YEAR 1',
            'chapter': 'SINGLE VARIABLE CALCULUS',
            'content': {'sections': ['Introduction', 'Basic Concepts']}
        }
        self.module = LearningModule.objects.create(
            created_by=self.user,
            **self.module_data
        )
        self.base_url = '/api/v1/modules/'

    def test_create_module(self):
        """Test creating a new learning module (happy path)"""
        new_module_data = {
            'subject': 'Science',
            'grade': '9th',
            'chapter': 'Physics',
            'content': {'sections': ['Motion', 'Forces']}
        }
        response = self.client.post(self.base_url, new_module_data, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(LearningModule.objects.count(), 2)
        self.assertEqual(response.data['subject'], 'Science')

    def test_get_module_list(self):
        response = self.client.get(self.base_url)
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(len(response.data), 1)

    def test_get_module_detail(self):
        response = self.client.get(f'{self.base_url}{self.module.id}/')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(response.data['subject'], 'Mathematics')

    def test_unauthorized_create(self):
        self.client.force_authenticate(user=None)
        response = self.client.post(self.base_url, self.module_data, format='json')
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_update_module(self):
        update_data = {
            'subject': 'Updated Math',
            'grade': '11th',
            'chapter': 'Calculus',
            'content': {'sections': ['Derivatives', 'Integrals']}
        }
        response = self.client.put(f'{self.base_url}{self.module.id}/', update_data, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.module.refresh_from_db()
        self.assertEqual(self.module.subject, 'Updated Math')
