
# ethflux

*ethflux* is an InfluxDB data gatherer for ethtool-style network interface information. It uses the Linux SIOCETHTOOL ioctl interface to obtain network interface statistics and other runtime data and outputs them in InfluxDB's [line protocol format]([https://docs.influxdata.com/influxdb/v1.4/write_protocols/line_protocol_reference/) for further propagation.

## Getting the code

```
go get github.com/DCSO/ethflux.git
```

## Building

```
go build -v
```

will produce an `ethflux` binary.

## Usage

```
Usage: ethflux [-measurement=string] <iface>
  -measurement string
    	name of measurement to use in line (default "ethtool")
```

## Example

```
$ ./ethflux enp175s0f1
ethtool,host=xxx.xxxxx.de,interface=enp175s0f1 collisions=0,driver=i40e,fcoe_bad_fccrc=0,fcoe_ddp_count=0,fcoe_last_error=0,port.VF_admin_queue_requests=0,port.arq_overflows=0,port.fdir_atr_match=0,port.fdir_atr_status=1,port.fdir_atr_tunnel_match=0,port.fdir_flush_cnt=0,port.fdir_sb_match=0,port.fdir_sb_status=1,port.illegal_bytes=0,port.link_xoff_rx=0,port.link_xoff_tx=0,port.link_xon_rx=0,port.link_xon_tx=0,port.mac_local_faults=0,port.mac_remote_faults=0,port.rx_broadcast=1503484,port.rx_bytes=273991760610806,port.rx_crc_errors=0,port.rx_csum_bad=10675196,port.rx_dropped=0,port.rx_fragments=0,port.rx_hwtstamp_cleared=0,port.rx_jabber=0,port.rx_length_errors=0,port.rx_lpi_count=0,port.rx_lpi_status=0,port.rx_multicast=6118060709,port.rx_oversize=0,port.rx_priority_0_xoff=0,port.rx_priority_0_xon=0,port.rx_priority_0_xon_2_xoff=0,port.rx_priority_1_xoff=0,port.rx_priority_1_xon=0,port.rx_priority_1_xon_2_xoff=0,port.rx_priority_2_xoff=0,port.rx_priority_2_xon=0,port.rx_priority_2_xon_2_xoff=0,port.rx_priority_3_xoff=0,port.rx_priority_3_xon=0,port.rx_priority_3_xon_2_xoff=0,port.rx_priority_4_xoff=0,port.rx_priority_4_xon=0,port.rx_priority_4_xon_2_xoff=0,port.rx_priority_5_xoff=0,port.rx_priority_5_xon=0,port.rx_priority_5_xon_2_xoff=0,port.rx_priority_6_xoff=0,port.rx_priority_6_xon=0,port.rx_priority_6_xon_2_xoff=0,port.rx_priority_7_xoff=0,port.rx_priority_7_xon=0,port.rx_priority_7_xon_2_xoff=0,port.rx_size_1023=74722350130,port.rx_size_127=368569804156,port.rx_size_1522=490051931062,port.rx_size_255=67399150443,port.rx_size_511=31852805837,port.rx_size_64=0,port.rx_size_big=0,port.rx_undersize=0,port.rx_unicast=1026476477402,port.tx_broadcast=0,port.tx_bytes=3294186,port.tx_dropped_link_down=0,port.tx_errors=0,port.tx_lpi_count=0,port.tx_lpi_status=0,port.tx_multicast=37911,port.tx_priority_0_xoff=0,port.tx_priority_0_xon=0,port.tx_priority_1_xoff=0,port.tx_priority_1_xon=0,port.tx_priority_2_xoff=0,port.tx_priority_2_xon=0,port.tx_priority_3_xoff=0,port.tx_priority_3_xon=0,port.tx_priority_4_xoff=0,port.tx_priority_4_xon=0,port.tx_priority_5_xoff=0,port.tx_priority_5_xon=0,port.tx_priority_6_xoff=0,port.tx_priority_6_xon=0,port.tx_priority_7_xoff=0,port.tx_priority_7_xon=0,port.tx_size_1023=0,port.tx_size_127=37911,port.tx_size_1522=0,port.tx_size_255=0,port.tx_size_511=0,port.tx_size_64=0,port.tx_size_big=0,port.tx_timeout=0,port.tx_unicast=0,rx-0.rx_bytes=25760113879251,rx-0.rx_packets=25164365487,rx-1.rx_bytes=13632265611556,rx-1.rx_packets=16152533578,rx-10.rx_bytes=68451372537673,rx-10.rx_packets=175032632872,rx-11.rx_bytes=17757173810258,rx-11.rx_packets=15835980314,rx-12.rx_bytes=8764805027779,rx-12.rx_packets=10140123021,rx-13.rx_bytes=22223849442838,rx-13.rx_packets=29276067160,rx-14.rx_bytes=10391035607365,rx-14.rx_packets=14932971047,rx-15.rx_bytes=23056778613056,rx-15.rx_packets=25284323910,rx-16.rx_bytes=16051019487062,rx-16.rx_packets=21259433681,rx-17.rx_bytes=10583831362520,rx-17.rx_packets=13614422175,rx-18.rx_bytes=22096971062957,rx-18.rx_packets=21956959527,rx-19.rx_bytes=18742982939435,rx-19.rx_packets=17696673457,rx-2.rx_bytes=14106298982560,rx-2.rx_packets=14531863773,rx-20.rx_bytes=19922195951231,rx-20.rx_packets=28680151786,rx-21.rx_bytes=19699293443204,rx-21.rx_packets=21210827777,rx-22.rx_bytes=12979243896441,rx-22.rx_packets=16005967104,rx-23.rx_bytes=23144497271190,rx-23.rx_packets=28930097115,rx-24.rx_bytes=12260062456916,rx-24.rx_packets=15883136613,rx-25.rx_bytes=19986459301888,rx-25.rx_packets=23650385573,rx-26.rx_bytes=11428212190208,rx-26.rx_packets=12908916174,rx-27.rx_bytes=10718217684165,rx-27.rx_packets=12732850563,rx-28.rx_bytes=147525686777846,rx-28.rx_packets=148249910396,rx-29.rx_bytes=9558245799695,rx-29.rx_packets=12600295813,rx-3.rx_bytes=11191403160576,rx-3.rx_packets=12879280559,rx-30.rx_bytes=13361776421559,rx-30.rx_packets=14147715684,rx-31.rx_bytes=19280248807075,rx-31.rx_packets=27950447621,rx-32.rx_bytes=14283505845843,rx-32.rx_packets=13678440382,rx-33.rx_bytes=11202494934134,rx-33.rx_packets=15966233450,rx-34.rx_bytes=10681453089257,rx-34.rx_packets=19005483221,rx-35.rx_bytes=13209635179110,rx-35.rx_packets=21859446329,rx-36.rx_bytes=7847061371638,rx-36.rx_packets=11645298740,rx-37.rx_bytes=18618156576591,rx-37.rx_packets=16170486726,rx-38.rx_bytes=15446372902723,rx-38.rx_packets=15934877929,rx-39.rx_bytes=16810922790994,rx-39.rx_packets=16584026771,rx-4.rx_bytes=14554570408487,rx-4.rx_packets=16112245969,rx-5.rx_bytes=13744045635641,rx-5.rx_packets=15680575214,rx-6.rx_bytes=24422177040450,rx-6.rx_packets=23549431627,rx-7.rx_bytes=13128092183974,rx-7.rx_packets=13992639624,rx-8.rx_bytes=30812727556210,rx-8.rx_packets=27043522983,rx-9.rx_bytes=19596466745320,rx-9.rx_packets=27078656463,rx_alloc_fail=0,rx_broadcast=1503284,rx_bytes=827031723770173,rx_crc_errors=0,rx_dropped=1410479433,rx_errors=0,rx_fcoe_dropped=0,rx_fcoe_dwords=0,rx_fcoe_packets=0,rx_length_errors=0,rx_multicast=6117053025,rx_packets=1031009698188,rx_pg_alloc_fail=0,rx_unicast=1026301621346,rx_unknown_protocol=0,tx-0.tx_bytes=0,tx-0.tx_packets=0,tx-1.tx_bytes=0,tx-1.tx_packets=0,tx-10.tx_bytes=0,tx-10.tx_packets=0,tx-11.tx_bytes=0,tx-11.tx_packets=0,tx-12.tx_bytes=0,tx-12.tx_packets=0,tx-13.tx_bytes=0,tx-13.tx_packets=0,tx-14.tx_bytes=0,tx-14.tx_packets=0,tx-15.tx_bytes=0,tx-15.tx_packets=0,tx-16.tx_bytes=0,tx-16.tx_packets=0,tx-17.tx_bytes=22738,tx-17.tx_packets=323,tx-18.tx_bytes=0,tx-18.tx_packets=0,tx-19.tx_bytes=0,tx-19.tx_packets=0,tx-2.tx_bytes=0,tx-2.tx_packets=0,tx-20.tx_bytes=0,tx-20.tx_packets=0,tx-21.tx_bytes=0,tx-21.tx_packets=0,tx-22.tx_bytes=0,tx-22.tx_packets=0,tx-23.tx_bytes=0,tx-23.tx_packets=0,tx-24.tx_bytes=0,tx-24.tx_packets=0,tx-25.tx_bytes=0,tx-25.tx_packets=0,tx-26.tx_bytes=0,tx-26.tx_packets=0,tx-27.tx_bytes=0,tx-27.tx_packets=0,tx-28.tx_bytes=0,tx-28.tx_packets=0,tx-29.tx_bytes=0,tx-29.tx_packets=0,tx-3.tx_bytes=0,tx-3.tx_packets=0,tx-30.tx_bytes=0,tx-30.tx_packets=0,tx-31.tx_bytes=0,tx-31.tx_packets=0,tx-32.tx_bytes=0,tx-32.tx_packets=0,tx-33.tx_bytes=0,tx-33.tx_packets=0,tx-34.tx_bytes=0,tx-34.tx_packets=0,tx-35.tx_bytes=0,tx-35.tx_packets=0,tx-36.tx_bytes=0,tx-36.tx_packets=0,tx-37.tx_bytes=0,tx-37.tx_packets=0,tx-38.tx_bytes=0,tx-38.tx_packets=0,tx-39.tx_bytes=0,tx-39.tx_packets=0,tx-4.tx_bytes=0,tx-4.tx_packets=0,tx-5.tx_bytes=0,tx-5.tx_packets=0,tx-6.tx_bytes=0,tx-6.tx_packets=0,tx-7.tx_bytes=0,tx-7.tx_packets=0,tx-8.tx_bytes=0,tx-8.tx_packets=0,tx-9.tx_bytes=0,tx-9.tx_packets=0,tx_broadcast=0,tx_bytes=22738,tx_dropped=0,tx_errors=0,tx_fcoe_dwords=0,tx_fcoe_packets=0,tx_force_wb=0,tx_linearize=0,tx_lost_interrupt=0,tx_multicast=323,tx_packets=323,tx_unicast=0 1519229432806537753
```
In Telegraf, for instance, one could use this software using the `exec` input:
```toml
[[inputs.exec]]
  commands = ["ethflux enp175s0f1"]
  timeout = "5s"
  data_format = "influx"
```

## Built Using

* [fluxline](https://github.com/DCSO/fluxline) - Encoder for InfluxDB LineProtocol
* [ethtool](https://github.com/safchain/ethtool) - Go library to talk to the kernel


## License

This project is licensed under a 3-clause BSD-like license.
