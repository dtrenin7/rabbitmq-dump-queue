package main

import (
  "io"
  "fmt"
  "bytes"
  "compress/gzip"
)

type GZipped struct {
  data bytes.Buffer
}

func (g *GZipped) Pack(id string, archived []byte) error {
  if len(id) < 32 {
    return fmt.Errorf("id must be uuid4")
  }
  g.data.WriteString(id) // по этой строке будем искать в кролике
  g.data.WriteByte(0)
  gz := gzip.NewWriter(&g.data)
	if _, err := gz.Write(archived); err != nil {
		return err
	}
  if err := gz.Close(); err != nil {
		return err
	}
  return nil
}

func (g *GZipped) Unpack() (string, []byte, error)  {
  id := ""
  var unpacked bytes.Buffer
  var err error
  if id, err = g.data.ReadString(0); err != nil {
    return id, []byte{}, err
  }
  gz, err := gzip.NewReader(&g.data)
  if err != nil {
    return id, []byte{}, err
  }
  for {
    n, _ := io.Copy(&unpacked, gz)
    // fmt.Printf("\nRead %d\n", n)
    _ = n
    err = gz.Reset(&g.data)
    		if err == io.EOF {
    			break
    		}
  }
  gz.Close()
  return id, unpacked.Bytes(), nil
}
