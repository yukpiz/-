#include <stdio.h>
#include <emscripten/emscripten.h>

int main(int argc, char ** argv) {
	printf("(」・ω・)」うー!\n");
	printf("(/・ω・)/にゃー!\n");
}

#ifdef __cplusplus
extern "C" {
#endif

int EMSCRIPTEN_KEEPALIVE Uhh() {
	printf("(」・ω・)」うー!\n");
	return 1;
}

char EMSCRIPTEN_KEEPALIVE *Nya() {
	return "(/・ω・)/にゃー!";
}

int EMSCRIPTEN_KEEPALIVE add(int a, int b) {
	return a + b;
}

#ifdef __cplusplus
}
#endif