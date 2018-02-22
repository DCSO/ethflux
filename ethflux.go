package main

// ethflux
// Copyright (c) 2018, DCSO GmbH

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/DCSO/fluxline"
	log "github.com/Sirupsen/logrus"
	"github.com/safchain/ethtool"
)

func main() {

	msmt := flag.String("measurement", "ethtool",
		"name of measurement to use in line")

	flag.Parse()

	argsWithoutProg := flag.Args()
	if len(argsWithoutProg) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: ethflux [-measurement=string] <iface>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	intf := argsWithoutProg[0]

	ethHandle, err := ethtool.NewEthtool()
	if err != nil {
		log.Fatal(err)
	}
	stats, err := ethHandle.Stats(intf)
	if err != nil {
		log.Fatal(err)
	}
	drivername, err := ethHandle.DriverName(intf)
	if err != nil {
		log.Fatal(err)
	}

	stringStats := make(map[string]string)
	for k, v := range stats {
		stringStats[k] = fmt.Sprintf("%d", v)
	}
	stringStats["driver"] = fmt.Sprintf(`"%s"`, drivername)
	tags := make(map[string]string)
	tags["interface"] = intf

	var b bytes.Buffer
	enc := fluxline.NewEncoder(&b)
	enc.EncodeMap(*msmt, stringStats, tags)
	fmt.Println(b.String())
}
