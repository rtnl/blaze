package main

import (
	"encoding/json"
	"os"

	"github.com/rtnl/blaze/pkg/provider/craft"
)

func main() {
	var (
		err error
	)

	report, err := craft.ReadPacketsReport("./packets.json")
	if err != nil {
		panic(err)
	}

	packets := report.GeneratePackets()

	packetsOutput, err := os.Create("packets_output.json")
	if err != nil {
		panic(err)
	}

	defer packetsOutput.Close()

	encoder := json.NewEncoder(packetsOutput)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(packets)
	if err != nil {
		panic(err)
	}

}
