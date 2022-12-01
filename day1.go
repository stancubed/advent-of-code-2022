// https://adventofcode.com/2022/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type (
	elfPack struct {
		ID            int
		CalorieItems  []int
		TotalCalories int
	}
)

func loadElves(path string) []elfPack {
	file, _ := os.Open(path)
	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)
	var rawData []string
	for reader.Scan() {
		rawData = append(rawData, reader.Text())
	}
	file.Close()

	var returnSlice []elfPack
	elf := elfPack{
		ID:            1,
		CalorieItems:  []int{},
		TotalCalories: 0,
	}
	for _, line := range rawData {
		calorieInt, err := strconv.Atoi(line)
		if err != nil {
			for _, item := range elf.CalorieItems {
				elf.TotalCalories = elf.TotalCalories + item
			}
			returnSlice = append(returnSlice, elf)
			elf = elfPack{
				ID:            elf.ID + 1,
				CalorieItems:  nil,
				TotalCalories: 0,
			}
			continue
		}
		elf.CalorieItems = append(elf.CalorieItems, calorieInt)
	}
	// fmt.Println(returnMap)
	return returnSlice
}

func getHeroElves(elves []elfPack, howMany int) []elfPack {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories
	})
	return elves[:howMany]
}

func main() {
	elves := loadElves("input.txt")
	heroElves := getHeroElves(elves, 3)
	var topThreeTotalCalories int

	for _, elf := range heroElves {
		topThreeTotalCalories = topThreeTotalCalories + elf.TotalCalories
	}

	fmt.Printf("The Ultimate Hero Elf is Elf %v, who is carrying %v calories!\n", heroElves[0].ID, heroElves[0].TotalCalories)
	fmt.Printf("The Top Three Hero Elves are carrying %v calories! Woah!\n", topThreeTotalCalories)
}
