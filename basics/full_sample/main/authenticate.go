package main

func authenticate(w http.ResponseWriter, r *http.Request) {
  // receives a request body via form url encoded submission
  r.ParseForm()
  // fetch the user from the database and load them into the 
  // User Struct
  user, _ := data.UserByEmail(r.PostFormValue("email"))
  // check that the user object's hashed password i.e. password is 
  // the same as the posted hashed version of the password
  if user.Password == data.Encrypt(r.PostFormValue("password")) {
    // create a new session
    // Session is a struct with a id, uuid, email, UserId, and CreatedAt
    session := user.CreateSession()
    // make a new cookie with a key value i.e. map
    cookie := http.Cookie{
      Name:      "_cookie",
      Value:     session.Uid,
    // http Only allows only http or https to access the cookie i.e. no js
      HttpOnly:  true,
    }
    http.SetCookie(w, &cookie)
    http.Redirect(w, r, "/" 302)
  } else { 
    http.Redirect(w, r, "/login" 302)
  }
}

