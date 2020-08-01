#include <unistd.h>
#include <stdlib.h>
#include <sys/sem.h>
#ifndef CGO_SEM_H_
#define CGO_SEM_H_

int cgo_semctl(int sem_id, int sem_num, int cmd,union semun s);

#endif