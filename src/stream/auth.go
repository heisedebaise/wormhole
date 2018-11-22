package stream

var producers = make(map[string]string)
var consumers = make(map[string]string)

// Producer 注册生产者。
func Producer(auth string, unique string) {
	producers[unique] = auth
}

func getProducer(unique string) string {
	if producer, ok := producers[unique]; ok {
		return producer
	}

	return ""
}

// Consumer 注册消费者。
func Consumer(auth string, unique string) {
	consumers[unique] = auth
}

func getConsumer(unique string) string {
	if consumer, ok := consumers[unique]; ok {
		return consumer
	}

	return ""
}
