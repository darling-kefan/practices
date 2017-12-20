"""
    >>> clip('banana ', 6)
    'banana'
    >>> clip('banana ', 7)
    'banana'
    >>> clip('banana ', 5)
    'banana'
    >>> clip('banana split', 6)
    'banana'
    >>> clip('banana split', 7)
    'banana'
    >>> clip('banana split', 10)
    'banana'
    >>> clip('banana split', 11)
    'banana'
    >>> clip('banana split', 12)
    'banana split'
"""

# 添加函数注解
# Python对注解所做的唯一的事情是，把它们存储在函数的__annotations__属性里。仅此而已，Python不做检查、不做强制、不做验证，什么都不做

def clip(text:str, max_len:'int > 0'=80) -> str:
    end = None
    if len(text) > max_len:
        if text[max_len] == ' ':
            end = max_len
        else:
            space_before = text.rfind(' ', 0, max_len)
            if space_before >= 0:
                end = space_before
            else:
                space_end = text.find(' ', max_len)
                if space_end >= 0:
                    end = space_end

    if end is None:
        end = len(text)

    return text[:end].rstrip()

if __name__ == '__main__':
    text = 'banana'
    #print(clip(text, 6))
    #print(clip(text, 7))
    #print(clip(text, 5))

    text = 'banana split test'
    print(clip(text, 6))
    print(clip(text, 7))
    print(clip(text, 10))
    print(clip(text, 11))
    print(clip(text, 12))
    print(clip(text, 13))
    print(clip(text, 14))
    print(clip(text, 15))
    print(clip(text, 16))
    print(clip(text, 17))
    print(clip(text, 18))
