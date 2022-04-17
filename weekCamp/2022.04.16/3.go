package weekCamp

type ATM struct {
    paperCount      map[int]int
    totalCount      int
    paperAmount     []int
    maxIndex      int
}


func Constructor() ATM {
    return ATM {
        paperCount : map[int]int {
            20  :   0,
            50  :   0,
            100 :   0,
            200 :   0,
            500 :   0,
        },
        totalCount  :   0,
        paperAmount :   []int{20,50,100,200,500},
        maxIndex   :   -1,
    }
}


func (this *ATM) Deposit(banknotesCount []int)  {
    this.paperCount[20] += banknotesCount[0]
    this.paperCount[50] += banknotesCount[1]
    this.paperCount[100] += banknotesCount[2]
    this.paperCount[200] += banknotesCount[3]
    this.paperCount[500] += banknotesCount[4]
    
    this.totalCount += 20*banknotesCount[0] + 50*banknotesCount[1] + 100*banknotesCount[2] + 200*banknotesCount[3] + 500*banknotesCount[4]
    for i:=4; i>=0; i-- {
        x := this.paperAmount[i]
        if this.paperCount[x] > 0 {
            this.maxIndex = i
            break
        }
    }
}


func (this *ATM) Withdraw(amount int) []int {
    if amount > this.totalCount {
        return []int{-1}
    }
    
    ret := make([]int, 5)
    for amount > 0 {
        
        for i:=0; i<5; i++ {
            if amount == this.paperAmount[i] && this.paperCount[this.paperAmount[i]] > 0 {
                ret[i]++
                x := this.paperAmount[i]
                this.totalCount -= x 
                this.paperCount[x]--
                if this.paperCount[x] == 0 {
                    for i:=this.maxIndex-1; i>=0; i-- {
                        y := this.paperAmount[i]
                        if this.paperCount[y] > 0 {
                            this.maxIndex = i
                            break
                        }
                    }
                }
                
                return ret
            }
        }
        
        f := 0
        if this.paperAmount[this.maxIndex] < amount {
            x := this.paperAmount[this.maxIndex]
            var n int
            if this.paperCount[x] >= amount/x {
                n = amount/x
            } else {
                n = this.paperCount[x]
            }
            amount -= x*n
            ret[this.maxIndex] += n
            this.paperCount[x] -= n
            if this.paperCount[x] == 0 {
                for i:=this.maxIndex-1; i>=0; i-- {
                    y := this.paperAmount[i]
                    if this.paperCount[y] > 0 {
                        this.maxIndex = i
                        break
                    }
                }
            }
            
            this.totalCount -= x *n
        } else {
            i := this.maxIndex-1
            for ; i>=0; i-- {
                x := this.paperAmount[i]
                if x < amount && this.paperCount[x] > 0 {
                    
                    amount -= x
                    this.totalCount -= x 
                    this.paperCount[x]--
                    ret[i]++
                    f = 1
                    break
                }
            }
            
            if f == 0 {
                this.Deposit(ret)
                
                return []int{-1}
            }
        }
    }
        
    
    if amount == 0 {
        return ret
    }
    
    this.Deposit(ret)
    return []int{-1}
}


/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */