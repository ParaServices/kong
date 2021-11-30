package plugin

import "github.com/ParaServices/kong/object"

// JWTCredentialMutatorFunc ...
type JWTCredentialMutatorFunc func(accessor JWTCredentialAccessor) error

// SetJWTCredentialKey ...
func SetJWTCredentialKey(key string) JWTCredentialMutatorFunc {
	return func(accessor JWTCredentialAccessor) error {
		return accessor.SetKey(key)
	}
}

// SetJWTCredentialSecret ...
func SetJWTCredentialSecret(secret string) JWTCredentialMutatorFunc {
	return func(accessor JWTCredentialAccessor) error {
		return accessor.SetSecret(secret)
	}
}

// SetJWTCredentialRSAPublicKey ...
func SetJWTCredentialRSAPublicKey(rsapublickey string) JWTCredentialMutatorFunc {
	return func(accessor JWTCredentialAccessor) error {
		return accessor.SetRSAPublicKey(rsapublickey)
	}
}

// SetJWTCredentialAlgorithm ...
func SetJWTCredentialAlgorithm(algorithm string) JWTCredentialMutatorFunc {
	return func(accessor JWTCredentialAccessor) error {
		return accessor.SetAlgorithm(algorithm)
	}
}

// SetJWTCredentialConsumer ...
func SetJWTCredentialConsumer(
	getter object.ConsumerGetter,
) JWTCredentialMutatorFunc {
	return func(accessor JWTCredentialAccessor) error {
		return accessor.SetConsumer(getter)
	}
}

// SetJWTCredentialNewConsumer ...
func SetJWTCredentialNewConsumer(
	newFn object.NewConsumerFunc,
	mutatorFuncs ...object.ConsumerMutatorFunc,
) JWTCredentialMutatorFunc {
	return func(accessor JWTCredentialAccessor) error {
		return accessor.SetNewConsumer(newFn, mutatorFuncs...)
	}
}
