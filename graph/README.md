## 回溯模版
```dotnetcli

void backtrack(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素) {
        处理节点;
        backtrack(路径，选择);
        回溯,撤销处理结果
    }
}
```



## DFS模版

```
void dfs(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本节点所连接的其他节点) {
        处理节点;
        dfs(图，选择的节点); // 递归
        回溯，撤销处理结果
    }
}
```
