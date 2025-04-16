func merge(intervals [][]int) [][]int {
   if len (intervals)==0{
        return nil
   }
   //起始位置排序
   sort.Slice(intervals, func (i,j int)bool{
        return intervals[i][0] < intervals[j][0]
   })
        res := [][]int{intervals[0]}
   for i:=1;i<len(intervals);i++{
        last:= res[len(res)-1]
        cur:= intervals[i]
        if(cur[0]<=last[1]){
            if(cur[1]>last[1]){
                res[len(res)-1][1]=cur[1]
            }
        }else {
            res = append(res, cur)
        }
   }
       return res
}

