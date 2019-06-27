import redis
import time
r = redis.StrictRedis(host='localhost', port=6379)
p = r.pubsub()                                                              
p.subscribe('mychannel1')         
PAUSE = True
while PAUSE:                                                                # Will stay in loop until START message received
	message = p.get_message()                                               # Checks for message
	if message:
		data = message['data']                                          # Get data from message
		print(data)                                                 # Breaks loop
	time.sleep(1)
