package stack

func isValid(s string) bool {
	m:=map[uint8]uint8{
		'}':'{',
		')':'(',
		']':'['}
	slice:=make([]uint8,0)
	for _,v:=range s{
		if v=='{'||v=='['||v=='('{
			slice=append(slice,uint8(v))
		} else{
			if len(slice)==0 || m[uint8(v)]!=slice[len(slice)-1]{
				return false
			}
			//f={ [ (
			slice=slice[:len(slice)-1]
		}
	}
	return len(slice)==0
}
