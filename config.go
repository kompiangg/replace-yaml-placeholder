package main

import (
	"bytes"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type config struct {
	FooBar foobar `yaml:"foobar"`
}

type foobar struct {
	FooName string `yaml:"foomeow"`
}

type availableData struct {
	FooConfig foo `yaml:"foo"`
}

type foo struct {
	Name string `yaml:"name"`
}

func loadConfig() config {
	return config{}
}

func (c *config) loadConfigFromYaml() ([]byte, error) {
	availableData, err := loadAvailableData("available_data.yaml")
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("config.tmpl").ParseFiles("./config.tmpl")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, availableData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func loadAvailableData(fileName string) (availableData, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return availableData{}, err
	}

	fs, err := os.Open(fileName)
	if err != nil {
		return availableData{}, err
	}

	defer fs.Close()

	var data availableData
	err = yaml.NewDecoder(fs).Decode(&data)
	if err != nil {
		return availableData{}, err
	}

	return data, nil
}
