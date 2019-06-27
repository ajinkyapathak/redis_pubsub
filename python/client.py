import redis
r = redis.StrictRedis(host='localhost', port=6379)
p = r.pubsub()
print("Starting main scripts...")
for i in range(1, 10):
	r.publish('mychannel1', 'python')
	print("Done")
