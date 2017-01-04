package main

import (
	"fmt"
	"os"

	"github.com/NothNoth/SSLSplitParser/spexplode"
	"github.com/NothNoth/SSLSplitParser/spparser"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "<logfile> <output dir>")
		return
	}

	//Make sur log file exists
	st, err := os.Stat(os.Args[1])
	if os.IsNotExist(err) || st.IsDir() {
		fmt.Println("Invalid log file:", os.Args[1])
		return
	}

	//Make sure output folder is ready
	st, err = os.Stat(os.Args[2])
	if os.IsExist(err) && (st.IsDir() == false) {
		fmt.Println("Invalid output dir:", os.Args[2])
		return
	}
	if os.IsNotExist(err) {
		os.MkdirAll(os.Args[2], os.ModePerm)
	}

	chunks, err := spparser.ParseLog(os.Args[1])
	if err != nil {
		fmt.Println("An error occured during parsing:", err)
	}

	fmt.Println(len(chunks), "SSLSplit chunks loaded")
	if len(chunks) != 0 {
		for idx, c := range chunks {
			fmt.Println("#", idx, "from", c.Descriptor.SrcIP, "to", c.Descriptor.DestIP, "(", c.Descriptor.Size, "bytes )")
			err := spexplode.Explode(&c, idx, os.Args[2])
			if err != nil {
				fmt.Println("Failed to explode chunk ", idx, ":", err)
			}
		}
	}
}
