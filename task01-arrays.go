package homework

func average(input [15]float32) (result float32) {

	var summ float32

   	for i := 0; i < 15; i++ 
   	{
   		summ += input[i]
   	}

   	return summ/15
}
