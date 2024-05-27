package pkg

import "fmt"

type Repo interface {
	Find(key string, defaultValue *string) (string, error)
	Get(key string) (string, error)
}

type KvRepo struct {
	filePath string
	parser   Parser
}

func NewKvRepo(filePath string, parser Parser) *KvRepo {
	return &KvRepo{
		filePath: filePath,
		parser:   parser,
	}
}

var _ Repo = &KvRepo{}

func (r *KvRepo) Find(key string, defaultValue *string) (string, error) {
	items, err := r.parser.ParseKeyValueFile(r.filePath)
	if err != nil {
		return "", err
	}

	value, ok := items[key]
	if !ok {
		return *defaultValue, nil
	}

	return value.Val, nil
}

func (r *KvRepo) Get(key string) (string, error) {
	items, err := r.parser.ParseKeyValueFile(r.filePath)
	if err != nil {
		return "", err
	}

	value, ok := items[key]
	if !ok {
		return "", fmt.Errorf("key doesn't exist")
	}

	return value.Val, nil
}
