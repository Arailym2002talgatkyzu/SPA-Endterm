package main

import (
   "bytes"
   "fmt"
   "io/ioutil"
   "sort"
)

type record2 struct {
   word    []byte
   counter int
}

func readWords( in []byte, out chan[]byte){
 var list []byte
 for _, i:=range in{
    if isLetter(i){
       if i>=65&&i<=90{
          i+=32
       }
       list=append(list, i)
    } else if len(list)!=0{
       out<-list
       list=[]byte{}
    }
 }
 close(out)
}


func isLetter(byteVal byte)bool{
   if ((byteVal >= 65 && byteVal <= 90)||(byteVal >= 97 && byteVal <= 122 ) ){
      return true
   }
   return false
}

func main()  {
   file, err := ioutil.ReadFile("mobydick.txt")
   if err != nil {
      panic(err)
   }
   output:=make(chan []byte, 20)
   go readWords(file, output )
   var records[]*record2
   for word:=range output{
      if records == nil {
         records=append(records, &record2{word: word, counter: 1})
      }else{
         found:=false
         for i, r:=range records{
            if bytes.Equal(r.word, word){
               records[i].counter++
               found=true
               break
            }
         }
         if found==false{
            records=append(records, &record2{word: word, counter: 1})
         }
      }
   }
   sort.Slice(records, func(i, j int) bool {
      return records[i].counter>records[j].counter
   })
   for i:=0;i<20;i++{
      _,_=fmt.Printf("%v %v\n", records[i].counter, string(records[i].word))
   }
}