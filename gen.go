package main

import (
	"fmt"
	"math/rand"
	"strconv"
 	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strings"
	"math"
)

var db, _ = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/avadata")

type Person struct {
	fullName string
	firstName string
	lastName string
	year int
	month int
	day int
	place string
}

type GeoData struct {
	name string
	x float64
	y float64
}

type Value struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (v *Value) toString() string {
	return "(" + v.Name + "[" + v.Type + "]" + ")"
}

type Module struct {
	Name   string `json:"name"`
	Input  []Value `json:"input"`
	Output []Value `json:"output"`
}

func (m *Module) toString() string {
	str := "Module: " + m.Name + "\n"
	str += "inputs: "
	for _, input := range m.Input {
		str += input.toString()
		str += ", "
	}
	str += "\n"
	str += "outputs: "
	for _, output := range m.Output {
		str += output.toString()
		str += ", "
	}
	str += "\n"
	return str
}

func concatModule(a string, b string) Module {
	return Module{
		"concat",
		[]Value{Value{a, "string"}, Value{b, "string"}},
		[]Value{Value{a + " " + b, "string"}}}
}

func isEvenModule(a int) Module {
	result := 0
	if a % 2 == 0 {
		result = 1
	}
	return Module{
		"even",
		[]Value{Value{strconv.Itoa(a), "int"}},
		[]Value{Value{strconv.Itoa(result), "int"}}}
}

func i2sModule(a int) Module {
	return Module{
		"i2s",
		[]Value{Value{strconv.Itoa(a), "int"}},
		[]Value{Value{strconv.Itoa(a), "string"}}}
}

func s2iModule(a int) Module {
	return Module{
		"s2i",
		[]Value{Value{strconv.Itoa(a), "string"}},
		[]Value{Value{strconv.Itoa(a), "int"}}}
}

func i2yModule(a int) Module {
	result := "yes"
	if a == 0 {
		result = "no"
	}
	return Module{
		"i2y",
		[]Value{Value{strconv.Itoa(a), "int"}},
		[]Value{Value{result, "string"}}}
}

func isPositiveModule(a int) Module {
	result := 1
	if a <= 0 {
		result = 0
	}
	return Module{
		"positive",
		[]Value{Value{strconv.Itoa(a), "int"}},
		[]Value{Value{strconv.Itoa(result), "int"}}}
}

func isSquareModule(a int) Module {
	result := 0
	for i := 1; i * i <= a; i++ {
		if i * i == a {
			result = 1
			break
		}
	}
	return Module{
		"square",
		[]Value{Value{strconv.Itoa(a), "int"}},
		[]Value{Value{strconv.Itoa(result), "int"}}}
}

func isPrimeModule(a int) Module {
	result := 1
	for i := 2; i * i <= a; i++ {
		if a % i == 0 {
			result = 0
			break
		}
	}
	return Module{
		"prime",
		[]Value{Value{strconv.Itoa(a), "int"}},
		[]Value{Value{strconv.Itoa(result), "int"}}}
}

func compareModule(a int, b int) Module {
	result := -1
	if a == b {
		result = 0;
	} else if a > b {
		result = 1;
	}
	return Module{
		"compare",
		[]Value{Value{strconv.Itoa(a), "int"}, Value{strconv.Itoa(b), "int"}},
		[]Value{Value{strconv.Itoa(result), "int"}}}
}

func sumModule(a int, b int) Module {
	return Module{
		"sum",
		[]Value{Value{strconv.Itoa(a), "int"}, Value{strconv.Itoa(b), "int"}},
		[]Value{Value{strconv.Itoa(a + b), "int"}}}
}

func diffModule(a int, b int) Module {
	return Module{
		"diff",
		[]Value{Value{strconv.Itoa(a), "int"}, Value{strconv.Itoa(b), "int"}},
		[]Value{Value{strconv.Itoa(a - b), "int"}}}
}

func getFirstAndLastNameModule(fullName string) Module {
	tokens := strings.Split(fullName, " ")
	return Module{
		"get_first_last_name",
		[]Value{Value{fullName, "string"}},
		[]Value{Value{tokens[0], "string"}, Value{tokens[1], "string"}}}
}

