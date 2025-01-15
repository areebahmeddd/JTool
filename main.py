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
        items = [f"{serialize(key)}:{serialize(value)}" for key, value in obj.items()]
        return '{' + ','.join(items) + '}'

    # Lists
    elif isinstance(obj, list):
        items = [serialize(item) for item in obj]
        return '[' + ','.join(items) + ']'

    else:
        raise TypeError(f'Unsupported type: {type(obj)}')
