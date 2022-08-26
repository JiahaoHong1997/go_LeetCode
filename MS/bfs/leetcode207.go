/*
 * @Descripttion: 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 
提示：
1 <= numCourses <= 10^5
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对互不相同
 * @Solution:
 1.inNodeNum用来记录每一门课程的前置课程
 2.outNode用来记录每一门课程的后置课程数量
 3.遍历inNodeNum找到前置课程数量为0的课程先学（入队）
 4.BFS搜索队列中的课程的后置课程，将他们的前置课程数量减1，并将学完了所有前置课程的课入队
 5.使用变量count统计所有学过的课程数来判断能否学完所有课程
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-08-26 16:24:58
 * @LastEditors: 洪笳淏
 */
package bfs

func canFinish(numCourses int, prerequisites [][]int) bool {
	outNode := make([][]int, numCourses)
	inNodeNum := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		pre, after := prerequisites[i][1], prerequisites[i][0]
		outNode[after] = append(outNode[after], pre)
		inNodeNum[pre]++
	}

	q := []int{}
	for i := 0; i < len(inNodeNum); i++ {
		if inNodeNum[i] == 0 {
			q = append(q, i)
		}
	}
	count := len(q)
	// fmt.Println(count)

	for len(q) > 0 {
		p := make([]int, 0)

		for i := 0; i < len(q); i++ {
			node := q[i]
			for _, v := range outNode[node] {
				inNodeNum[v]--
				if inNodeNum[v] == 0 {
					p = append(p, v)
				}
			}
		}
		q = p
		count += len(q)
	}
	// fmt.Println(count)
	if count == numCourses {
		return true
	}

	return false
}
