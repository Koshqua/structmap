# Struct-map

### Problem 
I'm working on a project where we use redis.
I have encountered a problem with values that redis returns.
Redis returns map[string]interface{} and then you need to assign values yourself.
I think it would be much easier if you could map fields in struct to the keys in map and simply decode it, as you do with json. 

