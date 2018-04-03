// Functions to deal with the backoffice.
package main

import (
//	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/heatxsink/go-gravatar"
	"html/template"
//	"io/ioutil"
	"net/http"
//	"strconv"
//	"strings"
)

// GoSLRentalTemplatesType expands on template.Template.
//	need to expand it so I can add a few more methods here 
type GoSLRentalTemplatesType struct{
	template.Template
}

// GoSLRentalTemplates stores all parsed templates for the backoffice.
var GoSLRentalTemplates GoSLRentalTemplatesType

// init parses all templates and puts it inside a (global) var.
//	This is supposed to be called just once! (in func main())
func (gt *GoSLRentalTemplatesType)init(globbedPath string) error {
	temp, err := template.ParseGlob(globbedPath)
	checkErr(err) // move out later, we just need it here to check what's wrong with the templates (20170706)
	//Log.Info("Path is (inside init):", globbedPath)
	gt.Template = *temp;
	return err
}

// GoSLRentalRenderer assembles the correct templates together and executes them.
//	this is mostly to deal with code duplication 
func (gt *GoSLRentalTemplatesType)GoSLRentalRenderer(w http.ResponseWriter, r *http.Request, tplName string, tplParams templateParameters) error {
	thisUserName :=  getUserName(r)
	
	// add cookie to all templates
	tplParams["SetCookie"] = thisUserName

	// Add URLPathPrefix
	tplParams["URLPathPrefix"] = URLPathPrefix
	
	// add Gravatar to templates (note that all logins are supposed to be emails)
	
	// calculate hash for the Gravatar hovercard
	hasher := md5.Sum([]byte(thisUserName))
	hash := hex.EncodeToString(hasher[:])
	tplParams["GravatarHash"] = hash // we ought to cache this somewhere
	
	// deal with sizes, we want to have a specific size for the top menu
	var gravatarSize, gravatarSizeMenu = 32, 32

	// if someone set the sizes, then use them; if not, use defaults
	// note that this required type assertion since tplParams is interface{}
	// see https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
	if tplParams["GravatarSize"] == nil {
		tplParams["GravatarSize"] = gravatarSize
	} else {
		gravatarSize = tplParams["GravatarSize"].(int)
	}
	if tplParams["GravatarSizeMenu"] == nil {
		tplParams["GravatarSizeMenu"] = gravatarSizeMenu
	} else {
		gravatarSizeMenu = tplParams["GravatarSizeMenu"].(int)
	}
	// for Retina displays; we could add a multiplication function for Go templates, but I'm lazy (20170706)
	tplParams["GravatarTwiceSize"] = 2 * gravatarSize
	tplParams["GravatarTwiceSizeMenu"] = 2 * gravatarSizeMenu
	
	// Now call the nice library function to get us the URL to the image, for the two sizes
	g := gravatar.New("identicon", gravatarSize, "g", true)
	tplParams["Gravatar"] = g.GetImageUrl(thisUserName) // we also ought to cache this somewhere
	
	g = gravatar.New("identicon", gravatarSizeMenu, "g", true)
	tplParams["GravatarMenu"] = g.GetImageUrl(thisUserName) // we also ought to cache this somewhere
	
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett") // do a tribute to one of my best fantasy authors (see http://www.gnuterrypratchett.com/) (20170807)
	return gt.ExecuteTemplate(w, tplName, tplParams)
}

// Auxiliary functions for session handling
//	see https://mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout/ (20170603)

var cookieHandler = securecookie.New(		// from gorilla/securecookie
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// setSession returns a new session cookie with an encoded username.
func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:	"session",
			Value: encoded,
			Path:	"/",
		}
		// Log.Debug("Encoded cookie:", cookie)
		http.SetCookie(response, cookie)
	} else {
		Log.Error("Error encoding cookie:", err)
	}
 }
 
 // getUserName sees if we have a session cookie with an encoded user name, returning nil if not found.
func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

// clearSession will remove a cookie by setting its MaxAge to -1 and clearing its value.
func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:	"session",
		Value:	 "",
		Path:	 "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// checkSession will see if we have a valid cookie; if not, redirects to login.
func checkSession(w http.ResponseWriter, r *http.Request) {
	// valid cookie and no errors?
	if getUserName(r) == "" {
		http.Redirect(w, r, URLPathPrefix + "/admin/login/", 302)	
	}
}


// Function handlers for HTTP requests (main functions for this file)

// backofficeMain is the main page, has some minor statistics, may do this fancier later on.
func backofficeMain(w http.ResponseWriter, r *http.Request) {
	checkSession(w, r) // make sure we've got a valid cookie, or else send to login page
	// let's load the main template for now, just to make sure this works
	
	tplParams := templateParameters{ "Title": "GoSLRental Administrator Panel - main",
	}
	err := GoSLRentalTemplates.GoSLRentalRenderer(w, r, "main", tplParams)
	checkErr(err)
	return
}

