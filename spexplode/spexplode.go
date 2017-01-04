package spexplode

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/NothNoth/SSLSplitParser/spparser"
)

func Explode(chunk *spparser.Chunk, destFolder string) (err error) {
	var ipLink string
	if chunk.Descriptor.SrcIP > chunk.Descriptor.DestIP {
		ipLink = fmt.Sprintf("%s-%s", chunk.Descriptor.SrcIP, chunk.Descriptor.DestIP)
	} else {
		ipLink = fmt.Sprintf("%s-%s", chunk.Descriptor.DestIP, chunk.Descriptor.SrcIP)
	}

	os.Mkdir(path.Join(destFolder, ipLink), os.ModePerm)

	//Write descriptor
	fName := fmt.Sprintf("%d.txt", chunk.Descriptor.Date)
	f, err := os.Create(path.Join(destFolder, ipLink, fName))
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString(fmt.Sprintf("src_ip: %s\n", chunk.Descriptor.SrcIP))
	f.WriteString(fmt.Sprintf("src_port: %d\n", chunk.Descriptor.SrcPort))
	f.WriteString(fmt.Sprintf("dest_ip: %s\n", chunk.Descriptor.DestIP))
	f.WriteString(fmt.Sprintf("dest_port: %d\n", chunk.Descriptor.DestPort))
	t := time.Unix(chunk.Descriptor.Date, 0)
	f.WriteString(fmt.Sprintf("date: %s\n", t.UTC().Format(time.UnixDate)))
	f.WriteString(fmt.Sprintf("size: %d\n", chunk.Descriptor.Size))

	//Write data
	if len(chunk.Data) != 0 {

		HTTPHeader := []byte{'H', 'T', 'T', 'P'}
		HTTPPOST := []byte{'P', 'O', 'S', 'T'}

		if bytes.HasPrefix(chunk.Data, HTTPHeader) {
			dataReader := bufio.NewReader(bytes.NewReader(chunk.Data))
			resp, _ := http.ReadResponse(dataReader, nil)
			if resp != nil && resp.ContentLength != 0 {
				fName := fmt.Sprintf("%d.http", chunk.Descriptor.Date)
				f, err := os.Create(path.Join(destFolder, ipLink, fName))
				defer f.Close()
				if err != nil {
					return err
				}
				body, err := ioutil.ReadAll(resp.Body)
				f.Write(body)
				resp.Body.Close()

				cookies := resp.Cookies()
				for idx, cook := range cookies {
					fName := fmt.Sprintf("%d.cookie%d", chunk.Descriptor.Date, idx)
					f, err := os.Create(path.Join(destFolder, ipLink, fName))
					defer f.Close()
					if err != nil {
						return err
					}
					f.WriteString(cook.String())
				}
			}
		} else if bytes.HasPrefix(chunk.Data, HTTPPOST) {
			dataReader := bufio.NewReader(bytes.NewReader(chunk.Data))
			req, _ := http.ReadRequest(dataReader)
			if req.ContentLength != 0 {
				fName := fmt.Sprintf("%d.post", chunk.Descriptor.Date)
				f, err := os.Create(path.Join(destFolder, ipLink, fName))
				defer f.Close()
				if err != nil {
					return err
				}
				body, err := ioutil.ReadAll(req.Body)
				f.Write(body)
				req.Body.Close()
			}

			cookies := req.Cookies()
			for idx, cook := range cookies {
				fName := fmt.Sprintf("%d.cookie%d", chunk.Descriptor.Date, idx)
				f, err := os.Create(path.Join(destFolder, ipLink, fName))
				defer f.Close()
				if err != nil {
					return err
				}
				f.WriteString(cook.String())
			}
		}

		fName := fmt.Sprintf("%d.raw", chunk.Descriptor.Date)
		f, err := os.Create(path.Join(destFolder, ipLink, fName))
		defer f.Close()
		if err != nil {
			return err
		}
		f.Write(chunk.Data)

	}

	return
}
