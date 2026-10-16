package homework01

import (
	"reflect";
	"testing";
	"fmt";
)

func TestSingleNumber(t *testing.T){
	// 这里输入参数
	nums := []int{4,1,2,1,2}
	
	// 调用homework.go 里的函数
	result := SingleNumber(nums)
	
	// 验证结果
	if result != 4 {
		t.Errorf("计算结果错误，期望 4，实际得到 %d", result)
		
	}
}


func TestIspalindrome(t *testing.T){
	tests := []struct{
		input int
		want  bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{12321, true},
		{0, true},
		{11, true},
	}
	
	for _,tt := range tests{
		got := IsPalindrome(tt.input)
		if got != tt.want{
			t.Errorf("IsPalindrome(%d) = %v; want %v", tt.input, got, tt.want)
		}
	}
}



func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"(", false},
		{"]", false},
	}
	
	for _, tt := range tests {
		got := IsValid(tt.input)
		if got != tt.want {
			t.Errorf("IsValid(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}



func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"cir", "car"}, "c"},
		{[]string{}, ""},
	}
	
	for _, tt:= range tests {
		got := LongestCommonPrefix(tt.strs)
		if got != tt.want{
			t.Errorf("LongestCommonPrefix(%q) = %q; want %q", tt.strs, got, tt.want)
		}
	}
}


func TestPlusOne(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{0}, []int{1}},
	}

	for _, tt := range tests {
		// 为了不影响下一次测试，最好复制一份传入，或者依赖每次循环新建的切片(因为Go的切片是引用类型，PlusOne会修改原数组)
		inputCopy := make([]int, len(tt.input))
		copy(inputCopy, tt.input)
		
		got := PlusOne(inputCopy)
		fmt.Println("\n------------------------------------------------")
		fmt.Printf("got: %d\n", got)
		if len(got) != len(tt.want) {
			t.Errorf("PlusOne(%v) = %v; want %v", tt.input, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("PlusOne(%v) = %v; want %v", tt.input, got, tt.want)
				break
			}
		}
	}
}



func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input []int
		wantK int
		want  []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 1, 1, 1}, 1, []int{1}},
		{[]int{1}, 1, []int{1}},
		{[]int{}, 0, []int{}},
	}

	for _, tt := range tests {
		// 关键步骤：因为函数会原地修改切片，所以必须拷贝一份传入，防止影响后续循环
		inputCopy := make([]int, len(tt.input))
		copy(inputCopy, tt.input)

		gotK := RemoveDuplicates(inputCopy)

		if gotK != tt.wantK {
			t.Errorf("RemoveDuplicates(%v) 返回长度 = %d; want %d", tt.input, gotK, tt.wantK)
			continue
		}

		// 截取前 gotK 个元素进行比较
		gotArr := inputCopy[:gotK]
		if len(gotArr) != len(tt.want) {
			t.Errorf("RemoveDuplicates(%v) 结果 = %v; want %v", tt.input, gotArr, tt.want)
			continue
		}
		for i := range gotArr {
			if gotArr[i] != tt.want[i] {
				t.Errorf("RemoveDuplicates(%v) 结果 = %v; want %v", tt.input, gotArr, tt.want)
				break
			}
		}
	}
}



func TestMerge(t *testing.T) {
	tests := []struct {
		input [][]int
		want  [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		{[][]int{{1, 4}}, [][]int{{1, 4}}},
		{[][]int{}, [][]int{}},
	}

	for _, tt := range tests {
		// 因为 sort.Slice 会原地修改输入切片，所以必须拷贝一份
		inputCopy := make([][]int, len(tt.input))
		for i := range tt.input {
			// 内层切片也要重新分配，实现深拷贝
			inputCopy[i] = make([]int, len(tt.input[i]))
			copy(inputCopy[i], tt.input[i])
		}

		got := Merge(inputCopy)

		// 使用 reflect.DeepEqual 比较二维切片内容是否一致
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Merge(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}



func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums []int
		target int
		want []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	
	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)
		// 因为返回的是切片，不能直接使用==比较，需要使用reflect.DeepEqual比较
		if !reflect.DeepEqual(got, tt.want){
			t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
