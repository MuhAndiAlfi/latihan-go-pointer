package main

import "fmt"

// type Student struct {
// 	ID   int
// 	Name string
// 	GPA  float32
// }

// func (student *Student) graduate() {
// 	student.Name = student.Name + " S.T"

// fmt.Println(student.Name)

// }

// func main() {
// numberA := 5
// numberB := &numberA

// fmt.Println(numberA)
// fmt.Println(numberB)
// fmt.Println(*numberB)

// *numberB = 10

// fmt.Println(*numberB)
// fmt.Println(numberA)

// & referensi
// var numberA int = 5
// var numberB *int = &numberA

// fmt.Println(numberA)
// fmt.Println(numberB)
// fmt.Println(*numberB)

// numberA = 20
// fmt.Println(numberA)
// fmt.Println(*numberB)

// number := 5

// fmt.Println(" alamat memori: ", &number)
// fmt.Println(" nilai awal:", number)

// change(&number, 100)

// fmt.Println(" nilai akhir:", number)
// fmt.Println(" alamat memori: ", &number)

// 	student := Student{1, "andi", 3.27}

// 	fmt.Println(student)

// 	student.graduate()

// 	fmt.Println(student.Name)

// }

// func change(old *int, new int) {
// 	*old = new

// 	fmt.Println(" alamat memori: ", old)
// 	fmt.Println("nilai old: ", *old)

// }

// func graduate(student *Student) {
// 	student.Name = student.Name + " S.T"

// }

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("Mario")
	gamer.AddGame("Fifa 2020")
	gamer.AddGame("Soccer 2020")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
