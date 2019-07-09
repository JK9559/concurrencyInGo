# 切片：  
切片和数组很类似，可以使用下标进行访问，如果越界，就会产生panic。但是比数组灵活，可以进行自动扩容。
```go
// runtime/slice.go
type slice struct {
    array unsafe.Pointer // 元素指针
    len   int // 长度 
    cap   int // 容量
}
```
1. slice共三个属性 指针是指向底层数组、长度表示切片可用元素的个数、容量表示底层数组的元素个数，其中容量>=长度。在底层的数组不进行扩容的情况下，容量是slice可以扩张的最大限度。
2. 底层数组是可以被多个slice同时指向的，因此对一个slice的元素进行操作是可能影响到其他的slice

##  slice的创建
1. 直接声明
```go
    var slice []int
```
go编译
```shell
go tool compile -S ch3\slice\main\create_slice.go
```
2. new方式