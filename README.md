# GO CSV tool

Goals:
- Program can read in a CSV file of quiz
- The tool will ask you the quiz
- Keep track of the questions corrected

Get Started:
- `git clone` this repo
- Run `go run main.go`
- Start answering the quiz
- `go run main.go -help` to look flags

Usage:
- `go run main.go -help`
- ![image](https://github.com/jun-hf/csv-quiz/assets/86782267/cbcfde9d-4cef-4771-aa05-2b49b050426d)


Actions:
- No csv provided will be default to a problems.csv
- Allow users need to provide how long does the quiz should last
- Next questions will be asked if the time is not up yet
- The provided CSV will have short answers

Format of CSV:
```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Edge case:
`"what 2+2, sir?",4`


[^1]:

Exercise from: https://github.com/gophercises/quiz
