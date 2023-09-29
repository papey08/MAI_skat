#include <iostream>
#include <unistd.h>
#include <fcntl.h>
#include <semaphore.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <signal.h>
#include <stdarg.h>

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

int main()
{
    int fdAC[2];
    int fdAB[2];
    int fdBC[2];
    pipe(fdAC);
    pipe(fdAB);
    pipe(fdBC);
    sem_unlink("_semA");
    sem_unlink("_semB");
    sem_unlink("_semC");
    sem_t* semA = sem_open("_semA", O_CREAT, 0777, 1);
    sem_t* semB = sem_open("_semB", O_CREAT, 0777, 0);
    sem_t* semC = sem_open("_semC", O_CREAT, 0777, 0);
    if ((semA == SEM_FAILED)||(semB == SEM_FAILED)||(semC == SEM_FAILED))
    {
        perror("sem_open");
        return -1;
    }
    std::cout << "Enter some strings:\n";
    pid_t C = fork();
    if (C == -1)
    {
        perror("fork");
        return -1;
    }
    if (C == 0)
    {
        pid_t B = fork();
        if (B == -1)
        {
            perror("fork");
            return -1;
        }
        if (B == 0)
        {
            execl("B", std::to_string(fdAB[0]).c_str(), std::to_string(fdAB[1]).c_str(), std::to_string(fdBC[0]).c_str(), std::to_string(fdBC[1]).c_str(), NULL);
        }
        else
        {
            execl("C", std::to_string(fdAC[0]).c_str(), std::to_string(fdAC[1]).c_str(), std::to_string(fdBC[0]).c_str(), std::to_string(fdBC[1]).c_str(), NULL);
        }
    }
    else
    {
        execl("A", std::to_string(fdAC[0]).c_str(), std::to_string(fdAC[1]).c_str(), std::to_string(fdAB[0]).c_str(), std::to_string(fdAB[1]).c_str(), NULL);
    }
    return 0;
}