func getDateOfBirthModule(fullName string, year int, month int, day int) Module {
	return Module{
		"get_DOB",
		[]Value{Value{fullName, "string"}},
		[]Value{Value{strconv.Itoa(year), "int"},
			    Value{strconv.Itoa(month), "int"},
				Value{strconv.Itoa(day), "int"}}}
}

func getDateOfBirthToStringModule(year int, month int, day int) Module {
	return Module{
		"get_DOB_to_string",
		[]Value{Value{strconv.Itoa(year), "int"},
				Value{strconv.Itoa(month), "int"},
				Value{strconv.Itoa(day), "int"}},
		[]Value{Value{strconv.Itoa(year) + "-" +
		              strconv.Itoa(month) + "-" +
					  strconv.Itoa(day), "string"}}}
}

func getCoordinatesModule(place string, x int, y int) Module {
	return Module{
		"get_coords",
		[]Value{Value{place, "string"}},
		[]Value{Value{strconv.Itoa(x), "int"}, Value{strconv.Itoa(y), "int"}}}
}

func getDistModule(x int, y int, x2 int, y2 int, dst int) Module {
	return Module{
		"get_dst",
		[]Value{Value{strconv.Itoa(x), "int"},
			    Value{strconv.Itoa(y), "int"},
				Value{strconv.Itoa(x2), "int"},
				Value{strconv.Itoa(y2), "int"}},
		[]Value{Value{strconv.Itoa(dst), "int"}}}
}

func getYearFromDate(year int, month int, day int) Module {
	return Module{
		"get_year_dob",
		[]Value{Value{strconv.Itoa(year), "int"},
				Value{strconv.Itoa(month), "int"},
				Value{strconv.Itoa(day), "int"}},
		[]Value{Value{strconv.Itoa(year), "int"}}}
}

func getMonthFromDate(year int, month int, day int) Module {
	return Module{
		"get_month_dob",
		[]Value{Value{strconv.Itoa(year), "int"},
			Value{strconv.Itoa(month), "int"},
			Value{strconv.Itoa(day), "int"}},
		[]Value{Value{strconv.Itoa(month), "int"}}}
}

func getDayFromDate(year int, month int, day int) Module {
	return Module{
		"get_day_dob",
		[]Value{Value{strconv.Itoa(year), "int"},
			Value{strconv.Itoa(month), "int"},
			Value{strconv.Itoa(day), "int"}},
		[]Value{Value{strconv.Itoa(day), "int"}}}
}

func getPlaceOfBirthModule(fullName string, place string) Module {
	return Module{
		"get_place",
		[]Value{Value{fullName, "string"}},
		[]Value{Value{place, "string"}}}
}

type SimpleStatement struct {
	result string;
	resultType string; // enum{int, string}
	generalType string; // enum{math, people, geo}
	negation bool; // *is* or *is not*
	what string;
	of string;
	representation string;
	simpleQuestion string;
	sequenceOfModules []Module;
}

func (a *SimpleStatement) addModule(m Module) {
	a.sequenceOfModules = append(a.sequenceOfModules, m);
}

func (a *SimpleStatement) getQuestion() string {
	if (a.simpleQuestion == "what") {
		return "What is " + a.what + " of " + a.of;
	} else {
		return "Is " + a.result + " " + a.what + " " + a.of;
	}
}

func (a *SimpleStatement) getResult() string {
	if a.simpleQuestion == "what" {
		return a.result
	} else {
		return a.result
	}

}

func (a *SimpleStatement) getModules() []string {
	modules := make([]string, 0)
	for _, module := range a.sequenceOfModules {
		modules = append(modules, module.toString())
	}
	return modules
}

func getNameDetails(name string) (string, string) {
	if strings.Index(name, " ") == -1 {
		return "", ""
	}
	tokens := strings.Split(name, " ")
	if len(tokens) != 2 {
		return "", ""
	}
	return tokens[0], tokens[1]
}

