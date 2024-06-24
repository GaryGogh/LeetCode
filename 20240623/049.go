func groupAnagrams(strs []string) [][]string {
    HM:=map[string][]string{}
    for _,i:=range strs{
        t:=[]byte(i)
        slices.Sort(t)
        afs:=string(t)
        HM[afs]=append(HM[afs],i)
    }
    res:=[][]string{}
    for _,n:=range HM{
        res=append(res,n)
    }
    return res
}
