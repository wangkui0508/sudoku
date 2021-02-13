#ifndef SUDOKUGEN_H
#define SUDOKUGEN_H

#ifdef __cplusplus
extern "C" {
#endif

	struct sudoku_puzzle {
		int grid[9][9];
		int solution[9][9];
		int difficulty;
	};
	struct sudoku_puzzle gen_sudoku_puzzle();

#ifdef __cplusplus
}
#endif

#endif

