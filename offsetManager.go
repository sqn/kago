package kago

import (
	"github.com/Shopify/sarama"
	"log"
)

type PartitionOffsetManager struct {
	pom sarama.PartitionOffsetManager
}

func (pom *PartitionOffsetManager) MarkOffset(offset int64, ifExactOnce bool) {
	//file
	pom.pom.MarkOffset(offset, "")
}

func (pom *PartitionOffsetManager) ResetOffset(offset int64, ifExactOnce bool) {
	//file
	pom.pom.ResetOffset(offset, "")
}

func (pom *PartitionOffsetManager) NextOffset() (offset int64) {
	offset, _ = pom.pom.NextOffset()
	return offset
}

func (pom *PartitionOffsetManager) Close() {
	pom.pom.AsyncClose()
}

func (pom *PartitionOffsetManager) Errors() <-chan *ConsumerError {
	return pom.pom.Errors()
}

func InitPartitionOffsetManager(addr []string, topic, groupId string, partition int32, conf *Config) (*PartitionOffsetManager, error) {
	client, err := sarama.NewClient(addr, &conf.Config.Config)
	if err != nil {
		log.Println("client create error")
		return nil, err
	}
	defer client.Close()

	offsetManager, err := sarama.NewOffsetManagerFromClient(groupId, client)
	if err != nil {
		log.Println("offsetManager create error")
		return nil, err
	}
	defer offsetManager.Close()

	partitionOffsetManager, err := offsetManager.ManagePartition(topic, partition)
	if err != nil {
		log.Println("partitionOffsetManager create error")
		return nil, err
	}
	var pom = PartitionOffsetManager{
		pom: partitionOffsetManager,
	}
	return &pom, nil
}
