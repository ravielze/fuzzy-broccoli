package token

import (
	"net/http"

	"github.com/ravielze/oculi/common/model/dto/user"
)

type (
	Encoder interface {
		Encode(claims Claims) (string, error)
		CreateClaims(credentials user.CredentialsDTO, exp int64) Claims
		CreateAndEncode(credentials user.CredentialsDTO, exp int64) (string, error)
	}

	Decoder interface {
		Decode(token string) (Claims, error)
		DecodeHttpRequest(req *http.Request) (Claims, error)
	}

	Claims interface {
		Credentials() user.CredentialsDTO
		Valid() error
	}

	Tokenizer interface {
		Encoder
		Decoder
	}
)