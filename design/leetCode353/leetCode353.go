package leetcode353

type SnakeGame struct {
    snakePos    [][]int
    width       int
    height      int
    food        [][]int
    foodIndex   int
}


func Constructor(width int, height int, food [][]int) SnakeGame {
    return SnakeGame {
        snakePos :  [][]int{{0,0}},
        width    :   width,
        height   :   height,
        food     :   food,
        foodIndex :  0,
    }
}


func (this *SnakeGame) Move(direction string) int {
    
    h, w := 0, 0
    switch direction {
        case "R":
            h, w = this.snakePos[0][0], this.snakePos[0][1]+1
        case "L":
            h, w = this.snakePos[0][0], this.snakePos[0][1]-1
        case "U":
            h, w = this.snakePos[0][0]-1, this.snakePos[0][1]
        case "D":
            h, w = this.snakePos[0][0]+1, this.snakePos[0][1]
    }

    if (h < 0 || h >= this.height) || (w < 0 || w >= this.width) {
        return -1
    }

    this.snakePos = append([][]int{{h,w}}, this.snakePos...)
    if this.foodIndex < len(this.food) {
        if h == this.food[this.foodIndex][0] && w == this.food[this.foodIndex][1] {   
            this.foodIndex++
            return this.foodIndex
        } else {
            this.snakePos = this.snakePos[:len(this.snakePos)-1]
            for i:=1; i<len(this.snakePos); i++ {
                if this.snakePos[i][0] == h && this.snakePos[i][1] == w {
                    return -1
                }
            }
            return this.foodIndex
        }
    } else {
        this.snakePos = this.snakePos[:len(this.snakePos)-1]
        for i:=1; i<len(this.snakePos); i++ {
            if this.snakePos[i][0] == h && this.snakePos[i][1] == w {
                return -1
            }
        }
    }

    return this.foodIndex
}


/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */