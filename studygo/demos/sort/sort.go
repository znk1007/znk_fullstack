/*
 * @Author: your name
 * @Date: 2019-12-19 22:27:34
 * @LastEditTime : 2019-12-19 22:45:39
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demos/sort/main.go
 */
package sort

import "fmt"

func main() {

	testArr := []int{21, 11, 32, 54, 60, 1, 2, 100, 30, 20, 30, 15, 15, 15}
	fmt.Println("origin arr:\n ", testArr)

	// swap(testArr, 2, 4)
	// fmt.Println("after swap: \n", testArr)

	// bubbleSort(testArr, true)
	// fmt.Println("after ascend bubble sort: \n", testArr)

	// bubbleSort(testArr, false)
	// fmt.Println("after descend bubble sort: \n", testArr)

	// selectSort(testArr, true)
	// fmt.Println("after ascend select sort: \n", testArr)

	// selectSort(testArr, false)
	// fmt.Println("after descend select sort: \n", testArr)

	// insertSort(testArr, true)
	// fmt.Println("after ascend insert sort: \n", testArr)

	// insertSort(testArr, false)
	// fmt.Println("after descend insert sort: \n", testArr)

	// shellSort(testArr, len(testArr)/2)
	// fmt.Println("after shell sort: \n", testArr)

	// heapSort(testArr)
	// fmt.Println("after heap sort: \n", testArr)

	// fmt.Println("after merge sort: \n", mergeSort(testArr))

	// quickSort(testArr)
	// fmt.Println("after quick sort: \n", testArr)

}

/*
冒泡排序
基本思想：它重复地走访过要排序的数列，
		一次比较两个元素，
		如果他们的顺序错误就把他们交换过来。
		走访数列的工作是重复地进行直到没有再需要交换，
		也就是说该数列已经排序完成。
		这个算法的名字由来是因为越小的元素会经由交换慢慢"浮"到数列的顶端。
算法步骤
比较相邻的元素。
如果第一个比第二个大，就交换他们两个。
对每一对相邻元素作同样的工作，
从开始第一对到结尾的最后一对。
这步做完后，最后的元素会是最大的数。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/
func bubbleSort(slice []int, ascend bool) {
	sliceLen := len(slice)
	if sliceLen < 2 {
		return
	}
	loop := sliceLen - 1
	for i := 0; i < loop; i++ {
		for j := 0; j < loop-i; j++ {
			if ascend {
				if slice[j] > slice[j+1] {
					swap(slice, j, j+1)
				}
			} else {
				if slice[j] < slice[j+1] {
					swap(slice, j, j+1)
				}
			}
		}
	}
}

/*
选择排序
基本思想：选择排序是一种简单直观的排序算法，
		无论什么数据进去都是 O(n²) 的时间复杂度。
		所以用到它的时候，数据规模越小越好。
		唯一的好处可能就是不占用额外的内存空间
算法步骤
1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
3.重复第二步，直到所有元素均排序完毕。
*/
func selectSort(slice []int, ascend bool) {
	sliceLen := len(slice)
	if sliceLen < 2 {
		return
	}
	for i := 0; i < sliceLen-1; i++ {
		tempIdx := i
		for j := i + 1; j < sliceLen; j++ {
			if ascend {
				if slice[j] < slice[tempIdx] {
					tempIdx = j
				}
			} else {
				if slice[j] > slice[tempIdx] {
					tempIdx = j
				}
			}
		}
		swap(slice, i, tempIdx)
	}

}

/*
插入排序
基本思想：
		是一种最简单直观的排序算法，
		它的工作原理是通过构建有序序列，
		对于未排序数据，
		在已排序序列中从后向前扫描，
		找到相应位置并插入。
插入排序和冒泡排序一样，也有一种优化算法，叫做拆半插入

算法步骤
将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
*/
func insertSort(slice []int, ascend bool) {
	for idx := range slice {
		preIdx := idx - 1
		cur := slice[idx]
		if ascend {
			for preIdx >= 0 && slice[preIdx] > cur {
				slice[preIdx+1] = slice[preIdx]
				preIdx--
			}
		} else {
			for preIdx >= 0 && slice[preIdx] < cur {
				slice[preIdx+1] = slice[preIdx]
				preIdx--
			}
		}
		slice[preIdx+1] = cur
	}
}

