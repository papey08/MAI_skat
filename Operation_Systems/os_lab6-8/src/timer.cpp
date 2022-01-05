#include "timer.h"

void Timer::start_timer()
{
    is_timer_started = true;
    start_ = std::chrono::steady_clock::now();
}

void Timer::stop_timer()
{
    if (is_timer_started)
    {
        is_timer_started = false;
        finish_ = std::chrono::steady_clock::now();
    }
}

int Timer::get_time()
{
    if (is_timer_started)
    {
        finish_ = std::chrono::steady_clock::now();
    }
    return std::chrono::duration_cast<std::chrono::milliseconds>(finish_ - start_).count();
}
