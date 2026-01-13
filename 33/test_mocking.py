import unittest
from unittest.mock import patch, Mock
from mocking import get_todo_item

class TestMocking(unittest.TestCase):
    @patch('mocking.requests.get')
    def test_get_todo_item(self, mock_get):
        mock_get.return_value.status_code = 200
        mock_get.return_value.json.return_value = {'title': 'Test Item'}
        self.assertEqual(get_todo_item(1), 'Test Item')
        mock_get.assert_called_once_with('https://jsonplaceholder.typicode.com/todos/1')

if __name__ == '__main__':
    unittest.main()