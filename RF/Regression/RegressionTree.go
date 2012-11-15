//a random forest implemtation in GoLang
package Regression

import (
	"math/rand"
	//"fmt"
)

const CAT = "cat"
const NUMERIC = "numeric"

type TreeNode struct{
	ColumnNo int //column number
	Value interface{}
	Left *TreeNode
	Right *TreeNode
	Label float64
}

type Tree struct{
	Root *TreeNode
}

func getRandomRange(N int, M int) []int{
	tmp := make([]int,N)
	for i:=0;i<N;i++{
		tmp[i]=i
	}
	for i:=0;i<M;i++{
		j := i + int(rand.Float64()*float64(N-i))
		tmp[i],tmp[j] = tmp[j],tmp[i]
	}

	return tmp[:M]
}

func getSamples(ary [][]interface{}, index []int)  [][]interface{} {
	result := make([][]interface{}, len(index))
	for i:=0;i<len(index);i++{
		result[i] = ary[index[i]]
	}
	return result
}


func getLabels(ary []float64, index []int ) []float64{
	result := make([]float64,len(index))
	for i:=0;i<len(index);i++{
		result[i] = ary[index[i]]
	}
	return result
}

func getMSE(labels []float64) float64 {
	total := 0.0
	for _,x := range labels{
		total += x
	}
	avg := total/float64(len(labels))
	mse := 0.0
	for _,x := range labels{
		delta := x - avg
		mse += delta*delta
	}
	mse = mse/float64(len(labels))
	return mse
}


func getBestGain(samples [][]interface{}, c int, samples_labels []float64, column_type string) (float64,interface{},[]int,[]int){
	var best_part_l []int
	var best_part_r []int
	var best_value interface{}
	best_gain := 0.0

	current_mse := getMSE(samples_labels)

	uniq_values := make(map[interface{}]int)
	for i:=0;i<len(samples);i++{
		uniq_values[samples[i][c]] = 1
	}

	for value,_ := range uniq_values{
		labels_l := make([]float64,0)
		labels_r := make([]float64,0)
		part_l := make([]int,0)
		part_r := make([]int,0)
		if column_type==CAT{
			for j:=0;j<len(samples);j++{
				if samples[j][c]==value{
					part_l = append(part_l,j)
					labels_l = append(labels_l,samples_labels[j])
				}else{
					part_r = append(part_r,j)
					labels_r = append(labels_r,samples_labels[j])
				}
			}
		}
		if column_type==NUMERIC{
			for j:=0;j<len(samples);j++{
				if samples[j][c].(float64)<=value.(float64){
					part_l = append(part_l,j)
					labels_l = append(labels_l,samples_labels[j])
				}else{
					part_r = append(part_r,j)
					labels_r = append(labels_r,samples_labels[j])
				}
			}
		}

		p1 := float64(len(part_r)) / float64(len(samples))
		p2 := float64(len(part_l)) / float64(len(samples))

		new_mse := p1*getMSE(labels_r) + p2*getMSE(labels_l)

		mse_gain := current_mse - new_mse
		
		if mse_gain>=best_gain{
			best_gain = mse_gain
			best_value = value
			best_part_l = part_l
			best_part_r = part_r
		}
	}

	return best_gain, best_value, best_part_l,best_part_r
}


func buildTree(samples [][]interface{}, samples_labels []float64, selected_feature_count int) *TreeNode{
	//fmt.Println(len(samples))
	//find a best splitter
	column_count := len(samples[0])
	//split_count := int(math.Log(float64(column_count)))
	split_count := selected_feature_count
	columns_choosen := getRandomRange(column_count,split_count)
	
	best_gain := 0.0
	var best_part_l []int
	var best_part_r []int
	var best_value interface{}
	var best_column int

	for _,c := range columns_choosen{
		column_type := CAT
		if _,ok := samples[0][c].(float64) ; ok{
			column_type = NUMERIC
		}

		gain,value,part_l,part_r := getBestGain(samples,c,samples_labels,column_type)
		//fmt.Println("kkkkk",gain,part_l,part_r)
		if gain>=best_gain{
			best_gain = gain
			best_part_l = part_l
			best_part_r = part_r
			best_value = value
			best_column = c
		}
	}

	if best_gain>0 && len(best_part_l)>0 && len(best_part_r)>0 {
		node := &TreeNode{}
		node.Value = best_value
		node.ColumnNo = best_column
		node.Left = buildTree(getSamples(samples,best_part_l),getLabels(samples_labels,best_part_l), selected_feature_count)
		node.Right = buildTree(getSamples(samples,best_part_r),getLabels(samples_labels,best_part_r), selected_feature_count)
		return node
	}
		
	return genLeafNode(samples_labels)
	
}

func genLeafNode(labels []float64) *TreeNode{
	total := 0.0
	for _,x := range labels{
		total += x
	}
	avg := total /float64(len(labels))
	node := &TreeNode{}
	node.Label = avg
	//fmt.Println(node)
	return node
}


func predicate(node *TreeNode, input []interface{}) float64{
	if node.Value == nil{ //leaf node
		return node.Label
	}

	c := node.ColumnNo
	value := input[c]

	switch value.(type){
	case float64:
		if value.(float64)<=node.Value.(float64) && node.Left!=nil{
			return predicate(node.Left,input)
		}else if node.Right!=nil{
			return predicate(node.Right,input)
		}
	case string:
		if value==node.Value && node.Left!=nil{
			return predicate(node.Left,input)
		}else if node.Right != nil{
			return predicate(node.Right,input)
		}
	}

	return 0
}


func BuildTree(inputs [][]interface{}, labels []float64, samples_count,selected_feature_count int) *Tree{

	samples := make([][]interface{},samples_count)
	samples_labels := make([]float64,samples_count)
	for i:=0;i<samples_count;i++{
		j := int(rand.Float64()*float64(len(inputs)))
		samples[i] = inputs[j]
		samples_labels[i] = labels[j]
	}

	tree := &Tree{}
	tree.Root = buildTree(samples,samples_labels, selected_feature_count)

	return tree
}



func PredicateTree(tree *Tree, input []interface{}) float64{
	return predicate(tree.Root,input)
}
