type SnapshotArray struct {
    length int
    modifications [][][]int
    snapshots int
}


func Constructor(length int) SnapshotArray {
    return SnapshotArray{
        length: length,
        modifications: make([][][]int, length),
        snapshots: -1,
    }
}


func (this *SnapshotArray) Set(index int, val int)  {
    if len(this.modifications[index]) == 0 {
        this.modifications[index] = append(this.modifications[index], []int{this.snapshots, val})
        return
    }

    if this.modifications[index][len(this.modifications[index])-1][0] == this.snapshots {
        this.modifications[index][len(this.modifications[index])-1][1] = val
    } else {
        this.modifications[index] = append(this.modifications[index], []int{this.snapshots, val})
    }
}


func (this *SnapshotArray) Snap() int {
    this.snapshots++
    return this.snapshots
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    if len(this.modifications[index]) == 0 {
        return 0
    }

    start, end := 0, len(this.modifications[index])-1
    for start < end {
        half := (start + end) / 2

        if this.modifications[index][half][0] >= snap_id {
            end = half
        } else {
            start = half + 1
        }
    }

    if this.modifications[index][start][0] < snap_id  {
        return this.modifications[index][start][1]
    } else {
        if start == 0 {
            return 0
        } else {
            return this.modifications[index][start-1][1]
        }
    }
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

