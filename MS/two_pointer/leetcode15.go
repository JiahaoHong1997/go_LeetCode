/*
 * @Descripttion: 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = [0,1,1]
输出：[]

示例 3：
输入：nums = [0,0,0]
输出：[[0,0,0]]
 * @Solution: 
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-25 23:18:40
 * @LastEditors: 洪笳淏
 */
package twopointer
import "sort"

func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    if len(nums) == 3 {
        if nums[0]+nums[1]+nums[2] == 0 {
            res = append(res, nums)
        }   
        return res
    }
    sort.Ints(nums)

   for first:=0; first<len(nums)-2; first++ {
       if first != 0 && nums[first] == nums[first-1] {
           continue
       }

        num1 := nums[first]
        l, r := first+1, len(nums)-1
       for l < r {
           num2, num3 := nums[l], nums[r]
           sum := num1 + num2 + num3
           if sum == 0 {
               res = append(res, []int{num1,num2,num3})
               for l < r && nums[l] <= num2 {
                   l++
               }
               for l < r && nums[r] >= num3 {
                   r--
               }
           } else if sum < 0 {
               l++
           } else {
               r--
           }
       }
   } 

   return res
}