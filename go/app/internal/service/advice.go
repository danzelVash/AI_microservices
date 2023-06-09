package service

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"microservices/app/libraries/grpc_client"
	"microservices/app/libraries/logging"
	"os"
)

const (
	prompt        = "Please write me recommendations on how to cope or what is better to do or not do when you experience an emotion of"
	sorryError    = "Извините, во время генерации советов произошла ошибка"
	promptRussian = "Дай мне пожалуйста основные рекомендации что нужно делать, а что не нужно делать во время того когда ты испытываешь эмоцию "
)

type Adviser struct {
	logger *logging.Logger
}

func NewAdviser(logger *logging.Logger) *Adviser {
	return &Adviser{
		logger: logger,
	}
}

func (a *Adviser) GetAdvise(ctx context.Context, req *grpc_client.GetAdviceRequest) ([]string, []string, error) {
	c := context.WithValue(ctx, "logger", a.logger)
	description, err := grpc_client.TellAboutPhoto(c, req)
	if err != nil {
		a.logger.Errorf("error while getting response from TellAboutPhoto: %s", err.Error())
		return nil, nil, err
	}

	content := make([]string, 0, len(description.MediaTellRespArr))
	emotions := make([]string, 0, len(description.MediaTellRespArr))

	client := openai.NewClient(os.Getenv("OPEN_AI_SECRET"))
	for _, media := range description.MediaTellRespArr {
		if len(media.Data) == 0 {
			emotions = append(emotions, "")
			content = append(content, media.Description)
			continue
		}

		emotions = append(emotions, media.Description)

		promptReady := promptRussian + media.Description
		resp, err := client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: promptReady,
					},
				},
			},
		)

		if err != nil {
			a.logger.Errorf("error while trying create chat completion: %s", err.Error())
			content = append(content, sorryError)
			continue
		}

		content = append(content, resp.Choices[0].Message.Content)

	}
	// TODO сделать проверку на существование в цикле для каждого фото и записать в структура ответа(создать ее)

	return content, emotions, nil
}
