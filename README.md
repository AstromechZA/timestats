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

## More to come:

- JSON data dump
- Show mean/p50 on histogram
- Coloured output?
- Make instructions, release binaries
