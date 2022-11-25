package testing

import "testing"

func TestSumOfNumbers(t *testing.T) {
	t.Run("SimpleTest", func(t *testing.T) {
		t.Parallel() //асинхронный запуск
		col := SumOfNumbers(1, 2, 3, 4, 5)
		result := 15
		if col != result {
			t.Errorf("expected value %d does not match %d", col, result)
		}
	})
	t.Run("MediumTest", func(t *testing.T) {
		t.Parallel() //асинхронный запуск
		col := SumOfNumbers(10, 24, 388, 48, 965)
		result := 15145
		if col != result {
			t.Errorf("expected value %d does not match %d", col, result)
		}
	})

}

//группа тестов
// func Test”NameFunctionOfTest” (t *testing.T) {
// 	t.Run(“SimpleTest”,func(t*testing.T){
// 	col:= NameFunction(2,2 )
// 	expectedValue:=4
// 	if col !=expectedValue{
// 			t.Error(“expected value %d does not match %d ”,col,expectedValue”)
// 	   }

// 	})
// 	t.Run(“MediumTest”,func(t*testing.T){
// 	col:= NameFunction(223,245 )
// 	expectedValue:=468
// 	if col !=expectedValue{
// 			t.Error(““expected value %d does not match %d ”,col,expectedValue”)
// 	   }

// 	})
// 	t.Run(“HardTest”,func(t*testing.T){
// 	col:= NameFunction(24422433,265483 )
// 	expectedValue:=4
// 	if col !=expectedValue{
// 			t.Error(“expected value %d does not match %d ”,col,expectedValue)
// 	   }

// 	})

// 	}
