func getAverage(arr []int, int size) float32 {
   var i int
   var avg, sum float32  

   for i = 0; i < size; ++i {
      sum += arr[i]
   }

   avg = sum / size
   return avg;
}

//This program is not for execution
