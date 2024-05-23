package workcashe

import (
	"fmt"
	"time"
)

const WordCount = 1000

type Words struct {
	Word           string
	NoW            int
	Lastusedgetime time.Time
}

type Cashe struct {
	Word map[string]Words
	TTL  uint32
}

func NewCash(ttlMs uint32) *Cashe {
	Words := make(map[string]Words, WordCount)

	return &Cashe{
		Word: Words,
		TTL:  ttlMs,
	}

}

func (c *Cashe) GetWordByID(word string) (Words, error) {
	if wordValue, ok := c.Word[word]; ok {

		return wordValue, nil
	} else {
		return Words{}, fmt.Errorf("Asd")
	}
}
func (c *Cashe) AddWord(word Words) {
	if _, ok := c.Word[word.Word]; ok {
		c.Word[word.Word] = word
	} else {
		c.Word[word.Word] = word
	}
}
func (c *Cashe) Clean() {
	for {
		now := time.Now()
		for _, val := range c.Word {
			timeUse := val.Lastusedgetime
			timecheck := timeUse.Add(time.Duration(c.TTL) * time.Minute)
			if int(now.Sub(timecheck)) > 0 {
				delete(c.Word, val.Word)
			}

		}
		time.Sleep(5 * time.Second)
	}

}
