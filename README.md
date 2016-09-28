# Genetic Chess

genetic-chess is a self-improving genetic chess engine written in Go.

## Install

genetic-chess requires Go 1.7.1 or later.

```
$ go get -u github.com/clex/genetic-chess
```

## Usage

All arguments are optional.

```
Usage of genetic-chess:
  -children uint
    	number of children for each qualified (default 2)
  -file string
    	data file, created if necessary (default "/tmp/genetic-chess-phenotype.json")
  -games uint
    	number of games to play against each other (must be a multiple of 2 to keep white/black even) (default 2)
  -max-depth uint
    	maximum allowed depth in games tree (default 3)
  -mutation-size float
    	maximum mutation size for a gene between two generations (default 0.25)
  -mutations uint
    	number of mutations between two generations (default 1)
  -parallel-games uint
    	number of parallel games (each game uses a go routine) (default 1)
  -play
    	play against ai
  -qualified uint
    	number of ai qualified for the next tournament (default 1)
  -quiet
    	disables all output
  -rounds uint
    	number of rounds (0 means infinite)
  -time-to-think duration
    	maximum time to think for a move (suffix with "ms", "s", "m" or "h" (default 10ms)
```

### Self-improving mode

genetic-chess will run endlessly trying to improve. You can specify a number of
rounds (-rounds options) or stop it at any moment with ^C.

```
$ genetic-chess --file ./phenotype.json --rounds 2
Starting new tournament
Playing game 12 of 12...

Tournament finished in 32.036142239s (2.669678519s/game)
Results:
G1-YgyiZGto: 5.00/8
G0-ZGto0CN6: 3.50/8
G1-bWC7ZGto: 3.50/8

Qualified for next round: G1-YgyiZGto
Ex-champion G0-ZGto0CN6 is replaced by G1-YgyiZGto in phenotype.json

Starting new tournament
Playing game 12 of 12...

Tournament finished in 51.371172656s (4.280931054s/game)
Results:
G2-m2hAYgyi: 5.50/8
G2-vHiDYgyi: 4.00/8
G1-YgyiZGto: 2.50/8

Winner: G2-m2hAYgyi
Ex-champion G1-YgyiZGto is replaced by G2-m2hAYgyi in phenotype.json
```

The best phenotype will be kept in the file defined by the -file parameter.
The file format is JSON, it contains the genes values.

If the file is empty a random phenotype will be generated.

### Playing mode

You can play against the AI with the --play option.

```
$ genetic-chess --file ./phenotype.json --play
|R|N|B|Q|K|B|N|R|
|P|P|P|P|P|P|P|P|
| | | | | | | | |
| | | | | | | | |
| | | | | | | | |
| | | | | | | | |
|p|p|p|p|p|p|p|p|
|r|n|b|q|k|b|n|r|

Move:
```

Pieces are described by chars:
* p: pawn
* n: knight
* b: bishop
* r: rook
* q: queen
* k: king

Lowercase pieces are white and uppercase pieces are black.

The move format is the following: &lt;origin&gt;:&lt;destination&gt;[promotion].

The squares are numeroted according to this table:

|        | A   | B   | C   | D   | E   | F   | G   | H   |
| ------ | --- | --- | --- | --- | --- | --- | --- | --- |
| **8**  | 00  | 01  | 02  | 03  | 04  | 05  | 06  | 07  |
| **7**  | 08  | 09  | 10  | 11  | 12  | 13  | 14  | 15  |
| **6**  | 16  | 17  | 18  | 19  | 20  | 21  | 22  | 23  |
| **5**  | 24  | 25  | 26  | 27  | 28  | 29  | 30  | 31  |
| **4**  | 32  | 33  | 34  | 35  | 36  | 37  | 38  | 39  |
| **3**  | 40  | 41  | 42  | 43  | 44  | 45  | 46  | 47  |
| **2**  | 48  | 49  | 50  | 51  | 52  | 53  | 54  | 55  |
| **1**  | 56  | 57  | 58  | 59  | 60  | 61  | 62  | 63  |

Promotions are the following:
* n: knight
* b: bishop
* r: rook
* q: queen

Examples of moves:

