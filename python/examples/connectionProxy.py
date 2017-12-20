DEFAULT_DB_ATIAS = 'default'

connections = {'default': 1, 'second': 2}

class DefaultConnectionProxy:

    def __getattr__(self, item):
        return getattr(connections[DEFAULT_DB_ATIAS], item)

    def __setattr__(self, name, value):
        return setattr(connections[DEFAULT_DB_ATIAS], name, value)

    def __delattr__(self, name):
        return delattr(connections[DEFAULT_DB_ATIAS], name)

    def __eq__(self, other):
        return connections[DEFAULT_DB_ALIAS] == other

connection = DefaultConnectionProxy()

print(connection, connections)
