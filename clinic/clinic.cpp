#include <iostream>
#include <list>
#include <set>
#include <algorithm>
#include <fstream>
#include <queue>

class Clinic;

class Patient
{
public:
    std::string name;
    int64_t base_priority;

    Patient(const std::string& name, int64_t base_priority) : name(name), base_priority(base_priority) {}

    bool operator< (const Patient& other) const {
        if (this->base_priority < other.base_priority)
            return true;
        if (this->base_priority > other.base_priority)
            return false;
        return this->name > other.name;
    }
};

class Clinic
{
    friend class Patient;
    int time_factor;
    int current_time;
    std::priority_queue<Patient> patients;
    std::set<std::string> dead_patients;

    void handle_arrival(const std::string &name, int severity)
    {
        // P = S + K * W = S + K * (current - arrival) = S + K * current - K * arrival
        // P_base = S - K * arrival
        int64_t base_priority = int64_t(severity) - int64_t(this->time_factor) * int64_t(this->current_time);

        this->patients.push(Patient(name, base_priority));
    }

    void handle_call() {
        if (patients.empty()) {
            std::cout << "doctor takes a break" << std::endl;
            return;
        }

        std::string name = patients.top().name;
        patients.pop();

        if (this->dead_patients.count(name)) {
            // Patient is dead, retry
            dead_patients.erase(name);
            return handle_call();
        }

        std::cout << name << std::endl;
    }

    void handle_leave(const std::string& name) {
        this->dead_patients.insert(name);
    }

public:
    void run()
    {
        std::cin.sync_with_stdio(false);
        std::cin.tie(NULL);
        std::istream &is = std::cin;
        //std::ifstream is("test_huge.txt");
        int n;
        is >> n >> this->time_factor;

        for (int i = 0; i < n; i++)
        {
            int q;
            is >> q;

            is >> this->current_time;

            switch (q)
            {
            case 1: {
                int severity;
                std::string name;

                is >> name >> severity;

                this->handle_arrival(name, severity);
                break;
            }
            case 2:
                this->handle_call();
                break;
            case 3: {
                std::string name;

                is >> name;
                this->handle_leave(name);
                break;
            }

            default:
                throw new std::runtime_error("Unknown event " + std::to_string(q) + "!");
            }
        }
    }
};

int main()
{
    Clinic().run();
    return 0;
}