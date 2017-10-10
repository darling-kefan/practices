#! /usr/bin/env python3.5
# -*- coding: utf-8 -*-

import re
from datetime import datetime, timezone, timedelta

def to_timestamp(dt_str, tz_str):
    dt = datetime.strptime(dt_str, '%Y-%m-%d %H:%M:%S')

    tz = timezone.utc
    re_tz = re.match('^UTC([\+\-])(\d+)\:\d+$', tz_str)
    if re_tz:
        h = int(re_tz.group(2).lstrip('0'))
        if re_tz.group(1) == '-':
            h = -h
        tz = timezone(timedelta(hours=h))
    else:
        print('tz_str is error')
        exit()

    return dt.replace(tzinfo=tz).timestamp()

t1 = to_timestamp('2015-6-1 08:10:30', 'UTC+7:00')
assert t1 == 1433121030.0, t1

t2 = to_timestamp('2015-5-31 16:10:30', 'UTC-09:00')
assert t2 == 1433121030.0, t2

print('Pass')
