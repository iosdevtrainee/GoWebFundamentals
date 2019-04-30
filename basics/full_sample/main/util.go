package main 

func session(w http.ResponseWriter, r *http.Request) (sess data.Session,
 err error) {
  // fetch the cookie previous set from the request
  cookie, err := r.Cookie("_cookie")
  if err == nil {
   // fetch the associate session with the cookie
    sess = data.Session{Uuid: cookie.Value}
   // if the session doesn't exist send an error of invalid session
    if ok, _ := sess.Check(); !ok {
      err = errors.New("Invalid session")
    }
  }
  return
}

