package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/integrii/flaggy"
	"github.com/sirupsen/logrus"
)

var topicRegex = regexp.MustCompile(`^[a-zA-Z0-9-_]+:.*$`)

func main() {
	var (
		file              = "topic-list.txt"
		hoursToAccountFor = 175
	)

	flaggy.String(&file, "", "standups", "file containing a non-unique list of topics that have been worked on")
	flaggy.Int(&hoursToAccountFor, "", "hours", "hours that need to be accounted for")
	flaggy.Parse()

	account(file, hoursToAccountFor)
}

func account(file string, hours int) {

	// read the file
	bt, err := ioutil.ReadFile(file)
	logAndExitOnError(err, fmt.Sprintf("couldn't read file %s: %s", file, err), 1)
	lines := strings.Split(string(bt), "\n")

	// fine relevant lines
	topicLines := filterForTopicLines(lines)

	// extract the actual topic
	topics := extractTopicCounts(topicLines)

	// calculate hours for each topic
	topicHours := scaleTopicsToHours(topics, float64(hours))
	fmt.Println(marshalTopicHours(roundTopicHours(topicHours)))

}

func filterForTopicLines(lines []string) []string {
	topicLines := []string{}
	for _, line := range lines {
		if topicRegex.MatchString(line) {
			topicLines = append(topicLines, line)
		}
	}
	return topicLines
}

func extractTopicCounts(topicLines []string) map[string]int {
	topics := make(map[string]int)
	for _, line := range topicLines {
		topic := strings.Split(line, ":")[0]
		if _, ok := topics[topic]; ok {
			topics[topic]++
		} else {
			topics[topic] = 1
		}
	}
	return topics
}

func marshalTopicHours(topicHours map[string]int) string {
	bt, err := yaml.Marshal(topicHours)
	logAndExitOnError(err, fmt.Sprintf("couldn't marshal topics: %s", err), 2)
	return string(bt)
}

func roundTopicHours(topicHours map[string]float64) map[string]int {

	roundedTopicHours := make(map[string]int)
	for topic, hours := range topicHours {
		roundedTopicHours[topic] = int(hours)
	}
	return roundedTopicHours
}

func scaleTopicsToHours(topics map[string]int, hours float64) map[string]float64 {
	taskSum := 0
	for _, count := range topics {
		taskSum += count
	}

	hoursPerTask := hours / float64(taskSum)

	topicHours := make(map[string]float64)

	for task, count := range topics {
		topicHours[task] = float64(count) * hoursPerTask
	}
	return topicHours
}

func logAndExitOnError(err error, msg string, code int) {
	if err != nil {
		logrus.Error(fmt.Sprintf("%s: %s", msg, err))
		os.Exit(code)
	}
}
