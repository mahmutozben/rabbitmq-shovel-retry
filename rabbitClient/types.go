package rabbitClient

import "fmt"

type MoveQueueRequest struct {
	Component string    `json:"component"`
	Vhost     string    `json:"vhost"`
	Name      string    `json:"name"`
	Value     QueueBody `json:"value"`
}

type QueueBody struct {
	SrcURI                string `json:"src-uri"`
	SrcQueue              string `json:"src-queue"`
	SrcProtocol           string `json:"src-protocol"`
	SrcPrefetchCount      int    `json:"src-prefetch-count"`
	SrcDeleteAfter        string `json:"src-delete-after"`
	DestProtocol          string `json:"dest-protocol"`
	DestURI               string `json:"dest-uri"`
	DestAddForwardHeaders bool   `json:"dest-add-forward-headers"`
	AckMode               string `json:"ack-mode"`
	DestQueue             string `json:"dest-queue"`
}

type RabbitMqQueueInfo struct {
	ConsumerDetails []interface{} `json:"consumer_details"`
	Arguments GoArguments `json:"arguments"`
	AutoDelete bool `json:"auto_delete"`
	BackingQueueStatus GoBackingQueueStatus `json:"backing_queue_status"`
	ConsumerUtilisation interface{} `json:"consumer_utilisation"`
	Consumers int `json:"consumers"`
	Deliveries []interface{} `json:"deliveries"`
	Durable bool `json:"durable"`
	EffectivePolicyDefinition GoEffectivePolicyDefinition `json:"effective_policy_definition"`
	Exclusive bool `json:"exclusive"`
	ExclusiveConsumerTag interface{} `json:"exclusive_consumer_tag"`
	GarbageCollection GoGarbageCollection `json:"garbage_collection"`
	HeadMessageTimestamp int `json:"head_message_timestamp"`
	IdleSince string `json:"idle_since"`
	Incoming []interface{} `json:"incoming"`
	Memory int `json:"memory"`
	MessageBytes int `json:"message_bytes"`
	MessageBytesPagedOut int `json:"message_bytes_paged_out"`
	MessageBytesPersistent int `json:"message_bytes_persistent"`
	MessageBytesRAM int `json:"message_bytes_ram"`
	MessageBytesReady int `json:"message_bytes_ready"`
	MessageBytesUnacknowledged int `json:"message_bytes_unacknowledged"`
	MessageStats GoMessageStats `json:"message_stats"`
	Messages int `json:"messages"`
	MessagesDetails GoMessagesDetails `json:"messages_details"`
	MessagesPagedOut int `json:"messages_paged_out"`
	MessagesPersistent int `json:"messages_persistent"`
	MessagesRAM int `json:"messages_ram"`
	MessagesReady int `json:"messages_ready"`
	MessagesReadyDetails GoMessagesReadyDetails `json:"messages_ready_details"`
	MessagesReadyRAM int `json:"messages_ready_ram"`
	MessagesUnacknowledged int `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails GoMessagesUnacknowledgedDetails `json:"messages_unacknowledged_details"`
	MessagesUnacknowledgedRAM int `json:"messages_unacknowledged_ram"`
	Name string `json:"name"`
	Node string `json:"node"`
	OperatorPolicy interface{} `json:"operator_policy"`
	Policy string `json:"policy"`
	RecoverableSlaves []string `json:"recoverable_slaves"`
	Reductions int `json:"reductions"`
	ReductionsDetails GoReductionsDetails `json:"reductions_details"`
	SingleActiveConsumerTag interface{} `json:"single_active_consumer_tag"`
	SlaveNodes []string `json:"slave_nodes"`
	State string `json:"state"`
	SynchronisedSlaveNodes []string `json:"synchronised_slave_nodes"`
	Type string `json:"type"`
	Vhost string `json:"vhost"`
}

type GoArguments struct {
	XQueueType string `json:"x-queue-type"`
}

type GoBackingQueueStatus struct {
	AvgAckEgressRate float64 `json:"avg_ack_egress_rate"`
	AvgAckIngressRate float64 `json:"avg_ack_ingress_rate"`
	AvgEgressRate float64 `json:"avg_egress_rate"`
	AvgIngressRate float64 `json:"avg_ingress_rate"`
	Delta []interface{} `json:"delta"`
	Len int `json:"len"`
	MirrorSeen int `json:"mirror_seen"`
	MirrorSenders int `json:"mirror_senders"`
	Mode string `json:"mode"`
	NextSeqID int `json:"next_seq_id"`
	Q1 int `json:"q1"`
	Q2 int `json:"q2"`
	Q3 int `json:"q3"`
	Q4 int `json:"q4"`
	TargetRAMCount string `json:"target_ram_count"`
}

type GoEffectivePolicyDefinition struct {
	HaMode string `json:"ha-mode"`
	HaParams int `json:"ha-params"`
	HaSyncBatchSize int `json:"ha-sync-batch-size"`
	HaSyncMode string `json:"ha-sync-mode"`
}

type GoGarbageCollection struct {
	FullsweepAfter int `json:"fullsweep_after"`
	MaxHeapSize int `json:"max_heap_size"`
	MinBinVheapSize int `json:"min_bin_vheap_size"`
	MinHeapSize int `json:"min_heap_size"`
	MinorGcs int `json:"minor_gcs"`
}

type GoAckDetails struct {
	Rate float64 `json:"rate"`
}

type GoDeliverDetails struct {
	Rate float64 `json:"rate"`
}

type GoDeliverGetDetails struct {
	Rate float64 `json:"rate"`
}

type GoDeliverNoAckDetails struct {
	Rate float64 `json:"rate"`
}

type GoGetDetails struct {
	Rate float64 `json:"rate"`
}

type GoGetEmptyDetails struct {
	Rate float64 `json:"rate"`
}

type GoGetNoAckDetails struct {
	Rate float64 `json:"rate"`
}

type GoPublishDetails struct {
	Rate float64 `json:"rate"`
}

type GoRedeliverDetails struct {
	Rate float64 `json:"rate"`
}

type GoMessageStats struct {
	Ack int `json:"ack"`
	AckDetails GoAckDetails `json:"ack_details"`
	Deliver int `json:"deliver"`
	DeliverDetails GoDeliverDetails `json:"deliver_details"`
	DeliverGet int `json:"deliver_get"`
	DeliverGetDetails GoDeliverGetDetails `json:"deliver_get_details"`
	DeliverNoAck int `json:"deliver_no_ack"`
	DeliverNoAckDetails GoDeliverNoAckDetails `json:"deliver_no_ack_details"`
	Get int `json:"get"`
	GetDetails GoGetDetails `json:"get_details"`
	GetEmpty int `json:"get_empty"`
	GetEmptyDetails GoGetEmptyDetails `json:"get_empty_details"`
	GetNoAck int `json:"get_no_ack"`
	GetNoAckDetails GoGetNoAckDetails `json:"get_no_ack_details"`
	Publish int `json:"publish"`
	PublishDetails GoPublishDetails `json:"publish_details"`
	Redeliver int `json:"redeliver"`
	RedeliverDetails GoRedeliverDetails `json:"redeliver_details"`
}

type GoMessagesDetails struct {
	Rate float64 `json:"rate"`
}

type GoMessagesReadyDetails struct {
	Rate float64 `json:"rate"`
}

type GoMessagesUnacknowledgedDetails struct {
	Rate float64 `json:"rate"`
}

type GoReductionsDetails struct {
	Rate float64 `json:"rate"`
}

func createRequest(queueName string) MoveQueueRequest {
	errorQueue := fmt.Sprintf("%s_error",queueName)
	name := fmt.Sprintf("Move from %s",errorQueue)
	req := &MoveQueueRequest{
		Component: "shovel",
		Vhost: "/",
		Name: name,
		Value: QueueBody{
			AckMode: "on-confirm",
			DestProtocol: "amqp091",
			SrcProtocol: "amqp091",
			DestURI: "amqp:///%2F",
			SrcURI: "amqp:///%2F",
			SrcQueue:errorQueue,
			SrcDeleteAfter: "queue-length",
			DestAddForwardHeaders: false,
			DestQueue: queueName,
			SrcPrefetchCount: 50,
		},
	}
	return *req
}
