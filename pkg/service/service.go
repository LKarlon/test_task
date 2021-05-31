package service

import (
	"fmt"

	"github.com/LKarlon/test_task/pkg/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Conversion interface {
	Convert(file string) (fileOut string, err error)
}

type Service struct {
	Conversion
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Convert(file []byte) (fileOut string, err error){
	t := models.Yaml2Go{}
	err = yaml.Unmarshal(file, &t)
    	if err != nil {
			logrus.Errorf("YAML unmarshall error: %s", err.Error())
			return
       	}
	fileOut = fmt.Sprintf("currency{%s=%d, %s=%d}", t.Currencies[0].Name, t.Currencies[0].Value, t.Currencies[1].Name, t.Currencies[1].Value)
	return
}