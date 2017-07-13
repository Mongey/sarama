package sarama

import (
	"testing"
)

var (
	deleteTopicsRequestSingleTopic = []byte{
		0, 0, 0, 1,
		0, 1, 'a', // topic string
		0, 0, 0, 0, // timeout int32
	}
)

func TestDeleteTopicsRequest(t *testing.T) {
	reqs := new(DeleteTopicsRequest)
	req := DTopic{}
	req.Topic = "a"
	reqs.DeleteTopicRequests = append(reqs.DeleteTopicRequests, req)
	testRequest(t, "delete single topic", reqs, deleteTopicsRequestSingleTopic)
}
