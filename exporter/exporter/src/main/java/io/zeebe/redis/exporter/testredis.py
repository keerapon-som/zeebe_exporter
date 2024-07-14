import redis

# Replace these variables with your Redis instance details
redis_host = 'localhost'
redis_port = 6379
redis_password = 'exampleRedisPassword'

# Creating a Redis connection
r = redis.Redis(host=redis_host, port=redis_port, password=redis_password, decode_responses=True)

# Testing the connection
# print(r.get('gem_key_test'))
r.set('gem_key_test', 0)

# def getAllkeys():
#     keys = r.keys('*')
#     for key in keys:
#         print(key)
        
# getAllkeys()

# def getDatafromstream():
#     stream = r.xrange('streamExporter', count=100)
#     c = 0
#     for message in stream:
        
#         # print(message[0])
#         c += 1
#         print(c)
#         # delete the message
#         r.xdel('streamExporter', message[0])
        
# def getDataFromStreamInfinite():
#     last_id = '0-0'  # Start from the beginning of the stream
#     while True:
#         # Use xread in a blocking mode, waiting for new messages
#         stream = r.xread({'streamExporter': last_id}, count=100, block=1000)
#         if stream:
#             for stream_name, messages in stream:
#                 for message in messages:
#                     message_id, message_data = message
#                     print(message_data["data"])
#                     # Optionally, delete the message after processing
#                     r.xdel('streamExporter', message_id)
#                     last_id = message_id  # Update last_id to the latest processed message ID

# getDataFromStreamInfinite()



