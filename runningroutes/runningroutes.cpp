#include <iostream>
#include <vector>
#include <map>

std::map<std::pair<int, int>, int> solve_cache;
std::vector<std::vector<int>> adj;

int solve(int l, int r)
{
    if (l >= r)
        return 0;

    if (r - l == 1)
        return adj[l][r];

    if (solve_cache.count(std::make_pair(l, r)))
    {
        return solve_cache[std::make_pair(l, r)];
    }

    int res = 0;

    for (int k = l; k <= r; k++)
    {
        res = std::max(res, solve(l + 1, k - 1) + solve(k + 1, r) + adj[l][k]);
    }

    solve_cache[std::make_pair(l, r)] = res;
    return res;
}

int main()
{
    int n;
    std::cin >> n;

    adj.resize(n);

    for (int r = 0; r < n; ++r)
    {
        adj[r].resize(n);
        for (int c = 0; c < n; ++c)
        {
            int adjbit;
            std::cin >> adjbit;
            adj[r][c] = adjbit;
        }
    }

    std::cout << solve(0, n - 1) << std::endl;

    return 0;
}