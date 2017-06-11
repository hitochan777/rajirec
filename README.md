# rajirec 
![CircleCI](https://circleci.com/gh/hitochan777/rajirec.svg?style=svg&circle-token=47fa7b83155124988315c5c41358563a4da33350)

A tool written in golang to record internet radio streams from [らじるらじる (rajirurajiru)](http://www.nhk.or.jp/radio).

## Requirements

- Go >= 1.8 (may also work in the lower versions)
- Ubuntu 16.04 (may also work in other Unix-like OS)

## Install

```
go get github.com/hitochan777/rajirec/cmd/rajirec
```

## Usage

```
Usage: rajirec <flags> <subcommand> <subcommand args>
 
Subcommands:
       area             Show area information
       book             Book record
       help             describe subcommands and their syntax
       record           record live stream
       server           Run server to record
```

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
$ rajirec record -areaid tokyo -channel r2 -duration 5 -output output.m4a
```

If you want to record regularly, you should first book records with the following command.

```
$ rajirec book -start "on sat at 22:00" -duration 15 -station_id tokyo -channel r2 -prefix r2
```
This will book a recording that starts on every Saturday at 22:00, and lasts for 15 minutes.
`-prefix` is the prefix of the output filenames. 
Each output file will be of the form `{prefix}_{start}.m4a` where `start` is the start time of the recording. 

The record schedules are saved to a JSON file. By default, the file is created
at `$HOME/rajirec/rajirec/book`. You can change it by specifying the location of the custom configuration file (YAML format).
To do that you need to export `RAJIREC_CONFIG` that points to the config file location.
The configuration file should look like the follows:

```
general:
  # The URL of a XML file that specifies channel information. You basically should not change this.
  api_url: "http://www3.nhk.or.jp/netradio/app/config_pc_2016.xml"

db:
  db_dir: "/home/foo/rajirec_db"
  db_name: "rajirec"
  book_table: "book"
```

## TODO
- [ ] Users can delete bookings
- [ ] Merge default and custom config

## License

The MIT License

## Author
Hitoshi Otsuki
