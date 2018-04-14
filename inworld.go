// This deals with calls coming from Second Life or OpenSimulator.
// it's essentially a RESTful thingy
package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
//	"github.com/cznic/ql"
	"net/http"
//	"strconv"
	"strings"
)

// GetMD5Hash takes a string which is to be encoded using MD5 and returns a string with the hex-encoded MD5 sum.
// Got this from https://gist.github.com/sergiotapia/8263278
func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

// registerObject saves a HTTP URL for a single rental object, making it persistent.
// POST parameters:
//  permURL: a permanent URL from llHTTPServer 
//  signature: to make spoofing harder
//  timestamp: in-world timestamp retrieved with llGetTimestamp()
//  request: currently only delete (to remove entry from database when the rental object is deleted)
func registerObject(w http.ResponseWriter, r *http.Request) {
	// get all parameters in array
	err := r.ParseForm()
	checkErrPanicHTTP(w, http.StatusServiceUnavailable, funcName() + ": Extracting parameters failed:", err)
	
	if r.Header.Get("X-Secondlife-Object-Key") == "" {
		// Log.Debugf("Got '%s'\n", r.Header["X-Secondlife-Object-Key"])
		logErrHTTP(w, http.StatusForbidden, funcName() + ": Only in-world requests allowed.")
		return		
	}
	
	if r.Form.Get("signature") == "" {
		logErrHTTP(w, http.StatusForbidden, funcName() + ": Signature not found") 
		return
	}
	
	signature := GetMD5Hash(r.Header.Get("X-Secondlife-Object-Key") + r.Form.Get("timestamp") + ":" + LSLSignaturePIN)
						
	if signature != r.Form.Get("signature") {
		logErrHTTP(w, http.StatusForbidden, funcName() + ": Signature does not match - hack attempt?")
		return
	}
	
	// open database connection and see if we can update the inventory for this object
	db, err := sql.Open(PDO_Prefix, GoSLRentalDSN)
	checkErrPanicHTTP(w, http.StatusServiceUnavailable, funcName() + ": Connect failed:", err)
	
	defer db.Close()

	if r.Form.Get("permURL") != "" { // object registration
		// Try to update first; if it fails, insert record.
		// This is sadly the only way to do it with the SQL that ql supports, it should however work on any database
		
		tx, err := db.Begin()
		checkErrPanicHTTP(w, http.StatusServiceUnavailable, "Transaction begin failed: %s\n", err)
		
		defer tx.Commit()
		
		stmt, err := tx.Prepare("UPDATE Objects SET Name=?2, OwnerKey=?3, OwnerName=?4, PermURL=?5, Location=?6, Position=?7, Rotation=?8, Velocity=?9, LastUpdate=?10) WHERE UUID=?1");
		checkErrPanicHTTP(w, http.StatusServiceUnavailable, "Update prepare failed: %s\n", err)

		defer stmt.Close()
		
		_, err = stmt.Exec(		
			r.Header.Get("X-Secondlife-Object-Key"),
			r.Header.Get("X-Secondlife-Object-Name"),
			r.Header.Get("X-Secondlife-Owner-Key"),
			r.Header.Get("X-Secondlife-Owner-Name"),
			r.Form.Get("permURL"),
			r.Header.Get("X-Secondlife-Region"),
			strings.Trim(r.Header.Get("X-Secondlife-Local-Position"), "<>()"),
			strings.Trim(r.Header.Get("X-Secondlife-Local-Rotation"), "<>()"),
			strings.Trim(r.Header.Get("X-Secondlife-Local-Velocity"), "<>()"),
			r.Form.Get("timestamp"),
		)
		if (err != nil) {
			// Update failed, means it's a new object, insert it instead
			stmt, err := tx.Prepare("INSERT INTO Objects (UUID, Name, OwnerKey, OwnerName, PermURL, Location, Position, Rotation, Velocity, LastUpdate) VALUES (?1,?2,?3,?4,?5,?6,?7,?8,?9,?10)");
			checkErrPanicHTTP(w, http.StatusServiceUnavailable, "Insert prepare failed: %s\n", err)
	
			_, err = stmt.Exec(
				r.Header.Get("X-Secondlife-Object-Key"),
				r.Header.Get("X-Secondlife-Object-Name"),
				r.Header.Get("X-Secondlife-Owner-Key"),
				r.Header.Get("X-Secondlife-Owner-Name"),
				r.Form.Get("permURL"),
				r.Header.Get("X-Secondlife-Region"),
				strings.Trim(r.Header.Get("X-Secondlife-Local-Position"), "<>()"),
				strings.Trim(r.Header.Get("X-Secondlife-Local-Rotation"), "<>()"),
				strings.Trim(r.Header.Get("X-Secondlife-Local-Velocity"), "<>()"),
				r.Form.Get("timestamp"),
			)
			checkErrPanicHTTP(w, http.StatusServiceUnavailable, funcName() + ": Insert exec failed:", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		replyText := "'" + r.Header.Get("X-Secondlife-Object-Name") +
				"' successfully updated object '" +
				r.Header.Get("X-Secondlife-Owner-Name") + "' (" +
				r.Header.Get("X-Secondlife-Owner-Key") + ")."
		
		fmt.Fprint(w, replyText)
		// log.Printf(replyText) // debug
	} else if r.Form.Get("request") == "delete" { // other requests, currently only deletion		
		stmt, err := db.Prepare("DELETE FROM Objects WHERE UUID=?")
		checkErrPanicHTTP(w, http.StatusServiceUnavailable, funcName() + ": Delete object prepare failed:", err)

		defer stmt.Close()
		
		_, err = stmt.Exec(r.Header.Get("X-Secondlife-Object-Key"))
		checkErrPanicHTTP(w, http.StatusServiceUnavailable, funcName() + ": Delete object exec failed:", err)
		
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "'%s' (%s) successfully deleted.", r.Header.Get("X-Secondlife-Object-Name"), r.Header.Get("X-Secondlife-Object-Key"))
		return
	}
}