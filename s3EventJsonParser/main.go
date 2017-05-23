package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Event struct {
	Type             string  `json:"Type,omitempty"`
	MessageId        string  `json:"MessageId,omitempty"`
	TopicArn         string  `json:"TopicArn,omitempty"`
	Subject          string  `json:"Subject,omitempty"`
	Timestamp        string  `json:"Timestamp,omitempty"`
	SignatureVersion string  `json:"SignatureVersion,omitempty"`
	SigningCertURL   string  `json:"SigningCertURL,omitempty"`
	UnsubscribeURL   string  `json:"UnsubscribeURL,omitempty"`
	Message          Message `json:"Message,omitempty"`
}

type Message struct {
	Records []Records `json:"Records,omitempty"`
}

type Records struct {
	S3 S3 `json:"S3,omitempty"`
}
type S3 struct {
	S3SchemaVersion string `json:"s3SchemaVersion,omitempty"`
	ConfigurationId string `json:"configurationId,omitempty"`
	Bucket          Bucket `json:"bucket,omitempty"`
	Object          Object `json:"object",omitempty`
}
type Bucket struct {
	Name          string        `json:"name,omitempty"`
	Arn           string        `json:"arn,omitempty"`
	OwnerIdentity OwnerIdentity `json:"owneridentity,omitempty"`
}

type OwnerIdentity struct {
	PrincipalId string `json:principalid,omitempty`
}

type Object struct {
	Key       string `json:key`
	Size      int    `json:size`
	Etag      string `json:etag`
	Sequencer string `json:sequencer`
}

func main() {
	var str Event
	b, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("error in reading file", err)
	}
	err = json.Unmarshal(b, &str)
	if err != nil {
		fmt.Println("Error is", err)
	}
	fmt.Println("Value is", str.Message.Records[0].S3.Object.Etag)
}
