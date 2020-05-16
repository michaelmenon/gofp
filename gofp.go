package gofp

import (
	"errors"
)

type Mapper func(interface{})interface{}
type Filterfunc func(interface{})bool
type Reducefunc func(a,b interface{})interface{}
//type Filter

type fpCollection interface{
	Map(Mapper) []interface{}
	Filter() []interface{}
}

type collection struct{
	list []interface{}
}

func NewCollection(sl []interface{})(*collection,error) {
	if sl == nil{
		return nil,errors.New("Inavild input")
	}
	 cln := new(collection)
	 cln.list = sl
	 return cln,nil
	
}

func (c *collection)Map(f Mapper)*collection{

	if f ==nil{
		return nil
	}
	cln := new(collection)

	sl := make([]interface{}, 0, len(c.list))
	
	for _,val := range c.list{
		sl = append(sl,f(val))
	}
	cln.list = sl
	return cln

}

func (c *collection)Filter(f Filterfunc) *collection{
	if f ==nil{
		return nil
	}
	cln := new(collection)

	sl := make([]interface{}, 0, len(c.list))
	
	for _,val := range c.list{
		if f(val){
			sl = append(sl,val)
		}
		
	}
	cln.list = sl
	return cln

}

func (c *collection)Reduce(identity interface{},reducer Reducefunc)interface{}{

	res := identity
	for _,val := range c.list{

		res = reducer(val,res)
	}
	return res
}