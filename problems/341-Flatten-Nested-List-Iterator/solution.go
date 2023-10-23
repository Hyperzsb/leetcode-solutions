/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    Idx int
    Nums []int
}

func generate(nestedList []*NestedInteger) []int {
    result := make([]int, 0)
    for i := range nestedList {
        if nestedList[i].IsInteger() {
            result = append(result, nestedList[i].GetInteger())
        } else {
            result = append(result, generate(nestedList[i].GetList())...)
        }
    }

    return result
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
        Idx: 0,
        Nums: generate(nestedList),
    }
}

func (this *NestedIterator) Next() int {
    result := this.Nums[this.Idx]
    this.Idx++

    return result
}

func (this *NestedIterator) HasNext() bool {
    return this.Idx < len(this.Nums)
}

