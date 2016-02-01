# GoWorker Recipe

Background Workers written in Golang, tasked from Golang and Python

## Need

It is often required to do a lot of computation asynchronously. Sometimes in a 
language like python, there is a need  to 'get overthe GIL' and for  concurrency,
computation horsepower that a langauge like Golang offers.

In the python/Django world, there are mature solutions like Celery. However they
are limited by 
- Worker code being in python
- Performance dependence of backends

This recipe aims to solve the above issues by
- Writing workers in Golang
- Keeeping 'results backend' open, rather than forcing communication over a 
'backend'

The recipe follows the Resque protocol, mainly because of it's simplicity

## Structure 

### src/worker 
This is the Golang worker example. It is based on goworker. You read about this 
here : https://github.com/benmanns/goworker

Important command line options include


* `-queues="comma,delimited,queues"` — This is the only required flag. The recommended practice is to separate your Resque workers from your goworkers with different queues. Otherwise, Resque worker classes that have no goworker analog will cause the goworker process to fail the jobs. Because of this, there is no default queue, nor is there a way to select all queues (à la Resque's `*` queue). If you have multiple queues you can assign them weights. A queue with a weight of 2 will be checked twice as often as a queue with a weight of 1: `-queues='high=2,low=1'`. Note this is jsut the queue name, not the namespace
* `-interval=5.0` — Specifies the wait period between polling if no job was in the queue the last time one was requested. Make it smaller (say 0.1) to make response faster 
* `-uri=redis://localhost:6379/` — Specifies the URI of the Redis database from which goworker polls for jobs. Accepts URIs of the format `redis://user:pass@host:port/db` or `unix:///path/to/redis.sock`. The flag may also be set by the environment variable `$($REDIS_PROVIDER)` or `$REDIS_URL`. E.g. set `$REDIS_PROVIDER` to `REDISTOGO_URL` on Heroku to let the Redis To Go add-on configure the Redis database.
* `-namespace=resque:` — Specifies the namespace from which goworker retrieves jobs and stores stats on workers.

For the recipe, namespace  is "goworker" and queue name for the add sample is "sampleadd"

To build server, 
- Set Gopath to <where_you_cloned>/vendors:<where_you_cloned>
- go build worker

### src/client

This is a Golang Resque client written from scratch. The simplicity of the protocol
makes this happen


To build client, 
- Set Gopath to <where_you_cloned>/vendors:<where_you_cloned>
- go build client

### python

This is a Python Resque client written from scratch . 

## TODO

* Fork and submit an MR for goworker to use redis-cluster. This will solve SPOF and
improve perf
* Update clients
