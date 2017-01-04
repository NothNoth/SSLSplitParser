package spparser

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

// ChunkDescriptor describes a chunk metadata
type ChunkDescriptor struct {
	Date     int64
	SrcIP    string
	SrcPort  uint16
	DestIP   string
	DestPort uint16
	Size     uint32
}

// Chunk holds metadata and data
type Chunk struct {
	Descriptor ChunkDescriptor
	Data       []byte
}

func parseDate(dateString string) (ts int64, err error) {
	layout := "2006-01-02 15:04:05 UTC"
	t, err := time.Parse(layout, dateString)

	if err != nil {
		return
	}
	ts = t.Unix()
	return
}

func parseIP(ipString string) (ip string, port uint16, err error) {
	exploded := strings.Split(ipString, ":")

	ip = exploded[0]
	if ip[0] != '[' || ip[len(ip)-1] != ']' {
		err = errors.New("Invalid ip field")
		return
	}
	ip = strings.TrimPrefix(ip, "[")
	ip = strings.TrimSuffix(ip, "]")

	p, err := strconv.ParseUint(exploded[1], 10, 16)
	if err != nil {
		return
	}
	port = uint16(p)

	return
}

func parseSize(sizeString string) (size uint32, err error) {

	sizeString = strings.TrimPrefix(sizeString, "(")
	sizeString = strings.TrimSuffix(sizeString, "):\n")

	s, err := strconv.ParseUint(sizeString, 10, 32)
	if err != nil {
		return
	}

	size = uint32(s)
	return
}

func parseDescriptor(dString string) (*ChunkDescriptor, error) {
	var d ChunkDescriptor
	// 2017-01-03 15:17:09 UTC [172.16.42.153]:53084 -> [91.190.216.81]:80 (65):
	exploded := strings.Split(dString, " ")
	if len(exploded) != 7 {
		return nil, errors.New("Invalid descriptor format")
	}

	d.Date, _ = parseDate(strings.Join(exploded[0:3], " "))
	d.SrcIP, d.SrcPort, _ = parseIP(exploded[3])
	d.DestIP, d.DestPort, _ = parseIP(exploded[5])
	d.Size, _ = parseSize(exploded[6])

	return &d, nil
}

func parseChunk(reader *bufio.Reader) (*Chunk, error) {

	chunkDescriptor, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Invalid chunk descriptor")
	}

	desc, err := parseDescriptor(chunkDescriptor)
	if err != nil || desc == nil {
		return nil, err
	}

	var d []byte
	for i := 0; i < int(desc.Size); i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		d = append(d, b)
	}

	return &Chunk{Descriptor: *desc, Data: d}, nil
}

// ParseLog parses the given SSLSplit logfile and returns a list of parsed chunks
func ParseLog(logfile string) (chunks []Chunk, err error) {

	f, err := os.Open(logfile)
	if err != nil {
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		c, perr := parseChunk(reader)
		if perr != nil {
			return chunks, perr
		}

		chunks = append(chunks, *c)
	}
}
