/**
 * Subject: Free Ice Cream
 * URL: https://codeforces.com/problemset/problem/686/A
*/
#include <iostream>
typedef long long ll;

int main() {
    std::ios::sync_with_stdio(0);
	std::cin.tie(0);
	std::cout.tie(0);

    ll n, x;
    std::cin >> n >> x;
    ll dk = 0;
    for (ll i = 0; i < n; i++) {
        char c;
        ll d;
        std::cin >> c >> d;
        if (c == '+') {
            x += d;
        } else {
            if (d > x) dk++;
            else x -= d;
        }
    }
    std::cout << x << " " << dk << "\n";

    return 0;
}