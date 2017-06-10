# rajirec

A tool written in golang to record internet radio streams from らじるらじる (rajirurajiru).

## Requirements

- Go >= 1.8 (may also work in the lower versions)
- Ubuntu 16.04 (may also work in other Unix-like OS)

## Install

```
go get github.com/hitochan777/rajirec
```

## Usage

First you should check the available areas with the following command.

```
$ rajirec area

area    area code
=================
仙台    sendai
東京    tokyo
名古屋  nagoya
大阪    osaka
広島    hiroshima
松山    matsuyama
福岡    fukuoka
札幌    sapporo
```
Then you can use the following command to record.
It starts to record on channel r2 in tokyo area for 5 minites, and saves to output.m4a.
Currently we support m4a only.
Supported channels are r1, r2, and fm.

```
rajirec record -areaid tokyo -channel r2 -duration 5 -output output.m4a
```

If you want to record regularly, you should first book records with the following command.

```
$ rajirec book -start "on sat at 22:00" -duration 15 -station_id tokyo -channel r2 -prefix r2
```
This will book a recording that starts on every Saturday at 22:00, and lasts for 15 minutes.
`-prefix` is the prefix of the output filenames. 
Each output file will be of the form `{prefix}_{start}.m4a` where `start` is the start time of the recording. 
