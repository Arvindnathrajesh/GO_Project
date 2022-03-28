/*
Created by : Arvind Nath Rajesh
contact : arvindnathr@gmail.com

*/
package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	// "reflect"
)

var s int
var ans string=""
var words_index = map[string][]int{}
var index_words = map[int][]string{}
var index_order = map[int]int{}
var count int=0
var Total_Logs_Count int=0
var initial_log_no int=1

func removeIndexesInWords_index(index int){ 
	oldWordsArr := index_words[index]
	noOldWords := len(oldWordsArr)
	for i:=0;i<noOldWords;i++{				// To remove each old word(indexes) present in words_index
		oldWord := oldWordsArr[i]
		indexArr := words_index[oldWord]
		words_index[oldWord] = removeElement(indexArr,index) // To remove index from the old words 
		if len(words_index[oldWord])==0{
			delete(words_index, oldWord)
		}
	}

}
//TO add the new indexes for the new words into the maps words_index if the index value is repeated
func addElemToWords_index(noOfWords int,index int,words []string){ 
	for i:=2;i<noOfWords;i++{

		position := index_order[index]
		word := words[i]
		indexList := words_index[word]
		indexListLen:=len(indexList)
		j:=0

		for ;j<indexListLen;j++{                     // To add new index in the correct position of words_index
			if  position>index_order[indexList[j]]{
				break
			}
		}

		if indexListLen == j{
			words_index[word] = append(words_index[word], index)
		} else if j==0{
			words_index[word] = append([]int{index}, words_index[word]...)
		}else {
			words_index[word] = append(words_index[word][:j+1], words_index[word][j:]...)
			words_index[word][j] = index
		}

	}
}
func removeElement[T comparable](l []T, item T) []T { //Removing an element from an array
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
//Function to ADD elements into the maps when the user ADD. eg: ADD 25 the first
func ADD(words []string, index int) { 
	
	_, isPresent := index_words[index]
	noOfWords:=len(words)

    if isPresent {                                       //if the index(key in qstn) is already present
		removeIndexesInWords_index(index)
		addElemToWords_index(noOfWords, index, words) // to add new values to the words_index

	} else {										//if the index(key in qstn) is not present
		Total_Logs_Count++;										
		count++
		index_order[index]=count
		for i:=2;i<noOfWords;i++{					// To add each words 
			word := words[i]
			words_index[word] = append([]int{index}, words_index[word]...)  //adding indexes to words_index for each word
		}
		if Total_Logs_Count>s{                       //if the max logs are more than s
			var minKey int
			for key, value := range index_order {    // get the log number(index) which is initially added
				if value==initial_log_no {
					minKey=key
					initial_log_no++
					break
				}
			}
			removeIndexesInWords_index(minKey)    // deleteing the initial most log
			delete(index_words, minKey)
			delete(index_order, minKey)
			Total_Logs_Count--   
		}		
	}
	index_words[index] = append(words[:0], words[2:]...) //Adding words to each index(key in qstn)
}
// To search for the given word in maps. eg: SEARCH second 2
func SEARCH(word string ,limit int) {

	_, found := words_index[word]
	if found{
		if len(words_index[word])>0{
			
			for i:=0;i<limit && i<len(words_index[word] );i++{
				value := strconv.Itoa(words_index[word][i])
				ans = ans + value +" "
			}
			
			ans=ans+ "\n"
			return
		}	
	}
	ans=ans+"NONE"+"\n"
}
func main(){                                       // main function

    _, err := fmt.Scanf("%d", &s)
	if err != nil{
		os.Exit(4)
	}
	for ;;{

		inputReader := bufio.NewReader(os.Stdin)   //Taking input value
		input, _ := inputReader.ReadString('\n')

		words := []string{}
		word :=""
		for i:=0;i<len(input);i++{               //Getting each word by word from input
			if input[i]==' '{
				words=append(words,word)
				word=""
			} else{
				word=word + string(input[i])				
			}			
		}
		if len(word)>=2{
			word = word[:len(word)-2]
		}
		words=append(words,word)

		if words[0]=="END"{                   // if input is END, 
			break
		} else if words[0]=="ADD"{                // if input to ADD
			index, err := strconv.Atoi(words[1])
			if err == nil{
				ADD(words, index)
			}			
		} else if words[0]=="SEARCH"{           // if input to SEARCH
			index, err := strconv.Atoi(words[2])
			if err == nil{
				SEARCH(words[1], index)
			}			
		}
	}
	// To print the MAP values if necessary 

	// for key, value := range index_words {
	// 	fmt.Println(key, ":", value)
	// }
	// for key, value := range words_index {
	// 	fmt.Println(key, ":", value)
	// }
	// for key, value := range index_order {
	// 	fmt.Println(key, ":", value)
	// }

	ans=ans+"END"
	fmt.Println(ans)
}