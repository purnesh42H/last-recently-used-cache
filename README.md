## Requirements

- Cache should have a fixed capacity
- Cache should allow adding a key, value pair if within capacity
    - If capacity is full, cache should evict the least recently used key
    - If key already exits, value of key should be updated
- Cache should return the value if key is present in the cache
- Usage of the key is determined by how many times value of key is read/updated
- Writers should be able to add key value pair in parallel
- Multiple readers should be able to read the value of key in parallel
- Add and Get should be O(1)