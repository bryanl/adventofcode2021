package support

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func ReadData(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	raw := strings.Split(strings.TrimSpace(string(data)), "\n")
	return raw, nil
}

func ParseBinary(in string) int64 {
	out, err := strconv.ParseInt(in, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("%s cannot be converted to binary", in))
	}

	return out
}

func ParseInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(fmt.Sprintf("%s is not an integer", in))
	}

	return out
}

func ParseStringList(in string) []int {
	var out []int

	for _, s := range strings.Split(in, ",") {
		out = append(out, ParseInt(s))
	}

	return out
}

func SetupInput(run func(input []string) error) {
	var input string
	flag.StringVar(&input, "input", "sample.txt", "input")
	flag.Parse()

	data, err := ReadData(input)
	if err != nil {
		log.Fatalf("read data: %v", err)
	}

	if err := run(data); err != nil {
		log.Fatalf("failed: %v", err)
	}
}

func ReadFromDisk(t *testing.T, filename string) []string {
	data, err := ReadData(filename)
	require.NoError(t, err)

	return data
}

func ContainsString(s string, sl []string) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}

	return false
}

func CountString(s string, sl []string) int {
	sum := 0

	for i := range sl {
		if s == sl[i] {
			sum += 1
		}
	}

	return sum
}
