class Solution {
    public int[] asteroidCollision(int[] asteroids) {
        Stack<Integer> stack = new Stack<>();
        for (int asteroid : asteroids) {
            if (stack.empty() || asteroid > 0 || asteroid * stack.peek() > 0) {
                stack.push(asteroid);
                continue;
            }

            boolean flag = true;
            while (!stack.empty() && stack.peek() > 0) {
                if (Math.abs(asteroid) > Math.abs(stack.peek())) {
                    stack.pop();
                    continue;
                }

                if (Math.abs(asteroid) == Math.abs(stack.peek())) {
                    stack.pop();
                    flag = false;
                    break;
                }

                flag = false;
                break;
            }

            if (flag) {
                stack.push(asteroid);
            }
        }

        return stack.stream().mapToInt(Integer::intValue).toArray();
    }
}

