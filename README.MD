RF.go
=====

* RF.go is a implementaion of Random Forest in GoLang. 

	Random forest (or random forests) is an ensemble classifier that consists of many decision trees and outputs the class that is the mode of the classes output by individual trees.

* RF.go can train trees in parallel mode with assigning go-routine to each decision tree, which can utilize multiple core efficently.


* On the famous dataset MNIST, RF.go can get 2.8% error rate with configuration of 100 trees 

* RF.go supports both Classification and Regression. The examples can be found in the repository.


