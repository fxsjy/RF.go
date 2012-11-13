package RF

import (
	"log"
	"math"
	//"fmt"
)
type Forest struct{
	Trees []*Tree
}

func BuildForest(inputs [][]interface{},labels []string, treesAmount, samplesAmount, selectedFeatureAmount int) *Forest{
	forest := &Forest{}
	forest.Trees = make([]*Tree,treesAmount)
	for i:=0;i<treesAmount;i++{
		log.Printf("building the %vth tree\n", i)
		forest.Trees[i] = BuildTree(inputs,labels,samplesAmount,selectedFeatureAmount)
	}
	log.Println("done.")
	return forest
}

func DefaultForest(inputs [][]interface{},labels []string, treesAmount int) *Forest{
	m := int( math.Sqrt( float64( len(inputs[0]) ) ) ) 
	n := int( math.Log( float64( len(inputs) ) ) / math.Log(1.3) )
	return BuildForest(inputs,labels, treesAmount,n,m)
}

func (self *Forest) Predicate(input []interface{}) string{
	counter := make(map[string]int)
	for i:=0;i<len(self.Trees);i++{
		label := PredicateTree(self.Trees[i],input)
		counter[label] += 1
	}
	max_c := 0
	max_label := ""
	for k,v := range counter{
		if v>=max_c{
			max_c = v
			max_label = k
		}
	}
	return max_label
}

