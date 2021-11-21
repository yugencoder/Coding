package _8XX

import (
	"math"
	"sort"
)

/*
You are given n​​​​​​ tasks labeled from 0 to n - 1 represented by a 2D integer array tasks, where tasks[i] = [enqueueTimei, processingTimei] means that the i​​​​​​th​​​​ task will be available to process at enqueueTimei and will take processingTimei to finish processing.

You have a single-threaded CPU that can process at most one task at a time and will act in the following way:

If the CPU is idle and there are no available tasks to process, the CPU remains idle.
If the CPU is idle and there are available tasks, the CPU will choose the one with the shortest processing time. If multiple tasks have the same shortest processing time, it will choose the task with the smallest index.
Once a task is started, the CPU will process the entire task without stopping.
The CPU can finish a task then start a new one instantly.
Return the order in which the CPU will process the tasks.



Example 1:

Input: tasks = [[1,2],[2,4],[3,2],[4,1]]
Output: [0,2,3,1]
*/

func getOrder(tasks [][]int) []int {
	res := []int{}
	q := []int{}
	firstTaskST := math.MaxInt32
	sortedTasks := make([][]int, len(tasks))

	for i, task := range tasks {
		firstTaskST = min(firstTaskST, task[0])
		task = append(task, i)
		tasks[i] = task

		sortedTasks[i] = make([]int, len(tasks[i]))
		copy(sortedTasks[i], tasks[i])
	}

	sort.Slice(sortedTasks, func(i, j int) bool {
		return sortedTasks[i][0] < sortedTasks[j][0]
	})

	next := firstTaskST

	for {
		nSortedTasks, currTasks := getAllTasksBeforeNext(next, sortedTasks)

		if len(currTasks) == 0 {
			if len(nSortedTasks) == 0 {
				return append(res, q...)
			} else if len(nSortedTasks) > 0 && len(q)==0{
				next = nSortedTasks[0][0]
				continue
			}
		}

		sortedTasks = nSortedTasks
		q = addToQ(q, currTasks, tasks)

		if len(q) == 0 {
			continue
		}

		curr := q[0]
		q = q[1:]

		next = next + tasks[curr][1]
		res = append(res, curr)
	}

	return res
}

func addToQ(q, items []int, tasks [][]int) []int {
	q = append(q, items...)
	sort.Slice(q, func(i, j int) bool {
		i = q[i]
		j = q[j]

		if tasks[i][1] == tasks[j][1] {
			return i < j
		}

		return tasks[i][1] < tasks[j][1]
	})

	return q
}

func getAllTasksBeforeNext(next int, sortedTasks [][]int) ([][]int, []int) {
	var res []int
	for i, t := range sortedTasks {
		if t[0] <= next {
			res = append(res, t[2])
		} else {
			return sortedTasks[i:], res
		}

	}

	return [][]int{}, res
}
