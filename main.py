def serialize(obj):
    # String
    if isinstance(obj, str):
        return f'"{obj}"'

    # Integer or Float
    elif isinstance(obj, (int, float)):
        return str(obj)

    # Boolean
    elif isinstance(obj, bool):
        return 'true' if obj else 'false'

    # None se null karna h
    elif obj is None:
        return 'null'

    # Dictionary
    elif isinstance(obj, dict):
        items = []
        for key, value in obj.items():
            serialized_key = serialize(key)
            serialized_value = serialize(value)
            items.append(f'{serialized_key}:{serialized_value}')
            return '{' + ','.join(items) + '}'
