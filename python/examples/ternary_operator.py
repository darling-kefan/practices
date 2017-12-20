for i in range(1, 3):
    if i == 1:
        plural = ''
    else:
        plural = 's'
    print("The loop ran %d time%s" % (i, plural))

for i in range(1, 3):
    print("The loop ran {} time{}".format(i, (i != 1 and 's' or '')))

for i in range(1, 3):
    print("The loop run {} time{}".format(i, ('', 's')[i!=1]))
