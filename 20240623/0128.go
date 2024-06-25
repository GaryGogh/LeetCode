func longestConsecutive(nums []int) int {
    HM:=map[int]bool{}
    for _,num:=range nums{
        HM[num]=true
    }
    res:=0
    for i:=range HM{
        if!HM[i-1]{
            cur:=1
            for HM[i+1]{
                i=i+1
                cur=cur+1
            }
        res=max(res,cur)
    }
    }
    return res
}