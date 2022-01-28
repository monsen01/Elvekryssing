package main

import(
	"fmt"
	"testing"
)

//array med fire inne i seg
var Character_Names[4] string 
var Character_Positions[4] int
//båt på venstre eller høyre side
var BoatLeftSide bool


func main(){
Init()
	pws()
	PutIn("Mann")
	pws()
	CrossRiver()
	pws()
}

func pws(){
	fmt.Println(MakeWorldState())
}
//båten på venstre siden
func Init(){
	BoatLeftSide = true
	//karakter navn
	Character_Names[0] = "Mann"
	Character_Names[1] = "Korn"
	Character_Names[2] = "Kylling"
	Character_Names[3] = "Rev"

	//alle posisjonene er på venstre side 
	//Mann
	Character_Positions[0] = 0
	//Korn
	Character_Positions[1] = 0
	//Kylling
	Character_Positions[2] = 0
	//Rev
	Character_Positions[3] = 0
}
//testing
func ReturnWorldState() string{
	Init()
	return MakeWorldState()
}
// lager en state på verden og blir satt sammen i en string
func MakeWorldState() string{
	rightSide := MakeRightSide()
	boat := MakeBoat()
	leftSide := MakeLeftSide()
	res := leftSide + boat + rightSide
	return res
}
//array med 4 strenger 
//"scanner" posisjonen til karakterene
func MakeRightSide() string{
	var endResult[4]string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 2){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	res := fmt.Sprintf("%s-%s-%s-%s-", endResult[0], endResult[1], endResult[2], endResult[3])
	return res
}

func MakeLeftSide() string{
	var endResult[4]string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 0){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	res := fmt.Sprintf("%s-%s-%s-%s-", endResult[0], endResult[1], endResult[2], endResult[3])
	return res
}
// oppdaterer karakterposisjonen
func PutIn(item string) string{
	for i := 0; i < len(Character_Names); i++{
		if(Character_Names[i] == item){
			Character_Positions[i] = 1
		}
	}
	return MakeWorldState()
}

func MakeBoat() string{
	var endResult[4]string
	var res string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 1){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	if(BoatLeftSide){
		res = fmt.Sprintf(" \\_%s_%s_/________", endResult[0], endResult[1])
	}else{
		res = fmt.Sprintf("________\\_%s_%s_/  ", endResult[0], endResult[1])
	}
	
	return res
}
//oppdaterer state om hvor karakteren er  
func CrossRiver() string{
	BoatLeftSide = !BoatLeftSide
	
	return MakeWorldState()
}
