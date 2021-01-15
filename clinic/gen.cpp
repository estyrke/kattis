#include <random>
#include <fstream>

int main() {
    std::ofstream os("test_huge.txt");

    int n = 200000;
    int k = 10000000;

    os << n << " " << k << std::endl;

    std::random_device rd;  //Will be used to obtain a seed for the random number engine
    std::mt19937 gen(rd()); //Standard mersenne_twister_engine seeded with rd()
    std::uniform_int_distribution<> event(1, 3);
    std::uniform_int_distribution<> severity(0, 10000000);

    int t=0;
    for (int i=0; i < n; i++) {
        int q = event(gen);
        os << q << " " << t;
        t += 50;

        switch (q) {
            case 1:
                os << " Alice" << " " << severity(gen);
                break;
            case 3:
                os << " Alice";
                break;
        }
        os << std::endl;
    }
    return 0;
}