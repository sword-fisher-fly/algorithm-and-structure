package hard

// TODO:
// https://leetcode.cn/problems/evaluate-division/solutions/548634/399-chu-fa-qiu-zhi-nan-du-zhong-deng-286-w45d/

// 1) 并查集
// 2) 广度优先算法
// 3) 深度优先算法

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	return nil
}

/*
vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values, vector<vector<string>>& queries) {
        unordered_map<string, unordered_map<string, double>> graph;
        int n = equations.size();
        for (int i = 0; i < n; i++) {
            string u = equations[i][0];
            string v = equations[i][1];
            double weight = values[i];
            graph[u][v] = weight;
            graph[v][u] = 1.0 / weight;
        }
        int m = queries.size();
        vector<double> results(m, -1.0);
        for (int i = 0; i < m; i++) {
            string start = queries[i][0];
            string target = queries[i][1];
            if (!graph.count(start) || !graph.count(target)) {
                continue;
            }
            if (start == target) {
                results[i] = 1.0;
                continue;
            }
            queue<pair<string, double>> q;
            unordered_map<string, bool> visited;
            q.push({start, 1.0});
            visited[start] = true;
            while (!q.empty()) {
                string curr = q.front().first;
                double currVal = q.front().second;
                q.pop();
                if (curr == target) {
                    results[i] = currVal;
                    break;
                }
                for (auto neighbor : graph[curr]) {
                    string next = neighbor.first;
                    double nextVal = neighbor.second;
                    if (!visited[next]) {
                        visited[next] = true;
                        q.push({next, currVal * nextVal});
                    }
                }
            }
        }
        return results;
    }
*/
