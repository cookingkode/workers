from redis import StrictRedis
import simplejson

class Resque(object):
    """Dirt simple Resque client in Python. Can be used to create jobs."""
    redis_server = 'localhost:6379'
    namespace = 'goworkerqueue'

    def __init__(self):
        host, port = self.redis_server.split(':')
        self.redis = StrictRedis(host=host, port=int(port))

    def push(self, queue, object):
        key = "%s:%s" % (self.namespace, queue)
        self.redis.rpush(key, simplejson.dumps(object))


queue = Resque()
queue.push('sampleadd', {'class':'SampleAddJobClass', 'args':[ 1, 2, 3, 4]})