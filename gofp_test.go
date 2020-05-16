package gofp

import (
	"testing"
)

func TestMap(t *testing.T){
	s := []interface{}{2,5,1,6,12,32,11,9,10}
	o := []interface{}{4,10,2,12,24,64,22,18,20}

	doubler := func(a interface{})interface{}{
		val,ok := a.(int)
		if !ok{
			t.Fatal("Wrong type")
		}
		return 2 * val
	}
	
	c,err := NewCollection(s)
	//t.Log(c.list...)
	if err!=nil{
		t.Fatal("Unable to create a collection.")
	}

	res := c.Map(doubler)
	if res== nil{
		t.Fatalf("Invalid Mapper function")
	}
	if  len(s) != len(res.list){
		t.Fatalf("Different length of output %v",res.list...)
	}
	
	for i,val := range res.list{
		if val != o[i]{
			t.Fatalf("Different value found at %d for result %v",i,res.list)
		}
	}
	t.Logf("Test Passed.... %v",res)
}

func TestFilter(t *testing.T){

	s := []interface{}{2,5,1,6,12,32,11,9,10}
	o := []interface{}{2,6,12,32,10}	

	even := func(a interface{})bool{
		val,ok := a.(int)
		if !ok{
			t.Fatal("Invalid input type.")
		}
		if val %2 ==0{
			return true
		}
		return false
	}

	c,err := NewCollection(s)
	//t.Log(c.list...)
	if err!=nil{
		t.Fatal("Unable to create a collection.")
	}

	res := c.Filter(even)
	if res== nil{
		t.Fatalf("Invalid Filter function")
	}
	if  len(o) != len(res.list){
		t.Fatalf("Different length of output %v",res.list)
	}
	
	for i,val := range res.list{
		if val != o[i]{
			t.Fatalf("Different value found at %d for result %v",i,res.list)
		}
	}
	t.Logf("Test Passed.... %v",res)
}

func TestReduce(t *testing.T){

	s := []interface{}{21,5,14,6,12,32,11,9,10}
	res := 5
	
	reducer := func(a,b interface{})interface{}{

		if a.(int) < b.(int) {
			return a.(int)
		}
		return b.(int)
	}
	c,err := NewCollection(s)
	//t.Log(c.list...)
	if err!=nil{
		t.Fatal("Unable to create a collection.")
	}
	o := c.Reduce(21,reducer)
	if o!=res{
		t.Fatalf("Test failed got %d, expected %d",o,res)
	}
	t.Logf("Test Reducer passed got %d",o)

}