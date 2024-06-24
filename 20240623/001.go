func twoSum(nums []int, target int) []int {
    HM:=map[int]int{}
    for m,n:=range nums{
        if i,ok:=HM[target-n];ok{
            return []int{m,i}
        }else{
            HM[n]=m
        }
    }
    return nil
}
