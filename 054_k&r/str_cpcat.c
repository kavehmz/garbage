#include <stdio.h>
#include <ctype.h>


void strcat_custom(char *s, char *t);
void strcp_custom(char *s, char *t);

int main() {
    char *s="abcd";
    char t[20];
    strcp_custom(t,"start: ");
    strcat_custom(t,s);
    strcat_custom(t,"efg");
    printf("%s", t);
}

void strcp_custom(char *t, char *s) {
    while (*t++=*s++);
}

void strcat_custom(char *t, char *s) {
    while (*t!='\0')
        t++;
    while (*t++=*s++);
}