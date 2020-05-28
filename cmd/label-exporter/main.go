package main

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	exporter "github.com/micnncim/label-exporter"
)

var (
	debug  = kingpin.Flag("debug", "Enable debug mode.").Bool()
	owner  = kingpin.Arg("owner", "Owner of the repository.").Required().String()
	repo   = kingpin.Arg("repo", "Repository whose wanted labels.").Required().String()
	output = kingpin.Flag("output", "Output format. One of: json|yaml|table - default is table").Default("table").Short('o').String()
)

func main() {
	kingpin.Parse()

	client, err := exporter.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	labels, err := client.ListLabels(context.Background(), *owner, *repo)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case "yaml":
		b, err := exporter.LabelsToYAML(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	case "json":
		b, err := exporter.LabelsToJSON(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	case "table":
		b, err := exporter.LabelsToTable(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	default:
		log.Fatal(fmt.Errorf("output format not recognized: %s", *output))
		return
	}
}
