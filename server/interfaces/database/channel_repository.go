package database

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type ChannelRepository struct {
	mongoHandler   MongoHandler
	collectionName string
}

func NewChannelRepository(mongoHandler MongoHandler) *ChannelRepository {
	channelRepository := ChannelRepository{
		mongoHandler:   mongoHandler,
		collectionName: "Channels",
	}
	return &channelRepository
}

func (repository *ChannelRepository) Create(channel domain.Channel) (domain.Channel, error) {
	doc := []KV{
		{"name", channel.Name},
		{"concept", channel.Concept},
	}
	id, err := repository.mongoHandler.Insert(repository.collectionName, doc)
	channel.Id = id
	return channel, err
}

/*
func (repository *ChannelRepository) Update(channel domain.Channel) (domain.Channel, error) {
	query := []KV{
		{"_id", channel.Id},
	}
	update := []KV{
		{"channel_id", channel.ChannelId},
		{"name", channel.ChannelName},
		{"concept", channel.ChannelConcept},
	}
	_, err := repository.mongoHandler.Update(repository.collectionName, query, update)
	return channel, err
}

func (repository *ChannelRepository) Delete(id string) error {
	query := []KV{
		{"_id", id},
	}
	_, err := repository.mongoHandler.Delete(repository.collectionName, query)
	return err
}

func (repository *ChannelRepository) GetAll() ([]domain.collectionName, error) {
	query := []KV{}
	raw, err := repository.mongoHandler.Find(repository.collectionName, query)
	if err != nil {
		return nil, err
	}
	channels := []domain.Channel{}
	for _, kv := range raw {
		for _, u := range kv {
			channel := domain.User{}
			switch u.key {
			case "channel_id":
				channel.ChannelId = u.value.(string)
			case "channel_name":
				channel.ChannelName = u.value.(string)
			case "channel_concept":
				channel.ChannelConcept = u.value.(string)
			}
			channels = append(channels, channel)
		}
	}
	return channels, nil
}
*/
