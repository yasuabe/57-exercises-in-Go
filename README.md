# *Exercises for Programmers* in Go

## Overview
A personal project to get started with Go by solving the exercises from the book *Exercises for Programmers* in Go.

### Solved Exercises
#### Chapter 2: Input, Processing, and Output
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex01 [x] [Saying Hello](ex01/main.go)
| Ex02 [x] [Counting the Number of Characters](ex02/main.go)
| Ex03 [x] [Printing Quotes](ex03/main.go)
| Ex04 [x] [Mad Lib](ex04/main.go)
| Ex05 [x] [Simple Math](ex05/main.go)
| Ex06 [x] [Retirement Calculator](ex06/main.go)
#### Chapter 3: Calculations
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex07 [x] [Area of a Rectangular Room](ex07/main.go)
| Ex08 [x] [Pizza Party](ex08/main.go)
| Ex09 [x] [Paint Calculator](ex09/main.go)
| Ex10 [x] [Self-Checkout](ex10/main.go)
| Ex11 [x] [Currency Conversion](ex11/main.go)
| Ex12 [x] [Computing Simple Interest](ex12/main.go)
| Ex13 [x] [Determining Compound Interest](ex13/main.go)
#### Chapter 4: Making Decisions
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex14 [x] [Tax Calculator](ex14/main.go)
| Ex15 [x] [Password Validation](ex15/main.go)
| Ex16 [x] [Legal Driving Age](ex16/main.go)
| Ex17 [x] [Blood Alcohol Calculator](ex17/main.go)
| Ex18 [x] [Temperature Converter](ex18/main.go)
| Ex19 [x] [BMI Calculator](ex19/main.go)
| Ex20 [x] [Multistate Sales Tax Calculator](ex20/main.go)
| Ex21 [x] [Numbers to Names](ex21/main.go)
| Ex22 [x] [Comparing Numbers](ex22/main.go)
| Ex23 [x] [Troubleshooting Car Issues](ex23/main.go)
#### Chapter 5: Functions
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex24 [x] [Anagram Checker](ex24/main.go)                 | exp |
| Ex25 [x] [Password Strength Indicator](ex25/main.go)     |     |
| Ex26 [x] [Months to Pay Off a Credit Card](ex26/main.go) |     |
| Ex27 [x] [Validating Inputs](ex27/main.go)               |     |
#### Chapter 6: Repetition
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex28 [x] [Adding Numbers](ex28/main.go)
| Ex29 [x] [Handling Bad Input](ex29/main.go)
| Ex30 [x] [Multiplication Table](ex30/main.go)
| Ex31 [x] [Karvonen Heart Rate](ex31/main.go)
| Ex32 [x] [Guess the Number Game](ex32/main.go)
#### Chapter 7: Data Structures
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex33 [x] [Magic 8 Ball](ex33/main.go)
| Ex34 [x] [Employee List Removal](ex34/main.go)
| Ex35 [x] [Picking a Winner](ex35/main.go)
| Ex36 [x] [Computing Statistics](ex36/main.go)
| Ex37 [x] [Password Generator](ex37/main.go)
| Ex38 [x] [Filtering Values](ex39/main.go)
| Ex39 [x] [Sorting Records](ex39/main.go)
| Ex40 [x] [Filtering Records](ex40/main.go)
#### Chapter 8: Working with Files
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex41 [x] [Name Sorter](ex41/main.go)
| Ex42 [x] [Parsing a Data File](ex42/main.go)
| Ex43 [x] [Website Generator](ex43/main.go)
| Ex44 [x] [Product Search](ex44/main.go)
| Ex45 [x] [Word Finder](ex45/main.go)
| Ex46 [x] [Word Frequency Finder](ex46/main.go)
#### Chapter 9: Working with External Services
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex47 [x] [Who’s in Space?](ex47/main.go)                |      |     
| Ex48 [x] [Grabbing the Weather](ex48/main.go)           |      |          
| Ex49 [x] [Flickr Photo Search](ex49/main.go)            | fyne, resize | GUI, goroutine       |        
| Ex50 [x] [Movie Recommendations](ex50/main.go)          |      |           
| Ex51 [x] [Pushing Notes to Firebase](ex51/main.go)      |      | Firebase Realtime Database              
| Ex52 [x] [Creating Your Own Time Service](ex52/main.go) |      | REST Server & Client |
#### Chapter 10: Full Programs
| Exercise | dependencies | Memo  |
| -------- | -------------| ----- |
| Ex53 [x] [Todo List](ex53/main.go)          | redis                  | Redis |
| Ex54 [x] [URL Shortener](ex54/main.go)      | mux, redis             | Web App + Redis|
| Ex55 [x] [Text Sharing](ex55/main.go)       | mongo-driver, uuid, context| Web App + Mongo|
| Ex56 [x] [Tracking Inventory](ex56/main.go) |                        | Web App |
| Ex57 [x] [Trivia App](ex57/main.go)         |                        | |

※ [x] Completed, [ ] Pending

## Technologies Used
- go1.24.2
### Dependency
- fyne.io/fyne/v2
- golang.org/x/exp
- golang.org/x/net/context
- github.com/go-redis/redis/v8
- github.com/gorilla/mux
- github.com/google/uuid
- github.com/nfnt/resize
- go.mongodb.org/mongo-driver

## How to Run
Run the following directly under the project.
```
$ go run ./ex[nn]
```
Individual details are described in the README.md file of each respective folder.

### Example
```
$ go run ./ex07
What is the length of the room in feet? 15
What is the width of the room in feet? 20
You entered dimensions of 15 feet by 20 feet.
The area is
300 square feet
27.871 square meters
```

## Notes
- GitHub Copilot in VSCode was extensively used for assistance.
- Copilot generated a large amount of boilerplate code, but I deliberately left the duplicate code as is without refactoring.

## References
- [Exercises for Programmers](https://www.oreilly.com/library/view/exercises-for-programmers/9781680501513/)
- [Learning Go, 2nd Edition](https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/)
