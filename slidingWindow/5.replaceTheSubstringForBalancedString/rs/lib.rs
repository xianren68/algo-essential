pub fn balanced_string(s: String) -> i32 {
    // 获取字符串长度
    let n = s.len();
    // 创建一个长度为4的数组，用于存放字符串中每个字符的索引
    let mut arr = Vec::with_capacity(n);
    // 创建一个长度为4的数组，用于存放每个字符的计数
    let mut cnts:[usize;4] = [0;4];
    // 遍历字符串中的每一个字符
    for item in s.chars(){
      match item{
        'Q' => {
          // 将字符串中Q的索引存入arr数组中
          arr.push(0);
          // 将Q的计数加1
          cnts[0]+=1;
        },
        'W' => {
          // 将字符串中W的索引存入arr数组中
          arr.push(1);
          // 将W的计数加1
          cnts[1]+=1;
        },
        'E' =>{
          // 将字符串中E的索引存入arr数组中
          arr.push(2);
          // 将E的计数加1
          cnts[2]+=1;
        },
        'R' => {
          // 将字符串中R的索引存入arr数组中
          arr.push(3);
          // 将R的计数加1
          cnts[3] += 1;
        },
        _ => ()
      };
    }
    // 初始化左右指针
    let mut l = 0;
    let mut r = 0;
    // 初始化答案
    let mut ans = n;
    // 计算每个字符的期望值
    let target = n/4;
    // 如果每个字符的计数都大于期望值，则返回0
    if ok(&cnts,0,target){
      return 0;
    }
    // 遍历字符串中的每一个字符
    while r < n {
      // 将当前字符的计数减1
      cnts[arr[r]] -= 1;
      // 如果当前字符的计数小于期望值，则更新答案
      while ok(&cnts,r-l+1,target){
        ans = min(ans,r-l+1);
        // 将当前字符的计数加1
        cnts[arr[l]] += 1;
        // 将左指针右移
        l += 1;
      }
      // 将右指针右移
      r += 1;
    }
    // 返回答案
    return ans as i32
      
    
  }
  // 判断当前字符的计数是否小于期望值
  fn ok(cnts:&[usize],mut leng:usize,target:usize)->bool{
    // 遍历每一个字符
    for i in 0..4{
      // 如果当前字符的计数大于期望值，则返回false
      if cnts[i] > target {
        return false
      }
      // 将当前字符的计数减1
      leng -= cnts[i];
    }
    // 如果当前字符的计数都小于期望值，则返回true
    leng == 0
  }
  // 返回两个数字中的最小值
  fn min(a:usize,b:usize)->usize{
    if a > b {b} else {a}
  }