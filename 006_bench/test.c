#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <time.h>

struct timespec timer_start(){
    struct timespec start_time;
    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &start_time);
    return start_time;
}

long timer_end(struct timespec start_time){
    struct timespec end_time;
    clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &end_time);
    long diffInNanos = end_time.tv_nsec - start_time.tv_nsec;
    return diffInNanos;
}


double const sqrtE=2.506628274631;
double pdf(double x, double m, double s) {
	double z = (x - m) / s;
	return exp(z*z) / (sqrtE * s);
}


int main ()
{
	struct timespec vartime = timer_start();
	for(double a = 0; a < 10000000; a++ ){
		exp(a);
	}
	long time_elapsed_nanos = timer_end(vartime);
	printf("exp taken (nanoseconds): %ld\n", time_elapsed_nanos);

	vartime = timer_start();
	srand(time(NULL));
	double x = (double)rand();
	double m = (double)rand();
	double s = (double)rand();
	int n=0;
	for(double i = 1; i < 10000000; i++){
		pdf(x,m,s);
	}	
	time_elapsed_nanos = timer_end(vartime);
	printf("pdf taken (nanoseconds): %ld\n", time_elapsed_nanos);
	return(0);
}
