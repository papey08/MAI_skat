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
    int fdAB[2];
    fdAB[0] = atoi(argv[0]);
    fdAB[1] = atoi(argv[1]);
    int fdBC[2];
    fdBC[0] = atoi(argv[2]);
    fdBC[1] = atoi(argv[3]);
    sem_t* semA = sem_open("_semA", O_CREAT, 0777, 1);
    sem_t* semB = sem_open("_semB", O_CREAT, 0777, 0);
    sem_t* semC = sem_open("_semC", O_CREAT, 0777, 0);
    while (1)
    {
        while(human_get(semB) == 0)
        {
            continue;
        }
        if (human_get(semB) == 2)
        {
            break;
        }
        int size;
        read(fdAB[0], &size, sizeof(int));
        std::cout << "Number of input symbols is " << size << std::endl;
        human_set(semC, 1);
        human_set(semB, 0);
        while (human_get(semB) == 0)
        {
            continue;
        }
        if (human_get(semB) == 2)
        {
            break;
        }
        read(fdBC[0], &size, sizeof(int));
        std::cout << "Number of output symbols is " << size << std::endl;
        human_set(semA, 1);
        human_set(semB, 0);
        while(human_get(semB) == 0)
        {
            continue;
        }
        if (human_get(semB) == 2)
        {
            break;
        }
    }
    sem_close(semA);
    sem_close(semB);
    sem_close(semC);
    close(fdAB[0]);
    close(fdAB[1]);
    close(fdBC[0]);
    close(fdBC[1]);
    return 0;
}
