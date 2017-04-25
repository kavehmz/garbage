#include <stdio.h>
#include <math.h>


struct benchTest {
  long double value;
};

long double bench(long double f) {
	struct benchTest s;
	s.value = sinl(f);
	return s.value;
}

int main() {
    long double tt=1.0;
    for(int i=0;i<10000000;i++) {
        tt=bench(tt);
    }
    printf("%0.19Lf\n", tt);
}