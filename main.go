package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/SurkovIlya/workCashe/workcashe"
)

func main() {
	const ttlMs = uint32(1)
	var WordsforCheck string
	var NoL int
	casheFile, err := os.ReadFile("casheWord.json")
	if err != nil {
		os.Create("casheWord.json")

	}
	var casheWord workcashe.Cashe
	err = json.Unmarshal(casheFile, &casheWord)
	if err != nil {
		fmt.Println("Нет кэша")
	}
	cashe := workcashe.NewCash(ttlMs)

	for _, val := range casheWord.Word {
		cashe.AddWord(val)
	}
	go cashe.Clean()
	fmt.Println("Введите слово для расчета кол-ва букв: ")
	for {

		fmt.Fscan(os.Stdin, &WordsforCheck)
		lowWords := strings.ToLower(WordsforCheck)

		answer, err := cashe.GetWordByID(lowWords)
		if err != nil {
			NoL = len([]rune(lowWords))
			answerForCashe := workcashe.Words{Word: lowWords, NoW: NoL, Lastusedgetime: time.Now()}
			cashe.AddWord(answerForCashe)
			Savecashe, err := json.Marshal(cashe)
			if err != nil {
				fmt.Println(err)
			}

			err = os.WriteFile("./casheWord.json", Savecashe, 0666)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Мы посчитали кол-во букв в слове --->%v<---. Ответ - %v \n", WordsforCheck, NoL)

		} else {
			fmt.Printf("Мы помним что в слове -->%v<-- ответ - %v  \n", answer.Word, answer.NoW)

		}

	}

}
