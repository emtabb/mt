package main

import(
	"fmt"
	// "errors"
	// . "./data"
	. "github.com/emtabb/space"
	. "github.com/emtabb/field/statistic"
	"github.com/emtabb/field/calculation/operator"
	//. "./space/element"
	// "reflect"
	// "os"
	// "bufio"
	// "log"
	// "strings"
	// "regexp"
	"strconv"
	"math"

)

// func TestMatrix() {
// 	var m *matrix.Matrix = new(matrix.Matrix)
// 	m.Zeros(10, 10)
// 	var ma *matrix.Matrix = new(matrix.Matrix)
// 	ma.Ones(10, 10)
// 	m.Add(ma)
// 	m.Add(ma)
// 	m.Show()
// }

func estimateCoef(data []float64, expect []float64) (float64, float64) {
	len_data := len(data)
	len_expect := len(expect)
	if len_data != len_expect {
		return 0.0, 0.0
	}
	m_x, m_y := calculate.Mean(data), calculate.Mean(expect)
	SS_xy := calculate.Multi(data, expect) - float64(len_data) * m_y * m_x
	SS_xx := calculate.Multi(data, data) - float64(len_data) * m_x * m_x
	intercept_ := SS_xy / SS_xx
	coef_ := m_y - intercept_ * m_x
	return coef_, intercept_
}

func LinearRegression(predict float64,data []float64, expect []float64) (float64, float64) {
	/*	*NOTE* 
	* - MAE : Mean Absolute Error
	* - MSE : Mean Squared Error
	* - RMSE : Root Mean Squared Error
	* - MAPE : Mean Absolute Percentage Error
	* - MPE : Mean Percentage Error
	*/
	coef_, intercept_ := estimateCoef(data, expect)
	y_predict := coef_ + predict * intercept_ 
	mse := 0.0
	for i := 0; i < len(expect); i++ {
		mse += math.Abs(expect[i] - y_predict)
	}
	mse = mse / float64(len(expect))
	return y_predict, mse
}

// func TestPredict() {
// 	var initCsv *Space = new(Space)
// 	initCsv = initCsv.Init()
// 	initCsv.CsvSpace("./x.csv")
// 	initCsv.View(true, true)
// 	fmt.Println(initCsv.FieldsName())

// 	list_data1 := initCsv.Field("ToaDo").Float()
// 	initCsv.ViewColumn("ToaDo")
// 	list_data2 := initCsv.Field("VanToc").Float()
// 	initCsv.ViewColumn("VanToc")
// 	fmt.Println("", len(list_data1), Max(list_data1))
// 	x, y := estimateCoef(list_data1, list_data2)
// 	fmt.Println("", x, y)
// 	predicted := LinearRegression(4.5, list_data1, list_data2)
// 	fmt.Println("", predicted)
// }

// func TestVector() {
// 	var x *Space = new(Space)
// 	x = x.Init()
// 	x.CsvSpace("./x.csv")
// 	x.View(true, true)
// 	elements := x.Fields([]string{})
// 	for _, ele := range elements {
// 		fmt.Println(ele.Float())
// 	}
// }

// func TestError() (err error) {
// 	err = errors.New("Hello Error")
// 	return err
// }

// func TestElement() {
// 	ele1 := new(DefaultElement).Init()
// 	ele2 := new(NumericalElement).Init()
// 	fmt.Println("Default Element", reflect.TypeOf(reflect.TypeOf(ele1)))
// 	fmt.Println("Numerical Element", reflect.TypeOf(ele2))
// }

// func TestGroupElement() {
// 	var initCsv *Space = new(Space)
// 	initCsv = initCsv.Init()
// 	initCsv.CsvSpace("./DS_TT.xlsx")
// 	initCsv.View(true, true)

// }

// func ReadDataPhysic(path string) ([][]float64, [][]float64) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	string_element := make([][]string, 0)
// 	start_table := false
// 	number_table := 0
// 	re := regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+")

// 	for scanner.Scan() {
// 		if start_table == false {
// 			if strings.Contains(scanner.Text(), "+-------+------------------+") {
// 				start_table = true
// 				temp := make([]string, 0)
// 				string_element = append(string_element, temp)
// 				string_element[number_table] = append(string_element[number_table], scanner.Text())
// 				continue
// 			}
// 		}
// 		if start_table == true {
// 			if strings.Contains(scanner.Text(), "+-------+------------------+") {
// 				start_table = false
// 				string_element[number_table] = append(string_element[number_table], scanner.Text())
// 				number_table = number_table + 1
// 				continue
// 			} 
// 			string_element[number_table] = append(string_element[number_table], scanner.Text())
// 		}
// 	}
	
// 	ListElement := make([][]Element, len(string_element))
// 	for i := 0; i < len(string_element); i++ {
// 		ListElement[i] = make([]Element, len(string_element[0]))
// 		for j, element := range string_element[i] {
// 			ListElement[i][j] = new(DefaultElement).Init()
// 			data_temp := re.FindAllString(element, -1)
// 			if len(data_temp) != 0 {
// 				ListElement[i][j].Set(data_temp)
// 			} 
// 		}
// 	}

