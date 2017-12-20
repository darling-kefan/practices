import sys
import unittest

def parse_int(s):
    return int(s)

class TestConversion(unittest.TestCase):
    def test_bad_int(self):
        self.assertRaises(ValueError, parse_int, 'N/A')

    def test_bad_int2(self):
        try:
            r = parse_int('N/A')
        except ValueError as e:
            self.assertEqual(type(e), ValueError)

    def test_bad_int3(self):
        try:
            r = parse_int(1)
        except ValueError as e:
            self.assertEqual(type(e), ValueError)
        else:
            self.fail('ValueError not raised')

    # assertRaisesRegex()方法，它可同时测试异常的存在以及通过正则表达式
    # 匹配异常的字符串表示
    def test_bad_int4(self):
        self.assertRaisesRegex(ValueError, 'invalid literal .*',
                                   parse_int, '1')

    def test_bad_int5(self):
        with self.assertRaisesRegex(ValueError, 'invalid literal .*'):
            r = parse_int('N/A')

def main(out=sys.stderr, verbosity=2):
    loader = unittest.TestLoader()
    suite = loader.loadTestsFromModule(sys.modules[__name__])
    unittest.TextTestRunner(out, verbosity=verbosity).run(suite)

if __name__ == '__main__':
    # unittest.main()
    with open('test.out', 'w') as f:
        main(f)
