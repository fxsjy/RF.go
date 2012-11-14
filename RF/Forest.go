package RF

import (
	"math"
	"time"
	"math/rand"
	"fmt"
)
type Forest struct{
	Trees []*Tree
}

func BuildForest(inputs [][]interface{},labels []string, treesAmount, samplesAmount, selectedFeatureAmount int) *Forest{
	rand.Seed(time.Now().UnixNano())
	forest := &Forest{}
	forest.Trees = make([]*Tree,treesAmount)
	done_flag := make(chan bool)
	for i:=0;i<treesAmount;i++{
		go func(x int){
			fmt.Printf("buiding %vth tree...\n", x)
			forest.Trees[x] = BuildTree(inputs,labels,samplesAmount,selectedFeatureAmount)
			fmt.Printf("the %vth tree is done.\n", x)
			done_flag <- true
		}(i)
	}

	for i:=1;i<=treesAmount;i++{
		<-done_flag
		fmt.Printf("tranning progress %v%%\n",float64(i)/float64(treesAmount)*100)
	}

	fmt.Println("all done.")
	return forest
}

func DefaultForest(inputs [][]interface{},labels []string, treesAmount int) *Forest{
	m := int( math.Sqrt( float64( len(inputs[0]) ) ) ) 
	n := int( math.Sqrt( float64( len(inputs) ) )  )
	return BuildForest(inputs,labels, treesAmount,n,m)
}

func (self *Forest) Predicate(input []interface{}) string{
	counter := make(map[string]float64)
	for i:=0;i<len(self.Trees);i++{
		tree_counter := PredicateTree(self.Trees[i],input)
		total := 0.0
		for _,v := range tree_counter{
			total += float64(v)
		}
		for k,v := range tree_counter{
			counter[k] += float64(v) / total
		}
	}

	max_c := 0.0
	max_label := ""
	for k,v := range counter{
		if v>=max_c{
			max_c = v
			max_label = k
		}
	}
	return max_label
}

