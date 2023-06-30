class Solution {
    class Pair {
        int u;
        int v;

        Pair(int u, int v) {
            this.u = u;
            this.v = v;
        }
    }

    class PairComparator implements Comparator<Pair> {
        public int compare(Pair a, Pair b) {
            return (b.u + b.v) - (a.u + a.v);
        }
    }

    public List<List<Integer>> kSmallestPairs(int[] nums1, int[] nums2, int k) {
        PriorityQueue<Pair> ppq = new PriorityQueue<Pair>(new PairComparator());

        for (int i = 0; i < nums1.length && i < k; i++) {
            for (int j = 0; j < nums2.length && j < k; j++) {
                if (ppq.size() < k) {
                    ppq.add(new Pair(nums1[i], nums2[j]));
                } else {
                    Pair p = ppq.peek();
                    if (p.u + p.v > nums1[i] + nums2[j]) {
                        ppq.add(new Pair(nums1[i], nums2[j]));
                        ppq.poll();
                    } else {
                        break;
                    }   
                }
            }
        }

        List<List<Integer>> result = new ArrayList<>();

        Iterator ppqi = ppq.iterator();
        while (ppqi.hasNext()) {
            Pair p = (Pair)ppqi.next();
            result.add(Arrays.asList(p.u, p.v));
        }

        return result;
    }
}

