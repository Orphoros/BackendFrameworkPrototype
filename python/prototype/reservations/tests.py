import json
from django.test import TestCase


class TestProtectedEndpoint(TestCase):
    """Validate that endpoint protection is working"""

    def call_status_anonymously(self):
        """Call a protected endpoint without auth. Should fail."""
        response = self.client.get('/api/v1/status', follow=True)
        self.assertEqual(response.status_code, 401)

    def call_registration(self):
        """Register user and save ID. Should succeed and return data"""
        user = {
            'name': 'test',
            'email': 'test@local.host',
            'password': 'Testing1234'
        }
        response = self.client.post(
            '/api/v1/users', json.dumps(user), content_type='application/json')
        self.assertContains(response, 'test', status_code=201)

    def call_login(self):
        """Log in and save JWT. Should succeed and return token"""
        user = {
            'email': 'test@local.host',
            'password': 'Testing1234'
        }
        response = self.client.post(
            '/api/v1/auth/login', json.dumps(user), content_type='application/json')
        self.assertContains(response, 'access')
        self.access = response.json()['access']

    def call_status_authenticated(self):
        """Call admin-only endpoint while logged in as a non-admin user. Should fail"""
        auth_headers = {'HTTP_AUTHORIZATION': f'Bearer {self.access}'}
        response = self.client.get(
            '/api/v1/status', follow=True, **auth_headers)
        self.assertEqual(response.status_code, 200)

    def test_api(self):
        """Run all tests in sequence"""
        self.call_status_anonymously()
        self.call_registration()
        self.call_login()
        self.call_status_authenticated()
