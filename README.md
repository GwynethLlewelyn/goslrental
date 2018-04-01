# goslrental
## A Go implementation of a rental system for Second Life and OpenSimulator.

![goslrental logo](apple-touch-icon.png)

It's still under construction. It uses `github.com/cznic/ql` as a simple database implemented in Go (as opposed to SQLite which is buggy with goroutines) but accessing the database is made through 'standard' commands via the `database/sql` package for portability with other database engines.

Licensed under the [https://opensource.org/licenses/BSD-3-Clause](BSD 3-clause license) (you can basically do whatever you wish with this code so long as you keep the attribution).