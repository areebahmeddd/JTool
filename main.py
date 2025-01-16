from serialize import serialize
from deserialize import deserialize

if __name__ == '__main__':
    # Test 1: Serialization
    data = {'name': 'Areeb', 'age': 21, 'is_student': True}
    serialized_data = serialize(data)
    print(f'Serialized Data: {serialized_data}')

    # Test 2: Deserialization
    json_str = '{"name": "Areeb", "age": 21, "is_student": true}'
    deserialized_data = deserialize(json_str)
    print(f'\nDeserialized Data: {deserialized_data}')

    # Need to fix this error later
    # # Test 3: Serialization and then Deserialization
    # deserialized_back = deserialize(serialized_data)
    # print(f'\nDeserialized Back from Serialized Data: {deserialized_back}')
