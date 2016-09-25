# `timestats`

I wanted a more versatile `time` command that could be used to run many
timed runs of the same command and gather statistics on the results. I also
needed a tool that could save the raw data gathered to disk for recording
and further analysis.

I wrote this fairly simple Golang binary to do this. Some of the features are:

### Statistics summary:

Some basic, commonly asked for stats:

```
Statistics (Nanoseconds):
-------------------------
Count 1000
Min   3057538.000000 (3.06 milliseconds)
Mean  6353627.262000 (6.35 milliseconds)
Max   20099715.000000 (20.10 milliseconds)

P25   4217045.500000 (4.22 milliseconds)
P50   5459587.500000 (5.46 milliseconds)
P75   7426019.000000 (7.43 milliseconds)
P90   10052752.000000 (10.05 milliseconds)
P95   13415480.000000 (13.42 milliseconds)
```

### Distribution graph:

A basic graph to show the distribution of the elapsed times. This graph
component was really nice to make and I'm quite pleased with the result. I will
definitely be reusing the same code in other tools.

```
Distribution (normalized):
-------------------------
 █
 █_         ▄
 ██         ██
 ██         ██
 ██        ███▄
_███       ████
████▄    _▄████  ___  ▄▄
█████_ _▄███████▄████ ██ █_ _ _
█████████████████████▄█████████▄ ▄ _
█████████████████████████████████████▄█_▄▄_▄▄▄_▄_▄_▄█__▄▄_____▄_▄▄__▄_  ____  _  __ ▄ ▄__  ____   __
----------------------------------------------------------------------------------------------------
3.06 milliseconds                                                                 20.10 milliseconds
```

## Example: Pinging `facebook.com`

Networks are funny animals, they don't behave as you expect. So while diagnosing
a connection why not do some stats on the ping time?

```
$ ./timestats --count 100 --output facebook.json ping -c1 facebook.com
PING facebook.com (66.220.146.36): 56 data bytes
64 bytes from 66.220.146.36: icmp_seq=0 ttl=70 time=345.304 ms

--- facebook.com ping statistics ---
1 packets transmitted, 1 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 345.304/345.304/345.304/0.000 ms
#0 succeeded after 351.390991ms
PING facebook.com (66.220.146.36): 56 data bytes
64 bytes from 66.220.146.36: icmp_seq=0 ttl=70 time=402.874 ms

--- facebook.com ping statistics ---
1 packets transmitted, 1 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 402.874/402.874/402.874/0.000 ms
#1 succeeded after 409.90095ms

...<snip>...

--- facebook.com ping statistics ---
1 packets transmitted, 1 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 403.092/403.092/403.092/0.000 ms
#99 succeeded after 410.022172ms

Statistics (Nanoseconds):
-------------------------
Count 100
Min   312390716.000000 (312.39 milliseconds)
Mean  398728423.750000 (398.73 milliseconds)
Max   412830311.000000 (412.83 milliseconds)

P25   409004564.250000 (409.00 milliseconds)
P50   409433110.500000 (409.43 milliseconds)
P75   409697745.000000 (409.70 milliseconds)
P90   409943612.000000 (409.94 milliseconds)
P95   410154294.000000 (410.15 milliseconds)

Distribution (normalized):
-------------------------
                                                                                                █
                                                                                                █
                                                                                                █
                                                                                                █
                                                                                                █
                                                                                                █
                                                                                                █
                                                                                                ██
_                                                                                               ██
█___                                  _                                             _   __▄_ _  ██ _
----------------------------------------------------------------------------------------------------
312.55 milliseconds                                                              410.48 milliseconds
```

It is interesting to see 2 clusters of times: one at 312ms and one at 410ms. One could predict that
these are 2 slightly different routing paths, one of which is 100ms faster than the other. Or maybe
its just gremlins on the wire, time for other tools to check it out.

## Future work:

- Show mean/p50 on histogram
- Coloured output?
- Provide analysis of exit codes as well