/*
希尔排序
基本思想：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，
		待整个序列中的记录"基本有序"时，
		再对全体记录进行依次直接插入排序

算法步骤
选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
按增量序列个数 k，对序列进行 k 趟排序；
每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，
分别对各子表进行直接插入排序。
仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/
func shellSort(slice []int, space int) {
	if space < 2 {
		space = 2
	}
	sliceLen := len(slice)
	gap := 1
	for gap < gap/space {
		gap = gap * space
	}
	for gap > 0 {
		for i := gap; i < sliceLen; i++ {
			temp := slice[i]
			j := i - gap
			for j >= 0 && slice[j] > temp {
				slice[j+gap] = slice[j]
				j -= gap
			}
			slice[j+gap] = temp
		}
		gap = gap / space
	}
}

/*
堆排序
基本思想：利用堆这种数据结构所设计的一种排序算法。
		堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：
		即子结点的键值或索引总是小于（或者大于）它的父节点。
		堆排序可以说是一种利用堆的概念来排序的选择排序。
		分为两种方法：
		大顶堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；
		小顶堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列；
		堆排序的平均时间复杂度为 Ο(nlogn)。
算法步骤：
		1.创建一个堆 H[0……n-1]；
		2.把堆首（最大值）和堆尾互换；
		3.把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
		4.重复步骤 2，直到堆的尺寸为 1。
*/
func heapSort(slice []int) {
	sliceLen := len(slice)
	buildMaxHeap(slice, sliceLen)
	for i := sliceLen - 1; i >= 0; i-- {
		swap(slice, 0, i)
		sliceLen--
		heapify(slice, 0, sliceLen)
	}
}

func buildMaxHeap(slice []int, sliceLen int) {
	for i := sliceLen / 2; i >= 0; i-- {
		heapify(slice, i, sliceLen)
	}
}

func heapify(slice []int, i, sliceLen int) {
	prev := 2*i + 1
	next := 2*i + 2
	max := i
	if prev < sliceLen && slice[prev] > slice[max] {
		max = prev
	}
	if next < sliceLen && slice[next] > slice[max] {
		max = next
	}
	if max != i {
		swap(slice, i, max)
		heapify(slice, max, sliceLen)
	}
}

/*
递归排序
基本思想：建立在归并操作上的一种有效的排序算法。
		该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
		作为一种典型的分而治之思想的算法应用，归并排序的实现由两种方法：
		1.自上而下的递归（所有递归的方法都可以用迭代重写，所以就有了第 2 种方法）；
		2.自下而上的迭代；

算法步骤：
		1.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
		2.设定两个指针，最初位置分别为两个已经排序序列的起始位置；
		3.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
		4.重复步骤 3 直到某一指针达到序列尾；
		5.将另一序列剩下的所有元素直接复制到合并序列尾

*/
func mergeSort(slice []int) []int {
	sliceLen := len(slice)
	if sliceLen < 2 {
		return slice
	}
	middle := sliceLen / 2 //二分序列
	prev := slice[0:middle]
	next := slice[middle:]
	return merge(mergeSort(prev), mergeSort(next))
}

func merge(prev []int, next []int) []int {
	var result []int
	for len(prev) != 0 && len(next) != 0 {
		if prev[0] <= next[0] {
			result = append(result, prev[0])
			prev = prev[1:]
		} else {
			result = append(result, next[0])
			next = next[1:]
		}
	}
	for len(prev) != 0 {
		result = append(result, prev[0])
		prev = prev[1:]
	}
	for len(next) != 0 {
		result = append(result, next[0])
		next = next[1:]
	}
	return result
}

/*
快速排序
基本思想：快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。
		快速排序又是一种分而治之思想在排序算法上的典型应用。
		本质上来看，快速排序应该算是在冒泡排序基础上的递归分治法。
		快速排序的最坏运行情况是 O(n²)，比如说顺序数列的快排。
		但它的平摊期望时间是 O(nlogn)，且 O(nlogn) 记号中隐含的常数因子很小，
		比复杂度稳定等于 O(nlogn) 的归并排序要小很多。
		所以，对绝大多数顺序性较弱的随机数列而言，快速排序总是优于归并排序。
算法步骤
1.从数列中挑出一个元素，称为 "基准"（pivot）;
2.重新排序数列，所有元素比基准值小的摆放在基准前面，
所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。
这个称为分区（partition）操作；
3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/
func quickSort(slice []int) {
	selfQuickSort(slice, 0, len(slice)-1)
}

func selfQuickSort(slice []int, prev, next int) {
	if prev < next {
		partitionIdx := partitionIndex(slice, prev, next)
		selfQuickSort(slice, prev, partitionIdx-1)
		selfQuickSort(slice, partitionIdx+1, next)
	}
}

func partitionIndex(slice []int, prev, next int) int {
	pivot := prev
	idx := pivot + 1
	for i := idx; i <= next; i++ {
		if slice[i] < slice[pivot] {
			swap(slice, i, idx)
			idx++
		}
	}
	swap(slice, pivot, idx-1)
	return idx - 1
}

/* 交换数组元素 */
func swap(arr []int, lhs int, rhs int) {
	if len(arr) <= lhs || len(arr) <= rhs {
		return
	}
	arr[lhs], arr[rhs] = arr[rhs], arr[lhs]
}
