def serialize(obj):
    if isinstance(obj, str):
        return f'"{obj}"'

    elif isinstance(obj, int):
        return str(obj)

    elif isinstance(obj, bool):
        return 'true' if obj else 'false'

    elif obj is None:
        return 'null'
