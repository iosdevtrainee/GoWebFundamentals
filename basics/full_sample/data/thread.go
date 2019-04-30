
import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	// Query the database and get a result set or an err
	rows, err := Db.Query(`Select id, uuid, topic, user_id, created_at
	 FROM threads ORDER BY created_at DESC`)
	if err != nil {
		return
	}
	// iterate through the result set by calling next and pass the values
	// at each row into the empty Thread object via address.
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		// append the current thread to the thread slice i.e. vectorized array
		// expandable and doesn't need fixed size.
		threads = append(threads, th)
	}
	rows.Close()
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("Select count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}

	rows.Close()
	return

}