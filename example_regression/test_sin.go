package main

import (
	"fmt"
	"../RF/Regression"
	"math"
	"os"
)

func main(){
	out_f,_ := os.OpenFile("sin.out",os.O_CREATE | os.O_RDWR,0777)
	defer out_f.Close()

	train_inputs := make([][]interface{},100)
	train_targets := make([]float64,100)
	
	
	for i:=0;i<len(train_inputs);i++{
		train_inputs[i] = []interface{} {float64(i) / 20.0}
		train_targets[i] = math.Sin(train_inputs[i][0].(float64) )
	}

	forest := Regression.BuildForest(train_inputs,train_targets,100,len(train_inputs),1)

	for i:=0;i<200;i++{
		x := []interface{} {float64(i) / 40.0}
		fmt.Fprintln(out_f,x[0], forest.Predicate(x))
	}

}

