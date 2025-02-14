package shares

import (
	"fmt"
	"github.com/SSSaaS/sssa-golang"
	"privileges-management/errors"
	"privileges-management/model"
	"regexp"
	"strings"
	"sync"
)

func CreateSecretsFromCredentials(minimum, createdShares int, credentials model.Credentials) ([]string, error) {
	raw := createCredentialsRawString(credentials)
	return sssa.Create(minimum, createdShares, raw)
}

func createCredentialsRawString(credentials model.Credentials) string {
	return fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)
}

func RetrieveCredentialsFromSecrets(shares []string) (model.Credentials, error) {
	raw, err := sssa.Combine(shares)
	if err != nil {
		return model.Credentials{}, err
	}

	if !isValidRawCredentials(raw) {
		return model.Credentials{}, &errors.WrongRawCredentialsFormatError{}
	}

	return makeCredentialsFromRaw(raw), nil
}

func DistributeSecrets(sharesMapping map[string]string) error {
	var wg sync.WaitGroup

	for approver, share := range sharesMapping {
		wg.Add(1)

		go func(topicSuffix, secretShare string) {
			wg.Done()
			//	kafkaWriter := broker.CreateKafkaWriter(topicSuffix)
			//	err := broker.WriteMessage(kafkaWriter, topicSuffix, secretShare)
			//	if err != nil {
			//		log.Printf("failed to publish event to Kafka: %v\n", err)
			//	}
		}(approver, share)
	}
	wg.Wait()

	return nil
}

func makeCredentialsFromRaw(raw string) model.Credentials {
	parts := strings.Split(raw, ":")
	user, pass := parts[0], parts[1]

	return model.Credentials{Username: user, Password: pass}
}

func isValidRawCredentials(input string) bool {
	regexPattern := `^[a-zA-Z0-9]+:[a-zA-Z0-9]+$`
	match, _ := regexp.MatchString(regexPattern, input)
	return match
}
