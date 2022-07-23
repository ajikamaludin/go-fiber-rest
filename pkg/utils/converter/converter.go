package converter

import "github.com/ajikamaludin/go-fiber-rest/app/models"

func MapNoteToNoteRes(notes []models.Note) (noteRes []models.NoteRes) {
	for _, v := range notes {
		noteRes = append(noteRes, *v.ToNoteRes())
	}
	return
}

func MapUserToUserRes(users []models.User) (userRes []models.UserRes) {
	for _, v := range users {
		userRes = append(userRes, *v.ToUserRes())
	}
	return
}
