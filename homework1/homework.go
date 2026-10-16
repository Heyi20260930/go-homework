package homework01

import ("fmt"; "sort")

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	counts := make(map[int]int)
	for _,num :=range nums{
		counts[num]++
	}
	for num,count := range counts{
		fmt.Println("\n-----------------------------------------------")
		fmt.Printf("num: %d\n", num)
		fmt.Printf("count: %d\n", count)
		if count == 1 {
			return num
		}
	}
	return -1
}

// 2. Is Palindrome (回文数)
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool{
	if x<0 || (x%10 ==0 && x !=0){
		fmt.Println("\n$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
		fmt.Printf("x: %d\n", x)
		return false
	}
	
	revertedNumbler := 0
	for x > revertedNumbler {
			revertedNumbler = revertedNumbler*10 + x%10
			x /= 10
	}
	fmt.Println("\n**************************************************")
	fmt.Printf("x: %d\n", x)
	fmt.Printf("revertedNumbler: %d\n", revertedNumbler)
	return x == revertedNumbler || x == revertedNumbler/10
}

// 3. Is Valid Parentheses (有效的括号)
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool{
	// 创建一个rune切片来模拟栈
	stack := []rune{}
	
	// 定义括号匹配的映射关系（右括号 -> 左括号）
	// pairs这个map解决的核心问题是：给定一个右括号，快速查出它对应的左括号是什么。
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	// 遍历字符串：
	for _, char := range s {
		// 如果是左括号，压入栈
		if char == '(' || char == '{' || char =='['{
			stack = append(stack, char)
		}else if char == ')' || char =='}' || char ==']'{
			// 如果是右括号，先检查栈是否为空？
			if len(stack) == 0{
				return false // 说明没有对应的左括号
			}
			
			// 如果不为空，再检查栈顶元素是否是对应的左括号
			if stack[len(stack)-1] != pairs[char]{
				return false // 括号不匹配
			}
			
			// 匹配成功，弹出栈顶元素（切片截取，去掉最后一个元素）
			stack = stack[:len(stack)-1]
		}
	}
	
	// 遍历结束后，如果栈为空，说明全部匹配成功
	return len(stack) == 0
	
}

// 4.最长公共前缀
// 查询字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string{
	// 1.边界条件：如果切片为空，直接返回空字符串
	if len(strs) == 0 {
		return ""
	}
	
	// 2.纵向扫描：以第一个字符串为基准，一列一列地比较
	for i:=0; i< len(strs[0]); i++ {
		// 取出第一个字符串的第i个字符（字节）
		char := strs[0][i]
		
		// 3.遍历其余的字符串，检查它们的第i个字符是否也是char
		for j:=1; j < len(strs); j++ {
			// 如果当前字符串的长度不够或者字符不匹配
			if i >= len(strs[j]) || strs[j][i] != char {
				// 截取第一个字符串的前i个字符作为结果返回
				return strs[0][:i]
			}
		}
	}
	
	// 4.如果第一个字符串全部匹配完成，说明它就是最长的公共前缀
	return strs[0]
}


// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// 1.从数组的最后一位（最低位）开始向前遍历
	for i := len(digits)-1; i >= 0; i-- {
		// 2.如果当前位小于9，直接加1并返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		
		// 2.如果当前位是9，加1会变成10，需要进位：把当前位设置位0，并继续循环去处理前一位
		digits[i] = 0
	}
	
	// 如果循环全部都没有返回，说明数组中数字全是9（比如[9,9,9]）
	// 3.加1后变成[1,0,0,0],需要在签名插入1
	return append([]int{1},digits...)
}



// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// 1.边界条件：如果数组为空，直接返回0
	if len(nums) == 0 {
		return 0
	}
	
	// 2.初始化慢指针slow，指向当前已去重部分的最后一个元素
	slow := 0
	
	// 3.快指针fast，从第二个元素开始遍历整个数组
	for fast := 1; fast < len(nums); fast++ {
		// 如果快指针指向的元素与慢指针指向的元素不同，说明发现了新元素
		if nums[fast] != nums[slow]{
			// 慢指针加1
			slow++
			// 新元素复制到慢指针的位置
			nums[slow] = nums[fast]
		}
	}
	
	// 4.新数组的长度就是慢指针的索引+1
	return slow + 1
}


// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 1.边界条件：如果数组为空或只有一个区间，直接返回
	if len(intervals) <= 1 {
		return intervals
	}
	
	// 2.按照区间的左端点（start）进行升序排列
	sort.Slice(intervals, func(i,j int) bool {
		// 这里是排序规则：第一个区间的左端值一定小于第二个区间的左端值
		return intervals[i][0] < intervals[j][0]
	})
	
	// 3. 初始化结果集，先把第一个区间放进去
	result := [][]int{intervals[0]}
	
	// 4.从第二个区间开始遍历
	for i:=1; i < len(intervals); i++ {
		curr := intervals[i]
		last := result[len(result)-1]
		
		if curr[0] <= last[1]{
			if curr[1] > last[1]{
				last[1] = curr[1]
			}
		}else {
			result = append(result,curr)
		}
	}
	
	return result
}



// 8. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值target的那两个整数，并返回它们的数组下标。
func TwoSum(nums []int, target int) []int {
	// 创建一个map，key 存放数组中的数字，value 存放该数组对应的索引
	m := make(map[int]int)
	
	// 遍历数组
	for i, num := range nums {
		// 计算当前数字距离目标还差多少？
		complement := target -num
		
		// 检查这个差值是否已经在 map 中出现过
		// Go 语言特有的语法：j, ok := m[complement]
		// 如果存在，ok 为 true，j 就是它之前的索引
		if j, ok := m[complement]; ok{
			return []int{j,i}
		}
		
		// 如果没有找到，把当前数字和索引存入map，供后面的数字匹配
		m[num] = i
	}
	
	// 兜底返回
	return nil
}