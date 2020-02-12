package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

//"os"
//"github.com/irisqaz/practice-go/test"

func main() {
	// employee_list = read_employees('<file_location>')
	employeeList := readEmployees("data/employees.csv")
	fmt.Printf("%#v\n", employeeList)
	// print(employee_list)
}

func readEmployees(filename string) []map[string]string {
	//csv.register_dialect('empDialect', skipinitialspace=True, strict=True)
	//employee_file = csv.DictReader(open(csv_file_location), dialect = 'empDialect')
	f, _ := os.Open(filename)
	defer f.Close()
	reader := csv.NewReader(f)
	employeeList, _ := reader.ReadAll()
	// employee_list = []
	// for data in employee_file:
	//   employee_list.append(data)
	// return employee_list
	list := []map[string]string{}
	for _, l := range employeeList {
		record := map[string]string{}
		//clefmt.Println(l)
		record["Department"] = l[0]
		record["Username"] = l[1]
		record["Fullname"] = l[2]
		list = append(list, record)
	}
	return list

}
