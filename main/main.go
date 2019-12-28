package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
)

func parseSyscallDuration(line string) (duration time.Duration, err error) {
	l := strings.LastIndex(line, "<")
	r := strings.LastIndex(line, ">")
	if l < 0 || r < 0 || l+1 >= r {
		err = errors.New("could not find duration tag")
		return
	}
	durationStr := line[l+1 : r]
	durationSec, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		err = fmt.Errorf("could not parse duration '%s': %v", durationStr, err)
		return
	}
	duration = time.Duration(durationSec*1e9) * time.Nanosecond
	return
}

func main() {
	flag.CommandLine.SortFlags = false
	argHelp := flag.BoolP(
		"help", "h", false, "Just print help message and exit",
	)
	argMinDuration := flag.DurationP(
		"duration", "d", time.Second, "Minimum syscall duration to pass the filter",
	)
	flag.Parse()
	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Unknown argument: '%s'\n", flag.Arg(0))
		flag.Usage()
		return
	}
	if *argHelp {
		flag.Usage()
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	minDuration := *argMinDuration
	for scanner.Scan() {
		line := scanner.Text()
		duration, err := parseSyscallDuration(line)
		if err != nil || duration < minDuration {
			continue
		}
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
