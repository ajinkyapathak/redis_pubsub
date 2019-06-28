import time

import redis

r = redis.StrictRedis(host='localhost', port=6379)
p = r.pubsub()
p.subscribe('mychannel1')
while True:
    message = p.get_message()
    if message:
        data = message['data']
        print(data)
    time.sleep(1)

