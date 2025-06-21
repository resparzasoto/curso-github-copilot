import unittest
from src.app import app

class GreetEndpointTestCase(unittest.TestCase):
    def setUp(self):
        self.app = app.test_client()
        self.app.testing = True

    def test_greet_default(self):
        response = self.app.get('/greet')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json, {'message': 'Hello World'})

    def test_greet_with_name(self):
        response = self.app.get('/greet?name=Alice')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json, {'message': 'Hello Alice'})

if __name__ == '__main__':
    unittest.main()
