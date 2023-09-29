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
    int fdAB[2];
    fdAB[0] = atoi(argv[2]);
    fdAB[1] = atoi(argv[3]);
    sem_t* semA = sem_open("_semA", O_CREAT, 0777, 1);
    sem_t* semB = sem_open("_semB", O_CREAT, 0777, 0);
    sem_t* semC = sem_open("_semC", O_CREAT, 0777, 0);
    while(1)
    {
        std::string str;
        getline(std::cin, str);
        if (str == "END")
        {
            human_set(semA, 2);
            human_set(semB, 2);
            human_set(semC, 2);
            break;
        }
        int size = str.length();
        write(fdAC[1], &size, sizeof(int));
        write(fdAB[1], &size, sizeof(int));
        for (int i = 0; i < size; ++i)
        {
            write(fdAC[1], &str[i], sizeof(char));
        }
        human_set(semB, 1);
        human_set(semA, 0);
        while (human_get(semA) == 0)
        {
            continue;
        }
    }
    sem_close(semA);
    sem_destroy(semA);
    sem_close(semB);
    sem_destroy(semB);
    sem_close(semC);
    sem_destroy(semC);
    close(fdAC[0]);
    close(fdAC[1]);
    close(fdAB[0]);
    close(fdAB[1]);
    return 0;
}
