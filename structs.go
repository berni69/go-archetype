package main

type Database struct {
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Database `yaml:"database"`
}

type configuration struct {
	Config `yaml:"config"`
}
