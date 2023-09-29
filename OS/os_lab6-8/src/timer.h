#pragma once
#include <chrono>

class Timer
{
public:
    Timer() = default;
    ~Timer() = default;
    void start_timer();
    void stop_timer();
    int get_time();
private:
    bool is_timer_started = false;
    std::chrono::steady_clock::time_point start_;
    std::chrono::steady_clock::time_point finish_;
};
