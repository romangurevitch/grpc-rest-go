package dictionary

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

var defaultWords = []string{"hello", "goodbye", "simple", "list", "search", "filter", "yes", "no"}

func New() (Dictionary, error) {
	dic := &dictionary{words: map[string]int64{}}
	err := loadDefault(dic)
	if err != nil {
		return nil, err
	}
	return dic, nil
}

// A simple dictionary
type Dictionary interface {
	Find(word string) (bool, error)
	Add(word string) error
	Delete(word string) error
	MostPopular() (map[string]int64, error)
}

func loadDefault(dic Dictionary) error {
	for _, word := range defaultWords {
		err := dic.Add(word)
		if err != nil {
			return err
		}
	}
	return nil
}

type dictionary struct {
	// a map between word and search count
	words map[string]int64
	// lock to prevent race conditions
	lock sync.Mutex
}

func (dic *dictionary) Find(word string) (bool, error) {
	dic.lock.Lock()
	defer dic.lock.Unlock()

	// support case insensitivity
	lowerCaseWord := strings.ToLower(word)
	counter, exists := dic.words[lowerCaseWord]
	if !exists {
		return false, nil
	}

	dic.words[lowerCaseWord] = counter + 1
	return true, nil
}

func (dic *dictionary) Add(word string) error {
	dic.lock.Lock()
	defer dic.lock.Unlock()

	// support case insensitivity
	lowerCaseWord := strings.ToLower(word)
	_, exists := dic.words[lowerCaseWord]
	if exists {
		return fmt.Errorf("error adding word: %s, already exists", lowerCaseWord)
	}

	dic.words[lowerCaseWord] = 0
	return nil
}

func (dic *dictionary) Delete(word string) error {
	dic.lock.Lock()
	defer dic.lock.Unlock()

	// support case insensitivity
	lowerCaseWord := strings.ToLower(word)
	_, exists := dic.words[lowerCaseWord]
	if !exists {
		return fmt.Errorf("error removing word: %s, the word does not exists", lowerCaseWord)
	}

	delete(dic.words, lowerCaseWord)
	return nil
}

func (dic *dictionary) MostPopular() (map[string]int64, error) {
	dic.lock.Lock()
	defer dic.lock.Unlock()

	var allWords []pair
	for k, v := range dic.words {
		allWords = append(allWords, pair{key: k, value: v})
	}
	sort.Slice(allWords, func(i, j int) bool {
		return allWords[i].value > allWords[j].value
	})
	popular := make(map[string]int64)
	for i := 0; i < min(len(allWords), 5); i++ {
		popular[allWords[i].key] = allWords[i].value
	}
	return popular, nil
}

type pair struct {
	key   string
	value int64
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
