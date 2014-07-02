RF.go
=====

* RF.go is an implementation of Random Forest in GoLang. 

		Random forest (or random forests) is an ensemble classifier that consists of many
		 decision trees and outputs the class that is the mode of the classes output by 
		 individual trees. http://en.wikipedia.org/wiki/Random_forest


* RF.go can train trees in parallel mode with assigning go-routine to each decision tree, which can utilize multiple-core CPU efficiently.


* On the famous dataset MNIST, RF.go can get 2.8% error rate with configuration of 100 trees 

* RF.go supports both Classification and Regression. The examples can be found in the repository.

* RF.go supports dumpping and loading the forest data structure between RAM and disk, in a JSON format file

### Installation
1. [Install Go](http://www.golang.org) 
2. ```$ go get github.com/fxsjy/RF.go/RF ``` This will put the binary in ```$GOROOT/bin```
