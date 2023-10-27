function ListNode(val, next) {
     this.val = (val===undefined ? 0 : val)
     this.next = (next===undefined ? null : next)
 }

/**
* @param {ListNode} head
* @param {number} n
* @return {ListNode}
*/
var removeNthFromEnd = function(head, n) {
    if(!head){
        return Null
    }
    // 前后两个指针
    let before = head
    let after = head
    // 前面先走n步
    let i = 0
    while(before && i < n){
        before = before.next
    }
    if(!before){
        if(r < n){
            // 不足n,原样返回
            return head
        }else{
            // 恰好是n,删除头节点
            return head.next
        }
    }
    // 指向前一个节点的指针
    let pre
    // 两个指针一起走，前面的到末尾，后面的到倒数第n个
    while(before){
        pre = after
        after = after.next
        before = before.next
    }
    pre.next = after.next
    return head
};