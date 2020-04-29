package main

import "fmt"
import "encoding/json"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Employee struct {
	Person Person `json:"person"`
	Salary int    `json:"salary"`
}

type Lawyer struct {
	Employee   Employee `json:"employee"`
	IsAttorney bool     `json:"is_attorney"`
}

type Builder struct {
	Employee            Employee `json:"employee"`
	CanWorkInUzbekistan bool     `json:"can_work_in_uzbekistan"`
}

type Coder struct {
	Employee         Employee `json:"employee"`
	FavoriteLanguage string   `json:"favorite_language"`
}

type DrugDealer struct {
	Employee     Employee `json:"employee"`
	PurityOfCoke float64  `json:"purity_of_coke"`
}

type SystemAdmin struct {
	Employee        Employee `json:"employee"`
	HasAwsKnowledge bool     `json:"has_aws_knowledge"`
}

func main() {
	coder := &Coder{
		Employee:         Employee{Salary: 500, Person: Person{Name: "ad", Age: 500}},
		FavoriteLanguage: "js",
	}
	out, err := json.MarshalIndent(coder, "", "  ")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(out))
}
