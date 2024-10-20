#include <time.h>
#include <sys/time.h>

double get_time(struct timeval *tstart, struct timeval *tend) {
    return (tend->tv_sec - tstart->tv_sec) + (tend->tv_usec - tstart->tv_usec) / 1000000.0;
}