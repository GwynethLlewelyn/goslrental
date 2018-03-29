package main

import (
	"fmt"
	"github.com/cznic/ql"
)

/*
				  .__			 
  _____ _____  |__| ____  
 /		  \\__	 \ |	 |/	\ 
|  Y Y	\/ __ \|	 |		 |	 \
|__|_|	(____  /__|___|	 /
	  \/	 \/			  \/ 
*/

// main() starts here.
func main() {
	db, err := ql.OpenMem()
	if err != nil {
	    panic(err)
	}
	
	rss, _, err := db.Run(ql.NewRWCtx(), `
	    BEGIN TRANSACTION;
	        CREATE TABLE t (i int, s string);
	        INSERT INTO t VALUES
	        	(1, "seafood"),
	        	(2, "A fool on the hill"),
	        	(3, NULL),
	        	(4, "barbaz"),
	        	(5, "foobar"),
	        ;
	    COMMIT;
	    
	    SELECT * FROM t WHERE s LIKE "foo" ORDER BY i;
	    SELECT * FROM t WHERE s LIKE "^bar" ORDER BY i;
	    SELECT * FROM t WHERE s LIKE "bar$" ORDER BY i;
	    SELECT * FROM t WHERE !(s LIKE "foo") ORDER BY i;`,
	)
	if err != nil {
	    panic(err)
	}
	
	for _, rs := range rss {
	    if err := rs.Do(false, func(data []interface{}) (bool, error) {
	        fmt.Println(data)
	        return true, nil
	    }); err != nil {
	        panic(err)
	    }
	    fmt.Println("----")
	}
}