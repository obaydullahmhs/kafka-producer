package producer

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
)

var messages = []string{
	"His father owns the restaurant.",
	"He is the person to see.",
	"She’s now wearing headphones.",
	"He likes to swim in the lake.",
	"My bandaid wasn't sticky anymore so it fell off on the way to school.",
	"I know that there was a big church here.",
	"I think I could fall asleep really quickly.",
	"The castle is on top of a cliff.",
	"An ENT removed the bug from his ear.",
	"The tables were made of fake wood.",
	"Tom made a big donation to the hospital.",
	"He is respectful to his elders.",
	"It was so cold that no one wanted to go outside.",
	"She was interested in photography but couldn't afford a camera.",
	"He is only about six feet tall.",
	"When she heard they were moving to China, she was distraught.",
	"There used to be big trees around my house.",
	"Good gracious!",
	"Being fashionable is easy.",
	"That's all!",
}

type Producer struct {
	sarama.SyncProducer
	Config *Config
}

func validateConfig(config *Config) error {
	if config.Brokers == nil {
		return errors.New("brokers not set")
	}
	if config.Topic == "" {
		return errors.New("topics not set")
	}

	if config.SecurityProtocol == "" {
		config.SecurityProtocol = "PLAINTEXT"
	}

	return nil
}

func NewProducer() (*Producer, error) {
	config := NewConfig()
	saramaConfig := config.GetConfig()
	saramaConfig.Producer.Return.Successes = true
	if err := validateConfig(config); err != nil {
		return nil, err
	}
	producer, err := sarama.NewSyncProducer(config.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}
	return &Producer{producer, config}, nil
}

func (p *Producer) SendRandomMessages() {
	for i := int64(0); i < p.Config.MaxMessages; i++ {
		key := sarama.StringEncoder(fmt.Sprintf("Random-%d", i))
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := r.Intn(len(messages))
		value := sarama.StringEncoder(messages[index])
		msg := &sarama.ProducerMessage{
			Topic: p.Config.Topic,
			Key:   key,
			Value: value,
		}
		_, _, err := p.SendMessage(msg)
		if err != nil {
			p.Config.Log.Error("failed to send message", err)
		}
	}
}
