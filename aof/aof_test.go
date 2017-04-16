package aof

import (
	"testing"
	"os"
	"fmt"
	"io"
)

func TestRead(t *testing.T) {
	file, err := os.Open("test.aof")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	reader := NewBufioReader(file)
	for {
		op, err := reader.ReadOperation()
		if err == io.EOF {
			break
		} else {
			fmt.Printf("cmd:%s, key:%s, args:%s\n",  op.Command, op.Key, op.Arguments)
		}

	}

}
