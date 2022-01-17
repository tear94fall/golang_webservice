package member

import (
	"errors"
	"main/database"
)

type MemberInfo struct {
	Id       string
	Password string
	Name     string
	Tel      string
}

func NewMember(register *MemberInfo) *database.Member {
	member := &database.Member{}
	member.UserId = register.Id
	member.Password = register.Password
	member.Name = register.Name
	member.Tel = register.Tel

	return member
}

func UpdateMember(origin *database.Member, update *MemberInfo) {
	if len(update.Id) != 0 && origin.UserId != update.Id {
		origin.UserId = update.Id
	} else if len(update.Password) != 0 && origin.Password != update.Password {
		origin.Password = update.Password
	} else if len(update.Name) != 0 && origin.Name != update.Name {
		origin.Name = update.Name
	} else if len(update.Tel) != 0 && origin.Tel != update.Tel {
		origin.Tel = update.Tel
	}
}

func CheckMemberValue(member *MemberInfo) error {
	var err error

	if len(member.Id) == 0 {
		err = errors.New("아이디를 입력해주세요")
	} else if len(member.Password) == 0 {
		err = errors.New("비밀번호를 입력해 주세요")
	} else if len(member.Name) == 0 {
		err = errors.New("이름을 입력해 주세요")
	} else if len(member.Tel) == 0 {
		err = errors.New("전화번호를 입력해 주세요")
	}

	return err
}

func CheckChangeMember(origin *database.Member, update *MemberInfo) bool {
	var flag bool = false
	if origin.UserId != update.Id {
		flag = true
	} else if origin.Password != update.Password {
		flag = true
	} else if origin.Name != update.Name {
		flag = true
	} else if origin.Tel != update.Tel {
		flag = true
	}

	return flag
}
