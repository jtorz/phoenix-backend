package fndbiz

import (
	"context"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
)

// pass1 the password is de number 1.
var pass1 = fndmodel.Password{
	Type: fndmodel.PassTypeScrypt2017,
	Data: base.JSONObject{
		"password": "76bd4a1abf79b2ae1a0d90207c9fcd0cf279a3e0ab804717ac41dcee2fbe3beb01e04775478ffa1de1e1912721541647ce8d2faa21503829248acf7854a5d5e64d41653ef837d6bac141abe30f4f67a966f02e89f19bf524deb9ebbe1a47e89148946df9e1a640d5f3b543d3bf875fa673376f3c7ab264f5ddf3e6135fe28d6ee63842ad5d8a42e80f42cf473e0ccae94ad2cfa16138b7698acd5625cf649ad5eeab17c94c74f4e7f675c53c6387e65dc07223a7e5797ee673fe33bb450d336eed333e7aa9d748e34722f087d9e2c3a496becd63c6b51ae5668441258fcb2a506ee31704a797ffae9b1d5d68491cef71eab02278fc5203aa0509c425b43dac6b",
		"salt":     "bd2feff895ac4aa15b2ee47203b75fa079da1a67f200f84d1580d54d8c6505daf8a2c68083c745b21e3db5cfd0639a9e30acc821d87a0951c8696533d7b350b6d1c86e55d6360f20fcf0b580b2ef5251f2f8ac8727dda2fee8cc5e48fb21df050ad700e5a05bbd7a62eb7319edd51040a799154104080e1103c2ccdf2e71e1c5",
	},
}

func newBizUserTest() BizUser {
	return BizUser{dao: &daoUserTest{}}
}

type daoUserTest struct{}

func (dao *daoUserTest) Login(ctx context.Context,
	user string,
) (*fndmodel.User, error) {
	if user == "john.doe@gmail.com" || user == "john.doe" {
		return &fndmodel.User{
			ID:         "0bc01afe-ae7d-4f14-a4a9-8f5a1f4e7e79",
			Name:       "John",
			MiddleName: "",
			LastName:   "Doe",
			Email:      "john.doe@gmail.com",
			Username:   "john.doe",
			Status:     base.StatusActive,
			Password:   &pass1,
		}, nil
	}

	if user == "jane.doe@gmail.com" || user == "jane.doe" {
		return &fndmodel.User{
			ID:         "0bc01afe-ae7d-4f14-a4a9-8f5a1f4e7e79",
			Name:       "John",
			MiddleName: "",
			LastName:   "Doe",
			Email:      "john.doe@gmail.com",
			Username:   "john.doe",
			Status:     base.StatusInactive,
			Password:   &pass1,
		}, nil
	}

	return nil, baseerrors.ErrAuth
}

// GetUserByMail returns a user given its email.
func (dao *daoUserTest) GetUserByMail(ctx context.Context,
	email string,
) (*fndmodel.User, error) {
	return nil, nil
}

// GetUserByID retrives the record information using its ID.
func (dao *daoUserTest) GetUserByID(ctx context.Context,
	userID string,
) (*fndmodel.User, error) {
	return nil, nil
}

func (dao *daoUserTest) New(context.Context, *fndmodel.User) error {
	return nil
}
