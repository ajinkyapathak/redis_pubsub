import redis

r = redis.StrictRedis(host='localhost', port=6379)
print("Starting main scripts...")
for i in range(1, 10):
    r.publish('mychannel1', 'python')
    print("Done")