```
$ genetic-chess --file ./phenotype.json --play
|R|N|B|Q|K|B|N|R|
|P|P|P|P|P|P|P|P|
| | | | | | | | |
| | | | | | | | |
| | | | | | | | |
| | | | | | | | |
|p|p|p|p|p|p|p|p|
|r|n|b|q|k|b|n|r|

Move: 52:36

|R|N|B|Q|K|B|N|R|
|P|P|P|P|P|P|P|P|
| | | | | | | | |
| | | | | | | | |
| | | | |p| | | |
| | | | | | | | |
|p|p|p|p| |p|p|p|
|r|n|b|q|k|b|n|r|

looking for best move
best move found: 12:28

|R|N|B|Q|K|B|N|R|
|P|P|P|P| |P|P|P|
| | | | | | | | |
| | | | |P| | | |
| | | | |p| | | |
| | | | | | | | |
|p|p|p|p| |p|p|p|
|r|n|b|q|k|b|n|r|

Move: 57:42

|R|N|B|Q|K|B|N|R|
|P|P|P|P| |P|P|P|
| | | | | | | | |
| | | | |P| | | |
| | | | |p| | | |
| | |n| | | | | |
|p|p|p|p| |p|p|p|
|r| |b|q|k|b|n|r|

looking for best move
best move 6:21

|R|N|B|Q|K|B| |R|
|P|P|P|P| |P|P|P|
| | | | | |N| | |
| | | | |P| | | |
| | | | |p| | | |
| | |n| | | | | |
|p|p|p|p| |p|p|p|
|r| |b|q|k|b|n|r|

Move:
```

Example of a promotion:

```
| | | | | | | | |
| | | | |P| |k|p|
| | | | | |P| | |
| | | | |K|p|P| |
| | | | | | |p| |
| | | | | | | | |
| | | | | | | | |
| | | | | | | | |

Move: 15:7q

| | | | | | | |q|
| | | | |P| |k| |
| | | | | |P| | |
| | | | |K|p|P| |
| | | | | | |p| |
| | | | | | | | |
| | | | | | | | |
| | | | | | | | |
```

Algebraic notation is not supported yet.

If you feel the AI is too weak for you, let it self-improve a little bit more.

## Algorithm

### Genes

A phenotype has the following genes:

| Name                | Description                                      | Minimum | Maximum |
| ------------------- | ------------------------------------------------ | ------- | ------- |
| PieceValuePawn      | Value of a pawn                                  | 0.1     | 10.0    |
| PieceValueKnight    | Value of a knight                                | 0.1     | 10.0    |
| PieceValueBishop    | Value of a bishop                                | 0.1     | 10.0    |
| PieceValueRook      | Value of a rook                                  | 0.1     | 10.0    |
| PieceValueQueen     | Value of a queen                                 | 0.1     | 10.0    |
| PiecePositionPawn   | Pawn position factor                             | 0.0     | 1.0     |
| PiecePositionKnight | Knight position factor                           | 0.0     | 1.0     |
| PiecePositionBishop | Bishop position factor                           | 0.0     | 1.0     |
| PiecePositionRook   | Rook position factor                             | 0.0     | 1.0     |
| PiecePositionQueen  | Queen position factor                            | 0.0     | 1.0     |
| PiecePositionKing   | King position factor                             | 0.0     | 1.0     |
| NbMovesFactor       | Number of moves factor                           | 0.0     | 1.0     |
| EndgameNbPieces     | Number of pieces that makes the board an endgame | 2.0     | 16.0    |
| PruneRatio          | Ratio of branches being pruned by alpha-beta     | 0.0     | 0.99    |
| MinKeptNodes        | Minimal number of nodes kept by alpha-beta       | 2.0     | 16.0    |

Genes are floating numbers. When a decimal makes no sense, the value is
rounded.

### Tournament

Each round is actually a tournament between all the phenotypes. It determines
a winner (that will replace the phenotype file) and the qualified individuals
that will participate in the next tournament.

In a tournament each player will play against each other a number of times
defined by the -games parameter.

#### Players

The number of qualified phenotypes for a tournament is defined by the
-qualified parameter. Each qualified phenotypes will generate a number of
children defined by the -children parameter.

Therefore, the number of players participating in the tournament is equal to:

```
qualified * (children + 1)
```

#### Mutations

Each qualified child will get a number of mutations defined by the -mutations
parameter.

Example:
* Current gene value: 3.14
* Minimum gene value: 0
* Maximum gene value: 10
* -mutation-size parameter: 0.2

This makes the value being able to vary from -2.0 to 2.0. The child value
for this gene will therefore be somewhere between 1.14 and 5.14.

Mutation are random among the possible values.

#### Names

Tournament players have names, like "G42-r5Igw77P".

The first part "G42" is the generation number of the player.
The last part "r5Igw77P" is the clone identifier. The first 4 characters
("r5Ig") are random and the last 4 characters ("w77P") are the first 4
characters of its parent.

### Engine

The chess engine itself is a classical minmax with an alpha-beta pruning.
Genes can affect both the alpha-beta pruning strategy (PruneRatio and
MinKeptNodes) and the evaluation function (all the remaining genes).

The time to think for a move is both limited by the -time-to-think parameter
and the -max-depth parameter.
