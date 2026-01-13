import unittest
from addition import addition

class TestAddition(unittest.TestCase):
    def test_addition_positive(self):
        self.assertEqual(addition(1, 2), 3)
    def test_addition_negative(self):
        self.assertEqual(addition(-1, 1), 0)
    def test_addition_negatives(self):
        self.assertEqual(addition(-1, -1), -2)
    def test_addition_float(self):
        self.assertAlmostEqual(addition(1.1, 2.2), 3.3)
    def test_addition_int_float(self):
        self.assertAlmostEqual(addition(1, 2.2), 3.2)
    def test_addition_float_int(self):
        self.assertAlmostEqual(addition(2.2, 1), 3.2)
    def test_addition_int_string(self):
        with self.assertRaises(TypeError):
            addition(1, '2')
    def test_addition_string_int(self):
        with self.assertRaises(TypeError):
            addition('1', 2)
    def test_addition_string(self):
        with self.assertRaises(TypeError):
            addition('1', '2')
    def test_addition_none(self):
        with self.assertRaises(TypeError):
            addition(None, None)
    def test_addition_list(self):
        with self.assertRaises(TypeError):
            addition([1, 2], [3, 4])

if __name__ == '__main__':
    unittest.main()
