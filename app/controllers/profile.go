package controllers

import (
	"github.com/revel/revel"
	"PPL2-ITBCareerCenter/app/models"
)

type Profile struct {
	App
}

func (c Profile) List(page int) revel.Result {
	profiles := true
	numUserPerPage := 6
	if (page == 0) {
		page = 1
	}
	startUserLimit := (page-1)*numUserPerPage
	endUserLimit := page*numUserPerPage

	userCount := CountUsers(Dbm)

	startUserLimit = max(startUserLimit, 0)

	if (startUserLimit >= userCount) {
		return c.NotFound("Invalid Page: ", page);
	}

	endUserLimit = min(userCount, endUserLimit)

	users := SelectLatestUsersInRange(Dbm, startUserLimit, endUserLimit - startUserLimit)
	currentPageNum := page
	return c.Render(profiles, page, users, userCount, numUserPerPage, currentPageNum)
}

func (c Profile) Edit(id int, user models.Users) revel.Result {
	oldUser := SelectUsersByUserid(Dbm, id)
	oldUser.CompanyName = user.CompanyName
	oldUser.Name = user.Name
	oldUser.CompanyDescription = user.CompanyDescription
	oldUser.Visi = user.Visi
	oldUser.Misi = user.Misi
	oldUser.Jurusan = user.Jurusan
	oldUser.Angkatan = user.Angkatan
	UpdateUsers(Dbm, oldUser)
	return c.Redirect("/ProfilePage/%d", id)
}

func (c Profile) Form(id int) revel.Result {
	profiles := true
	user := SelectUsersByUserid(Dbm, id)
	return c.Render(user, profiles)
}

func (c Profile) Page(id int) revel.Result {
	profiles := true
	authorized := false
	user := SelectUsersByUserid(Dbm, id)
	namaPerusahaan := user.CompanyName
	deskripsiPerusahaan := user.CompanyDescription
	visiPerusahaan := user.Visi
	misiPerusahaan := user.Misi
	namaPemilik := user.Name
	jurusan := user.Jurusan
	angkatanPMW := user.Angkatan
	userContact := SelectAllUserContactByUserId(Dbm, id)
	userSocialMedia := SelectAllUserSocialMediaByUserID(Dbm, id)
	userVideo := SelectVideoByUserId(Dbm, id)
	userImage := SelectUserImage(Dbm, id)
	return c.Render(id, profiles, namaPerusahaan, deskripsiPerusahaan, visiPerusahaan, misiPerusahaan, namaPemilik, jurusan, angkatanPMW, userContact, userSocialMedia, authorized, userVideo, userImage)
}
