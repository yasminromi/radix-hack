package service

import (
	"context"
	"errors"

	"github.com/yasminromi/radix-hack/model"

	"gopkg.in/olivere/elastic.v6"
)

const (
	indexName = "chat"
	docType   = "log"
	appName   = "back-end"
)

// ElasticService Exported
type ElasticService struct {
	ElasticCLI *elastic.Client
}

// SaveToElastic Exported
func (e *ElasticService) SaveToElastic(ctx context.Context, payload model.Message) error {

	exists, err := e.ElasticCLI.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		res, error := e.ElasticCLI.CreateIndex(indexName).Do(ctx)
		if error != nil {
			return error
		}
		if !res.Acknowledged {
			return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct")
		}
	}

	var user = model.User{
		Name:   payload.User.Name,
		Msisdn: payload.User.Msisdn,
	}

	var message = model.Message{
		Text: payload.Text,
		User: user,
	}

	_, error := e.ElasticCLI.Index().
		Index(indexName).
		Type(docType).
		BodyJson(message).
		Do(ctx)

	if error != nil {
		return error
	}

	return nil
}
