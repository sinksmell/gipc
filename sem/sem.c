#include <unistd.h>
#include <stdlib.h>
#include <sys/sem.h>
#include "sem.h"

int cgo_semctl(int sem_id, int sem_num, int cmd,union semun s){
	return semctl(sem_id,sem_num,cmd,s);
}