func getDateDetails(date string) (int, int, int) {
	if date == "NULL" {
		return -1, -1, -1
	}
	if len(date) == 0 {
		return -1, -1, -1
	}
	if strings.Index(date, " ") != -1 {
		return -1, -1, -1
	}
	if date[0] == '-' {
		y, _ := strconv.Atoi(date)
		return y, -1, -1
	}
	if strings.Index(date, "-") == -1 {
		y, _ := strconv.Atoi(date)
		return y, -1, -1
	} else {
		tokens := strings.Split(date, "-")
		y, _ := strconv.Atoi(tokens[0])
		m, _ := strconv.Atoi(tokens[1])
		d, _ := strconv.Atoi(tokens[2])
		return y, m, d
	}
}

func getCoords(coords string) (float64, float64) {
	tokens := strings.Split(coords, " ")
	x, _ := strconv.ParseFloat(tokens[0], 64)
	y, _ := strconv.ParseFloat(tokens[1], 64)
	return x, y
}

func getDataFromDBPeople() ([]Person, []GeoData) {
	fmt.Println("Loading databases...")
	rows, _ := db.Query("select name, date, place from birthdays")
	defer rows.Close()
	persons := make([]Person, 0)
	for rows.Next() {
		var name, place, date string;
		rows.Scan(&name, &date, &place)
		firstName, lastName := getNameDetails(name)
		year, month, day := getDateDetails(date)
		person := Person{name, firstName, lastName, year, month, day, place}
		persons = append(persons, person)
	}
	rows2, _ := db.Query("select name, coords from geomapping")
	defer rows2.Close()
	places := make([]GeoData, 0)
	for rows2.Next() {
		var name, coords string;
		rows2.Scan(&name, &coords)
		x, y := getCoords(coords)
		place := GeoData{name, x, y}
		places = append(places, place)
	}
	fmt.Println("Database was loaded!")
	return persons, places
}

var persons []Person;
var places []GeoData;

