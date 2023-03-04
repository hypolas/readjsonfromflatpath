package hypolasjsondecomposer

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// TestResolve test different strings format
func TestResolve(t *testing.T) {
	jso, _ := os.ReadFile("test/test.json")
	logf.Info.Println(ReadJSONFromFlatPath("services", jso))
	logf.Info.Println(ReadJSONFromFlatPath("services__1__status", jso))
	logf.Info.Println(ReadJSONFromFlatPath("services__0__nom", jso))
	logf.Info.Println(ReadJSONFromFlatPath("services__0__nom", jso))
	logf.Info.Println(ReadJSONFromFlatPath("services__2__level", jso))
	logf.Info.Println(ReadJSONFromFlatPath("services__2__bad", jso))

	readFile, err := os.Open(logf.LogFile.Name())

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
