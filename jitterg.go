package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

type log struct {
	timestampe string
	id         string
	src        string
	src_port   int
	dest       string
	dest_port  int
	bytes      int
	proto      string
	duration   string
}

var (
	sourceIP        string
	destinationIP   string
	destinationPort int
	protocol        string
	bytes           int
	startTime       string
	duration        time.Duration
	jitter          string
	starttime       time.Time
	jitter_type     string
)

func main() {

	flag.StringVar(&sourceIP, "source", "", "Source IP address")
	flag.IntVar(&bytes, "bytes", 221, "number of bytes")
	flag.StringVar(&destinationIP, "destination", "", "Destination IP address")
	flag.IntVar(&destinationPort, "port", 443, "Destination port")
	flag.StringVar(&protocol, "protocol", "tls", "Protocol")
	flag.StringVar(&startTime, "starttime", "", "Start Time, Example: 2023-10-20T15:10:10")
	flag.DurationVar(&duration, "duration", time.Hour*168, "duration in hours (Default 1 week)")
	flag.StringVar(&jitter, "jitter", "", "Jitter: sleep-percentage -> (e.g 60s-10%)")

	flag.Parse()
	if flag.NFlag() < 1 {
		flag.Usage()
		os.Exit(0)
	}
	if jitter == "" {
		fmt.Println("Error: Use -jitter to provide jitter value")
		return
	}
	if sourceIP == "" {
		sourceIP = generateInternalIP().String()

	}
	if destinationIP == "" {
		destinationIP = generatePublicIP().String()
	}

	layout := "2006-01-02T15:04:05"

	parsedStartTime, err := time.Parse(layout, startTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	starttime = parsedStartTime

	beacon()

}

func parseJitter(jitter string) (float64, float64) {

	//example "60s-10%"

	pattern := `(\d+)s-(\d+)%`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(jitter)
	if len(matches) != 3 {
		fmt.Println("jitter format is not correct")
		os.Exit(0)
	}

	sleep_seconds_str := matches[1]

	if sleep_seconds_str == "" || sleep_seconds_str == "0" {
		fmt.Println("jitter format is not correct")
		os.Exit(0)
	}
	sleep_percentage_str := matches[2]
	if sleep_percentage_str == "" || sleep_percentage_str == "0" {
		fmt.Println("jitter format is not correct")
		os.Exit(0)
	}
	seconds, err1 := strconv.ParseFloat(sleep_seconds_str, 32)
	percentage, err2 := strconv.ParseFloat(sleep_percentage_str, 32)

	//fmt.Println("parsed jitter", seconds, percentage)

	if err1 != nil || err2 != nil {
		fmt.Println("Failed to convert strings to integers")

	}
	return seconds, (float64(percentage) / float64(100))
}

func beacon() {

	time_reached := false
	firstBeaconTime := starttime
	fmt.Println("timestamp id src_ip src_port dest_ip dest_port bytes protocol")
	for time_reached == false {
		log := log{}

		log.id = generateUID()
		log.src = sourceIP
		log.dest = destinationIP
		log.dest_port = destinationPort
		log.proto = protocol
		log.bytes = bytes
		log.src_port = rand.Intn(65534-1024) + 1024

		seconds, percentage := parseJitter(jitter)
		//fmt.Println(seconds, percentage)

		//example 60s sleep with 10% jitter -> 60 * 0.1 = 6
		// that is 60 plus or minus 6 -> sleep randomly between 54 and 66
		sleep_percentage := percentage * seconds // number of seconds to add or substract

		sleep_seconds_min := seconds - sleep_percentage
		sleep_seconds_max := seconds + sleep_percentage

		random_sleep := rand.Intn((int(sleep_seconds_max) - int(sleep_seconds_min))) + int(sleep_seconds_min)

		next_beacon := starttime.Add(time.Duration(random_sleep) * time.Second)

		// adding some network delay
		next_beacon_with_netDelay := next_beacon.Add(time.Millisecond * time.Duration(randomNetworkDelay()))

		log.timestampe = next_beacon_with_netDelay.Format("2006-01-02T15:04:05Z.000")

		starttime = next_beacon
		writelog(log)

		time_reached = next_beacon.After(firstBeaconTime.Add(duration))

	}
}

func writelog(l log) {

	fmt.Printf("%s %s %s %d %s %d %d %s\n", l.timestampe, l.id, l.src, l.src_port, l.dest, l.dest_port, l.bytes, l.proto)

}

func generateUID() string {
	const charset = "abcdefghiJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}

func randomNetworkDelay() int {
	rand.Seed(time.Now().UnixNano())
	delayMilliseconds := rand.Intn(400) + 100
	return delayMilliseconds
}

// Use Biased Triangle Random
func generateBiasedTriangularRandom(min, max float64) float64 {
	// Calculate the mode (the peak of the distribution) closer to the middle.
	mode := (min + max) / 2.0
	modeBias := 0.8 // Adjust this value to control the bias towards the middle.

	// Generate a random value between 0 and 1.
	r := rand.Float64()

	// Calculate the random number following the biased triangular distribution.
	if r < modeBias {
		return min + math.Sqrt(r*(max-min)*(mode-min))
	} else {
		return max - math.Sqrt((1-r)*(max-min)*(max-mode))
	}
}

func generateInternalIP() net.IP {

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1<<24 - 1)
	ip := net.IPv4(10, byte(randomNum>>16&0xFF), byte(randomNum>>8&0xFF), byte(randomNum&0xFF))

	return ip
}

func generatePublicIP() net.IP {
	rand.Seed(time.Now().UnixNano())
	publicRanges := [][]byte{
		{8, 120, 2, 10},
		{17, 64, 4, 20},
		{45, 32, 8, 30},
		{64, 16, 16, 40},
		{129, 8, 32, 50},
		{191, 4, 64, 60},
		{203, 2, 113, 70},
	}

	randomRange := publicRanges[rand.Intn(len(publicRanges))]

	ip := net.IPv4(randomRange[0], randomRange[1], randomRange[2], randomRange[3])

	return ip
}
