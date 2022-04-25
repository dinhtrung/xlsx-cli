package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("<===========================================>")
	fmt.Println("Welcome to yellowpages I/EF gate")
	fmt.Println("<===========================================>\n")
	var name string
	fmt.Println("What is your name???\n")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Hello %s!\n",name)
	time.Sleep(1 * time.Second)
	fmt.Println("What can i help you??? Choose a number to make your choice!!!\n")
	fmt.Println("1. Convert XLSX file to CSV file")
	fmt.Println("2. Convert CSV file to XLSX file")
	fmt.Println("3. Convert XLSX file to JSON file")
	fmt.Println("4. Convert JSON file to XLSX file")
	fmt.Println("5. Convert CSV file to JSON file")
	fmt.Println("6. Convert JSON file to CSV file\n")
	var action int
	time.Sleep(1 * time.Second)
	fmt.Print("I want number: ")
	fmt.Scanf("%d",&action)
	time.Sleep(1 * time.Second)
	for{
		switch action {
		case 1:
			fmt.Println("You Chose number 1 which is Convert XLSX file to CSV file")
			fmt.Println("Here are some tutorial:\nRun command: ./xlsx2csv [flag option] <xlsx file>")
			fmt.Println("Example: ./xlsx2csv -i 0 -o outputfile.csv inputfile.xlsx")
			fmt.Println("-i stand for the sheet of the excel file you want to convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
		case 2:
			fmt.Println("You Chose number 2 which is Convert CSV file to XLSX file")
			fmt.Println("Here are some tutorial:\nRun command: ./csv2xlsx -i=<CSV Input File> -o=<XLSX Output File> -d=<Delimiter>")
			fmt.Println("Example: ./csv2xlsx -i <csv file> -o <xlsx file> -d ';'")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
		case 3:
			fmt.Println("You Chose number 3 which is Convert XLSX file to JSON file")
			fmt.Println("Here are some tutorial:\n Run command: ./xlsx2json2xlsx -i=<file.xlsx> -o=<file.json>")
			fmt.Println("Example: ./xlsx2json2xlsx -o outputfile.csv -i inputfile.xlsx")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
		case 4:
			fmt.Println("You Chose number 4 which is Convert JSON file to XLSX file")
			fmt.Println("Here are some tutorial:\nRun command: ./xlsx2json2xlsx -i=<file.json> -o=<file.xlsx>")
			fmt.Println("Example: ./xlsx2json2xlsx -o outputfile.csv -i inputfile.xlsx")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
		case 5:
			fmt.Println("You Chose number 5 which is Convert CSV file to JSON file")
			fmt.Println("Here are some tutorial:\nRun command: ./csv2json -path=path/to/you/csv/file")
			fmt.Println("Example : ./csv2json -i=/home/csv/file.csv")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
		case 6:
			fmt.Println("You chose number 6 which is Convert JSON file to CSV file")
			fmt.Println("Here are some tutorial:\nRun command: ./json2csv -i=<file.json> -o=<file.csv>")
			fmt.Println("Example: ./json2csv -i filetest.json  -o filetest.csv -url http:...")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
		default:
			fmt.Printf("Wrong one , there is no number %d!!! Chose again\n",action);
			fmt.Scanf("%d",&action)
		}
		if action == 1 {
			time.Sleep(1 * time.Second)
			fmt.Println("You chose number 1 which is Convert XLSX file to CSV file")
			fmt.Println("Here are some tutorial:\nRun command: ./xlsx2csv [flag option] <xlsx file>")
			fmt.Println("Example: ./xlsx2csv -i 0 -o outputfile.csv inputfile.xlsx")
			fmt.Println("-i stand for the sheet of the excel file you want to convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
			break
		}else if action ==2{
			time.Sleep(1 * time.Second)
			fmt.Println("You Chose number 2 which is Convert CSV file to XLSX file")
			fmt.Println("Here are some tutorial:\nRun command: ./csv2xlsx -i=<CSV Input File> -o=<XLSX Output File> -d=<Delimiter>")
			fmt.Println("Example: ./csv2xlsx -i <csv file> -o <xlsx file> -d ';'")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
			break
		}else if action==3{
			time.Sleep(1 * time.Second)
			fmt.Println("You Chose number 3 which is Convert XLSX file to JSON file")
			fmt.Println("Here are some tutorial:\n Run command: ./xlsx2json2xlsx -i=<file.xlsx> -o=<file.json>")
			fmt.Println("Example: ./xlsx2json2xlsx -o outputfile.csv -i inputfile.xlsx")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
			break
		}else if action==4{
			time.Sleep(1 * time.Second)
			fmt.Println("You Chose number 4 which is Convert JSON file to XLSX file")
			fmt.Println("Here are some tutorial:\nRun command: ./xlsx2json2xlsx -i=<file.json> -o=<file.xlsx>")
			fmt.Println("Example: ./xlsx2json2xlsx -o outputfile.csv -i inputfile.xlsx")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
			break

		}else if action==5{
			time.Sleep(1 * time.Second)
			fmt.Println("You Chose number 5 which is Convert CSV file to JSON file")
			fmt.Println("Here are some tutorial:\nRun command: ./csv2json -path=path/to/you/csv/file")
			fmt.Println("Example : ./csv2json -i=/home/csv/file.csv")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
			break
		}else if action==6{
			time.Sleep(1 * time.Second)
			fmt.Println("You Chose number 6 which is Convert JSON file to CSV file")
			fmt.Println("Here are some tutorial:\nRun command: ./json2csv -i=<file.json> -o=<file.csv>")
			fmt.Println("Example: ./json2csv -i filetest.json  -o filetest.csv -url http:...")
			fmt.Println("-i stand for the file you want to be convert, -o stand for the file you want after converted\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("Goodluck! %s\n",name)
			break
		}
	}

} 