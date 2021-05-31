package models
type Yaml2Go struct {
	Currencies []Currencies `yaml:"currencies"`
}

// Currencies
type Currencies struct {
	Name  string `yaml:"name"`
	Value int    `yaml:"value"`
}