// 	len_listelement := len(ListElement)

// 	label := make([][]float64, len_listelement)
// 	result := make([][]float64, len_listelement)

// 	for i := 0; i < len_listelement; i++ {
// 		label[i] = make([]float64, 0)
// 		result[i] = make([]float64, 0)
// 		for j := 0; j < len(ListElement[0]); j++ {
// 			list_y := ListElement[i][j].Float()
// 			y := 0.0
// 			if len(list_y) == 3 {
// 				y = list_y[0] * math.Pow(10, list_y[2])
// 			} else {
// 				continue
// 			}
// 			x, _ := strconv.ParseFloat(ListElement[i][j].Label(),64)
// 			label[i] = append(label[i], x)
// 			result[i] = append(result[i], y)
// 		}	
// 	}
// 	return label, result
// }

// func TestRead(path string) {
// 	var x *Space = new(Space)
// 	x = x.Init()
// 	x.HighEnergySpace(path)
// 	x.View(true, true)
// }

//func main() {
	// label, hard := ReadDataPhysic("hardEEZZ.result")
	// _, loop := ReadDataPhysic("loopEEZZ.result")
	// var tools ComposeTools = VectorTools{}
	// result := tools.Add(hard[1], loop[1])
	// str_save := "x,dsigma\n"
	// for i := 0; i < len(label[1]); i++ {
	// 	str_save += fmt.Sprintf("%f", label[1][i]) + "," + fmt.Sprintf("%f", result[i]) + "\n"
	// }
	// file, err := os.Create("test1.csv")
	// if err != nil {
	// 	log.Fatal("failed creating file:", err)
	// }
	// defer file.Close()

	// file.WriteString(str_save)
	//TestRead("loopEEZZ.result")
	//TestVector()
//}

// type IStruct struct {
//     Data map[string]interface{ }
// }

func CleanPhysicData(path string) {
	var space *Space = new(Space).Init()
	space.CsvSpace("x.csv")
	space.HighEnergySpace(path)
	for _, element := range space.Elements {
		value, _ := strconv.ParseFloat(element.Field("FIELD2").(string), 64)
		pow, _ := strconv.ParseFloat(element.Field("FIELD4").(string), 64)
		temp_result := fmt.Sprintf("%f", value * math.Pow(10, pow))
		column := element.Field("FIELD1").(string)
		fmt.Println(column + "," + temp_result)
	}
}

func TestGroup() {
	var space *Space = new(Space).Init()
	space.CsvSpace("x.csv")
	fg := []string{"ToaDo","VanToc"}
	xspace, _ := space.Group(fg)
	fmt.Println(space)
	fmt.Println(xspace)
	fmt.Println(len(xspace.Elements))
	for i := 0; i < len(xspace.Elements); i++ {
		fmt.Println(i, xspace.Elements[i])
	}
	xreshape, _ := xspace.Reshape()
	for i := 0; i < len(xreshape.Elements); i++ {
		fmt.Println(i, xreshape.Elements[i])
	}

	predict, error := LinearRegression(0.4, xreshape.Float("ToaDo"), xreshape.Float("VanToc"))
	fmt.Println(predict, error)
}

func TestBigdata() {
	var space *Space = new(Space).Init()
	space.CsvSpace("train.csv")
	for _, label := range space.FieldsName() {
		fmt.Print(label + " | ")
	}
	fmt.Println()
	for _, data := range space.Head() {
		for _, label := range data.Label() {
			fmt.Print(data.Field(label).(string) + " | ")
		}
		fmt.Println()
	}

	group := []string {"LotArea", "SalePrice"}
	LotFrontageString := []string {"LotArea"}

	LotFrontage := space.Field(LotFrontageString)[0].Double()
	totalLotFrontage := 0.0
	countLostData := 0
	for _, lotFrontage := range LotFrontage {
		if lotFrontage != 0.0 {
			data := lotFrontage
			totalLotFrontage += data
		} else {
			countLostData++
		}
	}

	meanLotFrontage := totalLotFrontage / float64(len(LotFrontage)- countLostData)

	for i := 0; i < len(LotFrontage); i++ {
		if space.Elements[i].Field("LotArea") == "NA" {
			space.Elements[i].SetField("LotArea", meanLotFrontage)
		}
		
	}
	fmt.Println()
	
	LotFrontage = space.Field(LotFrontageString)[0].Double()
	fmt.Println("LotArea")

	subSpace, _ := space.Group(group)
	expect := 9600.0
	predict, error := LinearRegression(expect, subSpace.Float("LotArea"), subSpace.Float("SalePrice"))
	fmt.Println("\nPredict %f LotFrontage follow SalePrice is: %f error %f", expect, predict, error)
}

func testDistinct() {
	var space *Space = new(Space).Init()
	space.CsvSpace("train.csv")
	fname := []string {"Neighborhood"}
	f := space.Field(fname)[0]
	c := Distinct(f.GetData())
	for _, a := range c {
		fmt.Println(a)
	}
}

func main () {
	//TestGroup()
	TestBigdata()
	//testDistinct()
	//"MSSubClass", "LotFrontage"
}