import java.util.*;

public class Solution {
    static List<List<Integer>> tree;
    static long[][] dp;
    static List<Integer> weight;

    public static long findMaximumSum(int tree_nodes, List<Integer> treeFrom, List<Integer> treeTo, List<Integer> weightList) {
        weight = weightList;
        tree = new ArrayList<>();
        for (int i = 0; i <= tree_nodes; i++) tree.add(new ArrayList<>());

        // build adjacency list
        for (int i = 0; i < treeFrom.size(); i++) {
            int u = treeFrom.get(i);
            int v = treeTo.get(i);
            tree.get(u).add(v);
            tree.get(v).add(u);
        }

        dp = new long[tree_nodes + 1][2];
        boolean[] visited = new boolean[tree_nodes + 1];

        dfs(1, -1);

        return Math.max(dp[1][0], dp[1][1]);
    }

    private static void dfs(int u, int parent) {
        dp[u][0] = 0;                       // not taking u
        dp[u][1] = weight.get(u - 1);       // taking u (1-based index for nodes, but weights list is 0-based)

        for (int v : tree.get(u)) {
            if (v == parent) continue;
            dfs(v, u);

            dp[u][0] += Math.max(dp[v][0], dp[v][1]);  // if u not taken, child can be taken or not
            dp[u][1] += dp[v][0];                      // if u taken, child cannot be taken
        }
    }

    public static void main(String[] args) {
        int tree_nodes = 3;
        List<Integer> tree_from = Arrays.asList(1, 1);
        List<Integer> tree_to = Arrays.asList(2, 3);
        List<Integer> weight = Arrays.asList(2, 2, 1);

        System.out.println(findMaximumSum(tree_nodes, tree_from, tree_to, weight)); // Output: 3
    }
}
