package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	//"strconv"
	"os"

	"time"

	//"math"
	"../RF"
)


func main(){

	start := time.Now()
	f,_ := os.Open("car.data")
	defer f.Close()
	content,_ := ioutil.ReadAll(f)
	s_content := string(content)
	lines := strings.Split(s_content,"\n")
	
	inputs := make([][]interface{},0)
	targets := make([]string,0)
	for _,line := range lines{

		line = strings.TrimRight(line,"\r\n")

		if len(line)==0{
			continue
		}
		tup := strings.Split(line,",")
		pattern := tup[:len(tup)-1]
		target := tup[len(tup)-1]
		X := make([]interface{},0)
		for _,x := range pattern{
			X = append(X,x)	
		}
		inputs = append(inputs,X)
	
		targets = append(targets,target)
	}
	train_inputs := make([][]interface{},0)

	train_targets := make([]string,0)

	test_inputs := make([][]interface{},0)
	test_targets := make([]string,0)

	for i,x := range inputs{
		if i%2==1{
			test_inputs = append(test_inputs, x)
		}else{
			train_inputs = append(train_inputs, x)
		}
	}

	for i,y := range targets{
		if i%2==1{
			test_targets = append(test_targets,y)
		}else{
			train_targets = append(train_targets,y)
		}
	}

	forest := RF.BuildForest(inputs,targets,10,500,len(train_inputs[0]))//100 trees

	test_inputs = train_inputs
	test_targets = train_targets
	err_count := 0.0
	for i:=0;i<len(test_inputs);i++{
		output := forest.Predicate(test_inputs[i])
		expect := test_targets[i]
		//fmt.Println(output,expect)
		if output!=expect{
			err_count += 1
		}
	}
	fmt.Println("success rate:",1.0 - err_count/float64(len(test_inputs)))

	fmt.Println(time.Since(start))

}
