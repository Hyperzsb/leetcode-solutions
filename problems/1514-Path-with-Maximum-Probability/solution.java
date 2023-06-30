class Solution {
    class Node {
        LinkedList<Integer> neighbors;
        LinkedList<Double> probs;

        Node() {
            this.neighbors = new LinkedList<Integer>();
            this.probs = new LinkedList<Double>();
        }
    }

    public double maxProbability(int n, int[][] edges, double[] succProb, int start, int end) {
        Node[] nodes = new Node[n];
        for (int i = 0; i < n; i++) {
            nodes[i] = new Node();
        }

        for (int i = 0; i < edges.length; i++) {
            nodes[edges[i][0]].neighbors.add(edges[i][1]);
            nodes[edges[i][0]].probs.add(succProb[i]);
            nodes[edges[i][1]].neighbors.add(edges[i][0]);
            nodes[edges[i][1]].probs.add(succProb[i]);
        }

        double[] probs = new double[n];
        for (int i = 0; i < n; i++) {
            probs[i] = 0;
        }
        probs[start] = 1d;

        Queue<Integer> iq = new LinkedList<Integer>();
        iq.offer(start);

        while(!iq.isEmpty()) {
            Integer index = iq.poll();

            for (int i = 0; i < nodes[index].neighbors.size(); i++) {
                int idx = nodes[index].neighbors.get(i);
                double prob = nodes[index].probs.get(i);
                if (probs[idx] < probs[index] * prob) {
                    probs[idx] = probs[index] * prob;
                    iq.offer(idx);
                }
            }
        }

        return probs[end];
    }
}

