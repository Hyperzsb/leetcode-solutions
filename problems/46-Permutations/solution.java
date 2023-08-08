class Solution {
    private void backtrace(LinkedList<Integer> curr, LinkedList<Integer> remain, List<List<Integer>> result) {
        if (remain.size() == 1) {
            curr.add(remain.getFirst());
            result.add(new LinkedList(curr));
            curr.removeLast();

            return;
        }

        for (int i = 0; i < remain.size(); i++) {
            Integer removed = remain.remove(i);
            curr.add(removed);
            backtrace(curr, remain, result);
            curr.removeLast();
            remain.add(i, removed);
        }
    }
    
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> result = new LinkedList<>();

        LinkedList<Integer> curr = new LinkedList<>(), remain = new LinkedList<>();
        for (int num : nums) {
            remain.add(num);
        }
        
        backtrace(curr, remain, result);

        return result;
    }
}

