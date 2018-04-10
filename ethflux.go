package main

// ethflux
// Copyright (c) 2018, DCSO GmbH

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"

	"github.com/DCSO/fluxline"
	log "github.com/Sirupsen/logrus"
	"github.com/safchain/ethtool"
)

func main() {

	msmt := flag.String("measurement", "ethtool",
		"name of measurement to use in line")
	help := flag.Bool("help", false, "show help")
	quiet := flag.Bool("quiet", false, "do not show warnings")
	flag.Parse()

	argsWithoutProg := flag.Args()

	ifs, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	if *help {
		fmt.Fprintln(os.Stderr, "Usage: ethflux [-help] [-quiet] [-measurement=string] <iface pattern>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	pattern := ".*"
	if len(argsWithoutProg) > 0 {
		pattern = argsWithoutProg[0]
	}

	ethHandle, err := ethtool.NewEthtool()
	if err != nil {
		log.Fatal(err)
	}
	for _, intf := range ifs {
		if len(pattern) > 0 {
			match, err := regexp.MatchString(pattern, intf.Name)
			if err != nil {
				log.Fatal(err)
			}
			if !match {
				continue
			}
		}
		stats, err := ethHandle.Stats(intf.Name)
		if err != nil {
			if !*quiet {
				log.Warn(err)
			}
			continue
		}
		drivername, err := ethHandle.DriverName(intf.Name)
		if err != nil {
			log.Fatal(err)
		}

		stringStats := make(map[string]string)
		for k, v := range stats {
			stringStats[k] = fmt.Sprintf("%d", v)
		}
		stringStats["driver"] = fmt.Sprintf(`"%s"`, drivername)
		tags := make(map[string]string)
		tags["interface"] = intf.Name

		var b bytes.Buffer
		enc := fluxline.NewEncoder(&b)
		enc.EncodeMap(*msmt, stringStats, tags)
		fmt.Println(b.String())
	}
}
