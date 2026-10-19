package main

import "testing"

func TestIncreaseByTen(t *testing.T){
	num := 5
	expected := 15
	
	IncreaseByTen(&num)
	
	if num != expected{
		t.Errorf("IncreaseByTen(&num) 修改后 num = %d; want %d", num, expected)
	}
}