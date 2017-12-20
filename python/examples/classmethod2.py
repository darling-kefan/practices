class Date(object):

    def __init__(self, day=0, month=0, year=0):
        self.day = day
        self.month = month
        self.year = year

    @classmethod
    def from_string(cls, date_as_string):
        day, month, year = map(int, date_as_string.split('-'))
        date1 = cls(day, month, year)
        return date1

    @staticmethod
    def is_date_valid(date_as_string):
        day, month, year = map(int, date_as_string.split('-'))
        return day <= 31 and month <= 12 and year <= 3999

if __name__ == '__main__':
    dt = Date(28, 11, 2017)
    print(dt.day, dt.month, dt.year)

    dt2 = Date.from_string('28-11-2017')
    print(dt2.day, dt2.month, dt2.year)

    is_date = Date.is_date_valid('12-11-2322')
    print(is_date)

