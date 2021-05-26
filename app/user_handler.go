package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Spores-Labs/spores-nft-backend/app/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//UserHandler provides function handlers of user
type UserHandler struct {
	DB *gorm.DB
}

type LoginRequest struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

//Attach login function to router
func (uh *UserHandler) Attach(r *mux.Router) {
	r.HandleFunc("/users/login", uh.login).Methods("POST")
	r.HandleFunc("/users/signup", uh.signup).Methods("POST")
	r.Handle("/users/{address}", AuthMiddleware(http.HandlerFunc(uh.update))).Methods("PUT")
}

func (uh *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addr := vars["address"]

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		NewJSendJSONBuilder().
			Code(http.StatusBadRequest).
			Message("Invalid request payload").
			Build().
			Send(w)
		return
	}
	defer r.Body.Close()

	updatedUser, err := user.UpdateUser(uh.DB, addr)
	if err != nil {
		NewJSendJSONBuilder().
			Code(http.StatusInternalServerError).
			Message(err.Error()).
			Build().
			Send(w)
		return
	}
	NewJSendJSONBuilder().
		Code(http.StatusOK).
		Data(updatedUser).
		Message("Updated user successfully").
		Build().
		Send(w)
}

func (uh *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		NewJSendJSONBuilder().
			Code(http.StatusBadRequest).
			Message(err.Error()).
			Build().
			Send(w)
		return
	}

	verified := verifyUserSignature(request)
	if verified {
		// Generate JWT Token
		newToken, err := genertateToken(request.Address)
		if err != nil {
			NewJSendJSONBuilder().
				Code(http.StatusBadRequest).
				Message(err.Error()).
				Build().
				Send(w)
			return
		} else {
			token := models.Token{}
			token.Token = newToken
			_, err := token.SaveNewToken(uh.DB)
			if err != nil {
				NewJSendJSONBuilder().
					Code(http.StatusInternalServerError).
					Message(err.Error()).
					Build().
					Send(w)
				return
			}
		}
		// Check if the user is already exist
		user := models.User{}
		_, err = user.FindUserByAddress(uh.DB, request.Address)
		if err != nil {
			// When user first login, add to database
			user.Address = request.Address
			user.ImageUrl = ""
			user.CreatedTime = time.Now()
			user.UpdatedTime = time.Now()
			_, err := user.SaveNewUser(uh.DB)
			if err != nil {
				NewJSendJSONBuilder().
					Code(http.StatusInternalServerError).
					Message(err.Error()).
					Build().
					Send(w)
				return
			}
		}
		var resp = map[string]interface{}{"token": newToken}
		NewJSendJSONBuilder().
			Code(http.StatusOK).
			Data(resp).
			Message("Logged in successfully").
			Build().
			Send(w)

	} else {
		NewJSendJSONBuilder().
			Code(http.StatusBadRequest).
			Message("Invalid request payload").
			Build().
			Send(w)
		return
	}
}

//Signup ...
func (uh *UserHandler) signup(w http.ResponseWriter, r *http.Request) {

	// Example for sending a error response
	// NewJSendJSONBuilder().
	// 		Code(http.StatusBadRequest).
	// 		Message(err.Error()).
	// 		Build().
	// 		Send(w)

	// Example for sending success response
	// NewJSendJSONBuilder().
	// 	Code(http.StatusOK).
	// 	Build().
	// 	Send(w)
}

func verifyUserSignature(request LoginRequest) bool {
	addr := common.HexToAddress(request.Address)
	sig := hexutil.MustDecode(request.Signature)

	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27

	data := []byte(request.Message)

	publicKey, err := crypto.SigToPub(signHash(data), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*publicKey)
	return addr == recoveredAddr
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
