func longestConsecutive(nums []int) int {
    HM:=map[int]bool {}
    for _,i:=range nums{
        HM[i]=true
        
    }
    res:=0
    for num:=range HM{
        if !HM[num-1]{
            cur:=1
           for HM[num+1]{
                num+=1
                cur+=1           
           } 
           if res<cur{
            res=cur
           }
        }
    }
    return res
}