package leetcode341

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool { return true }

func (this NestedInteger) GetInteger() int { return 0 }

func (this NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
    array   []int
    next    int    
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    
    nums := []int{}

    var initialize func([]*NestedInteger)
    initialize = func(nL []*NestedInteger) {

        for i:=0; i<len(nL); i++ {
            x := nL[i]
            if (*x).IsInteger() {
                t := (*x).GetInteger()
                nums = append(nums, t)
            } else {
                t := (*x).GetList()
                initialize(t)
            }
        }
    }
    
    initialize(nestedList)
    return &NestedIterator{
        array:  nums,
        next:   0,
    }
}

func (this *NestedIterator) Next() int {
    x := this.array[this.next]
    this.next++
    return x
}

func (this *NestedIterator) HasNext() bool {
    if this.next == len(this.array) {
        return false
    }
    return true
}