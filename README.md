# Dragon

<!-- MarkdownTOC autolink=true autoanchor=true-->

- [start dragon](#start-dragon)
- [ab performance](#ab-performance)

<!-- /MarkdownTOC -->



![CI Status](https://travis-ci.org/azerothyang/dragon.svg?branch=master)

 Dragon 🐲 is a lightweight high performance web framework with [Go](https://golang.org/) for the feature and comfortable develop

# start dragon
 dragon is mvc go framework. So you can hack with controller and model.  
 
 __build__: 
 >Just compile your src file and move bin file to directory dragon/release/
 
 >you might develop with [gobuild](https://github.com/caixw/gobuild) for hot rebuilding.
 just like:
 ```
 linux:
 gobuild -o ../release/go
 
 windows:
 gobuild -o ../release/go.exe
 ```
 
 __config__:
> dragon/release/conf/  

# ab performance
 ``` 
 cpu: 1 core, ram: 1 G
 Server Software:        nginx/1.12.2
 Server Hostname:        test.com
 Server Port:            80
 
 Document Path:          /
 Document Length:        13 bytes
 
 Concurrency Level:      100
 Time taken for tests:   9.341 seconds
 Complete requests:      100000
 Failed requests:        0
 Write errors:           0
 Total transferred:      17700000 bytes
 HTML transferred:       1300000 bytes
 Requests per second:    10705.75 [#/sec] (mean)
 Time per request:       9.341 [ms] (mean)
 Time per request:       0.093 [ms] (mean, across all concurrent requests)
 Transfer rate:          1850.51 [Kbytes/sec] received
 
 Connection Times (ms)
               min  mean[+/-sd] median   max
 Connect:        0    1   1.2      1       9
 Processing:     0    8   2.7      8      30
 Waiting:        0    7   2.5      6      29
 Total:          0    9   2.9      9      31
 
 Percentage of the requests served within a certain time (ms)
   50%      9
   66%     10
   75%     11
   80%     11
   90%     13
   95%     14
   98%     16
   99%     18
  100%     31 (longest request)

 ```