func intersection(nums1 []int, nums2 []int) []int {
   var arr []int
   
   for i := 0; i < len(nums1); i++ {
    for j := 0; j < len(nums2); j++ {
        if nums1[i] == nums2[j] {
            found := false
            for k := 0; k < len(arr); k++ {
                if arr[k] == nums1[i] {
                    found = true
                }
            }
            if !found {
                arr = append(arr, nums1[i])
            }
        }
    }
   }

   fmt.Println(arr)
   return arr 
}