func GeneratePeopleStatements(r *rand.Rand) SimpleStatement {
	/**
		Types:
			1) First/last name of a person
			2) DOB/year/day/month of a person
			3) Place of birth of a person
	 */
	functions := make([]func(r *rand.Rand) SimpleStatement, 3)
	generalType := "people"

	functions[0] = func(r *rand.Rand) SimpleStatement {
		// First/last name of a person
		type_of_name := r.Intn(2)
		what := "first name"
		firstName, lastName, personName := "", "", ""
		name := ""
		// We need to find a person with first and last name
		for ;; {
			idx := r.Intn(len(persons))
			if persons[idx].firstName == "" || persons[idx].lastName == "" {
				continue
			}
			personName, firstName, lastName = persons[idx].fullName, persons[idx].firstName, persons[idx].lastName
			break
		}
		if (type_of_name == 0) {
			// First name
			name = firstName
		} else {
			// Last name
			what = "last name"
			name = lastName
		}
		result := SimpleStatement{
			name,
			"string",
			generalType,
			false,
			what,
			personName,
			name + " is " + what + " of " + personName,
			"what",
			make([]Module, 0)}
		result.addModule(concatModule(firstName, lastName))
		result.addModule(getFirstAndLastNameModule(personName))
		return result;
	}

	functions[1] = func(r *rand.Rand) SimpleStatement {
		// DOB
		type_of_dob := r.Intn(4)
		var year, month, day int;
		var personName string;
		// We need to find a person with DOB
		for ;; {
			idx := r.Intn(len(persons))
			if persons[idx].year == -1 || persons[idx].month == -1 || persons[idx].day == -1 {
				continue
			}
			year, month, day = persons[idx].year, persons[idx].month, persons[idx].day
			personName = persons[idx].fullName
			break
		}
		if (type_of_dob == 0) {
			// full DOB
			dob := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day)
			result := SimpleStatement{
				dob,
				"string",
				generalType,
				false,
				"DOB",
				personName,
				dob + " is DOB of " + personName,
				"what",
				make([]Module, 0)}
			tokens := strings.Split(personName, " ")
			str := tokens[0]
			for i := 1; i < len(tokens); i++ {
				result.addModule(concatModule(str, tokens[i]))
				str += " "
				str += tokens[i]
			}
			result.addModule(getDateOfBirthModule(personName, year, month, day))
			result.addModule(getDateOfBirthToStringModule(year, month, day))
			return result
		} else if (type_of_dob == 1) {
			// year
			result := SimpleStatement{
				strconv.Itoa(year),
				"int",
				generalType,
				false,
				"year of birth",
				personName,
				strconv.Itoa(year) + " is year of birth of " + personName,
				"what",
				make([]Module, 0)}
			tokens := strings.Split(personName, " ")
			str := tokens[0]
			for i := 1; i < len(tokens); i++ {
				result.addModule(concatModule(str, tokens[i]))
				str += " "
				str += tokens[i]
			}
			result.addModule(getDateOfBirthModule(personName, year, month, day))
			result.addModule(getYearFromDate(year, month, day))
			result.addModule(i2sModule(year))
			return result
		} else if (type_of_dob == 2) {
			// month
			result := SimpleStatement{
				strconv.Itoa(month),
				"int",
				generalType,
				false,
				"month of birth",
				personName,
				strconv.Itoa(month) + " is month of birth of " + personName,
				"what",
				make([]Module, 0)}
			tokens := strings.Split(personName, " ")
			str := tokens[0]
			for i := 1; i < len(tokens); i++ {
				result.addModule(concatModule(str, tokens[i]))
				str += " "
				str += tokens[i]
			}
			result.addModule(getDateOfBirthModule(personName, year, month, day))
			result.addModule(getMonthFromDate(year, month, day))
			result.addModule(i2sModule(month))
			return result
		} else {
			// day
			result := SimpleStatement{
				strconv.Itoa(day),
				"int",
				generalType,
				false,
				"day of birth",
				personName,
				strconv.Itoa(day) + " is day of birth of " + personName,
				"what",
				make([]Module, 0)}
			tokens := strings.Split(personName, " ")
			str := tokens[0]
			for i := 1; i < len(tokens); i++ {
				result.addModule(concatModule(str, tokens[i]))
				str += " "
				str += tokens[i]
			}
			result.addModule(getDateOfBirthModule(personName, year, month, day))
			result.addModule(getDayFromDate(year, month, day))
			result.addModule(i2sModule(day))
			return result
		}
	}

	functions[2] = func(r *rand.Rand) SimpleStatement {
		// Place of birth
		var personName, place string
		// We need to find a person with place of birth not null
		for ;; {
			idx := r.Intn(len(persons))
			if persons[idx].place == "" {
				continue
			}
			personName, place = persons[idx].fullName, persons[idx].place
			break
		}
		result := SimpleStatement{
			place,
			"string",
			generalType,
			false,
			"place of birth",
			personName,
			place + " is place of birth of " + personName,
			"what",
			make([]Module, 0)}
		tokens := strings.Split(personName, " ")
		str := tokens[0]
		for i := 1; i < len(tokens); i++ {
			result.addModule(concatModule(str, tokens[i]))
			str += " "
			str += tokens[i]
		}
		result.addModule(getPlaceOfBirthModule(personName, place))
		return result
	}
	id := r.Intn(3)
	return functions[id](r)
}

func GenerateGeoStatements(r *rand.Rand) SimpleStatement {
	/**
		Types:
			1) D is distance between A and B
	 */
	functions := make([]func(r *rand.Rand) SimpleStatement, 1)
	generalType := "places"

	functions[0] = func(r *rand.Rand) SimpleStatement {
		for {
			idx := r.Intn(len(places))
			idx2 := r.Intn(len(places))
			if idx == idx2 {
				continue
			}
			name1, name2 := places[idx].name, places[idx2].name
			dst := int(math.Hypot(places[idx].x - places[idx2].x, places[idx].y - places[idx2].y) + 0.5)
			result := SimpleStatement{
				strconv.Itoa(dst),
				"int",
				generalType,
				false,
				"distance between",
				name1 + " and " + name2,
				strconv.Itoa(dst) + " is distance between " + name1 + " and " + name2,
				"what",
				make([]Module, 0)}
			tokens := strings.Split(places[idx].name, " ")
			str := tokens[0]
			for i := 1; i < len(tokens); i++ {
				result.addModule(concatModule(str, tokens[i]))
				str += " "
				str += tokens[i]
			}

			tokens = strings.Split(places[idx2].name, " ")
			str = tokens[0]
			for i := 1; i < len(tokens); i++ {
				result.addModule(concatModule(str, tokens[i]))
				str += " "
				str += tokens[i]
			}

			result.addModule(getCoordinatesModule(places[idx].name, int(places[idx].x), int(places[idx].y)))
			result.addModule(getCoordinatesModule(places[idx2].name, int(places[idx2].x), int(places[idx2].y)))
			result.addModule(getDistModule(int(places[idx].x),
				int(places[idx].y), int(places[idx2].x), int(places[idx2].y), dst))
			result.addModule(i2sModule(dst))

			return result
		}
	}
	return functions[0](r)
}

