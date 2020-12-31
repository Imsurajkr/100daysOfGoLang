/*
================== Arrays ============
1. creating array is simple

===== declaration of array ==========
anArray := [8]int
array[0] := 9
for i:= 0 ; i < len(array) ; i++ {
	array[i] = len-1
}
anArray := [4]int{1, 2, 3, 4}
============Multi Dimensional Array ===========
twoD := [4][4]int{{1,2,3,4}, {1,2,3,4}}
threeD := [2][2][2]int{{1,0},{-2,4},{5,-1}}

=============== Disadvantages of Array =========
1. Array cannot change its size which go array not dynamic
2. When you pass an array to a function as a parameter, you actually pass a copy of the array
3. passing a large array to a function can be pretty slow
*/

