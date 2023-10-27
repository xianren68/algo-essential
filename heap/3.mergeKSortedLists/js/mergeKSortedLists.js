function ListNode(val, next) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
}

// 堆
const heap = []

/**
 * 入堆
 * @param {ListNode} node
 */
function pushHeap(node) {
    let i = heap.length
    heap.push(node)
    while (i > 0 && heap[i].val < heap[(i - 1) >> 1].val) {
        swap(i, (i - 1) >> 1)
        i = (i - 1) >> 1
    }
}

/**
 * 出堆
 * @return {ListNode}
 */
function popHeap() {
    const res = heap[0]
    heap[0] = heap.at(-1)
    heap.pop()
    let i = 0
    let l = i * 2 + 1
    let best
    while (l < heap.length) {
        if (l + 1 < heap.length && heap[l + 1].val < heap[l].val) {
            best = l + 1
        } else {
            best = l
        }
        if (heap[i].val > heap[best].val) {
            swap(i, best)
        }
        i = best
        l = i * 2 + 1
    }
    return res

}

/**
 * 交换位置
 * @param {number} i
 * @param {number} j
 */
function swap(i, j) {
    let temp = heap[i]
    heap[i] = heap[j]
    heap[j] = temp
}

/**
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function (lists) {
    if (lists.length === 0) {
        return null
    }
    // 将链表全部入小根堆
    for (let i of lists) {
        if (i) {
            pushHeap(i)
        }
    }
    if (heap.length === 0) {
        return null
    }
    const head = popHeap()
    if (head.next) {
        pushHeap(head.next)
    }
    let node = head
    while (heap.length > 0) {
        let n = popHeap()
        if (n.next) {
            pushHeap(n.next)
        }
        node.next = n
        node = node.next
    }
    return head
};