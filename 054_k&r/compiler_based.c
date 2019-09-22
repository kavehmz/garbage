#include <stdio.h>

int main() {

    char a[100];
    int i=0;

    for (int i = 0; i < 100; i++)
    {
        a[i]=0;
    }
    
    a[ ++i +  ++i]=   ++i + ++i;

    printf("[%d],", i);
    for (int i = 0; i < 30; i++)
    {
        printf("%d,", a[i]);

    }
    printf("\n");
    return 0;
}
