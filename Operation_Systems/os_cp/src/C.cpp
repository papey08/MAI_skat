#include <iostream>
#include <unistd.h>
#include <fcntl.h>
#include <semaphore.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <stdarg.h>
#include <signal.h>

int human_get(sem_t *semaphore)
{
    int s;
    sem_getvalue(semaphore, &s);
    return s;
}

void human_set(sem_t *semaphore, int n)
{
    while (human_get(semaphore) < n)
    {
        sem_post(semaphore);
    }
    while (human_get(semaphore) > n)
    {
        sem_wait(semaphore);
    }
}

int main(int args, char* argv[])
{
    int fdAC[2];
    fdAC[0] = atoi(argv[0]);
    fdAC[1] = atoi(argv[1]);
    int fdBC[2];
    fdBC[0] = atoi(argv[2]);
    fdBC[1] = atoi(argv[3]);
    sem_t* semA = sem_open("_semA", O_CREAT, 0777, 1);
    sem_t* semB = sem_open("_semB", O_CREAT, 0777, 0);
    sem_t* semC = sem_open("_semC", O_CREAT, 0777, 0);
    while(1)
    {
        while(human_get(semC) == 0)
        {
            continue;
        }
        if (human_get(semC) == 2)
        {
            break;
        }
        int size;
        std::string str;
        read(fdAC[0], &size, sizeof(int));
        int t = 0;
        for (int i = 0; i < size; ++i)
        {
            char c;
            read(fdAC[0], &c, sizeof(char));
            str.push_back(c);
            t = i;
        }
        ++t;
        std::cout << str << std::endl;
        write(fdBC[1], &t, sizeof(int));
        human_set(semB, 1);
        human_set(semC, 0);
    }
    sem_close(semA);
    sem_close(semB);
    sem_close(semC);
    close(fdAC[0]);
    close(fdAC[1]);
    close(fdBC[0]);
    close(fdBC[1]);
    return 0;
}