func GenerateMathStatements(r *rand.Rand) SimpleStatement {
	/**
	    Types:
	    	1) Prime
	    	2) Even/odd
	    	3) Perfect square
	    	4) Positive/negative
	    	5) Sum/Subtraction
	    	6) Greater/Equal/Smaller
	 */
	generalType := "math"
	maxIntLimit := 100
	functions := make([]func(r *rand.Rand) SimpleStatement, 6)

	// Function is prime
	functions[0] = func(r *rand.Rand) SimpleStatement {
		toCheck := r.Intn(maxIntLimit) + 1;
		isPrime := true;
		isPrimeInt := 1
		for i := 2; i * i <= toCheck; i++ {
			if toCheck % i == 0 {
				isPrime = false;
				isPrimeInt = 0
				break;
			}
		}

		result := SimpleStatement{
			strconv.Itoa(isPrimeInt),
			"int",
			generalType,
			!isPrime,
			"prime",
			"number",
			"",
			"is",
			make([]Module, 0)}
		result.addModule(s2iModule(toCheck))
		result.addModule(isPrimeModule(toCheck))
		result.addModule(i2yModule(isPrimeInt))
		if isPrime {
			result.representation = strconv.Itoa(toCheck) + " is prime number"
		} else {
			result.representation = strconv.Itoa(toCheck) + " is not prime number"
		}
		return result;
	}

	// Function even/odd
	functions[1] = func(r *rand.Rand) SimpleStatement {
		toCheck := r.Intn(maxIntLimit) + 1;
		even := toCheck % 2 == 0
		result := SimpleStatement{
			strconv.Itoa(int(1 - toCheck % 2)),
			"int",
			generalType,
			false,
			"",
			"number",
			"",
			"is",
			make([]Module, 0)}
		if even {
			result.what = "even"
			result.representation = strconv.Itoa(toCheck) + " is even number"
		} else {
			result.what = "odd"
			result.representation = strconv.Itoa(toCheck) + " is odd number"
		}
		result.addModule(s2iModule(toCheck))
		result.addModule(isEvenModule(toCheck))
		result.addModule(i2yModule(1 - toCheck % 2))
		return result;
	}

	// Function perfect square
	functions[2] = func(r *rand.Rand) SimpleStatement {
		toCheck := r.Intn(maxIntLimit) + 1;
		perfectSquare := false
		perfectSquareInt := 0
		for i := 1; i * i <= toCheck; i++ {
			if i * i == toCheck {
				perfectSquare = true
				perfectSquareInt = 1
				break
			}
		}
		result := SimpleStatement{
			strconv.Itoa(perfectSquareInt),
			"int",
			generalType,
			!perfectSquare,
			"perfect square",
			"number",
			"",
			"is",
			make([]Module, 0)}
		if perfectSquare {
			result.representation = strconv.Itoa(toCheck) + " is perfect square number"
		} else {
			result.representation = strconv.Itoa(toCheck) + " is not perfect square number"
		}
		result.addModule(s2iModule(toCheck))
		result.addModule(isSquareModule(toCheck))
		result.addModule(i2yModule(perfectSquareInt))
		return result;
	}

	// Function positive/negative
	functions[3] = func(r *rand.Rand) SimpleStatement {
		toCheck := r.Intn(maxIntLimit) + 1
		mul := r.Intn(2)
		isPositive := 0
		if toCheck > 0 {
			isPositive = 1
		}
		result := SimpleStatement{
			strconv.Itoa(isPositive),
			"int",
			generalType,
			false,
			"",
			"number",
			"",
			"is",
			make([]Module, 0)}
		if mul != 0 {
			result.what = "negative"
			result.representation = strconv.Itoa(toCheck) + " is negative number"
		} else {
			result.what = "positive"
			result.representation = strconv.Itoa(toCheck) + " is positive number"
		}
		result.addModule(s2iModule(toCheck))
		result.addModule(isPositiveModule(toCheck))
		result.addModule(i2yModule(isPositive))
		return result;
	}

	// Function sum and subtraction
	functions[4] = func(r *rand.Rand) SimpleStatement {
		a := r.Intn(maxIntLimit) + 1
		b := r.Intn(maxIntLimit) + 1
		var result SimpleStatement
		if r.Intn(2) != 0 {
			sub := a - b
			result = SimpleStatement{
				strconv.Itoa(sub),
				"int",
				generalType,
				false,
				"difference",
				strconv.Itoa(a) + " and " + strconv.Itoa(b),
				strconv.Itoa(sub) + " is difference of " + strconv.Itoa(a) + " and " + strconv.Itoa(b),
				"what",
				make([]Module, 0)}
			result.addModule(s2iModule(a))
			result.addModule(s2iModule(b))
			result.addModule(diffModule(a, b))
			result.addModule(i2sModule(a - b))
		} else {
			sum := a + b
			result = SimpleStatement{
				strconv.Itoa(sum),
				"int",
				generalType,
				false,
				"sum",
				strconv.Itoa(a) + " and " + strconv.Itoa(b),
				strconv.Itoa(sum) + " is sum of " + strconv.Itoa(a) + " and " + strconv.Itoa(b),
				"what",
				make([]Module, 0)}
			result.addModule(s2iModule(a))
			result.addModule(s2iModule(b))
			result.addModule(sumModule(a, b))
			result.addModule(i2sModule(a + b))
		}
		return result
	}

	// Function compare
	functions[5] = func(r *rand.Rand) SimpleStatement {
		a := r.Intn(maxIntLimit) + 1
		b := r.Intn(maxIntLimit) + 1
		var result SimpleStatement
		var compareResult int;
		if a < b {
			compareResult = -1
			result = SimpleStatement{
				"-1",
				"int",
				generalType,
				true,
				"greater",
				strconv.Itoa(b),
				strconv.Itoa(a) + " is not greater than " + strconv.Itoa(b),
				"is",
				make([]Module, 0)}
		} else if a == b {
			compareResult = 0
			result = SimpleStatement{
				"0",
				"int",
				generalType,
				false,
				"equal",
				strconv.Itoa(b),
				strconv.Itoa(a) + " is equal to " + strconv.Itoa(b),
				"is",
				make([]Module, 0)}
		} else {
			compareResult = 1
			result = SimpleStatement{
				"1",
				"int",
				generalType,
				false,
				"greater",
				strconv.Itoa(b),
				strconv.Itoa(a) + " is greater than " + strconv.Itoa(b),
				"is",
				make([]Module, 0)}
		}

		result.addModule(s2iModule(a))
		result.addModule(s2iModule(b))
		result.addModule(compareModule(a, b))
		result.addModule(i2sModule(compareResult))
		return result
	}
	id := r.Intn(6)
	return functions[id](r)
}

var statements []SimpleStatement

func GenerateStatements(r *rand.Rand) {
	fmt.Println("Generating statements...")
	for iter := 0; iter <= 10000; iter++ {
		statements = append(statements, GenerateMathStatements(r))
		statements = append(statements, GeneratePeopleStatements(r))
		statements = append(statements, GenerateGeoStatements(r))
	}
	fmt.Println("Simple statements were generated!")
}

type Result struct {
	Question string `json:"question"`
	Answer string `json:"answer"`
	Modules []Module `json:"modules"`
}

func convertToJson(question string, answer string, modules []Module) string {
	result := Result{question, answer, modules}
	res, _ := json.Marshal(result)
	return string(res)
}

func main() {
	r := rand.New(rand.NewSource(3731))
	persons, places = getDataFromDBPeople()
	GenerateStatements(r)
	for _, statement := range statements {
		fmt.Println(convertToJson(statement.getQuestion() + "?",
								  statement.getResult(),
								  statement.sequenceOfModules))
	}
}
