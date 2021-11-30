package object

// SetConsumerID ...
func SetConsumerID(kongID string) ConsumerMutatorFunc {
	return func(accessor ConsumerAccessor) error {
		return accessor.SetID(kongID)
	}
}

// SetConsumerCreatedAt ...
func SetConsumerCreatedAt(createdAt int64) ConsumerMutatorFunc {
	return func(accessor ConsumerAccessor) error {
		return accessor.SetCreatedAt(createdAt)
	}
}

// SetConsumerTags ...
func SetConsumerTags(tags ...string) ConsumerMutatorFunc {
	return func(accessor ConsumerAccessor) error {
		return accessor.SetTags(tags...)
	}
}

// SetConsumerCustomID ...
func SetConsumerCustomID(customID string) ConsumerMutatorFunc {
	return func(accessor ConsumerAccessor) error {
		return accessor.SetCustomID(customID)
	}
}

// SetConsumerUsername ...
func SetConsumerUsername(username string) ConsumerMutatorFunc {
	return func(accessor ConsumerAccessor) error {
		return accessor.SetUsername(username)
	}
}
