package services

import (
	"context"
	"go-srv/internal/adapters/classifier_client"
	htmlparser "go-srv/internal/adapters/html_parser"
	"go-srv/internal/entities"
	"go-srv/internal/repo"
	"log"
	"time"
)

type ResponsesCollector struct {
	classifierCli *classifier_client.ClassifierCli
	repo          *repo.Repo
}

func NewResponsesCollector(classifier *classifier_client.ClassifierCli, repo *repo.Repo) *ResponsesCollector {
	return &ResponsesCollector{
		classifierCli: classifier,
		repo:          repo,
	}
}

func (s *ResponsesCollector) Run(ctx context.Context, firstId int, amount int) {
	for {
		responsesTexts, err := htmlparser.Parse(firstId, amount)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		responses := make([]entities.Response, 0, len(responsesTexts))

		for i, text := range responsesTexts {
			response := entities.Response{
				Id:           firstId + i,
				OriginalText: text,
				Category:     "",
			}

			category, err := s.classifierCli.Predict(ctx, response.OriginalText)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			response.Category = category
			responses = append(responses, response)
		}

		if err := s.repo.SaveResponses(ctx, responses); err != nil {
			log.Println(err.Error())
		}
		time.Sleep(time.Hour * 24) // собираем отзывы раз в сутки
		firstId += amount
	}
}