// backofficeUserManagement deals with adding/removing application users. Just login(email) and password right now, no profiles, no email confirmations, etc. etc. etc.
//	This is basically a stub for more complex user management, to be reused by other developments...
//	I will not develop this further, except perhaps to link usernames to in-world avatars (may be useful)
func backofficeUserManagement(w http.ResponseWriter, r *http.Request) {
	checkSession(w, r)
	tplParams := templateParameters{ "Title": "GoSLRental Administrator Panel - User Management",
			"Content": "Hi there, this is the User Management template",
			"URLPathPrefix": URLPathPrefix,
			"GoSLRentalJS": "user-management.js",
	}
	err := GoSLRentalTemplates.GoSLRentalRenderer(w, r, "user-management", tplParams)
	checkErr(err)
	return
}


// backofficeLogin deals with authentication.
func backofficeLogin(w http.ResponseWriter, r *http.Request) {
	// Log.Debug("Entered backoffice login for URL:", r.URL, "using method:", r.Method)
	if r.Method == "GET" {
		tplParams := templateParameters{ "Title": "GoSLRental Administrator Panel - login",
				"URLPathPrefix": URLPathPrefix,
		}
		err := GoSLRentalTemplates.GoSLRentalRenderer(w, r, "login", tplParams)
		checkErr(err)
	} else { // POST is assumed
		r.ParseForm()
		// logic part of logging in
		email		:= r.Form.Get("email")
		password	:= r.Form.Get("password")
		
		// Log.Debug("email:", email)
		// Log.Debug("password:", password)
		
		if email == "" || password == "" { // should never happen, since the form checks this
			 http.Redirect(w, r, URLPathPrefix + "/", 302)		  
		}
		
		// Check username on database
		db, err := sql.Open(PDO_Prefix, GoSLRentalDSN)
		checkErr(err)
		
		defer db.Close()
	
		// query
		rows, err := db.Query("SELECT Email, Password FROM Users")
		checkErr(err)

		defer rows.Close()
		
		var (
			Email string
			Password string
		)
	 
		// enhash the received password; I just use MD5 for now because there is no backoffice to create
		//  new users, so it's easy to generate passwords manually using md5sum;
		//  however, MD5 is not strong enough for 'real' applications, it's just what we also use to
		//  communicate with the in-world scripts (20170604)
		pwdmd5 := fmt.Sprintf("%x", md5.Sum([]byte(password))) //this has the hash we need to check
	  
		authorised := false // outside of the for loop because of scope
	
		for rows.Next() {	// we ought just to have one entry, but...
			_ = rows.Scan(&Email, &Password)
			// ignore errors for now, either it checks true or any error means no authentication possible
			if Password == pwdmd5 {
				authorised = true
				break
			}		
		}
		
		 if authorised {
			 // we need to set a cookie here
			 setSession(email, w)
			 // redirect to home
			 http.Redirect(w, r, URLPathPrefix + "/admin", 302)
		} else {
			// possibly we ought to give an error and then redirect, but I don't know how to do that (20170604)
			http.Redirect(w, r, URLPathPrefix + "/", 302) // will ask for login again
		}
		return
	}
}

// backofficeLogout clears session and returns to login prompt.
func backofficeLogout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, URLPathPrefix + "/", 302)
}

// backofficeLSLRegisterObject creates a LSL script for registering cubes, using the defaults set by the user.
//	This is better than using 'template' LSL scripts which people may fill in wrongly, this way at least
//	 we won't get errors about wrong signature PIN or hostnames etc.
func backofficeLSLRegisterObject(w http.ResponseWriter, r *http.Request) {
	checkSession(w, r)
	tplParams := templateParameters{ "Title": "GoSLRental LSL Generator - register object.lsl",
			"URLPathPrefix": URLPathPrefix,
			"Host": Host,
			"ServerPort": ServerPort,
			"LSLSignaturePIN": LSLSignaturePIN,
			"LSL": "lsl-register-object", // this will change some formatting on the 'main' template (20170706)
	}
	// check if we have a frontend (it's configured on the config.toml file); if no, use the ServerPort
	//	 the 'frontend' will be nginx, Apache, etc. to cache replies from Go and serve static files from port 80 (20170706)
	if FrontEnd == "" {
		tplParams["ServerPort"] = ServerPort
	}
	err := GoSLRentalTemplates.GoSLRentalRenderer(w, r, "main", tplParams)
	checkErr(err)
	return
}