package exporter

import (
	"bytes"
	"encoding/json"

	"github.com/olekukonko/tablewriter"
	"sigs.k8s.io/yaml"
)

// LabelsToJSON will return the results in JSON format
func LabelsToJSON(labels []*Label) ([]byte, error) {
	return json.Marshal(labels)
}

// LabelsToYAML will return the results in YAML format
func LabelsToYAML(labels []*Label) ([]byte, error) {
	return yaml.Marshal(labels)
}

// LabelsToTable will return the results in TABLE format
func LabelsToTable(labels []*Label) ([]byte, error) {
	labelRows := make([][]string, 0, len(labels))
	for _, l := range labels {
		labelRows = append(labelRows, []string{l.Name, l.Description, l.Color})
	}

	b := &bytes.Buffer{}
	t := tablewriter.NewWriter(b)
	t.SetHeader([]string{"Name", "Description", "Color"})
	t.AppendBulk(labelRows)
	t.Render()

	return b.Bytes(), nil
}
