package main

import "fmt"

//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//
// 实现 MyQueue 类：
//
//
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
//
//
//
//
// 说明：
//
//
// 你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
//
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//
//
//
//
// 进阶：
//
//
// 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
//
//
//
//
// 示例：
//
//
//输入：
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出：
//[null, null, null, 1, 1, false]
//
//解释：
//MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
//
//
//
//
//
//
//
// 提示：
//
//
// 1 <= x <= 9
// 最多调用 100 次 push、pop、peek 和 empty
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
//
// Related Topics 栈 设计 队列
// 👍 429 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
	stack []int
	back  []int
}

/** Initialize your data structure here. */
//思路：双栈队列，其中一个栈用于push，另一个栈在pop和peek的时候，将push栈中的元素依次弹出来 push到当前栈，然后再弹出元素
func ConstructorV1() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	//back not nil, return
	if len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		return val
	}

	//back is nil, pop stack to back
	for len(this.stack) > 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]

		this.back = append(this.back, val)
	}

	//递归pop，使得方法显得简洁明了
	return this.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	//back not nil, return
	if len(this.back) != 0 {
		return this.back[len(this.back)-1]
	}

	//back is nil, pop stack to back
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]

		this.back = append(this.back, val)
	}

	return this.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	queue := ConstructorV1()

	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
