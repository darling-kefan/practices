import cfg
from cfg import *

if __name__ == '__main__':
    print(cfg.var, var)
    var = 20
    print(cfg.var, var)

    print('-----------')
    print(cfg.somelist, somelist)
    somelist.append(1)

    print(cfg.somelist, somelist)

    print('-----------')
    var2 = 2
    print(var2, cfg.var2)
    show_var2()

    # 10 10
    # 10 20
    # -----------------
    # [] []
    # [1] [1]
    # -----------------
    # 1 1 1
