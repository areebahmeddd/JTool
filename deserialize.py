def deserialize(json_str):
    # Strip leading and trailing whitespace
    json_str = json_str.strip()
    
    # Handle null
    if json_str == 'null':
        return None
    
    # Handle booleans
    elif json_str == 'true':
        return True
    elif json_str == 'false':
        return False
    
    # Handle numbers (integers or floats)
    elif json_str.isdigit() or (json_str[0] == '-' and json_str[1:].isdigit()):
        return int(json_str)  # Handle integers
    
    elif '.' in json_str and json_str.replace('.', '', 1).isdigit():
        return float(json_str)  # Handle floats
    
    # Handle strings
    elif json_str.startswith('"') and json_str.endswith('"'):
        return json_str[1:-1]  # Remove surrounding quotes
    
    # Handle objects (dictionaries)
    elif json_str.startswith('{') and json_str.endswith('}'):
        obj = {}
        inner_str = json_str[1:-1].strip()  # Remove the curly braces
        key_value_pairs = inner_str.split(',')
        
        for pair in key_value_pairs:
            key, value = pair.split(":")
            obj[deserialize(key.strip())] = deserialize(value.strip())
        
        return obj
    
    # Handle arrays (lists)
    elif json_str.startswith('[') and json_str.endswith(']'):
        arr = []
        inner_str = json_str[1:-1].strip()  # Remove square brackets
        elements = inner_str.split(',')
        
        for element in elements:
            arr.append(deserialize(element.strip()))
        
        return arr
    
    else:
        raise ValueError(f"Invalid JSON string: {json_str}")
