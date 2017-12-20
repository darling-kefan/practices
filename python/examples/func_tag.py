#! /usr/bin/env python3
# -*- coding:utf-8 -*-

def tag(name, *content, cls=None, **attrs):
    """生成一个或多个HTML标签"""
    if cls is not None:
        attrs['class'] = cls
    if attrs:
        attr_str = ''.join(' %s="%s"' % (attr, value)
                           for attr, value
                           in sorted(attrs.items()))
    else:
        attr_str = ''

    if content:
        return '\n'.join('<%s%s>%s</%s>' %
                         (name, attr_str, c, name) for c in content)
    else:
        return '<%s%s />' % (name, attr_str)

if __name__ == '__main__':
    print(tag('br'))
    print(tag('p', 'hello'))
    print(tag('p', 'hello', 'world'))
    print(tag('p', 'hello', id=33))
    print(tag('p', 'hello', 'world', cls='sidebar'))
    # 即便第一个是定位参数也能作为关键字参数传入
    print(tag(content='testing', name="img"))
    my_tag = {'name': 'img', 'title': 'Sunset nBoulevard', 'src': 'sunset.jpg',
              'cls': "framed"}
    print(tag(**my_tag))
