class Solution {
    class Pair {
        int idx;
        int val;

        Pair(int i, int v) {
            this.idx = i;
            this.val = v;
        }
    }

    class PairComparator implements Comparator<Pair> {
        public int compare(Pair a, Pair b) {
            if (a.val < b.val) {
                return -1;
            } else if (a.val > b.val) {
                return 1;
            } else {
                return a.idx - b.idx;
            }
        }
    }

    public long totalCost(int[] costs, int k, int candidates) {
        PriorityQueue<Pair> leftPQ = new PriorityQueue<>(candidates, 
            new PairComparator());
        PriorityQueue<Pair> rightPQ = new PriorityQueue<>(candidates, 
            new PairComparator());

        int leftIdx = 0, rightIdx = costs.length - 1;
        for (int i = 0; i < costs.length && i < candidates; i++) {
            leftPQ.add(new Pair(leftIdx, costs[leftIdx]));
            leftIdx++;

            rightPQ.add(new Pair(rightIdx, costs[rightIdx]));
            rightIdx--;
        }

        long result = 0;
        while (k > 0) {
            Pair left = leftPQ.peek(), right = rightPQ.peek();

            if (left.val < right.val || 
                (left.val == right.val && left.idx < right.idx)) {
                result += left.val;

                costs[left.idx] = -1;
                leftPQ.poll();

                while (leftIdx < costs.length && costs[leftIdx] == -1) {
                    leftIdx++;
                }

                if (leftIdx < costs.length) {
                    leftPQ.add(new Pair(leftIdx, costs[leftIdx]));
                    leftIdx++;
                }
            } else if (left.val > right.val ||
                (left.val == right.val && left.idx > right.idx)) {
                result += right.val;

                costs[right.idx] = -1;
                rightPQ.poll();

                while (rightIdx >= 0 && costs[rightIdx] == -1) {
                    rightIdx--;
                }

                if (rightIdx >= 0) {
                    rightPQ.add(new Pair(rightIdx, costs[rightIdx]));
                    rightIdx--;
                }
            } else {
                result += left.val;

                costs[left.idx] = -1;
                leftPQ.poll();
                rightPQ.poll();

                while (leftIdx < costs.length && costs[leftIdx] == -1) {
                    leftIdx++;
                }

                if (leftIdx < costs.length) {
                    leftPQ.add(new Pair(leftIdx, costs[leftIdx]));
                    leftIdx++;
                }

                while (rightIdx >= 0 && costs[rightIdx] == -1) {
                    rightIdx--;
                }

                if (rightIdx >= 0) {
                    rightPQ.add(new Pair(rightIdx, costs[rightIdx]));
                    rightIdx--;
                }
            }

            k--;
        }

        return result;
    }
}

