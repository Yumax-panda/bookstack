package validator

import (
	"errors"
	"net/url"
	"regexp"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofrs/uuid"

	"bookstack/utils"
	"bookstack/utils/optional"
)

var (
	// ChannelRegex チャンネル名正規表現
	ChannelRegex = regexp.MustCompile(`^[a-zA-Z0-9-_]{1,20}$`)
	// TwitterIDRegex ツイッターIDの正規表現
	TwitterIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{1,15}$`)
	// PKCERegex PKCE文字列の正規表現
	PKCERegex = regexp.MustCompile("^[a-zA-Z0-9~._-]{43,128}$")
	// UserRoleNameRegex ユーザーロール名の正規表現
	UserRoleNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{1,30}$`)
)

// NotInternalURL 内部ネットワーク宛のURLでない
var NotInternalURL = vd.By(func(value interface{}) error {
	var s string
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		if len(s) == 0 {
			return nil
		}
		s = v
	case optional.Of[string]:
		if len(v.ValueOrZero()) == 0 {
			return nil
		}
		s = v.V
	}
	u, _ := url.Parse(s)
	if utils.IsPrivateHost(u.Hostname()) {
		return errors.New("must not be internal url")
	}
	return nil
})

// NotNilUUID uuid.Nilでない
var NotNilUUID = vd.By(func(value interface{}) error {
	switch u := value.(type) {
	case nil:
		return nil
	case uuid.UUID:
		if u == uuid.Nil {
			return errors.New("invalid uuid")
		}
	case optional.Of[uuid.UUID]:
		if u.Valid && u.V == uuid.Nil {
			return errors.New("invalid uuid")
		}
	case string:
		if v := uuid.FromStringOrNil(u); v == uuid.Nil {
			return errors.New("invalid uuid")
		}
	case []byte:
		if v := uuid.FromBytesOrNil(u); v == uuid.Nil {
			return errors.New("invalid uuid")
		}
	default:
		return errors.New("invalid uuid")
	}
	return nil
})
